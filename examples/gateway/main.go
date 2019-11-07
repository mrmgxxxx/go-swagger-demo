package main

import (
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "go-swagger-demo/examples/protocol/echoserver"
)

const (
	// gateway endpoint.
	endpoint = ":8081"

	// echo server endpoint.
	echoEndpoint = "localhost:9527"
)

// Gateway is gRPC gateway.
type Gateway struct {
	// gRPC gateway server mux.
	gwmux *runtime.ServeMux

	// http server mux.
	mux *http.ServeMux

	// echo server client.
	echoSvrConn *grpc.ClientConn
	echoSvrCli  pb.EchoClient
}

// NewGateway creates a new gateway instance.
func NewGateway() *Gateway {
	return &Gateway{}
}

// create a mux for gateway server.
func (gw *Gateway) initGWMux() {
	opt := runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&runtime.JSONPb{EnumsAsInts: true, EmitDefaults: true, OrigName: true})
	gw.gwmux = runtime.NewServeMux(opt)
}

// create a mux for http server.
func (gw *Gateway) initSvrMux() {
	gw.mux = http.NewServeMux()

	gw.mux.Handle("/", gw.gwmux)
	gw.mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join("swagger", strings.TrimPrefix(r.URL.Path, "/swagger/")))
	})
}

// initialize  echo server gRPC client.
func (gw *Gateway) initEchoServerClient() {
	conn, err := grpc.Dial(echoEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial: %v", err)
	}
	gw.echoSvrConn = conn
	gw.echoSvrCli = pb.NewEchoClient(conn)
}

// initMods initialize the server modules.
func (gw *Gateway) initMods() {
	// initialize echo server gRPC client.
	gw.initEchoServerClient()

	// initialize gateway mux.
	gw.initGWMux()

	// initialize server mux.
	gw.initSvrMux()
}

// Run runs gateway.
func (gw *Gateway) Run() {
	// initialize server modules.
	gw.initMods()

	// register echo server handler.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := pb.RegisterEchoHandlerClient(ctx, gw.gwmux, gw.echoSvrCli); err != nil {
		log.Fatal("gateway register echo server handler, %+v", err)
	}

	// gateway service listen and serve.
	if err := http.ListenAndServe(endpoint, gw.mux); err != nil {
		log.Fatal("listen and serve, %+v", err)
	}
}

func main() {
	// new gateway instance.
	gw := NewGateway()

	// run.
	gw.Run()
}
