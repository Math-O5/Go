package model

type Transaction struct {
	ID string `json:"id" validate:"required,uuid4"`
	AccountID string `json:"account_id" validate:"required,uuid4"`
	Amount float64  `json:"amount" validate:"required,numeric"`
	PixKeyTo string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string `json:"pixKeyKindTo" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status string `json:"status" validate:"required"`
	Error string `json:"description"`
}

func (t* Transaction) isValid() error {
	v := validator.New()

	err := v.Struct(t)

	if err != nil {
		fmt.Errorf("Error during Transaction validation: %s", err.Error())
		return err
	}

	return nil 
}

func (t *Transaction) PaserJson(data []byte) error {
	err := json.Unmarshal(data, t)

	if err != nil {
		return err
	}

	if err = t.isValid(); err != nil {
		return err 
	}

	return nil
} 

func NewTransaction() *transaction {
	return &Transaction{}
}

func (t *Transaction) ToJson() ([]byte, error) {
	err := t.isValid()

	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}
	
	return result
}