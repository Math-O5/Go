package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey ` valid:"-"`
}

func (account *Account) isValid() error {

	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

func newAccount(bank *Bank, number string, owenerName string) (*Account, error) {
	account := Account{
		OwnerName: owenerName,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.isValid(); err != nil {
		return nil, err
	}

	return &account, nil
}
