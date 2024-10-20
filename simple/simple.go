package simple

import "errors"

type SimpleRepository struct {
	Err bool
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository(ErrBool bool) *SimpleRepository {
	return &SimpleRepository{
		Err: ErrBool,
	}
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Err {
		return nil, errors.New("Provider Error")
	}
	return &SimpleService{
		SimpleRepository: repository,
	}, nil
}