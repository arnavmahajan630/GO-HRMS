# 🧾 Go HRMS API

A fast and lightweight Human Resource Management System (HRMS) API built using [Go Fiber](https://gofiber.io/) and [MongoDB](https://www.mongodb.com/). This project provides basic CRUD operations for managing employee records.

---

## 🚀 Features

 A simple Golang Practice Project part of my golang learning series https://github.com/arnavmahajan630/learn-go

- ✅ RESTful API with Fiber (Express-style Go framework)
- 🧩 MongoDB integration via the official Go driver
- 📄 CRUD operations on employees
- ⚡ Fast and minimal API design
- 🧪 Ready for expansion into full HRMS (authentication, payroll, etc.)

---

## 📁 Project Structure

```
/hrms-api
├── main.go                # App entry point
├── DB/
│   └── db.go 
    └── db.models.go      # MongoDB connection logic
├── models/
│   └── models.go         # Employee struct and schema
├── handlers/
│   └── models.go       # Route handlers for /employee
├── helpers/             # Helpers for handlers
├── go.mod / go.sum       # Go modules and dependencies
```

---

## 🛠️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-username/hrms-api.git
cd hrms-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run MongoDB

Make sure MongoDB is running locally on default port (`mongodb://localhost:27017`).

### 4. Run the application

```bash
go run main.go
```

Server will start at: `http://localhost:3000`

---

## 📚 API Endpoints

| Method | Endpoint         | Description              |
|--------|------------------|--------------------------|
| GET    | `/employee`      | Get all employees        |
| POST   | `/employee`      | Add a new employee       |
| PUT    | `/employee/:id`  | Update an employee       |
| DELETE | `/employee/:id`  | Delete an employee       |

> All endpoints use `application/json` format for request/response.

---

## 📌 Sample Request

### POST `/employee`

```json
{
  "name": "Alice Smith",
  "salary": 60000,
  "age": 28
}
```

---

## 🔮 Future Improvements

- JWT Authentication & RBAC
- Pagination and filtering
- Department/project association
- Unit testing with `testify`
- Docker support
- Swagger/OpenAPI documentation

---

## 🧑‍💻 Tech Stack

- **Language**: Go
- **Framework**: [Fiber](https://gofiber.io/)
- **Database**: MongoDB
- **Driver**: Official MongoDB Go driver

---

## 📝 License

MIT License. See `LICENSE` file for details.

---

## 🤝 Contributing

Feel free to fork, raise issues, and submit PRs.

---

## 📬 Contact

Built with ❤️ by [Arnav Mahajan](https://github.com/arnavmahajan630)
