REPO := github.com/mashiike/restrating

.PHONY: install goagen regen

install:
	@go get -u goa.design/goa/v3/...

goagen:
	@goa gen $(REPO)/design
