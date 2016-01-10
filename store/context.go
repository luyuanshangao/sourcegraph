// GENERATED CODE - DO NOT EDIT!
//
// Generated by:
//
//   go run gen_context_and_mock.go -o1 context.go -o2 mockstore/mockstores.go
//
// Called via:
//
//   go generate
//

package store

import (
	"golang.org/x/net/context"
	srcstore "sourcegraph.com/sourcegraph/srclib/store"
)

// Stores has a field for each store interface.
type Stores struct {
	Accounts                        Accounts
	Authorizations                  Authorizations
	BuildLogs                       BuildLogs
	Builds                          Builds
	Changesets                      Changesets
	Directory                       Directory
	ExternalAuthTokens              ExternalAuthTokens
	Graph                           srcstore.MultiRepoStoreImporterIndexer
	Invites                         Invites
	MirroredRepoSSHKeys             MirroredRepoSSHKeys
	Orgs                            Orgs
	Password                        Password
	RegisteredClients               RegisteredClients
	RepoConfigs                     RepoConfigs
	RepoCounters                    RepoCounters
	RepoOrigin                      RepoOrigin
	RepoOriginWithAuthorizedSSHKeys RepoOriginWithAuthorizedSSHKeys
	RepoOriginWithCommitStatuses    RepoOriginWithCommitStatuses
	RepoOriginWithPushHooks         RepoOriginWithPushHooks
	RepoStatuses                    RepoStatuses
	RepoVCS                         RepoVCS
	Repos                           Repos
	Storage                         Storage
	UserKeys                        UserKeys
	UserPermissions                 UserPermissions
	Users                           Users
}

type contextKey int

const (
	_AccountsKey contextKey = iota
	_AuthorizationsKey
	_BuildLogsKey
	_BuildsKey
	_ChangesetsKey
	_DirectoryKey
	_ExternalAuthTokensKey
	_GraphKey
	_InvitesKey
	_MirroredRepoSSHKeysKey
	_OrgsKey
	_PasswordKey
	_RegisteredClientsKey
	_RepoConfigsKey
	_RepoCountersKey
	_RepoOriginKey
	_RepoOriginWithAuthorizedSSHKeysKey
	_RepoOriginWithCommitStatusesKey
	_RepoOriginWithPushHooksKey
	_RepoStatusesKey
	_RepoVCSKey
	_ReposKey
	_StorageKey
	_UserKeysKey
	_UserPermissionsKey
	_UsersKey
)

// WithStores returns a copy of parent with the given stores. If a store's field value is nil, its previous value is inherited from parent in the new context.
func WithStores(ctx context.Context, s Stores) context.Context {
	if s.Accounts != nil {
		ctx = WithAccounts(ctx, s.Accounts)
	}
	if s.Authorizations != nil {
		ctx = WithAuthorizations(ctx, s.Authorizations)
	}
	if s.BuildLogs != nil {
		ctx = WithBuildLogs(ctx, s.BuildLogs)
	}
	if s.Builds != nil {
		ctx = WithBuilds(ctx, s.Builds)
	}
	if s.Changesets != nil {
		ctx = WithChangesets(ctx, s.Changesets)
	}
	if s.Directory != nil {
		ctx = WithDirectory(ctx, s.Directory)
	}
	if s.ExternalAuthTokens != nil {
		ctx = WithExternalAuthTokens(ctx, s.ExternalAuthTokens)
	}
	if s.Graph != nil {
		ctx = WithGraph(ctx, s.Graph)
	}
	if s.Invites != nil {
		ctx = WithInvites(ctx, s.Invites)
	}
	if s.MirroredRepoSSHKeys != nil {
		ctx = WithMirroredRepoSSHKeys(ctx, s.MirroredRepoSSHKeys)
	}
	if s.Orgs != nil {
		ctx = WithOrgs(ctx, s.Orgs)
	}
	if s.Password != nil {
		ctx = WithPassword(ctx, s.Password)
	}
	if s.RegisteredClients != nil {
		ctx = WithRegisteredClients(ctx, s.RegisteredClients)
	}
	if s.RepoConfigs != nil {
		ctx = WithRepoConfigs(ctx, s.RepoConfigs)
	}
	if s.RepoCounters != nil {
		ctx = WithRepoCounters(ctx, s.RepoCounters)
	}
	if s.RepoOrigin != nil {
		ctx = WithRepoOrigin(ctx, s.RepoOrigin)
	}
	if s.RepoOriginWithAuthorizedSSHKeys != nil {
		ctx = WithRepoOriginWithAuthorizedSSHKeys(ctx, s.RepoOriginWithAuthorizedSSHKeys)
	}
	if s.RepoOriginWithCommitStatuses != nil {
		ctx = WithRepoOriginWithCommitStatuses(ctx, s.RepoOriginWithCommitStatuses)
	}
	if s.RepoOriginWithPushHooks != nil {
		ctx = WithRepoOriginWithPushHooks(ctx, s.RepoOriginWithPushHooks)
	}
	if s.RepoStatuses != nil {
		ctx = WithRepoStatuses(ctx, s.RepoStatuses)
	}
	if s.RepoVCS != nil {
		ctx = WithRepoVCS(ctx, s.RepoVCS)
	}
	if s.Repos != nil {
		ctx = WithRepos(ctx, s.Repos)
	}
	if s.Storage != nil {
		ctx = WithStorage(ctx, s.Storage)
	}
	if s.UserKeys != nil {
		ctx = WithUserKeys(ctx, s.UserKeys)
	}
	if s.UserPermissions != nil {
		ctx = WithUserPermissions(ctx, s.UserPermissions)
	}
	if s.Users != nil {
		ctx = WithUsers(ctx, s.Users)
	}
	return ctx
}

