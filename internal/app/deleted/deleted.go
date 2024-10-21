// Модуль deleted содержит ошибку об уже удаленных данных.
package deleted

type DeletedError struct {
	Err error
}

func (e *DeletedError) Error() string {
	return e.Err.Error()
}

func NewDeletedError(err error) error {
	return &DeletedError{Err: err}
}
