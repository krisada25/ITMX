ให้ run main.go ก่อนครับ แล้วใช้ postman ทดสอบครับ
หลัง http://localhost:8080/customers/{id} autorun ครับ

1. POST for create
curl --location --request DELETE 'http://localhost:8080/customers/1' \
--data ''

2. PUT for update -put will update every field by id, if no record found reject error
404
curl --location --request PUT 'http://localhost:8080/customers/1' \
--header 'Content-Type: application/json' \
--data '{"name":"nut","age":40}'


3. DELETE /customers/{id} – delete by provide record id
curl --location 'http://localhost:8080/customers/1' \
--data ''


4. GET /customers /{id} – get by provided employee id
curl --location 'http://localhost:8080/customers/1' \
--data ''
