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
