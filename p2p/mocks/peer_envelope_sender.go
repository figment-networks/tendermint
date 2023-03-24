// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	log "github.com/tendermint/tendermint/libs/log"
	conn "github.com/tendermint/tendermint/p2p/conn"

	mock "github.com/stretchr/testify/mock"

	net "net"

	p2p "github.com/tendermint/tendermint/p2p"
)

// PeerEnvelopeSender is an autogenerated mock type for the PeerEnvelopeSender type
type PeerEnvelopeSender struct {
	mock.Mock
}

// CloseConn provides a mock function with given fields:
func (_m *PeerEnvelopeSender) CloseConn() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushStop provides a mock function with given fields:
func (_m *PeerEnvelopeSender) FlushStop() {
	_m.Called()
}

// Get provides a mock function with given fields: _a0
func (_m *PeerEnvelopeSender) Get(_a0 string) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetRemovalFailed provides a mock function with given fields:
func (_m *PeerEnvelopeSender) GetRemovalFailed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *PeerEnvelopeSender) ID() p2p.ID {
	ret := _m.Called()

	var r0 p2p.ID
	if rf, ok := ret.Get(0).(func() p2p.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(p2p.ID)
	}

	return r0
}

// IsOutbound provides a mock function with given fields:
func (_m *PeerEnvelopeSender) IsOutbound() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsPersistent provides a mock function with given fields:
func (_m *PeerEnvelopeSender) IsPersistent() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRunning provides a mock function with given fields:
func (_m *PeerEnvelopeSender) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsSidecarPeer provides a mock function with given fields:
func (_m *PeerEnvelopeSender) IsSidecarPeer() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NodeInfo provides a mock function with given fields:
func (_m *PeerEnvelopeSender) NodeInfo() p2p.NodeInfo {
	ret := _m.Called()

	var r0 p2p.NodeInfo
	if rf, ok := ret.Get(0).(func() p2p.NodeInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeInfo)
		}
	}

	return r0
}

// OnReset provides a mock function with given fields:
func (_m *PeerEnvelopeSender) OnReset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStart provides a mock function with given fields:
func (_m *PeerEnvelopeSender) OnStart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStop provides a mock function with given fields:
func (_m *PeerEnvelopeSender) OnStop() {
	_m.Called()
}

// Quit provides a mock function with given fields:
func (_m *PeerEnvelopeSender) Quit() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// RemoteAddr provides a mock function with given fields:
func (_m *PeerEnvelopeSender) RemoteAddr() net.Addr {
	ret := _m.Called()

	var r0 net.Addr
	if rf, ok := ret.Get(0).(func() net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Addr)
		}
	}

	return r0
}

// RemoteIP provides a mock function with given fields:
func (_m *PeerEnvelopeSender) RemoteIP() net.IP {
	ret := _m.Called()

	var r0 net.IP
	if rf, ok := ret.Get(0).(func() net.IP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.IP)
		}
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *PeerEnvelopeSender) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0, _a1
func (_m *PeerEnvelopeSender) Send(_a0 byte, _a1 []byte) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(byte, []byte) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SendEnvelope provides a mock function with given fields: _a0
func (_m *PeerEnvelopeSender) SendEnvelope(_a0 p2p.Envelope) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Envelope) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *PeerEnvelopeSender) Set(_a0 string, _a1 interface{}) {
	_m.Called(_a0, _a1)
}

// SetLogger provides a mock function with given fields: _a0
func (_m *PeerEnvelopeSender) SetLogger(_a0 log.Logger) {
	_m.Called(_a0)
}

// SetRemovalFailed provides a mock function with given fields:
func (_m *PeerEnvelopeSender) SetRemovalFailed() {
	_m.Called()
}

// SocketAddr provides a mock function with given fields:
func (_m *PeerEnvelopeSender) SocketAddr() *p2p.NetAddress {
	ret := _m.Called()

	var r0 *p2p.NetAddress
	if rf, ok := ret.Get(0).(func() *p2p.NetAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.NetAddress)
		}
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *PeerEnvelopeSender) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *PeerEnvelopeSender) Status() conn.ConnectionStatus {
	ret := _m.Called()

	var r0 conn.ConnectionStatus
	if rf, ok := ret.Get(0).(func() conn.ConnectionStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(conn.ConnectionStatus)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *PeerEnvelopeSender) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *PeerEnvelopeSender) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TrySend provides a mock function with given fields: _a0, _a1
func (_m *PeerEnvelopeSender) TrySend(_a0 byte, _a1 []byte) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(byte, []byte) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TrySendEnvelope provides a mock function with given fields: _a0
func (_m *PeerEnvelopeSender) TrySendEnvelope(_a0 p2p.Envelope) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Envelope) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewPeerEnvelopeSender interface {
	mock.TestingT
	Cleanup(func())
}

// NewPeerEnvelopeSender creates a new instance of PeerEnvelopeSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPeerEnvelopeSender(t mockConstructorTestingTNewPeerEnvelopeSender) *PeerEnvelopeSender {
	mock := &PeerEnvelopeSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
