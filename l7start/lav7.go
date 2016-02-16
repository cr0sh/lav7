package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"

	"github.com/L7-MCPE/lav7"
	"github.com/L7-MCPE/lav7/command"
	"github.com/L7-MCPE/lav7/level/format/dummy"
	"github.com/L7-MCPE/lav7/level/gen"
	"github.com/L7-MCPE/lav7/raknet"
)

func main() {
	go http.ListenAndServe(":8080", nil)
	runtime.GOMAXPROCS(runtime.NumCPU())
	p := new(dummy.Provider)
	p.Init(new(gen.SampleGenerator).Gen, lav7.GetDefaultLevel().Name)
	lav7.GetDefaultLevel().Init(p)
	log.Println("Generating chunks")
	wg := new(sync.WaitGroup)
	wg.Add(25)
	start := time.Now()
	for x := int32(-2); x <= 2; x++ {
		for z := int32(-2); z <= 2; z++ {
			go func(x, z int32) { lav7.GetDefaultLevel().GetChunk(x, z, true); wg.Done() }(x, z)
		}
	}
	wg.Wait()
	log.Println("Elapsed time:", time.Since(start).Seconds(), "seconds")
	var r *raknet.Router
	var err error
	if r, err = raknet.CreateRouter(lav7.RegisterPlayer, lav7.UnregisterPlayer, 19132); err != nil {
		fmt.Println(err)
		return
	}
	r.Start()
	command.HandleCommand()
}
