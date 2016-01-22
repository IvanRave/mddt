package mddt

type MasterServ struct {

	// Use this field to change lo name to our id
	// it is filled during convertation from loServ to appServ
	ServRubricId    string      `json:"serv_rubric_id" summary:"rubric id from serv_rubric"`
	Id              int32       `json:"id" summary:"autoincrement id"`
	MasterProfileId int32       `json:"master_profile_id" summary:"master id"`
	PriceFrom       *int32      `json:"price_from" summary:"min price"`
	PriceTo         *int32      `json:"price_to" summary:"max price"`
	CurrencyId      *string     `json:"currency_id" summary:"currency for a price value, 3-letter id"`
	TimeFrom        *int32      `json:"time_from" summary:"min time for work, in seconds"`
	TimeTo          *int32      `json:"time_to" summary:"max time for work, in seconds"`
	IsOut           bool        `json:"is_out" summary:"whether the service provided out of home"`
	ArrServWork     []*ServWork `json:"arr_serv_work,omitempty" summary:"portfolio"`
	MasterProfile   *MasterProfile `json:"master_profile,omitempty" summary:"parent: master profile"`
}

