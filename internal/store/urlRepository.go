package store

import (
	"ShortLink/internal/core"
	"ShortLink/internal/genErr"
	"ShortLink/internal/models"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=urlRepository.go -destination=mocks/mock.go

type UrlRepository struct {
	store *Store
}

type Urler interface {
	CreateTx(tx *sqlx.Tx, u *models.URLStorage) error
	Create(u *models.URLStorage) error
	GetByShortLink(shortLink string) (*models.URLStorage, error)
	GetByShortLinkTx(tx *sqlx.Tx, shortLink string) (*models.URLStorage, error)
	GetByLongLink(longLink string) (*models.URLStorage, error)
	GetByLongLinkTx(tx *sqlx.Tx, longLink string) (*models.URLStorage, error)
}

func (r *UrlRepository) Create(u *models.URLStorage) error {
	return r.CreateTx(nil, u)
}

func (r *UrlRepository) CreateTx(tx *sqlx.Tx, u *models.URLStorage) error {
	if _, err := r.store.db.Exec(tx, "INSERT INTO short_link.public.url_data (id, long, short) VALUES ($1, $2, $3);",
		u.ID, u.LongURL, u.ShortURL); err != nil {
		return r.store.Rollback(tx, err)
	}

	return nil
}

func (r *UrlRepository) GetByShortLink(shortLink string) (*models.URLStorage, error) {
	return r.GetByShortLinkTx(nil, shortLink)
}

func (r *UrlRepository) GetByShortLinkTx(tx *sqlx.Tx, shortLink string) (*models.URLStorage, error) {
	u := models.URLStorage{}
	row := r.store.db.QueryRow(tx, "SELECT id, short, long FROM short_link.public.url_data WHERE short=$1", shortLink)

	err := row.StructScan(&u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = genErr.NewError(err, core.ErrNotFound, "shortUrl", shortLink)
			return nil, err
		}
		err = genErr.NewError(err, ErrScanStructFailed)

		return nil, r.store.Rollback(tx, err)
	}

	return &u, nil
}

func (r *UrlRepository) GetByLongLink(longLink string) (*models.URLStorage, error) {
	return r.GetByLongLinkTx(nil, longLink)
}

func (r *UrlRepository) GetByLongLinkTx(tx *sqlx.Tx, longLink string) (*models.URLStorage, error) {
	u := models.URLStorage{}
	row := r.store.db.QueryRow(tx, "SELECT id, short, long FROM short_link.public.url_data WHERE long=$1", longLink)

	err := row.StructScan(&u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, genErr.NewError(err, core.ErrNotFound)
		}
		err = genErr.NewError(err, ErrScanStructFailed)

		return nil, r.store.Rollback(tx, err)
	}

	return &u, nil
}
