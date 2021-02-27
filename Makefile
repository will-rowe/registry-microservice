all: proto docs fmt lint vet test build

proto:
		protoc -I.  --go_out=plugins=grpc:internal api/proto/v1/registryService.proto

docs:
		protoc -I. --doc_out=api/docs/v1 --doc_opt=markdown,registryService.md api/proto/v1/registryService.proto	

fmt:
		go list ./... | grep -v /api/ | go fmt

lint:
		go list ./... | grep -v /api/ | xargs -L1 golint -set_exit_status

vet:
		go vet ./...

test:
		go test -v ./...

build: proto
		go mod tidy
		CGO_ENABLED=0 go build -o ./bin/registry .

pack: build
		docker build -t willrowe/registry-service:latest .

push:
		docker push willrowe/registry-service:latest
	
serve:
		docker run -p 9090:9090 willrowe/registry-service

clean:
		rm -r bin