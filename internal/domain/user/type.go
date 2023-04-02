package user

import "time"

type (
	User struct {
		IdUser   string `json:"id_user"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Pembeli struct {
		IdPembeli    string    `json:"id_pembeli"`
		Nama         string    `json:"nama"`
		Alamat       string    `json:"alamat"`
		TanggalLahir time.Time `json:"tanggal_lahir"`
		NoHp         string    `json:"no_hp"`
		JenisKelamin string    `json:"jenis_kelamin"`
	}
	Penjual struct {
		IdPenjual      string    `json:"id_penjual"`
		Nama           string    `json:"nama"`
		Alamat         string    `json:"alamat"`
		TanggalLahir   time.Time `json:"tanggal_lahir"`
		NoHp           string    `json:"no_hp"`
		JenisKelamin   string    `json:"jenis_kelamin"`
		KartuIdentitas string    `json:"kartu_identitas"`
		RekeningBank   string    `json:"rekening_bank"`
		NomorRekening  int       `json:"nomor_rekening"`
	}
)
