SHELL=/usr/bin/env bash
VENV=bazel-venv

all: bazel-venv
	. $(VENV)/bin/activate; bazel build //...

test: bazel-venv
	. $(VENV)/bin/activate; bazel test //...

gofmt:
	gofmt -w -s pkg/ cmd/

gazelle: bazel-venv
	. $(VENV)/bin/activate; bazel run //:gazelle

gazelle-update-repos: bazel-venv
	. $(VENV)/bin/activate; bazel run //:gazelle -- update-repos -from_file=go.mod

bazel-venv: $(VENV)/bin/activate

bazel-clean:
	rm -rf bazel-*

$(VENV)/bin/activate:
	test -d $(VENV) || virtualenv $(VENV) --python=python2
	touch $(VENV)/bin/activate

