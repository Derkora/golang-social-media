# golang-social-media

## Setup Docker
```sh
sudo docker run --name go-sm \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=socialmedia_db \
  -p 5432:5432 \
  -d postgres:latest
```

## Dokumentasi Testing
### USERS
- `POST /users`
```sh
{
  "username": "test",
  "email": "test@mail.com",
  "bio": "Dino enthusiast"
}
```

- `GET /users`

- `GET /users/{id}`

- `PUT /users/{id}`
```sh
{
  "username": "test",
  "email": "test@test.com",
  "bio": "Updated bio"
}
```

- `DELETE /users/{id}`

### POSTS
- `POST /posts`
```sh
{
  "user_id": "{id}",
  "content": "Hello World"
}
```

- `GET /posts`

- `GET /posts/{id}`

- `GET /users/{id}/posts`

- `DELETE /posts/{id}`

### LIKES
- `POST /likes`
```sh
{
  "user_id": "{id}",
  "post_id": "{id}"
}
```

- `GET likes`


- `GET /posts/{id}/likes`

- `GET /users/{id}/likes`

- `DELETE /likes/{id}`

### COMMENTS
- `POST /comments`
```sh
{
  "user_id": "{id}",
  "post_id": "{id}",
  "content": "Hallo test"
}
```

- `GET /comments`

- `GET /posts/{id}/comments`

- `DELETE /comments/{id}`
