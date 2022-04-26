// Package kubernetes provides a micro service using k8s registry plugin
package kubernetes

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/gotoblink/go-plugins/registry/kubernetes"

	// static selector offloads load balancing to k8s services
	// note: requires user to create k8s services
	"github.com/gotoblink/go-plugins/client/selector/static"
)

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// set the registry and selector
	options := []micro.Option{
		micro.Registry(r),
		micro.Selector(s),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return grpc.NewService(options...)
}
