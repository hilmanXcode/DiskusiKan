package memberships

import (
	"context"

	"github.com/hilmanXcode/DiskusiKan/internal/configs"
	"github.com/hilmanXcode/DiskusiKan/internal/model/memberships"
)

type membershiRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershiRepository
}

func NewService(cfg *configs.Config, membershipRepo membershiRepository) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
