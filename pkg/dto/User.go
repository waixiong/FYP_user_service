package dto

// import "time"

// User Details
type User struct {
	UserId   string `json:"userId" bson:"_id"`
	UserName string `json:"userName" bson:"userName"`
	Email    string `json:"email" bson:"email"`
	Img      string `json:"img" bson:"img"`
}

// type User struct {
// 	UserId         string `json:"userId" bson:"_id"`
// 	RegisterMethod string `json:"registerMethod" bson:"registerMethod,omitempty"`
// 	Name           string `json:"name" bson:"name"`
// 	// LastName       string `json:"lastName" bson:"lastName,omitempty"`
// 	Email string `json:"email" bson:"email,omitempty"`
// 	// EmailVerify       bool      `json:"emailVerify" bson:"emailVerify"`
// 	Mobile string `json:"mobile" bson:"mobile,omitempty"`
// 	// MobileVerify      bool      `json:"mobileVerify" bson:"mobileVerify"`
// 	NewEmail  string   `json:"newEmail" bson:"newEmail,omitempty"`
// 	NewMobile string   `json:"newMobile" bson:"newMobile,omitempty"`
// 	Emails    []string `json:"-" bson:"emails"`
// 	Mobiles   []string `json:"-" bson:"mobiles"`
// 	IC        string   `json:"IC" bson:"IC,omitempty"`
// 	IC_Path   string   `json:"ICPath" bson:"ICPath,omitempty"`
// 	IC_Verify bool     `json:"ICVerify" bson:"ICVerify"`
// 	// BYear             uint32    `json:"bYear" bson:"bYear,omitempty"`
// 	// BMonth            uint32    `json:"bMonth" bson:"bMonth,omitempty"`
// 	// BDay              uint32    `json:"bDay" bson:"bDay,omitempty"`
// 	Birthday          time.Time `json:"birthday" bson:"birthday"`
// 	RegisterTimestamp time.Time `json:"registerTimestamp" bson:"registerTimestamp"`
// 	ImgPath           string    `json:"imgPath" bson:"imgPath"`
// }

// func UserToAttributeMap(user *User) (map[string]*dynamodb.AttributeValue, error) {
// 	return dynamodbattribute.MarshalMap(user)
// }

// func AttributeMapToUser(m map[string]*dynamodb.AttributeValue) (*User, error) {
// 	user := &User{}
// 	err := dynamodbattribute.UnmarshalMap(m, user)
// 	return user, err
// }

// func UserToBsonD(user *User) {

// }
