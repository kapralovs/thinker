docker pull postgres:latest
docker run --name thinker-pg -e POSTGRES_PASSWORD=qwerty -dp 5432:5432 postgres:latest