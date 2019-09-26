# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

APIS := audit/v1 test/v1 tcc/core/v1 tcc/workflows/loadbalancer/v1 regsource/v1

.PHONY: all
all: $(APIS)

.PHONY: $(APIS)
$(APIS):
	@echo "--- $@: all ---"
	$(MAKE) all -C $@

.PHONY: clean
clean:
	@for API in $(APIS); do \
		echo "--- $$API: $@ ---"; \
		$(MAKE) $@ -C $$API; \
		if [ $$? -ne 0 ]; then \
			exit 1; \
		fi; \
	done
