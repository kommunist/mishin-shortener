package deleted

// ошибка о том, что запись удалена.
// Так как использую в нескольких пакетах, то вытащил в отдельный
type DeletedError struct {
	Err error
}

func (e *DeletedError) Error() string {
	return e.Err.Error()
}

func NewDeletedError(err error) error {
	return &DeletedError{Err: err}
}
