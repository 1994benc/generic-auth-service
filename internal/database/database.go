package database

import (
	"database/sql"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create a new database connection
func New() (*gorm.DB, error) {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mysqlConnectionString := envs["MYSQL_CONNECTION_STRING"]
	log.Println("Setting up new database connection")
	sqlDB, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	retryCount := 30
	for {
		if err != nil {
			if retryCount == 0 {
				log.Fatalln("Out of tries - couldn't connect to db!")
				break
			}
			log.Printf("Still not connected to db - retrying...", err)
			gormDB, err = gorm.Open(mysql.New(mysql.Config{
				Conn: sqlDB,
			}), &gorm.Config{})
			retryCount--
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}
	return gormDB, err
}
