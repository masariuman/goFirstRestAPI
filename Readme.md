<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 4.0 Attributions license][cc4-by].*
[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/gin-gonic/gin/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin)](https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin)


pertama tama bikin folder project

----

bikin file main.go dan masukkan project awal
    package main

    func main() {


    }

----

lakukan
    go mod init (nama project, disini saya pakai nama github saja misal github.com/masariuman/project)

----

install gin dan gorm dan setting driver mysql/mariadb dengan cara
```sh
go get github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/mysql
```

----

bikin folder controllers, models, dan routes

----

pada models bikin file setup.go utk setup koneksi ke database (kode ikutkan seperti di kode filenya yg ada sekarang)

----

bikin models untuk membuat koneksi ke tabe databasenya dan migrationnya

----

jangan lupa bikin databasenya di db

----

bikin file routes.go pada folder route sebagai superadmin route nya, dan bikin file bernama route untuk route yg diperlukan, disini memakai product.go. ikuti kodenya

----

bikin folder lagi didalam controller utk controller yg digunakan seperti productcontroller pada contoh, dan bikin file productcontroller didalam folder productcontroller

----

ikuti kodenya
