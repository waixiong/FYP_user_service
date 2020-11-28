package dao

import (
	"context"
	"fmt"

	"getitqec.com/server/user/pkg/commons"
	"getitqec.com/server/user/pkg/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StockConfigDAO struct {
	mongodb commons.MongoDB
}

// InitPortfolioDAO function
func InitStockConfigDAO(m commons.MongoDB) IStockConfigDAO {
	r, err := m.Client().Database(commons.StockConfigDatabase).Collection(commons.StockConfigCollection).Indexes().CreateOne(
		context.TODO(),
		mongo.IndexModel{
			Keys: bson.M{
				"_id":   "text",
				"alias": "text",
				"name":  "text",
			},
			Options: options.Index().SetName("QuerySearch"),
		},
	)
	if err != nil {
		fmt.Println("\tIndex existed")
	} else {
		fmt.Println("\tIndex " + r)
	}
	return &StockConfigDAO{mongodb: m}
}

// Get function
func (v *StockConfigDAO) Get(ctx context.Context, code string) (*dto.StockConfig, error) {
	result := v.mongodb.Read(ctx, commons.StockConfigDatabase, commons.StockConfigCollection, bson.D{{Key: "_id", Value: code}})
	if result.Err() != nil {
		return nil, result.Err()
	}
	stock := &dto.StockConfig{}
	err := result.Decode(stock)
	return stock, err
}

func (v *StockConfigDAO) Upsert(ctx context.Context, stock *dto.StockConfig) error {
	return v.mongodb.Upsert(ctx, commons.StockConfigDatabase, commons.StockConfigCollection, bson.D{{Key: "_id", Value: stock.Code}}, bson.D{{"$set", stock}})
}

func (v *StockConfigDAO) Delete(ctx context.Context, id string) (*dto.StockConfig, error) {
	result := v.mongodb.Delete(ctx, commons.StockConfigDatabase, commons.StockConfigCollection, bson.D{{Key: "_id", Value: id}})
	if result.Err() != nil {
		return nil, result.Err()
	}
	stock := &dto.StockConfig{}
	err := result.Decode(stock)
	return stock, err
}

// Query function
func (v *StockConfigDAO) QueryByText(ctx context.Context, text string) ([]*dto.StockConfig, error) {
	collection := v.mongodb.Client().Database(commons.StockConfigDatabase).Collection(commons.StockConfigCollection)

	var cursor *mongo.Cursor
	var err error
	// var raws []*bson.Raw

	findOptions := options.Find()
	// set sorter base on score
	findOptions.SetProjection(bson.D{{Key: "score", Value: bson.M{"$meta": "textScore"}}})
	findOptions.SetSort(bson.D{{Key: "score", Value: bson.M{"$meta": "textScore"}}})

	// set filter
	// "/{text}/gi"
	pattern := "/" + text + "/gi"
	fmt.Println(pattern)

	cursor, err = collection.Find(
		ctx,
		bson.M{
			"$text": bson.M{
				"$search": text,
				// bson.M{
				// 	"$regex": pattern,
				// },
			},
		},
		findOptions,
	)
	// cursor, err = collection.Find(
	// 	ctx,
	// 	bson.M{
	// 		"$or": bson.A{
	// 			bson.M{
	// 				"_id": bson.M{
	// 					"$regex": primitive.Regex{
	// 						Pattern: text,
	// 					},
	// 				},
	// 			},
	// 			bson.M{
	// 				"alias": bson.M{
	// 					"$regex": primitive.Regex{
	// 						Pattern: text,
	// 					},
	// 				},
	// 			},
	// 			bson.M{
	// 				"name": bson.M{
	// 					"$regex": primitive.Regex{
	// 						Pattern: text,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	findOptions,
	// )

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	stocks := []*dto.StockConfig{}
	for cursor.Next(ctx) {
		// raw := &bson.Raw{}
		// if err = cursor.Decode(&raw); err != nil {
		// 	return 0, nil, err
		// }
		stock := &dto.StockConfig{}
		err := cursor.Decode(stock)
		// raw := cursor.Current
		// raws = append(raws, &raw)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, err
}
