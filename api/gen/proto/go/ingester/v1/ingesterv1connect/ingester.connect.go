// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ingester/v1/ingester.proto

package ingesterv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v12 "github.com/grafana/pyroscope/api/gen/proto/go/ingester/v1"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/push/v1"
	v11 "github.com/grafana/pyroscope/api/gen/proto/go/types/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// IngesterServiceName is the fully-qualified name of the IngesterService service.
	IngesterServiceName = "ingester.v1.IngesterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IngesterServicePushProcedure is the fully-qualified name of the IngesterService's Push RPC.
	IngesterServicePushProcedure = "/ingester.v1.IngesterService/Push"
	// IngesterServiceLabelValuesProcedure is the fully-qualified name of the IngesterService's
	// LabelValues RPC.
	IngesterServiceLabelValuesProcedure = "/ingester.v1.IngesterService/LabelValues"
	// IngesterServiceLabelNamesProcedure is the fully-qualified name of the IngesterService's
	// LabelNames RPC.
	IngesterServiceLabelNamesProcedure = "/ingester.v1.IngesterService/LabelNames"
	// IngesterServiceProfileTypesProcedure is the fully-qualified name of the IngesterService's
	// ProfileTypes RPC.
	IngesterServiceProfileTypesProcedure = "/ingester.v1.IngesterService/ProfileTypes"
	// IngesterServiceSeriesProcedure is the fully-qualified name of the IngesterService's Series RPC.
	IngesterServiceSeriesProcedure = "/ingester.v1.IngesterService/Series"
	// IngesterServiceFlushProcedure is the fully-qualified name of the IngesterService's Flush RPC.
	IngesterServiceFlushProcedure = "/ingester.v1.IngesterService/Flush"
	// IngesterServiceMergeProfilesStacktracesProcedure is the fully-qualified name of the
	// IngesterService's MergeProfilesStacktraces RPC.
	IngesterServiceMergeProfilesStacktracesProcedure = "/ingester.v1.IngesterService/MergeProfilesStacktraces"
	// IngesterServiceMergeProfilesLabelsProcedure is the fully-qualified name of the IngesterService's
	// MergeProfilesLabels RPC.
	IngesterServiceMergeProfilesLabelsProcedure = "/ingester.v1.IngesterService/MergeProfilesLabels"
	// IngesterServiceMergeProfilesPprofProcedure is the fully-qualified name of the IngesterService's
	// MergeProfilesPprof RPC.
	IngesterServiceMergeProfilesPprofProcedure = "/ingester.v1.IngesterService/MergeProfilesPprof"
	// IngesterServiceMergeSpanProfileProcedure is the fully-qualified name of the IngesterService's
	// MergeSpanProfile RPC.
	IngesterServiceMergeSpanProfileProcedure = "/ingester.v1.IngesterService/MergeSpanProfile"
	// IngesterServiceBlockMetadataProcedure is the fully-qualified name of the IngesterService's
	// BlockMetadata RPC.
	IngesterServiceBlockMetadataProcedure = "/ingester.v1.IngesterService/BlockMetadata"
)

// IngesterServiceClient is a client for the ingester.v1.IngesterService service.
type IngesterServiceClient interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
	LabelValues(context.Context, *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error)
	LabelNames(context.Context, *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error)
	ProfileTypes(context.Context, *connect_go.Request[v12.ProfileTypesRequest]) (*connect_go.Response[v12.ProfileTypesResponse], error)
	Series(context.Context, *connect_go.Request[v12.SeriesRequest]) (*connect_go.Response[v12.SeriesResponse], error)
	Flush(context.Context, *connect_go.Request[v12.FlushRequest]) (*connect_go.Response[v12.FlushResponse], error)
	MergeProfilesStacktraces(context.Context) *connect_go.BidiStreamForClient[v12.MergeProfilesStacktracesRequest, v12.MergeProfilesStacktracesResponse]
	MergeProfilesLabels(context.Context) *connect_go.BidiStreamForClient[v12.MergeProfilesLabelsRequest, v12.MergeProfilesLabelsResponse]
	MergeProfilesPprof(context.Context) *connect_go.BidiStreamForClient[v12.MergeProfilesPprofRequest, v12.MergeProfilesPprofResponse]
	MergeSpanProfile(context.Context) *connect_go.BidiStreamForClient[v12.MergeSpanProfileRequest, v12.MergeSpanProfileResponse]
	BlockMetadata(context.Context, *connect_go.Request[v12.BlockMetadataRequest]) (*connect_go.Response[v12.BlockMetadataResponse], error)
}

