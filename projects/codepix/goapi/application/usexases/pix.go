package usecase

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind stirng, accountId string) (*model.PixKey, error)  {

}