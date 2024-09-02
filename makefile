# build proto
.PHONY: mockery
protobuild:
	cd $(shell pwd)/protos;  buf generate;