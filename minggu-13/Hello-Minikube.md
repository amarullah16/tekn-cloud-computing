# HELLO MINIKUBE

>**1. Membuat cluster minikube**
![Alt text](screenshoot/image.png)

>**2. Buka dashboard**
![Alt text](screenshoot/image-1.png)
![Alt text](screenshoot/image-2.png)

## BUAT DEPLOYMENT
>**1. Gunakan perintah `kubectl create` untuk membuat Deployment yang mengelola sebuah Pod. Pod menjalankan container berdasarkan image Docker yang disediakan**
![Alt text](screenshoot/image-3.png)

>**2. Lihat hasil deployment**
![Alt text](screenshoot/image-4.png)

>**3. Lihat pod**
![Alt text](screenshoot/image-5.png)

>**4. Melihat event cluster**
![Alt text](screenshoot/image-6.png)
![Alt text](screenshoot/image-7.png)

>**5. Melihat konfigurasi pada `kubectl`**
![Alt text](screenshoot/image-8.png)
![Alt text](screenshoot/image-9.png)

>**6. Melihat log aplikasi untuk container dalam pod**
![Alt text](screenshoot/image-10.png)

## Buat sebuah service
>**1. Mengekspos Pod ke internet publik menggunakan perintah kubectl expose**
![Alt text](screenshoot/image-11.png)

>**2. Melihat service yang telah dibuat sebelumnya**
![Alt text](screenshoot/image-12.png)

>**3. Jalankan perintah `minikube service hello-node`. setelah itu buka browser yang akan menyajikan aplikasi dan menunjukkan respons aplikasi**
![Alt text](screenshoot/image-13.png)
![Alt text](screenshoot/image-14.png)

## MENGAKTIFKAN ADD-ONS
>**1. Lihat daftar add-on yang saat ini didukung**
![Alt text](screenshoot/image-15.png)
![Alt text](screenshoot/image-16.png)

>**2. Aktifkan add-on, sebagai contoh kita akan mengaktifkan addons `metrics-server`**
![Alt text](screenshoot/image-17.png)

>**3. Melihat Pod dan Layanan yang dibuat dengan menginstal add-ons tersebut**
![Alt text](screenshoot/image-18.png)

>**4. Menonaktifkan `metrics-server`**
![Alt text](screenshoot/image-19.png)

## Cleaning Up
>**1. Sekarang kita dapat membersihkan sumber daya yang tadi sudah dibuat di dalam cluster**
![Alt text](screenshoot/image-20.png)

>**2. Menghentikan Minikube cluster**
![Alt text](screenshoot/image-21.png)

>**3. Jika memang tidak dibutuhkan, bisa juga dengan menghapus Minikube VM**
![Alt text](screenshoot/image-22.png)


