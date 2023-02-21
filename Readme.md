pertama tama bikin folder project
================================

bikin file main.go dan masukkan project awal
package main

func main() {


}
================================

lakukan
go mod init (nama project, disini saya pakai nama github saja misal github.com/masariuman/project)

================================

install gin dan gorm dan setting driver mysql/mariadb dengan cara

go get github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/mysql

================================

bikin folder controllers, models, dan routes

================================

pada models bikin file setup.go utk setup koneksi ke database (kode ikutkan seperti di kode filenya yg ada sekarang)

================================

bikin models untuk membuat koneksi ke tabe databasenya dan migrationnya

================================

jangan lupa bikin databasenya di db

================================

bikin file routes.go pada folder route sebagai superadmin route nya, dan bikin file bernama route untuk route yg diperlukan, disini memakai product.go. ikuti kodenya

================================

bikin folder lagi didalam controller utk controller yg digunakan seperti productcontroller pada contoh, dan bikin file productcontroller didalam folder productcontroller

================================

ikuti kodenya
