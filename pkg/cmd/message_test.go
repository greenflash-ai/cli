// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/greenflash-ai/greenflash-cli/internal/mocktest"
	"github.com/greenflash-ai/greenflash-cli/internal/requestflag"
)

func TestMessagesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "messages", "create",
			"--api-key", "string",
			"--external-user-id", "user-123",
			"--message", "{content: Hello!, context: context, createdAt: '2019-12-27T18:11:19.117Z', externalMessageId: user-msg-1, input: {foo: bar}, messageType: user_message, model: model, output: {foo: bar}, parentExternalMessageId: parentExternalMessageId, parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, properties: {foo: bar}, role: user, toolName: toolName}",
			"--message", "{content: Hi there! How can I help you?, context: context, createdAt: '2019-12-27T18:11:19.117Z', externalMessageId: assistant-msg-1, input: {foo: bar}, messageType: user_message, model: model, output: {foo: bar}, parentExternalMessageId: parentExternalMessageId, parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, properties: {foo: bar}, role: assistant, toolName: toolName}",
			"--message", "{content: Calling search tool, context: context, createdAt: '2019-12-27T18:11:19.117Z', externalMessageId: tool-call-1, input: {query: bar}, messageType: tool_call, model: model, output: {foo: bar}, parentExternalMessageId: parentExternalMessageId, parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, properties: {foo: bar}, role: user, toolName: web_search}",
			"--message", "{content: Search completed, context: context, createdAt: '2019-12-27T18:11:19.117Z', externalMessageId: tool-result-1, input: {foo: bar}, messageType: observation, model: model, output: {results: bar}, parentExternalMessageId: tool-call-1, parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, properties: {foo: bar}, role: user, toolName: toolName}",
			"--message", "{content: 'Based on the search, today will be sunny with a high of 75°F.', context: context, createdAt: '2019-12-27T18:11:19.117Z', externalMessageId: final-1, input: {foo: bar}, messageType: final_response, model: model, output: {foo: bar}, parentExternalMessageId: parentExternalMessageId, parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, properties: {foo: bar}, role: user, toolName: toolName}",
			"--conversation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--external-conversation-id", "conv-456",
			"--external-organization-id", "org-789",
			"--force-sample=true",
			"--model", "gpt-greenflash-1",
			"--product-id", "123e4567-e89b-12d3-a456-426614174001",
			"--properties", "{campaign: bar}",
			"--sample-rate", "0",
			"--system-prompt", "{components: [{content: You are a helpful assistant., componentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, externalComponentId: externalComponentId, isDynamic: true, name: name, source: customer, type: system}], content: x, externalPromptId: externalPromptId, promptId: 123e4567-e89b-12d3-a456-426614174004, variables: {foo: string}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(messagesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "messages", "create",
			"--api-key", "string",
			"--external-user-id", "user-123",
			"--message.content", "Hello!",
			"--message.context", "context",
			"--message.created-at", "2019-12-27T18:11:19.117Z",
			"--message.external-message-id", "user-msg-1",
			"--message.input", "{foo: bar}",
			"--message.message-type", "user_message",
			"--message.model", "model",
			"--message.output", "{foo: bar}",
			"--message.parent-external-message-id", "parentExternalMessageId",
			"--message.parent-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.properties", "{foo: bar}",
			"--message.role", "user",
			"--message.tool-name", "toolName",
			"--message.content", "Hi there! How can I help you?",
			"--message.context", "context",
			"--message.created-at", "2019-12-27T18:11:19.117Z",
			"--message.external-message-id", "assistant-msg-1",
			"--message.input", "{foo: bar}",
			"--message.message-type", "user_message",
			"--message.model", "model",
			"--message.output", "{foo: bar}",
			"--message.parent-external-message-id", "parentExternalMessageId",
			"--message.parent-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.properties", "{foo: bar}",
			"--message.role", "assistant",
			"--message.tool-name", "toolName",
			"--message.content", "Calling search tool",
			"--message.context", "context",
			"--message.created-at", "2019-12-27T18:11:19.117Z",
			"--message.external-message-id", "tool-call-1",
			"--message.input", "{query: bar}",
			"--message.message-type", "tool_call",
			"--message.model", "model",
			"--message.output", "{foo: bar}",
			"--message.parent-external-message-id", "parentExternalMessageId",
			"--message.parent-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.properties", "{foo: bar}",
			"--message.role", "user",
			"--message.tool-name", "web_search",
			"--message.content", "Search completed",
			"--message.context", "context",
			"--message.created-at", "2019-12-27T18:11:19.117Z",
			"--message.external-message-id", "tool-result-1",
			"--message.input", "{foo: bar}",
			"--message.message-type", "observation",
			"--message.model", "model",
			"--message.output", "{results: bar}",
			"--message.parent-external-message-id", "tool-call-1",
			"--message.parent-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.properties", "{foo: bar}",
			"--message.role", "user",
			"--message.tool-name", "toolName",
			"--message.content", "Based on the search, today will be sunny with a high of 75°F.",
			"--message.context", "context",
			"--message.created-at", "2019-12-27T18:11:19.117Z",
			"--message.external-message-id", "final-1",
			"--message.input", "{foo: bar}",
			"--message.message-type", "final_response",
			"--message.model", "model",
			"--message.output", "{foo: bar}",
			"--message.parent-external-message-id", "parentExternalMessageId",
			"--message.parent-message-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--message.properties", "{foo: bar}",
			"--message.role", "user",
			"--message.tool-name", "toolName",
			"--conversation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--external-conversation-id", "conv-456",
			"--external-organization-id", "org-789",
			"--force-sample=true",
			"--model", "gpt-greenflash-1",
			"--product-id", "123e4567-e89b-12d3-a456-426614174001",
			"--properties", "{campaign: bar}",
			"--sample-rate", "0",
			"--system-prompt", "{components: [{content: You are a helpful assistant., componentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, externalComponentId: externalComponentId, isDynamic: true, name: name, source: customer, type: system}], content: x, externalPromptId: externalPromptId, promptId: 123e4567-e89b-12d3-a456-426614174004, variables: {foo: string}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"externalUserId: user-123\n" +
			"messages:\n" +
			"  - content: Hello!\n" +
			"    context: context\n" +
			"    createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"    externalMessageId: user-msg-1\n" +
			"    input:\n" +
			"      foo: bar\n" +
			"    messageType: user_message\n" +
			"    model: model\n" +
			"    output:\n" +
			"      foo: bar\n" +
			"    parentExternalMessageId: parentExternalMessageId\n" +
			"    parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    properties:\n" +
			"      foo: bar\n" +
			"    role: user\n" +
			"    toolName: toolName\n" +
			"  - content: Hi there! How can I help you?\n" +
			"    context: context\n" +
			"    createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"    externalMessageId: assistant-msg-1\n" +
			"    input:\n" +
			"      foo: bar\n" +
			"    messageType: user_message\n" +
			"    model: model\n" +
			"    output:\n" +
			"      foo: bar\n" +
			"    parentExternalMessageId: parentExternalMessageId\n" +
			"    parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    properties:\n" +
			"      foo: bar\n" +
			"    role: assistant\n" +
			"    toolName: toolName\n" +
			"  - content: Calling search tool\n" +
			"    context: context\n" +
			"    createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"    externalMessageId: tool-call-1\n" +
			"    input:\n" +
			"      query: bar\n" +
			"    messageType: tool_call\n" +
			"    model: model\n" +
			"    output:\n" +
			"      foo: bar\n" +
			"    parentExternalMessageId: parentExternalMessageId\n" +
			"    parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    properties:\n" +
			"      foo: bar\n" +
			"    role: user\n" +
			"    toolName: web_search\n" +
			"  - content: Search completed\n" +
			"    context: context\n" +
			"    createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"    externalMessageId: tool-result-1\n" +
			"    input:\n" +
			"      foo: bar\n" +
			"    messageType: observation\n" +
			"    model: model\n" +
			"    output:\n" +
			"      results: bar\n" +
			"    parentExternalMessageId: tool-call-1\n" +
			"    parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    properties:\n" +
			"      foo: bar\n" +
			"    role: user\n" +
			"    toolName: toolName\n" +
			"  - content: Based on the search, today will be sunny with a high of 75°F.\n" +
			"    context: context\n" +
			"    createdAt: '2019-12-27T18:11:19.117Z'\n" +
			"    externalMessageId: final-1\n" +
			"    input:\n" +
			"      foo: bar\n" +
			"    messageType: final_response\n" +
			"    model: model\n" +
			"    output:\n" +
			"      foo: bar\n" +
			"    parentExternalMessageId: parentExternalMessageId\n" +
			"    parentMessageId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    properties:\n" +
			"      foo: bar\n" +
			"    role: user\n" +
			"    toolName: toolName\n" +
			"conversationId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"externalConversationId: conv-456\n" +
			"externalOrganizationId: org-789\n" +
			"forceSample: true\n" +
			"model: gpt-greenflash-1\n" +
			"productId: 123e4567-e89b-12d3-a456-426614174001\n" +
			"properties:\n" +
			"  campaign: bar\n" +
			"sampleRate: 0\n" +
			"systemPrompt:\n" +
			"  components:\n" +
			"    - content: You are a helpful assistant.\n" +
			"      componentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"      externalComponentId: externalComponentId\n" +
			"      isDynamic: true\n" +
			"      name: name\n" +
			"      source: customer\n" +
			"      type: system\n" +
			"  content: x\n" +
			"  externalPromptId: externalPromptId\n" +
			"  promptId: 123e4567-e89b-12d3-a456-426614174004\n" +
			"  variables:\n" +
			"    foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "messages", "create",
			"--api-key", "string",
		)
	})
}
