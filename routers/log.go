package routers

import (
	"ticket/database"
	"ticket/util"
	"time"
)

func GetLog(userId int64, operation string) *database.Log {
	var userDao database.UserImpl
	user, _ := userDao.GetUserById(userId)
	log := &database.Log{
		Id:         util.GetSnowflakeId(),
		UserId:     userId,
		CreateTime: time.Now(),
		Operation:  operation,
		Type:       int(user[0].Type),
	}
	return log
}
