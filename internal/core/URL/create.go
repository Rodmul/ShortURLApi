package URL

import (
	"ShortLink/internal/core"
	"ShortLink/internal/genErr"
	"ShortLink/internal/models"
	"ShortLink/internal/store"
)

func DBCreate(s *store.Store, u models.URLStorage) error {
	if err := s.Url().Create(&u); err != nil {
		return genErr.NewError(err, core.ErrRepository)
	}

	return nil
}

func IMCreate(s *store.Store, u models.URLStorage) {
	s.IMStorage = append(s.IMStorage, u)
}
