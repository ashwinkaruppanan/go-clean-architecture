package main

import (
	"log"
	"os"

	"github.com/cleancode/db"
	"github.com/cleancode/repository"
	"github.com/cleancode/router"
	"github.com/cleancode/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewMongoDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	rep := repository.NewUserRepository(db, "user")
	ser := service.NewUserService(rep, 5)

	gin := gin.Default()
	router.NewRouter(gin, ser)

	if err = gin.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}

}
