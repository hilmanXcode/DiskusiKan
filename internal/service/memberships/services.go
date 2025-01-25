package memberships

import (
	"context"

	"github.com/hilmanXcode/DiskusiKan/internal/model/memberships"
)

type membershiRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	membershipRepo membershiRepository
}

func NewService(membershipRepo membershiRepository) *service {
	return &service{
		membershipRepo: membershipRepo,
	}
}
