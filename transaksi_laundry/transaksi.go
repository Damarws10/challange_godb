package transaksi_laundry

import (
	"fmt"
	"time"

	"challenge-godb/dbconnect"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Transaksi struct{
	Id_transaksi int
	Nama_cust string
	No_hp string
	Item string
	Jumlah_satuan int
	Total_harga int
	Tgl_masuk string
	Tgl_keluar string
	Admin string
}

type AddTransaksi struct{
	Id_transaksi int
	Nama_cust string `validate:"required"`
	No_hp string `validate:"required"`
	Item int `validate:"required"`
	Jumlah_satuan int `validate:"required"`
	Total_harga int `validate:"required"`
	Tgl_masuk string `validate:"required"`
	Tgl_keluar string `validate:"required"`
	Admin int `validate:"required"`
}

type TotalTransaksi struct{
	Name_cust string
	Total_harga int
}

func GetallTransaksi() []Transaksi{
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	statementSql := "select tranx.id_transaksi, tranx.nama_cust, tranx.no_hp, pelayan.nama_pelayanan, tranx.jml_satuan, tranx.total_harga, tranx.tgl_masuk, tranx.tgl_keluar, adminlaun.name as Admin from tb_transaksi_detail as tranx join tb_pelayanan as pelayan on tranx.id_pelayanan = pelayan.id_pelayanan join tb_admin_laundry as adminlaun on tranx.id_admin = adminlaun.id_admin"

	rows, err := db.Query(statementSql)

	defer rows.Close()

	dataTransaksi := []Transaksi{}

	for rows.Next(){
		transaksi := Transaksi{}
		err := rows.Scan(&transaksi.Id_transaksi, &transaksi.Nama_cust, &transaksi.No_hp, &transaksi.Item, &transaksi.Jumlah_satuan, &transaksi.Total_harga, &transaksi.Tgl_masuk, &transaksi.Tgl_keluar, &transaksi.Admin)

		if err != nil{
			panic(err)
		}
		dataTransaksi = append(dataTransaksi, transaksi)
	}

	err = rows.Err()
	if err != nil{
		panic(err)
	}

	return dataTransaksi
}

func GetTotalHarga()[]TotalTransaksi{
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	statementSql := "select nama_cust, sum(total_harga) as Total from tb_transaksi_detail group by nama_cust"

	rows, err := db.Query(statementSql)

	defer rows.Close()

	dataTransaksi := []TotalTransaksi{}

	for rows.Next(){
		transaksi := TotalTransaksi{}
		err := rows.Scan(&transaksi.Name_cust, &transaksi.Total_harga)

		if err != nil{
			panic(err)
		}
		dataTransaksi = append(dataTransaksi, transaksi)
	}

	err = rows.Err()
	if err != nil{
		panic(err)
	}

	return dataTransaksi
}

func GetdataHargaitem(id int) int{
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	statemenSql := "SELECT harga FROM tb_pelayanan WHERE id_pelayanan=$1"

	type Harga struct{
		Get_harga int
	}

	harga := Harga{}

	err = db.QueryRow(statemenSql, id).Scan(&harga.Get_harga)

	if err != nil {
		panic(err)
	}

	return harga.Get_harga
}

func (a *AddTransaksi) AddDataTransaksi(){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())

	dataTransaksi := AddTransaksi{
		Nama_cust: a.Nama_cust,
		No_hp: a.No_hp,
		Item: a.Item,
		Jumlah_satuan: a.Jumlah_satuan,
		Total_harga: a.Total_harga,
		Tgl_masuk: a.Tgl_masuk,
		Tgl_keluar: a.Tgl_keluar,
		Admin: a.Admin, 
	}

	err = validate.Struct(dataTransaksi)

	if err != nil{
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsErr := range validationErrors{
			fmt.Println("Error", fieldsErr.Error())
		}
		fmt.Println("Gagal Insert")
		time.Sleep(time.Second * 3)
	}else{
		statementSql := "insert into tb_transaksi_detail(nama_cust, no_hp, id_pelayanan, jml_satuan, total_harga, tgl_masuk, tgl_keluar, id_admin) values($1, $2, $3, $4, $5, $6, $7, $8);"

		_, err = db.Exec(statementSql, dataTransaksi.Nama_cust, dataTransaksi.No_hp, dataTransaksi.Item, dataTransaksi.Jumlah_satuan, dataTransaksi.Total_harga, dataTransaksi.Tgl_masuk, dataTransaksi.Tgl_keluar, dataTransaksi.Admin)
		if err != nil{
			panic(err)
		}else{
			fmt.Println("Succes insert")
		}
	}
}

