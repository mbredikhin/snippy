# Snippy

Backend of the Snippy service - lightning-fast solution for managing code snippets.

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
| `id`      | `number` | **Required**. Collection ID |

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
| `id`      | `number` | **Required**. Collection ID |

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

#### Add snippet to favourites

```http
  POST /api/favourite-snippets
```

| Parameter | Type     | Description              |
| :-------- | :------- | :----------------------- |
| `id`      | `number` | **Required**. Snippet ID |

Response

```
{
    "status": "ok"
}
```

#### Get ids of favourite snippets

```http
  GET /api/favourite-snippets
```

Response

```
{
    "data": [number] | null
}
```

#### Delete snippet from favourites

```http
  DELETE /api/favourite-snippets/{id}
```

Response

```
{
    "status": "ok"
}
```

#### Add tag

```http
  POST /api/tags
```

| Parameter | Type     | Description          |
| :-------- | :------- | :------------------- |
| `name`    | `string` | **Required**. Tag ID |

Response

```
{
    "id": number
}
```

#### Get list of tags

```http
  GET /api/tags
```

Response

```
[
    {
        "id": number,
        "name": string,
        "user_id": number
    } | null
]
```

#### Get tag

```http
  GET /api/tags/{id}
```

Response

```
{
    "id": number,
    "name": string,
    "user_id": number
}
```

#### Update tag

```http
  PUT /api/tags/{id}
```

| Parameter | Type     | Description          |
| :-------- | :------- | :------------------- |
| `id`      | `number` | **Required**. Tag ID |

Response

```
{
    "status": "ok"
}
```

#### Delete tag

```http
  DELETE /api/tags/{id}
```

Response

```
{
    "status": "ok"
}
```
