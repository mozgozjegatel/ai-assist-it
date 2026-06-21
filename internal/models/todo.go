package models

import "time"

// Todo представляет модель задачи.
// @Description Задача с привязкой к пользователю, заголовком, описанием и статусом.
type Todo struct {
	// ID уникальный идентификатор задачи
	// Example: 1
	ID int `json:"id" example:"1"`

	// ID пользователя, которому принадлежит задача
	// Example: 5
	UserID int `json:"user_id" example:"5"`

	// Заголовок задачи
	// Example: Купить продукты
	Title string `json:"title" example:"Купить продукты"`

	// Детальное описание задачи
	// Example: Молоко, хлеб, яйца, масло
	Description string `json:"description" example:"Молоко, хлеб, яйца, масло"`

	// Статус выполнения (true - выполнена, false - не выполнена)
	// Example: false
	Done bool `json:"done" example:"false"`

	// Дата и время создания задачи
	// Example: 2024-03-20T10:30:00Z
	CreatedAt time.Time `json:"created_at" example:"2024-03-20T10:30:00Z"`

	// Дата и время последнего обновления задачи
	// Example: 2024-03-20T12:00:00Z
	UpdatedAt time.Time `json:"updated_at" example:"2024-03-20T12:00:00Z"`
}

// CreateTodoRequest представляет запрос на создание задачи.
// @Description Модель для создания новой задачи.
type CreateTodoRequest struct {
	// ID пользователя (обязательное поле)
	// Required: true
	// Example: 5
	UserID int `json:"user_id" validate:"required" example:"5"`

	// Заголовок задачи (обязательное поле, не более 200 символов)
	// Required: true
	// Example: Купить продукты
	Title string `json:"title" validate:"required,max=200" example:"Купить продукты"`

	// Описание задачи (не более 1000 символов)
	// Example: Молоко, хлеб, яйца, масло
	Description string `json:"description" validate:"max=1000" example:"Молоко, хлеб, яйца, масло"`

	// Статус выполнения (по умолчанию false)
	// Example: false
	Done bool `json:"done" example:"false"`
}

// UpdateTodoRequest представляет запрос на обновление задачи.
// @Description Модель для обновления существующей задачи.
type UpdateTodoRequest struct {
	// Заголовок задачи (опционально, не более 200 символов)
	// Example: Обновленный заголовок
	Title *string `json:"title" validate:"omitempty,max=200" example:"Обновленный заголовок"`

	// Описание задачи (опционально, не более 1000 символов)
	// Example: Новое подробное описание
	Description *string `json:"description" validate:"omitempty,max=1000" example:"Новое подробное описание"`

	// Статус выполнения (опционально)
	// Example: true
	Done *bool `json:"done" example:"true"`
}
