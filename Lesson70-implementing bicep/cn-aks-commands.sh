#!/bin/bash

# References
# ----------------------------

# https://denniszielke.medium.com/securing-ingress-with-azureappgateway-and-egress-traffic-with-azurefirewall-for-azure-kubernetes-41af94051347
# https://techcommunity.microsoft.com/t5/itops-talk-blog/how-to-query-azure-resources-using-the-azure-cli/ba-p/360147

# ----------------------------

source cn-aks.env

# Create resource groups
#-------------------------------------
az group create -n $PROD_VNET_GROUP -l $LOCATION # create the resource group that holds all VNET resources
az group create -n $PROD_HUB_GROUP -l $LOCATION # create the resource group that holds Application Gateway and Firewall resources
az group create -n $PROD_AKS_GROUP -l $LOCATION  # create the resource group that holds AKS resources
az group create -n $SITE_AKS_GROUP -l $LOCATION # create the resource group that holds AKS resources for Site

az group create -n $DEV_AKS_GROUP -l $LOCATION_DEV # create the resource group that holds AKS resources for Development and Testing

# Setup networks and Setup VNET Peering
#-------------------------------------
# we are creatingt the following vnets and peering them
# HUB_VNET which contains Bastion, Jumphost, Firewall and Application Gateway
# PROD_AKS_VNET which contains cn-prod-aks-cn and cn-prod-aks-cow clusters
# DEV_AKS_VNET which contains cn-dev-aks-cn and cn-dev-aks-cow clusters
# SITE_AKS_VNET which contains cn-site-aks-com

az network vnet create -g $PROD_VNET_GROUP -n $PROD_HUB_VNET --address-prefixes $PROD_HUB_VNET_IP_SEG
az network vnet create -g $PROD_VNET_GROUP -n $PROD_AKS_VNET --address-prefixes $PROD_AKS_VNET_IP_SEG
az network vnet create -g $PROD_VNET_GROUP -n $SITE_AKS_VNET --address-prefixes $SITE_AKS_VNET_IP_SEG

az network vnet create -g $PROD_VNET_GROUP -n $DEV_AKS_VNET --address-prefixes $DEV_AKS_VNET_IP_SEG

az network vnet peering create -g $PROD_VNET_GROUP -n hubtoaks1 --vnet-name $PROD_HUB_VNET --remote-vnet $PROD_AKS_VNET --allow-vnet-access
az network vnet peering create -g $PROD_VNET_GROUP -n akstohub1 --vnet-name $PROD_AKS_VNET --remote-vnet $PROD_HUB_VNET --allow-vnet-access

az network vnet peering create -g $PROD_VNET_GROUP -n hubtosite1 --vnet-name $PROD_HUB_VNET --remote-vnet $SITE_AKS_VNET --allow-vnet-access
az network vnet peering create -g $PROD_VNET_GROUP -n sitetohub1 --vnet-name $SITE_AKS_VNET --remote-vnet $PROD_HUB_VNET --allow-vnet-access

az network vnet peering create -g $PROD_VNET_GROUP -n hubtodev1 --vnet-name $PROD_HUB_VNET --remote-vnet $DEV_AKS_VNET --allow-vnet-access
az network vnet peering create -g $PROD_VNET_GROUP -n devtohub1 --vnet-name $DEV_AKS_VNET --remote-vnet $PROD_HUB_VNET --allow-vnet-access

# Creating Subnets
#-------------------------------------
az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $PROD_HUB_VNET -n AzureFirewallSubnet --address-prefix $PROD_HUB_FW_SUBNET_IP_SEG    
az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $PROD_HUB_VNET -n AzureBastionSubnet --address-prefix $PROD_HUB_BASTION_SUBNET_IP_SEG
az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $PROD_HUB_VNET -n $PROD_HUB_APPGW_SUBNET --address-prefix $PROD_HUB_APPGW_SUBNET_IP_SEG
az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $PROD_HUB_VNET -n $PROD_HUB_JUMPHOST_SUBNET --address-prefix $PROD_HUB_JUMPHOST_SUBNET_IP_SEG

az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $PROD_AKS_VNET -n $PROD_AKS_CN_SUBNET --address-prefix $PROD_AKS_CN_SUBNET_IP_SEG 
az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $PROD_AKS_VNET -n $PROD_AKS_COW_SUBNET --address-prefix $PROD_AKS_COW_SUBNET_IP_SEG 

az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $SITE_AKS_VNET -n $SITE_AKS_COM_SUBNET --address-prefix $SITE_AKS_COM_SUBNET_IP_SEG 

