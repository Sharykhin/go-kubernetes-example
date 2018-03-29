```bash
docker build -t go-hello-world . 
```

```bash
docker run -p 3002:3002 --name ttt --rm go-hello-world
```

TEST CASE:
```bash
kubectl create -f k8s/mysql1-1.yml
```

```bash
kubectl create -f k8s/mysql1-2.yml
```

Go to the dashboard and through the terminal use the following
commands for each mysql pod:
```bash
mysql -uroot -proot
create database test;
use test;
create table users(id int auto_increment primary key, name varchar(255));
insert into users(name) values("foo")
```

for a second pod use another name

```bash
kubectl create -f k8s/go1-1.yml
```

```bash
minikube service entrypoint
```