// NewIngesterServiceClient constructs a client for the ingester.v1.IngesterService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIngesterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IngesterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ingesterServiceClient{
		push: connect_go.NewClient[v1.PushRequest, v1.PushResponse](
			httpClient,
			baseURL+IngesterServicePushProcedure,
			opts...,
		),
		labelValues: connect_go.NewClient[v11.LabelValuesRequest, v11.LabelValuesResponse](
			httpClient,
			baseURL+IngesterServiceLabelValuesProcedure,
			opts...,
		),
		labelNames: connect_go.NewClient[v11.LabelNamesRequest, v11.LabelNamesResponse](
			httpClient,
			baseURL+IngesterServiceLabelNamesProcedure,
			opts...,
		),
		profileTypes: connect_go.NewClient[v12.ProfileTypesRequest, v12.ProfileTypesResponse](
			httpClient,
			baseURL+IngesterServiceProfileTypesProcedure,
			opts...,
		),
		series: connect_go.NewClient[v12.SeriesRequest, v12.SeriesResponse](
			httpClient,
			baseURL+IngesterServiceSeriesProcedure,
			opts...,
		),
		flush: connect_go.NewClient[v12.FlushRequest, v12.FlushResponse](
			httpClient,
			baseURL+IngesterServiceFlushProcedure,
			opts...,
		),
		mergeProfilesStacktraces: connect_go.NewClient[v12.MergeProfilesStacktracesRequest, v12.MergeProfilesStacktracesResponse](
			httpClient,
			baseURL+IngesterServiceMergeProfilesStacktracesProcedure,
			opts...,
		),
		mergeProfilesLabels: connect_go.NewClient[v12.MergeProfilesLabelsRequest, v12.MergeProfilesLabelsResponse](
			httpClient,
			baseURL+IngesterServiceMergeProfilesLabelsProcedure,
			opts...,
		),
		mergeProfilesPprof: connect_go.NewClient[v12.MergeProfilesPprofRequest, v12.MergeProfilesPprofResponse](
			httpClient,
			baseURL+IngesterServiceMergeProfilesPprofProcedure,
			opts...,
		),
		mergeSpanProfile: connect_go.NewClient[v12.MergeSpanProfileRequest, v12.MergeSpanProfileResponse](
			httpClient,
			baseURL+IngesterServiceMergeSpanProfileProcedure,
			opts...,
		),
		blockMetadata: connect_go.NewClient[v12.BlockMetadataRequest, v12.BlockMetadataResponse](
			httpClient,
			baseURL+IngesterServiceBlockMetadataProcedure,
			opts...,
		),
	}
}

