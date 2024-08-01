package money

import (
	"fmt"
	"testing"
	"time"
)

func TestTransactionGetReq_GetDateRange(t *testing.T) {

	r := TransactionGetReq{TimeRange: "昨天"}
	fmt.Println(time.Unix(1722468840, 0))
	fmt.Println(r.GetDateRange())
}
