package dao

import (
	"context"

	"getitqec.com/server/user/pkg/commons"
	"getitqec.com/server/user/pkg/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserDAO details
type UserDAO struct {
	// db      commons.DB
	mongodb commons.MongoDB
}

// InitUserDAO function
func InitUserDAO(m commons.MongoDB) IUserDAO {
	return &UserDAO{mongodb: m}
}

// Create function
func (v *UserDAO) Create(ctx context.Context, user *dto.User) error {
	// data, err := dto.UserToAttributeMap(user)
	// if err != nil {
	// 	logger.Log.Error(fmt.Sprintf("UserDAO dto convert : %v", err))
	// 	return err
	// }
	// // get checking
	// err = v.db.Write("User", data)
	// if err != nil {
	// 	return err
	// }
	// return nil
	option := options.Find()
	option = option.SetLimit(1)
	c, err := v.mongodb.Client().Database(commons.UserDatabase).Collection(commons.UserColection).CountDocuments(ctx, bson.D{{Key: "_id", Value: user.UserId}})
	if err != nil {
		return err
	}
	if c == 0 {
		return v.mongodb.Create(ctx, commons.UserDatabase, commons.UserColection, user)
	}
	return commons.UserAlreadyExist
}

// Get function
func (v *UserDAO) Get(ctx context.Context, id string) (*dto.User, error) {
	// data, err := v.db.Get("User", "_id", id)
	// if err != nil {
	// 	return nil, err
	// }
	// return dto.AttributeMapToUser(data)
	result := v.mongodb.Read(ctx, commons.UserDatabase, commons.UserColection, bson.D{{Key: "_id", Value: id}})
	if result.Err() != nil {
		return nil, result.Err()
	}
	user := &dto.User{}
	err := result.Decode(user)
	return user, err
}

func (v *UserDAO) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	// data, err := v.db.Get("User", "_id", id)
	// if err != nil {
	// 	return nil, err
	// }
	// return dto.AttributeMapToUser(data)
	result := v.mongodb.Read(ctx, commons.UserDatabase, commons.UserColection, bson.D{{Key: "email", Value: email}})
	if result.Err() != nil {
		return nil, result.Err()
	}
	user := &dto.User{}
	err := result.Decode(user)
	return user, err
}

func (v *UserDAO) Update(ctx context.Context, user *dto.User) error {
	return v.mongodb.Update(ctx, commons.UserDatabase, commons.UserColection, bson.D{{Key: "_id", Value: user.UserId}}, bson.D{{"$set", user}})
}

// func (v *UserDAO) Delete(ctx context.Context, id string) (*dto.User, error) {
// 	// data, err := v.db.Delete("User", "_id", id)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// if data == nil {
// 	// 	return nil, commons.ErrExpiredToken
// 	// }
// 	// return dto.AttributeMapToUser(data)
// 	result := v.mongodb.Delete(ctx, commons.UserDatabase, commons.UserColection, bson.D{{Key: "_id", Value: id}})
// 	if result.Err() != nil {
// 		return nil, result.Err()
// 	}
// 	user := &dto.User{}
// 	err := result.Decode(user)
// 	return user, err
// }

// func (v *UserDAO) CheckEmailUsed(ctx context.Context, email string) (bool, error) {
// 	count1, err1 := v.mongodb.Client().Database(commons.UserDatabase).Collection(commons.UserColection).CountDocuments(ctx, bson.D{{"email", email}})
// 	if err1 != nil {
// 		return false, err1
// 	}
// 	count2, err2 := v.mongodb.Client().Database(commons.UserDatabase).Collection(commons.UserColection).CountDocuments(ctx, bson.D{{"newEmail", email}})
// 	if err2 != nil {
// 		return false, err1
// 	}
// 	return count1+count2 > 0, nil
// }

// func (v *UserDAO) CheckMobileUsed(ctx context.Context, mobile string) (bool, error) {
// 	count1, err1 := v.mongodb.Client().Database(commons.UserDatabase).Collection(commons.UserColection).CountDocuments(ctx, bson.D{{"mobile", mobile}})
// 	if err1 != nil {
// 		return false, err1
// 	}
// 	count2, err2 := v.mongodb.Client().Database(commons.UserDatabase).Collection(commons.UserColection).CountDocuments(ctx, bson.D{{"newMobile", mobile}})
// 	if err2 != nil {
// 		return false, err1
// 	}
// 	return count1+count2 > 0, nil
// }

// // admin use for getting total number of users
// func (v *UserDAO) CountTotal(ctx context.Context) (int64, error) {
// 	return v.mongodb.Client().Database(commons.UserDatabase).Collection(commons.UserColection).CountDocuments(ctx, bson.D{{}})
// }

