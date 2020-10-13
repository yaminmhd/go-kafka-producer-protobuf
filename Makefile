APP=go-kafka-producer-protobuf
APP_EXECUTABLE="./out/$(APP)"
PROJ_DIR="."

setup:
	go get golang.org/x/tools/cmd/goimports

generate:
	@echo "Generate person proto..."
	@protoc --go_out=. -I=$(PROJ_DIR) $(PROJ_DIR)/protos/Person.proto

clean:
	rm generatedProtos/person/*.go

compile:
	@echo "Building executable..."
	@mkdir -p out/
	@go build -o $(APP_EXECUTABLE)

vet:
	@echo "Running vet..."
	@go vet ./...

build: vet compile

run: compile
	@echo "Publishing message to kafka"
	@$(APP_EXECUTABLE)