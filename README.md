# IMDB Top 250 Scrapper

---

### Table of Contents

- [Description](#description)
- [How To Use](#how-to-use)
- [Author Info](#author-info)

---

## Description

This project scraps the top 250 movies [from](https://www.imdb.com/chart/top/), inserts the data to PostgreSQL and handles a single or all movies HTTP requests.

## Technologies

### Main Technologies

- [Go](https://go.dev/)
- [Fiber Framework](https://docs.gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)

### Libraries

- [gocolly/colly](https://github.com/gocolly/colly)
- [jackc/pgconn](https://github.com/jackc/pgconn)
- [jackc/pgx](https://github.com/jackc/pgx/v4)

[Back To The Top](#IMDB-Top-250-Scrapper)

---

## How To Use

### Tools

- [Go](https://go.dev/dl/)
- [Soda CLI](https://gobuffalo.io/documentation/database/soda/)
- [DataGrip](https://www.jetbrains.com/datagrip/)

### Setup Database

- Create Database

```
CREATE DATABASE <your database name>
```

- Create your dsn.go file in dsn folder

```
var DSN = "host=localhost port=5432 dbname=<your db name> user=<your username> password=<your password>"
```

- Create your own database.yml file and run in terminal

```
soda migrate
```

- For dropping migration run

```
soda migrate down
```

### Start App

- Start app directly

```
go run cmd/main/*.go
```

### Give it a try

#### Routes

| Request     | URL           |
| ----------- | ------------- |
| Get Movie   | :3000/api/:id |
| List Movies | :3000/api/all |

[Back To The Top](#IMDB-Top-250-Scrapper)

## Author Info

- Twitter - [@dev_bck](https://twitter.com/dev_bck)

[Back To The Top](#IMDB-Top-250-Scrapper)
