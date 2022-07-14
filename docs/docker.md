```shell
docker build . -f cmd/meme/Dockerfile -t memepen/meme
```

```shell
docker run --rm -ti -p 8080:8080 memepen/meme
```