// // ignore pagination
// func (v *UserDAO) QueryBirthday(ctx context.Context, m time.Month, day int, from int, to int) (int64, []*dto.User, error) {
// 	// pipeline := mongo.Pipeline{}
// 	// pipeline = append(pipeline, bson.D{{
// 	// 	Key: "$project",
// 	// 	Value: bson.M{
// 	// 		"bMonth": bson.M{"$month": bson.M{
// 	// 			"date":     "$birthday",
// 	// 			"timezone": "+0800",
// 	// 		}},
// 	// 		"bDay": bson.M{"$dayOfMonth": bson.M{
// 	// 			"date":     "$birthday",
// 	// 			"timezone": "+0800",
// 	// 		}},
// 	// 	},
// 	// }})
// 	// {
// 	// 	date: ISODate("2000-01-01T00:00:00Z"),
// 	// 	timezone: "-0500"
// 	// }
// 	// v.mongodb.Client().Database(commons.UserTable).Collection(commons.UserColection).Aggregate(ctx, pipeline)
// 	count, raws, err := v.mongodb.Query(ctx, commons.UserDatabase, commons.UserColection, nil, nil, &commons.FilterData{Value: bson.M{
// 		"$expr": bson.M{
// 			"$and": bson.A{
// 				bson.M{"$eq": bson.A{bson.M{"$dayOfMonth": bson.M{
// 					"date":     "$birthday",
// 					"timezone": "+0800",
// 				}}, day}},
// 				bson.M{"$eq": bson.A{bson.M{"$month": bson.M{
// 					"date":     "$birthday",
// 					"timezone": "+0800",
// 				}}, m}},
// 			},
// 		},
// 	}})
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	users := []*dto.User{}
// 	for _, raw := range raws {
// 		user := &dto.User{}
// 		err = bson.Unmarshal(*raw, user)
// 		if err != nil {
// 			return 0, nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return count, users, nil
// }

// func (v *UserDAO) QueryBirthMonth(ctx context.Context, m time.Month, from int, to int) (int64, []*dto.User, error) {
// 	count, raws, err := v.mongodb.Query(ctx, commons.UserDatabase, commons.UserColection, nil, nil, &commons.FilterData{Value: bson.M{
// 		"$expr": bson.M{"$eq": bson.A{bson.M{"$month": bson.M{
// 			"date":     "$birthday",
// 			"timezone": "+0800",
// 		}}, m}},
// 	}})
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	users := []*dto.User{}
// 	for _, raw := range raws {
// 		user := &dto.User{}
// 		err = bson.Unmarshal(*raw, user)
// 		if err != nil {
// 			return 0, nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return count, users, nil
// }

// // birthday d1 <= t <= d2
// func (v *UserDAO) QueryBirthdayBetween(ctx context.Context, startMonth time.Month, startDay int, endMonth time.Month, endDay int, from int, to int) (int64, []*dto.User, error) {
// 	pipeline := mongo.Pipeline{}
// 	pipeline = append(pipeline, bson.D{{
// 		Key: "$project",
// 		Value: bson.M{
// 			"name":  true,
// 			"email": true,
// 			"bMonth": bson.M{"$month": bson.M{
// 				"date":     "$birthday",
// 				"timezone": "+0800",
// 			}},
// 			"bDay": bson.M{"$dayOfMonth": bson.M{
// 				"date":     "$birthday",
// 				"timezone": "+0800",
// 			}},
// 		},
// 	}})
// 	pipeline = append(pipeline, bson.D{{
// 		Key: "$match",
// 		Value: bson.M{
// 			"$or": bson.A{
// 				// equal startMonth
// 				bson.M{
// 					"$and": bson.A{
// 						bson.M{
// 							"bMonth": bson.M{"$gte": startDay},
// 						},
// 						bson.M{
// 							"bMonth": startMonth,
// 						},
// 					},
// 				},
// 				// between
// 				bson.M{
// 					"$and": bson.A{
// 						bson.M{
// 							"bMonth": bson.M{"$gt": startMonth},
// 						},
// 						bson.M{
// 							"bMonth": bson.M{"$lt": endMonth},
// 						},
// 					},
// 				},
// 				// equal endMonth
// 				bson.M{
// 					"$and": bson.A{
// 						bson.M{
// 							"bMonth": bson.M{"$lte": endDay},
// 						},
// 						bson.M{
// 							"bMonth": endMonth,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}})
// 	cursor, err := v.mongodb.Client().Database(commons.UserDatabase).Collection(commons.UserColection).Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	defer cursor.Close(ctx)
// 	users := []*dto.User{}
// 	for cursor.Next(ctx) {
// 		var user *dto.User
// 		if cursor.Decode(user) != nil {
// 			return 0, nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return int64(len(users)), users, nil
// }
