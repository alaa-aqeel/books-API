# books API

Created to practice Go programming as starting point.
<p align="center">
<img min-height="350" src="https://firebasestorage.googleapis.com/v0/b/shop-f-ba8a0.appspot.com/o/Component%201.png?alt=media&token=47624ca8-dab0-40dc-9ace-2296f6b3cda8" alt="Let's learn golang"/>
</p>

## Data source

original data : https://github.com/benoitvallon/100-best-books/blob/master/books.json

* currently used data has some modifications

## Data

A JSON file with array of book objects as sample shown below.

```json
{
    "books": [
        {
          "author": "Chinua Achebe",
          "language": "English",
          "link": "https://en.wikipedia.org/wiki/Things_Fall_Apart",
          "pages": 209,
          "title": "Things Fall Apart"
        }
    ]
}
```



## API Endpoints

### All books

```http
GET http://localhost:8080/all
```

### Sample response
Status: 200 OK
```json
{
    "books": [
        {
            "id": 1,
            "author": "Chinua Achebe",
            "language": "English",
            "link": "https://en.wikipedia.org/wiki/Things_Fall_Apart",
            "pages": 209,
            "title": "Things Fall Apart"
        }
    ],
    "message": ""
}
```



### Find by title

```http
GET http://localhost:8080/title={title}
```

##### Request String Parameters

| Request String Parameter | Description         | Type   |
| ------------------------ | ------------------- | ------ |
| title                    | part of  book title | String |


### Sample request

```http
GET http://localhost:8080/title=ALL
```

### Sample response
Status: 200 OK
```json
{
    "books": [
        {
            "id": 1,
            "author": "Chinua Achebe",
            "language": "English",
            "link": "https://en.wikipedia.org/wiki/Things_Fall_Apart",
            "pages": 209,
            "title": "Things Fall Apart"
        }
    ],
    "message": ""
}
```



### Find by author

```http
GET http://localhost:8080/author={author}
```

##### Request String Parameters

| Request String Parameter | Description                | Type   |
| ------------------------ | -------------------------- | ------ |
| author                   | part of book's author name | String |


### Sample request

```http
GET http://localhost:8080/author=po
```

### Sample response
Status: 200 OK
```json
{
    "books": [
        {
            "id": 73,
            "author": "Edgar Allan Poe",
            "language": "English",
            "link": "https://en.wikipedia.org/wiki/Edgar_Allan_Poe_bibliography#Tales",
            "pages": 842,
            "title": "Tales"
        }
    ],
    "message": ""
}
```



### Find by language

```http
GET http://localhost:8080/lang={language}
```

##### Request String Parameters

| Request String Parameter  | Description      | Type   |
| ------------------------- | ---------------- | ------ |
| lang                      | part of book language | String |


### Sample request

```http
GET http://localhost:8080/lang=norSe
```

### Sample response
Status: 200 OK
```json
{
    "books": [
        {
            "id": 7,
            "author": "Unknown",
            "language": "Old Norse",
            "link": "https://en.wikipedia.org/wiki/Nj%C3%A1ls_saga",
            "pages": 384,
            "title": "Njál's Saga"
        }
    ],
   "message": ""
}
```



### Find by id

```http
GET http://localhost:8080/id={id}
```

##### Request String Parameters

| Request String Parameter | Description | Type    |
| ------------------------- | ----------- | ------- |
| id                        | id of book  | integer |

* id parameter value should be starting from 1

### Sample request

```http
GET http://localhost:8080/id=33
```

### Sample response
Status: 200 OK
```json
{
    "book": {
            "id": 33,
            "author": "Gustave Flaubert",
            "language": "French",
            "link": "https://en.wikipedia.org/wiki/Madame_Bovary",
            "pages": 528,
            "title": "Madame Bovary"
        },
    "message": ""
}
```



### Find by page range

```http
GET http://localhost:8080/page/min={min}&max={max}
```

##### Request String Parameters

| Request String Parameter | Description                         | Type    |
| :----------------------- | ----------------------------------- | ------- |
| min                      | minimum number of pages in the book | integer |
| max                      | maximum number of pages in the book | integer |

* both parameter value should be starting from 1

* minimum pages should not exceed maximum pages

### Sample request

```http
GET http://localhost:8080/page/min=7&max=70
```

### Sample response
Status: 200 OK
```json
{
    "books": [
        {
            "id": 46,
            "author": "Henrik Ibsen",
            "language": "Norwegian",
            "link": "https://en.wikipedia.org/wiki/A_Doll%27s_House",
            "pages": 68,
            "title": "A Doll's House"
        }
    ],
    "message": ""
}
```



### Add new book

```http
POST http://localhost:8080/book
```

### Sample request

```http
POST http://localhost:8080/book
```
##### Request body
```json
{
    "author": "Marguerite Yourcenar",
    "language": "French",
    "link": "https://en.wikipedia.org/wiki/Memoirs_of_Hadrian",
    "pages": 195,
    "title": "Memoirs of Hadrian"
}
```

### Sample response
Status: 201 Created
```json
{
    "book": {
        "id": 101,
        "author": "Marguerite Yourcenar",
        "language": "French",
        "link": "https://en.wikipedia.org/wiki/Memoirs_of_Hadrian",
        "pages": 195,
        "title": "Memoirs of Hadrian"
    },
    "message": ""
}
```



### Update a book

```http
PATCH http://localhost:8080/id={id}
```

##### Request String Parameters

| Request String Parameter | Description | Type    |
| ------------------------ | ----------- | ------- |
| id                       | id of book  | integer |

* book of specified id should exist

### Sample request

```http
PATCH http://localhost:8080/id=103
```
##### Request body
```json
{
    "language": "English",
    "link": "https://en.wikipedia.org/wiki/Things_Fall_Apart",
    "pages": 99,
    "title": "Things fall apart"
}
```

### Sample response
Status: 200 OK
```json
{
    "book": {
        "id": 103,
        "author": "Marguerite Yourcenar",
        "language": "English",
        "link": "https://en.wikipedia.org/wiki/Things_Fall_Apart",
        "pages": 99,
        "title": "Things fall apart"
    },
    "message": ""
}
```
* only passed fields will be updated.
* only non-empty field values after trimming will be updated
* pages must be starting from 1

### Delete a book

```http
DELETE http://localhost:8080/id={id}
```

##### Request String Parameters

| Request String Parameter | Description | Type    |
| ------------------------ | ----------- | ------- |
| id                       | id of book  | integer |

* book of specified id should exist

### Sample request

```http
DELETE http://localhost:8080/id=103
```

### Sample response
Status: 200 OK
```json
{
    "message": "Book was successfully deleted"
}
```



## Status code

200 : request e.g. get books, update book and delete book successfully processed

201 : new book created

400 : bad input 

- non-integer entered as book id
- book id is less than one
- book with specified id does not exist
- page count is non-integer value
- page count is less than one
- minimum page count exceed maximum page count
- one of required fileds (all except link of book) is not specified when creating
- invalid json sent as request body
- request body size exceed 2 KB


500 : error occurred while processing

* error details will be returned



## How to run

```go
go run main.go
```

