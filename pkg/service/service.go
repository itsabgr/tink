package service

import (
	"github.com/itsabgr/tink/pkg/storage"
	"github.com/itsabgr/tink/pkg/uid"
	"github.com/itsabgr/tink/pkg/validator"
	"github.com/valyala/fasthttp"
	"net"
	"time"
)

func Serve(st storage.Storage, listener net.Listener) error {
	server := &fasthttp.Server{}
	server.Handler = func(ctx *fasthttp.RequestCtx) {
		if ctx.IsGet() {
			dest, err := st.GetByKey(ctx.Path()[1:])
			if err != nil {
				ctx.Error(err.Error(), fasthttp.StatusNotFound)
				return
			}
			ctx.SetStatusCode(fasthttp.StatusTemporaryRedirect)
			ctx.Response.Header.Set("Location", string(dest))
			return
		}
		uri := ctx.PostBody()
		err := validator.ValidateURI(uri)
		if err != nil {
			server.ErrorHandler(ctx, err)
			return
		}
		name := ctx.Path()[1:]
		if len(name) == 0 {
			name = []byte(uid.Gen().String())
		}
		err = st.Add(name, uri)
		if err != nil {
			server.ErrorHandler(ctx, err)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(name)
	}
	server.ErrorHandler = func(ctx *fasthttp.RequestCtx, err error) {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	}
	server.MaxRequestBodySize = 2024
	server.WriteTimeout = 2 * time.Second
	server.ReadTimeout = 2 * time.Second
	server.IdleTimeout = 2 * time.Second
	server.NoDefaultContentType = true
	server.DisableHeaderNamesNormalizing = false
	server.DisablePreParseMultipartForm = true
	server.Name = "tink/0.1"
	server.TCPKeepalivePeriod = time.Second
	server.TCPKeepalive = true
	server.SleepWhenConcurrencyLimitsExceeded = time.Second * 5
	server.SecureErrorLogMessage = true
	return server.Serve(listener)
}
