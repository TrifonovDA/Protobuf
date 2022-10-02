.PHONY: gen-py
gen-py:
	protoc -I=protos --python_out=python/pb protos/my/my.proto

.PHONY: gen-go
gen-go:
	protoc -I=protos --go_out=go/pb protos protos/my/my.proto --go_out=paths=source_relative

.PHONY: gen
gen:
	buf generate protos

.PHONY: run-py
run-py:
	python3 /mnt/c/Go/JSON_AND_HTTP/protobuf/123.py

.PHONY: run-go
run-go:
	cd go && go run /mnt/c/Go/JSON_AND_HTTP/main.go