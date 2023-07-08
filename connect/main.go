package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	greetv1 "github.com/ryutah/learn-golang/connect/gen/greet/v1"
	"github.com/ryutah/learn-golang/connect/gen/greet/v1/greetv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreetServer struct {
	greetv1connect.GreetServiceHandler
}

func (g *GreetServer) Greet(ctx context.Context, req *connect_go.Request[greetv1.GreetRequest]) (*connect_go.Response[greetv1.GreetResponse], error) {
	log.Println("request headres: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = http.ListenAndServe(":"+port, h2c.NewHandler(mux, &http2.Server{}))
}
