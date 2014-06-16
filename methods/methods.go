// Copyright (c) 2014 The meeko-poblano-directory AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package methods

import (
	"github.com/meeko/go-meeko/meeko/services/rpc"
	"github.com/salsita/go-poblano/v1/poblano"
)

const MethodPrefix = "Poblano@1"

type MethodRegistry interface {
	RegisterMethod(method string, handler rpc.RequestHandler) error
	UnregisterMethod(method string) error
}

type Logger interface {
	Infof(format string, v ...interface{})
	Critical(v ...interface{}) error
}

type Methods struct {
	log    Logger
	client *poblano.Client
}

func New(logger Logger, client *poblano.Client) *Methods {
	return &Methods{logger, client}
}

func (ms *Methods) getUsers(request rpc.RemoteRequest) {
	// Unmarshal the arguments.
	var args GetUsersArgs
	if err := request.UnmarshalArgs(&args); err != nil {
		ms.resolve(request, 1, &GetUsersReturnValue{Error: err})
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
		ms.resolve(request, 2, &GetUsersReturnValue{Error: err})
		return
	}

	// Send the HTTP request.
	var users []*data.User
	resp, err := ms.client.Do(req, &users)

	// Resolve the RPC request.
	ms.resolve(request, 0, &GetUsersReturnValue{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Users:      users,
		Error:      err,
	})
}

func (ms *Method) getProjects(req rpc.RemoteRequest) {
	// Unmarshal the arguments.
	var args GetProjectsArgs
	if err := request.UnmarshalArgs(&args); err != nil {
		ms.resolve(request, 1, &GetProjectsReturnValue{Error: err})
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
		ms.resolve(request, 2, &GetProjectsReturnValue{Error: err})
		return
	}

	// Send the HTTP request.
	var users []*data.User
	resp, err := ms.client.Do(req, &users)

	// Resolve the RPC request.
	ms.resolve(request, 0, &GetProjectsReturnValue{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Users:      users,
		Error:      err,
	})
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
