package app

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
)

// captureLogger returns a logger that writes to buf, so tests can assert
// on the exact WARN messages emitted.
func captureLogger() (*slog.Logger, *bytes.Buffer) {
	var buf bytes.Buffer
	return slog.New(slog.NewTextHandler(&buf, nil)), &buf
}

func TestCheckBINDataFreshness_Fresh_NoWarning(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	recent := time.Now().Add(-1 * 24 * time.Hour) // 1 day old
	mock.ExpectQuery("SELECT MAX\\(last_updated\\) FROM bin_data").
		WillReturnRows(pgxmock.NewRows([]string{"max"}).AddRow(&recent))

	logger, buf := captureLogger()
	checkBINDataFreshness(context.Background(), mock, logger)

	require.Empty(t, buf.String(), "no debería loguear WARN para datos frescos")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckBINDataFreshness_Stale_LogsWarning(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	old := time.Now().Add(-7 * 30 * 24 * time.Hour) // ~7 meses
	mock.ExpectQuery("SELECT MAX\\(last_updated\\) FROM bin_data").
		WillReturnRows(pgxmock.NewRows([]string{"max"}).AddRow(&old))

	logger, buf := captureLogger()
	checkBINDataFreshness(context.Background(), mock, logger)

	require.Contains(t, buf.String(), "bin_data is stale")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckBINDataFreshness_EmptyTable_LogsWarning(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mock.ExpectQuery("SELECT MAX\\(last_updated\\) FROM bin_data").
		WillReturnRows(pgxmock.NewRows([]string{"max"}).AddRow(nil))

	logger, buf := captureLogger()
	checkBINDataFreshness(context.Background(), mock, logger)

	require.Contains(t, buf.String(), "bin_data table is empty")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckBINDataFreshness_QueryError_LogsWarningNotFatal(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	mock.ExpectQuery("SELECT MAX\\(last_updated\\) FROM bin_data").
		WillReturnError(errors.New("connection reset"))

	logger, buf := captureLogger()
	checkBINDataFreshness(context.Background(), mock, logger)

	require.Contains(t, buf.String(), "bin_data freshness check failed")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckBINDataFreshness_ExactlySixMonths_Boundary(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	// Un segundo antes del límite -> NO debe advertir
	justUnder := time.Now().Add(-binDataStaleAfter + time.Minute)
	mock.ExpectQuery("SELECT MAX\\(last_updated\\) FROM bin_data").
		WillReturnRows(pgxmock.NewRows([]string{"max"}).AddRow(&justUnder))

	logger, buf := captureLogger()
	checkBINDataFreshness(context.Background(), mock, logger)

	require.Empty(t, buf.String())
	require.NoError(t, mock.ExpectationsWereMet())
}
