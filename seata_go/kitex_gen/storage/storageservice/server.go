// Code generated by Kitex v0.9.1. DO NOT EDIT.
package storageservice

import (
	storage "github.com/cloudwego/kitex-examples/seata_go/kitex_gen/storage"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler storage.StorageService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler storage.StorageService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
