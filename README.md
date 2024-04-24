# Snippy

Backend of the Snippy service - lightning-fast solution for managing code snippets.

## Installation and run

```sh
git clone git@github.com:mbredikhin/snippy.git && cd snippy
# Build services
make build
# Create and start containers
make run
# Run database migrations
make migrate-up
# Seed database
make db-seed
# API host name: http://localhost:8001
# curl -X POST --data '{"username":"username", "password":"password"}' http://localhost:8001/auth/sign-in
# curl -X POST --data '{"name":"My snippets collection"}' --header 'Authorization: Bearer {{token}}' http://localhost:8001/api/lists
```

## API Reference

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://god.gw.postman.com/run-collection/12758470-78ec42b6-b8e1-4935-ba76-d7e701cf4f37?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D12758470-78ec42b6-b8e1-4935-ba76-d7e701cf4f37%26entityType%3Dcollection%26workspaceId%3D26ed6946-1a78-42a3-a05e-6b50744b4675)

#### Register new user

```http
  POST /auth/sign-up
```

Body:

| Key        | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `username` | `string` | **Required**. Your username |
| `password` | `string` | **Required**. Your password |
| `name`     | `string` | Your full name              |

Response

```
{
    "status": "ok"
}
```

#### Login with username and password

```http
  POST /auth/sign-in
```

Body:

| Key        | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `username` | `string` | **Required**. Your username |
| `password` | `string` | **Required**. Your password |

Response

```
{
    "data": {
        "token": string
    }
}
```

Send given token in Authorization HTTP header – `Authorization: "Bearer %s" `

#### Create a new list of snippets

```http
  POST /api/lists
```

Body:

| Key    | Type     | Description                   |
| :----- | :------- | :---------------------------- |
| `name` | `string` | **Required**. Collection name |

Response

```
{
    "data": {
        "id": number,
        "name": string
    }
}
```

#### Get all lists

```http
  GET /api/lists
```

Query parameters:

| Parameter | Type     | Description      |
| :-------- | :------- | :--------------- |
| `page`    | `number` | Page             |
| `limit`   | `number` | Pagination limit |

Response

```
{
    "data": [
        {
            "id": number,
            "name": string
        }
    ]
}
```

#### Get list by id

```http
  GET /api/lists/:id
```

Query parameters:

| Parameter | Type     | Description                 |
| :-------- | :------- | :-------------------------- |
| `id`      | `number` | **Required**. Collection ID |

Response

```
{
    "data": {
        "id": number,
        "name": string
    }
}
```

#### Update list

```http
  PUT /api/lists/:id
```

Body:

| Key    | Type     | Description                   |
| :----- | :------- | :---------------------------- |
| `name` | `string` | **Required**. Collection name |

Response

```
{
    "data": {
        "id": number,
        "name": string
    }
}
```

#### Delete list

```http
  DELETE /api/lists/:id
```

Body:

| Key  | Type     | Description                 |
| :--- | :------- | :-------------------------- |
| `id` | `number` | **Required**. Collection ID |

Response

```
{
    "data": {
        "id": number
    }
}
```

#### Add language

```http
  POST /api/languages
```

Body:

| Key    | Type     | Description                 |
| :----- | :------- | :-------------------------- |
| `name` | `string` | **Required**. Language name |

Response

```
{
    "data": {
        "id": number
    }
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
        }
    ]
}
```

#### Create new snippet

```http
  POST /api/lists/:id/snippets
```

Body:

| Key           | Type     | Description                    |
| :------------ | :------- | :----------------------------- |
| `name`        | `string` | **Required**. Snippet name     |
| `language_id` | `number` | **Required**. Language ID      |
| `description` | `string` | **Required**. Text description |
| `content`     | `string` | **Required**. Snippet content  |

Response

```
{
    "data": {
        "id": number
    }
}
```

#### Get all snippets of the list

```http
  GET /api/lists/:id/snippets
```

Query parameters:

| Parameter | Type       | Description                                        |
| :-------- | :--------- | :------------------------------------------------- |
| `tag_ids` | `number[]` | Filters snippets with any of entered tags assigned |
| `page`    | `number`   | Page                                               |
| `limit`   | `number`   | Pagination limit                                   |

Response

```
{
    "data": [
        {
            "id": number,
            "list_id": number,
            "name": string,
            "language_id": number,
            "description": string,
            "content": string
        }
    ]
}
```

#### Get snippet

```http
  GET /api/snippets/:id
```

Response

```
{
    "data": {
        "id": number,
        "list_id": number,
        "name": string,
        "language_id": number,
        "description": number,
        "content": string
    }
}
```

#### Update snippet

```http
  PUT /api/snippets/:id
```

Body:

| Key           | Type     | Description                    |
| :------------ | :------- | :----------------------------- |
| `name`        | `string` | **Required**. Snippet name     |
| `language_id` | `number` | **Required**. Language ID      |
| `list_id`     | `number` | **Required**. Collection ID    |
| `description` | `string` | **Required**. Text description |
| `content`     | `string` | **Required**. Snippet content  |

Response

```
{
    "status": "ok"
}
```

#### Delete snippet

```http
  DELETE /api/snippets/:id
```

Response

```
{
    "status": "ok"
}
```

#### Add snippet to favourites

```http
  POST /api/snippets/favourites
```

Body:

| Key  | Type     | Description              |
| :--- | :------- | :----------------------- |
| `id` | `number` | **Required**. Snippet ID |

Response

```
{
    "status": "ok"
}
```

#### Remove snippet from favourites

```http
  DELETE /api/snippets/favourites
```

Body:

| Key  | Type     | Description              |
| :--- | :------- | :----------------------- |
| `id` | `number` | **Required**. Snippet ID |

Response

```
{
    "status": "ok"
}
```

#### Get ids of favourite snippets

```http
  GET /api/snippets/favourites
```

Response

```
{
    "data": [number]
}
```

#### Add tag

```http
  POST /api/tags
```

Body:

| Key    | Type     | Description          |
| :----- | :------- | :------------------- |
| `name` | `string` | **Required**. Tag ID |

Response

```
{
    "data": {
        "id": number
    }
}
```

#### Update tag

```http
  PUT /api/tags/:id
```

Body:

| Key  | Type     | Description          |
| :--- | :------- | :------------------- |
| `id` | `number` | **Required**. Tag ID |

Response

```
{
    "status": "ok"
}
```

#### Delete tag

```http
  DELETE /api/tags/:id
```

Response

```
{
    "status": "ok"
}
```

#### Get tag

```http
  GET /api/tags/:id
```

Response

```
{
    "data": {
        "id": number,
        "name": string
    }
}
```

#### Get list of tags

```http
  GET /api/tags
```

Response

```
{
    "data": [
        {
            "id": number,
            "name": string
        }
    ]
}
```

#### Get list of tags assigned to the snippet

```http
  GET /api/snippets/:id/tags
```

Response

```
{
    "data": [number]
}
```

#### Assign tag to the snippet

```http
  POST /api/snippets/:id/tags
```

Body:

| Key      | Type     | Description          |
| :------- | :------- | :------------------- |
| `tag_id` | `number` | **Required**. Tag ID |

Response

```
{
    "status": "ok"
}
```

#### Unassign tag from the snippet

```http
  DELETE /api/snippets/:id/tags
```

Body:

| Key      | Type     | Description          |
| :------- | :------- | :------------------- |
| `tag_id` | `number` | **Required**. Tag ID |

Response

```
{
    "status": "ok"
}
```
