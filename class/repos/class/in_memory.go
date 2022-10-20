package class

import (
	"errors"
	"fmt"
	"time"

	"github.com/lucsky/cuid"
	"go.h4n.io/openschool/class/models"
)

type InMemoryClassRepository struct {
	Items []models.Class
}

func NewInMemoryClassRepository(itemCount int) InMemoryClassRepository {
	var items []models.Class

	for i := 0; i < itemCount; i++ {
		items = append(items, models.Class{
			Id:          cuid.New(),
			Name:        fmt.Sprintf(`class-%v`, i),
			DisplayName: fmt.Sprintf(`Class %v`, i),
			Description: fmt.Sprintf(`This is class %v`, i),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	return InMemoryClassRepository{Items: items}
}

func (r *InMemoryClassRepository) GetAll(perPage int, page int) ([]models.Class, error) {
	return r.Items, nil
}

func (r *InMemoryClassRepository) Get(id string) (*models.Class, error) {
	var found *models.Class

	for _, v := range r.Items {
		if v.Id == id {
			found = &v
		}
	}

	return found, nil
}

func (r *InMemoryClassRepository) Update(class *models.Class) (*models.Class, error) {
	var found *models.Class
	for _, v := range r.Items {
		if v.Id == found.Id {
			if class.Name != v.Name {
				return nil, errors.New("you cannot update Name after creation")
			}

			v.DisplayName = class.DisplayName
			v.Description = class.Description
			v.UpdatedAt = time.Now()

			found = &v
		}
	}

	if found == nil {
		return nil, errors.New("no existing class found by that id")
	}

	return found, nil
}

func (r *InMemoryClassRepository) Create(class models.Class) (*models.Class, error) {
	model := models.Class{
		Id:          cuid.New(),
		Name:        class.Name,
		DisplayName: class.DisplayName,
		Description: class.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	r.Items = append(r.Items, model)

	return &model, nil
}

func (r *InMemoryClassRepository) Delete(class models.Class) error {
	var newItems []models.Class

	for _, c := range r.Items {
		if c.Id != class.Id {
			newItems = append(newItems, c)
		}
	}

	r.Items = newItems

	return nil
}
