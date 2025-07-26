package traceroute

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type M_ICMP struct {
	TTL      int     `json:"ttl"`
	Peer     string  `json:"peer"`
	Duration float64 `json:"duration"`
}

func Traceroute(dest string) ([]M_ICMP, error) {
	maxHops := 30
	timeout := time.Second * 2
	var result []M_ICMP

	destAddr, err := net.ResolveIPAddr("ip4", dest)
	if err != nil {
		return result, fmt.Errorf("[ERROR] Resolve error: %v", err)
	}

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return result, fmt.Errorf("[ERROR] Listen error: %v", err)
	}
	defer conn.Close()

	for ttl := 1; ttl <= maxHops; ttl++ {
		conn.IPv4PacketConn().SetTTL(ttl)

		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  ttl,
				Data: []byte("HELLO-RTR"),
			},
		}
		b, _ := msg.Marshal(nil)
		start := time.Now()
		_, err := conn.WriteTo(b, &net.IPAddr{IP: destAddr.IP})
		if err != nil {
			return result, fmt.Errorf("[ERROR] WriteTo error: %v", err)
		}

		conn.SetReadDeadline(time.Now().Add(timeout))
		reply := make([]byte, 1500)
		n, peer, err := conn.ReadFrom(reply)
		duration := time.Since(start)

		peerStr := ""
		if peer != nil {
			peerStr = peer.String()
		}

		if err != nil {
			result = append(result, M_ICMP{TTL: ttl, Peer: "", Duration: -1})
			continue
		}

		rm, err := icmp.ParseMessage(1, reply[:n])
		if err != nil {
			return result, fmt.Errorf("[ERROR] ParseMessage error: %v", err)
		}

		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			result = append(result, M_ICMP{TTL: ttl, Peer: peerStr, Duration: float64(duration.Microseconds()) / 1000})
		case ipv4.ICMPTypeEchoReply:
			result = append(result, M_ICMP{TTL: ttl, Peer: peerStr, Duration: float64(duration.Microseconds()) / 1000})
			return result, nil
		default:
			result = append(result, M_ICMP{TTL: ttl, Peer: "", Duration: -1})
		}
	}
	return result, nil
}
