// package config provides  configuration by environment
package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-gin-mysql-boilerplate/migrations"
	"log"
	"os"

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // mysql golang driver
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {

	// display info about config
	printInfo(env)

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)

	if env == "test" {
		v.AddConfigPath("../config/")
	} else {
		v.AddConfigPath("./config/")
	}

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	config = v
}

// main function call this to get config
func GetConfig() *viper.Viper {
	return config
}

// print some info on init
func printInfo(env string) {
	fmt.Println("")
	fmt.Println("***************************************************")
	fmt.Println("Application started with config:", env)
	fmt.Println("***************************************************")
	fmt.Println("")
}

//for check connection db progress  ----------------------------------------------------------------------------------------

var db_res *gorm.DB

// Check Connection Golang with mysql Database
func OpenConnection() *gorm.DB {

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
	mysqlInfo := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local" //if not use mamp
	fmt.Println("mysqlInfo", mysqlInfo)
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// check the connection
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Connecting Database " + host + ":" + port)

	//create migration table
	migrations.CreateMigrationTable()

	fmt.Println("db first", db)

	db_res = db

	// return the connection
	return db
}

func GetConnection() *gorm.DB {
	return db_res
}

// NullString marshal json -------------------------------------------------------------------------------------------------

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
