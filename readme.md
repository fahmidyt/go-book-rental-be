<h1 align="center">GO: Book Rental API Service</h1>
<h3 align="center">( GIN Gorm <GG> )</h3>
<br/>

> Basic API Service (Book Rental) using Golang

## Prerequisites
- go v1.17.5 (tested & develop on this version)
- MySQL
- familiar with Gin & Gorm
  
## Feature
- [Golang](https://go.dev) `v1.17.5`
- [Gin](https://github.com/gin-gonic/gin) `v1.7.7`
- [Gorm](https://gorm.io)
  
## How to use it
clone this repo with `https` / `ssh` / `github cli`
  
```sh
git clone https://github.com/fahmidyt/go-book-rental-be.git
```
  
After cloning this repo, make sure you have `duplicated` the `.env.example` file to `.env`, don't let the .env.example file be deleted or renamed.
  
## Install
Run `go install`
  
## Run Application
Makesure mysql database and .env configuration all are correct.
after that, run the application
```sh
go run main.go
```
app will start & gorm will migrate automaticaly into the selected database (note: only structre)
  
## Environment
you could change `ENV` into "PRODUCTION" when the app is about to run in released mode
and you also could change `PORT` to your desire.
