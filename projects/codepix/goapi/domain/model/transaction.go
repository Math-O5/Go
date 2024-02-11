package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "peding"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed        = "confirmed"
)

// type Transaction struct {
// 	Transaction []Transaction
// }

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyToID        string   `gorm:"column:pix_key_id_to;type:uuid;not null;" valid:"-"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"notnull"`
}

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

func (transaction *Transaction) isValid() error {

	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if transaction.Status != TransactionCompleted && transaction.Status != TransactionConfirmed && transaction.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	if transaction.PixKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func newTransaction(accountFrom *Account, description string, amount float64, account *Account, key *PixKey) (*Transaction, error) {
	t := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    key,
		Status:      TransactionPending,
		Description: description,
	}

	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()

	if err := t.isValid(); err != nil {
		return nil, err
	}

	return &t, nil

}

/* Regras de negócio: sem set, gets */
func (t *Transaction) Complete() error {

	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()

	err := t.isValid()

	return err
}

func (t *Transaction) Cancel(description string) error {

	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description

	err := t.isValid()

	return err
}

func (t *Transaction) Confirm() error {

	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()

	err := t.isValid()

	return err
}
