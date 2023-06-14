package data

import (
	"chatgpt-admin-server/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) Create(ctx context.Context, user *biz.User) error {
	return nil
}

func (u *userRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	return nil, nil
}

func (u *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	return nil, nil
}
