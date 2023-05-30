// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Git Credentials, Repos, Secrets, Workspace, etc.
package workspace

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewGitCredentials(client *client.DatabricksClient) *GitCredentialsAPI {
	return &GitCredentialsAPI{
		impl: &gitCredentialsImpl{
			client: client,
		},
	}
}

// Registers personal access token for Databricks to do operations on behalf of
// the user.
//
// See [more info].
//
// [more info]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
type GitCredentialsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(GitCredentialsService)
	impl GitCredentialsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *GitCredentialsAPI) WithImpl(impl GitCredentialsService) *GitCredentialsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level GitCredentials API implementation
func (a *GitCredentialsAPI) Impl() GitCredentialsService {
	return a.impl
}

// Create a credential entry.
//
// Creates a Git credential entry for the user. Only one Git credential per user
// is supported, so any attempts to create credentials if an entry already
// exists will fail. Use the PATCH endpoint to update existing credentials, or
// the DELETE endpoint to delete existing credentials.
func (a *GitCredentialsAPI) Create(ctx context.Context, request CreateCredentials) (*CreateCredentialsResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a credential.
//
// Deletes the specified Git credential.
func (a *GitCredentialsAPI) Delete(ctx context.Context, request DeleteGitCredentialRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a credential.
//
// Deletes the specified Git credential.
func (a *GitCredentialsAPI) DeleteByCredentialId(ctx context.Context, credentialId int64) error {
	return a.impl.Delete(ctx, DeleteGitCredentialRequest{
		CredentialId: credentialId,
	})
}

// Get a credential entry.
//
// Gets the Git credential with the specified credential ID.
func (a *GitCredentialsAPI) Get(ctx context.Context, request GetGitCredentialRequest) (*CredentialInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a credential entry.
//
// Gets the Git credential with the specified credential ID.
func (a *GitCredentialsAPI) GetByCredentialId(ctx context.Context, credentialId int64) (*CredentialInfo, error) {
	return a.impl.Get(ctx, GetGitCredentialRequest{
		CredentialId: credentialId,
	})
}

// Get Git credentials.
//
// Lists the calling user's Git credentials. One credential per user is
// supported.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GitCredentialsAPI) ListAll(ctx context.Context) ([]CredentialInfo, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Credentials, nil
}

// CredentialInfoGitProviderToCredentialIdMap calls [GitCredentialsAPI.ListAll] and creates a map of results with [CredentialInfo].GitProvider as key and [CredentialInfo].CredentialId as value.
//
// Returns an error if there's more than one [CredentialInfo] with the same .GitProvider.
//
// Note: All [CredentialInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GitCredentialsAPI) CredentialInfoGitProviderToCredentialIdMap(ctx context.Context) (map[string]int64, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]int64{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.GitProvider
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .GitProvider: %s", key)
		}
		mapping[key] = v.CredentialId
	}
	return mapping, nil
}

// GetByGitProvider calls [GitCredentialsAPI.CredentialInfoGitProviderToCredentialIdMap] and returns a single [CredentialInfo].
//
// Returns an error if there's more than one [CredentialInfo] with the same .GitProvider.
//
// Note: All [CredentialInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GitCredentialsAPI) GetByGitProvider(ctx context.Context, name string) (*CredentialInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]CredentialInfo{}
	for _, v := range result {
		key := v.GitProvider
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("CredentialInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of CredentialInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update a credential.
//
// Updates the specified Git credential.
func (a *GitCredentialsAPI) Update(ctx context.Context, request UpdateCredentials) error {
	return a.impl.Update(ctx, request)
}

func NewRepos(client *client.DatabricksClient) *ReposAPI {
	return &ReposAPI{
		impl: &reposImpl{
			client: client,
		},
	}
}

// The Repos API allows users to manage their git repos. Users can use the API
// to access all repos that they have manage permissions on.
//
// Databricks Repos is a visual Git client in Databricks. It supports common Git
// operations such a cloning a repository, committing and pushing, pulling,
// branch management, and visual comparison of diffs when committing.
//
// Within Repos you can develop code in notebooks or other files and follow data
// science and engineering code development best practices using Git for version
// control, collaboration, and CI/CD.
type ReposAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ReposService)
	impl ReposService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ReposAPI) WithImpl(impl ReposService) *ReposAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Repos API implementation
func (a *ReposAPI) Impl() ReposService {
	return a.impl
}

