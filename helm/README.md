    
    // packing
    helm package ./dbservice-factory -d ./package
    
    // install from local
    helm install . --name dbservice-factory --namespace cloud-svc
    
    // install from github
    version=dbservice-factory-0.1.0.tgz && \
    helm install https://github.com/samqintw/dbservice-factory/raw/master/helm/package/$version \
    --name dbservice-factory \
    --namespace cloud-svc