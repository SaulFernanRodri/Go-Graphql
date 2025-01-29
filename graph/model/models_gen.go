// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewOrderInput struct {
	User     string  `json:"user"`
	Product  string  `json:"product"`
	Quantity int32   `json:"quantity"`
	Total    float64 `json:"total"`
}

type NewUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Activate bool   `json:"activate"`
}

type Order struct {
	ID       string  `json:"id"`
	User     string  `json:"user"`
	Product  string  `json:"product"`
	Quantity int32   `json:"quantity"`
	Total    float64 `json:"total"`
}

type Query struct {
}

type Subscription struct {
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Activate bool   `json:"activate"`
}
