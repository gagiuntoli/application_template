# Application Template: React + Go + MySQL

This is full dockerize application with a frontend (in React) and a backend server written in Go.

The backend server communicates with a MySQL database.

# Docker & Docker Compose

Docker compose allows to set the environment to run the application on any machine. First
check that you have Docker & Docker Compose installed:

```bash
$ docker --version
Docker version 25.0.4, build 1a576c5
$ docker-compose --version
Docker Compose version v2.10.2
```

Build docker images and run them as containers in the background:

```bash
docker-compose up -d
```

Stop the containers

```bash
docker-compose down
```
