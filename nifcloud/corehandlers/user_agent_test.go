package corehandlers

import (
	"net/http"
	"os"
	"testing"

	"github.com/shztki/nifcloud-sdk-go/nifcloud/request"
)

func TestAddHostExecEnvUserAgentHander(t *testing.T) {
	cases := []struct {
		ExecEnv string
		Expect  string
	}{
		{ExecEnv: "Lambda", Expect: "exec_env/Lambda"},
		{ExecEnv: "", Expect: ""},
		{ExecEnv: "someThingCool", Expect: "exec_env/someThingCool"},
	}

	for i, c := range cases {
		os.Clearenv()
		os.Setenv(execEnvVar, c.ExecEnv)

		req := &request.Request{
			HTTPRequest: &http.Request{
				Header: http.Header{},
			},
		}
		AddHostExecEnvUserAgentHander.Fn(req)

		if err := req.Error; err != nil {
			t.Fatalf("%d, expect no error, got %v", i, err)
		}

		if e, a := c.Expect, req.HTTPRequest.Header.Get("User-Agent"); e != a {
			t.Errorf("%d, expect %v user agent, got %v", i, e, a)
		}
	}
}
