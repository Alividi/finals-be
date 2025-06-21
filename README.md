# finals-be

Backend repository for my final project **Digisatlink**

## Migrations

Migration using [golang-migrate](https://github.com/golang-migrate/migrate/tree/master)

Install the CLI : https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md

**Run script :**
```sql
migrate -path migrations -database postgres://postgres:root@localhost:5432/finals_db?sslmode=disable up
```

**Run script for zsh terminal :**
```sql
migrate -path migrations -database "postgres://postgres:root@localhost:5432/finals_db?sslmode=disable" up
```