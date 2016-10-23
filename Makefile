target := ''

help:
	@echo "Usage:"
	@echo "  make setup                      # install required modules"
	@echo "  make server                     # run server"
	@echo "  make migrate [target=portNo]    # execute data migration. target is model name."
	@echo "  make code-generate              # auto generate master_service and master_model by yaml"

setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

server:
	go run main.go

migrate:
	./_tools/migrate.sh $(target)

code-generate:
	./_tools/code_generate.sh

.PHONY: setup help server migrate code-generate
