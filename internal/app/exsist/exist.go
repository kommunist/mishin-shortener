// Модуль exist содержит ошибку об уже существующих данных.
package exsist

type ExistError struct {
	Err error
}

func (e *ExistError) Error() string {
	return e.Err.Error()
}

func NewExistError(err error) error {
	return &ExistError{Err: err}
}
