include go.mk/init.mk

GO_MAIN_PKG_PATH := ./cmd/packer-post-processor-exoscale-import

EXTRA_ARGS := -parallel 3 -count=1 -failfast

.PHONY: test-acc test-verbose test
test: GO_TEST_EXTRA_ARGS=${EXTRA_ARGS}
test-verbose: GO_TEST_EXTRA_ARGS+=$(EXTRA_ARGS)
