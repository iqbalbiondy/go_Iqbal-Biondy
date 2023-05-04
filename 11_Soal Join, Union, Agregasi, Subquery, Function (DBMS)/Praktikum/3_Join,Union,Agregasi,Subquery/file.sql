# 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM transactions WHERE user_id IN (1, 2);

# 2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) AS total_harga FROM transactions WHERE user_id = 1;

# 3. Tampilkan total transaksi dengan product type 2.
SELECT COUNT(*) AS total_transaksi FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
WHERE p.product_type_id = 2;


# 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT p.*, pt.name AS product_type_name FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

# 5. Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT t.*, p.name AS product_name, u.nama AS user_name FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id
JOIN products p ON td.product_id = p.id
JOIN users u ON t.user_id = u.id;


# 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER $$

CREATE TRIGGER after_transactions_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
  DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END $$

DELIMITER ;
# Penerapan funtion
SELECT update_total_qty_transaction_detail_deleted(3);



# 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER $$

CREATE FUNCTION update_total_qty_transaction_detail_deleted(transaction_id INT)
RETURNS BOOLEAN
BEGIN
    DECLARE qty_deleted INT;
    SELECT SUM(qty) INTO qty_deleted FROM transaction_details WHERE transaction_id = transaction_id;
    UPDATE transactions SET total_qty = total_qty - qty_deleted WHERE id = transaction_id;
    RETURN TRUE;
END$$

DELIMITER ;

# Penerapan Function
DELETE FROM transactions WHERE id = 3;
SELECT update_total_qty_transaction_detail_deleted(3);



# 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT *
FROM products
WHERE id NOT IN (
  SELECT DISTINCT product_id
  FROM transaction_details
);