// WithAccounts returns a copy of parent with the given Accounts store.
func WithAccounts(parent context.Context, s Accounts) context.Context {
	return context.WithValue(parent, _AccountsKey, s)
}

// AccountsFromContext gets the context's Accounts store. If the store is not present, it panics.
func AccountsFromContext(ctx context.Context) Accounts {
	s, ok := ctx.Value(_AccountsKey).(Accounts)
	if !ok || s == nil {
		panic("no Accounts set in context")
	}
	return s
}

// AccountsFromContextOrNil returns the context's Accounts store if present, or else nil.
func AccountsFromContextOrNil(ctx context.Context) Accounts {
	s, ok := ctx.Value(_AccountsKey).(Accounts)
	if ok {
		return s
	}
	return nil
}

// WithAuthorizations returns a copy of parent with the given Authorizations store.
func WithAuthorizations(parent context.Context, s Authorizations) context.Context {
	return context.WithValue(parent, _AuthorizationsKey, s)
}

// AuthorizationsFromContext gets the context's Authorizations store. If the store is not present, it panics.
func AuthorizationsFromContext(ctx context.Context) Authorizations {
	s, ok := ctx.Value(_AuthorizationsKey).(Authorizations)
	if !ok || s == nil {
		panic("no Authorizations set in context")
	}
	return s
}

// AuthorizationsFromContextOrNil returns the context's Authorizations store if present, or else nil.
func AuthorizationsFromContextOrNil(ctx context.Context) Authorizations {
	s, ok := ctx.Value(_AuthorizationsKey).(Authorizations)
	if ok {
		return s
	}
	return nil
}

// WithBuildLogs returns a copy of parent with the given BuildLogs store.
func WithBuildLogs(parent context.Context, s BuildLogs) context.Context {
	return context.WithValue(parent, _BuildLogsKey, s)
}

// BuildLogsFromContext gets the context's BuildLogs store. If the store is not present, it panics.
func BuildLogsFromContext(ctx context.Context) BuildLogs {
	s, ok := ctx.Value(_BuildLogsKey).(BuildLogs)
	if !ok || s == nil {
		panic("no BuildLogs set in context")
	}
	return s
}

// BuildLogsFromContextOrNil returns the context's BuildLogs store if present, or else nil.
func BuildLogsFromContextOrNil(ctx context.Context) BuildLogs {
	s, ok := ctx.Value(_BuildLogsKey).(BuildLogs)
	if ok {
		return s
	}
	return nil
}

