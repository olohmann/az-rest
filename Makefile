GOFMT_FILES?=$$(find . -name '*.go')

default: vet

tools:
	go get -u github.com/kardianos/govendor
	go get -u golang.org/x/tools/cmd/stringer
	go get -u golang.org/x/tools/cmd/cover

# bin generates the releaseable binaries for Terraform
bin: fmtcheck
	@AZ_REST_RELEASE=1 sh -c "'$(CURDIR)/scripts/build.sh'"

# dev creates binaries for testing Terraform locally. These are put
# into ./bin/ as well as $GOPATH/bin
dev: fmtcheck
	@AZ_REST_DEV=1 sh -c "'$(CURDIR)/scripts/build.sh'"

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@echo 'go vet ./...'
	@go vet ./... ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

clean:
	rm -rf bin/
	rm -rf pkg/

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: bin default dev fmt fmtcheck tools vet
