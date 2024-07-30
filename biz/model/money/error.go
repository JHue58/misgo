package money

import "fmt"

var (
	ErrAmountError            = fmt.Errorf("金额不能为0")
	ErrCategoryError          = fmt.Errorf("用途/来源 不能为空")
	ErrTypeError              = fmt.Errorf("类型请填写收入或支出")
	ErrTimeError              = fmt.Errorf("时间戳不正确，应该是Unix秒级时间戳")
	ErrStartTimeBeforeEndTime = fmt.Errorf("开始时间不能大于结束时间")
	ErrTransactionError       = fmt.Errorf("transaction字段不正确")
)

func (x *Transaction) Validation() error {
	if x.Amount == 0 {
		return ErrAmountError
	}
	if x.Category == "" {
		return ErrCategoryError
	}
	if x.Type != "收入" && x.Type != "支出" {
		return ErrTypeError
	}
	if x.Time <= 0 {
		return ErrTimeError
	}
	return nil
}
