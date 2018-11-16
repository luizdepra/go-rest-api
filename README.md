# go-rest-api

An example of REST API made with Go. A simple Task API.

## Requirements

- Go 1.11+
- `sqlite3` commandline tool

## Usage

1. Copy the `example.env` to `.env` and change its content as you want.
2. Run `make db` to create a sqlite3 database with an initial data. (Optional, if you already have a database file.)
3. Build and run the project with `make run`.
4. Enjoy the API.

## Endpoints

Available endpoints and usage example.

### List Tasks

Returns all stored tasks.

#### URL

`GET /tasks/`

#### Response Status


| Code | Description           |
| ---- | --------------------- |
| 200  | Ok                    |
| 500  | Internal Server Error |

#### Response Data

``` json
[
    {
        "id": 1,
        "title": "Task 01",
        "priority": 0,
        "done": true
    },
    {
        "id": 2,
        "title": "Task 02",
        "priority": 1,
        "done": false
    },
    {
        "id": 3,
        "title": "Task 03",
        "priority": 99,
        "done": false
    }
]
```

### Create a Task

Stores a new task.

#### URL

`POST /tasks/`

#### Payload

``` json
{
    "title": "A Task",
    "priority": 5,
    "done": false
}
```

#### Response Status

| Code | Description           |
| ---- | --------------------- |
| 201  | Created               |
| 400  | Bad Request           |
| 500  | Internal Server Error |

#### Response Data

``` json
{
    "id": 2,
    "title": "A Task",
    "priority": 5,
    "done": false
}
```

### Get a Task

Returns a tasks by its id.

#### URL

`GET /tasks/:id/`

#### Response Status

| Code | Description           |
| ---- | --------------------- |
| 200  | Ok                    |
| 400  | Bad Request           |
| 404  | Not Found             |
| 500  | Internal Server Error |

#### Response Data

``` json
{
    "id": 1,
    "title": "Task 01",
    "priority": 0,
    "done": true
}
```


### Update a Task

Updates a stored task with the provided data by its id.

#### URL

`PUT /tasks/:id/`

#### Payload

``` json
{
    "title": "Updated Task",
    "priority": 7,
    "done": false
}
```

#### Response Status

| Code | Description           |
| ---- | --------------------- |
| 200  | Ok                    |
| 400  | Bad Request           |
| 404  | Not Found             |
| 500  | Internal Server Error |

#### Response Data

``` json
{
    "id": 3,
    "title": "Updated Task",
    "priority": 7,
    "done": false
}
```

### Delete a Task

Removes a stored task by its id.

#### URL

`DELETE /tasks/:id/`

#### Response Status

| Code | Description           |
| ---- | --------------------- |
| 200  | Ok                    |
| 400  | Bad Request           |
| 404  | Not Found             |
| 500  | Internal Server Error |

#### Response Data

``` json
{
    "id": 4,
    "title": "Deleted Task",
    "priority": 11,
    "done": true
}
```
