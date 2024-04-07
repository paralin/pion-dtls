// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo
// +build tinygo

package dtls

import (
	"errors"
	"net"
)

// Listen creates a DTLS listener
func Listen(network string, laddr *net.UDPAddr, config *Config) (net.Listener, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return nil, errors.New("tinygo: dtls listener not implemented")
}

// listenUDP implements net.ListenUDP
func listenUDP(network string, laddr *net.UDPAddr) (*net.UDPConn, error) {
	return nil, errors.New("tinygo: udp listener not implemented")
}
