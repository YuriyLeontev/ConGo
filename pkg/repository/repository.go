package repository

type AccountsList interface {
}

type Repository struct {
	AccountsList
}

func NewRepository() *Repository {
	return &Repository{}
}
