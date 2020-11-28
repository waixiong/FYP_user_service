package model

import (
	"context"
	"fmt"

	"getitqec.com/server/user/pkg/commons"
	"getitqec.com/server/user/pkg/dto"
)

func (m *UserModel) GoogleSignIn(ctx context.Context, idToken string, accessToken string) (*dto.User, bool, error) {
	commons.MalaysiaTimeNow()
	tokenInfo, userInfo, err := commons.VerifyGoogleIDToken(ctx, idToken, accessToken)
	if err != nil {
		return nil, false, err
	}
	fmt.Printf("UserId: %s\n", tokenInfo.UserId)
	// 111377874023438022187
	fmt.Printf("Audience: %s\n", tokenInfo.Audience)
	// 996344380155-alqg0crulhacekd7s45eha4fbgr5s5lg.apps.googleusercontent.com
	fmt.Printf("Expired: %d\n", tokenInfo.ExpiresIn)
	// 3599
	fmt.Printf("UserE: %s\n", userInfo.Email)
	fmt.Printf("UserI: %s\n", userInfo.Link) // GivenName, FamilyName
	// userInfo.Gender
	// userInfo.Link  //user profile img
	// userInfo.Locale

	// TODO: return user if exist, else ask user for additional info
	user, err := m.GetUser(ctx, tokenInfo.UserId)
	if err != nil {
		// register user
		fmt.Println("\tRegister User")
		user = &dto.User{
			UserId:   tokenInfo.UserId,
			Email:    tokenInfo.Email,
			UserName: userInfo.Name,
			Img:      userInfo.Picture,
		}
		err := m.UserDAO.Create(ctx, user)
		if err != nil {
			return nil, false, err
		}
		return user, false, nil
	}
	fmt.Println("\tSign In User")
	// update
	user.Img = userInfo.Picture
	user.UserName = userInfo.Name
	user.Email = userInfo.Email
	m.UpdateUser(ctx, user)
	return user, true, nil
}
