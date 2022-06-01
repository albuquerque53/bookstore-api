# BookStore API 📚

> An API for a fictionary library's management

[![Run in Insomnia}](https://insomnia.rest/images/run.svg)](https://insomnia.rest/run/?label=BookStore%20API&uri=https%3A%2F%2Fgithub.com%2Falbuquerque53%2Fbookstore-api%2Fblob%2Fmain%2Fdoc%2Finsomnia.json)

## Tecnical specification 🔍

- Simple CRUD for three repositories:
  - Authors
  - Categories
  - Books
- Used the relational Database MySQL for this (see the [ER Diagram here](.github/bookstore_db.png)).
- The application structure is built based on DDD (Domain-driven design) pattern.
- All application environment is encapsuled into docker (see [build folder](build/) to understand).
- This API have a Swagger/Open API documentation, [take a look here](doc/swagger.yaml).

## Project set-up 🏗️

0. Create the `.env` file inside root (you must use the `.env.example` as base)

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
