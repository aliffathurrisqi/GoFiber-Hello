package migration

import (
	"fmt"
	"log"

	"github.com/aliffathurrisqi/GoFiber-MyApp/database"
	"github.com/aliffathurrisqi/GoFiber-MyApp/models"
)

func Migration(){
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil{
		log.Println(err)
	}

	fmt.Println("database migrated")
}