// ingesterServiceClient implements IngesterServiceClient.
type ingesterServiceClient struct {
	push                     *connect_go.Client[v1.PushRequest, v1.PushResponse]
	labelValues              *connect_go.Client[v11.LabelValuesRequest, v11.LabelValuesResponse]
	labelNames               *connect_go.Client[v11.LabelNamesRequest, v11.LabelNamesResponse]
	profileTypes             *connect_go.Client[v12.ProfileTypesRequest, v12.ProfileTypesResponse]
	series                   *connect_go.Client[v12.SeriesRequest, v12.SeriesResponse]
	flush                    *connect_go.Client[v12.FlushRequest, v12.FlushResponse]
	mergeProfilesStacktraces *connect_go.Client[v12.MergeProfilesStacktracesRequest, v12.MergeProfilesStacktracesResponse]
	mergeProfilesLabels      *connect_go.Client[v12.MergeProfilesLabelsRequest, v12.MergeProfilesLabelsResponse]
	mergeProfilesPprof       *connect_go.Client[v12.MergeProfilesPprofRequest, v12.MergeProfilesPprofResponse]
	mergeSpanProfile         *connect_go.Client[v12.MergeSpanProfileRequest, v12.MergeSpanProfileResponse]
	blockMetadata            *connect_go.Client[v12.BlockMetadataRequest, v12.BlockMetadataResponse]
}

// Push calls ingester.v1.IngesterService.Push.
func (c *ingesterServiceClient) Push(ctx context.Context, req *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// LabelValues calls ingester.v1.IngesterService.LabelValues.
func (c *ingesterServiceClient) LabelValues(ctx context.Context, req *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error) {
	return c.labelValues.CallUnary(ctx, req)
}

// LabelNames calls ingester.v1.IngesterService.LabelNames.
func (c *ingesterServiceClient) LabelNames(ctx context.Context, req *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error) {
	return c.labelNames.CallUnary(ctx, req)
}

// ProfileTypes calls ingester.v1.IngesterService.ProfileTypes.
func (c *ingesterServiceClient) ProfileTypes(ctx context.Context, req *connect_go.Request[v12.ProfileTypesRequest]) (*connect_go.Response[v12.ProfileTypesResponse], error) {
	return c.profileTypes.CallUnary(ctx, req)
}

// Series calls ingester.v1.IngesterService.Series.
func (c *ingesterServiceClient) Series(ctx context.Context, req *connect_go.Request[v12.SeriesRequest]) (*connect_go.Response[v12.SeriesResponse], error) {
	return c.series.CallUnary(ctx, req)
}

// Flush calls ingester.v1.IngesterService.Flush.
func (c *ingesterServiceClient) Flush(ctx context.Context, req *connect_go.Request[v12.FlushRequest]) (*connect_go.Response[v12.FlushResponse], error) {
	return c.flush.CallUnary(ctx, req)
}

// MergeProfilesStacktraces calls ingester.v1.IngesterService.MergeProfilesStacktraces.
func (c *ingesterServiceClient) MergeProfilesStacktraces(ctx context.Context) *connect_go.BidiStreamForClient[v12.MergeProfilesStacktracesRequest, v12.MergeProfilesStacktracesResponse] {
	return c.mergeProfilesStacktraces.CallBidiStream(ctx)
}

// MergeProfilesLabels calls ingester.v1.IngesterService.MergeProfilesLabels.
func (c *ingesterServiceClient) MergeProfilesLabels(ctx context.Context) *connect_go.BidiStreamForClient[v12.MergeProfilesLabelsRequest, v12.MergeProfilesLabelsResponse] {
	return c.mergeProfilesLabels.CallBidiStream(ctx)
}

// MergeProfilesPprof calls ingester.v1.IngesterService.MergeProfilesPprof.
func (c *ingesterServiceClient) MergeProfilesPprof(ctx context.Context) *connect_go.BidiStreamForClient[v12.MergeProfilesPprofRequest, v12.MergeProfilesPprofResponse] {
	return c.mergeProfilesPprof.CallBidiStream(ctx)
}

// MergeSpanProfile calls ingester.v1.IngesterService.MergeSpanProfile.
func (c *ingesterServiceClient) MergeSpanProfile(ctx context.Context) *connect_go.BidiStreamForClient[v12.MergeSpanProfileRequest, v12.MergeSpanProfileResponse] {
	return c.mergeSpanProfile.CallBidiStream(ctx)
}

// BlockMetadata calls ingester.v1.IngesterService.BlockMetadata.
func (c *ingesterServiceClient) BlockMetadata(ctx context.Context, req *connect_go.Request[v12.BlockMetadataRequest]) (*connect_go.Response[v12.BlockMetadataResponse], error) {
	return c.blockMetadata.CallUnary(ctx, req)
}

// IngesterServiceHandler is an implementation of the ingester.v1.IngesterService service.
type IngesterServiceHandler interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
	LabelValues(context.Context, *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error)
	LabelNames(context.Context, *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error)
	ProfileTypes(context.Context, *connect_go.Request[v12.ProfileTypesRequest]) (*connect_go.Response[v12.ProfileTypesResponse], error)
	Series(context.Context, *connect_go.Request[v12.SeriesRequest]) (*connect_go.Response[v12.SeriesResponse], error)
	Flush(context.Context, *connect_go.Request[v12.FlushRequest]) (*connect_go.Response[v12.FlushResponse], error)
	MergeProfilesStacktraces(context.Context, *connect_go.BidiStream[v12.MergeProfilesStacktracesRequest, v12.MergeProfilesStacktracesResponse]) error
	MergeProfilesLabels(context.Context, *connect_go.BidiStream[v12.MergeProfilesLabelsRequest, v12.MergeProfilesLabelsResponse]) error
	MergeProfilesPprof(context.Context, *connect_go.BidiStream[v12.MergeProfilesPprofRequest, v12.MergeProfilesPprofResponse]) error
	MergeSpanProfile(context.Context, *connect_go.BidiStream[v12.MergeSpanProfileRequest, v12.MergeSpanProfileResponse]) error
	BlockMetadata(context.Context, *connect_go.Request[v12.BlockMetadataRequest]) (*connect_go.Response[v12.BlockMetadataResponse], error)
}

// NewIngesterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIngesterServiceHandler(svc IngesterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	ingesterServicePushHandler := connect_go.NewUnaryHandler(
		IngesterServicePushProcedure,
		svc.Push,
		opts...,
	)
	ingesterServiceLabelValuesHandler := connect_go.NewUnaryHandler(
		IngesterServiceLabelValuesProcedure,
		svc.LabelValues,
		opts...,
	)
	ingesterServiceLabelNamesHandler := connect_go.NewUnaryHandler(
		IngesterServiceLabelNamesProcedure,
		svc.LabelNames,
		opts...,
	)
	ingesterServiceProfileTypesHandler := connect_go.NewUnaryHandler(
		IngesterServiceProfileTypesProcedure,
		svc.ProfileTypes,
		opts...,
	)
	ingesterServiceSeriesHandler := connect_go.NewUnaryHandler(
		IngesterServiceSeriesProcedure,
		svc.Series,
		opts...,
	)
	ingesterServiceFlushHandler := connect_go.NewUnaryHandler(
		IngesterServiceFlushProcedure,
		svc.Flush,
		opts...,
	)
	ingesterServiceMergeProfilesStacktracesHandler := connect_go.NewBidiStreamHandler(
		IngesterServiceMergeProfilesStacktracesProcedure,
		svc.MergeProfilesStacktraces,
		opts...,
	)
	ingesterServiceMergeProfilesLabelsHandler := connect_go.NewBidiStreamHandler(
		IngesterServiceMergeProfilesLabelsProcedure,
		svc.MergeProfilesLabels,
		opts...,
	)
	ingesterServiceMergeProfilesPprofHandler := connect_go.NewBidiStreamHandler(
		IngesterServiceMergeProfilesPprofProcedure,
		svc.MergeProfilesPprof,
		opts...,
	)
	ingesterServiceMergeSpanProfileHandler := connect_go.NewBidiStreamHandler(
		IngesterServiceMergeSpanProfileProcedure,
		svc.MergeSpanProfile,
		opts...,
	)
	ingesterServiceBlockMetadataHandler := connect_go.NewUnaryHandler(
		IngesterServiceBlockMetadataProcedure,
		svc.BlockMetadata,
		opts...,
	)
	return "/ingester.v1.IngesterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IngesterServicePushProcedure:
			ingesterServicePushHandler.ServeHTTP(w, r)
		case IngesterServiceLabelValuesProcedure:
			ingesterServiceLabelValuesHandler.ServeHTTP(w, r)
		case IngesterServiceLabelNamesProcedure:
			ingesterServiceLabelNamesHandler.ServeHTTP(w, r)
		case IngesterServiceProfileTypesProcedure:
			ingesterServiceProfileTypesHandler.ServeHTTP(w, r)
		case IngesterServiceSeriesProcedure:
			ingesterServiceSeriesHandler.ServeHTTP(w, r)
		case IngesterServiceFlushProcedure:
			ingesterServiceFlushHandler.ServeHTTP(w, r)
		case IngesterServiceMergeProfilesStacktracesProcedure:
			ingesterServiceMergeProfilesStacktracesHandler.ServeHTTP(w, r)
		case IngesterServiceMergeProfilesLabelsProcedure:
			ingesterServiceMergeProfilesLabelsHandler.ServeHTTP(w, r)
		case IngesterServiceMergeProfilesPprofProcedure:
			ingesterServiceMergeProfilesPprofHandler.ServeHTTP(w, r)
		case IngesterServiceMergeSpanProfileProcedure:
			ingesterServiceMergeSpanProfileHandler.ServeHTTP(w, r)
		case IngesterServiceBlockMetadataProcedure:
			ingesterServiceBlockMetadataHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIngesterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIngesterServiceHandler struct{}

