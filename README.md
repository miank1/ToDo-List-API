docker build -t my-postgres .
docker run -d -p 5433:5432 --name postgres-container my-postgres
docker stop postgres-container
docker rm 2bab 

docker run -d --name postgres-todo-container -p 5432:5432 postgres-todo