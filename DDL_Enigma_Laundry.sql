CREATE TABLE "public.tb_pelayanan" (
	"id_pelayanan" serial NOT NULL,
	"nama_pelayanan" varchar(20),
	"harga" integer NOT NULL,
	"satuan" varchar(10),
	CONSTRAINT "tb_pelayanan_pk" PRIMARY KEY ("id_pelayanan")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.tb_transaksi_detail" (
	"id_transaksi" serial NOT NULL,
	"nama_cust" varchar NOT NULL,
	"no_hp" varchar NOT NULL,
	"id_pelayanan" integer NOT NULL,
	"jml_satuan" integer NOT NULL,
	"total_harga" integer NOT NULL,
	"tgl_masuk" DATE NOT NULL,
	"tgl_keluar" DATE NOT NULL,
	"id_admin" integer(10) NOT NULL,
	CONSTRAINT "tb_transaksi_detail_pk" PRIMARY KEY ("id_transaksi")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.tb_admin_laundry" (
	"id_admin" serial NOT NULL,
	"name" varchar(20) NOT NULL,
	"tgl_masuk" DATE NOT NULL,
	"status" varchar(10) NOT NULL,
	CONSTRAINT "tb_admin_laundry_pk" PRIMARY KEY ("id_admin")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "tb_transaksi_detail" ADD CONSTRAINT "tb_transaksi_detail_fk0" FOREIGN KEY ("id_pelayanan") REFERENCES "tb_pelayanan"("id_pelayanan");
ALTER TABLE "tb_transaksi_detail" ADD CONSTRAINT "tb_transaksi_detail_fk1" FOREIGN KEY ("id_admin") REFERENCES "tb_admin_laundry"("id_admin");





