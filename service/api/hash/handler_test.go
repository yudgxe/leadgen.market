package hash

import (
	"fmt"
	"github/yudgxe/leadgen.market/common/handler"
	"github/yudgxe/leadgen.market/service/cache"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Hash(t *testing.T) {
	payload := `{
	"text": "%s",
	"alg" : "%s"
	}`

	cache.SetupMock()

	for _, test := range []struct {
		name         string
		body         string
		expectedCode int
	}{
		{
			name: "valid",
			body: fmt.Sprintf(payload, "hello", "sha256"),
			expectedCode: 200,
		},
		{
			name: "invalid_alg",
			body: fmt.Sprintf(payload, "hello", "invalid"),
			expectedCode: 400,
		},
		{
			name: "empty_text",
			body: fmt.Sprintf(payload, "", "sha256"),
			expectedCode: 400,
		},
		{
			name: "empty_alg",
			body: fmt.Sprintf(payload, "", "sha256"),
			expectedCode: 400,
		},
		{
			name: "empty_body",
			expectedCode: 400,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			wr, req := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/hash", strings.NewReader(test.body))

			handler.CreateHandler(Hash)(wr, req)

			if got, want := wr.Code, test.expectedCode; got != want {
				t.Fatal(cmp.Diff(got, want))
			}
		})
	}
}
