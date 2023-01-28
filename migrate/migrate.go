package main

import (
	"github.com/Shanki5/simply-split-server/Jaunt"
	"github.com/Shanki5/simply-split-server/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config.LoadEnvVariables()
	DB = config.ConnectDB()
}

func main() {
	DB.AutoMigrate(&Jaunt.User{})
	DB.AutoMigrate(&Jaunt.Expense{})
	DB.AutoMigrate(&Jaunt.Jaunt{})
}
