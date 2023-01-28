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
	PaidByUserID uuid.UUID `json:"paid_by_user_id"`
	PaidBy       User      `json:"paid_by" gorm:"foreignKey:PaidByUserID"`
	PaidFor      []User    `json:"paid_for" gorm:"many2many:expense_for_users;"`
	JauntID      uuid.UUID `gorm:"<-:create"`
}

type User struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `json:"name"`
	Email string    `json:"email" gorm:"unique"`
}
