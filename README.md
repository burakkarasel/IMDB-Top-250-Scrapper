# IMDB Top 250 Scrapper Project

Scraps the IMDB Top 250 movies by there individual pages.
Insert them into PostgreSQL.
Used Fiber framework to create API from the data.

- Built in Go version 1.18
- Uses [Colly](https://github.com/gocolly/colly)
- Uses [jackc/pgx](https://github.com/jackc/pgx)
- Uses [buffalo/soda](https://github.com/gobuffalo/buffalo)
- Uses [fiber](https://github.com/gofiber/fiber)

## GIF of the Database
![imdb-scrapper](https://user-images.githubusercontent.com/99825584/181496788-865b1aa5-b7fe-41f1-97e4-e7fe285c4b16.gif)

Link of the [API](139-144-68-26.ip.linodeusercontent.com/all) <br/>
Also you can specify for only one movie such as (139-144-68-26.ip.linodeusercontent.com/api/24)
