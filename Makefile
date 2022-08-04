up-migrations:
	docker exec -i mysql mysql -uroot -proot db_accounting < ./database/database-up.sql

down-migrations:
	docker exec -i mysql mysql -uroot -proot db_accounting < ./database/database-down.sql
