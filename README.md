# Todo List Application

A modular, command-line interface (CLI) and REST API-based Todo List application built in Go with MySQL as the database. The application allows users to manage tasks efficiently through a CLI or programmatically via a REST API. It features a clean, scalable architecture with separated concerns, making it easy to maintain and extend.

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Prerequisites](#prerequisites)
- [Setup Instructions](#setup-instructions)
- [Running the Application](#running-the-application)
- [CLI Usage](#cli-usage)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Features
- **CLI Interface**: Add, list, complete, and delete tasks interactively via a command-line interface.
- **REST API**: Programmatically manage tasks through a simple HTTP API.
- **MySQL Storage**: Persistent storage of tasks in a MySQL database.
- **Modular Architecture**: Organized codebase with separation of concerns (configuration, database, models, repository, service, and handlers).
- **Error Handling**: Robust error handling for database operations, user input, and API requests.
- **Environment-Based Configuration**: Configurable via a .env file for easy setup.

## Project Structure

```
todo-app/
├── cmd/
│   ├── todo-cli/                # CLI application entry point
│   │   └── main.go
│   └── todo-api/                # API application entry point
│       └── main.go
├── internal/
│   ├── config/                  # Configuration management
│   │   └── config.go
│   ├── db/                      # Database connection and setup
│   │   └── db.go
│   ├── models/                  # Data models
│   │   └── task.go
│   ├── repository/              # Data access layer
│   │   └── task_repository.go
│   ├── service/                 # Business logic
│   │   └── task_service.go
│   └── handler/                 # CLI and API handlers
│       ├── cli_handler.go
│       └── api_handler.go
├── scripts/
│   └── init.sql                 # MySQL schema initialization
├── .env                         # Environment variables (not tracked in git)
├── go.mod                       # Go module file
├── go.sum                       # Go module dependencies
└── README.md                    # Project documentation
```

- **cmd/**: Contains the entry points for the CLI (todo-cli) and API (todo-api) applications.
- **internal/**: Houses the core application logic, split into modules for configuration, database, models, repository, service, and handlers.
- **scripts/**: Includes SQL scripts for database initialization.
- **.env**: Stores environment variables (not tracked in version control).
- **go.mod** and **go.sum**: Manage Go dependencies.

## Technologies Used
- **Go**: Programming language for the application.
- **MySQL**: Database for persistent task storage.
- **Gorilla Mux**: HTTP router for the REST API.
- **godotenv**: Loads environment variables from a .env file.
- **go-sql-driver/mysql**: MySQL driver for Go.

## Prerequisites
- **Go**: Version 1.18 or higher.
- **MySQL**: Version 8.0 or higher.
- **Git**: For cloning the repository.
- A text editor or IDE (e.g., VS Code, GoLand).

## Setup Instructions
Follow these steps to set up and run the application locally.

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/yourusername/todo-app.git
    cd todo-app
    ```

2. **Install Dependencies**:
    Install the required Go packages:
    ```bash
    go mod tidy
    ```

    This will download dependencies like `go-sql-driver/mysql`, `godotenv`, and `gorilla/mux`.

3. **Set Up MySQL**:
    Ensure MySQL is installed and running. Create the database and schema by running the provided SQL script:
    ```bash
    mysql -u root -p < scripts/init.sql
    ```

    The script creates a database named `todo_app` and a tasks table.

4. **Configure Environment Variables**:
    Create a `.env` file in the project root based on the example below:
    ```env
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=yourpassword
    DB_NAME=todo_app
    API_PORT=8080
    ```

    Replace `yourpassword` with your MySQL root password or the appropriate credentials.

5. **Verify Setup**:
    Ensure MySQL is running and the `.env` file is correctly configured.

## Running the Application
The application can be run in two modes: CLI or API. Both can run simultaneously if needed.

### CLI Mode
To run the CLI application:
```bash
go run cmd/todo-cli/main.go
```
This starts an interactive CLI where you can manage tasks using commands like add, list, complete, and delete.

### API Mode
To run the REST API:
```bash
go run cmd/todo-api/main.go
```
The API will start on the port specified in the `.env` file (default: 8080). You can interact with it using tools like curl, Postman, or a web browser.

## CLI Usage
The CLI supports the following commands:

| Command             | Description                                | Example                          |
|---------------------|--------------------------------------------|----------------------------------|
| `add <title> [description]`  | Add a new task                            | `add Buy groceries Get milk and bread` |
| `list`              | List all tasks                             | `list`                           |
| `complete <id>`     | Mark a task as complete                    | `complete 1`                     |
| `delete <id>`       | Delete a task                              | `delete 1`                       |
| `help`              | Display available commands                 | `help`                           |
| `exit`              | Exit the CLI application                   | `exit`                           |

### Example CLI Session
```bash
$ go run cmd/todo-cli/main.go
Todo List CLI - Type 'help' for commands
add Buy groceries Get milk and bread
Task added successfully
list
ID: 1, Title: Buy groceries, Description: Get milk and bread, Status: pending, Created: 2025-04-25 10:00:00
complete 1
Task marked as complete
list
ID: 1, Title: Buy groceries, Description: Get milk and bread, Status: completed, Created: 2025-04-25 10:00:00
delete 1
Task deleted
exit
```

## API Endpoints
The REST API provides endpoints to manage tasks programmatically. The base URL is `http://localhost:8080` (or the port specified in `.env`).

| Method | Endpoint                 | Description            | Request Body (if applicable)                  | Response           |
|--------|--------------------------|------------------------|-----------------------------------------------|--------------------|
| POST   | `/tasks`                 | Create a new task      | `{"title": "Task title", "description": "Task description"}` | `201 Created`       |
| GET    | `/tasks`                 | List all tasks         | None                                          | `200 OK, JSON array of tasks` |
| PUT    | `/tasks/{id}/complete`   | Mark a task as complete| None                                          | `200 OK`           |
| DELETE | `/tasks/{id}`            | Delete a task          | None                                          | `200 OK`           |

### Example API Requests (using curl)

- **Create a Task:**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"title":"Test Task","description":"Test Description"}' http://localhost:8080/tasks
    ```

- **List All Tasks:**
    ```bash
    curl http://localhost:8080/tasks
    ```
  Response:
    ```json
    [
        {
            "ID": 1,
            "Title": "Test Task",
            "Description": "Test Description",
            "Completed": false,
            "CreatedAt": "2025-04-25T10:00:00Z"
        }
    ]
    ```

- **Complete a Task:**
    ```bash
    curl -X PUT http://localhost:8080/tasks/1/complete
    ```

- **Delete a Task:**
    ```bash
    curl -X DELETE http://localhost:8080/tasks/1
    ```

### Error Responses
- **400 Bad Request**: Invalid input (e.g., malformed JSON or invalid task ID).
- **500 Internal Server Error**: Database or server errors.

## Environment Variables
The application uses the following environment variables, defined in the `.env` file:

| Variable         | Description               | Default Value  |
|------------------|---------------------------|----------------|
| `DB_HOST`        | MySQL host                | `localhost`    |
| `DB_PORT`        | MySQL port                | `3306`         |
| `DB_USER`        | MySQL username            | `root`         |
| `DB_PASSWORD`    | MySQL password            | (none)         |
| `DB_NAME`        | MySQL database name       | `todo_app`     |
| `API_PORT`       | Port for the REST API     | `8080`         |

## Contributing
Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please ensure your code follows Go best practices and includes appropriate tests.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

Built with ❤️ by Jimit Chavda. For questions or feedback, open an issue or contact [jimitchavdadev@gmail.com](mailto:jimitchavdadev@gmail.com).
