package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strconv"

	"challenge-godb/admin_laundry"
	"challenge-godb/item_laundry"
	"challenge-godb/transaksi_laundry"
)

var scanner = bufio.NewScanner(os.Stdin)

func main(){
	var opsi int

	forloop:for{
		fmt.Println("Silahkan Pilih Opsi")
		fmt.Println("1. Load Data Item")
		fmt.Println("2. Tambah Data Item")
		fmt.Println("3. Update Data Item")
		fmt.Println("4. Delete Data Item")
		fmt.Println("=============================")
		fmt.Println("5. Load Data Admin")
		fmt.Println("6. Tambah Data Admin")
		fmt.Println("7. Update Data Admin")
		fmt.Println("8. Delete Data Admin")
		fmt.Println("=============================")
		fmt.Println("9. Semua Transaksi")
		fmt.Println("10. Add Transaksi")
		fmt.Println("11. Exit Proggram")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Masukkan Pilihan : ")
		scanner.Scan()
		opsi, _ = strconv.Atoi(scanner.Text())

		switch opsi{
			case 1:
				fmt.Println("Get All Data Item")
				loadDataItem()
				time.Sleep(time.Second * 5)
				continue
			case 2:
				addDataItem()
				continue
			case 3:
				fmt.Println("=============Update Item=============")
				updateItem()
				time.Sleep(time.Second * 3)
			case 4:
				fmt.Println("================Delete Item============")
				deleteDataItem()
				time.Sleep(time.Second * 2)
			case 5:
				fmt.Println("Get All Data Admin")
				loadDataAdmin()
				time.Sleep(time.Second * 5)
				continue
			case 6:
				addData()
				time.Sleep(time.Second * 3)
				continue
			case 7:
				fmt.Println("=============Update Admin=============")
				updatedataByidAdmin()
				time.Sleep(time.Second * 3)
			case 8:
				fmt.Println("================Delete Admin============")
				deleteDataAdmin()
				time.Sleep(time.Second * 2)
			case 9:
				viewGetallTransaksi()
				time.Sleep(time.Second * 5)
			case 10:
				addTransaksi()
			case 11:
				fmt.Println("You are close the Proggram, please contact me if any bug exits <3")
				time.Sleep(time.Second * 5)
				break forloop
			default:
				fmt.Println("Opsi Tidak ada")
		}
	}
}


//Semua Fungsi Admin
func loadDataAdmin(){
	layout := "2006-01-02T15:04:05Z07:00"
	// Get all data Admin
	admin := admin_laundry.GetallData()

	for _, dataAdmin := range admin{
    	dateString := dataAdmin.Tgl_masuk
    	t, err  := time.Parse(layout, dateString)

    	if err != nil{
    		panic(err)
    	}

		fmt.Printf("| Id : %d | Nama : %s | Tanggal Masuk : %s | Status : %s | \n", dataAdmin.Id, dataAdmin.Name, t.Format("01-02-2006"), dataAdmin.Status)
	}
}

func addData(){
	fmt.Println("===========Masukkan data admin=============")
	fmt.Print("Masukkan Nama : ")
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Masukkan Tanggal Masuk : ")
	scanner.Scan()
	tgl_masuk := scanner.Text()
	fmt.Print("Masukkan Status (active/inactive): ")
	scanner.Scan()
	status := scanner.Text()

	admin := admin_laundry.Admin{Name : nama, Tgl_masuk: tgl_masuk, Status: status}
	admin.Addata()
}

