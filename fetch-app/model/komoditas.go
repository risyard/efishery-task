package model

type Komoditas struct {
	ID         string `json:"uuid"`
	Barang     string `json:"komoditas"`
	Provinsi   string `json:"area_provinsi"`
	Kota       string `json:"area_kota"`
	Size       string `json:"size"`
	Price      string `json:"price"`
	Tanggal    string `json:"tgl_parsed"`
	Timestampz string `json:"timestamp"`
	USDPrice   string `json:"price_usd"`
}

type KomData struct {
	Provinsi string `json:"area_provinsi"`
	Profit   map[string]map[string]int
	Max      float64 `json:"max_profit"`
	Min      float64 `json:"min_profit"`
	Avg      float64 `json:"average_profit"`
	Median   float64 `json:"median_profit"`
}
