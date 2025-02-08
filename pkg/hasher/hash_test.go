package hasher

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_GetHash(t *testing.T) {
	for _, test := range []struct {
		alg      HashAlg
		text     string
		expected string
		err      error
	}{
		{
			alg:      HashAlgMD5,
			text:     "MD5",
			expected: "7f138a09169b250e9dcb378140907378",
		},
		{
			alg:      HashAlgSHA256,
			text:     "SHA256",
			expected: "b3abe5d8c69b38733ad57ea75e83bcae42bbbbac75e3a5445862ed2f8a2cd677",
		},
		{
			alg:      "invalid",
			text:     "SHA256",
			expected: "b3abe5d8c69b38733ad57ea75e83bcae42bbbbac75e3a5445862ed2f8a2cd677",
			err: errors.New("unsupported hash alg"),
		},
	} {
		t.Run(string(test.alg), func(t *testing.T) {
			got, err := GetHash(test.alg, test.text)
			
			if err != nil && test.err != nil && err.Error() == test.err.Error() {
				return
			}

			if err != nil {
				t.Fatal(err)
			}

			if got, want := got, test.expected; got != want {
				t.Fatal(cmp.Diff(got, want))
			}
		})
	}
}
