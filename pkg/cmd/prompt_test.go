// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/greenflash-ai/cli/internal/mocktest"
	"github.com/greenflash-ai/cli/internal/requestflag"
)

func TestPromptsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "prompts", "create",
			"--api-key", "string",
			"--component", "{content: 'You are a helpful assistant for {{productName}}. Greet {{userName}} warmly.', componentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, externalComponentId: externalComponentId, isDynamic: false, name: Base Instructions, source: customer, type: system}",
			"--name", "Customer Support Prompt",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--role", "x",
			"--description", "Standard customer support  prompt",
			"--external-prompt-id", "support-v1",
			"--source", "customer",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(promptsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "prompts", "create",
			"--api-key", "string",
			"--component.content", "You are a helpful assistant for {{productName}}. Greet {{userName}} warmly.",
			"--component.component-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--component.external-component-id", "externalComponentId",
			"--component.is-dynamic=false",
			"--component.name", "Base Instructions",
			"--component.source", "customer",
			"--component.type", "system",
			"--name", "Customer Support Prompt",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--role", "x",
			"--description", "Standard customer support  prompt",
			"--external-prompt-id", "support-v1",
			"--source", "customer",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"components:\n" +
			"  - content: >-\n" +
			"      You are a helpful assistant for {{productName}}. Greet {{userName}}\n" +
			"      warmly.\n" +
			"    componentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    externalComponentId: externalComponentId\n" +
			"    isDynamic: false\n" +
			"    name: Base Instructions\n" +
			"    source: customer\n" +
			"    type: system\n" +
			"name: Customer Support Prompt\n" +
			"productId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"role: x\n" +
			"description: Standard customer support  prompt\n" +
			"externalPromptId: support-v1\n" +
			"source: customer\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "prompts", "create",
			"--api-key", "string",
		)
	})
}

func TestPromptsUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "prompts", "update",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--component", "{content: 'You are a helpful assistant for {{productName}}. Always be polite to {{userName}}.', componentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, externalComponentId: externalComponentId, isDynamic: true, name: Base Instructions V2, source: customer, type: system}",
			"--description", "Updated description",
			"--name", "Updated Customer Support Prompt",
			"--role", "role",
			"--source", "customer",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(promptsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "prompts", "update",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--component.content", "You are a helpful assistant for {{productName}}. Always be polite to {{userName}}.",
			"--component.component-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--component.external-component-id", "externalComponentId",
			"--component.is-dynamic=true",
			"--component.name", "Base Instructions V2",
			"--component.source", "customer",
			"--component.type", "system",
			"--description", "Updated description",
			"--name", "Updated Customer Support Prompt",
			"--role", "role",
			"--source", "customer",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"components:\n" +
			"  - content: >-\n" +
			"      You are a helpful assistant for {{productName}}. Always be polite to\n" +
			"      {{userName}}.\n" +
			"    componentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    externalComponentId: externalComponentId\n" +
			"    isDynamic: true\n" +
			"    name: Base Instructions V2\n" +
			"    source: customer\n" +
			"    type: system\n" +
			"description: Updated description\n" +
			"name: Updated Customer Support Prompt\n" +
			"role: role\n" +
			"source: customer\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "prompts", "update",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPromptsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "prompts", "list",
			"--api-key", "string",
			"--active-only=true",
			"--include-archived=true",
			"--limit", "100",
			"--page", "1",
			"--page-size", "1",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPromptsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "prompts", "delete",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPromptsGet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "prompts", "get",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
