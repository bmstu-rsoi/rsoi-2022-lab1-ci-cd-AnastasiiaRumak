package delivery

import (
	"context"

	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/models"
)

type usecase interface {
	CreatePerson(ctx context.Context, person models.Person) (int64, error)
}
