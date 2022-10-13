package users

import (
	"log"

	"github.com/bpradana/goilerplate/pkg/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.UserRepository {
	db.AutoMigrate(&domain.User{})

	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User

	err := u.db.Find(&users).Error
	if err != nil {
		log.Println("[userRepository] [GetAll] error querying all users, err: ", err.Error())
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetById(id int) (domain.User, error) {
	var user domain.User

	err := u.db.First(&user, id).Error
	if err != nil {
		log.Println("[userRepository] [GetById] error querying user, err: ", err.Error())
		return user, err
	}

	return user, nil
}

func (u *userRepository) Create(user *domain.User) (*domain.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		log.Println("[userRepository] [Create] error creating user, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Update(id int, user *domain.User) (*domain.User, error) {
	var oldUser domain.User

	err := u.db.First(&oldUser, id).Error
	if err != nil {
		log.Println("[userRepository] [Update] error querying user, err: ", err.Error())
		return nil, err
	}

	err = u.db.Model(&oldUser).Updates(user).Error
	if err != nil {
		log.Println("[userRepository] [Update] error updating user, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Delete(id int) error {
	var user domain.User

	err := u.db.Delete(&user, id).Error
	if err != nil {
		log.Println("[userRepository] [Delete] error deleting user, err: ", err.Error())
		return err
	}

	return nil
}
