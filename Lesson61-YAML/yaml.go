package main

// YAMLInput used in main.go
var YAMLInput = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: octopus-deployment
  labels:
    app: web
spec:
  selector:
    matchLabels:
      octopusexport: OctopusExport
  replicas: 1
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: web
        octopusexport: OctopusExport
    spec:
      dnsConfig:
        nameservers:
          - continube
      hostNetwork: false
      containers:
        - name: nginx
          image: nginx
          ports:
            - containerPort: 80
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
            - weight: 100
              podAffinityTerm:
                labelSelector:
                  matchExpressions:
                    - key: app
                      operator: In
                      values:
                        - web
                topologyKey: kubernetes.io/hostname


`

// Data used in main.go
var Data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`
