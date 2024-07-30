package model

import (
	"github.com/jhue/misgo/db/model/clipboard"
	"github.com/jhue/misgo/db/model/money"
	"github.com/jhue/misgo/db/model/record"
	"github.com/jhue/misgo/db/model/user"
)

var models = []any{
	new(user.User),
	new(record.Record),
	new(clipboard.ClipBoard),
	new(money.Transaction),
}

func GetEmptyModels() []any {
	return models
}
