package model

import (
	"github.com/jhue/misgo/db/model/clipboard"
	"github.com/jhue/misgo/db/model/money"
	"github.com/jhue/misgo/db/model/record"
	"github.com/jhue/misgo/db/model/user"
	"gorm.io/gorm"
)

type Boot interface {
	Inject(db *gorm.DB) error
}

var models = []any{
	new(user.User),
	new(record.Record),
	new(clipboard.ClipBoard),
	new(money.Transaction),
	new(money.TransactionPersonal),
}

func GetEmptyModels() []any {
	return models
}
