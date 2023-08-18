package bwlist
type BWlist struct{
	Suspects map[string]bool
	Suspects_with_count map[string]
}
func (bwlist *BWlist)Init()(){
	bwlist.Suspects=make(map[string]bool)
}
func (bwlist BWlist)Check(ip string)(bool){
	_,exist:=bwlist.Suspects[ip]
	if exist{
		return true
	}
	return false
}
func (bwlist *BWlist)Append(ip string)(bool){
	bwlist.Suspects[ip]=true
	return true
}
func (bwlist *BWlist)Remove(ip string)(bool){
	_,exist:=bwlist.Suspects[ip]
	if exist{
		delete(bwlist.Suspects,ip)
		return true
	}else{
		return false
	}
}