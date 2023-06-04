## Resume - Compute Service

---

Berikut adalah beberapa ringkasan materi yang saya buat mengenai Compute Service:

1. Compute Service adalah layanan cloud yang menyediakan kapasitas komputasi untuk menjalankan aplikasi dan beban kerja pada infrastruktur yang dikelola oleh penyedia layanan cloud.

2. Compute Service menyediakan berbagai jenis server virtual yang dapat dikonfigurasi dan dioperasikan secara fleksibel, termasuk virtual private server (VPS), instance EC2 di AWS, atau Compute Engine di GCP.

3. Compute Service memungkinkan pengguna untuk menyesuaikan spesifikasi komputasi seperti jumlah CPU, memori, penyimpanan, jaringan, dan sistem operasi yang digunakan.

4. Compute Service menyediakan mekanisme skalabilitas otomatis, sehingga aplikasi dan beban kerja dapat dijalankan pada infrastruktur yang dapat menyesuaikan diri dengan permintaan.

5. Compute Service menyediakan fitur pengelolaan beban kerja dan otomatisasi yang dapat membantu pengguna dalam menjalankan dan mengelola aplikasi, termasuk pengaturan jadwal, penyeimbangan beban, dan manajemen pemantauan.

6. Compute Service juga menyediakan layanan keamanan seperti firewall, akses jaringan, dan enkripsi data untuk menjaga keamanan beban kerja dan data pengguna.

7. Compute Service memungkinkan pengguna untuk membayar hanya untuk kapasitas yang mereka gunakan, dengan model pembayaran berbasis penggunaan yang fleksibel dan transparan.

8. Compute Service dapat digunakan dalam berbagai konteks, termasuk pengembangan, pengujian, dan produksi, dan dapat diintegrasikan dengan layanan cloud lainnya seperti penyimpanan, database, dan jaringan.

9. Compute Service telah menjadi salah satu layanan paling penting di lingkungan cloud, karena kemampuannya untuk menyediakan kapasitas komputasi yang mudah dikelola dan sangat fleksibel dengan biaya yang terjangkau.


## Praktikum - Docker & AWS Compute Service

---

1.  Melakukan konfigurasi EC2

    ![image](./Screenshoot/1_ConfigureEC2.png)

2.  Membuat Create Key Pair

    ![image](./Screenshoot/2_CreateKeyPair.png)

3.  Menyetting Network Setting

    ![image](./Screenshoot/3_NetworkSetting.png)

4.  EC2 Instance berhasil dibuat

    ![image](./Screenshoot/4_EC2Sucessfully.png)

5.  mengkoneksi _instance EC2_ dari lokal komputer

    ![image](./Screenshoot/5_ConnectLocalToEC2.png)

6.  melakukan instalasi docker

    ![image](./Screenshoot/6_InstallDocker.png)

7.  melakukan instalasi & konfigurasi docker compose

    ![image](./Screenshoot/7_DockerCompose.png)

8.  melakukan clonning github ke instance EC2

    ![image](./Screenshoot/8_CloneGitHub.png)

9.  Melakukan Build Docker Image

    ![image](./Screenshoot/9_BuildDockerImage.png)

10. melakukan deploymen AWS RDS dengan memilih MySQL sebagai database

    ![image](./Screenshoot/10_choosingRDSMySQL.png)

11. melakukan konfigurasi MySQL di AWS RDS

    ![image](./Screenshoot/11_ConfigureMySQL.png)

12. AWS RDS MySQL Berhasil dibuat

    ![image](./Screenshoot/12_RDSMySQLSuccessfully.png)

13. melakukan koneksi dari AWS RDS dengan database MySQL ke ubuntu intance EC2

    ![image](./Screenshoot/13_ConnectMySQLEC2.png)

14. Buid, Push _Image Docker_ ke docker hub dan Pull _Image Docker_ dari instance EC2

    ![image](./Screenshoot/14_Build&Pushdocker.png)

15. melakukan run Docker image

    ![image](./Screenshoot/15_DockerRunSuccesfully.png)

16. mengakses project GO ke postman yang telah berjalan di docker image 

    ![image](./Screenshoot/16_AccessedGoProjectInPostman.png)

17. menguji inputan data yang telah dimasukkan dari postman dengan melihat langsung isi dari databasenya


    ![image](./Screenshoot/17_TestingInputDataInDB.png)
        ![image](./Screenshoot/18_TestingInputDataInDB_`2.png)
