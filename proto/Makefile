default: build

build:
	python -m grpc_tools.protoc -I=. --python_out=../src/python/protogen --grpc_python_out=../src/python/protogen ./*.proto
	protoc --go_out=plugins=grpc:../src/go/protogen *.proto