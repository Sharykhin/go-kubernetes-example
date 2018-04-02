fix credentials
```bash
kubectl create clusterrolebinding --user system:serviceaccount:kube-system:default kube-system-cluster-admin --clusterrole cluster-admin
```

```bash
MYSQL_ADDRESS="root:root@tcp(mysql-service-2:3306)/test?charset=utf8"
```
asd
