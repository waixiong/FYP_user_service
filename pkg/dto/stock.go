package dto

import "time"

// mongo text index on _id, name, alias
type StockConfig struct {
	Code  string  `json:"code" bson:"_id"`
	Name  string  `json:"name" bson:"name"`
	Alias string  `json:"alias" bson:"alias"`
	Price float64 `json:"price" bson:"price"`

	// Stock data
	EarningPerShare  float64 `json:"eps" bson:"eps"`
	DividendPerShare float64 `json:"dps" bson:"dps"`
	NetTangibleAsset float64 `json:"nta" bson:"nta"`
	Sector           string  `json:"sector" bson:"sector"`

	LastUpdated time.Time    `json:"lastUpdated" bson:"lastUpdated"`
	Indicator   *Indicator5N `json:"i5n" bson:"i5n"`

	L *FinancialReport `json:"l" bson:"l"`
	P *FinancialReport `json:"r" bson:"r"`
}

// sector
// https://www.bursamalaysia.com/trade/our_products_services/indices/bursa_malaysia_index_series

type FinancialReport struct {
	NetReceivables         float64 `json:"netReceivables" bson:"netReceivables"`
	CurrentAssets          float64 `json:"currentAssets" bson:"currentAssets"`
	PropertyPlantEquipment float64 `json:"propertyPlantEquipment" bson:"propertyPlantEquipment"`
	Securities             float64 `json:"securities" bson:"securities"`
	TotalAssets            float64 `json:"totalAssets" bson:"totalAssets"`
	CurrentLiabilities     float64 `json:"currentLiabilities" bson:"currentLiabilities"`
	NonCurrentLiabilities  float64 `json:"nonCurrentLiabilities" bson:"nonCurrentLiabilities"`
	TotalLiabilities       float64 `json:"totalLiabilities" bson:"totalLiabilities"`

	Sales                                float64 `json:"sales" bson:"sales"`
	CostOfSales                          float64 `json:"costOfSales" bson:"costOfSales"`
	Depreciation                         float64 `json:"depreciation" bson:"depreciation"`
	OperatingIncome                      float64 `json:"operatingIncome" bson:"operatingIncome"`
	NetIncome                            float64 `json:"netIncome" bson:"netIncome"`
	SalesGeneralAndAdministrativeExpense float64 `json:"salesGeneralAndAdministrativeExpense" bson:"salesGeneralAndAdministrativeExpense"`
	ContinuingOperationsIncome           float64 `json:"continuingOperationsIncome" bson:"continuingOperationsIncome"`
	Y                                    float64 `json:"y" bson:"y"` // 1 if a net loss for the last two years, 0 otherwise

	OperationsCashFlow float64 `json:"operationsCashFlow" bson:"operationsCashFlow"`
}
