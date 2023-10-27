// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package trace

import (
	"context"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

// Service
type Service struct {
	tracerProvider *sdktrace.TracerProvider

	wg         sync.WaitGroup
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// Type 服务类型
func (s *Service) Type() string {
	return "trace"
}

// newHTTPClient
func (s *Service) newHTTPClient() otlptrace.Client {
	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(strings.Join([]string{otlpHost, otlpPort}, ":")),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithRetry(otlptracehttp.RetryConfig{
			Enabled:         true,
			InitialInterval: time.Nanosecond,
			MaxInterval:     time.Nanosecond,
			MaxElapsedTime:  5,
		}),
	}
	client := otlptracehttp.NewClient(opts...)
	return client
}

// newGrpcClient
func (s *Service) newGrpcClient() otlptrace.Client {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(strings.Join([]string{otlpHost, otlpPort}, ":")),
		otlptracegrpc.WithInsecure(),
	}
	client := otlptracegrpc.NewClient(opts...)
	return client
}

// newResource
func (s *Service) newResource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(ServiceName),
		attribute.Key("bk.data.token").String(otlpToken),
	)
}

// Start
func (s *Service) Start(ctx context.Context) {
	var (
		client   otlptrace.Client
		exporter *otlptrace.Exporter
		err      error
	)

	switch otlpType {
	case "grpc":
		client = s.newGrpcClient()
	default:
		client = s.newHTTPClient()
	}

	exporter, err = otlptrace.New(ctx, client)
	if err != nil {
		return
	}

	// 这里的wg不是用于goroutine运行判断，需要注意
	s.wg.Add(1)
	s.tracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(s.newResource()),
	)
	otel.SetTracerProvider(s.tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	s.ctx, s.cancelFunc = context.WithCancel(ctx)
}

// Reload 重启服务
func (s *Service) Reload(ctx context.Context) {
	s.Close()
	s.Start(ctx)
}

// Close 关闭服务
func (s *Service) Close() {
	if s.tracerProvider == nil {
		return
	}

	go func(tracerProvider *sdktrace.TracerProvider) {
		defer s.wg.Done()
		if err := tracerProvider.Shutdown(s.ctx); err != nil {
			return
		}

	}(s.tracerProvider)
}

// Wait
func (s *Service) Wait() {
	s.wg.Wait()
}
