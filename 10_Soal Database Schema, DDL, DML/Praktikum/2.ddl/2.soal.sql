-- Membuat database alta_online_shop
CREATE DATABASE alta_online_shop;

-- Tambahkan kolom ongkos_dasar di tabel shipping
ALTER TABLE kurir
ADD COLUMN ongkos_dasar FLOAT NOT NULL AFTER updated_at;

-- Mengganti nama tabel kurir menjadi shipping
ALTER TABLE kurir
RENAME TO shipping;

-- Menghapus tabel shipping
DROP TABLE IF EXISTS shipping;

-- Menambahkan tabel payment_method_description dengan relasi 1-to-1 ke tabel product
CREATE TABLE payment_method_description (
  idpayment_method_description INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  product_idproduct INTEGER UNSIGNED NOT NULL,
  description TEXT NOT NULL,
  PRIMARY KEY(idpayment_method_description),
  INDEX payment_method_description_FKIndex1(product_idproduct),
  FOREIGN KEY(product_idproduct)
    REFERENCES product(idproduct)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION
);

-- Menambahkan kolom alamat_id di tabel pelanggan sebagai foreign key ke tabel alamat
ALTER TABLE pelanggan
ADD COLUMN alamat_id INTEGER UNSIGNED NOT NULL AFTER alamat,
ADD INDEX pelanggan_FKIndex1(alamat_id),
FOREIGN KEY(alamat_id)
  REFERENCES alamat(id_alamat)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION;

-- Membuat tabel alamat
CREATE TABLE alamat (
  id_alamat INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
  jalan VARCHAR(50) NOT NULL,
  kota VARCHAR(50) NOT NULL,
  PRIMARY KEY(id_alamat)
);

-- Menambahkan tabel user_payment_method_detail dengan relasi many-to-many antara pelanggan dan payment_method_description
CREATE TABLE user_payment_method_detail (
  pelanggan_id_pelanggan INTEGER UNSIGNED NOT NULL,
  payment_method_description_idpayment_method_description INTEGER UNSIGNED NOT NULL,
  PRIMARY KEY(pelanggan_id_pelanggan, payment_method_description_idpayment_method_description),
  INDEX user_payment_method_detail_FKIndex1(pelanggan_id_pelanggan),
  INDEX user_payment_method_detail_FKIndex2(payment_method_description_idpayment_method_description),
  FOREIGN KEY(pelanggan_id_pelanggan)
    REFERENCES pelanggan(id_pelanggan)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION,
  FOREIGN KEY(payment_method_description_idpayment_method_description)
    REFERENCES payment_method_description(idpayment_method_description)
      ON DELETE NO ACTION
      ON UPDATE NO ACTION
)