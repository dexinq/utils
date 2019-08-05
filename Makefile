.PHONY: proto

proto:
	for d in proto; do \
		for f in $$d/**/*.proto; do \
		    protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
		    echo "protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f"; \
			echo compiled; $$f; \
		done \
	done \