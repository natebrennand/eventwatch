package admin

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
)

// AdminPortString is the string respresentation of the port that the /ping endpoint will listen on.
var AdminPortString string

// init run when the library is imported and adds the endpoints as well as starts the server.
func init() {
	AdminPortString = os.Getenv("ADMIN_PORT")
	if AdminPortString == "" {
		log.Print("ADMIN_PORT not set so listening on 8001")
		AdminPortString = "8001"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", healthcheck)
	mux.HandleFunc("/stats", memoryStats)

	go func() {
		http.ListenAndServe(":"+AdminPortString, mux)
	}()
}

// healthcheck is a simple endpoint that responds "ping" for checking that the service is
// indeed alive and responsive.
func healthcheck(resp http.ResponseWriter, req *http.Request) {
	n, err := resp.Write([]byte(`pong`))
	if n != 4 || err != nil {
		log.Print("Admin: Failed to properly respond to healthcheck")
		return
	}
	log.Print("Responded to healthcheck ping")
}

// memoryStats gets the runtime information and returns a JSON
func memoryStats(resp http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.Encode(buildProfile())

	n, err := resp.Write(buf.Bytes())
	if n == 0 || err != nil { // err if failed to write
		resp.WriteHeader(http.StatusInternalServerError)
	}
	log.Print("Responded to healthcheck ping")
}

// MemoryProfile contains data on the memory allocations, heap and stack of the application.
// It is returned from the /stats endpoint.
type MemoryProfile struct {
	// General statistics.
	Allocation      uint64 // bytes allocated and still in use
	TotalAllocation uint64 // bytes allocated (even if freed)
	SystemBytes     uint64 // bytes obtained from system (sum of XxxSys below)
	Lookups         uint64 // number of pointer lookups
	Mallocs         uint64 // number of mallocs
	Frees           uint64 // number of frees

	// Main allocation heap statistics.
	HeapAllocation uint64 // bytes allocated and still in use
	HeapFromSystem uint64 // bytes obtained from system
	HeapIdle       uint64 // bytes in idle spans
	HeapInUse      uint64 // bytes in non-idle span
	HeapReleased   uint64 // bytes released to the OS
	HeapObjects    uint64 // total number of allocated objects

	// Low-level fixed-size structure allocator statistics.
	// Inuse is bytes used now.
	// Sys is bytes obtained from system.
	StackInUse     uint64 // bootstrap stacks
	StackSystem    uint64
	MSpanInUse     uint64 // mspan structures
	MSpanSystem    uint64
	MCacheInUse    uint64 // mcache structures
	MCacheSystem   uint64
	BuckHashSystem uint64 // profiling bucket hash table
	OtherSystem    uint64 // other system allocations
}

// buildProfile generates a MemoryProfile from Memstats in the runtime package.
func buildProfile() MemoryProfile {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return MemoryProfile{
		Allocation:      stats.Alloc,
		TotalAllocation: stats.TotalAlloc,
		SystemBytes:     stats.Sys,
		Lookups:         stats.Lookups,
		Mallocs:         stats.Mallocs,
		Frees:           stats.Frees,
		HeapAllocation:  stats.HeapAlloc,
		HeapFromSystem:  stats.HeapSys,
		HeapIdle:        stats.HeapIdle,
		HeapInUse:       stats.HeapInuse,
		HeapReleased:    stats.HeapReleased,
		HeapObjects:     stats.HeapObjects,
		StackInUse:      stats.StackInuse,
		StackSystem:     stats.StackSys,
		MSpanInUse:      stats.MSpanInuse,
		MSpanSystem:     stats.MSpanSys,
		MCacheInUse:     stats.MCacheInuse,
		MCacheSystem:    stats.MCacheSys,
		BuckHashSystem:  stats.BuckHashSys,
		OtherSystem:     stats.OtherSys,
	}
}

// garbageCollect triggers a garbage collection
func garbageCollect(resp http.ResponseWriter, req *http.Request) {
	runtime.GC()
}
