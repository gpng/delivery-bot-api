package database

import (
	"fmt"
	"log"

	"github.com/gpng/delivery-bot-api/config"
	c "github.com/gpng/delivery-bot-api/constants"
	"github.com/gpng/delivery-bot-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // required for postgres dbs
)

// New db connection and trigger migrations
func New(conf *config.Config) (*gorm.DB, error) {
	// connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", conf.DB.Host, conf.DB.User, conf.DB.Name, conf.DB.Password)
	log.Println(dbURI)

	db, err := openDB(dbURI)

	if err != nil {
		log.Println("failed to connect to db")
		return nil, err
	}

	migrate(db)

	log.Println("db connection successful")

	return db, nil
}

func openDB(connString string) (*gorm.DB, error) {
	log.Println("trying to connect to db...")
	conn, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func migrate(db *gorm.DB) {
	// db migration, automatically updates tables and schemas
	db.Exec(c.FunctionExtensionUUID)

	db.AutoMigrate(&models.Postcode{}, &models.LastNotification{})
}
