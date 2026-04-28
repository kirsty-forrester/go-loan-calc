package loan

import "github.com/shopspring/decimal"

type Loan struct {
    Principal   decimal.Decimal
    AnnualRate  decimal.Decimal
    TermMonths  int
}

func (l Loan) MonthlyRate() decimal.Decimal {
    if l.AnnualRate.IsZero() {
        return decimal.Zero
    }

    // monthly rate = annual rate / 12
    return l.AnnualRate.Div(decimal.NewFromInt(100)).Div(decimal.NewFromInt(12))
}

func (l Loan) MonthlyPayment() decimal.Decimal {
    if l.AnnualRate.IsZero() {
        return l.Principal.Div(decimal.NewFromInt(int64(l.TermMonths))).Round(2)
    }

    // M = P × (r(1+r)^n) / ((1+r)^n - 1)
    // M = monthly payment, p = principal, r = monthly interest rate,  n = number of monthly payments - term

    r := l.MonthlyRate()

    // (1 + r)^n
    onePlusR := decimal.NewFromInt(1).Add(r)
    n := decimal.NewFromInt(int64(l.TermMonths))
    power := onePlusR.Pow(n)

    // P * r * (1+r)^n / ((1+r)^n - 1)
    numerator := l.Principal.Mul(r).Mul(power)
    denominator := power.Sub(decimal.NewFromInt(1))
    monthlyPayment := numerator.Div(denominator)

    // Banker's rounding
    return monthlyPayment.Round(2)
}