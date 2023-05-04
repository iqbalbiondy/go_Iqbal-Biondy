# -- 1. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT nama FROM users WHERE gender IN ('Laki-laki', 'M');
# -- 2. Tampilkan product dengan id = 3.
SELECT * FROM products WHERE id = 3;
# -- 3. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM users
WHERE nama LIKE '%a%'
AND created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY);
# -- 4. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) AS Total_Perempuan FROM users WHERE gender = 'F';
# -- 5. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users ORDER BY nama ASC;
# -- 6. Tampilkan 5 data pada data product
SELECT * FROM products LIMIT 5;