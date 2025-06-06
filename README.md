

# 🩺 Monitoring API Service

Project sederhana untuk **memantau kesehatan (health check)** dari layanan API lain secara berkala. Dibuat menggunakan [Gofr](https://gofr.dev) dan Go.

## ✨ Fitur

- Membaca daftar service dari file JSON (`config/url.json`)
- Melakukan pengecekan berkala (10 detik sekali)
- Logging status `up/down` dari tiap service
- Menyediakan endpoint `/todos` (contoh layanan yang dipantau)

## 📁 Struktur Proyek
```bash
.
├── config/
│   └── url.json        # Daftar service yang akan dimonitor
├── service/
│   └── monitor.go      # Logika monitoring service
├── utils/
│   └── logger.go       # Logger Monitor
├── .gitignore          
├── main.go             # Entry point + definisi endpoint
├── go.mod
└── README.md
```

## 🚀 Menjalankan Project

1. Clone repo ini:
   ```bash
   git clone https://github.com/Dziqha/monitoringApiservice.git
   cd monitoringApiservice
   ```

2. Jalankan program:

   ```bash
   go run main.go
   ```

3. Akses endpoint:

   * `GET /todos` – Menampilkan daftar todo
   * `POST /todo` – Menambahkan todo
   * `GET /healty` – Mengecek apakah monitoring berjalan

## 🧪 Contoh `url.json`

```json
[
  { "name": "service-todos", "url": "http://localhost:8000/todos" },
  { "name": "service-health", "url": "http://localhost:8000/healty" }
]
```

## 🛠️ Tools

* Go
* Gofr framework (`gofr.dev`)

## 📝 Catatan

* Log monitoring akan ditampilkan di terminal.
* Monitoring berjalan di background (pakai goroutine).
