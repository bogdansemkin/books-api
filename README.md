# Start with dockerfile

`
docker run books-api
`

`docker run --name=book-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres`

`migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
`

## Start with docker-compose
`docker-compose up`


# books-api
![image](https://user-images.githubusercontent.com/40574816/175290188-09ef8151-123c-4516-bfaf-318413053365.png)
