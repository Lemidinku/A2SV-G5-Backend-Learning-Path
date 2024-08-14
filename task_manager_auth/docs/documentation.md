# Task Management API Documentation

## Overview

The Task Management API allows you to manage tasks with basic CRUD (Create, Read, Update, Delete) operations. This API is built using Go and the Gin framework and uses MongoDB for persistent data storage.

### Postman documentation link
[Task Manager API](https://documenter.getpostman.com/view/31604198/2sA3s6FAAc)

## Setting up environmental variables

To set up the JWT secret key for authentication, follow these steps:

1. Create a new file in the root directory of your project called `.env`.

2. Open the `.env` file and add the following line:

```
JWT_SECRET="your_secret_key"
```

Replace `"your_secret_key"` with your desired secret key for JWT authentication.

3. Save the `.env` file.

## Running the API

To run the API, you need to have Go and MongoDB installed on your system. You can start the API by running the following command in the root directory of the project:

```bash
go run main.go
```

This will start the API server on `http://localhost:8080`.

# Task Manager API Documentation

## Overview

This API allows you to manage tasks. You can create, read, update, and delete tasks using this API.

## Base URL

`http://localhost:8080`

## Endpoints

### Register a user

**Endpoint:** `POST /register`

**Description:** creates a new user.

**Request Body:**

```json
{
  "username": "admin",
  "password": "admin"
}
```

**Response:**

- **Status Code:** `200 OK`
- **Body:**

```json
{
  "user": {
    "id": "cfa6d4a9-126b-4923-93a3-f729629630b2",
    "username": "user",
    "password": "$2a$10$GnCRRbyV/vTn8UiRtECeJeYVUwcJ7dzJsPgEgnEsEo8Y1PKAEUA3S",
    "role": "user"
  }
}
```

- **Status Code:** `400 Bad Request`
- **Body:**

```json
{
  "error": "username already taken"
}
```

### log in

**Endpoint:** `POST /login`

**Description:** logs in the user.

**Request Body:**

```json
{
  "username": "admin",
  "password": "admin"
}
```

**Response:**

- **Status Code:** `200 OK`
- **Body:**

```json
{
  "message": "User logged in successfully",
  "token": "exampleToken"
}
```

- **Status Code:** `404 Not Found`
- **Description:** `if the username is wrong`
- **Body:**

```json
{
  "message": "user not found"
}
```

- **Status Code:** `404 Not Found`
- **Description:** `if the password is wrong`
- **Body:**

```json
{
  "message": "invalid password"
}
```

### Promote a user to admin

**Endpoint:** `PUT /promote/:username`

**Description:** Promote a user to admin.

**Response:**

- **Status Code:** `200 OK`
- **Body:**

```json
{
  "message": "The user have been promoted"
}
```

### Get All Tasks

**Endpoint:** `GET /tasks`

**Description:** Retrieves all tasks.

**Response:**

- **Status Code:** `200 OK`
- **Body:**

```json
{
  "tasks": [
    {
      "id": "c71edb66-bc3b-449b-9386-c180860ea9d3",
      "title": "task 1",
      "description": "Description for new task",
      "dueDate": "2024-08-06T07:31:27.675Z",
      "status": "In progress"
    },
    {
      "id": "e18514d2-8d7e-463c-a5ca-75c68450c389",
      "title": "task 2",
      "description": "Description for new task",
      "dueDate": "2024-08-06T07:31:27.675Z",
      "status": "In progress"
    },
    {
      "id": "a52047a3-d1af-41d2-af59-7debd44da925",
      "title": "task 3",
      "description": "Description for new task",
      "dueDate": "2024-08-06T07:31:27.675Z",
      "status": "In progress"
    }
  ]
}
```

### Get a Task by ID

**Endpoint:** `GET /tasks/:id`

**Description:** Retrieves a task by its ID.

**Parameters:**

Path Parameter:

- `id` (string) - The ID of the task.

**Response:**

- **Status Code:** `200 OK`
- **Body:**

```json
{
  "task": {
    "id": "c71edb66-bc3b-449b-9386-c180860ea9d3",
    "title": "task 1",
    "description": "Description for new task",
    "dueDate": "2024-08-06T07:31:27.675Z",
    "status": "In progress"
  }
}
```

- **Status Code:** `404 Not Found`
- **Body:**

```json
{
  "message": "Task not found"
}
```

### Add a New Task

**Endpoint:** `POST /tasks`

**Description:** Creates a new task.

**Request Body:**

```json
{
  "Title": "new task",
  "Description": "Description for new task",
  "DueDate": "2024-08-06T10:31:27.6750959+03:00",
  "Status": "In progress"
}
```

**Response:**

- **Status Code:** `201 Created`
- **Body:**

```json
{
  "task": {
    "id": "5fb2e97e-a55a-4546-9eb7-03ba83baad60",
    "title": "new task",
    "description": "Description for new task",
    "dueDate": "2024-08-06T07:31:27.675Z",
    "status": "In progress"
  }
}
```

### Remove a Task by ID

**Endpoint:** `DELETE /tasks/:id`

**Description:** Deletes a task by its ID.

**Parameters:**

Path Parameter:

- `id` (string) - The ID of the task.

**Response:**

- **Status Code:** `200 OK`
- **Body:**

```json
{
  "message": "Task removed"
}
```

- **Status Code:** `404 Not Found`
- **Body:**

```json
{
  "message": "Task not found"
}
```

### Update a Task by ID

**Endpoint:** `PUT /tasks/:id`

**Description:** Updates a task by its ID.

**Parameters:**
Path Parameter:

- `id` (string) - The ID of the task.
  **Request Body:**

```json
{
  "Title": "updated title",
  "Status": "completed"
}
```

**Response:**

- **Status Code:** `200 OK`
- **Body:**

```json
{
  "task": {
    "ID": "1",
    "Title": "updated title",
    "Description": "Description for Task 1",
    "DueDate": "2024-08-07T16:26:09.0525595+03:00",
    "Status": "completed"
  }
}
```

- **Status Code:** `404 Not Found`
- **Body:**

```json
{
  "message": "Task not found"
}
```

Usage Example
Get All Tasks

```sh
curl -X GET http://localhost:8080/tasks
```

Get a Task by ID

```sh
curl -X GET http://localhost:8080/tasks/1
```

Add a New Task

```sh
curl -X POST http://localhost:8080/tasks -d '{"title": "New Task", "description": "Description for new task", "DueDate": "2024-08-07T16:26:09.0525595+03:00","status": "pending"}' -H "Content-Type: application/json"
```

Remove a Task by ID

```sh
curl -X DELETE http://localhost:8080/tasks/1
```

Update a Task by ID

```sh
curl -X PUT http://localhost:8080/tasks/1 -d '{"title": "Updated Task", "description": "Description for new task", "status": "pending"}' -H "Content-Type: application/json"
```