az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $DEV_AKS_VNET -n $DEV_AKS_CN_SUBNET --address-prefix $DEV_AKS_CN_SUBNET_IP_SEG 
az network vnet subnet create -g $PROD_VNET_GROUP --vnet-name $DEV_AKS_VNET -n $DEV_AKS_COW_SUBNET --address-prefix $DEV_AKS_COW_SUBNET_IP_SEG 

# Setup Jumphost in the HUB 
#-------------------------------------
az vm list-sizes --location $LOCATION --output table
#Standard_D4a_v4

# az network public-ip create --resource-group $PROD_HUB_GROUP --name $PROD_HUB_JUMPHOST-publicip --sku Standard --location $LOCATION
# PROD_HUB_JUMPHOST_PUBLIC_IP=$(az network public-ip show --resource-group $PROD_HUB_GROUP --name $PROD_HUB_JUMPHOST-publicip --query ipAddress)

az sshkey create \
--name $PROD_HUB_JUMPHOST-sshkey \
--resource-group $PROD_HUB_GROUP \
--location $LOCATION \
 --public-key "@$PROD_HUB_JUMPHOST_SSH_PUB_KEY"

# az network vnet subnet show --resource-group $PROD_VNET_GROUP --name $PROD_HUB_JUMPHOST_SUBNET --vnet-name $PROD_HUB_VNET --query id -o tsv

# az network nic create \
# --name $PROD_HUB_JUMPHOST-nic1 \
# --resource-group $PROD_HUB_GROUP \
# --location $LOCATION \
# --vnet-name $PROD_HUB_VNET \
# --subnet $PROD_HUB_JUMPHOST_SUBNET  

az vm create \
--name $PROD_HUB_JUMPHOST \
--resource-group $PROD_HUB_GROUP \
--location $LOCATION \
--size $PROD_HUB_JUMPHOST_SIZE \
--image UbuntuLTS \
--computer-name $PROD_HUB_JUMPHOST \
--public-ip-address "" \
--vnet-name $PROD_HUB_VNET \
--subnet $PROD_HUB_JUMPHOST_SUBNET \
--authentication-type ssh \
--admin-username $PROD_HUB_JUMPHOST_USERNAME \
--ssh-key-name $PROD_HUB_JUMPHOST-sshkey 

# Setup Bastion in the HUB 
#-------------------------------------
# Bastion requires a HARDCODED SUBNET "AzureBastionSubnet"
az network public-ip create --resource-group $PROD_HUB_GROUP --name $PROD_HUB_BASTION-publicip --sku Standard --location $LOCATION
az network bastion create --name $PROD_HUB_BASTION --public-ip-address $PROD_HUB_BASTION-publicip --resource-group $PROD_HUB_GROUP --vnet-name $PROD_HUB_VNET --location $LOCATION

# Setup Application Gateway in the HUB 
#-------------------------------------
az network public-ip create --resource-group $PROD_HUB_GROUP --name $PROD_HUB_APPGW-publicip --sku Standard
PROD_HUB_APPGW_PUBLIC_IP=$(az network public-ip show --resource-group $PROD_HUB_GROUP --name $PROD_HUB_APPGW-publicip --query ipAddress)

az network application-gateway create \
--name $PROD_HUB_APPGW \
--resource-group $PROD_HUB_GROUP \
--location $LOCATION \
--sku Standard_v2 \
--public-ip-address $PROD_HUB_APPGW-publicip \
--vnet-name $PROD_HUB_VNET_ID \
--subnet $PROD_HUB_APPGW_SUBNET_ID

PROD_HUB_APPGW_ID=$(az network application-gateway show --name $PROD_HUB_APPGW --resource-group $PROD_HUB_GROUP -o tsv --query "id") 

# Private IP not applicable. This is not similar to the firewall Private IP that you can use UDR to route from aks to firewall for egress
# PROD_HUB_APPGW_PRIVATE_IP=$(az network application-gateway show  --name $PROD_HUB_APPGW --resource-group $PROD_HUB_GROUP --query "ipConfigurations[0].privateIpAddress" -o tsv)

