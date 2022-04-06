package tink

import (
	"flag"
	"log"
	"net"
	"tink/pkg/service"
	"tink/pkg/storage"
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
	err = service.Serve(st, ln)
	if err != nil {
		log.Fatalln(err)
	}
}
