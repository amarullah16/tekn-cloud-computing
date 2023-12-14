# Docker Orchestration Hands-On Lab

## Section 1 : Configure Swarm Mode
>**1.Jalankan perintah `docker run -dt ubuntu sleep infinity` pada node1 atau terminal1**
![Alt text](screenshoot/image.png)
![Alt text](screenshoot/image-1.png)

>**2. Buat node Manager, jalankan perintah `docker swarm init pada node1`**
![Alt text](screenshoot/image-2.png)
![Alt text](screenshoot/image-3.png)
![Alt text](screenshoot/image-4.png)

>**3. Gabungkan worker node2 dan node3 ke Swarm pada node1**
![Alt text](screenshoot/image-5.png)
![Alt text](screenshoot/image-6.png)
**Hasilnya seperti gambar di bawah ini**
![Alt text](screenshoot/image-7.png)
![Alt text](screenshoot/image-8.png)

## Section 2 : Menyebarkan aplikasi di beberapa host
>**1.Menerapkan komponen aplikasi sebagai layanan Docker**
![Alt text](screenshoot/image-9.png)
![Alt text](screenshoot/image-10.png)
**verifikasi apakah `service create` sudah diterima oleh Swarm Manager**
![Alt text](screenshoot/image-11.png)

## Section 3 : Meluaskan cakupan aplikasi
>**1.Tingkatkan jumlah kontainer dalam layanan sleep-app menjadi 7 dengan perintah `docker service update --replicas 7 sleep-app`.**
![Alt text](screenshoot/image-12.png)
![Alt text](screenshoot/image-13.png)
**Periksa hasil perintah**
![Alt text](screenshoot/image-14.png)
![Alt text](screenshoot/image-15.png)
![Alt text](screenshoot/image-16.png)

>**2. Ubah layanan kembali menjadi hanya 4 container dengan perintah `docker service update --replicas 4 sleep-app`**
![Alt text](screenshoot/image-17.png)
![Alt text](screenshoot/image-18.png)
![Alt text](screenshoot/image-19.png)
![Alt text](screenshoot/image-20.png)

## Section 4 : bersihkan node and jadwalkan ulang container
>**1. Pada node1**
![Alt text](screenshoot/image-21.png)
![Alt text](screenshoot/image-22.png)

>**2. Pada node2**
![Alt text](screenshoot/image-23.png)

>**3. Pada langkah ini, kita akan mencoba mengeluarkan node2 dari service. Jalankan perintah `docker node update --availability drain node2` pada terminal node1**
![Alt text](screenshoot/image-24.png)
**Hasilnya seperti berikut**
![Alt text](screenshoot/image-25.png)
![Alt text](screenshoot/image-26.png)

>**4. Pindah ke terminal node2, jalankan perintah `docker ps` untuk melihat hasil dari node2 yang sudah dikeluarkan dari layanan tadi**
![Alt text](screenshoot/image-27.png)

>**5. Terakhir, periksa kembali layanan pada node1 untuk memastikan bahwa kontainer telah dijadwalkan ulang**
![Alt text](screenshoot/image-28.png)
![Alt text](screenshoot/image-29.png)
![Alt text](screenshoot/image-30.png)

## Cleaning Up
>**1. Jalankan perintah docker service rm sleep-app pada node1 untuk menghapus layanan bernama myservice**
![Alt text](screenshoot/image-31.png)
![Alt text](screenshoot/image-32.png)

>**2. Bisa juga menggunakan perintah docker kill <CONTAINER ID> pada node1 untuk mematikan kontainer yang telah kita mulai di awal**
![Alt text](screenshoot/image-33.png)
![Alt text](screenshoot/image-34.png)

>**3. Terakhir, mari kita hapus node1, node2, dan node3 dari Swarm. Kita dapat menggunakan perintah `docker swarm leave --force` untuk melakukannya**
**Pada node1**
![Alt text](screenshoot/image-35.png)
**Pada node2 dan node3**
![Alt text](screenshoot/image-36.png)
