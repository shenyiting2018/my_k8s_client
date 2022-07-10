# my_k8s_client
A simple client written in Go to interact with k8s

Features:

* connect to the k8s cluster
* print out the namespaces on the cluster
* create a new namespace
* create a pod in that namespace that runs a simple hello-world container
* print out pod names and the namespace they are in for any pods that have a label of ‘k8s-app=kube-dns’ or a similar label is ok as well
* delete the hello-world pod created from above
* extra credit - show how an client-go informer works