# Setup Firewall in the hub and connect it to Log Workspace
#-------------------------------------
az extension add --name azure-firewall
az network public-ip create -g $PROD_VNET_GROUP -n $PROD_HUB_FW-publicip --sku Standard
PROD_HUB_FW_PUBLIC_IP=$(az network public-ip show -g $PROD_VNET_GROUP -n $PROD_HUB_FW-publicip --query ipAddress)
az network firewall create --name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --location $LOCATION
az network firewall ip-config create --firewall-name $PROD_HUB_FW --name $PROD_HUB_FW --public-ip-address $PROD_HUB_FW-publicip --resource-group $PROD_VNET_GROUP --vnet-name $PROD_HUB_VNET
PROD_HUB_FW_PRIVATE_IP=$(az network firewall show -g $PROD_VNET_GROUP -n $PROD_HUB_FW --query "ipConfigurations[0].privateIpAddress" -o tsv)
az monitor log-analytics workspace create --resource-group $PROD_VNET_GROUP --workspace-name $PROD_HUB_FW-logworkspace --location $LOCATION

# Setup Outbound Firewall rules
#-------------------------------------
az network firewall network-rule create --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name "time" --destination-addresses "*"  --destination-ports 123 --name "allow network" --protocols "UDP" --source-addresses "*" --action "Allow" --description "aks node time sync rule" --priority 101
az network firewall network-rule create --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name "dns" --destination-addresses "*"  --destination-ports 53 --name "allow network" --protocols "Any" --source-addresses "*" --action "Allow" --description "aks node dns rule" --priority 102
az network firewall network-rule create --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name "servicetags" --destination-addresses "AzureContainerRegistry" "MicrosoftContainerRegistry" "AzureActiveDirectory" "AzureMonitor" --destination-ports "*" --name "allow service tags" --protocols "Any" --source-addresses "*" --action "Allow" --description "allow service tags" --priority 110
az network firewall network-rule create --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name "hcp" --destination-addresses "AzureCloud.$LOCATION" --destination-ports "1194" --name "allow master tags" --protocols "UDP" --source-addresses "*" --action "Allow" --description "allow aks link access to masters" --priority 120
az network firewall application-rule create --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name 'aksfwar' -n 'fqdn' --source-addresses '*' --protocols 'http=80' 'https=443' --fqdn-tags "AzureKubernetesService" --action allow --priority 101
az network firewall application-rule create  --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name "osupdates" --name "allow network" --protocols http=80 https=443 --source-addresses "*"  --action "Allow" --target-fqdns "download.opensuse.org" "security.ubuntu.com" "packages.microsoft.com" "azure.archive.ubuntu.com" "changelogs.ubuntu.com" "snapcraft.io" "api.snapcraft.io" "motd.ubuntu.com"  --priority 102

# Setup access to docker hub for demo purposes. Remove later
az network firewall application-rule create  --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name "dockerhub" --name "allow network" --protocols http=80 https=443 --source-addresses "*"  --action "Allow" --target-fqdns "*auth.docker.io" "*cloudflare.docker.io" "*cloudflare.docker.com" "*registry-1.docker.io" --priority 200

# Setup access to checkip.dyndns.org. Remove later
az network firewall application-rule create  --firewall-name $PROD_HUB_FW --resource-group $PROD_VNET_GROUP --collection-name "checkip" --name "allow network" --protocols http=80 https=443 --source-addresses "*"  --action "Allow" --target-fqdns "checkip.dyndns.org" --priority 210


# Setup routes to the firewall from each AKS Cluster Subnet
#----------------------------------------------------------------

# Setup routes from CN Cluster to the firewall
az network route-table create -g $PROD_VNET_GROUP --name $PROD_AKS_CN_CLUSTER-routetable
az network route-table route create --resource-group $PROD_VNET_GROUP --name $PROD_HUB_FW --route-table-name $PROD_AKS_CN_CLUSTER-routetable --address-prefix 0.0.0.0/0 --next-hop-type VirtualAppliance --next-hop-ip-address $PROD_HUB_FW_PRIVATE_IP
# Need to update the AKS Cluster subnet
az network vnet subnet update --route-table $PROD_AKS_CN_CLUSTER-routetable --ids $PROD_AKS_CN_SUBNET_ID
az network route-table route list --resource-group $PROD_VNET_GROUP --route-table-name $PROD_AKS_CN_CLUSTER-routetable

# Setup routes from COW Cluster to the firewall
az network route-table create -g $PROD_VNET_GROUP --name $PROD_AKS_COW_CLUSTER-routetable
az network route-table route create --resource-group $PROD_VNET_GROUP --name $PROD_HUB_FW --route-table-name $PROD_AKS_COW_CLUSTER-routetable --address-prefix 0.0.0.0/0 --next-hop-type VirtualAppliance --next-hop-ip-address $PROD_HUB_FW_PRIVATE_IP
# Need to update the AKS Cluster subnet
az network vnet subnet update --route-table $PROD_AKS_COW_CLUSTER-routetable --ids $PROD_AKS_COW_SUBNET_ID
az network route-table route list --resource-group $PROD_VNET_GROUP --route-table-name $PROD_AKS_COW_CLUSTER-routetable


