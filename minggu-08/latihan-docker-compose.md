# Mencoba Docker Compose

>**1. buat folder baru untuk mencoba `Docker Compose`, pada langkah ini buat nama folder `composetest`**
![Alt text](ss-docker/image.png)

>**2. Di dalam folder `composetest` tadi, tambahkan file `app.py` dan masukkan kode ini:**
![Alt text](ss-docker/image-1.png)
![Alt text](ss-docker/image-2.png)

>**3. Buat file dengan nama `requirements.txt` di dalam folder `composetes` dan masukkan perintah berikut:**
![Alt text](ss-docker/image-3.png)

>**4. Buat `Dockerfile` untuk membangun docker image, dan tambahkan perintah berikut:**
![Alt text](ss-docker/image-4.png)

>**5. Buat file dengan nama `compose.yaml` pada folder yang tadi sudah dibuat, kemudian masukkan perintah berikut:**
![Alt text](ss-docker/image-5.png)
![Alt text](ss-docker/image-6.png)

>**6. Mulai coba dan jalankan applikasi dengan Docker Compose**
![Alt text](ss-docker/image-7.png)
![Alt text](ss-docker/image-8.png)

>**7. Coba ketikkan alamat `http://localhost:8000/` pada browser dan perhatikan applikasi yang sudah berjalan**
![Alt text](ss-docker/image-9.png)

>**8. Apabila di refresh lagi maka akan berubah jumlah angkanya**
![Alt text](ss-docker/image-10.png)

>**9. Jika dicek dengan perintah `docker image ls` maka list repository yang muncul seperti tadi yang sudah dibuat**
![Alt text](ss-docker/image-11.png)

>**10. Coba dilakukan edit pada file `compose.yaml` untuk menambahkan bind mount**
![Alt text](ss-docker/image-12.png)

>**11. Kemudian coba jalankan lagi dengan perintah `docker compose up`**
![Alt text](ss-docker/image-13.png)

>**12. Check pada browser dengan link `http://localhost:8000`**
![Alt text](ss-docker/image-14.png)

>**13. Mengupdate aplikasi untuk memastikan bahwa perubahan pada output kode tidak mempengaruhi penambahan jumlah waktu**
![Alt text](ss-docker/image-15.png)
![Alt text](ss-docker/image-16.png)

>**14. Mencoba jalankan dengan perintah lain**
![Alt text](ss-docker/image-17.png)

>**15. Menghentikan aplikasi yang berjalan pada docker compose**
![Alt text](ss-docker/image-18.png)