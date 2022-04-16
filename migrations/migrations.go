// package config provides  configuration by environment
package migrations

import (
	"fmt"
	"go-gin-mysql-boilerplate/models"
	"log"
	"os"

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // mysql golang driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetConnectionMigration() *gorm.DB {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//get from your env
	var host string = os.Getenv("DB_HOST")
	var port string = os.Getenv("DB_PORT")
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var dbname string = os.Getenv("DB_NAME")

	// open Connection DB
	mysqlInfo := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// check the connection
	if err != nil {
		panic(err)
	}

	// return the connection
	return db
}

// Create Migration Table Golang with mysql Database
func CreateMigrationTable() string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// open connection db mysql
	db := GetConnectionMigration()

	//get from your env
	var all_migration string = os.Getenv("CREATE_MIGRATION_ALL")
	var migrate_users string = os.Getenv("CREATE_MIGRATION_USERS")
	var migrate_books string = os.Getenv("CREATE_MIGRATION_BOOKS")

	fmt.Println("Prepare Migration Table")

	// for migrations all
	if all_migration == "1" {

		//create users
		users := models.Users{}
		db.AutoMigrate(&users)
		fmt.Println("Create Migration Table 'users' Successfully")

		//create books
		detail_users := models.Books{}
		db.AutoMigrate(&detail_users)
		fmt.Println("Create Migration Table 'books' Successfully")

	} else if all_migration == "2" {

		//create users
		if migrate_users == "1" {
			users := models.Users{}
			db.AutoMigrate(&users)
			fmt.Println("Create Migration Table 'users' Successfully")
		}

		//create books
		if migrate_books == "1" {
			books := models.Books{}
			db.AutoMigrate(&books)
			fmt.Println("Create Migration Table 'books' Successfully")
		}

		fmt.Println("Successfully Migration Selecting Table")

	} else {

		fmt.Println("No Tables Are Migrated")

	}

	// return the connection
	return "working"
}
