package Jaunt

import (
	"fmt"
	"net/http"

	"github.com/Shanki5/simply-split-server/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		fmt.Println(err)
		context.Status(400)
		return
	}

	var newExpense Expense

	newExpense.JauntID = jauntId

	if err := context.BindJSON(&newExpense); err != nil {
		fmt.Println(err)
		return
	}
	result := config.DB.Create(&newExpense)
	if result.Error != nil {
		context.Status(400)
		return
	}

	context.IndentedJSON(http.StatusCreated, newExpense)
}

func GetAllJaunts(context *gin.Context) {
	var jaunts []Jaunt
	err := config.DB.Model(&Jaunt{}).Preload("Expenses").Preload("Expenses.PaidBy").Preload("Expenses.PaidFor").Find(&jaunts).Error
	if err != nil {
		context.Status(400)
		return
	}
	context.IndentedJSON(http.StatusOK, jaunts)
}

func GetJauntByID(context *gin.Context) {

	jauntId, err := uuid.Parse(context.Param("jauntId"))

	if err != nil {
		context.Status(400)
		return
	}

	var jaunt Jaunt
	err = config.DB.Model(&Jaunt{}).Preload("Expenses").Preload("Expenses.PaidBy").Preload("Expenses.PaidFor").First(&jaunt, jauntId).Error
	if err != nil {
		context.Status(400)
		return
	}

	context.IndentedJSON(http.StatusOK, jaunt)
}

func UpdateExpense(context *gin.Context) {
	jauntID, err := uuid.Parse(context.Param("jauntId"))

	if err != nil {
		context.Status(400)
		return
	}

	expenseID, err := uuid.Parse(context.Param("expenseId"))

	if err != nil {
		context.Status(400)
		return
	}

	var expense Expense
	if err := context.BindJSON(&expense); err != nil {
		return
	}

	expense.JauntID = jauntID
	expense.ID = expenseID

	err = config.DB.Save(&expense).Error
	if err != nil {
		fmt.Println(err)
		context.Status(401)
		return
	}

	context.Status(http.StatusAccepted)

}

func GetAllUsers(context *gin.Context) {
	var users []User

	err := config.DB.Model(&User{}).Find(&users).Error
	if err != nil {
		context.Status(500)
		return
	}

	context.IndentedJSON(http.StatusOK, users)
}
