package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/Karthika-Rajagopal/fullstack/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Karthika",
		Email:    "karthi@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error //DropTableIfExists() method is called on the db object to drop any existing tables that have the same names as models.Post and models.User
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error //AutoMigrate() method is called on the db object to create the models.Post and models.User tables in the database and apply any necessary database schema changes
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error //AddForeignKey() method is called on the db object to create a foreign key constraint between the author_id column in the models.Post table and the id column in the models.User table.
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error//executed that seeds the models.User and models.Post tables with sample data
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
