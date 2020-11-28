package model

import (
	"fmt"
	"math"
	mathRand "math/rand"
	"time"
	"unsafe"

	"getitqec.com/server/user/pkg/commons/normal_dist"
	"getitqec.com/server/user/pkg/dto"
)

// ---------------- Random String Genarator ---------------- //
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const digits = "1234567890"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	digitIdxBits = 4                   // 6 bits to represent a letter index
	digitIdxMask = 1<<digitIdxBits - 1 // All 1-bits, as many as letterIdxBits
	digitIdxMax  = 10 / digitIdxBits   // # of letter indices fitting in 63 bits
)

var src = mathRand.NewSource(time.Now().UnixNano())

// source : https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func randDigitBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), digitIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), digitIdxMax
		}
		if idx := int(cache & digitIdxMask); idx < len(digits) {
			b[i] = digits[idx]
			i--
		}
		cache >>= digitIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// GenerateUserID generate user id
func GenerateUserID() string {
	return randStringBytesMaskImprSrcUnsafe(16)
}

func GenerateOTP() string {
	return randDigitBytesMaskImprSrcUnsafe(6)
}

func ExpPercentFunction(s float64) float64 {
	if s > 11 {
		return 1.0
	}
	return math.Exp(s) / (1 + math.Exp(s))
}

func ZFunction(s float64) float64 {
	return normal_dist.Cdf(s, 0, 1)
}

// BMScore  L = Latest, P = Previous
// func BMScore(LNetReceivables float64, PNetReceivables float64, LSales float64, PSales float64, LCostOfGoodsSold float64, PCostOfGoodsSold float64, LCurrentAssets float64, PCurrentAssets float64, LPropertyPlantEquipment float64, PPropertyPlantEquipment float64, LSecurities float64, PSecurities float64, LTotalAssets float64, PTotalAssets float64, LDepreciation float64, PDepreciation float64, LSalesGeneralAndAdministrativeExpense float64, PSalesGeneralAndAdministrativeExpense float64, LCurrentLiabilities float64, PCurrentLiabilities float64, LTotalLongTermDebt float64, PTotalLongTermDebt float64, LContinuingOperationsIncome float64, LOperationsCashFlow float64) float64 {
func BMScore(L *dto.FinancialReport, P *dto.FinancialReport) float64 {
	DaysSalesInReceivablesIndex := (L.NetReceivables / L.Sales) / (P.NetReceivables / P.Sales)

	GrossMarginIndex := ((P.Sales - P.CostOfSales) / P.Sales) / ((L.Sales - L.CostOfSales) / L.Sales)

	AssetQualityIndex := (1 - (L.CurrentAssets+L.PropertyPlantEquipment+L.Securities)/L.TotalAssets) / (1 - ((P.CurrentAssets + P.PropertyPlantEquipment + P.Securities) / P.TotalAssets))

	SalesGrowthIndex := L.Sales / P.Sales

	DepreciationIndex := (P.Depreciation / (P.PropertyPlantEquipment + P.Depreciation)) / (L.Depreciation / (L.PropertyPlantEquipment + L.Depreciation))

	SalesGeneralAndAdministrativeExpensesIndex := (L.SalesGeneralAndAdministrativeExpense / L.Sales) / (P.SalesGeneralAndAdministrativeExpense / P.Sales)

	LeverageIndex := ((L.CurrentLiabilities + L.NonCurrentLiabilities) / L.TotalAssets) / ((P.CurrentLiabilities + P.NonCurrentLiabilities) / P.TotalAssets)

	TotalAccrualsToTotalAssets := (L.ContinuingOperationsIncome - L.OperationsCashFlow) / L.TotalAssets

	return (-4.84 + 0.92*DaysSalesInReceivablesIndex + 0.528*GrossMarginIndex + 0.404*AssetQualityIndex + 0.892*SalesGrowthIndex + 0.115*DepreciationIndex - 0.172*SalesGeneralAndAdministrativeExpensesIndex + 4.679*TotalAccrualsToTotalAssets - 0.327*LeverageIndex)
}

// OOScore L = Latest, P = Previous, Y = 1 if there is a net loss for the last 2 years, else 0
// func OOScore(LTotalAssets float64, LTotalLiabilities float64, GrossNationalProductPriceIndex float64, LCurrentAssets float64, LCurrentLiabilities float64, LNetIncome float64, PNetIncome float64, LFundsFromOperations float64, Y float64) float64 {
func OOScore(L *dto.FinancialReport, P *dto.FinancialReport, GrossNationalProductPriceIndex float64) float64 {
	var X float64
	if L.TotalLiabilities > L.TotalAssets {
		X = 1
	} else {
		X = 0
	}

	LWorkingCapital := L.CurrentAssets - L.CurrentLiabilities

	score := (-1.32 - 0.407*(math.Log(L.TotalAssets/GrossNationalProductPriceIndex)) + 6.03*L.TotalLiabilities/L.TotalAssets - 1.43*LWorkingCapital/L.TotalAssets + 0.0757*L.CurrentLiabilities/L.CurrentAssets - 1.72*X - 2.37*(L.NetIncome/L.TotalAssets) - 1.83*(L.OperatingIncome/L.TotalLiabilities) + 0.285*L.Y - 0.521*(L.NetIncome-P.NetIncome)/(math.Abs(L.NetIncome)+math.Abs(P.NetIncome)))
	fmt.Print("OScore: ")
	fmt.Println(score)
	fmt.Println(-0.407 * (math.Log(L.TotalAssets / GrossNationalProductPriceIndex)))
	fmt.Println(6.03 * L.TotalLiabilities / L.TotalAssets)
	fmt.Println(-1.43 * LWorkingCapital / L.TotalAssets)
	fmt.Println(0.0757 * L.CurrentLiabilities / L.CurrentAssets)
	fmt.Println(-1.72 * X)
	fmt.Println(-2.37 * (L.NetIncome / L.TotalAssets))
	fmt.Println(-1.83 * (L.OperatingIncome / L.TotalLiabilities))
	fmt.Println(0.285 * L.Y)
	fmt.Println(-0.521 * (L.NetIncome - P.NetIncome) / (math.Abs(L.NetIncome) + math.Abs(P.NetIncome)))
	return score
}
