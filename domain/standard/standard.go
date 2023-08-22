package standard

import (
	"todo_app_mux/db"
	"todo_app_mux/domain"
)

type domainService struct {
	db db.Idb
}

func New(dbService db.Idb) domain.IDomainService {
	return &domainService{
		db: dbService,
	}
}
