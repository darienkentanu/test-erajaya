# Test - Erajaya
Coding Test Erajaya

Framework used in this project: <br>
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)

# Table of Content
- [Description](#description)
- [How to use](#how-to-use)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
This is submition of coding test of erajaya
```
In this project i use MVC architecture because it's the most common achitecture used by programmer
```

# How to use
- Install Go and MySQL or (install docker and docker-compose)
- Clone this repository in your $PATH:
```
$ git clone https://github.com/darienkentanu/test-erajaya
```
- add your '.env' files containing following env variabel:
```
DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME,
```
- run this project
```
$ go run main.go or $ docker-compose up --build
```
# Endpoints

| Method | Endpoint | Description|
|:-----|:--------|:----------| 
| POST  | /products | add a product
|:-----|:--------|:----------| 
| GET  | /products{sortby} | get product sortby: newest, lowest, highest, asc, desc
|:-----|:--------|:----------| 


```
- example POST
http://localhost:8000/products
-JSON Input 
{
	"name": "pulsa 200000",
	"description": "pulsa 200 rb",
	"price": 200000,
	"quantity": 10
}
```

```
- example GET
http://localhost:8000/products?sortby=newest
```

<br>

## Credits

- [Darien Kentanu](https://github.com/darienkentanu) (Author and maintainer)
