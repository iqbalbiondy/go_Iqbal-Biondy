Berikut adalah ringkasan materi yang telah saya pelajari pada course Cuncurrent Programming, diantara lain:


1. Concurrent programming adalah teknik yang digunakan untuk mengeksekusi tugas secara bersamaan, dan terdiri dari beberapa subtopik.

2. Sequential programming adalah cara tradisional untuk mengeksekusi tugas secara berurutan satu per satu, sedangkan parallel programming memungkinkan beberapa tugas dijalankan secara bersamaan.

3. Goroutines adalah unit eksekusi ringan yang digunakan untuk mengimplementasikan concurrent programming di Go. Goroutine memiliki overhead yang rendah dan sangat efisien.

4. Channels dan select adalah cara untuk mentransfer data antar goroutine. Channel adalah struktur data yang dapat digunakan untuk mengirim dan menerima data antara goroutine. Select memungkinkan kita untuk memilih dari beberapa channel yang siap untuk digunakan.

5. Race condition terjadi ketika dua atau lebih goroutine mencoba mengakses dan memodifikasi variabel secara bersamaan tanpa sinkronisasi yang memadai, dan dapat menyebabkan data race.

Untuk mengatasi masalah race condition, kita dapat menggunakan mutex. Mutex memungkinkan kita untuk mengunci akses ke suatu sumber daya sehingga hanya satu goroutine yang dapat mengaksesnya pada satu waktu.

6. Channels dapat memblokir goroutine jika tidak ada data yang tersedia untuk dibaca atau jika channel penuh dan tidak dapat menerima data.

7. WaitGroup adalah struktur data yang digunakan untuk menunggu sekelompok goroutine selesai sebelum melanjutkan eksekusi program.

Kita juga belajar cara mengimplementasikan Big Search Website dengan menggunakan teknik concurrent programming di Go. Teknik tersebut mencakup penggunaan goroutine, channel, dan waitgroup. Hal ini memungkinkan Big Search untuk mencari data secara efisien dan menghasilkan hasil yang akurat dan cepat.