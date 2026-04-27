package output

import (
    "os"
    "fmt"
    "encoding/csv"
)
import "github.com/kirsty-forrester/loan-calc/loan"

func WriteCsv(schedule []loan.Payment) {
    w := csv.NewWriter(os.Stdout)

    w.Write([]string{"Month", "Payment", "Principal", "Interest", "Balance"})

    for _, p := range schedule {
         w.Write([]string{
            fmt.Sprintf("%d", p.Number),
            p.Payment.StringFixed(2),
            p.Principal.StringFixed(2),
            p.Interest.StringFixed(2),
            p.Balance.StringFixed(2),
         })
    }

    w.Flush()
}