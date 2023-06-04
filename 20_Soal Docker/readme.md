Berikut adalah beberapa ringkasan materi yang saya buat mengenai Docker:

1. Docker adalah platform open-source untuk mengembangkan, mengirim, dan menjalankan aplikasi dalam wadah terisolasi, yang disebut kontainer.

2. Kontainer Docker memungkinkan pengembang untuk memisahkan aplikasi dari lingkungan hosting dan infrastruktur, memungkinkan aplikasi untuk berjalan dengan konsisten di berbagai lingkungan.

3. Docker memungkinkan pengembang untuk mengemas semua dependensi aplikasi, termasuk sistem operasi, perpustakaan, dan kode aplikasi, ke dalam satu wadah yang dapat diinstal dan dijalankan di berbagai lingkungan.

4. Docker menggunakan file konfigurasi yang disebut Dockerfile untuk mengatur dan membangun kontainer. Dockerfile berisi instruksi untuk menambahkan dependensi aplikasi, memetakan port, dan menentukan entri point untuk menjalankan aplikasi.

5. Docker memiliki command-line interface (CLI) yang kuat, yang memungkinkan pengembang untuk membangun, menjalankan, dan mengelola kontainer dengan mudah.

6. Docker memiliki fitur jaringan dan penyimpanan terpisah, yang memungkinkan pengembang untuk membuat jaringan dan penyimpanan terisolasi untuk kontainer.

7. Docker menyediakan layanan orchestrasi yang memungkinkan pengembang untuk mengelola dan menyebarkan kontainer pada berbagai host dan lingkungan.

8. Docker memiliki ekosistem yang kaya dengan banyak alat dan layanan, termasuk Docker Compose untuk mengelola kontainer, Docker Swarm untuk orchestrasi, dan Docker Hub untuk berbagi dan menyimpan kontainer.

9. Docker dapat membantu pengembang mempercepat waktu pengembangan, meningkatkan portabilitas aplikasi, dan mengurangi biaya infrastruktur.

10. Docker dapat digunakan pada berbagai macam proyek, dari proyek kecil hingga proyek besar dan kompleks, dan pada berbagai lingkungan, termasuk lingkungan pengembangan, pengujian, dan produksi.


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
