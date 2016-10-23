cmd = '' # for migrate task

.PHONY: setup help server migrate code-generate

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
	migrate -path=./migrations -url='mysql://root:@tcp(localhost:3306)/turntable_build' $(cmd)

code-generate:
	for file in `ls $(GOPATH)/src/github.com/fujimisakari/turntable-build/yaml`; do \
	  target=`basename $$file .yaml`; \
	  echo "Code generate: $${target}"; \
	  go run $(GOPATH)/src/github.com/fujimisakari/turntable-build/_code_generator/generate.go $${target}; \
	  gofmt -w $(GOPATH)/src/github.com/fujimisakari/turntable-build/domain/$${target}/service_master.go; \
	  gofmt -w $(GOPATH)/src/github.com/fujimisakari/turntable-build/model/$${target}_master.go; \
	done
