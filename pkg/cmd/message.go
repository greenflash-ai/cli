// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/greenflash-ai/greenflash-cli/internal/apiquery"
	"github.com/greenflash-ai/greenflash-cli/internal/requestflag"
	"github.com/stainless-sdks/greenflash-public-api-go"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var messagesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Send us your AI conversations so we can analyze them for you. Works with\neverything from simple chatbots to complex agentic systems.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "external-user-id",
			Usage:    "Your external user ID that will be mapped to a user in our system.",
			Required: true,
			BodyPath: "externalUserId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			Usage:    "Array of conversation messages.",
			Required: true,
			BodyPath: "messages",
		},
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Usage:    "The Greenflash conversation ID. When provided, updates an existing conversation instead of creating a new one. Either conversationId, externalConversationId, productId must be provided.",
			BodyPath: "conversationId",
		},
		&requestflag.Flag[string]{
			Name:     "external-conversation-id",
			Usage:    "Your external identifier for the conversation. Either conversationId, externalConversationId, productId must be provided.",
			BodyPath: "externalConversationId",
		},
		&requestflag.Flag[string]{
			Name:     "external-organization-id",
			Usage:    "Your unique identifier for the organization this user belongs to. If provided, the user will be associated with this organization.",
			BodyPath: "externalOrganizationId",
		},
		&requestflag.Flag[bool]{
			Name:     "force-sample",
			Usage:    "When true, bypasses sampling and ensures this request is always ingested regardless of sampleRate. Use for critical conversations that must be captured.",
			BodyPath: "forceSample",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "The AI model used for the conversation.",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "The Greenflash product this conversation belongs to. Either conversationId, externalConversationId, productId must be provided.",
			BodyPath: "productId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Usage:    "Additional data about the conversation.",
			BodyPath: "properties",
		},
		&requestflag.Flag[float64]{
			Name:     "sample-rate",
			Usage:    "Controls the percentage of requests that are ingested (0.0 to 1.0). For example, 0.1 means 10% of requests will be stored. Defaults to 1.0 (all requests ingested). Sampling is deterministic based on conversation ID.",
			BodyPath: "sampleRate",
		},
		&requestflag.Flag[any]{
			Name:     "system-prompt",
			Usage:    "System prompt for the conversation. Can be a simple string or a prompt object with components.",
			BodyPath: "systemPrompt",
		},
	},
	Action:          handleMessagesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[string]{
			Name:       "message.content",
			Usage:      "The message content. Required for language-based analyses.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[any]{
			Name:       "message.context",
			Usage:      "Additional context (e.g., RAG data) used to generate the message.",
			InnerField: "context",
		},
		&requestflag.InnerFlag[any]{
			Name:       "message.created-at",
			Usage:      "When this message was created. If not provided, messages get sequential timestamps. Use for importing historical data.",
			InnerField: "createdAt",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.external-message-id",
			Usage:      "Your external identifier for this message. Used to reference the message in other API calls.",
			InnerField: "externalMessageId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.input",
			Usage:      "Structured input data for tool calls, retrievals, or other operations.",
			InnerField: "input",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.message-type",
			Usage:      "Detailed message type for agentic workflows. Cannot be used with role. Available types: user_message, assistant_message, system_message, thought, tool_call, observation, final_response, retrieval, memory_read, memory_write, chain_start, chain_end, embedding, tool_error, callback, llm, task, workflow",
			InnerField: "messageType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.model",
			Usage:      "The AI model used for this specific message. Use for multi-agent scenarios where different messages use different models. Overrides the conversation-level model for this message.",
			InnerField: "model",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.output",
			Usage:      "Structured output data from tool calls, retrievals, or other operations.",
			InnerField: "output",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.parent-external-message-id",
			Usage:      "The external ID of the parent message for threading. Cannot be used with parentMessageId.",
			InnerField: "parentExternalMessageId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.parent-message-id",
			Usage:      "The internal ID of the parent message for threading. Cannot be used with parentExternalMessageId.",
			InnerField: "parentMessageId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.properties",
			Usage:      "Custom message properties.",
			InnerField: "properties",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.role",
			Usage:      "Simple message role for basic chat: user, assistant, or system. Cannot be used with messageType.",
			InnerField: "role",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.tool-name",
			Usage:      "Name of the tool being called. Required for tool_call messages.",
			InnerField: "toolName",
		},
	},
})

func handleMessagesCreate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.MessageNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Messages.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messages create", obj, format, transform)
}
