package database

import (
	"golang-mini-project/server/entity"
	"log"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

//Migrate create/updates database table
func Migrate(table *entity.Tasking) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
