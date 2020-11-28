package main

import (
	"fmt"

	"getitqec.com/server/user/pkg/commons/normal_dist"
	"getitqec.com/server/user/pkg/dto"
	"getitqec.com/server/user/pkg/model"
)

func main() {
	// // NVDA
	// LNetReceivables := 1657000.0
	// PNetReceivables := 1424000.0
	// LSales := 10918000.0
	// PSales := 11716000.0
	// LCostOfGoodsSold := 4150000.0
	// PCostOfGoodsSold := 4545000.0
	// LCurrentAssets := 13690000.0
	// PCurrentAssets := 10557000.0
	// LPropertyPlantEquipment := 3303000.0
	// PPropertyPlantEquipment := 2171000.0
	// LSecurities := 10897000.0
	// PSecurities := 7422000.0
	// LTotalAssets := 17315000.0
	// PTotalAssets := 13292000.0
	// LDepreciation := 1011000.0
	// PDepreciation := 767000.0
	// LSalesGeneralAndAdministrativeExpense := 1093000.0
	// PSalesGeneralAndAdministrativeExpense := 991000.0
	// LCurrentLiabilities := 1784000.0
	// PCurrentLiabilities := 1329000.0
	// LTotalLongTermDebt := 3327000.0
	// PTotalLongTermDebt := 2621000.0
	// LContinuingOperationsIncome := 2796000.0
	// LOperationsCashFlow := 4761000.0

	// BMScore := model.BMScore(LNetReceivables, PNetReceivables, LSales, PSales, LCostOfGoodsSold, PCostOfGoodsSold, LCurrentAssets, PCurrentAssets, LPropertyPlantEquipment, PPropertyPlantEquipment, LSecurities, PSecurities, LTotalAssets, PTotalAssets, LDepreciation, PDepreciation, LSalesGeneralAndAdministrativeExpense, PSalesGeneralAndAdministrativeExpense, LCurrentLiabilities, PCurrentLiabilities, LTotalLongTermDebt, PTotalLongTermDebt, LContinuingOperationsIncome, LOperationsCashFlow)
	// // LTotalLiabilities := 511000.0
	// // LNetIncome := 2796000.0
	// // PNetIncome := 4141000.0

	// println(BMScore)

	// // VBHLF
	// LTotalAssets1 := 6468200.0
	// LTotalLiabilities1 := 5849300.0
	// GrossNationalProductPriceIndex := 4917997000000.0
	// LCurrentAssets1 := 2165400.0
	// LCurrentLiabilities1 := 3236900.0
	// LNetIncome1 := -349100.0
	// PNetIncome1 := -681000.0
	// Y := 1.0
	// OOScore := model.OOScore(LTotalAssets1, LTotalLiabilities1, GrossNationalProductPriceIndex, LCurrentAssets1, LCurrentLiabilities1, LNetIncome1, PNetIncome1, 4761000, Y)

	// println(OOScore)

	norm := normal_dist.New(0, 1)
	fmt.Println(normal_dist.Cdf(0.5, 0, 1))
	fmt.Println(norm.Cdf(-1))
	fmt.Println(norm.Cdf(0))

	fmt.Println(model.BMScore(&dto.FinancialReport{}, &dto.FinancialReport{}))
}
