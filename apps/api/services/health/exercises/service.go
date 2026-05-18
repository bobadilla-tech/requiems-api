package exercises

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/svcerr"
)

// dbRows is the minimal row-iteration interface used by query methods.
// pgx.Rows satisfies this interface in production; mockRows satisfies it in unit tests.
type dbRows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

// dbPool is the minimal database interface required by Service.
// poolWrapper adapts *pgxpool.Pool to this interface in production;
// mockDB satisfies it in unit tests.
type dbPool interface {
	Query(ctx context.Context, sql string, args ...any) (dbRows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// poolWrapper adapts *pgxpool.Pool to the dbPool interface so that NewService
// can accept the concrete pool type while still allowing a test double to be
// injected via the interface.
type poolWrapper struct {
	p *pgxpool.Pool
}

func (pw *poolWrapper) Query(ctx context.Context, sql string, args ...any) (dbRows, error) {
	return pw.p.Query(ctx, sql, args...)
}

func (pw *poolWrapper) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return pw.p.QueryRow(ctx, sql, args...)
}

// Service provides exercise lookups against the exercises PostgreSQL table.
type Service struct {
	db dbPool
}

// NewService creates a new Service backed by the given connection pool.
func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: &poolWrapper{p: db}}
}

// List returns a paginated, optionally-filtered list of exercises.
func (s *Service) List(ctx context.Context, p ListParams) (ExerciseList, error) {
	const base = `
		SELECT
		    id, name, body_parts, equipment, target_muscles, secondary_muscles, instructions,
		    COUNT(*) OVER() AS total
		FROM exercises
		WHERE ($1 = '' OR $1 = ANY(body_parts))
		  AND ($2 = '' OR $2 = ANY(equipment))
		  AND ($3 = '' OR $3 = ANY(target_muscles) OR $3 = ANY(secondary_muscles))
		  AND ($4 = '' OR to_tsvector('english', name) @@ plainto_tsquery('english', $4))
		ORDER BY name
		LIMIT $5 OFFSET $6
	`

	offset := (p.Page - 1) * p.PerPage
	rows, err := s.db.Query(ctx, base,
		p.BodyPart, p.Equipment, p.Muscle, p.Search,
		p.PerPage, offset,
	)
	if err != nil {
		return ExerciseList{}, err
	}
	defer rows.Close()

	var total int
	items := make([]Exercise, 0, p.PerPage)

	for rows.Next() {
		var e Exercise
		if err := rows.Scan(
			&e.ID, &e.Name, &e.BodyParts, &e.Equipment,
			&e.TargetMuscles, &e.SecondaryMuscles, &e.Instructions,
			&total,
		); err != nil {
			return ExerciseList{}, err
		}
		items = append(items, e)
	}
	if err := rows.Err(); err != nil {
		return ExerciseList{}, err
	}

	return ExerciseList{
		Items:   items,
		Total:   total,
		Page:    p.Page,
		PerPage: p.PerPage,
	}, nil
}

// Get returns a single exercise by its internal ID.
func (s *Service) Get(ctx context.Context, id int) (Exercise, error) {
	const q = `
		SELECT id, name, body_parts, equipment, target_muscles, secondary_muscles, instructions
		FROM exercises
		WHERE id = $1
	`

	var e Exercise
	err := s.db.QueryRow(ctx, q, id).Scan(
		&e.ID, &e.Name, &e.BodyParts, &e.Equipment,
		&e.TargetMuscles, &e.SecondaryMuscles, &e.Instructions,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Exercise{}, svcerr.NotFound("not_found", "exercise not found")
		}
		return Exercise{}, err
	}
	return e, nil
}

// Random returns a single exercise chosen at random, respecting any filters.
func (s *Service) Random(ctx context.Context, p ListParams) (Exercise, error) {
	const q = `
		SELECT id, name, body_parts, equipment, target_muscles, secondary_muscles, instructions
		FROM exercises
		WHERE ($1 = '' OR $1 = ANY(body_parts))
		  AND ($2 = '' OR $2 = ANY(equipment))
		  AND ($3 = '' OR $3 = ANY(target_muscles) OR $3 = ANY(secondary_muscles))
		  AND ($4 = '' OR to_tsvector('english', name) @@ plainto_tsquery('english', $4))
		ORDER BY random()
		LIMIT 1
	`

	var e Exercise
	err := s.db.QueryRow(ctx, q, p.BodyPart, p.Equipment, p.Muscle, p.Search).Scan(
		&e.ID, &e.Name, &e.BodyParts, &e.Equipment,
		&e.TargetMuscles, &e.SecondaryMuscles, &e.Instructions,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Exercise{}, svcerr.NotFound("not_found", "no exercises found matching the given filters")
		}
		return Exercise{}, err
	}
	return e, nil
}

// GetBatch returns the exercises matching the given IDs.
// IDs that do not exist are silently skipped.
// Results are ordered to match the input ID slice.
func (s *Service) GetBatch(ctx context.Context, ids []int) ([]Exercise, error) {
	const q = `
		SELECT id, name, body_parts, equipment, target_muscles, secondary_muscles, instructions
		FROM exercises
		WHERE id = ANY($1::int[])
		ORDER BY array_position($1::int[], id)
	`

	rows, err := s.db.Query(ctx, q, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Exercise, 0, len(ids))
	for rows.Next() {
		var e Exercise
		if err := rows.Scan(
			&e.ID, &e.Name, &e.BodyParts, &e.Equipment,
			&e.TargetMuscles, &e.SecondaryMuscles, &e.Instructions,
		); err != nil {
			return nil, err
		}
		items = append(items, e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// BodyParts returns a sorted list of all distinct body part values.
func (s *Service) BodyParts(ctx context.Context) (StringList, error) {
	return s.distinctValues(ctx, `
		SELECT DISTINCT unnest(body_parts) AS val FROM exercises ORDER BY val
	`)
}

// Equipment returns a sorted list of all distinct equipment values.
func (s *Service) Equipment(ctx context.Context) (StringList, error) {
	return s.distinctValues(ctx, `
		SELECT DISTINCT unnest(equipment) AS val FROM exercises ORDER BY val
	`)
}

// Muscles returns a sorted list of all distinct muscle values (target + secondary).
func (s *Service) Muscles(ctx context.Context) (StringList, error) {
	return s.distinctValues(ctx, `
		SELECT DISTINCT unnest(target_muscles || secondary_muscles) AS val FROM exercises ORDER BY val
	`)
}

func (s *Service) distinctValues(ctx context.Context, query string) (StringList, error) {
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return StringList{}, err
	}
	defer rows.Close()

	var items []string
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			return StringList{}, err
		}
		items = append(items, v)
	}
	if err := rows.Err(); err != nil {
		return StringList{}, err
	}

	if items == nil {
		items = []string{}
	}

	return StringList{Items: items, Total: len(items)}, nil
}
