package item_laundry

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"challenge-godb/dbconnect"
)

type Item struct{
	Id_pelayanan int
	Nama_pelayanan string `validate:"required"`
	Harga int `validate:"required"`
	Satuan string `validate:"oneof=kg buah prefer_not_to"`
}

var validate *validator.Validate

func GetalldataItem() []Item {
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	statementSql := "SELECT * FROM tb_pelayanan order by id_pelayanan asc"

	rows, err := db.Query(statementSql)

	defer rows.Close()

	dataItem := []Item{}

	for rows.Next(){
		item := Item{}
		err := rows.Scan(&item.Id_pelayanan, &item.Nama_pelayanan, &item.Harga, &item.Satuan)

		if err != nil{
			panic(err)
		}
		dataItem = append(dataItem, item)
	}

	err = rows.Err()
	if err != nil{
		panic(err)
	}

	return dataItem
}

func (i *Item) AddItem(){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())

	dataItem := Item{
		Nama_pelayanan: i.Nama_pelayanan,
		Harga: i.Harga,
		Satuan: i.Satuan,
	}

	err = validate.Struct(dataItem)

	if err != nil{
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsErr := range validationErrors{
			fmt.Println("Error", fieldsErr.Error())
		}
		fmt.Println("Gagal Insert")
		time.Sleep(time.Second * 3)
	}else{
		statementSql := "INSERT INTO tb_pelayanan (Nama_pelayanan, harga, satuan) values($1, $2, $3);"

		_, err = db.Exec(statementSql, dataItem.Nama_pelayanan, dataItem.Harga, dataItem.Satuan)
		if err != nil{
			panic(err)
		}else{
			fmt.Println("Succes insert")
		}
	}
} 

func GetdataItem(id int)(*Item, bool){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())

	item := &Item{}

	type ValidatorItem struct{
		Id_item int `validate:"required"`
	}

	itemCek := &ValidatorItem{
		Id_item: id,
	}

	err = validate.Struct(itemCek)

	if err != nil{
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsErr := range validationErrors{
			fmt.Println("Error", fieldsErr.Error())
		}
		fmt.Println("Data Miss")
		time.Sleep(time.Second * 3)
	}else{
	statemenSql := "SELECT * FROM tb_pelayanan WHERE id=$1"

	err = db.QueryRow(statemenSql, itemCek.Id_item).Scan(&item.Id_pelayanan, &item.Nama_pelayanan, &item.Harga, &item.Satuan)	
	}

	stats := isZero(*item)

	return item, stats
}

func (i *Item) UpdatedataItem(){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())

	item := Item{
		Nama_pelayanan: i.Nama_pelayanan,
		Harga: i.Harga,
		Satuan: i.Satuan,
	}

	err = validate.Struct(item)

	if err != nil{
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsErr := range validationErrors{
			fmt.Println("Error", fieldsErr.Error())
		}
		fmt.Println("Gagal Update")
		time.Sleep(time.Second * 3)
	}else{
		statementSql := "UPDATE tb_pelayanan set nama_pelayanan=$2, harga=$3, satuan=$4 WHERE id=$1;"

		_, err = db.Exec(statementSql, i.Id_pelayanan, item.Nama_pelayanan, item.Harga, item.Satuan)
		if err != nil{
			panic(err)
		}else{
			fmt.Println("Succes Update")
		}
	}
}

func DeletedataItem(id int){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	statementSql := "DELETE FROM tb_pelayanan WHERE id=$1"

	_ , err = db.Exec(statementSql, id)
	if err != nil {
		panic(err)
	}else{
		fmt.Println("Succes Delete")
	}
}

func isZero(i Item) bool{
	return i == Item{}
}