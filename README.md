# 📅TodoList



a simple TodoList builting with Vue ,Go and MySQL

## Why this project

- Learn html/css/js and a web framework as a gopher
- Understand how the frontend and backend  interact by a simple CRUD  project
- Try to deploy applications by using docker

## Demo

- https://todo.nc-77.top 

  <img src="https://img.nc-77.top/20210805143307.png" style="max-width:67%;" />

## Usage

### Docker

```bash
#1	Clone this repository
git clone https://github.com/nc-77/todoList.git && cd todoList
#2	Create docker network
docker create network todo
#3  Build App Server
docker-compose up
#4	Go to http://localhost:8081 and enjoy it  
```

tips: The web and api runs on 8081,8082 ports by defalut.

You can custom the port config by change the  backend/config/config.toml and dockerfile-compoes.yml

### Local

1. Clone this repository

2. Finish mysql config in backend/config/config.toml and run data/mysql-init/tasks.sql in MySQL.

3. Run backend

   ```bash
   cd backend && go run main.go
   ```

4. Run frontend

   ```bash
   cd frontend
   yarn install
   yarn serve
   ```

5. Go to http://localhost:8080 and enjoy it

