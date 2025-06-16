package dto

type UserRequest struct {
	Name       string `json:"name" validate:"required,min=5"`
	Profile    string `json:"profile" validate:"required"`
	Document   string `json:"document" validate:"required"`
	Pin        int    `json:"pin" validate:"required,len=6"`
	Coercion   int    `json:"coercion" validate:"required,len=7"`
	CardNumber int    `json:"card_number" validate:"required,len=6"`
	Status     bool   `json:"status" validate:"required"`
	WorkStart  string `json:"work_start" validate:"required"`
	WorkEnd    string `json:"work_end" validate:"required"`
}

type UserResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Profile    string `json:"profile"`
	Document   string `json:"document"`
	CardNumber int    `json:"card_number"`
	Status     bool   `json:"status"`
	WorkStart  string `json:"work_start"`
	WorkEnd    string `json:"work_end"`
}

type UserUpdateRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name" validate:"required,min=5"`
	Profile    string `json:"profile" validate:"required"`
	Document   string `json:"document"`
	Pin        int    `json:"pin" validate:"required,len=6"`
	Coercion   int    `json:"coercion" validate:"required,len=7"`
	CardNumber int    `json:"card_number" validate:"required,len=6"`
	Status     bool   `json:"status" validate:"required"`
	WorkStart  string `json:"work_start" validate:"required"`
	WorkEnd    string `json:"work_end" validate:"required"`
}

type QueryPinReponse struct {
	Name     string `json:"name" validate:"required,min=5"`
	Profile  string `json:"profile" validate:"required"`
	Document string `json:"document"`
}
