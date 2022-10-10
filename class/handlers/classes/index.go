package classes

import "go.h4n.io/openschool/class/models"

type ClassesIndexRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type ClassesIndexResponse struct {
	Classes []models.Class `json:"classes"`
}

func (h *ClassesHandler) ClassesIndex(r ClassesIndexRequest) (*ClassesIndexResponse, error) {
	items, err := h.Repository.GetAll(r.PerPage, r.Page)
	if err != nil {
		return nil, err
	}

	return &ClassesIndexResponse{Classes: items}, nil
}
