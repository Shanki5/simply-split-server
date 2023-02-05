package main

import (
	"github.com/Shanki5/simply-split-server/Jaunt"
	"github.com/Shanki5/simply-split-server/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}
func main() {
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	router.Use(cors.Default())

	superGroup := router.Group("/jaunts")
	{
		superGroup.GET("", Jaunt.GetAllJaunts)
		superGroup.POST("", Jaunt.AddJaunt)
		// superGroup.GET("/:jauntId/expenses", getExpenses)
		superGroup.POST("/:jauntId/expenses", Jaunt.AddExpense)
		superGroup.GET("/:jauntId", Jaunt.GetJauntByID)
		superGroup.PUT("/:jauntId/expenses/:expenseId", Jaunt.UpdateExpense)
	}

	router.GET("/users", Jaunt.GetAllUsers)
	
	router.Run("localhost:8080")
}
