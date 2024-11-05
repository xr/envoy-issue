package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	extproc "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
)

type service struct{}

func (s *service) Process(stream extproc.ExternalProcessor_ProcessServer) error {
	for {

		select {
		case <-stream.Context().Done():
			if err := stream.Context().Err(); err != nil {
				fmt.Printf("Error reading message from stream: %v\n", err)
				return status.Errorf(codes.Canceled, "context cancelled")
			}
			return nil
		default:
			msg, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("Stream closed by proxy\n")
				return nil
			}

			if err != nil {
				fmt.Printf("Error reading message from stream: %v\n", err)
				return err
			}

			res := &extproc.ProcessingResponse{}

			fmt.Printf("Received request type: %T\n", msg.Request)

			switch requestType := msg.Request.(type) {
			case *extproc.ProcessingRequest_RequestHeaders:
				headers := msg.GetRequestHeaders()
				fmt.Printf("Received headers: %+v\n", headers)

				res.Response = &extproc.ProcessingResponse_RequestHeaders{
					RequestHeaders: &extproc.HeadersResponse{
						Response: &extproc.CommonResponse{
							Status: extproc.CommonResponse_CONTINUE,
						},
					},
				}

			case *extproc.ProcessingRequest_RequestBody:
				fmt.Printf("Received request body size %v, endOfStream: %v\n", len(msg.GetRequestBody().GetBody()), msg.GetRequestBody().GetEndOfStream())
				if msg.GetRequestBody().GetEndOfStream() {
					fmt.Printf("Received end of stream\n")
				}
				res.Response = &extproc.ProcessingResponse_RequestBody{
					RequestBody: &extproc.BodyResponse{
						Response: &extproc.CommonResponse{
							Status: extproc.CommonResponse_CONTINUE,
						},
					},
				}

			case *extproc.ProcessingRequest_RequestTrailers:
				res.Response = &extproc.ProcessingResponse_RequestTrailers{
					RequestTrailers: &extproc.TrailersResponse{},
				}

			case *extproc.ProcessingRequest_ResponseHeaders:
				res.Response = &extproc.ProcessingResponse_ResponseHeaders{
					ResponseHeaders: &extproc.HeadersResponse{
						Response: &extproc.CommonResponse{
							Status: extproc.CommonResponse_CONTINUE,
						},
					},
				}

			case *extproc.ProcessingRequest_ResponseBody:
				res.Response = &extproc.ProcessingResponse_ResponseBody{
					ResponseBody: &extproc.BodyResponse{
						Response: &extproc.CommonResponse{
							Status: extproc.CommonResponse_CONTINUE,
						},
					},
				}

			case *extproc.ProcessingRequest_ResponseTrailers:
				res.Response = &extproc.ProcessingResponse_ResponseTrailers{
					ResponseTrailers: &extproc.TrailersResponse{},
				}

			default:
				return status.Errorf(codes.InvalidArgument, "unknown request type: %T", requestType)
			}

			err = stream.Send(res)
			if err != nil {
				fmt.Printf("Error sending response: %v\n", err)
				return err
			}
		}

	}

}

func main() {
	lis, err := net.Listen("tcp", ":20001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	extproc.RegisterExternalProcessorServer(s, &service{})

	log.Println("Starting gRPC server on port :20001")

	_ = s.Serve(lis)
}
