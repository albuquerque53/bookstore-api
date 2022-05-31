# BookStore API 📚

> An API for a fictionary library's management

## Tecnical specification 🔍

- Simple CRUD for three repositories:
  - Authors
  - Categories
  - Books
- Used the relational Database MySQL for this (see the [ER Diagram here](.github/bookstore_db.png)).
- The application structure is built based on DDD (Domain-driven design) pattern.
- All application environment is encapsuled into docker (see [build folder](build/) to understand).

## Project set-up 🏗️

0. Create the `.env` file inside `build` folder with your local configurations (you must use the `.env.example` as base)

1. Build the environment and get into API container:

```
make up
```

2. Install the dependencies:

```
make install
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
