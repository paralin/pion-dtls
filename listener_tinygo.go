// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo

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

	// net.ListenUDP is not defined on tinygo
	return nil, errors.New("tinygo: dtls listener not implemented")
}
