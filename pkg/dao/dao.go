package dao

import (
	"context"

	"getitqec.com/server/user/pkg/dto"
)

// IUserDAO APIs
type IUserDAO interface {
	// Set(*dto.User) error
	Create(ctx context.Context, user *dto.User) error
	Get(ctx context.Context, id string) (*dto.User, error)
	Update(ctx context.Context, user *dto.User) error
	GetByEmail(ctx context.Context, id string) (*dto.User, error)
}

// IPortfolioDAO APIs
type IPortfolioDAO interface {
	Create(ctx context.Context, portfolio *dto.Portfolio) error
	Get(ctx context.Context, id string) (*dto.Portfolio, error)
	Update(ctx context.Context, portfolio *dto.Portfolio) error
	Delete(ctx context.Context, id string) (*dto.Portfolio, error)
	QueryByUser(ctx context.Context, userId string) ([]*dto.Portfolio, error)
}

// IPortfolioDAO APIs
type IStockConfigDAO interface {
	Get(ctx context.Context, id string) (*dto.StockConfig, error)
	Upsert(ctx context.Context, stock *dto.StockConfig) error
	Delete(ctx context.Context, id string) (*dto.StockConfig, error)
	QueryByText(ctx context.Context, text string) ([]*dto.StockConfig, error)
}

// type IUserDAO interface {
// 	// Set(*dto.User) error
// 	Create(ctx context.Context, user *dto.User) error
// 	Get(ctx context.Context, id string) (*dto.User, error)
// 	Update(ctx context.Context, user *dto.User) error
// 	Delete(ctx context.Context, id string) (*dto.User, error)

// 	Query(ctx context.Context, filter interface{}) ([]*dto.User, error)

// 	CheckEmailUsed(ctx context.Context, email string) (bool, error)
// 	CheckMobileUsed(ctx context.Context, mobile string) (bool, error)
// }

// type IEmailVDAO interface {
// 	// Set(*dto.User) error
// 	Create(ctx context.Context, req *dto.VerifyRequest) error
// 	Get(ctx context.Context, id string) (*dto.VerifyRequest, error)
// 	Update(ctx context.Context, req *dto.VerifyRequest) error
// 	Delete(ctx context.Context, id string) (*dto.VerifyRequest, error)

// 	// Query(ctx context.Context, filter interface{}) ([]*dto.VerifyRequest, error)
// }

// type IPhoneVDAO interface {
// 	// Set(*dto.User) error
// 	Create(ctx context.Context, req *dto.VerifyRequest) error
// 	Get(ctx context.Context, id string) (*dto.VerifyRequest, error)
// 	Update(ctx context.Context, req *dto.VerifyRequest) error
// 	Delete(ctx context.Context, id string) (*dto.VerifyRequest, error)

// 	// Query(ctx context.Context, filter interface{}) ([]*dto.VerifyRequest, error)
// }
