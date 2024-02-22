package Basic_Functions

type PingResult struct {
	IP          string `json:"ip"`
	RTT         string `json:"rtt"`
	PacketsSent int    `json:"packets_sent"`
	PacketsRecv int    `json:"packets_recv"`
}
