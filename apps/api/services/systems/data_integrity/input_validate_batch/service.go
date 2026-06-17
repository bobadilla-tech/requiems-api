package inputvalidatebatch

import (
	"context"
	"sync"
	"time"

	inputvalidate "requiems-api/services/systems/data_integrity/input_validate"
)

// Item represents a single contact record to validate within a batch request.
type Item struct {
	Email string `json:"email" validate:"required" normalize:"trim"`
	Phone string `json:"phone" validate:"required" normalize:"trim"`
	Text  string `json:"text" normalize:"trim"`
}

// Request is the input payload for the batch validation endpoint.
type Request struct {
	Items []Item `json:"items" validate:"required,min=1,max=50,dive"`
}

// EmailResult carries the outcome of email validation and normalization.
// Normalized is nil when Valid is false — no canonical form can be produced
// for an invalid address.
type EmailResult struct {
	Valid        bool    `json:"valid"`
	Normalized   *string `json:"normalized"`
	QualityScore float64 `json:"quality_score"`
	Disposable   bool    `json:"disposable"`
}

// PhoneResult carries the outcome of phone validation and normalization.
// Normalized is nil when Valid is false.
type PhoneResult struct {
	Valid        bool    `json:"valid"`
	Normalized   *string `json:"normalized"`
	QualityScore float64 `json:"quality_score"`
}

type TextResult struct {
	IsSafe        bool     `json:"is_safe"`
	ToxicityScore *float64 `json:"toxicity_score"`
	Sentiment     string   `json:"sentiment"`
}

// ItemResult holds the validation outcome for a single item in the batch.
// Index mirrors the item's position in the original request array, preserving
// order regardless of the order in which workers completed.
// Error is non-nil only for processing failures (timeout, panic) — field-level
// failures (bad email format) are expressed via Email.Valid / Phone.Valid instead.
type ItemResult struct {
	Index               int          `json:"index"`
	Email               *EmailResult `json:"email"`
	Phone               *PhoneResult `json:"phone"`
	Text                *TextResult  `json:"text"`
	OverallQualityScore float64      `json:"overall_quality_score"`
	Error               *string      `json:"error"`
}

// BatchResponse is the top-level response envelope for the batch endpoint.
// Results preserves input order via the Index field on each ItemResult.
// AverageQualityScore is computed only over items that processed successfully;
// items with a non-nil Error are excluded from the average.
type BatchResponse struct {
	Results             []ItemResult `json:"results"`
	Total               int          `json:"total"`
	ValidCount          int          `json:"valid_count"`
	InvalidCount        int          `json:"invalid_count"`
	AverageQualityScore float64      `json:"average_quality_score"`
}

// InputValidateService is the interface for the single-item validation logic
type InputValidateService interface {
	Validate(ctx context.Context, emailAddress, phoneNumber, text string) inputvalidate.Response
}

// Service holds the dependencies for batch validation.
type Service struct {
	InputValidateSvc InputValidateService
}

// NewService constructs a Service with its required dependency.
func NewService(inputValidateSvc InputValidateService) *Service {
	return &Service{
		InputValidateSvc: inputValidateSvc,
	}
}

// ValidateBatch runs up to 8 concurrent workers to validate each item in the
// batch against the single-item validation service. Each item is given a
// 2-second timeout; items that exceed it are marked with an error and counted
// as invalid. Results preserve the original request order via the Index field.
func (s *Service) ValidateBatch(ctx context.Context, items []Item) BatchResponse {
	const workers = 8
	const perItemTimeout = 2 * time.Second

	sem := make(chan struct{}, workers)
	results := make([]ItemResult, len(items))

	var (
		mu              sync.Mutex
		wg              sync.WaitGroup
		processedCount  int
		validCount      int
		invalidCount    int
		sumQualityScore float64
	)

	for i, item := range items {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, item Item) {
			defer wg.Done()
			defer func() { <-sem }()

			itemCtx, cancel := context.WithTimeout(ctx, perItemTimeout)
			defer cancel()

			type validateOutcome struct {
				res      inputvalidate.Response
				panicMsg *string
			}
			ch := make(chan validateOutcome, 1)

			go func() {
				defer func() {
					if r := recover(); r != nil {
						msg := "panic: validation worker failed"
						ch <- validateOutcome{panicMsg: &msg}
					}
				}()
				ch <- validateOutcome{res: s.InputValidateSvc.Validate(itemCtx, item.Email, item.Phone, item.Text)}
			}()

			select {
			case out := <-ch:
				if out.panicMsg != nil {
					mu.Lock()
					invalidCount++
					mu.Unlock()
					results[i] = ItemResult{Index: i, Error: out.panicMsg}
					return
				}
				result := out.res

				emailResult := EmailResult{
					Valid:        result.Email.Valid,
					QualityScore: result.Email.QualityScore,
					Disposable:   result.Email.Disposable,
				}

				if result.Email.Normalized != nil {
					emailResult.Normalized = result.Email.Normalized
				}

				phoneResult := PhoneResult{
					Valid:        result.Phone.Valid,
					QualityScore: result.Phone.QualityScore,
				}

				if result.Phone.Normalized != nil {
					phoneResult.Normalized = result.Phone.Normalized
				}

				var textResult *TextResult
				if result.Text != nil {
					textResult = &TextResult{
						IsSafe:        result.Text.IsSafe,
						Sentiment:     result.Text.Sentiment,
						ToxicityScore: result.Text.ToxicityScore,
					}
				}

				mu.Lock()
				processedCount++
				sumQualityScore += result.OverallQualityScore
				if result.Email.Valid && result.Phone.Valid {
					validCount++
				} else {
					invalidCount++
				}
				mu.Unlock()

				results[i] = ItemResult{
					Index:               i,
					Email:               &emailResult,
					Phone:               &phoneResult,
					Text:                textResult,
					OverallQualityScore: result.OverallQualityScore,
				}
			case <-itemCtx.Done():
				mu.Lock()
				invalidCount++
				mu.Unlock()

				errMsg := "timeout: validation took longer than 2 seconds"
				results[i] = ItemResult{
					Index: i,
					Error: &errMsg,
				}
			}
		}(i, item)
	}

	wg.Wait()

	var averageQualityScore float64
	if processedCount > 0 {
		averageQualityScore = sumQualityScore / float64(processedCount)
	}

	return BatchResponse{
		Results:             results,
		ValidCount:          validCount,
		InvalidCount:        invalidCount,
		Total:               len(items),
		AverageQualityScore: averageQualityScore,
	}
}

// UsageCount implements the UsageCounter interface, allowing Handle to
// automatically set the X-Usage-Count response header with the total
// number of items processed in the batch.
func (b BatchResponse) UsageCount() int {
	return b.Total
}
