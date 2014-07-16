// Copyright (c) 2014 The meeko-poblano-directory AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package methods

import (
	"github.com/meeko/go-meeko/meeko/services/rpc"
	"github.com/salsita/go-poblano/v1/poblano"
)

const MethodPrefix = "PoblanoDirectory@1"

type MethodRegistry interface {
	RegisterMethod(method string, handler rpc.RequestHandler) error
	UnregisterMethod(method string) error
}

type Logger interface {
	Infof(format string, v ...interface{})
	Critical(v ...interface{}) error
}

type Methods struct {
	log      Logger
	client   *poblano.Client
	rpcToken string
}

func New(logger Logger, client *poblano.Client, rpcToken string) *Methods {
	return &Methods{logger, client, rpcToken}
}

func (ms *Methods) getUsers(request rpc.RemoteRequest) {
	// Unmarshal the arguments.
	var args GetUsersArgs
	if err := request.UnmarshalArgs(&args); err != nil {
		ms.resolve(request, 1, &GetUsersReturnValue{Error: err.Error()})
		return
	}

	// Check the RPC token.
	if args.RPCToken != ms.rpcToken {
		ms.resolve(request, 2, &GetUsersReturnValue{Error: "invalid RPC token"})
		return
	}

	// Prepare the HTTP request.
	var url string
	if args.Query == "" {
		url = "/users"
	} else {
		url = "/users?" + args.Query
	}
	req, err := ms.client.NewRequest("GET", url, nil)
	if err != nil {
		ms.resolve(request, 3, &GetUsersReturnValue{Error: err.Error()})
		return
	}

	// Send the HTTP request.
	var users []*data.User
	resp, err := ms.client.Do(req, &users)

	// Resolve the RPC request.
	retVal := &GetUsersReturnValue{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Users:      users,
	}
	if err == nil {
		ms.resolve(request, 0, retVal)
	} else {
		retVal.Error = err.Error()
		ms.resolve(request, 4, retVal)
	}
}

func (ms *Method) getProjects(req rpc.RemoteRequest) {
	// Unmarshal the arguments.
	var args GetProjectsArgs
	if err := request.UnmarshalArgs(&args); err != nil {
		ms.resolve(request, 1, &GetProjectsReturnValue{Error: err.Error()})
		return
	}

	// Check the RPC token.
	if args.RPCToken != ms.rpcToken {
		ms.resolve(request, 2, &GetProjectsReturnValue{Error: "invalid RPC token"})
		return
	}

	// Prepare the HTTP request.
	var url string
	if args.Query == "" {
		url = "/projects"
	} else {
		url = "/projects?" + args.Query
	}
	req, err := ms.client.NewRequest("GET", url, nil)
	if err != nil {
		ms.resolve(request, 3, &GetProjectsReturnValue{Error: err})
		return
	}

	// Send the HTTP request.
	var users []*data.User
	resp, err := ms.client.Do(req, &users)

	// Resolve the RPC request.
	retVal := &GetProjectsReturnValue{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Users:      users,
	}
	if err == nil {
		ms.resolve(request, 0, retVal)
	} else {
		retVal.Error = err.Error()
		ms.resolve(request, 4, retVal)
	}
}

func (ms *Methods) resolve(request rpc.RemoteRequest, retCode rpc.ReturnCode, retValue interface{}) {
	if err := request.Resolve(retCode, retValue); err != nil {
		ms.log.Critical(err)
	}
}

func (ms *Method) Export(registry MethodRegistry) error {
	// Export GetUser.
	userMethod := MethodPrefix + ".GetUsers"
	err := registry.RegisterMethod(userMethod, ms.getUser)
	if err != nil {
		return ms.log.Critical(err)
	}
	ms.log.Infof("Method %v exported", userMethod)

	// Export GetProject.
	projectMethod := MethodPrefix + ".GetProjects"
	err = registry.RegisterMethod(projectMethod, ms.getProject)
	if err != nil {
		ms.log.Critical(err)
		if ex := registry.UnregisterMethod(userMethod); ex != nil {
			ms.log.Critical(ex)
		}
		return err
	}
	ms.log.Infof("Method %v exported", projectMethod)

	return nil
}
