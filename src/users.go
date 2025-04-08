package src

import (
	"errors"
	"fmt"
)

type Usermanagement interface {
	createUser()
	findUser()
}

var AccountsUser [3]BancAccount
var Users []User

type User struct {
	Name string
	Id   int
}

func findUser(id int) bool {
	i := 0
	find := false
	for i <= len(Users) && find == false {
		if Users[i].Id == id {
			find = true
			return find
		} else {
			i++
		}
	}
	return find
}
func CreateUser(user User) error {
	if user.Id == 0 {
		return errors.New("ID is nil")
	}
	if user.Name == "" {
		return errors.New("user name is empty")
	}
	if findUser(user.Id) == true {
		return fmt.Errorf("user with %d already exist", user.Id)
	}
	Users = append(Users, user)
	return nil
}
