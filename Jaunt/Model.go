package Jaunt

import "github.com/google/uuid"

type Jaunt struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Expenses []Expense `json:"expenses"`
}

type Expense struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name         string    `json:"name"`
	Total        float64   `json:"total"`
	PaidByUserID uuid.UUID
	PaidBy       User   `json:"paid_by" gorm:"foreignKey:PaidByUserID"`
	PaidFor      []User `json:"paid_for" gorm:"many2many:expense_for_users;"`
	JauntID      uuid.UUID
}

type User struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// type Expense struct {
// 	ID       string   `json:"id"`
// 	Name     string   `json:"name"`
// 	Total    float64  `json:"total"`
// 	Paid_By  string   `json:"paid_by"`
// 	Paid_For []string `json:"paid_for"`
// }

// type Jaunt struct {
// 	ID       string    `json:"id"`
// 	Name     string    `json:"name"`
// 	Expenses []Expense `json:"expenses"`
// }

// var jaunts = []Jaunt{
// 	{
// 		ID: "1", Name: "Sandy's Bday", Expenses: []Expense{
// 			{ID: "1", Name: "Cake", Total: 1000, Paid_By: "Sai", Paid_For: []string{"MSK", "Shanki"}},
// 			{ID: "2", Name: "FunCity", Total: 1000, Paid_By: "MSK", Paid_For: []string{"Shanki", "Sai"}},
// 		},
// 	},
// }
