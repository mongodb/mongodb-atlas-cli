// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var errMissingEnvVar = errors.New("environment variable missing")

func execute() error {
	attrs := map[string]string{}

	rootCmd := &cobra.Command{
		Use:  "otel <span>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), args[0], attrs)
		},
	}

	rootCmd.Flags().StringToStringVar(&attrs, "attr", nil, "attributes to the span")

	return rootCmd.Execute()
}

func run(ctx context.Context, spanName string, attrs map[string]string) error {
	traceID := os.Getenv("otel_trace_id")
	if traceID == "" {
		return fmt.Errorf("%w: %s", errMissingEnvVar, "otel_trace_id")
	}

	parentID := os.Getenv("otel_parent_id")
	if parentID == "" {
		return fmt.Errorf("%w: %s", errMissingEnvVar, "otel_parent_id")
	}

	collectorEndpoint := os.Getenv("otel_collector_endpoint")
	if collectorEndpoint == "" {
		return fmt.Errorf("%w: %s", errMissingEnvVar, "otel_collector_endpoint")
	}

	projectID := os.Getenv("project_id")
	if collectorEndpoint == "" {
		return fmt.Errorf("%w: %s", errMissingEnvVar, "project_id")
	}

	projectIdentifier := os.Getenv("project_identifier")
	if projectIdentifier == "" {
		return fmt.Errorf("%w: %s", errMissingEnvVar, "project_identifier")
	}

	parsedTraceID, err := trace.TraceIDFromHex(traceID)
	if err != nil {
		return err
	}

	parsedSpanID, err := trace.SpanIDFromHex(parentID)
	if err != nil {
		return err
	}

	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    parsedTraceID,
		SpanID:     parsedSpanID,
		TraceFlags: trace.FlagsSampled,
	})

	traceCtx := trace.ContextWithRemoteSpanContext(ctx, spanContext)

	const serviceName = "mongodb-atlas-cli-master-tools-otel"

	res, err := resource.New(traceCtx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String("1.0.0"),
		),
		resource.WithAttributes(
			attribute.String("evergreen.project.id", projectID),
			attribute.String("evergreen.project.identifier", projectIdentifier)),
	)
	if err != nil {
		return err
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, //nolint:gosec //needed for evg
	}

	conn, err := grpc.NewClient(collectorEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	if err != nil {
		return err
	}
	defer conn.Close()

	otlpExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return err
	}

	processors := []sdktrace.SpanProcessor{sdktrace.NewBatchSpanProcessor(otlpExporter)}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
	)

	defer func() {
		_ = tp.Shutdown(traceCtx)
	}()

	for _, processor := range processors {
		tp.RegisterSpanProcessor(processor)
	}

	otel.SetTextMapPropagator(propagation.TraceContext{})

	otel.SetTracerProvider(tp)

	tracer := otel.Tracer(serviceName)

	_, span := tracer.Start(traceCtx, spanName)
	defer span.End()

	attrList := []attribute.KeyValue{}
	for k, v := range attrs {
		attrList = append(attrList, attribute.String(k, v))
	}

	span.SetAttributes(attrList...)

	span.SetStatus(codes.Ok, "")

	return nil
}

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}
