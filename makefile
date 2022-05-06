
help:
	@echo
	@echo Description of available options
	@echo
	@echo   'make env-up             "Start the project containers"'
	@echo   'make env-up-db          "Start the project container just for the database"'
	@echo   'make env-up-db-tools    "Start the project developer just for the database and pgadmin"'
	@echo   'make env-down           "Stopt the project developer containers"'
	@echo   'make env-prune          "Delete all volumes data, use when it is necessary to reset environment data"'
	@echo

env-up:
	@echo
	@echo Start the project containers
	docker-compose up -d
	@echo

env-up-db:
	@echo
	@echo Start the project container just for the database
	docker-compose up -d postgres
	@echo

env-up-db-tools:
	@echo
	@echo Start the project developer just for the database and pgadmin
	docker-compose up -d postgres pgadmin
	@echo

env-down:
	@echo
	@echo Stop the project developer containers 
	docker-compose down
	@echo

env-prune:
	@echo
	@echo Delete all volume data
	docker volume prune
	@echo

