package Jaunt

import (
	"net/http"

	"github.com/Shanki5/simply-split-server/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllJaunts(context *gin.Context) {

}

func AddJaunt(context *gin.Context) {

	var newJaunt Jaunt

	if err := context.BindJSON(&newJaunt); err != nil {
		return
	}

	result := config.DB.Create(&newJaunt)

	if result.Error != nil {
		context.Status(400)
		return
	}
	context.IndentedJSON(http.StatusCreated, newJaunt.ID)
}

func AddExpense(context *gin.Context) {

	jauntId, err := uuid.Parse(context.Param("jauntId"))

	if err != nil {
		context.Status(400)
		return
	}

	var newExpense Expense

	if err := context.BindJSON(&newExpense); err != nil {
		return
	}
	newExpense.JauntID = jauntId

	result := config.DB.Create(&newExpense)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.IndentedJSON(http.StatusCreated, newExpense)
}

// func getAllJaunts(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOK, jaunts)
// }

// func addJaunt(context *gin.Context) {
// 	var newJaunt Jaunt

// 	if err := context.BindJSON(&newJaunt); err != nil {
// 		return
// 	}

// 	jaunts = append(jaunts, newJaunt)

// 	context.IndentedJSON(http.StatusCreated, newJaunt)
// }

// func getExpenses(context *gin.Context) {
// 	jauntId := context.Param("jauntId")
// 	fmt.Println(jauntId)
// 	var expenses []Expense
// 	for _, jaunt := range jaunts {
// 		if jaunt.ID == jauntId {
// 			expenses = append(expenses, jaunt.Expenses...)
// 		}
// 	}
// 	context.IndentedJSON(http.StatusOK, expenses)
// }
// func addExpense(context *gin.Context) {
// 	jauntId := context.Param("jauntId")
// 	var newExpense Expense

// 	if err := context.BindJSON(&newExpense); err != nil {
// 		return
// 	}
// 	fmt.Println(newExpense)
// 	for _, jaunt := range jaunts {
// 		if jaunt.ID == jauntId {
// 			currJaunt := &jaunt
// 			currJaunt.Expenses = append(jaunt.Expenses, newExpense)
// 			fmt.Println(jaunt.Expenses)
// 		}
// 	}
// 	context.IndentedJSON(http.StatusCreated, newExpense)
// }
