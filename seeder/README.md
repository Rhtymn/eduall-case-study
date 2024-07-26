# Database Seeder

This golang application is used to seed the database. Seeding data is obtained from [Dataset](https://drive.google.com/file/d/1WsYUhEFxsJboSdr1-qU07oPTVcM0gfOy/view). Then the dataset is converted into .csv form to make it easier for the golang application to read.

## How to run

1. Make sure you have installed `golang v1.22.4` on the machine you are using.
2. Make sure you have `Postgre database server` running on the machine you are using.
3. Create an .env file in the root folder `/seeder`.
4. Add variables like those in the `.env.example` file.
5. Ensure compatibility between env variables to connect with the running Postgres server database.
6. Run the `go mod tidy` command in the terminal.
7. Run the `go build` command in the terminal.
8. Run the `.\seeder.exe` command in the terminal.
9. Successfully seeding database.
