clean:
	ls gen/ | grep -v __init__.py | xargs -I {} rm -rf gen/{}

buf:
	cd ../../.. && buf generate --template ./grpc/client/python/buf.gen.yaml
	cd ../../.. && buf generate --template ./grpc/client/python/buf.gen.yaml buf.build/beta/googleapis

setup:
	pip3 install grpclib protobuf

run:
	python3 client.py localhost 8080
