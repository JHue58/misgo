package money

import (
	"fmt"
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	tra := Transaction{
		Type:     "支出",
		Category: "餐饮",
		Amount:   100.00,
		Note:     "无",
		Time:     time.Now(),
	}

	fmt.Println(tra.String())
	fmt.Println(tra.Markdown())

}
