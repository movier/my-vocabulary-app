# Running the project
```
docker-compose up --build
```
http://localhost:8080/upload

# Database Migration
```
docker exec -it my_postgres bash
psql -U postgres
SELECT pid, pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'my_database' AND pid <> pg_backend_pid();
DROP database my_database;
Create database my_database;
```