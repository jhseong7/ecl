help: ## This help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

sample: ## Run the default sample
	@go run cmd/sample/main.go
sample-spring: ## Run the spring style sample
	@go run cmd/spring-style/main.go
sample-nestjs: ## Run the nestjs style sample
	@go run cmd/nestjs-style/main.go
sample-filestream: ## Run the filestream style sample
	@go run cmd/filestream/main.go