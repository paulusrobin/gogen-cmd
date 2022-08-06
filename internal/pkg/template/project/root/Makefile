help: ## This help dialog
help h:
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-30s %s\n" "target" "help" ; \
	printf "%-30s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
		IFS=$$':' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf '\033[36m'; \
		printf "%-30s %s" $$help_command ; \
		printf '\033[0m'; \
		printf "%s\n" $$help_info; \
	done

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yaml

environment:
	@echo "==========================="
	@echo "Prepare Environment"
	@echo "==========================="
	make go-private
	go mod tidy
	cp .env.sample .env

#===============#
#== App Build ==#
#===============#

build:
	# Build APP Binary
	@echo "==========================="
	@echo "Building APP binary"
	@echo "==========================="
	docker-compose exec -T {{.ProjectName}} go mod tidy && go mod vendor
	docker-compose exec -T {{.ProjectName}} sh -c "go build -mod=vendor -o bin/{{.ProjectName}} cmd/main.go"

run-subscriber:
	docker-compose exec -T {{.ProjectName}} bin/{{.ProjectName}} subscriber

run-http:
	docker-compose exec -T {{.ProjectName}} bin/{{.ProjectName}} http

run-grpc:
	docker-compose exec -T {{.ProjectName}} bin/{{.ProjectName}} grpc

clean: # clean executables
	rm -rf bin/*

#=======================#
#== ENVIRONMENT SETUP ==#
#=======================#

create-env-file:
ifeq (,$(wildcard .env))
	cp .env.sample .env
endif

remove-env-file:
ifneq (,$(wildcard .env))
	rm .env
endif

go-private:
	go env -w GOPRIVATE=github.com/paulusrobin

go-private-in-docker:
	docker-compose exec -T {{.ProjectName}} go env -w GOPRIVATE=github.com/paulusrobin

docker-start:
	@echo "=========================="
	@echo "Starting Docker Containers"
	@echo "=========================="
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d --build --remove-orphans
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

docker-stop:
	@echo "=========================="
	@echo "Stopping Docker Containers"
	@echo "=========================="
	docker-compose -f ${DOCKER_COMPOSE_FILE} stop
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

docker-clean: docker-stop
	@echo "=========================="
	@echo "Removing Docker Containers"
	@echo "=========================="
	docker-compose -f ${DOCKER_COMPOSE_FILE} rm -v -f

docker-restart: docker-stop docker-start

docker-environment: ## The only command needed to start a working environment
docker-environment: remove-env-file docker-restart create-env-file go-private build-native

docker-environment-clean: ## The only command needed to clean the environment
docker-environment-clean: docker-clean clean


#====================#
#== QUALITY CHECKS ==#
#====================#

static-analysis: ## Run all enabled linters
	@echo "======================="
	@echo "Running static analysis"
	@echo "======================="
	docker-compose exec -T {{.ProjectName}} golangci-lint run