# Create/Set Security Groups
#-------------------------------------
# set up security groups: prod-cn-admin, prod-cow-admin, prod-site-admin for administering ContiNube, ComplianceCow and Site clusters respectively
PROD_AD_CN_ADMIN_ID=$(az ad group list --query "[?contains(displayName, 'prod-cn-admin')].[objectId]" --output tsv)
PROD_AD_COW_ADMIN_ID=$(az ad group list --query "[?contains(displayName, 'prod-cow-admin')].[objectId]" --output tsv) 
SITE_AD_AKS_ADMIN_ID=$(az ad group list --query "[?contains(displayName, 'prod-site-admin')].[objectId]" --output tsv) 
DEV_AD_AKS_ADMIN_ID=$(az ad group list --query "[?contains(displayName, 'dev-site-admin')].[objectId]" --output tsv) 

az identity create --name $PROD_AKS_CN_CLUSTER_ID_NAME --resource-group $PROD_AKS_GROUP
MSI_RESOURCE_ID=$(az identity show --name $PROD_AKS_CN_CLUSTER_ID_NAME --resource-group $PROD_AKS_GROUP -o json | jq -r ".id")
MSI_CLIENT_ID=$(az identity show --name $PROD_AKS_CN_CLUSTER_ID_NAME --resource-group $PROD_AKS_GROUP -o json | jq -r ".clientId")

az role assignment create --role "Virtual Machine Contributor" --assignee $MSI_CLIENT_ID 
#--resource-group $PROD_VNET_GROUP


# Get ACR Info
#-------------------------------------
PROD_ACR_CN_ID=$(az acr show --name $PROD_ACR_CN --subscription $SUBSCRIPTION_ID --query "id" --output tsv)
PROD_ACR_COW_ID=$(az acr show --name $PROD_ACR_COW --subscription $SUBSCRIPTION_ID --query "id" --output tsv)
PROD_ACR_SITE_ID=$(az acr show --name $PROD_ACR_SITE --subscription $SUBSCRIPTION_ID --query "id" --output tsv)

# Setup Private AKS Cluster
#-------------------------------------
# az cli help: https://docs.microsoft.com/en-us/cli/azure/aks?view=azure-cli-latest#az_aks_create

# Register for AAD-V2
az feature register --name AAD-V2 --namespace Microsoft.ContainerService 
az feature list -o table --query "[?contains(name, 'Microsoft.ContainerService/AAD-V2')].{Name:name,State:properties.state}" 
# Once the feature is "registered" use the following command
az provider register --namespace Microsoft.ContainerService 

az aks create \
--name $PROD_AKS_CN_CLUSTER \
--resource-group $PROD_AKS_GROUP \
--location $LOCATION \
--vm-set-type VirtualMachineScaleSets \
--aad-admin-group-object-ids $PROD_AD_CN_ADMIN_ID \
--attach-acr $PROD_ACR_CN_ID \
--enable-aad \
--enable-private-cluster \
--enable-managed-identity \
--assign-identity $MSI_RESOURCE_ID \
--private-dns-zone system \
--node-count 2 \
--network-plugin $PROD_AKS_CNI_PLUGIN \
--vnet-subnet-id $PROD_AKS_CN_SUBNET_ID \
--docker-bridge-address 172.17.0.1/16 \
--dns-service-ip 10.2.0.10 \
--service-cidr 10.2.0.0/24 \
--load-balancer-sku standard \
--outbound-type userDefinedRouting \
--kubernetes-version $PROD_AKS_CLUSTER_VERSION 

# az aks enable-addons --name $PROD_AKS_CN_CLUSTER --resource-group $PROD_AKS_GROUP --addons ingress-appgw --appgw-id $PROD_HUB_APPGW_ID --verbose


# Get the MC_ resource group and the Private DNS Zone
PROD_AKS_NODE_GROUP=$(az aks show --name $PROD_AKS_CN_CLUSTER --resource-group $PROD_AKS_GROUP --query nodeResourceGroup --output tsv)
PROD_AKS_PRIVATE_DNS_ZONE=$(az network private-dns zone list --resource-group $PROD_AKS_NODE_GROUP --query "[0].name" --output tsv)

# List the Private DNS Zones in the MC_ resource group
az network private-dns link vnet list --resource-group $PROD_AKS_NODE_GROUP --zone-name $PROD_AKS_PRIVATE_DNS_ZONE

