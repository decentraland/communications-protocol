PROTOC ?= protoc

proto:
	${PROTOC} \
		--js_out=import_style=commonjs,binary:js \
		--go_out=go \
        --ts_out=js \
       ./worldcomm.proto

.PHONY: proto
