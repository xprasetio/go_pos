package db

import (
	"database/sql"
	"log"
)

func SeedRoles(db *sql.DB) {
	roles := []struct {
		code string
		name string
	}{
		{"SUPERADMIN", "Superadmin"},
		{"ADMIN", "Admin"},
		{"BARISTA", "Barista"},
		{"CUSTOMER", "Customer"},
	}

	for _, role := range roles {
		// Cek apakah role sudah ada
		var exists int
		checkQuery := `SELECT COUNT(*) FROM roles WHERE code = ?`
		err := db.QueryRow(checkQuery, role.code).Scan(&exists)
		if err != nil {
			log.Printf("Gagal cek role %s: %v", role.code, err)
			continue
		}

		// Jika role belum ada, insert
		if exists == 0 {
			insertQuery := `INSERT INTO roles (code, name, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`
			_, err := db.Exec(insertQuery, role.code, role.name)
			if err != nil {
				log.Printf("Gagal insert role %s: %v", role.code, err)
			} else {
				log.Printf("Seeder role %s sukses!", role.name)
			}
		} else {
			log.Printf("Role %s sudah ada, skip seeding", role.name)
		}
	}
	log.Println("Seeder roles selesai!")
}

func SeedMerchants(db *sql.DB) {
	merchants := []struct {
		code         string
		name         string
		merchantsApp string
		merchantsKey string
		baseUrl      string
		prefixTrxId  string
	}{
		{"MERCHANT001", "Coffee Shop Jakarta", "coffee_app", "key_123", "https://api.coffeeshop.com", "CSJ"},
		{"MERCHANT002", "Tea House Bandung", "tea_app", "key_456", "https://api.teahouse.com", "THB"},
	}

	for _, merchant := range merchants {
		// Cek apakah merchant sudah ada
		var exists int
		checkQuery := `SELECT COUNT(*) FROM merchants WHERE code = ?`
		err := db.QueryRow(checkQuery, merchant.code).Scan(&exists)
		if err != nil {
			log.Printf("Gagal cek merchant %s: %v", merchant.code, err)
			continue
		}

		// Jika merchant belum ada, insert
		if exists == 0 {
			insertQuery := `INSERT INTO merchants (code, name, merchants_app, merchants_key, base_url, prefix_trx_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`
			_, err := db.Exec(insertQuery, merchant.code, merchant.name, merchant.merchantsApp, merchant.merchantsKey, merchant.baseUrl, merchant.prefixTrxId)
			if err != nil {
				log.Printf("Gagal insert merchant %s: %v", merchant.code, err)
			} else {
				log.Printf("Seeder merchant %s sukses!", merchant.name)
			}
		} else {
			log.Printf("Merchant %s sudah ada, skip seeding", merchant.name)
		}
	}
	log.Println("Seeder merchants selesai!")
}

func SeedOutlets(db *sql.DB) {
	outlets := []struct {
		merchantID int64
		name       string
		province   string
		city       string
		address    string
		phone      string
		openTime   string
		closeTime  string
	}{
		{1, "Coffee Shop Jakarta Pusat", "DKI Jakarta", "Jakarta Pusat", "Jl. Sudirman No. 123", "021-1234567", "07:00", "22:00"},
		{1, "Coffee Shop Jakarta Selatan", "DKI Jakarta", "Jakarta Selatan", "Jl. Gatot Subroto No. 456", "021-7654321", "08:00", "23:00"},
		{2, "Tea House Bandung Center", "Jawa Barat", "Bandung", "Jl. Asia Afrika No. 789", "022-9876543", "09:00", "21:00"},
	}

	for _, outlet := range outlets {
		// Cek apakah outlet sudah ada
		var exists int
		checkQuery := `SELECT COUNT(*) FROM outlets WHERE name = ? AND merchant_id = ?`
		err := db.QueryRow(checkQuery, outlet.name, outlet.merchantID).Scan(&exists)
		if err != nil {
			log.Printf("Gagal cek outlet %s: %v", outlet.name, err)
			continue
		}

		// Jika outlet belum ada, insert
		if exists == 0 {
			insertQuery := `INSERT INTO outlets (merchant_id, name, province, city, address, phone, open_time, close_time, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
			_, err := db.Exec(insertQuery, outlet.merchantID, outlet.name, outlet.province, outlet.city, outlet.address, outlet.phone, outlet.openTime, outlet.closeTime)
			if err != nil {
				log.Printf("Gagal insert outlet %s: %v", outlet.name, err)
			} else {
				log.Printf("Seeder outlet %s sukses!", outlet.name)
			}
		} else {
			log.Printf("Outlet %s sudah ada, skip seeding", outlet.name)
		}
	}
	log.Println("Seeder outlets selesai!")
}

func RunSeeders(db *sql.DB) {
	log.Println("Memulai seeder...")

	SeedRoles(db)
	SeedMerchants(db)
	SeedOutlets(db)

	log.Println("Semua seeder selesai!")
}
