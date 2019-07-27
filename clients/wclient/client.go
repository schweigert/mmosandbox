package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/imroc/req"
	"github.com/schweigert/mmosandbox/clients/client"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/domain/inputs"
	"github.com/schweigert/mmosandbox/domain/outputs"
	"github.com/schweigert/mmosandbox/infra/routes"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// CreateAccountFlow implementation for client.CreateAccountFlow
type CreateAccountFlow struct {
}

// NewCreateAccountFlow constructor
func NewCreateAccountFlow() client.CreateAccountFlow {
	return &CreateAccountFlow{}
}

// CreateAccountOperation for willson flow
func (flow *CreateAccountFlow) CreateAccountOperation(in *inputs.CreateAccountInput) (*entities.Account, bool) {
	rec := req.New()
	rec.SetTimeout(5 * time.Second)

	resp, err := rec.Put(fmt.Sprintf("%s/%s", config.Client().Addr(), routes.Account), req.BodyJSON(in))
	dont.Panic(err)

	out := outputs.NewCreateAccountOutput()

	if resp.Response().StatusCode != http.StatusCreated {
		return nil, false
	}

	err = resp.ToJSON(out)
	dont.Panic(err)

	fmt.Println(out.Account)

	return out.Account, out.Success
}

func main() {
	client.UsedCreateAccountFlow = NewCreateAccountFlow()

	client.BotFlow()
}
