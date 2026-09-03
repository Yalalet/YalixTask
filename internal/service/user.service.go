package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

const defaultRoleID = 2

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Login(login, password string) (*models.User, string, error) {
	user, err := s.Repo.GetUserByLogin(login)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *UserService) Register(firstname, lastname, login, email, password string) (*models.User, error) {
	// проверка данных
	if lastname == "" || firstname == "" || login == "" || email == "" || password == "" {
		return nil, ErrMissingFields
	}

	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	if containsHTML(firstname) || containsHTML(lastname) || containsHTML(login) {
		return nil, ErrXSS
	}

	if len(password) < 8 {
		return nil, ErrInvalidPassword
	}

	// хешируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FirstName:    firstname,
		LastName:     lastname,
		Login:        login,
		Email:        email,
		RoleID:       defaultRoleID,
		PasswordHash: string(hash),
	}

	// сохраняем через репозиторий (сам service SQL не пишет!)
	if err := s.Repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}
