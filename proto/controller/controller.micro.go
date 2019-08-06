// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/controller/controller.proto

package controller

import (
	fmt "fmt"
	global "github.com/dexinq/utils/proto/global"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Controller service

type ControllerService interface {
	RegisterController(ctx context.Context, in *ControllerObj, opts ...client.CallOption) (*global.Response, error)
	UpdateController(ctx context.Context, in *ControllerObj, opts ...client.CallOption) (*global.Response, error)
	GetAppObject(ctx context.Context, in *global.Empty, opts ...client.CallOption) (*ControllerList, error)
}

type controllerService struct {
	c    client.Client
	name string
}

func NewControllerService(name string, c client.Client) ControllerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "controller"
	}
	return &controllerService{
		c:    c,
		name: name,
	}
}

func (c *controllerService) RegisterController(ctx context.Context, in *ControllerObj, opts ...client.CallOption) (*global.Response, error) {
	req := c.c.NewRequest(c.name, "Controller.RegisterController", in)
	out := new(global.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerService) UpdateController(ctx context.Context, in *ControllerObj, opts ...client.CallOption) (*global.Response, error) {
	req := c.c.NewRequest(c.name, "Controller.UpdateController", in)
	out := new(global.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerService) GetAppObject(ctx context.Context, in *global.Empty, opts ...client.CallOption) (*ControllerList, error) {
	req := c.c.NewRequest(c.name, "Controller.GetAppObject", in)
	out := new(ControllerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Controller service

type ControllerHandler interface {
	RegisterController(context.Context, *ControllerObj, *global.Response) error
	UpdateController(context.Context, *ControllerObj, *global.Response) error
	GetAppObject(context.Context, *global.Empty, *ControllerList) error
}

func RegisterControllerHandler(s server.Server, hdlr ControllerHandler, opts ...server.HandlerOption) error {
	type controller interface {
		RegisterController(ctx context.Context, in *ControllerObj, out *global.Response) error
		UpdateController(ctx context.Context, in *ControllerObj, out *global.Response) error
		GetAppObject(ctx context.Context, in *global.Empty, out *ControllerList) error
	}
	type Controller struct {
		controller
	}
	h := &controllerHandler{hdlr}
	return s.Handle(s.NewHandler(&Controller{h}, opts...))
}

type controllerHandler struct {
	ControllerHandler
}

func (h *controllerHandler) RegisterController(ctx context.Context, in *ControllerObj, out *global.Response) error {
	return h.ControllerHandler.RegisterController(ctx, in, out)
}

func (h *controllerHandler) UpdateController(ctx context.Context, in *ControllerObj, out *global.Response) error {
	return h.ControllerHandler.UpdateController(ctx, in, out)
}

func (h *controllerHandler) GetAppObject(ctx context.Context, in *global.Empty, out *ControllerList) error {
	return h.ControllerHandler.GetAppObject(ctx, in, out)
}
