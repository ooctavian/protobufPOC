.PHONY: compile

compile:
	go run .

generate-protobuf:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative schema/model.proto
