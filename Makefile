WORK_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
APPLICATION_DIR = ~/Library/Application\ Support/Alfred\ 2/Alfred.alfredpreferences/workflows
BUNDLE_ID = user.workflow.63F60794-BB56-4415-9372-BAF974C3A7E1
CLI_CMD = ./workflow/DeepL-Alfred-translator

default: build cli

deploy: clean build unlink link

release: clean build package

cli:
	@echo "--> Running CLI commands"
	@$(CLI_CMD) pos sta

build:
	@echo "--> Compiling packages and dependencies"
	@mkdir -p ./workflow/
	go build -ldflags '-s -w' -o $(CLI_CMD)

package:
	@echo "--> Package to Workflow"
	@- rm -rf ./build
	@mkdir -p ./build
	ditto -ck ./workflow ./build/DeepL-Alfred-translator.alfredworkflow

test:
	@echo "--> test binary is it works"
	./workflow/DeepL-Alfred-translator -q amiable

clean:
	@echo "--> Cleaning workflow execute files"
	@- rm -f "$(CLI_CMD)"

link:
	@echo "--> Linking workflow files"
	ln -snf $(WORK_DIR)/workflow $(APPLICATION_DIR)/$(BUNDLE_ID)

unlink:
	@echo "--> Unlinking workflow files"
	@- rm $(APPLICATION_DIR)/$(BUNDLE_ID)

.PHONY: default deploy release package build cli test clean link unlink