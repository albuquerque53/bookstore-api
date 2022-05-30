# BookStore API 📚

> An API for a fictionary library's management

## Tecnical specification 🔍

- Simple CRUD for three repositories:
  - Authors
  - Categories
  - Books
- Used the relational Database MySQL for this.
- The application structure is built based on DDD (Domain-driven design) pattern.
- All application environment is encapsuled into docker (see `build` folder to understand).

## Project set-up 🏗️

1. Build the containers:

```
make up
```

2. Get into API container:

```
make api
```

3. Run the tests to make sure that everyting's all right:

```
make test
```

4. Run the migrations to create the database tables:

```
make migrateup
```

5. Still into container, you can run the application with:

```
make serve
``` 

6. The application is running on the default port `2001` in `localhost`
