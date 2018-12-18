package service

import (
	"net/http"
	"os"

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

	http.Handle("/token", TokenHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}
