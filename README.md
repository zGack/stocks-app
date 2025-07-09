
# Stocks Recommendation App

This app makes recommendations on stock investments
based on `target`, `rating` and `action` stock properties.

## Technical Stack

Stops Recommendation App is built using the following technologies:
- Golang for the server side logic
- Vue3 for the frontend
- CockroachDB for storage

## Screenshots
![image](https://github.com/user-attachments/assets/c3ae0289-22e3-42fe-bbf9-fd11373e5c74)

![image](https://github.com/user-attachments/assets/7c0cc909-fbc6-485b-98d0-b64458092a30)

![image](https://github.com/user-attachments/assets/f5deac89-004a-465c-aff4-a1b0f0ad3391)

## Usage

Follow these steps to run the Stocks Recommendation App locally:

### Setup Environment Variables
`.envrc.template` has the necessary environment variables that need to be set in order run the app.

### Deploy the Local Cluster in Docker

In the root directory of the project run this docker compose command

```bash
docker compose up -d
```

Then open `http://localhost:9090` in your browser to check if the cluster is running.

### Run Database Migrations

After the cluster is running, you need to run the database migrations. You can do this by executing the following command:

```bash
make migrate-up
```

This command will create the necessary tables in the CockroachDB database.

### Start the Server

In a separate terminal, run the following command to start the server:
```bash
go run cmd/api/main.go
```

This will start the server on port you set before in the env file. You can access the API at `http://localhost:{PORT}`.

You can optionally run the server using `air` for live reloading during development

### Start the Frontend App

In a separate terminal, navigate to the `web/` directory and run the following commands:

```bash
npm install
npm run dev
```

This will start the Vue3 frontend app, and you can access it at `http://localhost:5173`.

### Author
Sebastian Mena
[Linkedin](www.linkedin.com/in/sebastian-mena-ferreira-0341b9171)
