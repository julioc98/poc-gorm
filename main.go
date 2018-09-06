package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Profile struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Profile   Profile `gorm:"foreignkey:ProfileID"` // use ProfileID as foreign key
	ProfileID uint
}

func main() {
	db, _ := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User{}, &Profile{})
	u := new(User)
	p := new(Profile)
	u.ProfileID = 1
	p.Name = "JCTESTEJC"
	db.Create(&u)
	db.Create(&p)

	user := new(User)
	profile := new(Profile)

	user.ProfileID = 1
	db.Model(&user).Related(&profile)
	//// SELECT * FROM profiles WHERE id = 111; // 111 is user's foreign key ProfileID

	log.Println(user.Profile)
	log.Println(profile)
}
