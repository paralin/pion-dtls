// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !tinygo
// +build !tinygo

package dtls

import (
	"net"

	"github.com/pion/dtls/v2/internal/net/udp"
	"github.com/pion/dtls/v2/pkg/protocol"
	"github.com/pion/dtls/v2/pkg/protocol/recordlayer"
)

// Listen creates a DTLS listener
func Listen(network string, laddr *net.UDPAddr, config *Config) (net.Listener, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	lc := udp.ListenConfig{
		AcceptFilter: func(packet []byte) bool {
			pkts, err := recordlayer.UnpackDatagram(packet)
			if err != nil || len(pkts) < 1 {
				return false
			}
			h := &recordlayer.Header{}
			if err := h.Unmarshal(pkts[0]); err != nil {
				return false
			}
			return h.ContentType == protocol.ContentTypeHandshake
		},
	}
	// If connection ID support is enabled, then they must be supported in
	// routing.
	if config.ConnectionIDGenerator != nil {
		lc.DatagramRouter = cidDatagramRouter(len(config.ConnectionIDGenerator()))
		lc.ConnectionIdentifier = cidConnIdentifier()
	}
	parent, err := lc.Listen(network, laddr)
	if err != nil {
		return nil, err
	}
	return &listener{
		config: config,
		parent: parent,
	}, nil
}

// listenUDP implements net.ListenUDP
func listenUDP(network string, laddr *net.UDPAddr) (*net.UDPConn, error) {
	return net.ListenUDP(network, laddr)
}
