package config

import (
	"os"

	"fmt"

	"github.com/darienkentanu/test-erajaya/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConfig() (config map[string]string) {
	conf, err := godotenv.Read()
	if err != nil {
		fmt.Println("cannot read '.env' files -> reading docker CONN_STRING")
		return nil
	}
	return conf
}

func InitDB() *gorm.DB {
	conf := GetConfig()
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf["DB_USERNAME"], conf["DB_PASSWORD"], conf["DB_HOST"],
		conf["DB_PORT"], conf["DB_NAME"],
	)

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot use '.env' files for db-connections -> using docker CONN_STRING")

		connStrDocker := os.Getenv("CONN_STRING")
		db, err2 := gorm.Open(mysql.Open(connStrDocker), &gorm.Config{})
		if err2 != nil {
			panic(err2)
		}
		initMigration(db)
		return db
	}
	initMigration(db)
	return db
}

func initMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Product{})
}
