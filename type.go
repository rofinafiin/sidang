package sidang

type Skt struct {
	Nomor string   `json:"nomor,omitempty" bson:"nomor,omitempty"`
	Judul []string `json:"judul,omitempty" bson:"judul,omitempty"`
}
