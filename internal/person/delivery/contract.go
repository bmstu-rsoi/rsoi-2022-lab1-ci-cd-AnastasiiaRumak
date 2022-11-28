package delivery

import (
	"context"

	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/models"
)

type usecase interface {
	CreatePerson(ctx context.Context, person models.Person) (int64, error)
	DeletePerson(ctx context.Context, id int64) error
	UpdatePerson(ctx context.Context, person models.Person) (models.Person, error)
	GetPersonID(ctx context.Context, id int64) (models.Person, error)
	GetAll(ctx context.Context) (*[]models.Person, error)
}

type request struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64    `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

type response struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64    `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}
type httpError struct {
	Message string `json:"message"`
}


func fromModel(m models.Person) response {
	return response{	
		ID:      m.ID,
		Name:    m.Name,
		Age:     m.Age,
		Address: m.Address,
		Work:    m.Work,
	}
}

func toModel(req request) models.Person {
	return models.Person{
		ID:      req.ID,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
		Work:    req.Work,
	}
}