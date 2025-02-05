# ESCommerce

ESCommerce adalah **REST API** yang menyediakan fitur marketplace sederhana seperti e-commerce pada umumnya. Proyek ini dibangun menggunakan **Golang** dengan framework **GIN**, dan **PostgreSQL** sebagai database utamanya.

## ✨ Fitur Utama

- **User Authentication**  
  Autentikasi menggunakan **JWT** dan **bcrypt** untuk keamanan. Terdapat dua jenis user:  
  - **Seller**: Dapat menambahkan dan mengelola produk.  
  - **Customer**: Dapat melihat produk dan menambahkan produk ke keranjang belanja.

- **Product Management**  
  CRUD (Create, Read, Update, Delete) untuk mengelola produk.

- **Shopping Cart**  
  Fitur keranjang belanja untuk pelanggan, termasuk menambahkan, menghapus, dan mengelola item di keranjang belanja.

## 🛠️ Tech Stack

- **Golang** dengan framework **GIN**  
- **GORM** sebagai ORM  
- **PostgreSQL** sebagai database  
- **Go Swagger** untuk dokumentasi API

## 🚀 Menjalankan Proyek

Pastikan Anda sudah menginstall **Golang** dan **PostgreSQL** di mesin Anda.

1. Clone repositori ini:
   git clone https://github.com/muhammadsaefulr/escommerce.git
   cd escommerce

2. Buat file `.env` untuk konfigurasi environment:
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name

3. Jalankan aplikasi:
   go run main.go

## 📚 Dokumentasi API

Dokumentasi API tersedia di endpoint `http://localhost:8080/swagger/index.html` setelah aplikasi dijalankan.

## 🤝 Kontribusi

Kontribusi sangat terbuka! Silakan buat pull request atau laporkan issue jika Anda menemukan bug atau memiliki ide untuk fitur baru.

## 📜 Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
