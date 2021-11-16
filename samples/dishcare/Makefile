clean:
	ls gen/ | grep -v __init__.py | xargs -I {} rm -rf gen/{}

buf:
	export PATH=${PATH}:$(shell python3 -m site --user-base)/bin && cd ../.. && buf generate --template ./samples/dishcare/buf.gen.yaml
	export PATH=${PATH}:$(shell python3 -m site --user-base)/bin && cd ../.. && buf generate --template ./samples/dishcare/buf.gen.yaml buf.build/googleapis/googleapis

setup:
	pip3 install grpclib protobuf

package: setup buf
	tar -czf viam.tgz gen dump.py viam.json

run_client:
	python3 client.py local.jetson.unyyjqwtoa.viam.cloud 8080
