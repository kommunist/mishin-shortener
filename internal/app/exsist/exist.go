package exsist

// ошибка о том, что запись уже существует.
// Так как использую в нескольких пакетах, то вытащил в отдельный
type ExistError struct {
	Err error
}

func (e *ExistError) Error() string {
	return e.Err.Error()
}

func NewExistError(err error) error {
	return &ExistError{Err: err}
}
