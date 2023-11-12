## Research Record
1. Find a intrution detection model
2. try to prune the features to make it feasible for the traffic
3. run it with iptable to block the ip with anomaly traffic
#### here's my flow
![image](https://hackmd.io/_uploads/HycgflCQT.png)
There's a detector running in one k8s node. The main server would capture the features from traffic and send them to the detector through http, then the detector would send a number back. 0 is for benign, while others would correspond to a type of intrusion.
### Find a intrution detection model
Referenced to  this github repo: https://github.com/alik604/cyber-security, i chose to use CIC-IDS dataset.

### Feature Pruning 
- I prune the features by PCA tool provided by sklearn at first, then do further remove by the correlation between features.
- This part is included in edge_nids/edge_anomaly_detection/CIC_IDS2017.ipynb, but the removal is kind of trial and error, thus the code might not be correct.

### run with iptable
- This section has 2 steps
1. understand how to capture the traffics
    - Understand how iptable work
    - here's some references
        - https://github.com/Asphaltt/go-nfnetlink-example
        - https://medium.com/skilluped/what-is-iptables-and-how-to-use-it-781818422e52 
        - https://home.regit.org/netfilter-en/using-nfqueue-and-libnetfilter_queue/

2. Extract the features data for the model to do detection




## Some Declaration
1. This repo not runnable due to the need of some setup in iptable and deployment in k8s.
2. The correctness of detection model is not justified, because i haven't run it on any real testing traffic.