func (UnimplementedIngesterServiceHandler) Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.Push is not implemented"))
}

func (UnimplementedIngesterServiceHandler) LabelValues(context.Context, *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.LabelValues is not implemented"))
}

func (UnimplementedIngesterServiceHandler) LabelNames(context.Context, *connect_go.Request[v11.LabelNamesRequest]) (*connect_go.Response[v11.LabelNamesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.LabelNames is not implemented"))
}

func (UnimplementedIngesterServiceHandler) ProfileTypes(context.Context, *connect_go.Request[v12.ProfileTypesRequest]) (*connect_go.Response[v12.ProfileTypesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.ProfileTypes is not implemented"))
}

func (UnimplementedIngesterServiceHandler) Series(context.Context, *connect_go.Request[v12.SeriesRequest]) (*connect_go.Response[v12.SeriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.Series is not implemented"))
}

func (UnimplementedIngesterServiceHandler) Flush(context.Context, *connect_go.Request[v12.FlushRequest]) (*connect_go.Response[v12.FlushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.Flush is not implemented"))
}

func (UnimplementedIngesterServiceHandler) MergeProfilesStacktraces(context.Context, *connect_go.BidiStream[v12.MergeProfilesStacktracesRequest, v12.MergeProfilesStacktracesResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.MergeProfilesStacktraces is not implemented"))
}

func (UnimplementedIngesterServiceHandler) MergeProfilesLabels(context.Context, *connect_go.BidiStream[v12.MergeProfilesLabelsRequest, v12.MergeProfilesLabelsResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.MergeProfilesLabels is not implemented"))
}

func (UnimplementedIngesterServiceHandler) MergeProfilesPprof(context.Context, *connect_go.BidiStream[v12.MergeProfilesPprofRequest, v12.MergeProfilesPprofResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.MergeProfilesPprof is not implemented"))
}

func (UnimplementedIngesterServiceHandler) MergeSpanProfile(context.Context, *connect_go.BidiStream[v12.MergeSpanProfileRequest, v12.MergeSpanProfileResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.MergeSpanProfile is not implemented"))
}

func (UnimplementedIngesterServiceHandler) BlockMetadata(context.Context, *connect_go.Request[v12.BlockMetadataRequest]) (*connect_go.Response[v12.BlockMetadataResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.BlockMetadata is not implemented"))
}
