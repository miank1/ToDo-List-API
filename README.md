docker build -t my-postgres .
docker run -d -p 5433:5432 --name postgres-container my-postgres
docker stop postgres-container
docker rm 2bab 

docker run -d --name postgres-todo-container -p 5432:5432 postgres-todo

http://localhost:8080/todos?sort=created_at&order=asc

http://localhost:8080/todos?sort=due_date&order=asc

Now /todos supports:

Pagination → ?page=1&limit=5

Filtering → ?status=pending

Sorting → ?sort=due_date&order=asc