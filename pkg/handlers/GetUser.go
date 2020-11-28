package handlers

// // GetUserHandler details
// type GetUserHandler struct {
// 	Model model.UserModelI
// }

// // GetUser function
// func (s *GetUserHandler) GetUser(ctx context.Context, id string) (*pb.User, error) {
// 	userID, err := commons.GetUserID(ctx)
// 	if err != nil {
// 		logger.Log.Error(err.Error())
// 		return nil, err
// 	}

// 	if id != "" {
// 		// get other
// 		user, err := s.Model.GetUser(ctx, id)

// 		if err != nil {
// 			logger.Log.Error(err.Error())
// 			return nil, err
// 		}
// 		// check if need hide
// 		if commons.IsGetItService((userID)) {
// 			return User2PBUser(user), nil
// 		}
// 		return UserHiding(User2PBUser(user)), nil
// 	}

// 	user, err := s.Model.GetUser(ctx, userID)
// 	if err != nil {
// 		logger.Log.Error(err.Error())
// 		return nil, err
// 	}
// 	return User2PBUser(user), nil
// }

// // GetPortfolios function
// func (s *GetUserHandler) GetPortfolios(ctx context.Context, id string) (*[]pb.Portfolio, error) {
// 	userID, err := commons.GetUserID(ctx)
// 	if err != nil {
// 		logger.Log.Error(err.Error())
// 		return nil, err
// 	}

// 	if id != "" {
// 		// get other
// 		portfolios, err := s.Model.GetPorfolios(ctx, id)

// 		if err != nil {
// 			logger.Log.Error(err.Error())
// 			return nil, err
// 		}

// 		return portfolios, nil
// 		// // check if need hide
// 		// if commons.IsGetItService((userID)) {
// 		// 	return User2PBUser(user), nil
// 		// }
// 		// return UserHiding(User2PBUser(user)), nil
// 	}

// 	user, err := s.Model.GetUser(ctx, userID)
// 	if err != nil {
// 		logger.Log.Error(err.Error())
// 		return nil, err
// 	}
// 	return User2PBUser(user), nil
// }
