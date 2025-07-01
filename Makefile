migrate-create:
	docker run --rm -v "$(PWD)/migrations:/migrations" \
	  migrate/migrate:v4.16.2 create -ext sql -dir /migrations -seq $(name)