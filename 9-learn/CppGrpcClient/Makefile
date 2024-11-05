clean:
	rm -rf build

proto:
	rm -rf src/user_service.*
	@protoc -I src/proto --grpc_out=src --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` src/proto/*.proto
	@protoc -I src/proto --cpp_out=src src/proto/*.proto
	echo 'generator file success, see src/*'

.PHONY: proto clean 
