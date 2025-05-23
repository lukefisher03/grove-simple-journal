### Creating a migration
```
migrate create -ext sql -dir db/migrations -seq create_users_table
```

### Applying a migration

`POSTGRESQL_URL` could look something like: `postgres://username:password@localhost:5432/grove`
```
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```