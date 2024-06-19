## To clone the repository:

git clone https://github.com/DDanielcoding/blog-project.git
cd blog-project

## To install dependencies:

go mod tidy


## To build and run the application:

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
## Logging In:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"email": "test1234@example.com", "password": "password123"}' http://localhost:8080/login
```
```JSON
{
  "email": "test1234@example.com",
  "password": "password123"
}
```
## To Create a New blog_entry:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"username": "testuser2", "title": "Blog post by testuser2", "content": "This is my first blog post", "author_id": 2}' http://localhost:8080/blog_entries
```
```JSON
{
    "id": 2,
    "username": "testuser2",
    "title": "Blog post by testuser2",
    "content": "This is my first blog post",
    "author_id": 2,
    "created_at": "2024-06-17T22:17:40.830339957+02:00",
    "updated_at": "2024-06-17T22:17:40.830339957+02:00"
  }
```
## To Create a new comment on a specific blog entry id:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"content": "Test comment", "username": "testuser2", "author_id": 2, "blog_id":2 }' http://localhost:8080/blog_entries/2/comments
```
```JSON
{
  "id": 2,
  "username": "testuser2",
  "title": "Blog post by testuser2",
  "content": "This is my first blog post",
  "author_id": 2,
  "created_at": "2024-06-17T22:17:40.830339957+02:00",
  "updated_at": "2024-06-17T22:17:40.830339957+02:00"
}
```

