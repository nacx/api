# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

# Allow consumers to add targets or override variables in this Makefile.
-include ../api.mk

APIS := audit/v1 test/v1 q/rbac/v1 tcc/core/v1 tcc/workflows/loadbalancer/v1 regsource/v1 spm authz/v1 \
	troubleshoot/cluster/v1 settings/v1 operator/common operator/tsboperator/v1alpha1 \
	tsb/gateway/v1 tsb/security/v1 tsb/traffic/v1 tsb/types/v1 tsb/v1 tsb/rbac/v1

.PHONY: all
all: format $(APIS)

.PHONY: $(APIS)
$(APIS):
	@echo "--- $@: all ---"
	$(MAKE) all -C $@

.PHONY: test
test:
	go test ./...

.PHONY: format
format:
	../scripts/format-protos.sh

.PHONY: clean
clean:
	@for API in $(APIS); do \
		echo "--- $$API: $@ ---"; \
		$(MAKE) $@ -C $$API; \
		if [ $$? -ne 0 ]; then \
			exit 1; \
		fi; \
	done
