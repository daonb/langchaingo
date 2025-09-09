package tools

import "context"

// Tool is a tool for the llm agent to interact with different applications.
type Tool interface {
	Name() string
	Description() string
	Call(ctx context.Context, input string) (string, error)
}

// SchematicTool is an optional interface for tools that can provide their own
// JSON schema for function calling. This enables proper multi-parameter tool
// support instead of falling back to the __arg1 workaround.
type SchematicTool interface {
	Tool
	// GetSchema returns the JSON schema for the tool's input parameters.
	// The schema should be a valid JSON Schema object with properties, required fields, etc.
	// Example: {"type": "object", "properties": {"path": {"type": "string"}}, "required": ["path"]}
	GetSchema() map[string]any
}
