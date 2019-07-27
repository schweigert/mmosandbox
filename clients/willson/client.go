package main

import (
	"github.com/schweigert/mmosandbox/clients/client"
	"github.com/schweigert/mmosandbox/domain/inputs"
)

// CreateAccountFlow implementation for client.CreateAccountFlow
type CreateAccountFlow struct {
}

// NewCreateAccountFlow constructor
func NewCreateAccountFlow() client.CreateAccountFlow {
	return &CreateAccountFlow{}
}

// CreateAccountOperation for willson flow
func (flow *CreateAccountFlow) CreateAccountOperation(in *inputs.CreateAccountInput) bool {
	return true
}

func main() {
	client.UsedCreateAccountFlow = NewCreateAccountFlow()

	client.BotFlow()
}
