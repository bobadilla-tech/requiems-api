package mortgage

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate_MonthlyPayment(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Calculate(300000, 6.5, 30)

	// Standard formula: 300k @ 6.5% for 30 years ≈ $1896.20/month
	want := 1896.20
	assert.True(t, math.Abs(result.MonthlyPayment-want) <= 0.01, "expected monthly payment ~%.2f, got %.2f", want, result.MonthlyPayment)
}

func TestCalculate_ScheduleLength(t *testing.T) {
	t.Parallel()
	tests := []struct {
		years int
		want  int
	}{
		{30, 360},
		{15, 180},
		{1, 12},
	}

	svc := NewService()
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			result := svc.Calculate(100000, 5.0, tt.years)
			assert.Len(t, result.Schedule, tt.want)
		})
	}
}

func TestCalculate_FinalBalanceNearZero(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Calculate(300000, 6.5, 30)

	last := result.Schedule[len(result.Schedule)-1]
	assert.True(t, last.Balance <= 1.0, "expected final balance < $1.00, got %.2f", last.Balance)
}

func TestCalculate_TotalsConsistent(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Calculate(200000, 4.0, 15)

	assert.True(t, result.TotalPayment >= result.Principal, "total payment (%.2f) must be >= principal (%.2f)", result.TotalPayment, result.Principal)
	assert.True(t, result.TotalInterest > 0, "expected positive total interest, got %.2f", result.TotalInterest)
	wantInterest := result.TotalPayment - result.Principal
	assert.True(t, math.Abs(result.TotalInterest-wantInterest) <= 1.0, "TotalInterest (%.2f) should equal TotalPayment - Principal (%.2f)", result.TotalInterest, wantInterest)
}

func TestCalculate_FieldsEchoed(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Calculate(250000, 7.25, 20)

	assert.Equal(t, float64(250000), result.Principal)
	assert.Equal(t, 7.25, result.Rate)
	assert.Equal(t, 20, result.Years)
}

func TestCalculate_ScheduleMonthNumbers(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Calculate(100000, 5.0, 1)

	for i, entry := range result.Schedule {
		assert.Equal(t, i+1, entry.Month)
	}
}

func TestService_CalculateBatch(t *testing.T) {
	svc := NewService()
	mortgages := []Request{
		{100000, 5.0, 1},
		{100000, 5.6, 3},
		{100000, 5.5, 5},
	}
	result := svc.CalculateBatch(mortgages)

	if len(result) != 3 {
		t.Errorf("expected 3, got %v", len(result))
	}

	if result[0].MonthlyPayment != 8560.75 {
		t.Errorf("expected MonthlyPayment = 8560.75, got %v", result[0].MonthlyPayment)
	}

	if result[1].MonthlyPayment != 3024.1 {
		t.Errorf("expected MonthlyPayment = 3024.1, got %v", result[1].MonthlyPayment)
	}

	if result[2].MonthlyPayment != 1910.12 {
		t.Errorf("expected MonthlyPayment = 1910.12, got %v", result[2].MonthlyPayment)
	}
}
