// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	context "context"

	info "github.com/Axia-Tech/axiav2/api/info"
	ids "github.com/Axia-Tech/axiav2/ids"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/Axia-Tech/axiav2/utils/rpc"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetBlockchainID provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) GetBlockchainID(_a0 context.Context, _a1 string, _a2 ...rpc.Option) (ids.ID, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, string, ...rpc.Option) ids.ID); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkID provides a mock function with given fields: _a0, _a1
func (_m *Client) GetNetworkID(_a0 context.Context, _a1 ...rpc.Option) (uint32, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) uint32); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkName provides a mock function with given fields: _a0, _a1
func (_m *Client) GetNetworkName(_a0 context.Context, _a1 ...rpc.Option) (string, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeID provides a mock function with given fields: _a0, _a1
func (_m *Client) GetNodeID(_a0 context.Context, _a1 ...rpc.Option) (string, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeIP provides a mock function with given fields: _a0, _a1
func (_m *Client) GetNodeIP(_a0 context.Context, _a1 ...rpc.Option) (string, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeVersion provides a mock function with given fields: _a0, _a1
func (_m *Client) GetNodeVersion(_a0 context.Context, _a1 ...rpc.Option) (*info.GetNodeVersionReply, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *info.GetNodeVersionReply
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) *info.GetNodeVersionReply); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.GetNodeVersionReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxFee provides a mock function with given fields: _a0, _a1
func (_m *Client) GetTxFee(_a0 context.Context, _a1 ...rpc.Option) (*info.GetTxFeeResponse, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *info.GetTxFeeResponse
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) *info.GetTxFeeResponse); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.GetTxFeeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVMs provides a mock function with given fields: _a0, _a1
func (_m *Client) GetVMs(_a0 context.Context, _a1 ...rpc.Option) (map[ids.ID][]string, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[ids.ID][]string
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) map[ids.ID][]string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ids.ID][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBootstrapped provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) IsBootstrapped(_a0 context.Context, _a1 string, _a2 ...rpc.Option) (bool, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, ...rpc.Option) bool); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peers provides a mock function with given fields: _a0, _a1
func (_m *Client) Peers(_a0 context.Context, _a1 ...rpc.Option) ([]info.Peer, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []info.Peer
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) []info.Peer); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]info.Peer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Uptime provides a mock function with given fields: _a0, _a1
func (_m *Client) Uptime(_a0 context.Context, _a1 ...rpc.Option) (*info.UptimeResponse, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *info.UptimeResponse
	if rf, ok := ret.Get(0).(func(context.Context, ...rpc.Option) *info.UptimeResponse); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.UptimeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...rpc.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