// WithBuilds returns a copy of parent with the given Builds store.
func WithBuilds(parent context.Context, s Builds) context.Context {
	return context.WithValue(parent, _BuildsKey, s)
}

// BuildsFromContext gets the context's Builds store. If the store is not present, it panics.
func BuildsFromContext(ctx context.Context) Builds {
	s, ok := ctx.Value(_BuildsKey).(Builds)
	if !ok || s == nil {
		panic("no Builds set in context")
	}
	return s
}

// BuildsFromContextOrNil returns the context's Builds store if present, or else nil.
func BuildsFromContextOrNil(ctx context.Context) Builds {
	s, ok := ctx.Value(_BuildsKey).(Builds)
	if ok {
		return s
	}
	return nil
}

// WithChangesets returns a copy of parent with the given Changesets store.
func WithChangesets(parent context.Context, s Changesets) context.Context {
	return context.WithValue(parent, _ChangesetsKey, s)
}

// ChangesetsFromContext gets the context's Changesets store. If the store is not present, it panics.
func ChangesetsFromContext(ctx context.Context) Changesets {
	s, ok := ctx.Value(_ChangesetsKey).(Changesets)
	if !ok || s == nil {
		panic("no Changesets set in context")
	}
	return s
}

// ChangesetsFromContextOrNil returns the context's Changesets store if present, or else nil.
func ChangesetsFromContextOrNil(ctx context.Context) Changesets {
	s, ok := ctx.Value(_ChangesetsKey).(Changesets)
	if ok {
		return s
	}
	return nil
}

// WithDirectory returns a copy of parent with the given Directory store.
func WithDirectory(parent context.Context, s Directory) context.Context {
	return context.WithValue(parent, _DirectoryKey, s)
}

// DirectoryFromContext gets the context's Directory store. If the store is not present, it panics.
func DirectoryFromContext(ctx context.Context) Directory {
	s, ok := ctx.Value(_DirectoryKey).(Directory)
	if !ok || s == nil {
		panic("no Directory set in context")
	}
	return s
}

// DirectoryFromContextOrNil returns the context's Directory store if present, or else nil.
func DirectoryFromContextOrNil(ctx context.Context) Directory {
	s, ok := ctx.Value(_DirectoryKey).(Directory)
	if ok {
		return s
	}
	return nil
}

// WithExternalAuthTokens returns a copy of parent with the given ExternalAuthTokens store.
func WithExternalAuthTokens(parent context.Context, s ExternalAuthTokens) context.Context {
	return context.WithValue(parent, _ExternalAuthTokensKey, s)
}

// ExternalAuthTokensFromContext gets the context's ExternalAuthTokens store. If the store is not present, it panics.
func ExternalAuthTokensFromContext(ctx context.Context) ExternalAuthTokens {
	s, ok := ctx.Value(_ExternalAuthTokensKey).(ExternalAuthTokens)
	if !ok || s == nil {
		panic("no ExternalAuthTokens set in context")
	}
	return s
}

// ExternalAuthTokensFromContextOrNil returns the context's ExternalAuthTokens store if present, or else nil.
func ExternalAuthTokensFromContextOrNil(ctx context.Context) ExternalAuthTokens {
	s, ok := ctx.Value(_ExternalAuthTokensKey).(ExternalAuthTokens)
	if ok {
		return s
	}
	return nil
}

// WithGraph returns a copy of parent with the given Graph store.
func WithGraph(parent context.Context, s srcstore.MultiRepoStoreImporterIndexer) context.Context {
	return context.WithValue(parent, _GraphKey, s)
}

// GraphFromContext gets the context's Graph store. If the store is not present, it panics.
func GraphFromContext(ctx context.Context) srcstore.MultiRepoStoreImporterIndexer {
	s, ok := ctx.Value(_GraphKey).(srcstore.MultiRepoStoreImporterIndexer)
	if !ok || s == nil {
		panic("no Graph set in context")
	}
	return s
}

// GraphFromContextOrNil returns the context's Graph store if present, or else nil.
func GraphFromContextOrNil(ctx context.Context) srcstore.MultiRepoStoreImporterIndexer {
	s, ok := ctx.Value(_GraphKey).(srcstore.MultiRepoStoreImporterIndexer)
	if ok {
		return s
	}
	return nil
}

