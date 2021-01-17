download:
	go mod download
run: 
	nodemon --exec go run app/*.go --signal SIGTERM