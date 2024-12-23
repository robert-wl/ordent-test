# Artikel Online - Backend

## Deskripsi
Backend yang dibuat adalah backend untuk platform Artikel Online, dibuat menggunakan Go dan framework Gin.

## Tools dan Package
1. Gin (Framework)
2. Gorm (ORM)
3. Air (Live Reload)
4. PostgreSQL (Database)
5. JWT (Authentication)
6. Swagger (API Documentation)
7. Docker (Container)

## Cara Menjalankan Aplikasi

Sebelum menjalankan aplikasi, pastikan sudah menginstall Go dan PostgreSQL.
Copy `.env.example` menjadi `.env` dan sesuaikan dengan konfigurasi yang ada.

Script untuk menjalankan dan mendevelop aplikasi ini ada didalam file `Makefile`.
1. `make run` - Menjalankan aplikasi dengan live reload
2. `make swagger` - Mengupdate JSON swagger
3. `make test` - Menjalankan unit testing (hanya beberapa)

Tanpa menggunakan Makefile, berikut adalah cara menjalankan aplikasi:
`go run ./cmd/app/main.go`

Untuk menggunakan `docker`, berikut adalah cara menjalankan aplikasi:
`docker compose up --build`


## Fitur
Berikut adalah fitur yang ada pada aplikasi ini, (catatan, semua endpoint memiliki baseUrl `/api/v1`):

1. Authentication dan Authorization

| Fitur        | Method | Endpoint         | Authenticated | Deskripsi                               |
|--------------|--------|------------------|---------------|-----------------------------------------|
| Register     | POST   | `/auth/register` | ❌             | Register user baru                      |
| Login        | POST   | `/auth/login`    | ❌             | Login user                              |
| Current User | GET    | `/auth/me`       | ✅             | Mendapatkan data user yang sedang login |


2. User Management

| Fitur          | Method | Endpoint         | Authenticated | Deskripsi                                                                                       |
|----------------|--------|------------------|---------------|-------------------------------------------------------------------------------------------------|
| View All User  | GET    | `/users`         | ✅ `admin`     | Untuk melihat semua user yang ada, <br/>dapat diberikan parameter `page`, `limit` dan `search`  |
| View All Admin | GET    | `/users/admins`  | ✅ `admin`     | Untuk melihat semua admin yang ada, <br/>dapat diberikan parameter `page`, `limit` dan `search` |
| Get User       | GET    | `/users/{id}`    | ✅ `admin`     | Untuk melihat sebuah user berdasarkan `id`                                                      |
| Promote User   | PUT    | `/users/promote` | ✅ `admin`     | Untuk mengubah role user menjadi admin                                                          |
| Demote Admin   | PUT    | `/users/demote`  | ✅ `admin`     | Untuk mengubah role admin menjadi user                                                          |


3. Article Management

| Fitur             | Method | Endpoint         | Authenticated       | Deskripsi                                                                                         |
|-------------------|--------|------------------|---------------------|---------------------------------------------------------------------------------------------------|
| View All Articles | GET    | `/articles`      | ✅                   | Untuk melihat semua artikel yang ada, <br/>dapat diberikan parameter `page`, `limit` dan `search` |
| Get Article       | GET    | `/articles/{id}` | ✅                   | Untuk mendapatkan sebuah artikel berdasarkan `id`                                                 |
| Create Article    | POST   | `/articles`      | ✅                   | Untuk membuat sebuah artikel baru                                                                 |
| Update Article    | PUT    | `/articles/{id}` | ✅ `admin` `creator` | Untuk mengubah sebuah artikel                                                                     |
| Delete Article    | DELETE | `/articles/{id}` | ✅ `admin` `creator` | Untuk menghapus sebuah artikel                                                                    |

3. Comment Management

| Fitur                 | Method | Endpoint                  | Authenticated       | Deskripsi                                                                                                  |
|-----------------------|--------|---------------------------|---------------------|------------------------------------------------------------------------------------------------------------|
| View Article Comments | GET    | `/articles/{id}/comments` | ✅                   | Untuk melihat semua komen dari sebuah artikel. <br/>Dapat diberikan parameter `page`, `limit` dan `search` |
| Get Comment           | GET    | `/comments/{id}`          | ✅                   | Untuk mendapatkan sebuah komentar berdasarkan `id`                                                         |
| Create Comment        | POST   | `/comments`               | ✅                   | Untuk membuat sebuah komentar baru                                                                         |
| Update Comment        | PUT    | `/articles/{id}`          | ✅ `admin` `creator` | Untuk mengubah sebuah komentar                                                                             |
| Delete Comment        | DELETE | `/articles/{id}`          | ✅ `admin` `creator` | Untuk menghapus sebuah komentar                                                                            |


## Database

Database yang digunakan adalah PostgreSQL, berikut adalah struktur database yang digunakan:

![Database Schema](./.github/images/database.png)'

# Arsitektur Aplikasi

Aplikasi yang dibuat menggunakan arsitektur mirip DDD (Domain Driven Design),
berikut adalah struktur dari aplikasi:



![Architecture](./.github/images/architecture.png)