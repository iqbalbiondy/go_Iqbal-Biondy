#     1. Insert 5 operators pada table operators.
INSERT INTO operators (name, created_at, updated_at)
VALUES
("Operator 1", NOW(), NOW()),
("Operator 2", NOW(), NOW()),
("Operator 3", NOW(), NOW()),
("Operator 4", NOW(), NOW()),
("Operator 5", NOW(), NOW());

#     2. Insert 3 product type.
INSERT INTO product_types (name, created_at, updated_at)
VALUES
("Product Type 1", NOW(), NOW()),
("Product Type 2", NOW(), NOW()),
("Product Type 3", NOW(), NOW());
#     3. Insert 2 product dengan product type id = 1, dan operators id = 3.
INSERT INTO products (product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
(1, 3, "PROD001", "Product 1", 1, NOW(), NOW()),
(1, 3, "PROD002", "Product 2", 1, NOW(), NOW());

#     4. Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO products (product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
(2, 1, "PROD003", "Product 3", 1, NOW(), NOW()),
(2, 1, "PROD004", "Product 4", 1, NOW(), NOW()),
(2, 1, "PROD005", "Product 5", 1, NOW(), NOW());

#     5. Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO products (product_type_id, operator_id, code, name, status, created_at, updated_at)
VALUES
(3, 4, "PROD006", "Product 6", 1, NOW(), NOW()),
(3, 4, "PROD007", "Product 7", 1, NOW(), NOW()),
(3, 4, "PROD008", "Product 8", 1, NOW(), NOW());

#     6. Insert product description pada setiap product.
INSERT INTO product_descriptions (id, description, created_at, updated_at)
VALUES
(1, "Description of Product 1", NOW(), NOW()),
(2, "Description of Product 2", NOW(), NOW()),
(3, "Description of Product 3", NOW(), NOW()),
(4, "Description of Product 4", NOW(), NOW()),
(5, "Description of Product 5", NOW(), NOW()),
(6, "Description of Product 6", NOW(), NOW()),
(7, "Description of Product 7", NOW(), NOW()),
(8, "Description of Product 8", NOW(), NOW());

#     7. Insert 3 payment methods.
INSERT INTO payment_methods (name, status, created_at, updated_at)
VALUES
("Payment Method 1", 1, NOW(), NOW()),
("Payment Method 2", 1, NOW(), NOW()),
("Payment Method 3", 1, NOW(), NOW());

#     8. Insert 5 user pada tabel user.
INSERT INTO users (status, dob, gender, created_at, updated_at)
VALUES
(1, '1990-01-01', 'M', NOW(), NOW()),
(1, '1992-05-10', 'F', NOW(), NOW()),
(1, '1985-11-15', 'M', NOW(), NOW()),
(1, '1998-07-20', 'F', NOW(), NOW()),
(1, '1978-03-05', 'M', NOW(), NOW());

#     9. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
SELECT u.id, pm.id, 'completed', 3, 5000000
FROM users u
CROSS JOIN payment_methods pm
WHERE u.id BETWEEN 1 AND 5;

#     10. Insert 3 product di masing-masing transaksi.
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
SELECT t.id, p.id, 'completed', 1, 1000000
FROM transactions t
CROSS JOIN products p
WHERE t.user_id BETWEEN 1 AND 5
LIMIT 3;

UPDATE users SET nama = 'Admaja Dwi Herlambang' WHERE id = 1;
UPDATE users SET nama = 'Faizatul Amalia' WHERE id = 2;
UPDATE users SET nama = 'Fitra A Bachtiar' WHERE id = 3;
UPDATE users SET nama = 'Asa Fitri Sabila' WHERE id = 4;
UPDATE users SET nama = 'Satrio Agung Wicaksono' WHERE id = 5;
