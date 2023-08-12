package standard

import (
	"todo_app_mux/db"
	"todo_app_mux/domain"
)

type domainService struct {
	db db.IMongoService
}

func New(dbService db.IMongoService) domain.IDomainService {
	return &domainService{
		db: dbService,
	}
}
