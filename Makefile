# Zeta-Omega Infinity: e2e Giga-Stack Verification
BINARY_NAME=libzeta_omega_infinity.a
GO_HARNESS=main.go

all: verify-symbols run-bench

verify-symbols:
	@echo "--- Checking Kernel Symbols ---"
	nm $(BINARY_NAME) | grep zeta_omega

run-bench:
	@echo "--- Running e2e Giga-Saturation Test ---"
	go run $(GO_HARNESS)

clean:
	go clean
	rm -f $(BINARY_NAME)
