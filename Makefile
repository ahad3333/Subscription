migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://postgres:0003@localhost:3003/subscriptionserver?sslmode=disable' up

swag-gen:
	swag init -g api/api.go -o api/docs

swag:
	export PATH=$(go env GOPATH)/bin:$PATH
run:
	go run cmd/main.go

test_Id:
	go test -run TestBookGetById -v

test_Insert:
	go test -run TestBookInsert -v  

migrate:
	migrate create -ext sql -dir ./migrations/postgres -seq -digits 2 create_trigger


