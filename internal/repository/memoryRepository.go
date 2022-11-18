package memory

import (
	"context"

	"github.com/Greenfielder/backend2-go/internal/entity"
	"github.com/google/uuid"
)

type InMemory struct {
	users      []entity.User
	groups     []entity.Group
	memberShip map[uuid.UUID][]uuid.UUID
}

func (i InMemory) CreateUser(ctx context.Context, u entity.User) error {
	return nil
}

func (i InMemory) CreateGroup(ctx context.Context, g entity.Group) error {
	return nil
}

func (i InMemory) AddToGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return nil
}

func (i InMemory) RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return nil
}

func (i InMemory) SearchUser(ctx context.Context, name string, gids []uuid.UUID) ([]entity.User, error) {
	var users []entity.User
	return users, nil
}

func (i InMemory) SearchGroup(ctx context.Context, name string, uids []uuid.UUID) ([]entity.Group, error) {
	var groups []entity.Group
	return groups, nil
}

func NewMemory() *InMemory {
	return &InMemory{}
}
