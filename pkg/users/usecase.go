package users

import (
	"log"

	"github.com/bpradana/goilerplate/pkg/domain"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) GetAll() ([]domain.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		log.Println("[userUsecase] [Create] error getting all users, err: ", err.Error())
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) GetById(id int) (domain.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		log.Println("[userUsecase] [GetById] error getting user, err: ", err.Error())
		return user, err
	}

	return user, nil
}

func (u *userUsecase) Create(user *domain.User) (*domain.User, error) {
	user, err := u.repo.Create(user)
	if err != nil {
		log.Println("[userUsecase] [Create] error creating user, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Update(id int, user *domain.User) (*domain.User, error) {
	user, err := u.repo.Update(id, user)
	if err != nil {
		log.Println("[userUsecase] [Update] error updating user, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Delete(id int) error {
	err := u.repo.Delete(id)
	if err != nil {
		log.Println("[userUsecase] [Delete] error deleting user, err: ", err.Error())
		return err
	}

	return nil
}
