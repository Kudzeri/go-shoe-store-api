# Shoe Store API

Shoe Store API is a RESTful API backend for managing products, orders, and users for a shoe store.

## Table of Contents

- [Installation](#installation)
- [Running the API](#running-the-api)
- [API Documentation](#api-documentation)
- [Example API Requests](#example-api-requests)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/shoe-store-api.git
   cd shoe-store-api
   ```

2. **Install dependencies:**

   Make sure you have Go installed. Then, install project dependencies:

   ```bash
   go mod download
   ```

3. **Set up the database:**

    - Create a PostgreSQL database named `shoe_store_db`.
    - Set up database connection details in `config/config.go`.

4. **Environment variables:**

   Create a `.env` file in the root directory and define environment variables:

   ```plaintext
   PORT=8080
   DATABASE_URL=postgres://user:password@localhost:5432/shoe_store_db?sslmode=disable
   JWT_SECRET=your_jwt_secret
   ```

   Adjust `DATABASE_URL` and `JWT_SECRET` as per your setup.

## Running the API

1. **Build and run the application:**

   ```bash
   go build -o main ./cmd/main
   ./main
   ```

   The API server will start running at `http://localhost:8080`.

2. **Docker (optional):**

   Alternatively, you can run the application using Docker:

   ```bash
   docker build -t shoe-store-api .
   docker run -p 8080:8080 shoe-store-api
   ```

## API Documentation

The API is documented using Swagger. Access the API documentation at `http://localhost:8080/swagger/index.html` after starting the server.

## Example API Requests

Here are some example API requests:

### Get all products

```http
GET /products
```

### Create a new product

```http
POST /products
Content-Type: application/json

{
  "name": "Nike Air Max",
  "description": "Running shoes",
  "price": 129.99
}
```

### Get a product by ID

```http
GET /products/{id}
```

### Register a new user

```http
POST /users/register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "johndoe@example.com",
  "password": "password"
}
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
```

### Explanation

- **Project Description**: Briefly describes what the project does.
- **Installation**: Provides steps to clone the repository, install dependencies (`go mod download`), set up the database, and configure environment variables.
- **Running the API**: Explains how to build and run the application locally (`go build`, `./main`) or using Docker (`docker build`, `docker run`).
- **API Documentation**: Mentions where to access the Swagger documentation once the server is running (`http://localhost:8080/swagger/index.html`).
- **Example API Requests**: Provides example HTTP requests (`GET`, `POST`) for interacting with the API endpoints (`/products`, `/users/register`, etc.).
- **Contributing**: Encourages contributions and provides a brief guide on how to contribute to the project.
- **License**: States the project's licensing terms (MIT License).

### Usage

- Customize the `README.md` template to fit your specific project details, endpoints, and requirements.
- Include any additional sections or details relevant to your project, such as authentication mechanisms, deployment instructions, or testing guidelines.

By following this template, you ensure that anyone accessing your repository can quickly understand, set up, and interact with your Go shoe store API project effectively. Adjust and expand the `README.md` as needed to provide comprehensive documentation for your users and contributors.