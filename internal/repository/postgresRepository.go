package postgres

import (
	"context"

	"github.com/Greenfielder/backend2-go/internal/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGRepo struct {
	db *pgxpool.Pool
}

var _ storage.Repo = &PGRepo{}

func NewDB(ctx context.Context, connStr string) (*PGRepo, error) {
	return &PGRepo{}, nil
}

func (pg *PGRepo) Close() {
	pg.db.Close()
}

func (pg *PGRepo) CreateUser(ctx context.Context, u entity.User) error {
	return nil
}

func (pg *PGRepo) CreateGroup(ctx context.Context, g entity.Group) error {
	return nil
}

func (pg *PGRepo) AddToGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return nil
}

func (pg *PGRepo) RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return nil
}

func (pg *PGRepo) SearchUser(ctx context.Context, name string, gids []uuid.UUID) ([]entity.User, error) {
	users := make([]entity.User, 0)
	return users, nil
}

func (pg *PGRepo) SearchGroup(ctx context.Context, name string, uids []uuid.UUID) ([]entity.Group, error) {
	groups := make([]entity.Group, 0)
	return groups, nil
}

func getPGXPoolConfig(connStr string) (*pgxpool.Config, error) {
	cfg, _ := pgxpool.ParseConfig(connStr)
	return cfg, nil
}
