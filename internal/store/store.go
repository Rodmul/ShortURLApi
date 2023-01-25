package store

import (
	"ShortLink/internal/genErr"
	"ShortLink/internal/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Store struct {
	IMStorage     []models.URLStorage
	db            *DB
	logger        *zap.SugaredLogger
	Itx           TX
	urlRepository Urler
}

type TX interface {
	Rollback(tx *sqlx.Tx, err error) error
	BeginTransaction() (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
}

func NewIM(l *zap.SugaredLogger) *Store {
	return &Store{
		IMStorage: []models.URLStorage{},
		logger:    l,
	}
}

func New(db *sqlx.DB, l *zap.SugaredLogger) *Store {
	return &Store{
		db:     &DB{db},
		logger: l,
	}
}

func NewTest(tx TX, ur Urler) *Store {
	return &Store{
		Itx:           tx,
		urlRepository: ur,
	}
}

func (s *Store) Url() Urler {
	if s.urlRepository != nil {
		return s.urlRepository
	}

	s.urlRepository = &UrlRepository{
		store: s,
	}

	return s.urlRepository
}

func (s *Store) BeginTransaction() (*sqlx.Tx, error) {
	if s.Itx != nil {
		return s.Itx.BeginTransaction()
	} else {
		tx, err := s.db.sqlxDB.Beginx()
		if err != nil {
			s.logger.Error(ErrTransactionCreationFailed)

			return nil, genErr.NewError(err, ErrTransactionCreationFailed)
		}

		return tx, nil
	}
}

func (s *Store) CommitTransaction(tx *sqlx.Tx) error {
	if s.Itx != nil {
		return s.Itx.CommitTransaction(tx)
	} else {
		if err := tx.Commit(); err != nil {
			// s.logger.Errorf(wErr.WrapError(ErrTransactionCommitFailed.Error(), err).Error())
			s.logger.Error(err)
			if err = s.Rollback(tx, err); err != nil {
				return err
			}
		}

		return nil
	}
}

func (s *Store) Rollback(tx *sqlx.Tx, err error) error {
	if s.Itx != nil {
		return s.Itx.Rollback(tx, err)
	} else {
		if tx == nil {
			return err
		}

		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return genErr.NewError(err, ErrTransactionCreationFailed, nil)
		}

		return genErr.NewError(err, ErrTransactionFailed, nil)
	}
}
