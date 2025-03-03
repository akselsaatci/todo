package repo

import (
	"authService/internal/core/domain"
	CustomErrors "authService/internal/core/errors"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthPostgresRepository struct {
	db *gorm.DB
}

type DbCredantials struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewAuthPostgresRepository(credantials DbCredantials) *AuthPostgresRepository {
	host := credantials.Host
	port := credantials.Port
	user := credantials.User
	password := credantials.Password
	dbname := credantials.DbName

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}

	return &AuthPostgresRepository{
		db: db,
	}
}

func (p *AuthPostgresRepository) FindUserByPassword(username string, password string) (*domain.User, error) {
	var user domain.User
	req := p.db.Where(&domain.User{UserName: username, Password: password}).First(&user)
	if req.Error != nil {
		if errors.Is(req.Error, gorm.ErrRecordNotFound) {
			return nil, nil // Kullanıcı bulunamadığında nil döndürün
		}
		return nil, fmt.Errorf("User couldn't found!: %v", req.Error)
	}
	return &user, nil
}
func (p *AuthPostgresRepository) FindUserById(id string) (*domain.User, error) {
	var user *domain.User
	uid, err := uuid.Parse(id)
	if err == nil {
		return nil, err
	}

	req := p.db.Find(user, domain.User{ID: uid})
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}
	return user, nil
}

func (p *AuthPostgresRepository) AddUserToDb(user *domain.User) error {
	doesUserExist := p.db.Where("email = ? OR user_name = ?", user.Email, user.UserName).First(&domain.User{}).Error

	if doesUserExist == nil {
		return CustomErrors.EmailOrUsernameExists
	}

	result := p.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
