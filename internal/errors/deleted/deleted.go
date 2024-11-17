// Модуль deleted содержит ошибку об уже удаленных данных.
package deleted

// Основная структура ошибки.
type DeletedError struct {
	Err error
}

// Вывод ошибки
func (e *DeletedError) Error() string {
	return e.Err.Error()
}

// Инициализация ошибки
func NewDeletedError(err error) error {
	return &DeletedError{Err: err}
}
