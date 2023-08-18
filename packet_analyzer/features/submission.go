package features

import (
	"packet_analyzer/features/context"
)
func Submit_all_features(flow Flow, p bool)([17]float64){
	dest_port:=flow.Dest_port
	flow_duration:=flow.Get_flow_duration()
	total_length_fwd,fwd_packets_length_mean:=flow.Get_stat_length(context.FORWARD)
	total_length_bwd,_:=flow.Get_stat_length(context.REVERSE)
	fwd_packets_length_max,_:=flow.Get_ext_length(context.FORWARD)
	flow_bytes_per_s:=flow.Get_bytes_per_second(context.NONE)
	flow_packets_per_s:=flow.Get_packets_per_second(context.NONE)
	fwd_iat_mean:=flow.Get_stat_iat(context.FORWARD)
	_,fwd_iat_min:=flow.Get_ext_iat(context.FORWARD)
	_,bwd_iat_min:=flow.Get_ext_iat(context.REVERSE)
	fwd_header_length:=flow.Get_header_length(context.FORWARD)
	bwd_packets_s:=flow.Get_packets_per_second(context.REVERSE)
	fwd_init_win_bytes:=flow.Get_init_win_bytes(context.FORWARD)
	bwd_init_win_bytes:=flow.Get_init_win_bytes(context.REVERSE)
	flow_iat_mean:=flow.Get_stat_flow_iat(context.NONE)
	_,flow_iat_min:=flow.Get_ext_flow_iat(context.NONE)
	var submission [17] float64=[17]float64{float64(dest_port),float64(flow_duration),float64(total_length_fwd),float64(total_length_bwd),float64(fwd_packets_length_max),fwd_packets_length_mean,flow_bytes_per_s,flow_packets_per_s,flow_iat_mean,float64(flow_iat_min),fwd_iat_mean,float64(fwd_iat_min),float64(bwd_iat_min),float64(fwd_header_length),bwd_packets_s,float64(fwd_init_win_bytes),float64(bwd_init_win_bytes)}
	if p{
		println("Destination_Port",dest_port)
		println("Flow_Duration",flow_duration)
		println("Total_length_of_fwd_packets",total_length_fwd)
		println("Total_length_of_bwd_packets",total_length_bwd,)
		println("Fwd Packet Length Max",fwd_packets_length_max)
		println("Fwd Packet Length Mean",fwd_packets_length_mean)
		println("Flow Bytes/s",flow_bytes_per_s)
		println("Flow Packets/s",flow_packets_per_s)
		println("fwd_iat_mean",fwd_iat_mean)
		println("fwd_iat_min",fwd_iat_min)
		println("bwd_iat_min",bwd_iat_min)
		println("fwd_header_length",fwd_header_length)
		println("bwd_packets_s",bwd_packets_s)
		println("fwd_init_win_bytes",fwd_init_win_bytes)
		println("bwd_init_win_bytes",bwd_init_win_bytes)
		println("flow_iat_mean",flow_iat_mean)
		println("flow_iat_min",flow_iat_min)
	}
	return submission
}