.PHONY: compile

compile:
	go run .

generate-protobuf:
	protoc --go_out=. --go_opt=paths=source_relative schema/model.proto
