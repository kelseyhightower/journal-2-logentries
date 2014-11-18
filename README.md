# journal-2-logentries 


## Usage

```
sudo LOGENTRIES_TOKEN=<token> journal-2-logentries
```

### Docker

```
GO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
```

```
docker build -t quay.io/kelseyhightower/journal-2-logentries .
```

```
docker run -d -e 'LOGENTRIES_TOKEN=<token>' quay.io/kelseyhightower/journal-2-logentries
```
