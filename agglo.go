// API definitions for geo units
package mddt

type GeoDistrict struct {
    Id          int32       `json:"id"`
    GeoRegionId int32       `json:"geo_region_id" summary:"id of a parent region"`
    LocalName   string      `json:"local_name" summary:"name in current country"`
    Extent      string      `json:"extent" summary:"polygon in WKT format"`
	Expanse     string      `json:"expanse" summary:"polygon in ((12.23,34.34),(23.65,54.76)) format"`
    Centroid    string      `json:"centroid" summary:"geo point of center"`
	Tag1        *string     `json:"tag1" summary:"primary tag"`	
	Tag2        *string     `json:"tag2" summary:"secondary tag"`
}

type GeoRegion struct {
    Id          int32       `json:"id"`
    GeoAggloId  int32       `json:"geo_agglo_id" summary:"id of a parent agglomeration"`
    LocalName   string      `json:"local_name" summary:"name in current country"`
    Extent      string      `json:"extent" summary:"polygon in WKT format"`
	Expanse     string      `json:"expanse" summary:"polygon in ((12.23,34.34),(23.65,54.76)) format"`
    Centroid    string      `json:"centroid" summary:"geo point of center"`
	Tag1        *string     `json:"tag1" summary:"primary tag"`	
	Tag2        *string     `json:"tag2" summary:"secondary tag"`
	//    ArrGeoDistrict  []*GeoDistrict  `json:"arr_geo_district" summary:"districts in this region"`
}

// rubric with locale
type GeoAgglo struct {
    Id          int32       `json:"id" summary:"agglomeration id"`
    InterCode   string      `json:"inter_code" summary:"code string, like paris"`
    LocalName   string      `json:"local_name" summary:"using country language"`
    Extent      string      `json:"extent" summary:"polygon in WKT format"`
	Expanse     string      `json:"expanse" summary:"polygon in ((12.23,34.34),(23.65,54.76)) format"`
    Centroid    string      `json:"centroid" summary:"geo point of center"`
    GeoId       string      `json:"geo_id"`
    CountryId   string      `json:"country_id" summary:"two-letter id"`
    LocaleId    string      `json:"locale_id" summary:"two-letter id"`
	Tag1        *string     `json:"tag1" summary:"primary tag"`	
	Tag2        *string     `json:"tag2" summary:"secondary tag"`
	//    ArrGeoRegion    []*GeoRegion    `json:"arr_geo_region,omitempty" summary:"regions in this agglomeration"`
}
