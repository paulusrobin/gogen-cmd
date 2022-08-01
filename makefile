build:
	# Build APP Binary
	@echo "==========================="
	@echo "Building APP binary"
	@echo "==========================="
	go mod tidy && go mod vendor
	go build -mod=vendor -o bin/gogen cmd/main.go

build-bin:
	make build
	sudo rm /usr/local/bin/gogen
	sudo cp bin/gogen /usr/local/bin/gogen

initialize-project:
	make build-bin
	gogen init -n example -m github.com/paulusrobin/example