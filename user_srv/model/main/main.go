package main

import (
	"fmt"
	"mxshop_srvs/user_srv/comm"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/model"
	"time"
)

func main() {

	err := global.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	var users []model.User
	for i := 0; i < 10; i++ {
		now := time.Now()
		user := model.User{
			NikeName: fmt.Sprintf("dengc%d", i),
			Mobile:   fmt.Sprintf("1822973519%d", i),
			Password: comm.Encode(fmt.Sprintf("dengc%d", i)),
			Birthday: &now,
			Gender:   1,
		}
		users = append(users, user)
	}
	global.DB.Create(&users)
}
