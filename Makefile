.PHONY: shooter
shooter:
	go run cmd/shooter/main.go

.PHONY: dungeon
dungeon:
	go run cmd/dungeon/main.go


.PHONY: format
format:
	@echo "Start formatting code..."
	@sh -c '\
		export PATH=$$PATH:$$(go env GOPATH)/bin; \
		if ! command -v gotools >/dev/null 2>&1; then \
			echo "gotools not found, installing..."; \
			go install github.com/zq-xu/gotools@latest; \
		fi; \
		gotools format \
	'
	@echo "Finished formatting code."