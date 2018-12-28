PROTOC ?= protoc

build: prepare
	${PROTOC} \
		--js_out=import_style=commonjs,binary:js \
		--go_out=./go \
        --ts_out=./js \
       ./commproto.proto

prepare:
	mkdir -p js go


.PHONY: build
