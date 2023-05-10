package model

type Page struct {
	ConsumptionRecode []ConsumptionRecode
	User              *User
	PayType           []PayType
	Msg               string
}
