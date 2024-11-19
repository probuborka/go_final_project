package entity

import "errors"

var (
	ErrFormatError   = errors.New("format error")
	ErrNotFound      = errors.New("not found")
	ErrNotInInterval = errors.New("not in the interval")
	ErrTaskNotFound  = errors.New("Задача не найдена")
	ErrNoID          = errors.New("Не указан идентификатор")
)

type Error struct {
	Error string `json:"error"`
}
