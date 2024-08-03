package money

import (
	"fmt"
	"github.com/jhue/misgo/db/model/user"
	"gorm.io/gorm"
	"time"
)

var model = Transaction{}

type Transaction struct {
	gorm.Model
	UserID uint `gorm:"index"`
	// 支出 or 收入
	Type string `gorm:"not null"`
	// 用途/来源
	Category string `gorm:"not null"`
	// 金额
	Amount float64 `gorm:"not null"`
	// 备注
	Note string    `gorm:"not null"`
	Time time.Time `gorm:"not null"`
}

func Model() *Transaction {
	return &model
}

func (t *Transaction) String() string {
	formatTime := t.Time.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s %s %s %.2f元 %s", formatTime, t.Category, t.Type, t.Amount, t.Note)
}

// Markdown 表格
//
//	|日期|类型（支出/收入）|用途/来源|金额|备注|
func (t *Transaction) Markdown() string {
	formatTime := t.Time.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("| %s | %s | %s | %.2f | %s | \n",
		formatTime, t.Type, t.Category, t.Amount, t.Note)
}

type TransactionPersonal struct {
	UserID uint `gorm:"primaryKey"`
	// 总共记账
	Count int64
	// 总共收入Count
	IncomeCount int64
	// 总共支出Count
	ExpenditureCount int64
	// 总收入
	Income float64
	// 总支出
	Expenditure float64
	// 剩余金额
	Balance float64
	// 开始记账的时间
	StartTime time.Time
}

func (t *TransactionPersonal) Exist(db *gorm.DB, u user.User) (bool, error) {
	var existing TransactionPersonal
	// 查找现有记录
	err := db.Where("user_id = ?", u.ID).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (t *TransactionPersonal) Update(db *gorm.DB, transaction Transaction) error {
	var existing TransactionPersonal
	// 查找现有记录
	err := db.Where("user_id = ?", transaction.UserID).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果记录不存在，创建一个新记录
			existing = TransactionPersonal{
				UserID:    transaction.UserID,
				StartTime: transaction.Time,
			}
		} else {
			return err
		}
	}

	// 累加字段
	existing.Count++
	if transaction.Type == "收入" {
		existing.IncomeCount++
		existing.Income += transaction.Amount
		existing.Balance += transaction.Amount
	} else if transaction.Type == "支出" {
		existing.ExpenditureCount++
		existing.Expenditure += transaction.Amount
		existing.Balance -= transaction.Amount
	}

	// 保存记录
	if err := db.Save(&existing).Error; err != nil {
		return err
	}
	return nil
}

func (t *TransactionPersonal) UpdateDelete(db *gorm.DB, transaction Transaction) error {
	var existing TransactionPersonal
	// 查找现有记录
	err := db.Where("user_id = ?", transaction.UserID).First(&existing).Error
	if err != nil {
		return err
	}

	// 累减字段
	existing.Count--
	if transaction.Type == "收入" {
		existing.IncomeCount--
		existing.Income -= transaction.Amount
		existing.Balance -= transaction.Amount
	} else if transaction.Type == "支出" {
		existing.ExpenditureCount--
		existing.Expenditure -= transaction.Amount
		existing.Balance += transaction.Amount
	}

	// 保存记录
	if err := db.Save(&existing).Error; err != nil {
		return err
	}
	return nil
}

// Inject 更新每个用户的所有对应值(应该在程序启动时调用)
func (t *TransactionPersonal) Inject(db *gorm.DB) error {
	var users []user.User
	err := db.Find(&users).Error
	if err != nil {
		return err
	}

	for _, u := range users {
		var transactions []Transaction
		err = db.Where("user_id=?", u.ID).Order("time").Find(&transactions).Error
		if err != nil {
			return err
		}
		exist, err := t.Exist(db, u)
		if err != nil {
			return err
		}
		if exist {
			continue
		}
		for _, transaction := range transactions {
			err = t.Update(db, transaction)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type TransactionPersonalDaily struct {
	UserID uint `gorm:"primaryKey"`
	// 总共记账
	Count int64
	// 总共收入Count
	IncomeCount int64
	// 总共支出Count
	ExpenditureCount int64
	// 总收入
	Income float64
	// 总支出
	Expenditure float64
	// 剩余金额
	Balance float64
	// 开始记账的时间
	StartTime time.Time
}
