CC=gcc
CFLAGS=-Wall -Wextra -std=c99 -O2
SRCDIR=src
SOLUTIONDIR=solutions
INPUTDIR=input

# Find all .c files in src directory
SOURCES=$(wildcard $(SRCDIR)/day*.c)
TARGETS=$(patsubst $(SRCDIR)/%.c,$(SOLUTIONDIR)/%,$(SOURCES))

.PHONY: all clean test

all: $(TARGETS)

$(SOLUTIONDIR)/%: $(SRCDIR)/%.c | $(SOLUTIONDIR)
	$(CC) $(CFLAGS) -o $@ $<

$(SOLUTIONDIR):
	mkdir -p $(SOLUTIONDIR)

clean:
	rm -rf $(SOLUTIONDIR)

test: all
	@echo "Running all solutions..."
	@for day in $(TARGETS); do \
		if [ -f $$day ]; then \
			echo "Running $$day:"; \
			$$day; \
			echo ""; \
		fi; \
	done

# Individual day targets
day%: $(SOLUTIONDIR)/day%
	./$(SOLUTIONDIR)/day$*

.SUFFIXES: