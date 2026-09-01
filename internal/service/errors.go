package service

import "errors"

var (
	ErrInvalidCredentials = errors.New("неверный login или пароль")
	ErrMissingFields      = errors.New("все поля обязательны")
	ErrInvalidEmail       = errors.New("некорректный email")
	ErrInvalidPassword    = errors.New("пароль должен быть не менее 8 символов")
)
