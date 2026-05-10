package mortgage

import (
	"math"
	"testing"
)

func TestCalculate_MonthlyPayment(t *testing.T) {
	svc := NewService()
	result := svc.Calculate(300000, 6.5, 30)

	// Standard formula: 300k @ 6.5% for 30 years ≈ $1896.20/month
	want := 1896.20
	if math.Abs(result.MonthlyPayment-want) > 0.01 {
		t.Errorf("expected monthly payment ~%.2f, got %.2f", want, result.MonthlyPayment)
	}
}

func TestCalculate_ScheduleLength(t *testing.T) {
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
			result := svc.Calculate(100000, 5.0, tt.years)
			if len(result.Schedule) != tt.want {
				t.Errorf("years=%d: expected %d schedule entries, got %d", tt.years, tt.want, len(result.Schedule))
			}
		})
	}
}

func TestCalculate_FinalBalanceNearZero(t *testing.T) {
	svc := NewService()
	result := svc.Calculate(300000, 6.5, 30)

	last := result.Schedule[len(result.Schedule)-1]
	if last.Balance > 1.0 {
		t.Errorf("expected final balance < $1.00, got %.2f", last.Balance)
	}
}

func TestCalculate_TotalsConsistent(t *testing.T) {
	svc := NewService()
	result := svc.Calculate(200000, 4.0, 15)

	if result.TotalPayment < result.Principal {
		t.Errorf("total payment (%.2f) must be >= principal (%.2f)", result.TotalPayment, result.Principal)
	}
	if result.TotalInterest <= 0 {
		t.Errorf("expected positive total interest, got %.2f", result.TotalInterest)
	}
	wantInterest := result.TotalPayment - result.Principal
	if math.Abs(result.TotalInterest-wantInterest) > 1.0 {
		t.Errorf("TotalInterest (%.2f) should equal TotalPayment - Principal (%.2f)", result.TotalInterest, wantInterest)
	}
}

func TestCalculate_FieldsEchoed(t *testing.T) {
	svc := NewService()
	result := svc.Calculate(250000, 7.25, 20)

	if result.Principal != 250000 {
		t.Errorf("expected principal 250000, got %v", result.Principal)
	}
	if result.Rate != 7.25 {
		t.Errorf("expected rate 7.25, got %v", result.Rate)
	}
	if result.Years != 20 {
		t.Errorf("expected years 20, got %v", result.Years)
	}
}

func TestCalculate_ScheduleMonthNumbers(t *testing.T) {
	svc := NewService()
	result := svc.Calculate(100000, 5.0, 1)

	for i, entry := range result.Schedule {
		if entry.Month != i+1 {
			t.Errorf("schedule[%d].Month = %d, want %d", i, entry.Month, i+1)
		}
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

	if result.Total != 3 {
		t.Errorf("expected 3, got %v", result.Total)
	}

	if result.Results[0].MonthlyPayment != 8560.75 {
		t.Errorf("expected MonthlyPayment = 8560.75, got %v", result.Results[0].MonthlyPayment)
	}

	if result.Results[1].MonthlyPayment != 3024.1 {
		t.Errorf("expected MonthlyPayment = 3024.1, got %v", result.Results[1].MonthlyPayment)
	}

	if result.Results[2].MonthlyPayment != 1910.12 {
		t.Errorf("expected MonthlyPayment = 1910.12, got %v", result.Results[2].MonthlyPayment)
	}
}
