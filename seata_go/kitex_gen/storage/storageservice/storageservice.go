// Code generated by Kitex v0.9.1. DO NOT EDIT.

package storageservice

import (
	"context"
	"errors"
	storage "github.com/cloudwego/kitex-examples/seata_go/kitex_gen/storage"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Deduct": kitex.NewMethodInfo(
		deductHandler,
		newStorageServiceDeductArgs,
		newStorageServiceDeductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Calculate": kitex.NewMethodInfo(
		calculateHandler,
		newStorageServiceCalculateArgs,
		newStorageServiceCalculateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	storageServiceServiceInfo                = NewServiceInfo()
	storageServiceServiceInfoForClient       = NewServiceInfoForClient()
	storageServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return storageServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return storageServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return storageServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "StorageService"
	handlerType := (*storage.StorageService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "storage",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func deductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*storage.StorageServiceDeductArgs)

	err := handler.(storage.StorageService).Deduct(ctx, realArg.CommodityCode, realArg.Count)
	if err != nil {
		return err
	}

	return nil
}
func newStorageServiceDeductArgs() interface{} {
	return storage.NewStorageServiceDeductArgs()
}

func newStorageServiceDeductResult() interface{} {
	return storage.NewStorageServiceDeductResult()
}

func calculateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*storage.StorageServiceCalculateArgs)
	realResult := result.(*storage.StorageServiceCalculateResult)
	success, err := handler.(storage.StorageService).Calculate(ctx, realArg.CommodityCode, realArg.Count)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newStorageServiceCalculateArgs() interface{} {
	return storage.NewStorageServiceCalculateArgs()
}

func newStorageServiceCalculateResult() interface{} {
	return storage.NewStorageServiceCalculateResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Deduct(ctx context.Context, commodityCode string, count int32) (err error) {
	var _args storage.StorageServiceDeductArgs
	_args.CommodityCode = commodityCode
	_args.Count = count
	var _result storage.StorageServiceDeductResult
	if err = p.c.Call(ctx, "Deduct", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) Calculate(ctx context.Context, commodityCode string, count int32) (r int32, err error) {
	var _args storage.StorageServiceCalculateArgs
	_args.CommodityCode = commodityCode
	_args.Count = count
	var _result storage.StorageServiceCalculateResult
	if err = p.c.Call(ctx, "Calculate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
