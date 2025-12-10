package validation

import (
	"encoding/json"
	"fmt"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

// incoming messages should be checked against these schemas
type jsonSchemas struct {
	message  *jsonschema.Schema
	error    *jsonschema.Schema
	filter   *jsonschema.Schema
	sorting  *jsonschema.Schema
	query    *jsonschema.Schema
	mutation *jsonschema.Schema
	account  *jsonschema.Schema
}

// var messageSchema *jsonschema.Schema
var schemas *jsonSchemas

// WebSocketValidator handles WebSocket message validation
type WebSocketValidator struct {
	schemas *jsonSchemas
}

// NewWebSocketValidator creates a new validator instance
func NewWebSocketValidator() *WebSocketValidator {
	jsonCompiler := jsonschema.NewCompiler()
	jsonCompiler.DefaultDraft(jsonschema.Draft7)

	schemas = &jsonSchemas{}

	const schemaDir = "internal/models/jsonSchemas/"

	var err error
	schemas.message, err = jsonCompiler.Compile(schemaDir + "message.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	schemas.error, err = jsonCompiler.Compile(schemaDir + "error.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	schemas.filter, err = jsonCompiler.Compile(schemaDir + "filter.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	schemas.sorting, err = jsonCompiler.Compile(schemaDir + "sorting.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	schemas.query, err = jsonCompiler.Compile(schemaDir + "query.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	schemas.mutation, err = jsonCompiler.Compile(schemaDir + "mutation.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	schemas.account, err = jsonCompiler.Compile(schemaDir + "account.json")
	if err != nil {
		panic(fmt.Errorf("failed to compile schema: %w", err))
	}

	return &WebSocketValidator{
		schemas: schemas,
	}
}

func (v *WebSocketValidator) ValidateMessage(message []byte) error {
	var msgData interface{}
	if err := json.Unmarshal(message, &msgData); err != nil {
		panic(err)
	}

	return schemas.message.Validate(msgData)
}

func (v *WebSocketValidator) ValidateAccount(message []byte) error {
	var msgData interface{}
	if err := json.Unmarshal(message, &msgData); err != nil {
		panic(err)
	}

	return schemas.account.Validate(msgData)
}

func (v *WebSocketValidator) ValidateQuery(message []byte) error {
	var msgData interface{}
	if err := json.Unmarshal(message, &msgData); err != nil {
		panic(err)
	}

	ok := schemas.query.Validate(msgData)
	if ok != nil {
		fmt.Printf("ValidateQuery Error: %s\n", ok)
	} else {
		fmt.Printf("ValidateQuery Success\n")
	}
	return ok
}

func (v *WebSocketValidator) ValidateMutation(message []byte) error {
	var msgData interface{}
	if err := json.Unmarshal(message, &msgData); err != nil {
		panic(err)
	}

	ok := schemas.mutation.Validate(msgData)
	if ok != nil {
		fmt.Printf("ValidateMutation Error: %s\n", ok)
	} else {
		fmt.Printf("ValidateMutation Success\n")
	}
	return ok
}

func (v *WebSocketValidator) ValidateFilter(message []byte) error {
	var msgData interface{}
	if err := json.Unmarshal(message, &msgData); err != nil {
		panic(err)
	}

	ok := schemas.filter.Validate(msgData)
	if ok != nil {
		fmt.Printf("ValidateFilter Error: %s\n", ok)
	} else {
		fmt.Printf("ValidateFilter Success\n")
	}
	return ok
}

func (v *WebSocketValidator) ValidateSorting(message []byte) error {
	var msgData interface{}
	if err := json.Unmarshal(message, &msgData); err != nil {
		panic(err)
	}

	ok := schemas.sorting.Validate(msgData)
	if ok != nil {
		fmt.Printf("ValidateSorting Error: %s\n", ok)
	} else {
		fmt.Printf("ValidateSorting Success\n")
	}
	return ok
}
