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

var usersCreate = cli.Command{
	Name:    "create",
	Usage:   "Keep track of who's talking to your AI by creating user profiles with contact\ninformation and custom properties.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "external-user-id",
			Usage:    "Your unique identifier for the user. Use this same ID in other API calls to reference this user.",
			Required: true,
			BodyPath: "externalUserId",
		},
		&requestflag.Flag[bool]{
			Name:     "anonymized",
			Usage:    "Whether to anonymize the user's personal information. Defaults to false.",
			BodyPath: "anonymized",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "The user's email address.",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "external-organization-id",
			Usage:    "Your unique identifier for the organization this user belongs to. If provided, the user will be associated with this organization.",
			BodyPath: "externalOrganizationId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The user's full name.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "organization-id",
			Usage:    "The Greenflash organization ID that the user belongs to.",
			BodyPath: "organizationId",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			Usage:    "The user's phone number.",
			BodyPath: "phone",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Usage:    "Additional data about the user (e.g., plan type, preferences).",
			BodyPath: "properties",
		},
	},
	Action:          handleUsersCreate,
	HideHelpCommand: true,
}

var usersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing user profile without changing everything.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "Your external user ID (the externalUserId used when creating the user)",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "anonymized",
			Usage:    "Whether to anonymize the user's personal information.",
			BodyPath: "anonymized",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "The user's email address.",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "external-organization-id",
			Usage:    "Your unique identifier for the organization this user belongs to. If provided, the user will be associated with this organization.",
			BodyPath: "externalOrganizationId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The user's full name.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "organization-id",
			Usage:    "The Greenflash organization ID that the user belongs to.",
			BodyPath: "organizationId",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			Usage:    "The user's phone number.",
			BodyPath: "phone",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Usage:    "Additional data about the user (e.g., plan type, preferences).",
			BodyPath: "properties",
		},
	},
	Action:          handleUsersUpdate,
	HideHelpCommand: true,
}

var usersList = cli.Command{
	Name:    "list",
	Usage:   "Browse through all the users in your workspace. Filter by organization to see\nwho belongs to specific teams or companies. Results are paginated for easy\nnavigation through large user bases.",
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
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Usage:     "Filter users by organization ID.",
			QueryPath: "organizationId",
		},
		&requestflag.Flag[float64]{
			Name:      "page",
			Usage:     "Page number (used to derive offset = (page-1)*limit).",
			QueryPath: "page",
		},
	},
	Action:          handleUsersList,
	HideHelpCommand: true,
}

var usersGetUserAnalytics = cli.Command{
	Name:    "get-user-analytics",
	Usage:   "Understand how a specific user engages with your AI across all their\nconversations. Track their satisfaction, identify pain points, and spot\nopportunities to improve their experience.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "The user ID to get analytics for",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "mode",
			Usage:     `Analysis mode: "simple" returns only numeric aggregates (no rate limiting), "insights" includes topics, keywords, and recommendations (rate limited per tenant plan).`,
			Default:   "insights",
			QueryPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:      "product-id",
			Usage:     "Filter analytics by product ID.",
			QueryPath: "productId",
		},
		&requestflag.Flag[string]{
			Name:      "version-id",
			Usage:     "Filter analytics by version ID.",
			QueryPath: "versionId",
		},
	},
	Action:          handleUsersGetUserAnalytics,
	HideHelpCommand: true,
}

func handleUsersCreate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.UserNewParams{}

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
	_, err = client.Users.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users create", obj, format, transform)
}

func handleUsersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.UserUpdateParams{}

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
	_, err = client.Users.Update(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users update", obj, format, transform)
}

func handleUsersList(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.UserListParams{}

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
	_, err = client.Users.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users list", obj, format, transform)
}

func handleUsersGetUserAnalytics(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.UserGetUserAnalyticsParams{}

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
	_, err = client.Users.GetUserAnalytics(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users get-user-analytics", obj, format, transform)
}
