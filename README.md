# Multi-Layer REST API Application

This project is a multi-layer web application developed using the Gin framework for serving a REST API, with PostgreSQL as the database management system. The database queries are handled using SQL, and the application is tested with Postman. The PostgreSQL database runs in a Docker container.

## Features
- Developed using the **Gin** framework for efficient HTTP routing and RESTful API development.
- **PostgreSQL** is used as the relational database management system.
- **SQL** queries handle database interactions.
- The database runs inside a **Docker** container for easy management.
- **Postman** is used for testing all API endpoints.
- Follows multi-layer architecture for better code organization and separation of concerns:
  - **Router Layer**: Handles HTTP requests and routes them to the appropriate handler functions.
  - **Service Layer**: Contains business logic for processing data.
  - **Repository Layer**: Manages interactions with the database.

## Technologies

- [Go (Golang)](https://golang.org/)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Postman](https://www.postman.com/)

