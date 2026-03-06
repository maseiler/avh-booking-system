package ws

import (
	"encoding/json"
	"fmt"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

type messageValidator struct {
	schema *jsonschema.Schema
}

func newMessageValidator() *messageValidator {
	compiler := jsonschema.NewCompiler()
	compiler.DefaultDraft(jsonschema.Draft7)

	schema, err := compiler.Compile("../schemas/message.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	return &messageValidator{schema: schema}
}

func (v *messageValidator) validate(message []byte) error {
	var data interface{}
	if err := json.Unmarshal(message, &data); err != nil {
		panic(err)
	}
	return v.schema.Validate(data)
}
