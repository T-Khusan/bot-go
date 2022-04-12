package storage

import (
	"contact_service/storage/postgres"
	"contact_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Contact() repo.ContactRepoI
}

type storagePg struct {
	db      *sqlx.DB
	contact repo.ContactRepoI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:      db,
		contact: postgres.NewContactRepo(db),
	}
}

func (s *storagePg) Contact() repo.ContactRepoI {
	return s.contact // check nil
}
