package databases

import (
	"log"

	"github.com/wandersonsantos01/go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	strConn := "host=host.docker.internal user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(strConn))
	if err != nil {
		log.Panic("Error on connect on database")
	}
	DB.AutoMigrate(&models.Animal{})
}
