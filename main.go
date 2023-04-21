package main

import (
	"final-project/app"
	"final-project/config"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {

	app.StartApplication()
}
