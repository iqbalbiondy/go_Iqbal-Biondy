CREATE TABLE pelanggan (
  id_pelanggan INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  nama_pelanggan VARCHAR(50) NOT NULL,
  alamat TEXT NOT NULL,
  tgl_lahir DATE NOT NULL,
  status_user CHAR(2) NOT NULL,
  gender CHAR(2) NOT NULL,
  created_at VARCHAR(25) NOT NULL,
  update_at VARCHAR(25) NOT NULL,
  PRIMARY KEY(id_pelanggan)
);

CREATE TABLE product (
  idproduct INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  product VARCHAR(50) NOT NULL,
  product_type VARCHAR(50) NOT NULL,
  product_desc TEXT NOT NULL,
  operator VARCHAR(25) NOT NULL,
  payment_method VARCHAR(25) NOT NULL,
  PRIMARY KEY(idproduct)
);

CREATE TABLE transaksi (
  id_transaksi INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  product_idproduct INTEGER UNSIGNED NOT NULL,
  pelanggan_id_pelanggan INTEGER UNSIGNED NOT NULL,
  tgl_transaksi DATE NOT NULL,
  total_harga FLOAT NOT NULL,
  PRIMARY KEY(id_transaksi),
  INDEX transaksi_FKIndex1(pelanggan_id_pelanggan),
  INDEX transaksi_FKIndex2(product_idproduct),
  FOREIGN KEY(pelanggan_id_pelanggan)
    REFERENCES pelanggan(id_pelanggan)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION,
  FOREIGN KEY(product_idproduct)
    REFERENCES product(idproduct)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION
);

CREATE TABLE pelanggan_membeli (
  pelanggan_id_pelanggan INTEGER UNSIGNED NOT NULL,
  product_idproduct INTEGER UNSIGNED NOT NULL,
  PRIMARY KEY(pelanggan_id_pelanggan, product_idproduct),
  INDEX pelanggan_has_product_FKIndex1(pelanggan_id_pelanggan),
  INDEX pelanggan_has_product_FKIndex2(product_idproduct),
  FOREIGN KEY(pelanggan_id_pelanggan)
    REFERENCES pelanggan(id_pelanggan)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION,
  FOREIGN KEY(product_idproduct)
    REFERENCES product(idproduct)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION
);

CREATE TABLE detail_transaksi (
  id_detail_transaksi INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  transaksi_id_transaksi INTEGER UNSIGNED NOT NULL,
  jumlah INTEGER UNSIGNED NOT NULL,
  subtotal_harga FLOAT NOT NULL,
  PRIMARY KEY(id_detail_transaksi),
  INDEX detail_transaksi_FKIndex1(transaksi_id_transaksi),
  FOREIGN KEY(transaksi_id_transaksi)
    REFERENCES transaksi(id_transaksi)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION
);

CREATE TABLE kurir (
  id_kurir INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  created_at VARCHAR(50) NOT NULL,
  updated_at VARCHAR(50) NOT NULL,
  PRIMARY KEY(id_kurir)
);

