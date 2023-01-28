package main

import (
	"github.com/Shanki5/simply-split-server/Jaunt"
	"github.com/Shanki5/simply-split-server/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}
func main() {
	router := gin.Default()

	superGroup := router.Group("/jaunts")
	{
		// superGroup.GET("", getAllJaunts)
		superGroup.POST("", Jaunt.AddJaunt)
		// superGroup.GET("/:jauntId/expenses", getExpenses)
		superGroup.POST("/:jauntId/expenses", Jaunt.AddExpense)
	}
	router.Run("localhost:8080")
}
