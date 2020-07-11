# Iris Rest Api (Sample)
[![Build Status](https://travis-ci.org/giansalex/iris-rest.svg?branch=master)](https://travis-ci.org/giansalex/iris-rest)     
Using [Iris](https://github.com/kataras/iris) deploy on Heroku.

Run application. 
```bash
go run ./
```

Run tests.
```bash
go test ./
```

## Docker

Deploy iris application on Docker.

**Build**

```bash
docker build -t goiris .
```

**Run**
```bash
docker run -d -p 80:8080 --name restapp goiris
```

### Deploy on AWS Elastic Beanstalk

In type application choose docker.
