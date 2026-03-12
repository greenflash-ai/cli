// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/greenflash-ai/cli/internal/apiquery"
	"github.com/greenflash-ai/cli/internal/requestflag"
	"github.com/greenflash-ai/go"
	"github.com/greenflash-ai/go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var interactionsList = cli.Command{
	Name:    "list",
	Usage:   "Browse through all conversations in your workspace to understand how users are\ninteracting with your AI. Filter by product or version to focus on specific\nareas of your application.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return.",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "offset",
			Usage:     "Offset for pagination.",
			QueryPath: "offset",
		},
		&requestflag.Flag[float64]{
			Name:      "page",
			Usage:     "Page number",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "product-id",
			Usage:     "Filter interactions by product ID.",
			QueryPath: "productId",
		},
		&requestflag.Flag[string]{
			Name:      "version-id",
			Usage:     "Filter interactions by version ID.",
			QueryPath: "versionId",
		},
	},
	Action:          handleInteractionsList,
	HideHelpCommand: true,
}

var interactionsGetInteractionAnalytics = cli.Command{
	Name:    "get-interaction-analytics",
	Usage:   "Understand what happened in a specific conversation with AI-powered analysis.\nSee sentiment shifts, detect frustration, identify commercial intent, and get\nactionable insights.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "interaction-id",
			Usage:    "The interaction ID to get analytics for",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "mode",
			Usage:     `Analysis mode: "simple" returns only numeric aggregates (no rate limiting), "insights" includes topics, keywords, and recommendations (rate limited per tenant plan).`,
			Default:   "insights",
			QueryPath: "mode",
		},
	},
	Action:          handleInteractionsGetInteractionAnalytics,
	HideHelpCommand: true,
}

func handleInteractionsList(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.InteractionListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Interactions.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "interactions list", obj, format, transform)
}

func handleInteractionsGetInteractionAnalytics(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("interaction-id") && len(unusedArgs) > 0 {
		cmd.Set("interaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.InteractionGetInteractionAnalyticsParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Interactions.GetInteractionAnalytics(
		ctx,
		cmd.Value("interaction-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "interactions get-interaction-analytics", obj, format, transform)
}
