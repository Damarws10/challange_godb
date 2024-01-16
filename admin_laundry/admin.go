package admin_laundry

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"challenge-godb/dbconnect"
)

type Admin struct{
	Id int
	Name string `validate:"required"`
	Tgl_masuk string `validate:"required"`
	Status string `validate:"oneof=active inactive prefer_not_to"`
}

var validate *validator.Validate

func (a *Admin) Addata(){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())

	dataAdmin := Admin{
		Name: a.Name,
		Tgl_masuk: a.Tgl_masuk,
		Status: a.Status,
	}

	err = validate.Struct(dataAdmin)

	if err != nil{
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsErr := range validationErrors{
			fmt.Println("Error", fieldsErr.Error())
		}
		fmt.Println("Gagal Insert")
		time.Sleep(time.Second * 3)
	}else{
		statementSql := "INSERT INTO tb_admin_laundry (name, tgl_masuk, status) values($1, $2, $3);"

		_, err = db.Exec(statementSql, dataAdmin.Name, dataAdmin.Tgl_masuk, dataAdmin.Status)
		if err != nil{
			panic(err)
		}else{
			fmt.Println("Succes insert")
		}
	}
}

func GetallData()[]Admin{
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	statementSql := "SELECT * FROM tb_admin_laundry order by id_admin asc"

	rows, err := db.Query(statementSql)

	defer rows.Close()

	dataAdmin := []Admin{}

	for rows.Next(){
		admin := Admin{}
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Tgl_masuk, &admin.Status)

		if err != nil{
			panic(err)
		}
		dataAdmin = append(dataAdmin, admin)
	}

	err = rows.Err()
	if err != nil{
		panic(err)
	}

	return dataAdmin
}

func Getdatabyid(id int)(*Admin, bool){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())

	dataIdAdmin := &Admin{}

	type DataId struct{
		IdData int `validate:"required"`
	}

	dataId := &DataId{
		IdData: id,
	}

	err = validate.Struct(dataId)

	if err != nil{
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsErr := range validationErrors{
			fmt.Println("Error", fieldsErr.Error())
		}
		fmt.Println("Data Miss")
		time.Sleep(time.Second * 3)
	}else{
	statemenSql := "SELECT * FROM tb_admin_laundry WHERE id_admin=$1"

	err = db.QueryRow(statemenSql, dataId.IdData).Scan(&dataIdAdmin.Id, &dataIdAdmin.Name, &dataIdAdmin.Tgl_masuk, &dataIdAdmin.Status)	
	}

	stat := isZero(*dataIdAdmin)

	return dataIdAdmin, stat
}

func (a *Admin) UpdatedAdmin(){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	validate = validator.New(validator.WithRequiredStructEnabled())

	dataAdminUpdate := Admin{
		Name: a.Name,
		Tgl_masuk: a.Tgl_masuk,
		Status: a.Status,
	}

	err = validate.Struct(dataAdminUpdate)

	if err != nil{
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldsErr := range validationErrors{
			fmt.Println("Error", fieldsErr.Error())
		}
		fmt.Println("Gagal Update")
		time.Sleep(time.Second * 3)
	}else{
		statementSql := "UPDATE tb_admin_laundry set name=$2, tgl_masuk=$3, status=$4 WHERE id_admin=$1;"

		_, err = db.Exec(statementSql, a.Id, dataAdminUpdate.Name, dataAdminUpdate.Tgl_masuk, dataAdminUpdate.Status)
		if err != nil{
			panic(err)
		}else{
			fmt.Println("Succes Update")
		}
	}
}

func DeletedataAdmin(id int){
	var db = dbconnect.ConnectDB()
	defer db.Close()
	var err error

	statementSql := "DELETE FROM tb_admin_laundry WHERE id_admin=$1"

	_ , err = db.Exec(statementSql, id)
	if err != nil {
		panic(err)
	}else{
		fmt.Println("Succes Delete")
	}
}

func isZero(a Admin) bool{
	return a == Admin{}
}