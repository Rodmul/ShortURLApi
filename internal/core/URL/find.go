package URL

import (
	"ShortLink/internal/core"
	"ShortLink/internal/models"
	"ShortLink/internal/store"
)

func GetByShortLinkIM(s *store.Store, shortLink string) (*models.URLStorage, error) {
	for _, v := range s.IMStorage {
		if v.ShortURL == shortLink {
			return &v, nil
		}
	}

	return nil, core.ErrNotFound
}
