package repository


type TransactionRepositoryDB struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryDB) Register(transaction *Transaction) error {
	err := t.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDB) Save(transaction *Transaction) error {
	err := t.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

Save(transaction *Transaction) error
Find(id string) (*Transaction, error)

func (t TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var trans model.Transaction

	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Error("no transaction was found")
	}

	if err != nil {
		return err
	}

	return &transaction, nil
}
