package sidang

import "time"

type Skt struct {
	Nomor string   `json:"nomor,omitempty" bson:"nomor,omitempty"`
	Judul []string `json:"judul,omitempty" bson:"judul,omitempty"`
}

type Mhsdetail struct {
	MhswId    int    `json:"mhsw_id" bson:"mhsw_id"`
	Nama      string `json:"nama" bson:"nama"`
	Handphone int    `json:"handphone" bson:"handphone"`
}

type Bimbingandata struct {
	NPM            string    `json:"npm" bson:"npm"`
	TahunId        string    `json:"tahun_id" bson:"tahun_id"`
	TipeBimbingan  string    `json:"tipe_bimbingan" bson:"tipe_bimbingan"`
	TotalBimbingan int       `json:"total_bimbingan" bson:"total_bimbingan"`
	LastBimbingan  time.Time `json:"last_bimbingan" bson:"last_bimbingan"`
}
