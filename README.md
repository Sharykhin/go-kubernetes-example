fix credentials
```bash
kubectl create clusterrolebinding --user system:serviceaccount:kube-system:default kube-system-cluster-admin --clusterrole cluster-admin
```

```bash
MYSQL_ADDRESS="root:root@tcp(mysql-service-2:3306)/test?charset=utf8"
```

Example of using sed
```bash
sed 's/${VERSION}/0.0.1/g' deployment.tpl.yml > deployment.yml
```