// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/appmanager/appmanager.proto

package appmanager

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

// Client API for AppManager service

type AppManagerService interface {
	GetApp(ctx context.Context, in *global.Empty, opts ...client.CallOption) (*AppList, error)
	AddApp(ctx context.Context, in *AppEntity, opts ...client.CallOption) (*global.Response, error)
	ModifyApp(ctx context.Context, in *AppEntity, opts ...client.CallOption) (*global.Response, error)
}

type appManagerService struct {
	c    client.Client
	name string
}

func NewAppManagerService(name string, c client.Client) AppManagerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "appmanager"
	}
	return &appManagerService{
		c:    c,
		name: name,
	}
}

func (c *appManagerService) GetApp(ctx context.Context, in *global.Empty, opts ...client.CallOption) (*AppList, error) {
	req := c.c.NewRequest(c.name, "AppManager.GetApp", in)
	out := new(AppList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerService) AddApp(ctx context.Context, in *AppEntity, opts ...client.CallOption) (*global.Response, error) {
	req := c.c.NewRequest(c.name, "AppManager.AddApp", in)
	out := new(global.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerService) ModifyApp(ctx context.Context, in *AppEntity, opts ...client.CallOption) (*global.Response, error) {
	req := c.c.NewRequest(c.name, "AppManager.ModifyApp", in)
	out := new(global.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppManager service

type AppManagerHandler interface {
	GetApp(context.Context, *global.Empty, *AppList) error
	AddApp(context.Context, *AppEntity, *global.Response) error
	ModifyApp(context.Context, *AppEntity, *global.Response) error
}

func RegisterAppManagerHandler(s server.Server, hdlr AppManagerHandler, opts ...server.HandlerOption) error {
	type appManager interface {
		GetApp(ctx context.Context, in *global.Empty, out *AppList) error
		AddApp(ctx context.Context, in *AppEntity, out *global.Response) error
		ModifyApp(ctx context.Context, in *AppEntity, out *global.Response) error
	}
	type AppManager struct {
		appManager
	}
	h := &appManagerHandler{hdlr}
	return s.Handle(s.NewHandler(&AppManager{h}, opts...))
}

type appManagerHandler struct {
	AppManagerHandler
}

func (h *appManagerHandler) GetApp(ctx context.Context, in *global.Empty, out *AppList) error {
	return h.AppManagerHandler.GetApp(ctx, in, out)
}

func (h *appManagerHandler) AddApp(ctx context.Context, in *AppEntity, out *global.Response) error {
	return h.AppManagerHandler.AddApp(ctx, in, out)
}

func (h *appManagerHandler) ModifyApp(ctx context.Context, in *AppEntity, out *global.Response) error {
	return h.AppManagerHandler.ModifyApp(ctx, in, out)
}
