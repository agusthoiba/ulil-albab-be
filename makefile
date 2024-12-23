build:
	go build -o ulil-albab-be.app src/project/main.go


# running
run:
	go run src/project/main.go


# testing
test:
	go test -v ./...

local-test:	
	go test -v ./... -coverprofile=cover.out
	go tool cover -html=cover.out


