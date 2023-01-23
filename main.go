package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/lzm-cli/gin-web-server-template/config"
	"github.com/lzm-cli/gin-web-server-template/durables"
	"github.com/lzm-cli/gin-web-server-template/jobs"
	"github.com/lzm-cli/gin-web-server-template/services"
)

func main() {
	service := flag.String("service", "http", "run a service")
	flag.Parse()

	// database := durable.NewDatabase()
	db := durables.NewDB()
	mixinClient := durables.GetMixinClient()
	log.Println(*service)

	switch *service {
	case "http":
		go func() {
			runtime.SetBlockProfileRate(1) // 开启对阻塞操作的跟踪
			_ = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.C.MonitorPort), nil)
		}()
		go jobs.StartWithHttpServiceJob()
		err := StartHTTP(db, mixinClient)
		if err != nil {
			log.Println(err)
		}
	default:
		hub := services.NewHub(db, mixinClient)
		err := hub.StartService(*service)
		if err != nil {
			log.Println(err)
		}
	}
}
