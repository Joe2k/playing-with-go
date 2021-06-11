
# playing-with-go

A very simple Notification Service in Go 💬

## API Reference

#### Get all notifications

```http
  GET /api/notification
```

#### Create notification

```http
  POST /api/notification
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | Optional |
| `content`      | `string` | Content of the notification |
| `receiver.name`      | `string` | Receiver's name |
| `receiver.number`      | `string` | Receiver's number |
