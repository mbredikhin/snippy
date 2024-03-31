# Snippy

A RESTful API example for snippets app with GO

## Installation and run

```sh
go get github.com/mbredikhin/snippy
cd snippy
# Build services
make build
# Create and start containers
make run
# Run database migrations
make migrate
# API Endpoint: http://127.0.0.1:8001
```

## API Reference

#### Register new user

```http
  POST /auth/sign-up
```

| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `username` | `string` | **Required**. Your username |
| `password` | `string` | **Required**. Your password |
| `name`     | `string` | Your full name              |

Response

```
{
    "id": number
}
```

#### Login with username and password

```http
  POST /auth/sign-in
```

| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `username` | `string` | **Required**. Your username |
| `password` | `string` | **Required**. Your password |

Response

```
{
    "token": string
}
```

Send given token in Authorization HTTP header – `Authorization: "Bearer %s" `

#### Create a new list of snippets

```http
  POST /api/lists
```

| Parameter | Type     | Description                   |
| :-------- | :------- | :---------------------------- |
| `name`    | `string` | **Required**. Collection name |

Response

```
{
    "data": {
        "id": number,
        "user_id": string,
        "name": string
    }
}
```

#### Get all lists

```http
  GET /api/lists
```

Response

```
{
    "data": [
        {
            "id": number,
            "user_id": string,
            "name": string
        }
    ] | null
}
```

#### Get list by id

```http
  GET /api/lists/{id}
```

| Parameter | Type     | Description                 |
| :-------- | :------- | :-------------------------- |
| `id`      | `number` | **Required**. Collection id |

Response

```
{
    "data": {
        "id": number,
        "user_id": string,
        "name": string
    }
}
```

#### Update list

```http
  PUT /api/lists/{id}
```

| Parameter | Type     | Description                   |
| :-------- | :------- | :---------------------------- |
| `name`    | `string` | **Required**. Collection name |

Response

```
{
    "data": {
        "id": number,
        "user_id": string,
        "name": string
    }
}
```

#### Delete list

```http
  DELETE /api/lists/{id}
```

| Parameter | Type     | Description                 |
| :-------- | :------- | :-------------------------- |
| `id`      | `number` | **Required**. Collection id |

Response

```
{
    "id": number
}
```

#### Add language

```http
  POST /api/languages
```

| Parameter | Type     | Description                 |
| :-------- | :------- | :-------------------------- |
| `name`    | `string` | **Required**. Language name |

Response

```
{
    "id": number
}
```

#### Get list of languages

```http
  GET /api/languages
```

Response

```
{
    "data": [
      {
          "id": number,
          "name": string
      } | null
    ]
}
```

#### Add new snippet

```http
  POST /api/lists/{list_id}/snippets
```

| Parameter     | Type     | Description                   |
| :------------ | :------- | :---------------------------- |
| `name`        | `string` | **Required**. Snippet name    |
| `language_id` | `number` | **Required**. Language ID     |
| `content`     | `string` | **Required**. Snippet content |

Response

```
{
    "id": number
}
```

#### Get all snippets of a list

```http
  GET /api/lists/{list_id}/snippets
```

Response

```
{
    "data": [
        {
            "id": number,
            "name": string,
            "user_id": number
        }
    ] | null
}
```

#### Get snippet

```http
  GET /api/snippets/{id}
```

Response

```
{
    "id": number,
    "name": string,
    "content": string,
    "language_id": number,
    "list_id": number
}
```

#### Update snippet

```http
  PUT /api/snippets/{id}
```

| Parameter     | Type     | Description                   |
| :------------ | :------- | :---------------------------- |
| `name`        | `string` | **Required**. Snippet name    |
| `language_id` | `number` | **Required**. Language ID     |
| `list_id`     | `number` | **Required**. Collection ID   |
| `content`     | `string` | **Required**. Snippet content |

Response

```
{
    "status": "ok"
}
```

#### Delete snippet

```http
  DELETE /api/snippets/{id}
```

Response

```
{
    "status": "ok"
}
```

#### /api/favourite-snippets

- `GET` : Get favourite snippet's IDs

#### /api/favourite-snippets/:id

- `POST` : Add snippet to favourites
- `DELETE` : Delete snippet from favourites

#### /api/tags

- `POST` : Create new tag
- `GET` : Get tags

#### /api/tags/:id

- `GET` : Get tag
- `PUT` : Update tag
- `DELETE` : Delete tag
