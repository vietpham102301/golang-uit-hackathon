package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type UserSQLRepo struct {
	db *gorm.DB
}

func NewUserSQLRepo(db *gorm.DB) IUserRepo {
	return &UserSQLRepo{
		db: db,
	}
}

func (u UserSQLRepo) Create(record *models.User) (*models.User, error) {
	err := u.db.Create(record).Error
	return record, err
}

func (u UserSQLRepo) Login(user *models.User) (*models.User, error) {
	err := u.db.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error
	return user, err
}

func (u UserSQLRepo) GetUserByToken(token string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Where("token = ?", token).First(user).Error
	return user, err
}
