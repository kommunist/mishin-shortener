package storage

type Abstract interface {
	Push(string, string)
	Get(string) (string, error)
}
