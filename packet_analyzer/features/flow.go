package features

import (
	"packet_analyzer/features/context"
	"strconv"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Flow struct{
	Fin bool
	Dest_ip string
	Src_ip string
	Dest_port int
	Src_port int
	//Src_mac string
	//Dest_mac string

	Packets [] Packet
	
	Start_active int64
	Last_active int64
	Init_window_bytes map[int8]int
	Flow_Iats [] int64
	dead_line int64
	//active []
	//idle []
	//self.forward_bulk_last_timestamp =0
	//self.forward_bulk_start_tmp = 0
	//self.forward_bulk_count = 0
	//self.forward_bulk_count_tmp = 0
	//self.forward_bulk_duration = 0
	//self.forward_bulk_packet_count = 0
	//self.forward_bulk_size = 0
	//self.forward_bulk_size_tmp = 0
	//self.backward_bulk_last_timestamp = 0
	//self.backward_bulk_start_tmp = 0
	//self.backward_bulk_count = 0
	//self.backward_bulk_count_tmp = 0
	//self.backward_bulk_duration = 0
	//self.backward_bulk_packet_count = 0
	//self.backward_bulk_size = 0
	//self.backward_bulk_size_tmp = 0

}
func (flow *Flow)Init(packet gopacket.Packet, direction int8,t int64){
	flow.Fin=false
	flow.Dest_ip=packet.NetworkLayer().NetworkFlow().Dst().String()
	flow.Src_ip=packet.NetworkLayer().NetworkFlow().Src().String()
	tpl:=packet.TransportLayer()
	if tpl!=nil{
		flow.Dest_port,_=strconv.Atoi(packet.TransportLayer().TransportFlow().Dst().String())
		flow.Src_port,_=strconv.Atoi(packet.TransportLayer().TransportFlow().Src().String())

	}
	flow.Start_active=t
	flow.Last_active=t
	flow.Init_window_bytes=make(map[int8]int)
	flow.Init_window_bytes[context.FORWARD]=0
	flow.Init_window_bytes[context.REVERSE]=0
}
func (flow *Flow)Add_packet(packet Packet, direction int8){
	flow.Packets=append(flow.Packets, packet)
	tcpLayer:=packet.Packet.Layer(layers.LayerTypeTCP)
	if tcpLayer!=nil{
		if flow.Init_window_bytes[direction]==0{
			flow.Init_window_bytes[direction]=int(tcpLayer.(*layers.TCP).Window)
		}
		if tcpLayer.(*layers.TCP).FIN || tcpLayer.(*layers.TCP).RST{
			flow.Fin=true
		}
	}

	flow.Last_active=packet.Time
	if len(flow.Packets)>=2{
		flow.Flow_Iats = append(flow.Flow_Iats, packet.Time-flow.Start_active)
	}

}
func (flow Flow)Get_init_win_bytes(direction int8)(int64){
	return int64(flow.Init_window_bytes[direction])
}
func (flow Flow)Get_flow_duration()(int64){
	return flow.Last_active-flow.Start_active
}
func (flow Flow)Get_packet_lengths(direction int8)([]int){
	var lengths[] int
	for _, p :=range flow.Packets{
		if p.Direction==direction{
			lengths = append(lengths, p.Get_packet_length())
		}
	}
	return lengths
}
func (flow Flow)Get_stat_length(direction int8)(sum int64,mean float64){
	lengths:=flow.Get_packet_lengths(direction)
	sum=0
	for _,l := range lengths{
		sum+=int64(l)
	}
	mean= float64(sum)
	mean=mean / float64(len(lengths))
	return sum,mean
}
func (flow Flow)Get_ext_length(direction int8)(max int64,min int64){
	lengths:=flow.Get_packet_lengths(direction)
	for i,l := range lengths{
		l_int64:=int64(l)
		if i==0{
			max=l_int64
			min=l_int64
		}else{
			if l_int64> max{
				max=l_int64
			}
			if l_int64< min{
				min =l_int64
			}
		}
	}
	return max,min
}
func (flow Flow)Get_bytes_per_second(direction int8)(float64){
	period:=flow.Get_flow_duration()
	if direction == context.FORWARD{
		sum1,_:=flow.Get_stat_length(context.FORWARD)
		return float64(sum1)/float64(period)
	} else if direction == context.REVERSE{
		sum2,_:=flow.Get_stat_length(context.REVERSE)
		return float64(sum2)/float64(period)
		} else{
			sum1,_:=flow.Get_stat_length(context.FORWARD)
		sum2,_:=flow.Get_stat_length(context.REVERSE)
		return (float64(sum1+sum2)/float64(period))*1e6
	}
}
func (flow Flow)Get_iats(direction int8)([] int64){
	var packets [] Packet
	var iats [] int64 
	if direction !=context.NONE{
		for _,p :=range flow.Packets{
			if p.Direction==direction{
				packets = append(packets, p)
			}
		}
		for i:=1;i<len(packets);i++{
			iats = append(iats, packets[i].Time-packets[i-1].Time)
		}
	}else{
		for i:=1;i<len(flow.Packets);i++{
			iats = append(iats, flow.Packets[i].Time-flow.Packets[i-1].Time)
		}

	}

	return iats
}
func (flow Flow)Get_dest_port()(int64){
	return int64(flow.Dest_port)
}
func (flow Flow)Get_stat_iat(direction int8)(mean_iat float64){
	var sum int64=0
	iats:=flow.Get_iats(direction)
	for _,iat :=range iats{
		sum+=iat
	}
	return float64(sum)/float64(len(iats))
}
func (flow Flow)Get_packets_per_second(direction int8)(float64){
	return float64(len(flow.Packets))/float64(flow.Get_flow_duration())
}
func (flow Flow)Get_ext_iat(direction int8)(max,min int64){
	iats:=flow.Get_iats(direction)
	for i,iat :=range iats{
		if i==0{
			max=iat
			min=iat
		}else{
			if iat>max{
				max=iat
			}
			if iat<min{
				min=iat
			}
		}
	}
	return max,min
}
func (flow* Flow)Set_dead_line(dead_line int64){
	flow.dead_line=dead_line
}
func (flow Flow)Get_header_length(direction int8)(length int64){
	length=0
	if direction !=context.NONE{

		for _,p := range flow.Packets{
			if p.Direction==direction{
				length+=p.Get_header_length()
			}
		}
	}else{
		for _,p := range flow.Packets{
			length+=p.Get_header_length()
		}
	}
	return length
}
func (flow Flow)Get_flow_iats(direction int8)(length [] int64){
	return flow.Flow_Iats
}
func (flow Flow)Get_ext_flow_iat(direction int8)(max, min int64){
	iats:=flow.Get_flow_iats(direction)
	if iats!=nil{
		for i,iat :=range iats{
			if i==0{
				max=iat
				min=iat
			}else{
				if iat>max{
					max=iat
				}
				if iat< min{
					min=iat
				}
			}
		}
	}
	return
}
func(flow Flow)Get_stat_flow_iat(direction int8)(mean float64){
	iats:=flow.Get_flow_iats(direction)
	mean=0
	if iats!=nil{
		for _,iat :=range iats{
			mean+=float64(iat)
		}
	}
	mean=mean/float64(len(iats))
	return
}
func (flow Flow)Get_fin()(bool){
	return flow.Fin
}