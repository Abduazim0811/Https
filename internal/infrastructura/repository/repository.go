package repository

import "Items/internal/entity/items"


type ItemsRepository interface{
	AddItems(req items.CreateItems)error
	GetbyidItems(id string) (*items.Items, error)
	GetAll()([]*items.Items, error)
	Update(id string, updateData items.Items) error
	Delete(id string)error
}