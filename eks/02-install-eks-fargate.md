# Install EKS

Please follow the prerequisites doc before this.

## Install a EKS cluster with EKSCTL

```
eksctl create cluster --name demo-cluster --region us-east-1 
```

## Delete the cluster

```
eksctl delete cluster --name demo-cluster --region us-east-1
```

## Adding More Nodes to the Cluster [Optinal]

```
eksctl scale nodegroup \
  --cluster=<your-cluster-name> \
  --name=<node-group-name> \
  --nodes=4 \
  --nodes-min=2 \
  --nodes-max=5
```
## Create EKS Cluster 
### If you face error above command

```
eksctl create cluster --name demo-ekscluster --region ap-south-1 --version 1.21 --nodegroup-name linux-nodes --node-type t2.micro --nodes 2
```


