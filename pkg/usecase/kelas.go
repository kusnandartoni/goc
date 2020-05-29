package usecase

import "github.com/kusnandartoni/goc/pkg/domain/repository"

type KelasUsecase struct {
	repo repository.KelasRepository
}

func NewKelasUsecase(r repository.KelasRepository) *KelasUsecase {
	return &KelasUsecase{r}
}

func (rec *KelasUsecase) Save(data interface{}) error {
	return nil
}

func (rec *KelasUsecase) Read(param interface{}) (interface{}, error) {
	return nil, nil
}

func (rec *KelasUsecase) Delete(id int) error {
	return nil
}
