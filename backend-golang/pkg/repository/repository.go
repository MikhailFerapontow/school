package repository

type Guardian interface {
}

type Repository struct {
	Guardian
}

func NewRepository() *Repository {
	return &Repository{}
}
