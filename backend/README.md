# Backend Application

This backend application is used to retrieve product data from the database. This application is made using golang programming language with gin framework.

## API Endpoint

1. GET `/api/v1/products` for retrieving products data.

## Postman Collections

You can see the postman collection in the root folder `/backend`.

## How to run

1. Make sure you have installed `golang v1.22.4` on the machine you are using.
2. Make sure you have `Postgre database server` running on the machine you are using.
3. Create an .env file in the root folder `/backend`.
4. Add variables like those in the `.env.example` file.
5. Ensure compatibility between env variables to connect with the running Postgres server database.
6. Run the `go mod tidy` command in the terminal.
7. Run the `go build` command in the terminal.
8. Run the `.\backend.exe` command in the terminal.
9. The backend application runs successfully.