// WithInvites returns a copy of parent with the given Invites store.
func WithInvites(parent context.Context, s Invites) context.Context {
	return context.WithValue(parent, _InvitesKey, s)
}

// InvitesFromContext gets the context's Invites store. If the store is not present, it panics.
func InvitesFromContext(ctx context.Context) Invites {
	s, ok := ctx.Value(_InvitesKey).(Invites)
	if !ok || s == nil {
		panic("no Invites set in context")
	}
	return s
}

// InvitesFromContextOrNil returns the context's Invites store if present, or else nil.
func InvitesFromContextOrNil(ctx context.Context) Invites {
	s, ok := ctx.Value(_InvitesKey).(Invites)
	if ok {
		return s
	}
	return nil
}

// WithMirroredRepoSSHKeys returns a copy of parent with the given MirroredRepoSSHKeys store.
func WithMirroredRepoSSHKeys(parent context.Context, s MirroredRepoSSHKeys) context.Context {
	return context.WithValue(parent, _MirroredRepoSSHKeysKey, s)
}

// MirroredRepoSSHKeysFromContext gets the context's MirroredRepoSSHKeys store. If the store is not present, it panics.
func MirroredRepoSSHKeysFromContext(ctx context.Context) MirroredRepoSSHKeys {
	s, ok := ctx.Value(_MirroredRepoSSHKeysKey).(MirroredRepoSSHKeys)
	if !ok || s == nil {
		panic("no MirroredRepoSSHKeys set in context")
	}
	return s
}

// MirroredRepoSSHKeysFromContextOrNil returns the context's MirroredRepoSSHKeys store if present, or else nil.
func MirroredRepoSSHKeysFromContextOrNil(ctx context.Context) MirroredRepoSSHKeys {
	s, ok := ctx.Value(_MirroredRepoSSHKeysKey).(MirroredRepoSSHKeys)
	if ok {
		return s
	}
	return nil
}

// WithOrgs returns a copy of parent with the given Orgs store.
func WithOrgs(parent context.Context, s Orgs) context.Context {
	return context.WithValue(parent, _OrgsKey, s)
}

// OrgsFromContext gets the context's Orgs store. If the store is not present, it panics.
func OrgsFromContext(ctx context.Context) Orgs {
	s, ok := ctx.Value(_OrgsKey).(Orgs)
	if !ok || s == nil {
		panic("no Orgs set in context")
	}
	return s
}

// OrgsFromContextOrNil returns the context's Orgs store if present, or else nil.
func OrgsFromContextOrNil(ctx context.Context) Orgs {
	s, ok := ctx.Value(_OrgsKey).(Orgs)
	if ok {
		return s
	}
	return nil
}

// WithPassword returns a copy of parent with the given Password store.
func WithPassword(parent context.Context, s Password) context.Context {
	return context.WithValue(parent, _PasswordKey, s)
}

// PasswordFromContext gets the context's Password store. If the store is not present, it panics.
func PasswordFromContext(ctx context.Context) Password {
	s, ok := ctx.Value(_PasswordKey).(Password)
	if !ok || s == nil {
		panic("no Password set in context")
	}
	return s
}

// PasswordFromContextOrNil returns the context's Password store if present, or else nil.
func PasswordFromContextOrNil(ctx context.Context) Password {
	s, ok := ctx.Value(_PasswordKey).(Password)
	if ok {
		return s
	}
	return nil
}

// WithRegisteredClients returns a copy of parent with the given RegisteredClients store.
func WithRegisteredClients(parent context.Context, s RegisteredClients) context.Context {
	return context.WithValue(parent, _RegisteredClientsKey, s)
}

// RegisteredClientsFromContext gets the context's RegisteredClients store. If the store is not present, it panics.
func RegisteredClientsFromContext(ctx context.Context) RegisteredClients {
	s, ok := ctx.Value(_RegisteredClientsKey).(RegisteredClients)
	if !ok || s == nil {
		panic("no RegisteredClients set in context")
	}
	return s
}

