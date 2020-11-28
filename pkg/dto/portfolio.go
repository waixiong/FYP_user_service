package dto

import "time"

// Portfolio Details
type Portfolio struct {
	PortfolioID   string `json:"portolioId" bson:"_id"`
	UserId        string `json:"uid" bson:"uid"`
	PortfolioName string `json:"portfolioName" bson:"portfolioName"`

	/// Use either one
	Stocks map[string]float64 `json:"stocks" bson:"stocks"`
	// Stocks      []*Stock  `json:"stocks" bson:"stocks"`
	LastUpdated time.Time `json:"lastUpdated" bson:"lastUpdated"`

	Indicator *Indicator5N `json:"i5n" bson:"i5n"`
}

// Stock Details, if above use map, this is not exist
type Stock struct {
	Code      string  `json:"code" bson:"code"`
	Weightage float64 `json:"weightage" bson:"weightage"`
}

type Indicator5N struct {
	Alpha      float64 `json:"alpha" bson:"alpha"`
	Beta       float64 `json:"beta" bson:"beta"`
	Sharpe     float64 `json:"sharpeRatio" bson:"sharpe"`
	Sortino    float64 `json:"sortinoRatio" bson:"sortino"`
	Treynor    float64 `json:"treynorRatio" bson:"treynor"`
	Volatility float64 `json:"volatility" bson:"volatility"`

	// FA
	MScore float64 `json:"mScore" bson:"mScore"`
	OScore float64 `json:"oScore" bson:"oScore"`
}
