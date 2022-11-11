generate:
	go generate ./...
doc:
	godoc -http=:8080
d_run:
	docker rm --force gonear_container
	docker run -dit -p 4000:4000 --env-file ./.env --name gonear_container gonear_image
d_build:
	docker build -t gonear_image .
d_exec:
	docker exec -it gonear_container /bin/bash
