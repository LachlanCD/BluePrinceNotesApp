# Blue Prince Notes App

This is a full-stack web application with:

- **Backend**: Go (Golang)
- **Frontend**: React with Tailwind CSS
- **Database**: SQLite
- **Containerization**: Docker

---

## Project Structure

```
.
├── backend/              # Go backend code
│   └── main.go
|   └──data/                  # SQL DB
│      └── notes.db
├── frontend/             # React + Tailwind frontend
│   ├── public/
│   └── src/
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md
```

---

##  Features

- REST API built with Go
- Tailwind-powered React frontend
- SQL database integration
- Easily deployable via Docker

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/LachlanCD/BluePrinceNotesApp
cd your-project
```

2. Move the .env-template to .env
```bash
mv .env-template .env
```

3. Build and start the containers:

```bash
docker-compose up --build
```

4. Access the app:

- Frontend: [http://localhost:4000](http://localhost:4000)


---
## Deployment

 the app is currently hosted at: [URL](https://blueprince.ashycoast-f2648366.australiaeast.azurecontainerapps.io/) on Azure and the changes in [deployment](https://github.com/LachlanCD/BluePrinceNotesApp/tree/deployment) reflect the changes to consolidate this (due to scaling to 0 instances the application may take a bit to load while the container spins up).

The .env-template file has extra fields DB_SERVER, DB_USER, DB_PASSWORD, DB_NAME to connect to the azure db instance.

---
## Docker Setup

### Dockerfile

Single multi-stage Dockerfile builds backend and frontend.

### docker-compose.yml

Combines:
- Go app container
- React build served with Go or Nginx
- SQL DB (e.g., Postgres or SQLite volume)

---

## Testing

Basic tests are setup on Go backend and can be run with:

```bash
cd backend
go run test
```

---

## License

MIT

---

## Author

Created by [Lachlan](https://github.com/LachlanCD)

