package model

import (

	//pb "./proto"

	// "github.com/syndtr/goleveldb/leveldb"

	"context"

	"getitqec.com/server/user/pkg/dto"
)

func (m *UserModel) GetUser(ctx context.Context, id string) (*dto.User, error) {
	user, err := m.UserDAO.Get(ctx, id)
	return user, err
}

func (m *UserModel) UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	err := m.UserDAO.Update(ctx, user)
	return user, err
}