// RegisteredClientsFromContextOrNil returns the context's RegisteredClients store if present, or else nil.
func RegisteredClientsFromContextOrNil(ctx context.Context) RegisteredClients {
	s, ok := ctx.Value(_RegisteredClientsKey).(RegisteredClients)
	if ok {
		return s
	}
	return nil
}

// WithRepoConfigs returns a copy of parent with the given RepoConfigs store.
func WithRepoConfigs(parent context.Context, s RepoConfigs) context.Context {
	return context.WithValue(parent, _RepoConfigsKey, s)
}

// RepoConfigsFromContext gets the context's RepoConfigs store. If the store is not present, it panics.
func RepoConfigsFromContext(ctx context.Context) RepoConfigs {
	s, ok := ctx.Value(_RepoConfigsKey).(RepoConfigs)
	if !ok || s == nil {
		panic("no RepoConfigs set in context")
	}
	return s
}

// RepoConfigsFromContextOrNil returns the context's RepoConfigs store if present, or else nil.
func RepoConfigsFromContextOrNil(ctx context.Context) RepoConfigs {
	s, ok := ctx.Value(_RepoConfigsKey).(RepoConfigs)
	if ok {
		return s
	}
	return nil
}

// WithRepoCounters returns a copy of parent with the given RepoCounters store.
func WithRepoCounters(parent context.Context, s RepoCounters) context.Context {
	return context.WithValue(parent, _RepoCountersKey, s)
}

// RepoCountersFromContext gets the context's RepoCounters store. If the store is not present, it panics.
func RepoCountersFromContext(ctx context.Context) RepoCounters {
	s, ok := ctx.Value(_RepoCountersKey).(RepoCounters)
	if !ok || s == nil {
		panic("no RepoCounters set in context")
	}
	return s
}

// RepoCountersFromContextOrNil returns the context's RepoCounters store if present, or else nil.
func RepoCountersFromContextOrNil(ctx context.Context) RepoCounters {
	s, ok := ctx.Value(_RepoCountersKey).(RepoCounters)
	if ok {
		return s
	}
	return nil
}

// WithRepoOrigin returns a copy of parent with the given RepoOrigin store.
func WithRepoOrigin(parent context.Context, s RepoOrigin) context.Context {
	return context.WithValue(parent, _RepoOriginKey, s)
}

// RepoOriginFromContext gets the context's RepoOrigin store. If the store is not present, it panics.
func RepoOriginFromContext(ctx context.Context) RepoOrigin {
	s, ok := ctx.Value(_RepoOriginKey).(RepoOrigin)
	if !ok || s == nil {
		panic("no RepoOrigin set in context")
	}
	return s
}

// RepoOriginFromContextOrNil returns the context's RepoOrigin store if present, or else nil.
func RepoOriginFromContextOrNil(ctx context.Context) RepoOrigin {
	s, ok := ctx.Value(_RepoOriginKey).(RepoOrigin)
	if ok {
		return s
	}
	return nil
}

// WithRepoOriginWithAuthorizedSSHKeys returns a copy of parent with the given RepoOriginWithAuthorizedSSHKeys store.
func WithRepoOriginWithAuthorizedSSHKeys(parent context.Context, s RepoOriginWithAuthorizedSSHKeys) context.Context {
	return context.WithValue(parent, _RepoOriginWithAuthorizedSSHKeysKey, s)
}

// RepoOriginWithAuthorizedSSHKeysFromContext gets the context's RepoOriginWithAuthorizedSSHKeys store. If the store is not present, it panics.
func RepoOriginWithAuthorizedSSHKeysFromContext(ctx context.Context) RepoOriginWithAuthorizedSSHKeys {
	s, ok := ctx.Value(_RepoOriginWithAuthorizedSSHKeysKey).(RepoOriginWithAuthorizedSSHKeys)
	if !ok || s == nil {
		panic("no RepoOriginWithAuthorizedSSHKeys set in context")
	}
	return s
}

