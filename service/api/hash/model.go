package hash

import (
	"errors"
	"github/yudgxe/leadgen.market/pkg/hasher"
)

type hashReqBody struct {
	Text string         `json:"text"`
	Alg  hasher.HashAlg `json:"alg"`
}

func (r hashReqBody) validate() error {
	if len(r.Text) == 0 {
		return errors.New("field text can't be empty")
	}

	if len(r.Alg) == 0 {
		return errors.New("field alg can't be empty")
	}

	if err := hasher.IsValidAlg(r.Alg); err != nil {
		return err
	}

	return nil
}
