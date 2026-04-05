#include <stdint.h>
#include <stddef.h>

typedef struct {
    const uint8_t *data;
    size_t len;
} ZetaTxRaw;

int32_t zeta_omega_singularity_quantum(
    const ZetaTxRaw *batch_ptr,
    size_t count,
    uint32_t *success_out,
    uint64_t *gas_out,
    uint64_t *ns_out,
    uint64_t *swaps_out,
    uint64_t *nfts_out,
    uint64_t *mev_out,
    uint64_t *_dao_out
);
