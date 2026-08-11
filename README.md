# Marks ![Static Badge](https://img.shields.io/badge/weather-%23b16ded?style=flat&logo=github&logoColor=black&labelColor=0%2C0%2C0&link=https%3A%2F%2Fgithub.com%2FweatherGod3218%2F)

### How Fast can ___YOU___ Eat a garbage plate?!?

A simple timer based system for tracking how fast members of CSH can eat a garbage plate. Ohhh the things we do

_This project uses Golang, [Gin](https://gin-gonic.com/en/), and Vite React._

## Documentation
All API Endpoints have swagger doucmentation, which can be viewed at `/swagger/index.html`

# Local Development
Pre - Requirements:
 - Node v20+
 - go 1.26+
   - air (`go install github.com/air-verse/air@latest`)
 - PostgreSQL database
 - ___OPTIONAL___: Redis cache

## Docker Compose

Marks uses docker compose for local development. The entire suite can be started by running
```
    docker compose up --build
```

Marks utilizes a local proxy for local development. Whenever air is used, it creates a proxy
for the Vite frontend development instead. This route allows all of Vite's features
along with running the application.

## Starting NPM Dev Server
For development, enter `web/` and complete the following:

- install web dependencies
  ```bash
  npm i
  ```

- run development server
  ```bash
  npm run dev
  ```

The development server will begin at port 5173. Do not use this port to access the project.

### Starting Backend
In a new terminal, complete the following:

- install dependencies
   ```bash
   go mod download
   ```

- begin development backend server
   ```bash
   air
   ```

Shoutout to Tyler Severino for the proxy + air approach for local development.
