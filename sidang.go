package sidang

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type StudentsDataTables struct {
	Con *sqlx.DB
	Ctx context.Context
}

func StudentsDataTablesMgr(db *sqlx.DB, ctx context.Context) *StudentsDataTables {
	return &StudentsDataTables{Con: db, Ctx: ctx}
}

func (std *StudentsDataTables) GetNPMbyPhone(phone int) (dest Mhsdetail, err error) {
	err = std.Con.GetContext(
		std.Ctx,
		&dest,
		"SELECT MhswId, Nama, Handphone FROM simak_mst_mahasiswa WHERE handphone = ?",
		phone,
	)
	return
}

func (std *StudentsDataTables) GetJumlahPertemuan(Mhswid int, tipe string) (dest Bimbingandata, err error) {
	err = std.Con.GetContext(
		std.Ctx,
		&dest,
		"SELECT npm, tahun_id, tipe_bimbingan, total_bimbingan, last_bimbingan FROM bimbingan_data WHERE npm = ? and tipe_bimbingan = ?",
		Mhswid,
		tipe,
	)
	return
}
