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

var eventsCreate = cli.Command{
	Name:    "create",
	Usage:   "Track timestamped events representing user or organization actions. Events are\nused to track important business outcomes (signups, conversions, upgrades,\ncancellations, etc.) and integrate them into Greenflash's quality metrics. Each\nevent can be optionally linked to a conversation, user, and organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "event-type",
			Usage:    `The specific name or category of the event being tracked (e.g., "trial_started", "signup", "feature_usage"). This helps categorize events for analysis and often pairs with "value" to define the outcome.`,
			Required: true,
			BodyPath: "eventType",
		},
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "The unique identifier of the Greenflash product associated with this event. This links the event to a specific product context.",
			Required: true,
			BodyPath: "productId",
		},
		&requestflag.Flag[string]{
			Name:     "value",
			Usage:    `The specific value associated with the event (e.g., "99.00", "5", "premium_plan"). This pairs with "valueType" and "eventType" to define the magnitude or content of the event.`,
			Required: true,
			BodyPath: "value",
		},
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Usage:    "The unique Greenflash identifier for the conversation. Links the event to a specific chat session in Greenflash.",
			BodyPath: "conversationId",
		},
		&requestflag.Flag[any]{
			Name:     "event-at",
			Usage:    "The ISO 8601 timestamp of when the event actually occurred. Defaults to the current time if not provided. Useful for backdating historical events.",
			BodyPath: "eventAt",
		},
		&requestflag.Flag[string]{
			Name:     "external-conversation-id",
			Usage:    "Your system's unique identifier for the conversation or thread where this event occurred.",
			BodyPath: "externalConversationId",
		},
		&requestflag.Flag[string]{
			Name:     "external-organization-id",
			Usage:    "Your system's unique identifier for the organization associated with this event. Used to map events to your customer accounts.",
			BodyPath: "externalOrganizationId",
		},
		&requestflag.Flag[string]{
			Name:     "external-user-id",
			Usage:    "Your system's unique identifier for the user associated with this event. Used to map Greenflash events back to your user records.",
			BodyPath: "externalUserId",
		},
		&requestflag.Flag[bool]{
			Name:     "force-sample",
			Usage:    "When true, bypasses sampling and ensures this event is always ingested regardless of sampleRate. Use for critical events that must be captured.",
			BodyPath: "forceSample",
		},
		&requestflag.Flag[string]{
			Name:     "influence",
			Usage:    `A high-level categorization of how this event generally "changed things" or influenced quality (positive, negative, or neutral). Use this for broad classification of outcomes.`,
			Default:  "neutral",
			BodyPath: "influence",
		},
		&requestflag.Flag[string]{
			Name:     "insert-id",
			Usage:    "A unique key for idempotency. If you retry a request with the same insertId, it prevents creating a duplicate event record.",
			BodyPath: "insertId",
		},
		&requestflag.Flag[string]{
			Name:     "organization-id",
			Usage:    "The unique Greenflash identifier for the organization. Provide this if you have the Greenflash Organization ID.",
			BodyPath: "organizationId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Usage:    `A key-value object for storing additional, unstructured context about the event (e.g., { source: "web_app", campaign_id: "123" }). Useful for custom filtering.`,
			BodyPath: "properties",
		},
		&requestflag.Flag[float64]{
			Name:     "quality-impact-score",
			Usage:    `A precise numeric score between -1.0 and 1.0 for direct control over the quality impact. If omitted, it is automatically derived from the "influence" field.`,
			Default:  0,
			BodyPath: "qualityImpactScore",
		},
		&requestflag.Flag[float64]{
			Name:     "sample-rate",
			Usage:    "Controls the percentage of requests that are ingested (0.0 to 1.0). For example, 0.1 means 10% of events will be stored. Defaults to 1.0 (all events ingested). Sampling is deterministic based on event type and organization.",
			BodyPath: "sampleRate",
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    `The unique Greenflash identifier for the user. Provide this if you already have the Greenflash User ID; otherwise, use "externalUserId".`,
			BodyPath: "userId",
		},
		&requestflag.Flag[string]{
			Name:     "value-type",
			Usage:    `Defines the format of the "value" field (currency, numeric, or text). This ensures the value is interpreted and processed correctly.`,
			Default:  "text",
			BodyPath: "valueType",
		},
	},
	Action:          handleEventsCreate,
	HideHelpCommand: true,
}

func handleEventsCreate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.EventNewParams{}

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
	_, err = client.Events.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "events create", obj, format, transform)
}