// RepoOriginWithAuthorizedSSHKeysFromContextOrNil returns the context's RepoOriginWithAuthorizedSSHKeys store if present, or else nil.
func RepoOriginWithAuthorizedSSHKeysFromContextOrNil(ctx context.Context) RepoOriginWithAuthorizedSSHKeys {
	s, ok := ctx.Value(_RepoOriginWithAuthorizedSSHKeysKey).(RepoOriginWithAuthorizedSSHKeys)
	if ok {
		return s
	}
	return nil
}

// WithRepoOriginWithCommitStatuses returns a copy of parent with the given RepoOriginWithCommitStatuses store.
func WithRepoOriginWithCommitStatuses(parent context.Context, s RepoOriginWithCommitStatuses) context.Context {
	return context.WithValue(parent, _RepoOriginWithCommitStatusesKey, s)
}

// RepoOriginWithCommitStatusesFromContext gets the context's RepoOriginWithCommitStatuses store. If the store is not present, it panics.
func RepoOriginWithCommitStatusesFromContext(ctx context.Context) RepoOriginWithCommitStatuses {
	s, ok := ctx.Value(_RepoOriginWithCommitStatusesKey).(RepoOriginWithCommitStatuses)
	if !ok || s == nil {
		panic("no RepoOriginWithCommitStatuses set in context")
	}
	return s
}

// RepoOriginWithCommitStatusesFromContextOrNil returns the context's RepoOriginWithCommitStatuses store if present, or else nil.
func RepoOriginWithCommitStatusesFromContextOrNil(ctx context.Context) RepoOriginWithCommitStatuses {
	s, ok := ctx.Value(_RepoOriginWithCommitStatusesKey).(RepoOriginWithCommitStatuses)
	if ok {
		return s
	}
	return nil
}

// WithRepoOriginWithPushHooks returns a copy of parent with the given RepoOriginWithPushHooks store.
func WithRepoOriginWithPushHooks(parent context.Context, s RepoOriginWithPushHooks) context.Context {
	return context.WithValue(parent, _RepoOriginWithPushHooksKey, s)
}

// RepoOriginWithPushHooksFromContext gets the context's RepoOriginWithPushHooks store. If the store is not present, it panics.
func RepoOriginWithPushHooksFromContext(ctx context.Context) RepoOriginWithPushHooks {
	s, ok := ctx.Value(_RepoOriginWithPushHooksKey).(RepoOriginWithPushHooks)
	if !ok || s == nil {
		panic("no RepoOriginWithPushHooks set in context")
	}
	return s
}

// RepoOriginWithPushHooksFromContextOrNil returns the context's RepoOriginWithPushHooks store if present, or else nil.
func RepoOriginWithPushHooksFromContextOrNil(ctx context.Context) RepoOriginWithPushHooks {
	s, ok := ctx.Value(_RepoOriginWithPushHooksKey).(RepoOriginWithPushHooks)
	if ok {
		return s
	}
	return nil
}

// WithRepoStatuses returns a copy of parent with the given RepoStatuses store.
func WithRepoStatuses(parent context.Context, s RepoStatuses) context.Context {
	return context.WithValue(parent, _RepoStatusesKey, s)
}

// RepoStatusesFromContext gets the context's RepoStatuses store. If the store is not present, it panics.
func RepoStatusesFromContext(ctx context.Context) RepoStatuses {
	s, ok := ctx.Value(_RepoStatusesKey).(RepoStatuses)
	if !ok || s == nil {
		panic("no RepoStatuses set in context")
	}
	return s
}

// RepoStatusesFromContextOrNil returns the context's RepoStatuses store if present, or else nil.
func RepoStatusesFromContextOrNil(ctx context.Context) RepoStatuses {
	s, ok := ctx.Value(_RepoStatusesKey).(RepoStatuses)
	if ok {
		return s
	}
	return nil
}

// WithRepoVCS returns a copy of parent with the given RepoVCS store.
func WithRepoVCS(parent context.Context, s RepoVCS) context.Context {
	return context.WithValue(parent, _RepoVCSKey, s)
}

// RepoVCSFromContext gets the context's RepoVCS store. If the store is not present, it panics.
func RepoVCSFromContext(ctx context.Context) RepoVCS {
	s, ok := ctx.Value(_RepoVCSKey).(RepoVCS)
	if !ok || s == nil {
		panic("no RepoVCS set in context")
	}
	return s
}

