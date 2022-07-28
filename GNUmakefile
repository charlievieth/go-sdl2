# Directories
SRCDIR=		${CURDIR}
BINDIR=     ${SRCDIR}/bin
ALLDIRS!=	find ${SRCDIR} -type d -not -path "./.git*"
# Packages
ROOTPKG=	$(subst ${GOPATH}/src/,,${SRCDIR})
PACKAGES=	sdl img mix ttf gfx
# Some cleanups
CLEANUP=	*.cache *.core *~ *.orig

GO?=		go

all: clean packages

clean:
	@echo "Cleaning up..."
	@for dir in ${ALLDIRS}; do \
		cd $$dir && rm -f ${CLEANUP} && cd ${SRCDIR}; \
	done
	@rm -rf ${BINDIR}

packages:
	@for pkg in ${PACKAGES}; do \
		echo "Building package ${ROOTPKG}/$$pkg..."; \
		${GO} build ${ROOTPKG}/$$pkg; \
	done

test:
	@for pkg in ${PACKAGES}; do \
		echo "Testing ${ROOTPKG}/$$pkg..."; \
		${GO} test -v ${ROOTPKG}/$$pkg; \
	done

coverage:
	@for pkg in ${PACKAGES}; do \
		echo "Coverage for ${ROOTPKG}/$$pkg..."; \
		${GO} test -cover ${ROOTPKG}/$$pkg; \
	done

bench:
	@for pkg in ${PACKAGES}; do \
		echo "Benchmarking ${ROOTPKG}/$$pkg..."; \
		${GO} test -bench=. -cpu 1,2,3,4 ${ROOTPKG}/$$pkg; \
	done

bin/stringer:
	@mkdir -p ${BINDIR}
	@${GO} build -o ${BINDIR}/stringer golang.org/x/tools/cmd/stringer

.PHONY: generate
generate: bin/stringer
	@for pkg in ${PACKAGES}; do \
		echo "Generating ${ROOTPKG}/$$pkg..."; \
		export PATH=${BINDIR}:$$PATH; \
		${GO} generate ${ROOTPKG}/$$pkg; \
	done
