# playing-with-go

A very simple Notification Service in Go with Postgres Database 💬

## How to use?

-   Clone this repository

```bash
  git clone https://github.com/Joe2k/playing-with-go
  cd playing-with-go
```

-   Create a .env file in the root directory and place all important DB info like below

```
  APP_DB_USERNAME=postgres
  APP_DB_PASSWORD="12345"
  APP_DB_NAME=postgres
```

-   Start a Postgres server with Docker with the make command below

```bash
  make postgres
```

-   Run the Go Server

```bash
  go run main.go
```

## Run Tests

To run unit tests, run the command below

```bash
  go test -v
```

## API Reference

#### Get all notifications

```http
  GET /notifications
```

#### Create a notification

```http
  POST /notification
```

#### Get a notification

```http
  GET /notification/:id
```

#### Update a notification

```http
  PUT /notification/:id
```

| Parameter | Type     | Description             |
| :-------- | :------- | :---------------------- |
| `id`      | `int`    | Optional                |
| `number`  | `int`    | Contact Number          |
| `message` | `string` | Message to be delivered |
