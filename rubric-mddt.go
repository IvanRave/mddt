// API definitions for serv rubrics, groups, areas and scopes
package mddt

// ServRubric with locale
type ServRubric struct {
	Id          string  `json:"id" summary:"like hairCut, nailExt"`
	ServGroupId string  `json:"serv_group_id" summary:"id of a parent group"`
	LocaleId    string  `json:"lang" summary:"two-letter language"`
	Name        *string `json:"name" summary:"name in current language"`
	Description *string `json:"description" summary:"rubric description"`
	Tag1        *string `json:"tag1" summary:"primary tag"`
	Tag2        *string `json:"tag2" summary:"secondary tag"`
}

// ServGroup with locale
type ServGroup struct {
	Id            string        `json:"id" summary:"like 'nail, hair'"`
	ServAreaId    string        `json:"serv_area_id" summary:"like 'beauty'"`
	LocaleId      string        `json:"lang" summary:"two-letter id"`
	Name          *string       `json:"name" summary:"in selected language"`
	Description   *string       `json:"description" summary:"group description"`
	Tag1          *string       `json:"tag1" summary:"primary tag"`
	Tag2          *string       `json:"tag2" summary:"secondary tag"`
	ArrServRubric []*ServRubric `json:"arr_serv_rubric" summary:"rubrics from current group"`
}

// ServArea: no locales for areas, just ids
type ServArea struct {
	Id           string       `json:"id" summary:"like beauty, health"`
	ArrServGroup []*ServGroup `json:"arr_serv_group" summary:"groups of this area"`
}

// ServScope: to send to client: this scope may contains groups from different areas
type ServScope struct {
	Total        int32        `json:"total" summary:"total number of items"`
	ArrServGroup []*ServGroup `json:"arr_serv_group" summary:"groups with rubrics"`
}
