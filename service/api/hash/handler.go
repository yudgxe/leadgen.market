package hash

import (
	"context"
	"encoding/json"
	"fmt"

	"github/yudgxe/leadgen.market/common/handler"
	"github/yudgxe/leadgen.market/pkg/hasher"
	"github/yudgxe/leadgen.market/service/cache"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func Bind(mux *chi.Mux) {
	mux.Post("/", handler.CreateHandler(Hash))
}

func Hash(s handler.CoreService) (any, error) {
	var body hashReqBody
	if err := json.NewDecoder(s.R.Body).Decode(&body); err != nil { // todo: wrap parse in core service
		return nil, handler.NewHttpErrorBadRequest(err.Error())
	}

	if err := body.validate(); err != nil {
		return nil, handler.NewHttpErrorBadRequest(err.Error())
	}

	find, result, err := cache.GetHashService().Get(context.TODO(), fmt.Sprintf("%s:%s", body.Alg, body.Text))
	if err != nil {
		log.Errorf("error on get from cache - %v", err)
	}

	if find {
		log.Debugf("get result from cache")
		return map[string]string{"text": result}, nil
	}

	hash, err := hasher.GetHash(body.Alg, body.Text)
	if err != nil {
		return nil, err
	}

	log.Debugf("get result from hasher")

	if err := cache.GetHashService().Save(context.TODO(), fmt.Sprintf("%s:%s", body.Alg, body.Text), hash); err != nil {
		log.Errorf("error on save in cache")
	}

	return map[string]string{"text": hash}, nil
}
