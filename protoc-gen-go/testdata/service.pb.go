// Code generated by protoc-gen-go.
// source: protoc-gen-go/testdata/service.proto
// DO NOT EDIT!

/*
Package svc is a generated protocol buffer package.

It is generated from these files:
	protoc-gen-go/testdata/service.proto

It has these top-level messages:
	Args
	Return
*/
package svc

import proto "github.com/golang/protobuf/proto"
import math "math"

import "net"
import "net/rpc"
import "github.com/bradhe/go-rpcgen/codec"
import "net/url"
import "net/http"
import "github.com/bradhe/go-rpcgen/webrpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Args struct {
	A                *string `protobuf:"bytes,1,req,name=a" json:"a,omitempty"`
	B                *string `protobuf:"bytes,2,req,name=b" json:"b,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}

func (m *Args) GetA() string {
	if m != nil && m.A != nil {
		return *m.A
	}
	return ""
}

func (m *Args) GetB() string {
	if m != nil && m.B != nil {
		return *m.B
	}
	return ""
}

type Return struct {
	C                *string `protobuf:"bytes,1,req,name=c" json:"c,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Return) Reset()         { *m = Return{} }
func (m *Return) String() string { return proto.CompactTextString(m) }
func (*Return) ProtoMessage()    {}

func (m *Return) GetC() string {
	if m != nil && m.C != nil {
		return *m.C
	}
	return ""
}

func init() {
}

// ConcatService is an interface satisfied by the generated client and
// which must be implemented by the object wrapped by the server.
type ConcatService interface {
	Concat(in *Args, out *Return) error
}

// internal wrapper for type-safe RPC calling
type rpcConcatServiceClient struct {
	*rpc.Client
}

func (this rpcConcatServiceClient) Concat(in *Args, out *Return) error {
	return this.Call("ConcatService.Concat", in, out)
}

// NewConcatServiceClient returns an *rpc.Client wrapper for calling the methods of
// ConcatService remotely.
func NewConcatServiceClient(conn net.Conn) ConcatService {
	return rpcConcatServiceClient{rpc.NewClientWithCodec(codec.NewClientCodec(conn))}
}

// ServeConcatService serves the given ConcatService backend implementation on conn.
func ServeConcatService(conn net.Conn, backend ConcatService) error {
	srv := rpc.NewServer()
	if err := srv.RegisterName("ConcatService", backend); err != nil {
		return err
	}
	srv.ServeCodec(codec.NewServerCodec(conn))
	return nil
}

// DialConcatService returns a ConcatService for calling the ConcatService servince at addr (TCP).
func DialConcatService(addr string) (ConcatService, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return NewConcatServiceClient(conn), nil
}

// ListenAndServeConcatService serves the given ConcatService backend implementation
// on all connections accepted as a result of listening on addr (TCP).
func ListenAndServeConcatService(addr string, backend ConcatService) error {
	clients, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := rpc.NewServer()
	if err := srv.RegisterName("ConcatService", backend); err != nil {
		return err
	}
	for {
		conn, err := clients.Accept()
		if err != nil {
			return err
		}
		go srv.ServeCodec(codec.NewServerCodec(conn))
	}
	panic("unreachable")
}

// ConcatServiceWeb is the web-based RPC version of the interface which
// must be implemented by the object wrapped by the webrpc server.
type ConcatServiceWeb interface {
	Concat(r *http.Request, in *Args, out *Return) error
}

// internal wrapper for type-safe webrpc calling
type rpcConcatServiceWebClient struct {
	remote   *url.URL
	protocol webrpc.Protocol
}

func (this rpcConcatServiceWebClient) Concat(in *Args, out *Return) error {
	return webrpc.Post(this.protocol, this.remote, "/ConcatService/Concat", in, out)
}

// Register a ConcatServiceWeb implementation with the given webrpc ServeMux.
// If mux is nil, the default webrpc.ServeMux is used.
func RegisterConcatServiceWeb(this ConcatServiceWeb, mux webrpc.ServeMux) error {
	if mux == nil {
		mux = webrpc.DefaultServeMux
	}
	if err := mux.Handle("/ConcatService/Concat", func(c *webrpc.Call) error {
		in, out := new(Args), new(Return)
		if err := c.ReadRequest(in); err != nil {
			return err
		}
		if err := this.Concat(c.Request, in, out); err != nil {
			return err
		}
		return c.WriteResponse(out)
	}); err != nil {
		return err
	}
	return nil
}

// NewConcatServiceWebClient returns a webrpc wrapper for calling the methods of ConcatService
// remotely via the web.  The remote URL is the base URL of the webrpc server.
func NewConcatServiceWebClient(pro webrpc.Protocol, remote *url.URL) ConcatService {
	return rpcConcatServiceWebClient{remote, pro}
}
