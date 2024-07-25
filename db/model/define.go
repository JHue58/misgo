package model

import (
	"github.com/jhue/misgo/db/model/record"
	"github.com/jhue/misgo/db/model/user"
)

var models = []any{
	new(user.User),
	new(record.Record),
}

func GetEmptyModels() []any {
	return models
}
