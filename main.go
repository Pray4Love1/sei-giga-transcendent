package main

/*
#cgo LDFLAGS: -L./target/release -lzeta_omega_infinity
#include "zeta_omega.h"
#include <stdlib.h>
*/
import "C"
import (
        "fmt"
        "time"
        "unsafe"
)

func main() {
        count := 1000000
        fmt.Printf("--- Saturating Rust Kernel from Go: %d TXs ---\n", count)

        // 1. Allocate a single block of C memory for the transaction metadata
        // This prevents the "unpinned Go pointer" panic.
        size := C.size_t(count) * C.size_t(unsafe.Sizeof(C.ZetaTxRaw{}))
        pTxs := (*C.ZetaTxRaw)(C.malloc(size))
        defer C.free(unsafe.Pointer(pTxs))

        // 2. Allocate the raw data in Go memory, but pass direct pointers
        data := make([]byte, count*64)
        
        // Convert C pointer to a Go slice header so we can iterate easily
        txsSlice := (*[1 << 28]C.ZetaTxRaw)(unsafe.Pointer(pTxs))[:count:count]

        for i := 0; i < count; i++ {
                txsSlice[i] = C.ZetaTxRaw{
                        data: (*C.uint8_t)(unsafe.Pointer(&data[i*64])),
                        len:  C.size_t(64),
                }
        }

        var success C.uint32_t
        var gas, ns, swaps, nfts, mev, dao C.uint64_t

        // 3. Execute the Bridge
        start := time.Now()
        C.zeta_omega_singularity_quantum(
                pTxs,
                C.size_t(count),
                &success,
                &gas,
                &ns,
                &swaps,
                &nfts,
                &mev,
                &dao,
        )
        elapsed := time.Since(start)

        tps := float64(count) / elapsed.Seconds()
        fmt.Printf("Bridge Throughput: %.2f tx/sec\n", tps)
        fmt.Printf("Kernel Internal Latency: %d ns\n", ns)
        fmt.Printf("Total Elapsed (Bridge + Kernel): %v\n", elapsed)
}
