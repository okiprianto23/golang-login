# Login Microservice

This is a Go-based microservice for handling user login functionality. It provides endpoints for user authentication, session management, and related features.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features

- User authentication
- JWT token generation and validation
- Session management
- Secure password handling with bcrypt

## Installation

### Prerequisites

- Go 1.18+
- Docker (for running the database and other dependencies)

### Steps

1. Clone the repository:
    ```sh
    git clone https://github.com/ZW-Rays/golang-login.git
    cd golang-login
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Start the database using Docker:
    ```sh
    docker-compose up -d
    ```

4. Run the application:
    ```sh
    go run main.go
    ```

## Configuration

Configuration can be managed through environment variables or a configuration file. The following environment variables are used:

- `PORT`: The port on which the server runs (default: `8080`)
- `DB_HOST`: Database host (default: `localhost`)
- `DB_PORT`: Database port (default: `5432`)
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `JWT_SECRET`: Secret key for JWT

Example `.env` file:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
JWT_SECRET=yourjwtsecret
