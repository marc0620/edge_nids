package features

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Packet struct{
	Packet gopacket.Packet 
	Time int64
	Direction int8
}

func (p Packet)Get_packet_length()(int){
	return len(p.Packet.Data())
}

func (p Packet)Get_header_length()(int64){
	ipLayer := p.Packet.Layer(layers.LayerTypeIPv4)
	if ipLayer!=nil{
		ip,_:=ipLayer.(*layers.IPv4)
		return int64(ip.IHL*4)
	}else{
		return 8
	}
}
func (p Packet)Get_window_bytes(direction int8)(int){
	tcpLayer:=p.Packet.Layer(layers.LayerTypeTCP)
	if tcpLayer!=nil{
		tcp,_:=tcpLayer.(*layers.TCP)
		return int(tcp.Window)
	}else{
		return -1
	}
}