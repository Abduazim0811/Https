package service

import (
	"Items/internal/entity/items"
	"Items/internal/infrastructura/repository"
)


type ItemsService struct{
	repo repository.ItemsRepository
}

func NewItems(repo repository.ItemsRepository) *ItemsService{
	return &ItemsService{repo: repo}
}

func (i *ItemsService) Create(req items.CreateItems)error{
	return i.repo.AddItems(req)
}

func (i *ItemsService) GetById(id string) (*items.Items, error){
	return i.repo.GetbyidItems(id)
}

func (i *ItemsService) Getall()([]*items.Items, error){
	return i.repo.GetAll()
}

func (i *ItemsService) Updateitems(id string, req items.Items) error{
	return i.repo.Update(id, req)
}

func (i *ItemsService) Deleteitems(id string) error{
	return i.repo.Delete(id)
}