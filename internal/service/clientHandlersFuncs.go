package service

import (
	"ShortLink/internal/core"
	"ShortLink/internal/core/URL"
	"ShortLink/internal/core/link"
	"ShortLink/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const baseURL = "/"

func (srv *Service) registerClientHandlers() {
	srv.router.HandleFunc(baseURL+"saveUrl/", srv.handleSaveUrl()).Methods(http.MethodPost, http.MethodOptions)
	srv.router.HandleFunc(baseURL, srv.handleOriginUrl()).Methods(http.MethodGet, http.MethodOptions)
}

func (srv *Service) handleSaveUrl() http.HandlerFunc {
	type Request struct {
		LongUrl string `json:"longUrl"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
		var resultUrl string
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			srv.error(w, http.StatusBadRequest, err)

			return
		}

		ID, err := uuid.NewRandom()
		if err != nil {
			srv.error(w, http.StatusInternalServerError, err)
		}

		shortUrl := link.NewShortUrl([]rune(req.LongUrl))
		u := models.URLStorage{
			ID:       ID,
			LongURL:  req.LongUrl,
			ShortURL: shortUrl,
		}
		if srv.conf.DBStorageMode {
			err = URL.DBCreate(srv.store, u)
			if err != nil {
				srv.error(w, http.StatusInternalServerError, err)
			}
		} else {
			URL.IMCreate(srv.store, u)
		}

		resultUrl = fmt.Sprintf("http://%s:%s%s?shortUrl=%s", srv.conf.Listen.IP, srv.conf.Listen.Port, baseURL, shortUrl)
		srv.respond(w, http.StatusOK, resultUrl)
	}
}

func (srv *Service) handleOriginUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &models.URLStorage{}
		var err error
		shortUrl := r.URL.Query().Get("shortUrl")
		if srv.conf.DBStorageMode {
			u, err = srv.store.Url().GetByShortLink(shortUrl)
			if err != nil {
				if errors.Is(err, core.ErrNotFound) {
					srv.error(w, http.StatusNotFound, err)
					return
				}
				srv.error(w, http.StatusInternalServerError, err)

				return
			}
		} else {
			u, err = URL.GetByShortLinkIM(srv.store, shortUrl)
			if err != nil {
				if errors.Is(err, core.ErrNotFound) {
					srv.error(w, http.StatusNotFound, err)
					return
				}
				srv.error(w, http.StatusInternalServerError, err)

				return
			}
		}

		http.Redirect(w, r, u.LongURL, http.StatusFound)
	}
}
