package models

import (
	"fmt"

	"github.com/chensanle/gin-example/models/cache"
	"github.com/chensanle/gin-example/models/mysql"
)

type User struct {
	*mysql.User

	age int
}

func (u *User) GetUser() (*User, error) {
	user := new(User)
	cUser, err := cache.GetUserById(0)
	if err == nil {
		user.User = cUser.(*mysql.User)
		return user, nil
	}
	fmt.Println(err)

	pUser, err := mysql.NewEmptyUser().Get()
	if err != nil {
		return nil, err
	}
	user.User = pUser
	user.age = pUser.Birthday

	return user, nil
}
