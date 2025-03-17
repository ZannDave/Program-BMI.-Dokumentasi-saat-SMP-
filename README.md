# Proyek BMI

**Dibuat oleh:** Zann Dave, 15 tahun  

## Deskripsi  
Proyek ini adalah aplikasi berbasis **Golang** untuk menghitung **BMI (Body Mass Index)** dan mengklasifikasikannya ke dalam kategori seperti **Underweight, Normal, Overweight, dan Obesitas yang disimpan didalam database **Postgresql****.  

## Cara Menjalankan  
### Persyaratan  
- **Go** harus terinstal di sistem Anda. Jika belum, unduh dari [golang.org](https://golang.org/dl/).  
- Pastikan ada koneksi ke database yang dikonfigurasi dalam folder `config`.  

### Langkah-langkah  
1. Clone atau ekstrak proyek ini.  
2. Buka terminal, lalu pindah ke direktori proyek:  
   ```sh
   cd proyekBMI
3. Jalankan perintah berikut untuk unduh dependensi:
   go mod tidy
4. Jalankan program:
   go run pBMI.go / go run .
   
## Struktur Proyek

Folder Utama : 

proyekBMI, Isi nya :

1. pBMI.go
2. go.mod
3. go.sum
4. .gitignore
5. .env
6. Folder 'config', Isi nya :
   config.go

## Depedensi
• Go Modules ( go.mod dan go.sum )
• Database (Di folder config.go ( Disembunyikan ) )

## License / Lisensi
Proyek ini bersifat open-source!
