    helm create dbservice-factory
    
    // packing
    helm package ./dbservice-factory -d ./package
    
    // install from local
    helm install . --name dbservice-factory --namespace cloud-svc
    
    // install from github
    helm install --repo https://github.com/samqintw/dbservice-factory/tree/master/helm dbservice-factory --name dbservice-factory --namespace cloud-svc