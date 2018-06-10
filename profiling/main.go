package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	sort "sort_algorithm/sort"
	"time"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := r.Perm(1000000)
	sort.MergeSort(l)
	if sort.IsSorted(l) != false {
		fmt.Println("array is sorted")
	} else {
		fmt.Println("failed")
	}
}
