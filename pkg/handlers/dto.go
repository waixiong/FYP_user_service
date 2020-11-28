package handlers

// // PBUser2User function
// func PBUser2User(req *pb.User) *dto.User {
// 	return &dto.User{
// 		UserId:    req.Id,
// 		Name:      req.Name,
// 		NewEmail:  req.Email,
// 		NewMobile: req.Phone,
// 		IC:        req.Ic,
// 		// BYear:     req.Year,
// 		// BMonth:    req.Month,
// 		// BDay:      req.Day,
// 		Birthday: commons.MilliToTime(req.Birthday),
// 		ImgPath:  req.ImgPath,
// 	}
// }

// // User2PBUser function
// func User2PBUser(user *dto.User) *pb.User {
// 	return &pb.User{
// 		Id:              user.UserId,
// 		Name:            user.Name,
// 		Email:           user.NewEmail,
// 		Phone:           user.NewMobile,
// 		Ic:              user.IC,
// 		EmailValidation: user.Email == user.NewEmail,
// 		PhoneValidation: user.Mobile == user.NewMobile,
// 		IcValidation:    user.IC_Verify,
// 		// Year:            user.BYear,
// 		// Month:           user.BMonth,
// 		// Day:             user.BDay,
// 		Birthday: commons.TimeToMilli(user.Birthday),
// 		ImgPath:  user.ImgPath,
// 	}
// }

// // UserHiding function
// func UserHiding(user *pb.User) *pb.User {
// 	// user.Day = 0
// 	// user.Month = 0
// 	// user.Year = 0
// 	user.Birthday = 0
// 	user.Email = ""
// 	user.Ic = ""
// 	user.Phone = ""
// 	return user
// }
