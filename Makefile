include go.mk/init.mk

PACKAGE := github.com/exoscale/packer-post-processor-exoscale-import

GO_LD_FLAGS := -ldflags "-s -w -X $(PACKAGE)/version.Version=${VERSION} \
									-X $(PACKAGE)/version.Commit=${GIT_REVISION}"
GO_MAIN_PKG_PATH := ./cmd/packer-post-processor-exoscale-import

EXTRA_ARGS := -parallel 3 -count=1 -failfast

.PHONY: test-verbose test
test: GO_TEST_EXTRA_ARGS=${EXTRA_ARGS}
test-verbose: GO_TEST_EXTRA_ARGS+=$(EXTRA_ARGS)
