# Iris Rest Api (Sample)
[![Build Status](https://travis-ci.org/giansalex/iris-rest.svg?branch=master)](https://travis-ci.org/giansalex/iris-rest)     
using [Iris](https://github.com/kataras/iris) with Docker

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