package main

import (
	"flag"
	"github.com/itsabgr/tink/pkg/service"
	"github.com/itsabgr/tink/pkg/storage"
	"log"
	"net"
)

var flagAddr = flag.String("addr", "localhost:8080", "server address")
var flagDir = flag.String("dir", "", "data directory path")

func main() {
	st, err := storage.Open(*flagDir)
	if err != nil {
		log.Fatalln(err)
	}
	defer st.Close()
	ln, err := net.Listen("tcp", *flagAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()
	log.Println("listen", ln.Addr())
	err = service.Serve(st, ln)
	if err != nil {
		log.Fatalln(err)
	}
}