// Create a repo.
//
// Creates a repo in the workspace and links it to the remote Git repo
// specified. Note that repos created programmatically must be linked to a
// remote Git repo, unlike repos created in the browser.
func (a *ReposAPI) Create(ctx context.Context, request CreateRepo) (*RepoInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a repo.
//
// Deletes the specified repo.
func (a *ReposAPI) Delete(ctx context.Context, request DeleteRepoRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a repo.
//
// Deletes the specified repo.
func (a *ReposAPI) DeleteByRepoId(ctx context.Context, repoId int64) error {
	return a.impl.Delete(ctx, DeleteRepoRequest{
		RepoId: repoId,
	})
}

// Get a repo.
//
// Returns the repo with the given repo ID.
func (a *ReposAPI) Get(ctx context.Context, request GetRepoRequest) (*RepoInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a repo.
//
// Returns the repo with the given repo ID.
func (a *ReposAPI) GetByRepoId(ctx context.Context, repoId int64) (*RepoInfo, error) {
	return a.impl.Get(ctx, GetRepoRequest{
		RepoId: repoId,
	})
}

// Get repos.
//
// Returns repos that the calling user has Manage permissions on. Results are
// paginated with each page containing twenty repos.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ReposAPI) ListAll(ctx context.Context, request ListReposRequest) ([]RepoInfo, error) {
	var results []RepoInfo
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Repos) == 0 {
			break
		}
		for _, v := range response.Repos {
			results = append(results, v)
		}
		request.NextPageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// RepoInfoPathToIdMap calls [ReposAPI.ListAll] and creates a map of results with [RepoInfo].Path as key and [RepoInfo].Id as value.
//
// Returns an error if there's more than one [RepoInfo] with the same .Path.
//
// Note: All [RepoInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ReposAPI) RepoInfoPathToIdMap(ctx context.Context, request ListReposRequest) (map[string]int64, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]int64{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Path
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Path: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByPath calls [ReposAPI.RepoInfoPathToIdMap] and returns a single [RepoInfo].
//
// Returns an error if there's more than one [RepoInfo] with the same .Path.
//
// Note: All [RepoInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ReposAPI) GetByPath(ctx context.Context, name string) (*RepoInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListReposRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]RepoInfo{}
	for _, v := range result {
		key := v.Path
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("RepoInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of RepoInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update a repo.
//
// Updates the repo to a different branch or tag, or updates the repo to the
// latest commit on the same branch.
func (a *ReposAPI) Update(ctx context.Context, request UpdateRepo) error {
	return a.impl.Update(ctx, request)
}

func NewSecrets(client *client.DatabricksClient) *SecretsAPI {
	return &SecretsAPI{
		impl: &secretsImpl{
			client: client,
		},
	}
}

// The Secrets API allows you to manage secrets, secret scopes, and access
// permissions.
//
// Sometimes accessing data requires that you authenticate to external data
// sources through JDBC. Instead of directly entering your credentials into a
// notebook, use Databricks secrets to store your credentials and reference them
// in notebooks and jobs.
//
// Administrators, secret creators, and users granted permission can read
// Databricks secrets. While Databricks makes an effort to redact secret values
// that might be displayed in notebooks, it is not possible to prevent such
// users from reading secrets.
type SecretsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(SecretsService)
	impl SecretsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *SecretsAPI) WithImpl(impl SecretsService) *SecretsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Secrets API implementation
func (a *SecretsAPI) Impl() SecretsService {
	return a.impl
}

// Create a new secret scope.
//
// The scope name must consist of alphanumeric characters, dashes, underscores,
// and periods, and may not exceed 128 characters. The maximum number of scopes
// in a workspace is 100.
func (a *SecretsAPI) CreateScope(ctx context.Context, request CreateScope) error {
	return a.impl.CreateScope(ctx, request)
}

// Delete an ACL.
//
// Deletes the given ACL on the given scope.
//
// Users must have the `MANAGE` permission to invoke this API. Throws
// `RESOURCE_DOES_NOT_EXIST` if no such secret scope, principal, or ACL exists.
// Throws `PERMISSION_DENIED` if the user does not have permission to make this
// API call.
func (a *SecretsAPI) DeleteAcl(ctx context.Context, request DeleteAcl) error {
	return a.impl.DeleteAcl(ctx, request)
}

// Delete a secret scope.
//
// Deletes a secret scope.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if the scope does not exist. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) DeleteScope(ctx context.Context, request DeleteScope) error {
	return a.impl.DeleteScope(ctx, request)
}

// Delete a secret scope.
//
// Deletes a secret scope.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if the scope does not exist. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) DeleteScopeByScope(ctx context.Context, scope string) error {
	return a.impl.DeleteScope(ctx, DeleteScope{
		Scope: scope,
	})
}

// Delete a secret.
//
// Deletes the secret stored in this secret scope. You must have `WRITE` or
// `MANAGE` permission on the secret scope.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope or secret exists.
// Throws `PERMISSION_DENIED` if the user does not have permission to make this
// API call.
func (a *SecretsAPI) DeleteSecret(ctx context.Context, request DeleteSecret) error {
	return a.impl.DeleteSecret(ctx, request)
}

// Get secret ACL details.
//
// Gets the details about the given ACL, such as the group and permission. Users
// must have the `MANAGE` permission to invoke this API.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) GetAcl(ctx context.Context, request GetAclRequest) (*AclItem, error) {
	return a.impl.GetAcl(ctx, request)
}

