package winrm

import (
	"github.com/Azure/go-ntlmssp"
	"github.com/optanix/winrm/soap"
	"net"
)

// ClientNTLM provides a transport via NTLMv2
type ClientNTLM struct {
	clientRequest
}

// Transport creates the wrapped NTLM transport
func (c *ClientNTLM) Transport(endpoint *Endpoint) error {
	c.clientRequest.Transport(endpoint)
	c.clientRequest.transport = &ntlmssp.Negotiator{RoundTripper: c.clientRequest.transport}
	return nil
}

// Post make post to the winrm soap service (forwarded to clientRequest implementation)
func (c ClientNTLM) Post(client *Client, request *soap.SoapMessage) (string, error) {
	return c.clientRequest.Post(client, request)
}


func NewClientNTLMWithDial(dial func(network, addr string) (net.Conn, error)) *ClientNTLM {
	return &ClientNTLM{
		clientRequest{
			dial:dial,
		},
	}
}