func updatedataByidAdmin(){
	fmt.Println("============Check Data Sebelum Update=========")
	fmt.Print("Masukan Id Admin : ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())
	showData, stat :=admin_laundry.Getdatabyid(id)
	if stat == true {
		fmt.Println("No Data Available")
	}else{
		fmt.Println("Data Available")
		fmt.Println(*showData)

		fmt.Print("Apakah Mau Update Data ? (y/t) :")
		scanner.Scan()
		pilihan := scanner.Text()

		if pilihan == "y"{
			fmt.Println("===========Masukkan data admin=============")
			fmt.Print("Masukkan Nama Update: ")
			scanner.Scan()
			nama := scanner.Text()
			fmt.Print("Masukkan Tanggal Masuk Update: ")
			scanner.Scan()
			tgl_masuk := scanner.Text()
			fmt.Print("Masukkan Status Update (active/inactive): ")
			scanner.Scan()
			status := scanner.Text()

			adminUpdate := admin_laundry.Admin{Id: id, Name : nama, Tgl_masuk: tgl_masuk, Status: status}

			adminUpdate.UpdatedAdmin()

		}else if pilihan == "t"{
			fmt.Println("Tidak Update")
		}else{
			fmt.Println("Tidak Ada Opsi")
		}
	}
}

func deleteDataAdmin(){
	fmt.Println("============Check Data Sebelum Delete=========")
	fmt.Print("Masukan Id Admin : ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())
	showData, stat := admin_laundry.Getdatabyid(id)
	if stat == true {
		fmt.Println("No Data Available")
	}else{
		fmt.Println("Data Available")
		fmt.Println(*showData)


		fmt.Print("Apakah Mau Hapus Data ? (y/t) :")
		scanner.Scan()
		pilihan := scanner.Text()

		if pilihan == "y"{
			admin_laundry.DeletedataAdmin(id)
		}else if pilihan == "t"{
			fmt.Println("Tidak Dihapus")
		}else{
			fmt.Println("Tidak Ada Opsi")
		}

	}
}


//Semua Fungsi Item
func loadDataItem(){
	dataItem := item_laundry.GetalldataItem()

	for _, data := range dataItem{
		fmt.Printf("| Id : %d | Nama Item : %s | Harga : %d | Satuan : %s | \n", data.Id_pelayanan, data.Nama_pelayanan, data.Harga, data.Satuan)
	}
}

func addDataItem(){
	fmt.Println("===========Masukkan data Item=============")
	fmt.Print("Masukkan Nama Item: ")
	scanner.Scan()
	nama_item := scanner.Text()
	fmt.Print("Masukkan Harga : ")
	scanner.Scan()
	harga, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Masukkan Satuan (kg/buah): ")
	scanner.Scan()
	satuan := scanner.Text()

	item := item_laundry.Item{Nama_pelayanan: nama_item, Harga: harga, Satuan: satuan}
	item.AddItem()
}

func updateItem(){
	fmt.Println("============Check Data Sebelum Update=========")
	fmt.Print("Masukan Id Item : ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())
	showData, stat := item_laundry.GetdataItem(id)
	if stat == true {
		fmt.Println("No Data Available")
	}else{
		fmt.Println(*showData)

		fmt.Print("Apakah Mau Update Data ? (y/t) :")
		scanner.Scan()
		pilihan := scanner.Text()

		if pilihan == "y"{
			fmt.Println("===========Masukkan data Item Update=============")
			fmt.Print("Masukkan Nama Item : ")
			scanner.Scan()
			nama := scanner.Text()
			fmt.Print("Masukkan Harga : ")
			scanner.Scan()
			harga, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Masukkan Satuan kg/buah: ")
			scanner.Scan()
			satuan := scanner.Text()

			item := item_laundry.Item{
				Id_pelayanan: id,
				Nama_pelayanan: nama,
				Harga: harga,
				Satuan: satuan,
			}

			item.UpdatedataItem()

		}else if pilihan == "t"{
			fmt.Println("Tidak Update")
		}else{
			fmt.Println("Tidak Ada Opsi")
		}		
	}
}

func deleteDataItem(){
	fmt.Println("============Check Data Sebelum Delete=========")
	fmt.Print("Masukan Id Item : ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())
	showData, stat := item_laundry.GetdataItem(id)
	if stat == true {
		fmt.Println("No Data Available")
	}else{
		fmt.Println("Data Available")
		fmt.Println(*showData)

		fmt.Print("Apakah Mau Hapus Data ? (y/t) :")
		scanner.Scan()
		pilihan := scanner.Text()

		if pilihan == "y"{
			item_laundry.DeletedataItem(id)
		}else if pilihan == "t"{
			fmt.Println("Tidak Dihapus")
		}else{
			fmt.Println("Tidak Ada Opsi")
		}
	}
}

func viewGetallTransaksi(){
	layout := "2006-01-02T15:04:05Z07:00"
	// Get all data Transaksi
	transaksi := transaksi_laundry.GetallTransaksi()

	for _, t := range transaksi{
		tgl_masuk := t.Tgl_masuk
		tgl_keluar := t.Tgl_keluar
    	t_masuk, err  := time.Parse(layout, tgl_masuk)
    	t_keluar, err  := time.Parse(layout, tgl_keluar)

    	if err != nil{
    		panic(err)
    	}

    	fmt.Printf("|Id Transaksi: %d, Nama: %s, No_hp: %s, Item: %s, Jumlah: %d, Total: %d, Tanggal Masuk: %s, Tanggal Keluar: %s, Admin: %s| \n",
    	 			t.Id_transaksi, t.Nama_cust, t.No_hp, t.Item, t.Jumlah_satuan, t.Total_harga, t_masuk.Format("01-02-2006"), t_keluar.Format("01-02-2006"), t.Admin)
	}

	allTotalharga := transaksi_laundry.GetTotalHarga()

	for _, total_harga := range allTotalharga{
		fmt.Printf("|Nama Customer: %s, Total Harga: %d |\n", total_harga.Name_cust, total_harga.Total_harga)
	}
}

func addTransaksi(){
	fmt.Println("=================Masukan Data Transaksi=================")
	fmt.Print("Masukkan Nama Customer: ")
	scanner.Scan()
	nama_cust := scanner.Text()

	fmt.Print("Masukkan No Hp: ")
	scanner.Scan()
	no_hp := scanner.Text()

	loadDataItem()

	fmt.Print("Masukkan Id Item Pelayanan: ")
	scanner.Scan()
	itemPelayanan,_ := strconv.Atoi(scanner.Text())

	fmt.Print("Masukkan Jumlah Satuan: ")
	scanner.Scan()
	jmlSatuan,_ := strconv.Atoi(scanner.Text())

	fmt.Print("Masukkan Tanggal Masuk: ")
	scanner.Scan()
	tglMasuk := scanner.Text()

	fmt.Print("Masukkan Tanggal Keluar: ")
	scanner.Scan()
	tglKeluar := scanner.Text()

	fmt.Print("Masukkan Id Admin: ")
	scanner.Scan()
	id_admin,_ := strconv.Atoi(scanner.Text())

	total_harga := jmlSatuan * transaksi_laundry.GetdataHargaitem(1)

	dataTransaksi := transaksi_laundry.AddTransaksi{
		Nama_cust: nama_cust,
		No_hp: no_hp,
		Item: itemPelayanan,
		Jumlah_satuan: jmlSatuan,
		Total_harga: total_harga,
		Tgl_masuk: tglMasuk,
		Tgl_keluar: tglKeluar,
		Admin: id_admin, 
	}

	dataTransaksi.AddDataTransaksi()
}