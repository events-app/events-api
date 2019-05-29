# events-api

- Start a PostgreSQL server running in a docker container
- Setup schema
- Seed database
- Connect to database from service

## Notes:


```
## Compile api and events-admin
go isntall ./...

## Start postgres:
docker-compose up -d

## Create the schema and insert some seed data.
go build
./events-admin migrate
./events-admin seed

## Run the app then make requests.
./api
```