package model

type ConsumptionRecode struct {
	ConsumptionRecodeId int
	Income              string
	Purpose             string
	PayType             PayType
	Price               float64
	Detail              string
	TradeTime           string
	RecodeTime          string
}
