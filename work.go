package mddt

type ServWork struct {
	Id           int32   `json:"id" summary:"work id, autoincrement"`
	// MasterServId int32   `json:"master_serv_id" summary:"id of a master serv"`
	Name         *string `json:"name" summary:"localized name"`
	Description  *string `json:"description" summary:"localized description"`
	Created      *int32  `json:"created" summary:"when a work finished, utc timestamp"`
	ActualTime   *int32  `json:"actual_time" summary:"used time for work, in seconds"`
	ActualPrice  *int32  `json:"actual_price" summary:"actual price for work, in currency"`
	CurrencyId   *string `json:"currency_id" summary:"currency for a price value, 3-letter id"`
	PreviewImg   *string `json:"preview_img" summary:"url to a small image to show as a preview"`
	MainImg      *string `json:"main_img" summary:"url to a main image to show in portfolio"`
	MainVideo    *string `json:"main_video" summary:"url to video with work result"`
	SideLink     *string `json:"side_link" summary:"url to side page with work description"`
	MasterProfileId int32  `json:"master_profile_id" summary:"master profile id"`
}

type WorkRubric struct {
	ServRubricId string `json:"serv_rubric_id" summary:"id of a service rubric, like nailExt"`
	ServWorkId   int32  `json:"serv_work_id" summary:"id of a work"`
	Priority     int16  `json:"priority" summary:"rubric priority in a work"`
}
