package service

import (
	"IMProject/dao"
	"IMProject/model"
)

func Login(username string, password string) string {
	user := model.User{Username: username, Password: password}
	var result string
	if dao.Login(user) {
		result = "success"
	} else {
		result = "fail"
	}
	return result
}
