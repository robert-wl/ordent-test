



run:
	air --build.cmd "go build -o bin/api cmd/app/main.go" --build.bin "./bin/api"


swagger:
	swag init -g .\cmd\app\main.go