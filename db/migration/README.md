## Run a docker container of postgres
```
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine
```

## Create migration
```
migrate create -ext sql -dir . -seq init_schema   
```

## Create the database
```
docker exec -it some-postgres createdb --username=postgres --owner=postgres simple_bank
```

## Run migrations up
```
migrate --path . --database "postgresql://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## Run migrations down
```
migrate --path . --database "postgresql://postgres:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down
```