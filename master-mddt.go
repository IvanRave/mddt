// mddt - models for masters
package mddt

type MasterAddress struct {
	MasterProfileId int32   `json:"master_profile_id" summary:"master id"`
	GeoDistrictId   int32   `json:"geo_district_id" summary:"id from geo districts"`
	Description     *string `json:"description" summary:"place, address, street etc."`
	Coords          *string `json:"coords" summary:"approximate coords of a place"`
}

type MasterProfile struct {
	Id            int32          `json:"id" summary:"master id"`
	Name          *string        `json:"name" summary:"nick name, or combined last+first names"`
	MainPhone     *string        `json:"main_phone" summary:"main phone number"`
	MainEmail     *string        `json:"main_email" summary:"main email"`
	Modified      int32          `json:"modified" summary:"when updated, unix timestamp"`
	Created       int32          `json:"created" summary:"when created, unix timestamp"`
	IsVerified    bool           `json:"is_verified" summary:"is correct information"`
	IsActive      bool           `json:"is_active" summary:"is not blocked in app"`
	ExpEntire     *int16         `json:"exp_entire" summary:"entire expirience, in months"`
	ExpPrivate    *int16         `json:"exp_private" summary:"exp as self employed, in months"`
	IsCompany     *bool   `json:"is_company" summary:"Whether the master is company"`
	NstgId        *string `json:"nstg_id" summary:"id of public profile from nstg"`
	UserName      *string `json:"username" summary:"unique username"`
	Website       *string `json:"website" summary:"website url"`
	Bio           *string `json:"bio" summary:"bio description"`

    Skype         *string `json:"skype"`
	Vkontakte     *string `json:"vkontakte"`
	Twitter       *string `json:"twitter"`
	Facebook      *string `json:"facebook"`
	Viber         *string `json:"viber"`
	Whatsapp      *string `json:"whatsapp"`
	Odnoklassniki *string `json:"odnoklassniki"`
	Landline      *string `json:"landline"`

	RefreshedWrk  *int32  `json:"refreshed_wrk" summary:"unix timestamp when master's works has been refreshed"`
	
	Gender        *string `json:"gender" summary:"f or m or nothing"`
	LastName      *string `json:"lname" summary:"last name"`
	FirstName     *string `json:"fname" summary:"first name"`
	MiddleName    *string `json:"mname" summary:"middle name"`
	Avatar        *string `json:"avatar" summary:"url to avatar image"`
	
	MasterAddress *MasterAddress `json:"master_address,omitempty" summary:"address with coords and place"`
	ArrMasterServ []*MasterServ  `json:"arr_master_serv,omitempty" summary:"master services"`
}
