// Code generated by protoc-gen-gogo.
// source: identify.proto
// DO NOT EDIT!

/*
Package identify_pb is a generated protocol buffer package.

It is generated from these files:
	identify.proto

It has these top-level messages:
	Identify
*/
package identify_pb

import proto "github.com/jbenet/go-ipfs/Godeps/_workspace/src/code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Identify struct {
	// protocolVersion determines compatibility between peers
	ProtocolVersion *string `protobuf:"bytes,5,opt,name=protocolVersion" json:"protocolVersion,omitempty"`
	// agentVersion is like a UserAgent string in browsers, or client version in bittorrent
	// includes the client name and client.
	AgentVersion *string `protobuf:"bytes,6,opt,name=agentVersion" json:"agentVersion,omitempty"`
	// publicKey is this node's public key (which also gives its node.ID)
	// - may not need to be sent, as secure channel implies it has been sent.
	// - then again, if we change / disable secure channel, may still want it.
	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey" json:"publicKey,omitempty"`
	// listenAddrs are the multiaddrs the sender node listens for open connections on
	ListenAddrs [][]byte `protobuf:"bytes,2,rep,name=listenAddrs" json:"listenAddrs,omitempty"`
	// oservedAddr is the multiaddr of the remote endpoint that the sender node perceives
	// this is useful information to convey to the other side, as it helps the remote endpoint
	// determine whether its connection to the local peer goes through NAT.
	ObservedAddr []byte `protobuf:"bytes,4,opt,name=observedAddr" json:"observedAddr,omitempty"`
	// protocols are the services this node is running
	Protocols        []string `protobuf:"bytes,3,rep,name=protocols" json:"protocols,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Identify) Reset()         { *m = Identify{} }
func (m *Identify) String() string { return proto.CompactTextString(m) }
func (*Identify) ProtoMessage()    {}

func (m *Identify) GetProtocolVersion() string {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return ""
}

func (m *Identify) GetAgentVersion() string {
	if m != nil && m.AgentVersion != nil {
		return *m.AgentVersion
	}
	return ""
}

func (m *Identify) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Identify) GetListenAddrs() [][]byte {
	if m != nil {
		return m.ListenAddrs
	}
	return nil
}

func (m *Identify) GetObservedAddr() []byte {
	if m != nil {
		return m.ObservedAddr
	}
	return nil
}

func (m *Identify) GetProtocols() []string {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func init() {
}
