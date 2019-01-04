package service

import (
	//	"fmt"
	//	ilog "log"
	"net/http"
	"os"
	//	"os/signal"
	//	"syscall"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/oauth/v2/service/request"
)

func Run() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := request.TokenService{}

	TokenHandler := httptransport.NewServer(
		request.MakeGetTokenEndpoint(svc),
		request.DecodeGetTokenRequest,
		request.EncodeResponse,
	)

	password := request.PassWordTokenService{}

	PassWordHandler := httptransport.NewServer(
		request.MakeGetPassWordTokenEndpoint(password),
		request.DecodeGetPassWordTokenRequest,
		request.EncodeResponse,
	)

	module := request.OauthModuleService{}

	ModuleHandler := httptransport.NewServer(
		request.MakeGetModuleEndpoint(module),
		request.DecodeGetModuleRequest,
		request.EncodeResponse,
	)

	authorize := request.AuthorizationTokenService{}

	AuthorizeHandler := httptransport.NewServer(
		request.MakeGetAuthorizationTokenEndpoint(authorize),
		request.DecodeGetAuthorizationTokenRequest,
		request.EncodeResponse,
	)

	http.HandleFunc("/authorize", request.AuthorizeHandler)
	http.HandleFunc("/login", request.LoginHandler)
	http.HandleFunc("/auth", request.AuthHandler)

	http.Handle("/access_token", AuthorizeHandler)
	http.Handle("/token", TokenHandler)
	http.Handle("/password", PassWordHandler)
	http.Handle("/module", ModuleHandler)

	//	errChan := make(chan error)

	//	go func() {
	//		c := make(chan os.Signal, 1)
	//		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	//		errChan <- fmt.Errorf("%s", <-c)
	//	}()
	//	error := <-errChan

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}
