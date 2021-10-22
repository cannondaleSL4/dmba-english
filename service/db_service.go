package service

import (
	"github.com/dmba-english/db"
)

func GetWords(userId int) []*db.Dict {
	userExist := db.CheckUserExist(userId)
	var dict []*db.Dict
	if userExist {
		dict = db.GetWords(userId)
	} else {
		db.SaveNewUser(userId)
	}
	return dict
}
