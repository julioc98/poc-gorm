package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func belongsTo() {

	type Profile struct {
		gorm.Model
		Name string
	}

	type User0 struct {
		gorm.Model
		Profile   Profile
		ProfileID uint
	}

	db, _ := gorm.Open("sqlite3", "test.db")
	db = db.Set("gorm:auto_preload", true)

	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User0{}, &Profile{})
	u := new(User0)
	p := new(Profile)

	u.ProfileID = 1
	p.Name = "JCTESTEJC"

	db.Create(&u)
	db.Create(&p)

	db.Find(&u)
	log.Println(u)
}

func hasMany() {

	type Email struct {
		gorm.Model
		Email   string
		User1ID uint
	}

	type User1 struct {
		gorm.Model
		Emails []Email
	}

	db, _ := gorm.Open("sqlite3", "test.db")
	db = db.Set("gorm:auto_preload", true)

	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User1{}, &Email{})
	u := new(User1)
	e := new(Email)

	e.Email = "jc@test.com"
	e.User1ID = 1

	db.Create(&u)
	db.Create(&e)

	db.Find(&u)
	log.Println(u)
}

func belongsToAndHasMany() {
	type Profile struct {
		gorm.Model
		Name string
	}

	type Email struct {
		gorm.Model
		Email   string
		User2ID uint
	}

	type User2 struct {
		gorm.Model
		Profile   Profile
		ProfileID uint
		Emails    []Email
	}

	db, _ := gorm.Open("sqlite3", "test.db")
	db = db.Set("gorm:auto_preload", true)

	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&User2{}, &Profile{}, &Email{})
	u := new(User2)
	p := new(Profile)
	e := new(Email)

	e.Email = "jc@test.com"
	e.User2ID = 1

	u.ProfileID = 1
	p.Name = "JCTESTEJC"

	db.Create(&p)
	db.Create(&e)
	db.Create(&u)

	db.Find(&u)
	log.Println(u)
}

func main() {
	// belongsTo()
	// hasMany()
	belongsToAndHasMany()
}
