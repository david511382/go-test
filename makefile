##########################################
## develop
##########################################

# Make Release Note
release:
	bash ./scripts/release_note.sh

MODULE := go-test
VERSION := $(shell bash ./scripts/merge_version.sh)

tag:
	git tag v$(VERSION)

deploy:
	git push origin v$(VERSION)
	GOPROXY=proxy.golang.org go list -m github.com/david511382/$(MODULE)@v$(VERSION)
