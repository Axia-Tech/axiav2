// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsubnetlookup

import (
	"context"

	"github.com/Axia-Tech/axiav2/ids"
	"github.com/Axia-Tech/axiav2/snow"

	subnetlookuppb "github.com/Axia-Tech/axiav2/proto/pb/subnetlookup"
)

var _ subnetlookuppb.SubnetLookupServer = &Server{}

// Server is a subnet lookup that is managed over RPC.
type Server struct {
	subnetlookuppb.UnsafeSubnetLookupServer
	aliaser snow.SubnetLookup
}

// NewServer returns a subnet lookup connected to a remote subnet lookup
func NewServer(aliaser snow.SubnetLookup) *Server {
	return &Server{aliaser: aliaser}
}

func (s *Server) SubnetID(
	_ context.Context,
	req *subnetlookuppb.SubnetIDRequest,
) (*subnetlookuppb.SubnetIDResponse, error) {
	chainID, err := ids.ToID(req.ChainId)
	if err != nil {
		return nil, err
	}
	id, err := s.aliaser.SubnetID(chainID)
	if err != nil {
		return nil, err
	}
	return &subnetlookuppb.SubnetIDResponse{
		Id: id[:],
	}, nil
}
