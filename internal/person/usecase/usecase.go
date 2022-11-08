package usecase

import (
	"context"

	"github.com/bmstu-rsoi/rsoi-2022-lab1-ci-cd-AnastasiiaRumak/internal/models"
)

type UseCase struct {
	repo repo
}

func New(r repo) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) CreatePerson(ctx context.Context, person models.Person) (int64, error) {
	return u.repo.CreatePerson(ctx, person)
}
