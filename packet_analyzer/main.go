package main

import (
	"fmt"
	"os"
	"packet_analyzer/bwlist"
	"packet_analyzer/features"
	"packet_analyzer/features/context"
	"packet_analyzer/submitter"
	"time"

	"github.com/google/gopacket"
	"github.com/joho/godotenv"
	"github.com/subgraph/go-nfnetlink/nfqueue"
)

func main() {
	godotenv.Load()
	q := nfqueue.NewNFQueue(1)
	var next_time=time.Now().UnixMilli()
	ps, err := q.Open()
	if err != nil {
		fmt.Printf("Error opening NFQueue: %v\n", err)
		os.Exit(1)
	}
	defer q.Close()
	var whitelist bwlist.BWlist
	var blacklist bwlist.BWlist
	whitelist.Init()
	blacklist.Init()


	flows:=make(map[gopacket.Flow]*features.Flow)
	var u=-10

	for p := range ps {
		print("in",time.Now().UnixMicro(),"\n")
		networkLayer := p.Packet.NetworkLayer()
		key:=networkLayer.NetworkFlow()
		src,dst:=key.Endpoints()
		if whitelist.Check(src.String()){
			p.Accept()
			continue
		}
		if blacklist.Check(src.String()){
			p.Drop()
			continue
		}
		reverse, err:=gopacket.FlowFromEndpoints(dst, src)
		if err!=nil{
			print("No")
		}
		var t int64=time.Now().UnixMilli()
		_, ok1 :=flows[key]
		_, ok2 :=flows[reverse]
		var direction int8
		var packet_case byte
		if ok1{
			packet_case='a'
			direction=context.FORWARD
			flows[key].Add_packet(features.Packet{Packet: p.Packet,Time: t,Direction: direction},direction)
			//println((*flows[key]).Src_ip," ",(*flows[key]).Dest_ip)
			//println(len(flows[key].Packets))
		} else if ok2{
			packet_case='b'
			direction=context.REVERSE
			flows[reverse].Add_packet(features.Packet{Packet: p.Packet,Time: t, Direction: direction},direction)
		} else {
			packet_case='c'
			direction=context.FORWARD
			flows[key]=&features.Flow{}
			flows[key].Init(p.Packet,direction,t)
			flows[key].Add_packet(features.Packet{Packet: p.Packet,Time: t},direction)
		}
		p.Accept()
		u+=1
		print("out",time.Now().UnixMicro(),"\n")
		if u==10{
			break
		}
		var current_time=time.Now().UnixMilli()
		var fin bool=false
		var tar *features.Flow
		if packet_case=='a' ||packet_case== 'c'{
			fin=(*flows[key]).Get_fin()
			tar=(flows[key])
		}else{
			fin=(*flows[reverse]).Get_fin()
			tar=(flows[reverse])
		}
		if current_time>next_time || fin{
			next_time=current_time+features.UPDATE_INTERVAL
			data:=features.Submit_all_features(*tar,false)
			result:=submitter.Submit_to_detector(data)
			print(result==0)
			if result!=0{
				blacklist.Append(src.String())
			}
			if fin{
				delete(flows,reverse)
			}
		}
	}
	//fo, err := os.Create("output.json")

	//if err != nil {
    //    panic(err)
    //}
	//defer func(){
	//	if err =fo.Close(); err!= nil{
	//		panic(err)
	//	}
	//}()
}