# Attach the DNS to the JUMPHOST VM Network
az network private-dns link vnet create \
--name $PROD_HUB_JUMPHOST-link \
--resource-group $PROD_AKS_NODE_GROUP \
--virtual-network $PROD_HUB_VNET_ID \
--zone-name $PROD_AKS_PRIVATE_DNS_ZONE \
--registration-enabled false

# Update the route table for each AKS Cluster. Already done. Refer above
# az network vnet subnet update --route-table $PROD_AKS_CN_CLUSTER-routetable --ids $PROD_AKS_CN_SUBNET_ID
# az network vnet subnet update --route-table $PROD_AKS_COW_CLUSTER-routetable --ids $PROD_AKS_COW_SUBNET_ID

# on the jumphost install az cli
# curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash

# on the jumphost install kubectl
# curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
# curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
# echo "$(<kubectl.sha256) kubectl" | sha256sum --check
# sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl


# Access cluster, deploy sample app and test
#-------------------------------------
az aks get-credentials --name $PROD_AKS_CN_CLUSTER --resource-group $PROD_AKS_GROUP



#Provision Applications in Kubernetes, wire them with the firewall for inbound traffic on 443 and expose them to the internet
#-----------------START-------------------------------------------------

# the svc yaml should have a type: LoadBalancer with specific Azure Annotations for internal load balancer
# kubectl apply -f https://raw.githubusercontent.com/Azure/application-gateway-kubernetes-ingress/master/docs/examples/aspnetapp.yaml
# kubectl get ingress
kubectl apply -f aspnetapp.yaml # modified the apsnetapp.yaml to remove ingress and route this through internal load balancer as a service
# you should see an 10.x.x.x ip addresss under external IP address that is fully routable outside the kubernetes domain

# we now need to create a frontend ip and route the l3 traffic to the kubernetes service
# we will need to create public ips for each service: 
# partner.continube.live, test.continube.live, dev.continube.live, 
# partner.compliancecow.live, test.compliancecow.live, dev.compliancecow.live,
# privacybison.com
# compliancecow.com, continube.com

SERVICE_PREFIX="partner.cn" # partner.cow, pb.com, cc.com, cn.com
K8S_SERVICE_NAME="cnreverseproxy"
K8S_SERVICE_NAMESPACE="partner-continube"

az network public-ip create -g $PROD_VNET_GROUP -n $PROD_HUB_FW-$SERVICE_PREFIX-publicip-1 --sku Standard
PROD_HUB_FW_SERVICE_PUBLIC_IP=$(az network public-ip show -g $PROD_VNET_GROUP -n $PROD_HUB_FW-publicip --query ipAddress)

# attach the created public ip to frontend of the firewall
az network firewall ip-config create \
--firewall-name $PROD_HUB_FW \
--name $PROD_HUB_FW \
--public-ip-address $PROD_HUB_FW-$SERVICE_PREFIX-publicip-1 \
--resource-group $PROD_VNET_GROUP \
--vnet-name $PROD_HUB_VNET

# Can priorities number be the same? If not get the last priority for the NAT rule in the Firewall
# TBD
NAT_RULE_PRIORITY="500"

# Get the load balancer ip
# kubectl get services --namespace cn-test --output json | jq -r '.items[0].status.loadBalancer.ingress[0].ip'
TARGET_K8S_ILB_IP=$(kubectl get service $K8S_SERVICE_NAME --namespace $K8S_SERVICE_NAMESPACE --output json | jq -r '.status.loadBalancer.ingress[0].ip')


# create the inbound NAT rule on the firewall to go from the internet into the load balancer service on the cluster/namespace
az network firewall nat-rule create \
--name $SERVICE_PREFIX-inbound-natrule-1 \
--resource-group $PROD_VNET_GROUP \
--description "inbound nat rule for $SERVICE_PREFIX from the internet into the kubernetes load balancer service" \
--priority $NAT_RULE_PRIORITY \
--firewall-name $PROD_HUB_FW \
--protocols "TCP UDP" \
--action Dnat \
--source-addresses "*" \
--dest-addr $PROD_HUB_FW_SERVICE_PUBLIC_IP \
--destination-ports "443"\
--translated-address $TARGET_K8S_ILB_IP \
--translated-port "443"

#-----------------END-------------------------------------------------



# Check egress IP
#-------------------------------------
kubectl run -it --rm aks-ip --image=mcr.microsoft.com/aks/fundamental/base-ubuntu:v0.0.11
apt-get update && apt-get install curl -y
curl -s checkip.dyndns.org
#------------------------

