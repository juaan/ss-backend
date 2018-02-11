package models

import (
	"time"
)

type (
	// Penjualan is struct for Barang kluar
	Penjualan struct {
		ID           int64     `json:"id" orm:"column(id);auto"`
		SKU          string    `json:"sku" orm:"column(sku)"`
		NamaItem     string    `json:"nama_item" orm:"column(nama_item)"`
		JumlahKeluar int64     `json:"jumlah_keluar" orm:"column(jumlah_keluar)"`
		HargaJual    int64     `json:"harga_jual" orm:"column(harga_jual)"`
		Catatan      string    `json:"catatan" orm:"column(catatan)"`
		Waktu        time.Time `json:"waktu" orm:"column(waktu);auto_now_add;type(datetime)"`
		Total        int64     `json:"total" orm:"column(total)"`
	}

	// RequestPenjualan ...
	RequestPenjualan struct {
		FromDate string
		ToDate   string
		Query    string
	}

	// RequestUpdatePenjualan ...
	RequestUpdatePenjualan struct {
		JumlahPengurangan int64  `json:"jumlah_pengurangan"`
		IDPenjualan       int64  `json:"id_penjualan"`
		Catatan           string `json:"catatan"`
	}

	// ResponseUpdatePemesanan ...
	ResponseUpdatePemesanan struct {
		SKU            string `json:"sku" orm:"column(sku)"`
		NamaItem       string `json:"nama_item" orm:"column(nama_item)"`
		NoKwitansi     string `json:"no_kwitansi" orm:"column(no_kwitansi)"`
		JumlahPesanan  int64  `json:"jumlah_pesanan" orm:"column(jumlah_pesanan)"`
		JumlahDiterima int64  `json:"jumlah_diterima" orm:"column(jumlah_diterima)"`
		Status         string `json:"status" orm:"column(status)"`
	}
)
