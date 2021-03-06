package user

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user UserModel) (UserModel, error)
	FindUserByEmail(email string) (UserModel, error)
}

type DefaultUserService struct {
	DB *gorm.DB
	T  *TokenValidator
}

// Returns a new User service
func NewUserService(db *gorm.DB, t *TokenValidator) *DefaultUserService {
	return &DefaultUserService{
		DB: db,
		T:  t,
	}
}

func (s *DefaultUserService) CreateUser(u UserModel) (UserModel, error) {
	result := s.DB.Save(&u)
	return u, result.Error
}

func (s *DefaultUserService) FindUserByEmail(email string) (UserModel, error) {
	var user UserModel
	result := s.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (s *DefaultUserService) GetAllUsers() ([]UserModel, error) {
	var users []UserModel
	result := s.DB.Find(&users)
	return users, result.Error
}

func (s *DefaultUserService) GenerateJWT(email string, role string) (string, error) {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var mySigningKey = []byte(envs["AUTH_SECRET"])
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Printf("Error generating JWT: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func (s *DefaultUserService) ValidateToken(token string) *TokenVerificationResultModel {
	return &TokenVerificationResultModel{IsTokenValid: s.T.ValidateToken(token)}
}
