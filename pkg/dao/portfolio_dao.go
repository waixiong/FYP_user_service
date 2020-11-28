package dao

import (
	"context"

	"getitqec.com/server/user/pkg/commons"
	"getitqec.com/server/user/pkg/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PortfolioDAO details
type PortfolioDAO struct {
	mongodb commons.MongoDB
}

// InitPortfolioDAO function
func InitPortfolioDAO(m commons.MongoDB) IPortfolioDAO {
	return &PortfolioDAO{mongodb: m}
}

// Create function
func (v *PortfolioDAO) Create(ctx context.Context, portfolio *dto.Portfolio) error {
	option := options.Find()
	option = option.SetLimit(1)
	// c, err := v.mongodb.Client().Database(commons.PortfolioDatabase).Collection(commons.PortfolioCollection).CountDocuments(ctx, bson.D{{Key: "_id", Value: portfolio.PortfolioID}})
	// if err != nil {
	// 	return err
	// }
	// if c == 0 {
	// 	return v.mongodb.Create(ctx, commons.PortfolioDatabase, commons.PortfolioCollection, portfolio)
	// }
	// return commons.PortfolioAlreadyExist
	return v.mongodb.Create(ctx, commons.PortfolioDatabase, commons.PortfolioCollection, portfolio)
}

// Get function
func (v *PortfolioDAO) Get(ctx context.Context, id string) (*dto.Portfolio, error) {
	result := v.mongodb.Read(ctx, commons.PortfolioDatabase, commons.PortfolioCollection, bson.D{{Key: "_id", Value: id}})
	if result.Err() != nil {
		return nil, result.Err()
	}
	portfolio := &dto.Portfolio{}
	err := result.Decode(portfolio)
	return portfolio, err
}

func (v *PortfolioDAO) Update(ctx context.Context, portfolio *dto.Portfolio) error {
	return v.mongodb.Update(ctx, commons.PortfolioDatabase, commons.PortfolioCollection, bson.D{{Key: "_id", Value: portfolio.PortfolioID}}, bson.D{{"$set", portfolio}})
}

func (v *PortfolioDAO) Delete(ctx context.Context, id string) (*dto.Portfolio, error) {
	result := v.mongodb.Delete(ctx, commons.PortfolioDatabase, commons.PortfolioCollection, bson.D{{Key: "_id", Value: id}})
	if result.Err() != nil {
		return nil, result.Err()
	}
	portfolio := &dto.Portfolio{}
	err := result.Decode(portfolio)
	return portfolio, err
}

// Query function
func (v *PortfolioDAO) QueryByUser(ctx context.Context, userId string) ([]*dto.Portfolio, error) {
	// count, raws, err := v.mongodb.Query(ctx, commons.UserTable, commons.UserColection, nil, nil, &commons.FilterData{Item: field, Value: data})
	raws, err := v.mongodb.BatchRead(ctx, commons.PortfolioDatabase, commons.PortfolioCollection, bson.D{{Key: "uid", Value: userId}})
	if err != nil {
		return nil, err
	}
	portfolios := []*dto.Portfolio{}
	for _, raw := range raws {
		portfolio := &dto.Portfolio{}
		err = bson.Unmarshal(*raw, portfolio)
		if err != nil {
			return nil, err
		}
		portfolios = append(portfolios, portfolio)
	}
	return portfolios, err
}
