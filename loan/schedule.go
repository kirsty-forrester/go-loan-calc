package loan

import "github.com/shopspring/decimal"

type Payment struct {
    Number    int
    Payment   decimal.Decimal
    Principal decimal.Decimal
    Interest  decimal.Decimal
    Balance   decimal.Decimal
}

func (l Loan) Schedule() []Payment {
    r := l.MonthlyRate()
    monthly := l.MonthlyPayment()
    balance := l.Principal

    payments := make([]Payment, 0, l.TermMonths)

    // Loop through each month
    for i := 1; i <= l.TermMonths; i++ {
        // Calculate the interest on the remaining balance
        interest := balance.Mul(r).Round(2)

        // Subtract that from the fixed payment to get the principal portion
        principal := monthly.Sub(interest)

        // Last payment - adjust for rounding so balance hits exactly zero
        if i == l.TermMonths {
            principal = balance
            monthly = principal.Add(interest)
        }

        // Reduce the balance
        balance = balance.Sub(principal)

        payments = append(payments, Payment {
            Number: i,
            Payment: monthly,
            Principal: principal,
            Interest: interest,
            Balance: balance,
        })
    }

    return payments
}