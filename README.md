# Test - Erajaya
test erajaya

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)

# Table of Content
- [Description](#description)
- [How to use](#how-to-use)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
This is submition of coding test of erajaya

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
$ go run main.go or $ docker-compose up --build -d

```


# Endpoints

| Method | Endpoint | Description|
|:-----|:--------|:----------| 
| POST  | /products | add a product
|:-----|:--------|:----------| 
| GET  | /newestproducts | get product from newest
| GET | /productslowest | get product from lowest price
| GET | /productshighest | get product from highest price
| GET | /productsasc | get products order by A-Z
| GET | /productsdesc | get products order by Z-A
|---|---|---|---|---|

<br>

## Credits

- [Darien Kentanu](https://github.com/darienkentanu) (Author and maintainer)
