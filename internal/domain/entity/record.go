package entity

import domain "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/value_objects"

type Record struct {
	DR_NO        string     `json:"DR_NO"`
	DateRptd     string     `json:"DateRptd"`
	DATEOCC      string     `json:"DATEOCC"`
	TIMEOCC      string     `json:"TIMEOCC"`
	AREA         string     `json:"AREA"`
	AREANAME     string     `json:"AREANAME"`
	RptDistNo    string     `json:"RptDistNo"`
	Part12       string     `json:"Part12"`
	CrmCd        string     `json:"CrmCd"`
	CrmCdDesc    string     `json:"CrmCdDesc"`
	Mocodes      string     `json:"Mocodes"`
	VictAge      string     `json:"VictAge"`
	VictSex      domain.Sex `json:"VictSex"`
	VictDescent  string     `json:"VictDescent"`
	PremisCd     string     `json:"PremisCd"`
	PremisDesc   string     `json:"PremisDesc"`
	WeaponUsedCd string     `json:"WeaponUsedCd"`
	WeaponDesc   string     `json:"WeaponDesc"`
	Status       string     `json:"Status"`
	StatusDesc   string     `json:"StatusDesc"`
	CrmCd1       string     `json:"CrmCd1"`
	CrmCd2       string     `json:"CrmCd2"`
	CrmCd3       string     `json:"CrmCd3"`
	CrmCd4       string     `json:"CrmCd4"`
	LOCATION     string     `json:"LOCATION"`
	CrossStreet  string     `json:"CrossStreet"`
	LAT          string     `json:"LAT"`
	LON          string     `json:"LON"`
}
