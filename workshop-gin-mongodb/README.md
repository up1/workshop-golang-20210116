# Workshop

### Start MongoDB
```
$docker-compose up -d
$docker-compose ps
```

### Run Server
```
$sh run.sh
```


### Generate API document
```
$go get -u github.com/swaggo/swag/cmd/swag
$swag init --dir cmd/ --parseDependency --output docs
```

Access to document at http://localhost:8080/docs/index.html