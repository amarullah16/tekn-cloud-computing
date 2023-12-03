# Stage Setup
>**1. Cloning repository demo, mengganti, direktori kerja, dan cek branch demo workout**
![Alt text](<ss/Screenshot 2023-12-03 123752.png>)

>**2. Basic Link Extractor Script**
**- Periksa branch `step0` dan daftar file yang ada disana**
![Alt text](<ss/Screenshot 2023-12-03 123915.png>)
![Alt text](<ss/Screenshot 2023-12-03 124051.png>)
![Alt text](<ss/Screenshot 2023-12-03 124220.png>)

# Step 1 : Containerized Link Extractor Script
>**1. Periksa brach `step1` dan lihat daftar list yang ada di dalamnya**
![Alt text](<ss/Screenshot 2023-12-03 124321.png>)

>**2. Cek file `Dockerfile`**
![Alt text](<ss/Screenshot 2023-12-03 124349.png>)

>**3. Lakukan Perintah ini untuk menyiapkan dockerfile**
![Alt text](<ss/Screenshot 2023-12-03 124742.png>)
![Alt text](<ss/Screenshot 2023-12-03 124851.png>)

 >**4. Coba jalankan 1 container dengan image dan extract link dari 1 halaman website**
 ![Alt text](<ss/Screenshot 2023-12-03 125004.png>)
 ![Alt text](<ss/Screenshot 2023-12-03 125112.png>)

 >**5. Update Script**
![Alt text](<ss/Screenshot 2023-12-03 125207.png>)

 >**6. Buat image lagi dan perhatikan perubahannya**
 ![Alt text](<ss/Screenshot 2023-12-03 125245.png>)
![Alt text](<ss/Screenshot 2023-12-03 125314.png>)

>**7. Jalankan container menggunakan script `linkextractor:step2` pada image**
![Alt text](<ss/Screenshot 2023-12-03 125347.png>)
![Alt text](<ss/Screenshot 2023-12-03 125428.png>)
![Alt text](<ss/Screenshot 2023-12-03 125443.png>)

# Step 2 : Link Extractor API Service**
>**1. Periksa `step3`**
![Alt text](<ss/Screenshot 2023-12-03 125524.png>)
**- buka `Dockerfile`**
![Alt text](<ss/Screenshot 2023-12-03 125618.png>)
**- kemudian buka file `main.py` untuk melihat isi file**
![Alt text](<ss/Screenshot 2023-12-03 125703.png>)

>**2. Buat image baru dengan beberapa perubahan, dan coba dijalankan**
![Alt text](<ss/Screenshot 2023-12-03 125730.png>)
![Alt text](<ss/Screenshot 2023-12-03 125856.png>)

>**3. Buat request HTTP pada form `/api/<url>` untuk menyambung kepada server dan mengambil respon yang berisi `extracted link`**
![Alt text](<ss/Screenshot 2023-12-03 130003.png>)
![Alt text](<ss/Screenshot 2023-12-03 130046.png>)

# Step 4 : Extractor API and Web Front End Services
>**1. Periksa `step4`**
![Alt text](<ss/Screenshot 2023-12-03 130131.png>)

>**2. Buka file `docker-compose.yml` dan perhatikan isinya**
![Alt text](<ss/Screenshot 2023-12-03 130205.png>)

>**3. Perhatikan file `index.php` pada folder www**
![Alt text](<ss/Screenshot 2023-12-03 130619.png>)
![Alt text](<ss/Screenshot 2023-12-03 130639.png>)
![Alt text](<ss/Screenshot 2023-12-03 130650.png>)
![Alt text](<ss/Screenshot 2023-12-03 130719.png>)
![Alt text](<ss/Screenshot 2023-12-03 130731.png>)
![Alt text](<ss/Screenshot 2023-12-03 130747.png>)
![Alt text](<ss/Screenshot 2023-12-03 130758.png>)

>**4. Sekarang coba jalankan servicenya menggunakan `docker compose`**
![Alt text](<ss/Screenshot 2023-12-03 130844.png>)
![Alt text](<ss/Screenshot 2023-12-03 130912.png>)
![Alt text](<ss/Screenshot 2023-12-03 130925.png>)

>**5. Sekarang periksa container yang sedang berjalan**
![Alt text](<ss/Screenshot 2023-12-03 131002.png>)

>**6. Sekarang sudah bisa berkomunikasi ke API service**
![Alt text](<ss/Screenshot 2023-12-03 131042.png>)
![Alt text](<ss/Screenshot 2023-12-03 131055.png>)

>**7. Dengan cara yang hampir mirip, mari kita coba ubah sedikit menjadi `Super Link Extractor`**
![Alt text](<ss/Screenshot 2023-12-03 131130.png>)
![Alt text](<ss/Screenshot 2023-12-03 131142.png>)

>**8. Clean the Git Tracking**
![Alt text](<ss/Screenshot 2023-12-03 131220.png>)

# Stage 5 : Redis Service for Caching
>**1. Periksa `step5`**
![Alt text](<ss/Screenshot 2023-12-03 131243.png>)

>**2. Tambahkan `Dockerfile` ke dalam folder `www`**
![Alt text](<ss/Screenshot 2023-12-03 131305.png>)

>**3. Buka file `api/main.py` pada API server yang mana akan menggunakan Redis Cache**
![Alt text](<ss/Screenshot 2023-12-03 131328.png>)
![Alt text](<ss/Screenshot 2023-12-03 131344.png>)

>**4. Buka file `docker-compose.yml`**
![Alt text](<ss/Screenshot 2023-12-03 131408.png>)

>**5. Nyalakan service**
![Alt text](<ss/Screenshot 2023-12-03 131441.png>)
![Alt text](<ss/Screenshot 2023-12-03 131456.png>)

>**6. Lakukan pengecekan pada Redis**
![Alt text](<ss/Screenshot 2023-12-03 132506.png>)

>**7. Apabila service sudah berjalan, matikan service kembali**
![Alt text](<ss/Screenshot 2023-12-03 132535.png>)

# Step 6 : Swap Python API Service with Ruby
>**1. Periksa daftar file yang ada di `step6`**
![Alt text](<ss/Screenshot 2023-12-03 132727.png>)

>**2. Buka file `linkextractor.rb` dan `Dockerfile` pada folder `api`**
![Alt text](<ss/Screenshot 2023-12-03 132827.png>)
![Alt text](<ss/Screenshot 2023-12-03 132839.png>)
![Alt text](<ss/Screenshot 2023-12-03 132859.png>)

>**3. Buka file `docker-compose.yml`**
![Alt text](<ss/Screenshot 2023-12-03 132922.png>)

>**4. Nyalakan service**
![Alt text](<ss/Screenshot 2023-12-03 132951.png>)
![Alt text](<ss/Screenshot 2023-12-03 133021.png>)

>**5. Sekarang kita sudah bisa meghubungkan akses ke API menggunakan nomor port yang sudah diupdate**
![Alt text](<ss/Screenshot 2023-12-03 133048.png>)
![Alt text](<ss/Screenshot 2023-12-03 133204.png>)

>**6. Kita sudah bisa mematikan server**
![Alt text](<ss/Screenshot 2023-12-03 133331.png>)

>**7. Kita juga masih bisa melihat riwayat service yang tadi sudah berjalan**
![Alt text](<ss/Screenshot 2023-12-03 133445.png>)
