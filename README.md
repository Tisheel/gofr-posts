# Gofr Posts API

A simple RESTful API built using [Gofr](https://gofr.dev), demonstrating CRUD operations on **Posts** and **Comments**.

## 🚀 Features

* Health check endpoint (`/ping`)
* CRUD operations for **Posts**

  * `GET /posts` → Get all posts
  * `GET /posts/{id}` → Get a post by ID
  * `POST /posts` → Create a new post
  * `PUT /posts/{id}` → Update a post
  * `PATCH /posts/{id}` → Partially update a post
  * `DELETE /posts/{id}` → Delete a post
* CRUD operations for **Comments**

  * `POST /comments` → Create a comment
  * `GET /posts/{id}/comments` → Get comments for a specific post
  * `GET /comments` → Get all comments

## 🛠 Tech Stack

* [Gofr](https://gofr.dev) – Go framework for building web apps
* Go (Golang)
* MySql

## 📂 Project Structure

```
.
├── main.go              # Application entrypoint
└── internal/
    └── handlers/        # Request handlers for posts & comments
```

## ▶️ Running the Project

### 1. Clone the repository

```bash
git clone https://github.com/your-username/gofr-posts.git
cd gofr-posts
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the application

```bash
go run main.go
```

By default, the server runs on **[http://localhost:8000](http://localhost:8000)** (unless configured otherwise).

### 4. Test the API

Check server health:

```bash
curl http://localhost:8000/ping
# -> pong
```

### Example: Create a Post

```bash
curl -X POST http://localhost:8000/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"John Snow","body":"Game of Thrones key character"}'
```

---

## 📖 API Documentation

Interactive API documentation is available at:

👉 [http://localhost:8000/.well-known/swagger](http://localhost:8000/.well-known/swagger)

This Swagger UI provides a visual interface to explore and test all available endpoints directly from your browser.

