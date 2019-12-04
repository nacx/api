# Copyright (c) Tetrate, Inc 2019 All Rights Reserved.

APIS := alert/v1 config/v1 metadata/v1

.PHONY: all
all: format $(APIS)

.PHONY: $(APIS)
$(APIS):
	@echo "--- $@: all ---"
	$(MAKE) all -C $@

.PHONY: format
format:
	../../scripts/format-protos.sh

.PHONY: clean
clean:
	@for API in $(APIS); do \
		echo "--- $$API: $@ ---"; \
		$(MAKE) $@ -C $$API; \
		if [ $$? -ne 0 ]; then \
			exit 1; \
		fi; \
	done
