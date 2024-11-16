package entity

import domain "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/value_objects"

type Record struct {
	DR_NO        string
	DateRptd     string
	DATEOCC      string
	TIMEOCC      string
	AREA         string
	AREANAME     string
	RptDistNo    string
	Part12       string
	CrmCd        string
	CrmCdDesc    string
	Mocodes      string
	VictAge      string
	VictSex      domain.Sex
	VictDescent  string
	PremisCd     string
	PremisDesc   string
	WeaponUsedCd string
	WeaponDesc   string
	Status       string
	StatusDesc   string
	CrmCd1       string
	CrmCd2       string
	CrmCd3       string
	CrmCd4       string
	LOCATION     string
	CrossStreet  string
	LAT          string
	LON          string
}
