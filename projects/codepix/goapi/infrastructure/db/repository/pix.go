package repository 

// type PixKey struct {
	// RegisterKey(pixKey *PixKey) (*PixKey, error)
	// FindKeyByKind(key string, kind string) (*PixKey, error)
	// AddBank(bank *Bank) error
	// AddAccount(account *Account) error
	// FindAccount(id string) (*Account, error)
// }

type PixKeyRepositoryDB struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDB) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDB) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error

	if err != nil {
		return err
	}

	return PixKey, nil
}


func (r PixKeyRepositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Error("no key was found")
	}

	if err != nil {
		return err
	}

	return &pixkey, nil
}


func (r PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Error("no account was found")
	}

	if err != nil {
		return err
	}

	return &account, nil
}

func (r PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Error("no bank was found")
	}

	if err != nil {
		return err
	}

	return &bank, nil
}

