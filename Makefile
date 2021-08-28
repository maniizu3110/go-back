sql:
	docker-compose up -d

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:13306)/simple_bank" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:13306)/simple_bank" -verbose down

login mysql:
	docker exec -it mysql  mysql -u root -p 

.PHONY: sql migrateup migratedown login mysql