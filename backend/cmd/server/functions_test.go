package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.nhat.io/grpcmock"
	"google.golang.org/grpc/codes"
	proto_bookAuthor_service "testKwd/proto"
	"testing"
)

type TestCase struct {
	scenario         string
	request          interface{}
	expectedResponse interface{}
	response         interface{}
	responseOut      interface{}
	expectedMsgError string
	MsgError         string
	ErrorStatus      codes.Code
	method           string
}

const (
	Success    = "success"
	Error      = "error"
	getBooks   = "/pb.bookAuthorService/GetBooks"
	getAuthors = "/pb.bookAuthorService/GetAuthors"
)

func createServiceServerMock(testCase TestCase) grpcmock.ServerMockerWithContextDialer {
	opts := grpcmock.RegisterService(proto_bookAuthor_service.RegisterBookAuthorServiceServer)
	if testCase.MsgError != "" {
		return grpcmock.MockServerWithBufConn(opts, func(s *grpcmock.Server) {
			s.ExpectUnary(testCase.method).
				WithPayload(testCase.request).
				ReturnError(testCase.ErrorStatus, testCase.MsgError)
		})
	}

	return grpcmock.MockServerWithBufConn(opts, func(s *grpcmock.Server) {
		s.ExpectUnary(testCase.method).
			WithPayload(testCase.request).
			Return(testCase.response)
	})
}

func TestConcert(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{
			scenario: Success,
			method:   getBooks,
			request:  &proto_bookAuthor_service.GetBookRequest{AuthorName: "AuthorName"},
			expectedResponse: &proto_bookAuthor_service.GetBookResponse{Books: []*proto_bookAuthor_service.Book{
				{
					BookName: "testName",
				}, {
					BookName: "testName2",
				},
			}},
			response: &proto_bookAuthor_service.GetBookResponse{Books: []*proto_bookAuthor_service.Book{
				{
					BookName: "testName",
				}, {
					BookName: "testName2",
				},
			}},
			responseOut: &proto_bookAuthor_service.GetBookResponse{Books: []*proto_bookAuthor_service.Book{{}}},
		},
		{
			scenario: Success,
			method:   getAuthors,
			request:  &proto_bookAuthor_service.GetAuthorRequest{BookName: "BookName"},
			expectedResponse: &proto_bookAuthor_service.GetAuthorResponse{Authors: []*proto_bookAuthor_service.Author{
				{
					AuthorName: "testName",
				}, {
					AuthorName: "testName2",
				},
			}},
			response: &proto_bookAuthor_service.GetAuthorResponse{Authors: []*proto_bookAuthor_service.Author{
				{
					AuthorName: "testName",
				}, {
					AuthorName: "testName2",
				},
			}},
			responseOut: &proto_bookAuthor_service.GetAuthorResponse{Authors: []*proto_bookAuthor_service.Author{{}}},
		},
		{
			scenario:         Error,
			method:           getBooks,
			request:          &proto_bookAuthor_service.GetBookRequest{AuthorName: "!"},
			ErrorStatus:      404,
			expectedMsgError: "rpc error: code = Code(404) desc = Record not found",
			MsgError:         "Record not found",
		},
		{
			scenario:         Error,
			method:           getAuthors,
			request:          &proto_bookAuthor_service.GetAuthorRequest{BookName: "!"},
			ErrorStatus:      404,
			expectedMsgError: "rpc error: code = Code(404) desc = Record not found",
			MsgError:         "Record not found",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			_, dialer := createServiceServerMock(tc)(t)

			err := grpcmock.InvokeUnary(context.Background(),
				tc.method, tc.request, tc.responseOut,
				grpcmock.WithInsecure(),
				grpcmock.WithContextDialer(dialer),
			)

			if tc.MsgError != "" {
				t.Log(tc.expectedMsgError)
				t.Log(err)

				assert.EqualError(t, err, tc.expectedMsgError)
				return
			}

			require.NoError(t, err)

			t.Log(tc.expectedResponse)
			t.Log(tc.responseOut)

			assert.Equal(t, tc.expectedResponse, tc.responseOut)
		})
	}
}
