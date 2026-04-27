package loan

import (
    "testing"
    "github.com/shopspring/decimal"
)

func TestMonthlyRate(t *testing.T) {
    l := Loan{
        Principal:  decimal.NewFromInt(1200),
        AnnualRate: decimal.NewFromInt(12),
        TermMonths: 12,
    }

    got := l.MonthlyRate()
    want := decimal.NewFromFloat(0.01)

    if !got.Equal(want) {
        t.Errorf("MonthlyRate() = %s, want %s", got, want)
    }
}

// TODO: finish tests. Test zero interest rate