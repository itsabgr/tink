package validator

import "github.com/valyala/fasthttp"

func ValidateURI(uri []byte) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURIBytes(uri)
	resp.SkipBody = true
	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return ErrNonOKResp
	}
	return nil
}
