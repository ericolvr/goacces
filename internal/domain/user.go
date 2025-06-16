package domain

import "time"

type User struct {
	ID             int    `json:"id"`
	Server         string `json:"server"`
	ChangePassword bool   `json:"change_password"`

	Name       string    `json:"name"`
	Profile    string    `json:"profile"` // 0 Admin - 1 Operator
	Document   string    `json:"document"`
	Pin        int       `json:"pin"`
	Coercion   int       `json:"coercion"`
	CardNumber int       `json:"card_number"`
	Status     bool      `json:"status"`
	WorkStart  time.Time `json:"work_start"`
	WorkEnd    time.Time `json:"work_end"`
}
