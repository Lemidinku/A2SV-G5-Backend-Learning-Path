# Let's write the provided API documentation to a markdown file.

# Task Manager API Documentation

## Postman documentation link

### [Task manager api](https://documenter.getpostman.com/view/31604198/2sA3s6F9rw)

## Overview

This API allows you to manage tasks. You can create, read, update, and delete tasks using this API.

## Base URL

`http://localhost:8080`

## Endpoints

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
      "ID": "1",
      "Title": "Task 1",
      "Description": "Description for Task 1",
      "DueDate": "2024-08-07T16:26:09.0525595+03:00",
      "Status": "In Progress"
    },
    {
      "ID": "2",
      "Title": "Task 2",
      "Description": "Description for Task 2",
      "DueDate": "2024-08-06T16:26:09.0525595+03:00",
      "Status": "Completed"
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
    "ID": "1",
    "Title": "Task 1",
    "Description": "Description for Task 1",
    "DueDate": "2024-08-07T16:26:09.0525595+03:00",
    "Status": "In Progress"
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
  "Title": "task 3",
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
    "ID": "7a9cd31b-ea5c-4c92-b3df-dabedc1b13cf",
    "Title": "task 3",
    "Description": "Description for new task",
    "DueDate": "2024-08-06T10:31:27.6750959+03:00",
    "Status": "In progress"
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
