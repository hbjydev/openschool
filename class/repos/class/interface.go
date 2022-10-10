package class

import "go.h4n.io/openschool/class/models"

type ClassRepository interface {
	GetAll(perPage int, page int) ([]models.Class, error)

	Get(id string) (*models.Class, error)

	Update(class *models.Class) (*models.Class, error)

	Create(class models.Class) (*models.Class, error)

	Delete(class models.Class) error
}
