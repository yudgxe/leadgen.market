test:
	go test ./... -timeout 60s

reup: 
	docker-compose up --build app