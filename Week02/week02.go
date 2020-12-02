package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
}

//dao层
type dao struct {
}

func newDbDao() dao {
	return dao{}
}

//dbGetUsers 模拟sql.ErrorNoRow
func dbGetUser(id int64) (user User, err error) {
	return User{}, sql.ErrNoRows
}

func (d dao) getUserMsg(id int64) (user User, err error) {
	user, err = dbGetUser(id)
	if err != nil {
		return User{}, errors.Wrap(err, "failed query")
	}
	return
}

//queryUserInfo 模拟service层
func queryUserInfo(id int64) (user User, err error) {
	d := newDbDao()
	if user, err = d.getUserMsg(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errors.WithMessagef(err, "query user: %d detail", id)
		}
	}
	return
}

func main() {
	var user_id int64 = 123
	user, err := queryUserInfo(user_id)
	if err != nil {
		fmt.Printf("query user failed,err:%v\n", err)
		return
	}
	fmt.Println(user)
}
