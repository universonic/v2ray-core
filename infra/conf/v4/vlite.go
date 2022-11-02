package v4

import (
	"github.com/golang/protobuf/proto"

	"github.com/v2fly/v2ray-core/v5/infra/conf/cfgcommon"
	"github.com/v2fly/v2ray-core/v5/proxy/vlite/inbound"
	"github.com/v2fly/v2ray-core/v5/proxy/vlite/outbound"
)

// VliteuClientConfig is configuration of trojan servers
type VliteuClientConfig struct {
	Address                     *cfgcommon.Address `json:"address"`
	Port                        uint16             `json:"port"`
	Password                    string             `json:"password"`
	ScramblePacket              bool               `json:"scramblePacket"`
	EnableFec                   bool               `json:"enableFec"`
	EnableStabilization         bool               `json:"enableStabilization"`
	EnableRenegotiation         bool               `json:"enableRenegotiation"`
	HandshakeMaskingPaddingSize uint32             `json:"handshakeMaskingPaddingSize"`
}

// Build implements Buildable
func (c *VliteuClientConfig) Build() (proto.Message, error) {
	config := new(outbound.UDPProtocolConfig)

	if c.Address == nil {
		return nil, newError("Vliteu server address is not set.")
	}
	if c.Port == 0 {
		return nil, newError("Invalid Vliteu port.")
	}
	if c.Password == "" {
		return nil, newError("Vliteu password is not specified.")
	}
	config.ScramblePacket = c.ScramblePacket
	config.EnableFec = c.EnableFec
	config.EnableStabilization = c.EnableStabilization
	config.EnableRenegotiation = c.EnableRenegotiation
	config.HandshakeMaskingPaddingSize = c.HandshakeMaskingPaddingSize

	return config, nil
}

// VliteuServerConfig is Inbound configuration
type VliteuServerConfig struct {
	Password                    string `json:"password"`
	ScramblePacket              bool   `json:"scramblePacket"`
	EnableFec                   bool   `json:"enableFec"`
	EnableStabilization         bool   `json:"enableStabilization"`
	EnableRenegotiation         bool   `json:"enableRenegotiation"`
	HandshakeMaskingPaddingSize uint32 `json:"handshakeMaskingPaddingSize"`
}

// Build implements Buildable
func (c *VliteuServerConfig) Build() (proto.Message, error) {
	config := new(inbound.UDPProtocolConfig)

	if c.Password == "" {
		return nil, newError("Vliteu password is not specified.")
	}
	config.ScramblePacket = c.ScramblePacket
	config.EnableFec = c.EnableFec
	config.EnableStabilization = c.EnableStabilization
	config.EnableRenegotiation = c.EnableRenegotiation
	config.HandshakeMaskingPaddingSize = c.HandshakeMaskingPaddingSize

	return config, nil
}
