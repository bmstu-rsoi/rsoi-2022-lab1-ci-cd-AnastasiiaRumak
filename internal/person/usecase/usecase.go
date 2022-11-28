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

func (u *UseCase) DeletePerson(ctx context.Context, id int64) error {
	return u.repo.DeletePerson(ctx, id)
}
func (u *UseCase) UpdatePerson(ctx context.Context, person models.Person) (models.Person, error) {
	c, err := u.repo.GetPersonID(ctx, person.ID)
	if err != nil {
		return models.Person{}, err
	}

	merge := mergePersons(c, person)
	err = u.repo.UpdatePerson(ctx, merge)
	if err != nil {
		return models.Person{}, err
	}

	return merge, nil
}

func mergePersons(c models.Person, update models.Person) models.Person {
	var name, address, work string
	var age int64

	name = c.Name
	if update.Name != "" {
		name = update.Name
	}
	age = c.Age
	if update.Age != 0 {
		age = update.Age
	}
	address = c.Address
	if update.Address != "" {
		address = update.Address
	}
	work = c.Work
	if update.Work != "" {
		work = update.Work
	}

	return models.Person{
		ID:      c.ID,
		Name:    name,
		Age:     age,
		Address: address,
		Work:    work,
	}
}

func (u *UseCase) GetPersonID(ctx context.Context, id int64) (models.Person, error) {
	return u.repo.GetPersonID(ctx, id)
}
func (u *UseCase) GetAll(ctx context.Context) (*[]models.Person, error) {
	return u.repo.GetAll(ctx)
}




//CreatePerson(ctx context.Context, person models.Person) (int64, error)
//DeletePerson(ctx context.Context, id int64) error
//UpdatePerson(ctx context.Context, person models.Person) (models.Person, error)
//GetPersonID(ctx context.Context, id int64) (models.Person, error)
//GetAll(ctx context.Context) (*[]models.Person, error)