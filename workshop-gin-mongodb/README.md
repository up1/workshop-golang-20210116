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

### Keep metric with [Prometheus](https://github.com/penglongli/gin-metrics)

Access to http://localhost:8080/metrics

### Generate API document
```
$go get -u github.com/swaggo/swag/cmd/swag
$swag init --dir cmd/ --parseDependency --output docs
```

Access to document at http://localhost:8080/docs/index.html