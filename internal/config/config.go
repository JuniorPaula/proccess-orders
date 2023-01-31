package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ConnectStringDatabase = ""
)

func InitVariables() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ConnectStringDatabase = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
}
