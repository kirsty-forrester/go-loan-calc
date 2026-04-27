package output

import (
    "fmt"
    "os"
    "text/tabwriter"
)
import "github.com/kirsty-forrester/loan-calc/loan"

func WriteTable(schedule []loan.Payment) {
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ' , 0)
    defer w.Flush()

    fmt.Fprintln(w, "Month\tPayment\tPrincipal\tInterest\tBalance")

    // Print out amortisation schedule
    for _, p := range schedule {
        fmt.Fprintf(w, "%d\t£%s\t£%s\t£%s\t£%s\n",
            p.Number, p.Payment.StringFixed(2), p.Principal.StringFixed(2),
            p.Interest.StringFixed(2), p.Balance.StringFixed(2))
    }
}