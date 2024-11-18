package entity

import "errors"

var (
	ErrTaskNotFound = errors.New("Задача не найдена")
	ErrNoID         = errors.New("Не указан идентификатор")
)

type Error struct {
	Error string `json:"error"`
}
