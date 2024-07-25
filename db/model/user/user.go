package user

import (
	"context"
	"github.com/jhue/misgo/db/model/key"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID  string `gorm:"unique"`
	Name string `gorm:"unique"`
}

func WithUser(parent context.Context, user User) context.Context {
	return context.WithValue(parent, key.User, user)
}

func ExtractUser(ctx context.Context) (user User, ok bool) {
	user, ok = ctx.Value(key.User).(User)
	return
}
