# eduall-case-study

This repository was created to fulfil the case study of eduall which consists of `backend`, `frontend`, and `seeder` applications. You can run each application by looking at the guide in README.md in each folder. You can also run it using Docker.

## How to run using Docker

1. Make sure you have installed docker on the machine you are using.
2. Create an `.env.docker` file in the root folder.
3. Add variables like those in the `.env.docker.example` file.
4. Ensure compatibility between env variables to connect with the running Postgres server database on `/seeder/.env` and `/backend/.env`. Please check carefully.
5. Ensure compatibility between env variables to connect backend application on `/frontend/.env`. Please check carefully.
6. Run the `docker compose up` command in the terminal.
7. Wait for all containers to run successfully.
8. Frontend application will running on port 3000, and the backend application on port 8080

## How to run manually

Please refer to the guidelines for each project in README.md.
