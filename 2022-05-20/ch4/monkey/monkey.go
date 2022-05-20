package p39

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func fooDatabaseCaseByFunc() int {
	return getUserAgeFunc()
}

func getUserAgeFunc() int {
	var user User
	db, err := gorm.Open(postgres.Open("host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword"), &gorm.Config{})
	if err != nil {
		panic("connect fail")
	}
	res := db.First(&user)
	if res.Error != nil {
		panic("error")
	}

	return user.Age
}

type User struct {
	Age int
}
