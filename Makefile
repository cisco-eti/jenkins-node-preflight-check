export PROJECT_ROOT=$(shell pwd)
export GO111MODULE=on
REPO_NAME = wwwin-github.cisco.com/eti/sre-go-helloworld

all: build

deps:
	go mod download
	go mod tidy

clean:
	@rm -rf coverage coverage.html

test:
	echo "Running tests"
	go test -v --cover ./...

sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"

test_debug: debug_kill
	@cd $(PROJECT_ROOT)
	dlv test $(REPO_NAME) --headless --api-version=2 --listen "0.0.0.0:2345" --log=true

debug_kill:
	-kill -9 `ps -ef | grep 'dlv debug' | grep -v grep | awk '{print $$2}'`
	-kill -9 `ps -ef | grep '/debug' | grep -v grep | awk '{print $$2}'`
