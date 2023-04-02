package transcation

import "time"

type (
	Produk struct {
		IdProduk   string  `json:"id_produk"`
		IdPenjual  string  `json:"id_penjual"`
		NamaProduk string  `json:"nama_produk"`
		Kategori   string  `json:"kategori"`
		Deskripsi  string  `json:"deskripsi"`
		Harga      float64 `json:"harga"`
		Stok       int     `json:"stok"`
		Berat      float64 `json:"berat"`
	}
	Pesanan struct {
		IdPesanan      string    `json:"id_pesanan"`
		IdPembeli      string    `json:"id_pembeli"`
		IdProduk       string    `json:"id_produk"`
		WaktuTransaksi time.Time `json:"waktu_transaksi"`
		Status         string    `json:"status"`
	}
	PesananProduk struct {
		IdPesananProduk string `json:"id_pesanan_produk"`
		IdPesanan       string `json:"id_pesanan"`
		IdProduk        string `json:"id_produk"`
		Jumlah          int    `json:"jumlah"`
	}
	Pembayaran struct {
		IdPembayaran     string    `json:"id_pembayaran"`
		IdPesanan        string    `json:"id_pesanan"`
		TanggalBayar     time.Time `json:"waktu_bayar"`
		TotalBayar       float64   `json:"total_bayar"`
		MetodePembayaran string    `json:"metode_pembayaran"`
	}
)
