// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/greenflash-public-api-cli/internal/apiquery"
	"github.com/stainless-sdks/greenflash-public-api-cli/internal/requestflag"
	"github.com/stainless-sdks/greenflash-public-api-go"
	"github.com/stainless-sdks/greenflash-public-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var organizationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Group your users by company, team, or any organizational structure that makes\nsense for your business.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "external-organization-id",
			Usage:    "Your unique identifier for the organization. Use this same ID in other API calls to reference this organization.",
			Required: true,
			BodyPath: "externalOrganizationId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The organization's name.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Usage:    "Custom organization properties.",
			BodyPath: "properties",
		},
	},
	Action:          handleOrganizationsCreate,
	HideHelpCommand: true,
}

var organizationsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update specific fields of an existing organization without changing everything.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "organization-id",
			Usage:    "Your external organization ID (the externalOrganizationId used when creating the organization)",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The organization's name.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Usage:    "Custom organization properties.",
			BodyPath: "properties",
		},
	},
	Action:          handleOrganizationsUpdate,
	HideHelpCommand: true,
}

var organizationsList = cli.Command{
	Name:    "list",
	Usage:   "Browse through all the organizations (companies, teams, etc.) in your workspace.\nSearch for specific organizations or paginate through the full list. Perfect for\nbuilding admin dashboards or organization management interfaces.",
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
			Usage:     "Page number (used to derive offset = (page-1)*limit).",
			QueryPath: "page",
		},
	},
	Action:          handleOrganizationsList,
	HideHelpCommand: true,
}

var organizationsGetOrganizationAnalytics = cli.Command{
	Name:    "get-organization-analytics",
	Usage:   "See how an entire organization (company, team, etc.) engages with your AI across\nall their users and conversations. Spot trends, measure satisfaction, and\nidentify opportunities to improve the experience for your biggest customers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "organization-id",
			Usage:    "The organization ID to get analytics for",
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
	Action:          handleOrganizationsGetOrganizationAnalytics,
	HideHelpCommand: true,
}

func handleOrganizationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.OrganizationNewParams{}

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
	_, err = client.Organizations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "organizations create", obj, format, transform)
}

func handleOrganizationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.OrganizationUpdateParams{}

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
	_, err = client.Organizations.Update(
		ctx,
		cmd.Value("organization-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "organizations update", obj, format, transform)
}

func handleOrganizationsList(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.OrganizationListParams{}

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
	_, err = client.Organizations.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "organizations list", obj, format, transform)
}

func handleOrganizationsGetOrganizationAnalytics(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.OrganizationGetOrganizationAnalyticsParams{}

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
	_, err = client.Organizations.GetOrganizationAnalytics(
		ctx,
		cmd.Value("organization-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "organizations get-organization-analytics", obj, format, transform)
}
