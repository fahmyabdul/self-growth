package komoditas

import (
	"strings"

	"github.com/fahmyabdul/efishery-task/fetch-app/configs"
)

type Komoditas struct {
	Uuid         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
	PriceUSD     string `json:"price_usd"`
}

type KomoditasAggregateContent struct {
	WeekNumber int            `json:"week_number"`
	Size       AggregateGroup `json:"size"`
	Price      AggregateGroup `json:"price"`
}

type AggregateGroup struct {
	Collection []int   `json:"collection"`
	Min        int     `json:"min"`
	Max        int     `json:"max"`
	Median     int     `json:"median"`
	Avg        float64 `json:"avg"`
}

func (p *Komoditas) TableName() string {
	tablename := "<schema>.t_komoditas"
	return strings.ReplaceAll(tablename, "<schema>", configs.Properties.Databases.Postgre.Schema)
}

func (p *Komoditas) KeyRedis() string {
	return "data:komoditas"
}
