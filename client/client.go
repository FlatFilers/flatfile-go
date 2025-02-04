// This file was auto-generated by Fern from our API Definition.

package client

import (
	accounts "github.com/FlatFilers/flatfile-go/accounts"
	actions "github.com/FlatFilers/flatfile-go/actions"
	agents "github.com/FlatFilers/flatfile-go/agents"
	apps "github.com/FlatFilers/flatfile-go/apps"
	assistant "github.com/FlatFilers/flatfile-go/assistant"
	commits "github.com/FlatFilers/flatfile-go/commits"
	core "github.com/FlatFilers/flatfile-go/core"
	dataretentionpolicies "github.com/FlatFilers/flatfile-go/dataretentionpolicies"
	documents "github.com/FlatFilers/flatfile-go/documents"
	entitlements "github.com/FlatFilers/flatfile-go/entitlements"
	environments "github.com/FlatFilers/flatfile-go/environments"
	events "github.com/FlatFilers/flatfile-go/events"
	files "github.com/FlatFilers/flatfile-go/files"
	guests "github.com/FlatFilers/flatfile-go/guests"
	internal "github.com/FlatFilers/flatfile-go/internal"
	jobs "github.com/FlatFilers/flatfile-go/jobs"
	mapping "github.com/FlatFilers/flatfile-go/mapping"
	option "github.com/FlatFilers/flatfile-go/option"
	records "github.com/FlatFilers/flatfile-go/records"
	roles "github.com/FlatFilers/flatfile-go/roles"
	secrets "github.com/FlatFilers/flatfile-go/secrets"
	sheets "github.com/FlatFilers/flatfile-go/sheets"
	snapshots "github.com/FlatFilers/flatfile-go/snapshots"
	spaces "github.com/FlatFilers/flatfile-go/spaces"
	users "github.com/FlatFilers/flatfile-go/users"
	versions "github.com/FlatFilers/flatfile-go/versions"
	views "github.com/FlatFilers/flatfile-go/views"
	workbooks "github.com/FlatFilers/flatfile-go/workbooks"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Accounts              *accounts.Client
	Actions               *actions.Client
	Agents                *agents.Client
	Apps                  *apps.Client
	Assistant             *assistant.Client
	Commits               *commits.Client
	DataRetentionPolicies *dataretentionpolicies.Client
	Documents             *documents.Client
	Entitlements          *entitlements.Client
	Environments          *environments.Client
	Events                *events.Client
	Files                 *files.Client
	Guests                *guests.Client
	Jobs                  *jobs.Client
	Mapping               *mapping.Client
	Records               *records.Client
	Roles                 *roles.Client
	Secrets               *secrets.Client
	Sheets                *sheets.Client
	Snapshots             *snapshots.Client
	Spaces                *spaces.Client
	Users                 *users.Client
	Versions              *versions.Client
	Views                 *views.Client
	Workbooks             *workbooks.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:                options.ToHeader(),
		Accounts:              accounts.NewClient(opts...),
		Actions:               actions.NewClient(opts...),
		Agents:                agents.NewClient(opts...),
		Apps:                  apps.NewClient(opts...),
		Assistant:             assistant.NewClient(opts...),
		Commits:               commits.NewClient(opts...),
		DataRetentionPolicies: dataretentionpolicies.NewClient(opts...),
		Documents:             documents.NewClient(opts...),
		Entitlements:          entitlements.NewClient(opts...),
		Environments:          environments.NewClient(opts...),
		Events:                events.NewClient(opts...),
		Files:                 files.NewClient(opts...),
		Guests:                guests.NewClient(opts...),
		Jobs:                  jobs.NewClient(opts...),
		Mapping:               mapping.NewClient(opts...),
		Records:               records.NewClient(opts...),
		Roles:                 roles.NewClient(opts...),
		Secrets:               secrets.NewClient(opts...),
		Sheets:                sheets.NewClient(opts...),
		Snapshots:             snapshots.NewClient(opts...),
		Spaces:                spaces.NewClient(opts...),
		Users:                 users.NewClient(opts...),
		Versions:              versions.NewClient(opts...),
		Views:                 views.NewClient(opts...),
		Workbooks:             workbooks.NewClient(opts...),
	}
}
