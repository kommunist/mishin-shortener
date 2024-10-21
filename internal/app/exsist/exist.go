// Модуль exist содержит ошибку об уже существующих данных.
package exsist

// Основная структура ошибки.
type ExistError struct {
	Err error
}

// Вывод ошибки
func (e *ExistError) Error() string {
	return e.Err.Error()
}

// Инициализация ошибки
func NewExistError(err error) error {
	return &ExistError{Err: err}
}
