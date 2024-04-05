# Golang Rest API Kit Menggunakan Gin Gonic + PostgreSQL

- [Tentang Kit](#tentang-kit)
- [Requirements](#requirements)
- [Fitur](#fitur)
- [Folder Struktur](#folder-struktur)
- [Notes](#notes)

## Tentang Kit

kit ini biasa saya gunakan untuk membuat projek Rest API (from scratch) menggunakan bahasa pemrograman golang. untuk database yang digunakan pada kit ini adalah PostgreSQL. dalam kit ini saya sertakan contoh CRUD yang berada di folder `domain/posts`.

## Requirements

- Golang minimum versi 1.18
- PostgreSQL minimum versi 15

## Fitur

- Autentikasi menggunakan JWT
- Konfigurasi menggunakan dotenv dan Viper
- Object Relational Mapping menggunakan Gorm
- uuid sebagai primary key
- Error handlers untuk response http
- Pagination

## Folder Struktur

```
├── domain
│   └── auth
│       ├── controllers
│       |    └── auth.controller.go
│       ├── models
│       ├── repositories
│       |   ├── signin.repo.go
│       |   └── signup.repo.go
│       ├── requests
│       |   ├── signin.request.go
│       |   └── signup.request.go
│       ├── responses
│       |   └── token.response.go
│       ├── routes
│       |   └── auth.routes.go
│       ├── services
│       |   ├── refresh.token.service.go
│       |   ├── signin.service.go
│       |   ├── signout.service.go
│       |   └── signup.service.go
│   └── posts
│       ├── controllers
│       |   └── post.controller.go
│       ├── models
│       |   └── post.model.go
│       ├── repositories
│       |   ├── create.repo.go
│       |   ├── delete.repo.go
│       |   ├── get.all.repo.go
│       |   ├── get.by.id.repo.go
│       |   ├── get.pagination.repo.go
│       |   └── update.repo.go
│       ├── requests
│       |   ├── create.request.go
│       |   ├── get.pagination.request.go
│       |   └── update.request.go
│       ├── responses
│       |   └── post.response.go
│       ├── routes
│       |   └── post.routes.go
│       ├── services
│       |   ├── create.service.go
│       |   ├── delete.service.go
│       |   ├── get.all.service.go
│       |   ├── get.by.id.service.go
│       |   ├── get.pagination.service.go
│       |   └── update.service.go
│   └── users
│       ├── controllers
│       |    └── user.controller.go
│       ├── models
│       |    └── user.model.go
│       ├── repositories
│       |   └── get.by.logged.in.repo.go
│       ├── requests
│       |   └── get.by.logged.in.request.go
│       ├── responses
│       |   └── user.response.go
│       ├── routes
│       |   └── user.routes.go
│       ├── services
│       |   └── get.by.logged.in.service.go
├── error-handlers
│   ├── bad.gateway.error.go
│   ├── bad.request.error.go
│   ├── conflict.error.go
│   ├── forbidden.error.go
│   ├── not.found.error.go
│   └── unauthorized.error.go
├── helpers
│   ├── meta.helper.go
│   ├── pagination.helper.go
│   └── seo.helper.go
├── initializers
│   ├── connectDB.go
│   └── loadEnv.go
├── middleware
│   └── deserialize-user.go
├── migrate
│   └── migrate.go
├── responses
│   └── pagination.response.go
├── utils
│   ├── password.go
│   └── token.go
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── README.md
```

Deskripsi dari Struktur Folder:

- `domain` berisi modul-modul yang digunakan di projek.
- `error-handlers` berisi tentang http error response.
- `helpers` berisi librari untuk penggunaan.
- `initializers` konfigurasi yang dijalankan di awal, seperti koneksi database, env, dll.
- `middleware` middleware yang digunakan di projek.
- `migrate` file-file migrasi untuk pembuatan schema database.
- `responses` berisi global response http.
- `utils` untuk third party, seperti jwt, dll.

## Notes

Untuk menjalankan projek ini, gunakan command seperti berikut:

- Clone dari github

  ```
  git clone https://github.com/agungsuprayitno/go-rest-api.git
  ```

- Download dependencies

  ```
  go mod tidy
  ```

- Buat file .env

  ```
  cp .env.example .env
  ```

- Isi `ACCESS_TOKEN_PRIVATE_KEY`, `ACCESS_TOKEN_PUBLIC_KEY`, `REFRESH_TOKEN_PRIVATE_KEY`, `REFRESH_TOKEN_PUBLIC_KEY`,
  dengan men-generate base64 encode public key.
  contoh format:

  ```
  base64encode(
    -----BEGIN PUBLIC KEY-----
    AAAAB3NzaC1yc2EAAAADAQABAAABgQDe2EpJLmStfjGlLADR2Xndl5s9lcaxE6POWgecSL62JvTbWwVZVraukQDK/NHq7HZ9OhoH4ZqHD1C4gvZTyK+avB
    -----END PUBLIC KEY-----)
  ```

  contoh hasil base64 encode :

  ```
  LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KQUFBQUIzTnphQzF5YzJFQUFBQURBUUFCQUFBQmdRRGUyRXBKTG1TdGZqR2xMQURSMlhuZGw1czlsY2F4RTZQT1dnZWNTTDYySnZUYld3VlpWcmF1a1FESy9OSHE3SFo5T2hvSDRacUhEMUM0Z3ZaVHlLK2F2QgotLS0tLUVORCBQVUJMSUMgS0VZLS0tLS0=
  ```

- Run migrasi database

  ```
  go run migrate/migrate.go
  ```

- Run

  ```
  go run main.go
  ```
