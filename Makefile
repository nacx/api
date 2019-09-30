# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

APIS := audit/v1 test/v1 q/rbac/v1 tcc/core/v1 tcc/workflows/loadbalancer/v1 regsource/v1
SWAGGERS := $(shell find . -name "*.swagger.json")
OPENAPI_GEN := $(or ${shell which openapi-generator},openapi-generator)

.PHONY: all
all: format $(APIS)

.PHONY: $(APIS)
$(APIS):
	@echo "--- $@: all ---"
	$(MAKE) all -C $@

.PHONY: ts
ts: $(SWAGGERS)

.PHONY: $(SWAGGERS)
$(SWAGGERS):
	@echo "--- generate *.ts from $@ ---"
	$(OPENAPI_GEN) generate -i $@ -g typescript-fetch -o $(shell dirname -- $@)/generated/ts

# TODO(dio): Handle errors when $(OPENAPI_GEN) is not found in an arbitrary env.

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
