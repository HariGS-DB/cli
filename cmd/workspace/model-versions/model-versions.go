// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package model_versions

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/spf13/cobra"
)

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var cmdOverrides []func(*cobra.Command)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model-versions",
		Short: `Databricks provides a hosted version of MLflow Model Registry in Unity Catalog.`,
		Long: `Databricks provides a hosted version of MLflow Model Registry in Unity
  Catalog. Models in Unity Catalog provide centralized access control, auditing,
  lineage, and discovery of ML models across Databricks workspaces.
  
  This API reference documents the REST endpoints for managing model versions in
  Unity Catalog. For more details, see the [registered models API
  docs](/api/workspace/registeredmodels).`,
		GroupID: "catalog",
		Annotations: map[string]string{
			"package": "catalog",
		},
	}

	// Apply optional overrides to this command.
	for _, fn := range cmdOverrides {
		fn(cmd)
	}

	return cmd
}

// start delete command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var deleteOverrides []func(
	*cobra.Command,
	*catalog.DeleteModelVersionRequest,
)

func newDelete() *cobra.Command {
	cmd := &cobra.Command{}

	var deleteReq catalog.DeleteModelVersionRequest

	// TODO: short flags

	cmd.Use = "delete FULL_NAME VERSION"
	cmd.Short = `Delete a Model Version.`
	cmd.Long = `Delete a Model Version.
  
  Deletes a model version from the specified registered model. Any aliases
  assigned to the model version will also be deleted.
  
  The caller must be a metastore admin or an owner of the parent registered
  model. For the latter case, the caller must also be the owner or have the
  **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
  privilege on the parent schema.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		deleteReq.FullName = args[0]
		_, err = fmt.Sscan(args[1], &deleteReq.Version)
		if err != nil {
			return fmt.Errorf("invalid VERSION: %s", args[1])
		}

		err = w.ModelVersions.Delete(ctx, deleteReq)
		if err != nil {
			return err
		}
		return nil
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range deleteOverrides {
		fn(cmd, &deleteReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newDelete())
	})
}

// start get command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var getOverrides []func(
	*cobra.Command,
	*catalog.GetModelVersionRequest,
)

func newGet() *cobra.Command {
	cmd := &cobra.Command{}

	var getReq catalog.GetModelVersionRequest

	// TODO: short flags

	cmd.Use = "get FULL_NAME VERSION"
	cmd.Short = `Get a Model Version.`
	cmd.Long = `Get a Model Version.
  
  Get a model version.
  
  The caller must be a metastore admin or an owner of (or have the **EXECUTE**
  privilege on) the parent registered model. For the latter case, the caller
  must also be the owner or have the **USE_CATALOG** privilege on the parent
  catalog and the **USE_SCHEMA** privilege on the parent schema.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		getReq.FullName = args[0]
		_, err = fmt.Sscan(args[1], &getReq.Version)
		if err != nil {
			return fmt.Errorf("invalid VERSION: %s", args[1])
		}

		response, err := w.ModelVersions.Get(ctx, getReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range getOverrides {
		fn(cmd, &getReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newGet())
	})
}

// start get-by-alias command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var getByAliasOverrides []func(
	*cobra.Command,
	*catalog.GetByAliasRequest,
)

func newGetByAlias() *cobra.Command {
	cmd := &cobra.Command{}

	var getByAliasReq catalog.GetByAliasRequest

	// TODO: short flags

	cmd.Use = "get-by-alias FULL_NAME ALIAS"
	cmd.Short = `Get Model Version By Alias.`
	cmd.Long = `Get Model Version By Alias.
  
  Get a model version by alias.
  
  The caller must be a metastore admin or an owner of (or have the **EXECUTE**
  privilege on) the registered model. For the latter case, the caller must also
  be the owner or have the **USE_CATALOG** privilege on the parent catalog and
  the **USE_SCHEMA** privilege on the parent schema.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		getByAliasReq.FullName = args[0]
		getByAliasReq.Alias = args[1]

		response, err := w.ModelVersions.GetByAlias(ctx, getByAliasReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range getByAliasOverrides {
		fn(cmd, &getByAliasReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newGetByAlias())
	})
}

// start list command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var listOverrides []func(
	*cobra.Command,
	*catalog.ListModelVersionsRequest,
)

func newList() *cobra.Command {
	cmd := &cobra.Command{}

	var listReq catalog.ListModelVersionsRequest

	// TODO: short flags

	cmd.Flags().IntVar(&listReq.MaxResults, "max-results", listReq.MaxResults, `Max number of model versions to return.`)
	cmd.Flags().StringVar(&listReq.PageToken, "page-token", listReq.PageToken, `Opaque token to send for the next page of results (pagination).`)

	cmd.Use = "list FULL_NAME"
	cmd.Short = `List Model Versions.`
	cmd.Long = `List Model Versions.
  
  List model versions. You can list model versions under a particular schema, or
  list all model versions in the current metastore.
  
  The returned models are filtered based on the privileges of the calling user.
  For example, the metastore admin is able to list all the model versions. A
  regular user needs to be the owner or have the **EXECUTE** privilege on the
  parent registered model to recieve the model versions in the response. For the
  latter case, the caller must also be the owner or have the **USE_CATALOG**
  privilege on the parent catalog and the **USE_SCHEMA** privilege on the parent
  schema.
  
  There is no guarantee of a specific ordering of the elements in the response.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		listReq.FullName = args[0]

		response, err := w.ModelVersions.ListAll(ctx, listReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range listOverrides {
		fn(cmd, &listReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newList())
	})
}

// start update command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var updateOverrides []func(
	*cobra.Command,
	*catalog.UpdateModelVersionRequest,
)

func newUpdate() *cobra.Command {
	cmd := &cobra.Command{}

	var updateReq catalog.UpdateModelVersionRequest
	var updateJson flags.JsonFlag

	// TODO: short flags
	cmd.Flags().Var(&updateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	cmd.Flags().StringVar(&updateReq.Comment, "comment", updateReq.Comment, `The comment attached to the model version.`)

	cmd.Use = "update FULL_NAME VERSION"
	cmd.Short = `Update a Model Version.`
	cmd.Long = `Update a Model Version.
  
  Updates the specified model version.
  
  The caller must be a metastore admin or an owner of the parent registered
  model. For the latter case, the caller must also be the owner or have the
  **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
  privilege on the parent schema.
  
  Currently only the comment of the model version can be updated.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		if cmd.Flags().Changed("json") {
			err = updateJson.Unmarshal(&updateReq)
			if err != nil {
				return err
			}
		}
		updateReq.FullName = args[0]
		_, err = fmt.Sscan(args[1], &updateReq.Version)
		if err != nil {
			return fmt.Errorf("invalid VERSION: %s", args[1])
		}

		response, err := w.ModelVersions.Update(ctx, updateReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range updateOverrides {
		fn(cmd, &updateReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newUpdate())
	})
}

// end service ModelVersions
