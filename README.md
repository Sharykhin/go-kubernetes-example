```bash
docker build -t go-hello-world . 
```

```bash
docker run -p 3002:3002 --name ttt --rm go-hello-world
```