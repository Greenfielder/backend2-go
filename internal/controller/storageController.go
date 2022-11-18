package controller

import (
	"context"

	"github.com/Greenfielder/backend2-go/internal/entity"
	"github.com/google/uuid"
)

type Repo interface {
	CreateUser(ctx context.Context, u entity.User) error
	CreateGroup(ctx context.Context, u entity.Group) error
	AddToGroup(ctx context.Context, uid, gid uuid.UUID) error
	RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error
	SearchUser(ctx context.Context, name string, gids []uuid.UUID) ([]entity.User, error)
	SearchGroup(ctx context.Context, name string, uids []uuid.UUID) ([]entity.Group, error)
}

type DB struct {
	repo Repo
}

func NewDB(repo Repo) *DB {
	return &DB{
		repo: repo,
	}
}

func (db *DB) CreateUser(ctx context.Context, u entity.User) (*entity.User, error) {
	return &u, nil
}

func (db *DB) CreateGroup(ctx context.Context, g entity.Group) (*entity.Group, error) {
	return &g, nil
}

func (db *DB) AddToGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return db.repo.AddToGroup(ctx, uid, gid)
}

func (db *DB) RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return db.repo.RemoveFromGroup(ctx, uid, gid)
}

func (db *DB) SearchUser(ctx context.Context, name string, gids []uuid.UUID) ([]entity.User, error) {
	users, _ := db.repo.SearchUser(ctx, name, gids)
	return users, nil
}

func (db *DB) SearchGroup(ctx context.Context, name string, uids []uuid.UUID) ([]entity.Group, error) {
	groups, _ := db.repo.SearchGroup(ctx, name, uids)
	return groups, nil
}
