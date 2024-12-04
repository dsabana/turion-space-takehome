package telemetryIngest

import (
	"fmt"
	"log"
	"net"
)

// Client struct containing the udp connection and the repository
type Client struct {
	conn *net.UDPConn
	repo *Repository
}

// NewTelemetryIngestClient initializes a new UDP client
func NewTelemetryIngestClient(cfg ClientConfiguration) (*Client, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%s", cfg.ClientPort))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve UDP address: %v", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen UDP address: %v", err)
	}

	repository, err := NewRepository(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create repository: %v", err)
	}

	return &Client{
		conn: conn,
		repo: repository,
	}, nil
}

// ListenAndServe receives and parses telemetry packets
func (tic *Client) ListenAndServe() {
	defer tic.conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := tic.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("error reading from UDP: %v", err)
			continue
		}

		log.Printf("[DEBUG] read %d bytes from %s", n, addr.String())

		packet, err := ParseTelemetryPacket(buffer[:n])
		if err != nil {
			log.Printf("error parsing packet: %v", err)
			continue
		}

		ValidatePacket(packet)

		tic.repo.SavePacketEntry(packet)
	}
}