// Lists ACLs.
//
// List the ACLs for a given secret scope. Users must have the `MANAGE`
// permission to invoke this API.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SecretsAPI) ListAclsAll(ctx context.Context, request ListAclsRequest) ([]AclItem, error) {
	response, err := a.impl.ListAcls(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}

// Lists ACLs.
//
// List the ACLs for a given secret scope. Users must have the `MANAGE`
// permission to invoke this API.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) ListAclsByScope(ctx context.Context, scope string) (*ListAclsResponse, error) {
	return a.impl.ListAcls(ctx, ListAclsRequest{
		Scope: scope,
	})
}

// List all scopes.
//
// Lists all secret scopes available in the workspace.
//
// Throws `PERMISSION_DENIED` if the user does not have permission to make this
// API call.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SecretsAPI) ListScopesAll(ctx context.Context) ([]SecretScope, error) {
	response, err := a.impl.ListScopes(ctx)
	if err != nil {
		return nil, err
	}
	return response.Scopes, nil
}

// List secret keys.
//
// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// The lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
// `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SecretsAPI) ListSecretsAll(ctx context.Context, request ListSecretsRequest) ([]SecretMetadata, error) {
	response, err := a.impl.ListSecrets(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Secrets, nil
}

// List secret keys.
//
// Lists the secret keys that are stored at this scope. This is a metadata-only
// operation; secret data cannot be retrieved using this API. Users need the
// READ permission to make this call.
//
// The lastUpdatedTimestamp returned is in milliseconds since epoch. Throws
// `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) ListSecretsByScope(ctx context.Context, scope string) (*ListSecretsResponse, error) {
	return a.impl.ListSecrets(ctx, ListSecretsRequest{
		Scope: scope,
	})
}

// Create/update an ACL.
//
// Creates or overwrites the Access Control List (ACL) associated with the given
// principal (user or group) on the specified scope point.
//
// In general, a user or group will use the most powerful permission available
// to them, and permissions are ordered as follows:
//
// * `MANAGE` - Allowed to change ACLs, and read and write to this secret scope.
// * `WRITE` - Allowed to read and write to this secret scope. * `READ` -
// Allowed to read this secret scope and list what secrets are available.
//
// Note that in general, secret values can only be read from within a command on
// a cluster (for example, through a notebook). There is no API to read the
// actual secret value material outside of a cluster. However, the user's
// permission will be applied based on who is executing the command, and they
// must have at least READ permission.
//
// Users must have the `MANAGE` permission to invoke this API.
//
// The principal is a user or group name corresponding to an existing Databricks
// principal to be granted or revoked access.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `RESOURCE_ALREADY_EXISTS` if a permission for the principal already exists.
// Throws `INVALID_PARAMETER_VALUE` if the permission is invalid. Throws
// `PERMISSION_DENIED` if the user does not have permission to make this API
// call.
func (a *SecretsAPI) PutAcl(ctx context.Context, request PutAcl) error {
	return a.impl.PutAcl(ctx, request)
}

// Add a secret.
//
// Inserts a secret under the provided scope with the given name. If a secret
// already exists with the same name, this command overwrites the existing
// secret's value. The server encrypts the secret using the secret scope's
// encryption settings before storing it.
//
// You must have `WRITE` or `MANAGE` permission on the secret scope. The secret
// key must consist of alphanumeric characters, dashes, underscores, and
// periods, and cannot exceed 128 characters. The maximum allowed secret value
// size is 128 KB. The maximum number of secrets in a given scope is 1000.
//
// The input fields "string_value" or "bytes_value" specify the type of the
// secret, which will determine the value returned when the secret value is
// requested. Exactly one must be specified.
//
// Throws `RESOURCE_DOES_NOT_EXIST` if no such secret scope exists. Throws
// `RESOURCE_LIMIT_EXCEEDED` if maximum number of secrets in scope is exceeded.
// Throws `INVALID_PARAMETER_VALUE` if the key name or value length is invalid.
// Throws `PERMISSION_DENIED` if the user does not have permission to make this
// API call.
func (a *SecretsAPI) PutSecret(ctx context.Context, request PutSecret) error {
	return a.impl.PutSecret(ctx, request)
}

func NewWorkspace(client *client.DatabricksClient) *WorkspaceAPI {
	return &WorkspaceAPI{
		impl: &workspaceImpl{
			client: client,
		},
	}
}

// The Workspace API allows you to list, import, export, and delete notebooks
// and folders.
//
// A notebook is a web-based interface to a document that contains runnable
// code, visualizations, and explanatory text.
type WorkspaceAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WorkspaceService)
	impl WorkspaceService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WorkspaceAPI) WithImpl(impl WorkspaceService) *WorkspaceAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Workspace API implementation
