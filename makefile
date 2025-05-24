##########################################
## develop
##########################################

# Make Release Note
release:
	bash ./scripts/release_note.sh

VERSION := $(shell bash ./scripts/merge_version.sh)
VERSION := 1.0.0
tag:
	git tag v$(VERSION)

deploy:
	git push origin v$(VERSION)
	GOPROXY=proxy.golang.org go list -m github.com/david511382/go-test@v$(VERSION)
