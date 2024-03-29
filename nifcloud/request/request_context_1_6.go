// +build !go1.7

package request

import "github.com/shztki/nifcloud-sdk-go/nifcloud"

// setContext updates the Request to use the passed in context for cancellation.
// Context will also be used for request retry delay.
//
// Creates shallow copy of the http.Request with the WithContext method.
func setRequestContext(r *Request, ctx nifcloud.Context) {
	r.context = ctx
	r.HTTPRequest.Cancel = ctx.Done()
}
