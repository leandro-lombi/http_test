# http_test

Nowadays, unit testing is essential to ensure software quality and maintainability. It plays a key role in identifying issues early and supporting confident code changes. With that in mind, I'm beginning my journey into unit testing to improve my skills and start applying this practice in my daily work and projects.

```bash
curl http://localhost:8080/hello-world
```

```bash
curl -X POST http://localhost:8080/hello-person -d '{"name": "leandro"}'
```

```bash
curl -F file=@./api/testdata/image.png http://localhost:8080/upload
```

```bash
http# go test ./... -v
```