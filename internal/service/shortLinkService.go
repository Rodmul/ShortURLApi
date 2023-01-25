package service

import (
	"ShortLink/internal/config"
	"ShortLink/internal/store"
	"ShortLink/internal/store/db"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"net/http"
)

type Service struct {
	Logger *zap.SugaredLogger
	conf   *config.Config
	router *mux.Router
	store  *store.Store
}

func NewService(cfg *config.Config) *Service {
	logger := NewLogger()
	srv := &Service{Logger: logger, conf: cfg, router: mux.NewRouter()}
	srv.registerClientHandlers()

	return srv
}

func (srv *Service) Start() error {
	if srv.conf.DBStorageMode {
		database, err := db.NewDB(srv.conf, srv.Logger.Desugar())
		if err != nil {
			srv.Logger.Fatal("Can't initialize connection to database", zap.Error(err))

			return fmt.Errorf("failed to initialize connection to database: %w", err)
		}
		checkMigrationVersion(srv, database)
		srv.store = store.New(database, srv.Logger)
	} else {
		srv.store = store.NewIM(srv.Logger)
		srv.Logger.Info("In-memory storage using")
	}

	return fmt.Errorf("failed to listen and serve: %w", http.ListenAndServe(":"+srv.conf.Listen.Port, srv.router))
}

func checkMigrationVersion(srv *Service, db *sqlx.DB) {
	type MigrationSchema struct {
		Version int  `db:"version"`
		Dirty   bool `db:"dirty"`
	}
	result := db.QueryRowx("SELECT version, dirty FROM short_link.public.schema_migrations;")
	if result.Err() != nil {
		srv.Logger.Fatal("failed to connect to database")
	}
	migration := &MigrationSchema{}
	err := result.StructScan(migration)
	if err != nil {
		srv.Logger.Fatal("failed to scan schema migrations:", zap.Error(err))
	}
	if migration.Dirty {
		srv.Logger.Fatal("schema migration is in dirty mode")
	}
	if srv.conf.VersionDB != migration.Version {
		srv.Logger.Fatalf("Mismatched db versions. Expected: %d, got: %d", srv.conf.VersionDB, migration.Version)
	}
}
