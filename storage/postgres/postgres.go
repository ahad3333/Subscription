package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"add/config"
	"add/storage"
)

type Store struct {
	db       *pgxpool.Pool
	user     *UserRepo
	subscription *SubscriptionRepo
	usersubscription *UserSubscriptionRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))

	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConn

	pool, err := pgxpool.ConnectConfig(ctx, config)

	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pool,
		user:     NewUserRepo(pool),
		subscription: NewSubscriptionRepo(pool),
		usersubscription: NewUserSubscriptionRepo(pool),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *Store) Subscription() storage.SubscriptionRepoI {

	if s.subscription == nil {
		s.subscription = NewSubscriptionRepo(s.db)
	}

	return s.subscription
}

func (s *Store) UserSubscription() storage.UserSubscriptionRepoI {

	if s.usersubscription == nil {
		s.usersubscription = NewUserSubscriptionRepo(s.db)
	}

	return s.usersubscription
}