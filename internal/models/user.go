package models

import "time"

// User представляет модель пользователя в системе.
// @Description Пользователь с email, именем и временными метками.
type User struct {
	// ID уникальный идентификатор пользователя
	// Example: 1
	ID int `json:"id" example:"1"`

	// Email пользователя (уникальный)
	// Example: user@example.com
	Email string `json:"email" example:"user@example.com"`

	// Имя пользователя
	// Example: Иван Петров
	Name string `json:"name" example:"Иван Петров"`

	// Дата и время создания записи
	// Example: 2024-03-20T10:30:00Z
	CreatedAt time.Time `json:"created_at" example:"2024-03-20T10:30:00Z"`
}

// CreateUserRequest представляет запрос на создание пользователя.
// @Description Модель для создания нового пользователя.
type CreateUserRequest struct {
	// Email пользователя (обязательное поле, должен быть валидным email)
	// Required: true
	// Example: user@example.com
	Email string `json:"email" validate:"required,email" example:"user@example.com"`

	// Имя пользователя (обязательное поле, не более 100 символов)
	// Required: true
	// Example: Иван Петров
	Name string `json:"name" validate:"required,max=100" example:"Иван Петров"`
}

// UpdateUserRequest представляет запрос на обновление пользователя.
// @Description Модель для обновления данных пользователя.
type UpdateUserRequest struct {
	// Email пользователя (опционально)
	// Example: newemail@example.com
	Email *string `json:"email" validate:"omitempty,email" example:"newemail@example.com"`

	// Имя пользователя (опционально, не более 100 символов)
	// Example: Иван Сидоров
	Name *string `json:"name" validate:"omitempty,max=100" example:"Иван Сидоров"`
}

// ErrorResponse представляет ответ с ошибкой.
// @Description Модель для сообщений об ошибках.
type ErrorResponse struct {
	// Текстовый код ошибки
	// Example: invalid_request
	Error string `json:"error" example:"invalid_request"`

	// Человекочитаемое сообщение об ошибке
	// Example: Неверный формат запроса
	Message string `json:"message" example:"Неверный формат запроса"`
}

// SuccessResponse представляет успешный ответ API.
// @Description Стандартный успешный ответ с данными.
type SuccessResponse struct {
	// Статус операции (всегда true при успехе)
	// Example: true
	Success bool `json:"success" example:"true"`

	// Данные ответа (может быть объектом или массивом)
	Data interface{} `json:"data,omitempty"`
}
