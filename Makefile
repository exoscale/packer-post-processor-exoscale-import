include go.mk/init.mk
include go.mk/public.mk

PACKAGE := github.com/exoscale/packer-post-processor-exoscale-import

PROJECT_URL = https://$(PACKAGE)

GO_LD_FLAGS := -ldflags "-s -w -X $(PACKAGE)/version.Version=${VERSION} \
									-X $(PACKAGE)/version.Commit=${GIT_REVISION}"
GO_MAIN_PKG_PATH := ./cmd/packer-post-processor-exoscale-import

EXTRA_ARGS := -parallel 3 -count=1 -failfast

.PHONY: test-acc test-verbose test
test: GO_TEST_EXTRA_ARGS=${EXTRA_ARGS}
test-verbose: GO_TEST_EXTRA_ARGS+=$(EXTRA_ARGS)
test-acc: GO_TEST_EXTRA_ARGS=-v $(EXTRA_ARGS)
test-acc: ## Runs acceptance tests (requires valid Exoscale API credentials)
	PACKER_ACC=1 $(GO) test         \
		-race                       \
		-mod $(GO_VENDOR_DIR)       \
		-timeout 60m                \
		--tags=testacc              \
		$(GO_TEST_EXTRA_ARGS)       \
		$(GO_TEST_PKGS)