func (a *WorkspaceAPI) Impl() WorkspaceService {
	return a.impl
}

// Delete a workspace object.
//
// Deletes an object or a directory (and optionally recursively deletes all
// objects in the directory). * If `path` does not exist, this call returns an
// error `RESOURCE_DOES_NOT_EXIST`. * If `path` is a non-empty directory and
// `recursive` is set to `false`, this call returns an error
// `DIRECTORY_NOT_EMPTY`.
//
// Object deletion cannot be undone and deleting a directory recursively is not
// atomic.
func (a *WorkspaceAPI) Delete(ctx context.Context, request Delete) error {
	return a.impl.Delete(ctx, request)
}

// Export a workspace object.
//
// Exports an object or the contents of an entire directory.
//
// If `path` does not exist, this call returns an error
// `RESOURCE_DOES_NOT_EXIST`.
//
// One can only export a directory in `DBC` format. If the exported data would
// exceed size limit, this call returns `MAX_NOTEBOOK_SIZE_EXCEEDED`. Currently,
// this API does not support exporting a library.
func (a *WorkspaceAPI) Export(ctx context.Context, request ExportRequest) (*ExportResponse, error) {
	return a.impl.Export(ctx, request)
}

// Get status.
//
// Gets the status of an object or a directory. If `path` does not exist, this
// call returns an error `RESOURCE_DOES_NOT_EXIST`.
func (a *WorkspaceAPI) GetStatus(ctx context.Context, request GetStatusRequest) (*ObjectInfo, error) {
	return a.impl.GetStatus(ctx, request)
}

// Get status.
//
// Gets the status of an object or a directory. If `path` does not exist, this
// call returns an error `RESOURCE_DOES_NOT_EXIST`.
func (a *WorkspaceAPI) GetStatusByPath(ctx context.Context, path string) (*ObjectInfo, error) {
	return a.impl.GetStatus(ctx, GetStatusRequest{
		Path: path,
	})
}

// Import a workspace object.
//
// Imports a workspace object (for example, a notebook or file) or the contents
// of an entire directory. If `path` already exists and `overwrite` is set to
// `false`, this call returns an error `RESOURCE_ALREADY_EXISTS`. One can only
// use `DBC` format to import a directory.
func (a *WorkspaceAPI) Import(ctx context.Context, request Import) error {
	return a.impl.Import(ctx, request)
}

// List contents.
//
// Lists the contents of a directory, or the object if it is not a directory.If
// the input path does not exist, this call returns an error
// `RESOURCE_DOES_NOT_EXIST`.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAPI) ListAll(ctx context.Context, request ListWorkspaceRequest) ([]ObjectInfo, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Objects, nil
}

// ObjectInfoPathToObjectIdMap calls [WorkspaceAPI.ListAll] and creates a map of results with [ObjectInfo].Path as key and [ObjectInfo].ObjectId as value.
//
// Returns an error if there's more than one [ObjectInfo] with the same .Path.
//
// Note: All [ObjectInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAPI) ObjectInfoPathToObjectIdMap(ctx context.Context, request ListWorkspaceRequest) (map[string]int64, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]int64{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Path
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Path: %s", key)
		}
		mapping[key] = v.ObjectId
	}
	return mapping, nil
}

// GetByPath calls [WorkspaceAPI.ObjectInfoPathToObjectIdMap] and returns a single [ObjectInfo].
//
// Returns an error if there's more than one [ObjectInfo] with the same .Path.
//
// Note: All [ObjectInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAPI) GetByPath(ctx context.Context, name string) (*ObjectInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListWorkspaceRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]ObjectInfo{}
	for _, v := range result {
		key := v.Path
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("ObjectInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of ObjectInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Create a directory.
//
// Creates the specified directory (and necessary parent directories if they do
// not exist). If there is an object (not a directory) at any prefix of the
// input path, this call returns an error `RESOURCE_ALREADY_EXISTS`.
//
// Note that if this operation fails it may have succeeded in creating some of
// the necessary parrent directories.
func (a *WorkspaceAPI) Mkdirs(ctx context.Context, request Mkdirs) error {
	return a.impl.Mkdirs(ctx, request)
}

// Create a directory.
//
// Creates the specified directory (and necessary parent directories if they do
// not exist). If there is an object (not a directory) at any prefix of the
// input path, this call returns an error `RESOURCE_ALREADY_EXISTS`.
//
// Note that if this operation fails it may have succeeded in creating some of
// the necessary parrent directories.
func (a *WorkspaceAPI) MkdirsByPath(ctx context.Context, path string) error {
	return a.impl.Mkdirs(ctx, Mkdirs{
		Path: path,
	})
}