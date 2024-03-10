## Copyright 2024 Edward Bennett.  All rights reserved.
## Use of this source code is governed by an Apache 2.0
## style license that can be found in the LICENSE file.

GOCMD=go
TARGET=dwagd

all: build

build:
	@echo "Building..."
	@$(GOCMD) build -o $(TARGET) .

run:
	@echo "Running..."
	@echo "Press CTRL-C to quit"
	@$(GOCMD) run .

test:
	@echo "Testing..."
	@$(GOCMD) test -v

test_coverage:
	@echo "Generating test coverage report..."
	@$(GOCMD) test -coverprofile=coverage.out

clean:
	@echo "Cleaning..."
	@$(GOCMD) clean
	@rm -f $(TARGET)
	@rm -f coverage.out

.PHONY: all build clean run test
