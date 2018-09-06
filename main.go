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
	Profile   Profile
	ProfileID uint
}

func main() {
	db, _ := gorm.Open("sqlite3", "test.db")
	db = db.Set("gorm:auto_preload", true)

	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User{}, &Profile{})
	u := new(User)
	p := new(Profile)

	u.ProfileID = 1
	p.Name = "JCTESTEJC"

	db.Create(&u)
	db.Create(&p)

	db.Find(&u)
	log.Println(u)

}
