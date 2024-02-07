package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `json:"account" valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

func (pixkey *PixKey) isValid() error {

	_, err := govalidator.ValidateStruct(pixkey)

	if pixkey.Kind != "email" && pixkey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixkey.Status != "active" && pixkey.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func newPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixkey := PixKey{
		Kind:    kind,
		Account: account,
		Key:     key,
	}

	pixkey.ID = uuid.NewV4().String()
	pixkey.CreatedAt = time.Now()

	if err := pixkey.isValid(); err != nil {
		return nil, err
	}

	return &pixkey, nil

}
