# dataekspor-be

Repository ini digunakan untuk development _back-end_ dari web
[dataekspor](https://dataekspor.com)

# Dokumentasi
Dokumentasi untuk API dataekspor dapat dilihat di

https://documenter.getpostman.com/view/22520213/VUxVpin3

# Stack
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm)
- [jwt-go](https://github.com/golang-jwt/jwt)

# Getting Started
1. Siapkan folder untuk meletakkan repository ini secara lokal, pastikan sudah menginstall [Git](https://git-scm.com/downloads) dan [Go](https://go.dev/doc/install).
2. Buka terminal dan jalankan `git init` untuk menginisialisasi repository
3. Lalu jalankan `git remote add origin https://github.com/Data-Ekspor/dataekspor-be.git`
4. Setelah berhasil, jalankan `git fetch origin master`
5. Setelah file-file terunduh, jalankan `git pull origin master`

# Development Lifecycle
1. Branch utama yang digunakan untuk production adalah branch `master`
2. Branch untuk pengembangan fitur dinamai dengan `nama/fitur`, contoh : `alex/daftar-eksportir`
3. Setelah membuat fitur, lakukan `git add .`, `git commit -m "message_commit"`, lalu `git push -u origin nama_branch`. (Gunakan [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/#specification) untuk referensi commit) 
4. Untuk mendapatkan update terbaru dari repository lakukan `git pull origin master`
5. Apabila perubahan pada branch individu telah benar dan akan diupload ke branch master, pindah ke branch master dengan `git checkout master`, kemudian `git merge nama_branch_individu` untuk menyatukan perubahan dari branch individu, lalu jangan lupa melakukan `git push origin master`

# Menjalankan Server
1. Pastikan Go versi 1.18 Sudah diinstall
2. Masuk ke folder dataekspor-be
3. Buka terminal dan jalankan `go mod tidy` untuk mendownload dependensi
4. Buka terminal dan jalankan perintah `go run main.go` pada root directory yang terdapat file main.go
5. API dapat diakses pada port :8080

<!-- Test aman -->