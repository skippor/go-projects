package main

import (
	"fmt"
	log "github.com/zeromicro/go-zero/core/logx"
)

ctx     context.Context
wg      *sync.WaitGroup

func test() {
    ticker := time.NewTicker(_tcpCollectInterval * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-t.ctx.Done():
            log.Infof("TcpRoutine: received cancel signal, exiting...")
            return
        case <-ticker.C:
            log.Infof("TcpRoutine: received cancel signal, exiting...")
            return
        }
    }
}

func main() {
    runtime.GOMAXPROCS(_macProc)

    ctx, cancel := context.WithCancel(context.Background())
    wg := &sync.WaitGroup{}

    tcp.Init(ctx, wg)
}