help:
	@echo "\033[32mAvailable commands:\033[0m"
	@grep "^[a-zA-Z_-]\+:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

# Goの実行をする
run:
	@echo "\033[32mRunning the application...\033[0m"
	go run src/main.go

stop:
	@echo "\033[32mStopping the application...\033[0m"
	
fmt:
	@echo "\033[32mFormatting the code...\033[0m"
	go fmt ./...