package domain

type Account struct {
	ID string `validate:"required,min=4,max=15,alphanumunicode"`
	Password string `validate:"required,min=8"`
	Name string `validate:"required"`
	EmailAddress string `validate:"required,email"`
	AccountType string `validate:"required,eq=Regular"`
}

type Accounts []Account
