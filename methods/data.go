// Copyright (c) 2014 The meeko-poblano-directory AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package methods

import "github.com/salsita/go-poblano/v1/poblano/data"

type GetUsersArgs struct {
	RPCToken string `codec:"rpc_token,omitempty"`
	Query    string `codec:"query,omitempty"`
}

type GetUsersReturnValue struct {
	Status     string             `codec:"status,omitempty"`
	StatusCode int                `codec:"status_code,omitempty"`
	Users      []*data.UserRecord `codec:"users,omitempty"`
	Error      string             `codec:"error,omitempty"`
}

type GetProjectsArgs struct {
	RPCToken string `codec:"rpc_token,omitempty"`
	Query    string `codec:"query,omitempty"`
}

type GetProjectsReturnValue struct {
	Status     string                `codec:"status,omitempty"`
	StatusCode int                   `codec:"status_code,omitempty"`
	Projects   []*data.ProjectRecord `codec:"projects,omitempty"`
	Error      string                `codec:"error,omitempty"`
}
