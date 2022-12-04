package user

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IUser interface {
	CreateUser(record *models.User) (*models.User, error)
	Login(user *models.User) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
}

type User struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewUser(appConfig *config.AppConfig, repo repos.IRepo) IUser {
	return &User{
		cfg:  appConfig,
		repo: repo,
	}
}

func (u User) CreateUser(record *models.User) (*models.User, error) {
	token := GenerateSecureToken(5)
	record.Token = token
	cz, err := u.repo.Citizen().Create(record.Citizen)
	record.Citizen.ID = cz.ID
	record, err = u.repo.User().Create(record)

	return record, err
}

func (u User) Login(user *models.User) (*models.User, error) {
	record, err := u.repo.User().Login(user)
	if record.Role == 1 {
		citizen, err := u.repo.Citizen().GetByID(record.CitizenID)
		if err != nil {
			fmt.Println("FAIL!!")
			return nil, err
		}
		record.Citizen = citizen
	}

	return record, err
}

func (u User) GetUserByToken(token string) (*models.User, error) {
	record, err := u.repo.User().GetUserByToken(token)
	return record, err
}