// RepoVCSFromContextOrNil returns the context's RepoVCS store if present, or else nil.
func RepoVCSFromContextOrNil(ctx context.Context) RepoVCS {
	s, ok := ctx.Value(_RepoVCSKey).(RepoVCS)
	if ok {
		return s
	}
	return nil
}

// WithRepos returns a copy of parent with the given Repos store.
func WithRepos(parent context.Context, s Repos) context.Context {
	return context.WithValue(parent, _ReposKey, s)
}

// ReposFromContext gets the context's Repos store. If the store is not present, it panics.
func ReposFromContext(ctx context.Context) Repos {
	s, ok := ctx.Value(_ReposKey).(Repos)
	if !ok || s == nil {
		panic("no Repos set in context")
	}
	return s
}

// ReposFromContextOrNil returns the context's Repos store if present, or else nil.
func ReposFromContextOrNil(ctx context.Context) Repos {
	s, ok := ctx.Value(_ReposKey).(Repos)
	if ok {
		return s
	}
	return nil
}

// WithStorage returns a copy of parent with the given Storage store.
func WithStorage(parent context.Context, s Storage) context.Context {
	return context.WithValue(parent, _StorageKey, s)
}

// StorageFromContext gets the context's Storage store. If the store is not present, it panics.
func StorageFromContext(ctx context.Context) Storage {
	s, ok := ctx.Value(_StorageKey).(Storage)
	if !ok || s == nil {
		panic("no Storage set in context")
	}
	return s
}

// StorageFromContextOrNil returns the context's Storage store if present, or else nil.
func StorageFromContextOrNil(ctx context.Context) Storage {
	s, ok := ctx.Value(_StorageKey).(Storage)
	if ok {
		return s
	}
	return nil
}

// WithUserKeys returns a copy of parent with the given UserKeys store.
func WithUserKeys(parent context.Context, s UserKeys) context.Context {
	return context.WithValue(parent, _UserKeysKey, s)
}

// UserKeysFromContext gets the context's UserKeys store. If the store is not present, it panics.
func UserKeysFromContext(ctx context.Context) UserKeys {
	s, ok := ctx.Value(_UserKeysKey).(UserKeys)
	if !ok || s == nil {
		panic("no UserKeys set in context")
	}
	return s
}

// UserKeysFromContextOrNil returns the context's UserKeys store if present, or else nil.
func UserKeysFromContextOrNil(ctx context.Context) UserKeys {
	s, ok := ctx.Value(_UserKeysKey).(UserKeys)
	if ok {
		return s
	}
	return nil
}

// WithUserPermissions returns a copy of parent with the given UserPermissions store.
func WithUserPermissions(parent context.Context, s UserPermissions) context.Context {
	return context.WithValue(parent, _UserPermissionsKey, s)
}

// UserPermissionsFromContext gets the context's UserPermissions store. If the store is not present, it panics.
func UserPermissionsFromContext(ctx context.Context) UserPermissions {
	s, ok := ctx.Value(_UserPermissionsKey).(UserPermissions)
	if !ok || s == nil {
		panic("no UserPermissions set in context")
	}
	return s
}

// UserPermissionsFromContextOrNil returns the context's UserPermissions store if present, or else nil.
func UserPermissionsFromContextOrNil(ctx context.Context) UserPermissions {
	s, ok := ctx.Value(_UserPermissionsKey).(UserPermissions)
	if ok {
		return s
	}
	return nil
}

// WithUsers returns a copy of parent with the given Users store.
func WithUsers(parent context.Context, s Users) context.Context {
	return context.WithValue(parent, _UsersKey, s)
}

// UsersFromContext gets the context's Users store. If the store is not present, it panics.
func UsersFromContext(ctx context.Context) Users {
	s, ok := ctx.Value(_UsersKey).(Users)
	if !ok || s == nil {
		panic("no Users set in context")
	}
	return s
}

// UsersFromContextOrNil returns the context's Users store if present, or else nil.
func UsersFromContextOrNil(ctx context.Context) Users {
	s, ok := ctx.Value(_UsersKey).(Users)
	if ok {
		return s
	}
	return nil
}
