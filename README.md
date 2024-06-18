To build and run the application:

go run main.go

## To Create User:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"username": "testuser2", "email": "test1234@example.com", "password": "password123"}' http://localhost:8080/users
```

```JSON
  {
    "id": 2,
    "username": "testuser2",
    "email": "test1234@example.com",
    "password": "password123"
  }
```

