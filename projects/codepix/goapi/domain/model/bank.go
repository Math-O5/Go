package model

import (
	"time"

	"github.com/google/uuid"
)

type Bank struct {
	Base
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (bank *Bank) isValid() error {
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}
