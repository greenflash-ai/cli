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

var promptsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new prompt that you can use across your AI applications. Build prompts\nfrom one or more components, and use handlebars-style variables like\n`{{userName}}` for personalization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "component",
			Usage:    "Array of component objects.",
			Required: true,
			BodyPath: "components",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Prompt name.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "Product this prompt will map to.",
			Required: true,
			BodyPath: "productId",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    `Role key in the product mapping (e.g. "agent tool").`,
			Required: true,
			BodyPath: "role",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Prompt description.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "external-prompt-id",
			Usage:    "Your external identifier for the prompt.",
			BodyPath: "externalPromptId",
		},
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    "Prompt source.",
			Default:  "customer",
			BodyPath: "source",
		},
	},
	Action:          handlePromptsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"component": {
		&requestflag.InnerFlag[string]{
			Name:       "component.content",
			Usage:      "The content of the component.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.component-id",
			Usage:      "The Greenflash component ID.",
			InnerField: "componentId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.external-component-id",
			Usage:      "Your external identifier for the component.",
			InnerField: "externalComponentId",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "component.is-dynamic",
			Usage:      "Whether the component content changes dynamically.",
			InnerField: "isDynamic",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.name",
			Usage:      "Component name.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.source",
			Usage:      "Component source: customer, participant, greenflash, or agent.",
			InnerField: "source",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.type",
			Usage:      "Component type: system, user, tool, guardrail, rag, agent, or a custom type (other).",
			InnerField: "type",
		},
	},
})

var promptsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a prompt with new content or properties. Your production prompt stays\nsafe—updates create new versions without affecting what's currently active.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The prompt ID to update",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "component",
			Usage:    "Updated components (if provided, creates new immutable prompt and version).",
			BodyPath: "components",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Updated prompt description.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Updated prompt name.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "Role key in the product mapping.",
			BodyPath: "role",
		},
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    "Prompt source.",
			Default:  "customer",
			BodyPath: "source",
		},
	},
	Action:          handlePromptsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"component": {
		&requestflag.InnerFlag[string]{
			Name:       "component.content",
			Usage:      "Updated component content.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.component-id",
			Usage:      "The Greenflash component ID.",
			InnerField: "componentId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.external-component-id",
			Usage:      "External component identifier.",
			InnerField: "externalComponentId",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "component.is-dynamic",
			Usage:      "Dynamic flag.",
			InnerField: "isDynamic",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.name",
			Usage:      "Component name.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.source",
			Usage:      "Component source.",
			InnerField: "source",
		},
		&requestflag.InnerFlag[string]{
			Name:       "component.type",
			Usage:      "Component type: system, user, tool, guardrail, rag, agent, or a custom type (other).",
			InnerField: "type",
		},
	},
})

var promptsList = cli.Command{
	Name:    "list",
	Usage:   "Browse through all your prompts to see what you're using across your AI\napplications. Returns all prompts by default (both active and inactive\nversions), or filter by product or version to narrow down the results.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "active-only",
			Usage:     "Filter to only show prompts that are part of active versions. When true, only returns prompts associated with versions where isActive=true.",
			Default:   false,
			QueryPath: "activeOnly",
		},
		&requestflag.Flag[bool]{
			Name:      "include-archived",
			Usage:     "Include archived prompts.",
			Default:   false,
			QueryPath: "includeArchived",
		},
		&requestflag.Flag[float64]{
			Name:      "limit",
			Usage:     "Page size limit (cursor-based pagination). Use either limit/cursor OR page/pageSize, not both.",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "page",
			Usage:     "Page number (page-based pagination). Use either page/pageSize OR limit/cursor, not both.",
			QueryPath: "page",
		},
		&requestflag.Flag[float64]{
			Name:      "page-size",
			Usage:     "Number of items per page (page-based pagination). Use either page/pageSize OR limit/cursor, not both.",
			QueryPath: "pageSize",
		},
		&requestflag.Flag[string]{
			Name:      "product-id",
			Usage:     "Filter prompts by product ID.",
			QueryPath: "productId",
		},
		&requestflag.Flag[string]{
			Name:      "version-id",
			Usage:     "Filter prompts by specific version ID.",
			QueryPath: "versionId",
		},
	},
	Action:          handlePromptsList,
	HideHelpCommand: true,
}

var promptsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Archive a prompt you no longer need. Archived prompts are soft-deleted (we set\nan `archived_at` timestamp) so you can still access them for historical data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The prompt ID to archive",
			Required: true,
		},
	},
	Action:          handlePromptsDelete,
	HideHelpCommand: true,
}

var promptsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a prompt and optionally personalize it with dynamic variables. Perfect\nfor fetching the prompt you want to use right before sending it to your AI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The prompt identifier. Can be: internal prompt ID (UUID), or externalPromptId (your custom ID).",
			Required: true,
		},
	},
	Action:          handlePromptsGet,
	HideHelpCommand: true,
}

func handlePromptsCreate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.PromptNewParams{}

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
	_, err = client.Prompts.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "prompts create", obj, format, transform)
}

func handlePromptsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.PromptUpdateParams{}

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
	_, err = client.Prompts.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "prompts update", obj, format, transform)
}

func handlePromptsList(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := greenflashpublicapi.PromptListParams{}

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
	_, err = client.Prompts.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "prompts list", obj, format, transform)
}

func handlePromptsDelete(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Prompts.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "prompts delete", obj, format, transform)
}

func handlePromptsGet(ctx context.Context, cmd *cli.Command) error {
	client := greenflashpublicapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Prompts.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "prompts get", obj, format, transform)
}
