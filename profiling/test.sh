#! /bin/sh
go build -o testSort
./testSort -cpuprofile cpu.prof -memprofile mem.prof
go tool pprof --http localhost:8080 cpu.prof 
