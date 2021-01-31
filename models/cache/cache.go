package cache

import "errors"

/*
缓存存放地，本项目主要用作 redis 缓存数据

*/

var NoRows = errors.New("no errors")

func GetUserById(uid int) (interface{}, error) {
	// do something
	if uid == 0 {
		return nil, NoRows
	}
	user := uid
	return user, nil
}
