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

var ratingsLog = cli.Command{
	Name:    "log",
	Usage:   "Record user feedback and ratings for conversations or individual messages.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "The Greenflash product ID to rate.",
			Required: true,
			BodyPath: "productId",
		},
		&requestflag.Flag[float64]{
			Name:     "rating",
			Usage:    "The rating value. Must be between ratingMin and ratingMax (inclusive).",
			Required: true,
			BodyPath: "rating",
		},
		&requestflag.Flag[float64]{
			Name:     "rating-max",
			Usage:    "The maximum possible rating value (e.g., 5 for a 1-5 scale).",
			Required: true,
			BodyPath: "ratingMax",
		},
		&requestflag.Flag[float64]{
			Name:     "rating-min",
			Usage:    "The minimum possible rating value (e.g., 1 for a 1-5 scale).",
			Required: true,
			BodyPath: "ratingMin",
		},
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Usage:    "The Greenflash conversation ID to rate. Either conversationId, externalConversationId, messageId, or externalMessageId must be provided.",
			BodyPath: "conversationId",
		},
		&requestflag.Flag[string]{
			Name:     "external-conversation-id",
			Usage:    "Your external conversation identifier to rate. Either conversationId, externalConversationId, messageId, or externalMessageId must be provided.",
			BodyPath: "externalConversationId",
		},
		&requestflag.Flag[string]{
			Name:     "external-message-id",
			Usage:    "Your external message identifier to rate. Either conversationId, externalConversationId, messageId, or externalMessageId must be provided.",
			BodyPath: "externalMessageId",
		},
		&requestflag.Flag[string]{
			Name:     "feedback",
			Usage:    "Optional text feedback accompanying the rating.",
			BodyPath: "feedback",
		},
		&requestflag.Flag[string]{
			Name:     "message-id",
			Usage:    "The Greenflash message ID to rate. Either conversationId, externalConversationId, messageId, or externalMessageId must be provided.",
			BodyPath: "messageId",
		},
		&requestflag.Flag[any]{
			Name:     "rated-at",
			Usage:    "When the rating was given. Defaults to current time if not provided.",
			BodyPath: "ratedAt",
		},
	},
	Action:          handleRatingsLog,
	HideHelpCommand: true,
}

func handleRatingsLog(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.RatingLogParams{}

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
	_, err = client.Ratings.Log(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ratings log", obj, format, transform)
}
