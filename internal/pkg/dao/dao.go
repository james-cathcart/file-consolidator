package dao

type DAO interface {
	Save(hash string, path string) error
	Delete(hash string) error
}
