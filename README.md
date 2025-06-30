# POS-GO

Aplikasi Point of Sale (POS) sederhana berbasis Go, Echo, Dependency Injection, dan MySQL.

## Fitur

- Registrasi dan login user
- JWT authentication
- Dependency Injection (Sarulabs/DI)
- Logger (Logrus)
- Validasi input
- Migrasi dan seeder database otomatis

## Struktur Project

- `cmd/server/main.go` — Entry point aplikasi
- `config/config.yaml` — Konfigurasi aplikasi & database
- `internal/user/` — Modul user (handler, service, repository)
- `pkg/` — Paket utilitas (db, jwt, logger, middleware, response, validator)
- `router/` — Routing aplikasi
- `container/` — Dependency Injection container

## Instalasi

1. **Clone repository**
   ```bash
   git clone <repo-url>
   cd pos-go
   ```
2. **Install dependencies**
   ```bash
   go mod tidy
   ```
3. **Konfigurasi**
   Edit file `config/config.yaml` sesuai kebutuhan:
   ```yaml
   app:
     name: pos-go
     port: 8081
     jwt_secret: your-secret-key-here
     jwt_exp: 24 # jam
   database:
     host: localhost
     port: 3306
     user: root
     password:
     name: pos_go
   log:
     level: debug
     format: json
   ```
4. **Jalankan aplikasi**
   ```bash
   go run cmd/server/main.go
   ```

## Endpoint Utama

| Method | Endpoint  | Deskripsi       |
| ------ | --------- | --------------- |
| POST   | /register | Registrasi user |
| POST   | /login    | Login user      |

### Contoh Request

#### Register

```http
POST /register
Content-Type: application/json

{
  "name": "Budi",
  "email": "budi@mail.com",
  "password": "rahasia"
}
```

#### Login

```http
POST /login
Content-Type: application/json

{
  "email": "budi@mail.com",
  "password": "rahasia"
}
```

## Lisensi

MIT
