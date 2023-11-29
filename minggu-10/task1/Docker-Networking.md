# SECTION 1 : Networking Basic
>**1. Perintah untuk koneksi di Docker**
jalankan perintah di bawah untuk melakukan konfigurasi pada container networks
![Alt text](ss/1.png)

>**2. Daftar Network pada Docker
jalankan perintah `docker network ls` untuk melihat container network yang ada pada Docker
![Alt text](ss/2.png)

>**3. Memeriksa Network pada docker**
jalanakn perintah `docker inspect network` untuk melihat konfigurasi detail pada container network
![Alt text](ss/3.png)
![Alt text](ss/22.png)

>**4. Daftar plugin driver pada network**
perintah `Docker info` bertujuan untuk menampilkan beberapa informasi mengenai instalasi docker
![Alt text](ss/4.png)

# SECTION 2 : Bridge Networking
>**1. Lakukan verifikasi untuk `bridge` dari docker network dengan menggunakan perintah `docker network ls`**
![Alt text](ss/5.png)
![Alt text](ss/23.png)
![Alt text](ss/24.png)
**Menampilkan juga bridge id nama bridge name**
![Alt text](ss/7.png)
![Alt text](ss/8.png)

>**2. Menghubungkan Container**
Buat Container baru dengan menjalankan perintah `docker run -dt ubuntu sleep infinity` tunggu hingga proses instalasi selesai
![Alt text](ss/9.png)
**lihat Container yang tadi sudah dibuat
![Alt text](ss/10.png)
![Alt text](ss/11.png)
**periksa `bridge connection` lagi, dengan menjalankan perintah `docker network inspect bridge` untuk melihat container sudah terlampir atau belum**
![Alt text](ss/12.png)

>**3. Mengecek konektivitas pada jaringan**
Menggunakan perintah `ping -c5 <IP4 Address`
![Alt text](ss/13.png)
![Alt text](ss/13.1.png)
**untuk memeriksa containerID, gunakan perintah berikut**
![Alt text](ss/14.png)
**jalankan shell di dalam container ubuntu dengan perintah `docker exec -it <CONTAINER ID> /bin/bash`**
![Alt text](ss/15.png)
**install ping program untuk mencoba perintah ping pada shell yang ada di dalam container ubuntu**
![Alt text](ss/16.png)
**jika sudah, maka coba lakukan ping ke `www.github.com`
![Alt text](ss/17.png)

>**4. Konfigurasi NAT untuk konektivitas eksternal**
Memulai container baru berdasar dari image NGINX dengan menjalankan perintah `docker run --name web1 -d -p 8080:80 nginx`**
![Alt text](ss/18.png)
**periksa status container dan port mappings dengan menggunakan perintah `docker ps`
![Alt text](ss/19.png)
![Alt text](ss/20.png)
![Alt text](ss/20.1.png)

# SECTION 3 : Overlay Networking
*Noted : pada step ini menggunakan IP address yang berbeda karena ada masalah dengan dockerclassrom sehingga harus merefresh halaman web.
>**1. Inisialisasi swarm dan mengecek operasi yang dilakukan berhasil**
![Alt text](ss/21.png)
![Alt text](ss/22.png)
**copykan perintah ini pada terminal kedua, sehingga terminal kedua bisa bergabung ke dalam 1 node dengan terminal yang pertama**
![Alt text](ss/25.png)
**cek apakah kedua node pada masing2 terminal sudah sukses dieksekusi dan aktif**
![Alt text](ss/26.png)

>**2. Membuat Overlay Network**
![Alt text](ss/27.png)
**verifikasi network apakah sudah sukses dibuat**
![Alt text](ss/28.png)
**cek juga pada terminal kedua, apakah sudah bisa dieksekusi**
![Alt text](ss/29.png)
**gunakan perintah `docker network inspect overnet` untuk melihat secara detail tentang network bernama `overnet` yang tadi sudah dibuat**
![Alt text](ss/30.png)
![Alt text](ss/31.png)
![Alt text](ss/32.png)

>**3. Membuat Service/Layanan**
**gunakan perintah `docker service create --name myservice \
--network overnet \
--replicas 2 \
ubuntu sleep infinity` di dalam network `overnet` dengan 2 taks**
![Alt text](ss/33.png)
![Alt text](ss/34.png)
**taks yang sudah berjalan di masing2 nodes pada swarm**
![Alt text](ss/35.png)
**apabila dicek dari node2(terminal kedua)**
![Alt text](ss/36.png)
**kita juga bisa menggunakan perintah `docker network inspect overnet` di terminal kedua untuk mendapatkan informasi detail tentang `overnet` network dan IP address yang berjalan pada terminal kedua**
![Alt text](ss/37.png)
![Alt text](ss/38.png)
![Alt text](ss/39.png)
![Alt text](ss/40.png)
![Alt text](ss/41.png)

>**4. Tes Network**
**eksekusi perintah `docket network inspect overnet` pada terminal 1(node1)**
![Alt text](ss/43.png)
![Alt text](ss/44.png)
**login ke service task dan install `ping command` di terminal kedua**
![Alt text](ss/45.png)
![Alt text](ss/46.png)
**kemudian coba lakukan ping ke alamat `10.0.1.5`**
![Alt text](ss/47.png)

>**5. Test Service Discovery**
![Alt text](image-1.png)
**Coba lakukan `ping` untuk `myservice`**
![Alt text](image.png)

