PROTOC ?= protoc

proto: build-worldcomm build-coordinator

prepare:
	mkdir -p js go


build-worldcomm: prepare
	${PROTOC} \
		--js_out=import_style=commonjs,binary:js \
		--go_out=./go \
        --ts_out=./js \
       ./worldcomm.proto


build-coordinator: prepare
	${PROTOC} \
		--js_out=import_style=commonjs,binary:js \
		--go_out=./go \
        --ts_out=./js \
       ./coordinator.proto

.PHONY: proto
