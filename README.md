# Go-Library-API
Library implementation with database (array)
1) Terminal >> go run main.go
2) Separate terminal >> send curl request from the list below:

1. To see all added books:
curl 'localhost:8080/books'

2. To see Book by id:
curl 'localhost:8080/books/3'

3. Add book from "body.json" in library:
curl 'localhost:8080/books' -i -H "Content-Type: application/json" -d "@body.json" --request "POST"

4. To buy 3 Books by id
curl 'localhost:8080/books/buy?id=3&quantity=3' --request "PATCH"