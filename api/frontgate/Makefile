# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

PROTO_FILES=$(sort $(wildcard ./*.proto))
PROTOC_FLAGS:=-I.

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  generate          to generate protobuf code"
	@echo "  tools             to install depends tools"
	@echo "  help              to show this help informatrion"

generate: $(PROTO_FILES) Makefile
	mkdir -p ../../pkg/pb/frontgate
	rm -rf ../../pkg/pb/frontgate/*
	protoc $(PROTOC_FLAGS) --go_out=../../pkg/pb/frontgate ${PROTO_FILES}
	protoc $(PROTOC_FLAGS) --stdrpc_out=../../pkg/pb/frontgate ${PROTO_FILES}
	@echo "ok"

tools:
	docker pull openpitrix/openpitrix:builder
	@echo "ok"

clean:
	@echo "ok"
