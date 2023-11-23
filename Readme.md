# Goodies
Goodies is a microservice that serve product functionalities such as posting a product, submit review, and search for products by it's SKU, title, and category. This microservice runs on RESTful API protocols.

## Dependencies
- Go 1.20
- MySQL 8.0
- Docker (optional)
- [golang-migrate](https://github.com/golang-migrate/migrate)

## Usage
1. Clone this repository
```bash
git clone git@github.com:rakhmatullahyoga/goodies.git
```
2. Setup dependencies via Docker (optional)
```bash
docker-compose up -d
```
3. Run database migration to setup database schema. To run the migration, please follow the instruction on the [developer page](https://pkg.go.dev/github.com/golang-migrate/migrate/cli#section-readme)
4. Setup environment variables. You can easily setup environment variables using `.env` file by copying from the `env.sample` file and modifying the `.env` file
```bash
make env
```
5. Compile the project
```bash
make compile
```
6. Run the application
```bash
./goodies
```
## Test the application
You can now test the application by making http request to goodies service as described in the attached [Postman collection](goodies.postman_collection.json)