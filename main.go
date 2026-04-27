package main

import (
    "flag"
    "fmt"
)
import "github.com/kirsty-forrester/loan-calc/loan"
import "github.com/kirsty-forrester/loan-calc/output"
import "github.com/shopspring/decimal"

func main() {
    var principal, annualrate string
    var term int

    // Check for CSV flag
    csvFlag := flag.Bool("csv", false, "Output schedule as CSV")
    flag.Parse()

    // could potentially use flag here rather than individual input lines but I like the prompt format
    fmt.Println("Please enter the principal (loan amount - in GBP):")
    fmt.Scanln(&principal)

    fmt.Println("Please enter the annual rate (in percentage - e.g. 5):")
    fmt.Scanln(&annualrate)

    fmt.Println("Please enter the term (in months):")
    fmt.Scanln(&term)

    principalDec, _ := decimal.NewFromString(principal)
    annualrateDec, _ := decimal.NewFromString(annualrate)

    l := loan.Loan{
        Principal:  principalDec,
        AnnualRate: annualrateDec,
        TermMonths: term,
    }

    monthlyPayment := l.MonthlyPayment()

    fmt.Printf("Monthly payment £%s\n", monthlyPayment)

    schedule := l.Schedule()

    if *csvFlag {
        output.WriteCsv(schedule)
    } else {
        output.WriteTable(schedule)
    }
}