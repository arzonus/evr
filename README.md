# GOEVR #
[![Build Status](https://travis-ci.org/arzonus/goevr.svg?branch=master)](https://travis-ci.org/arzonus/goevr)
[![Coverage Status](https://coveralls.io/repos/github/arzonus/goevr/badge.svg?branch=master)](https://coveralls.io/github/arzonus/goevr?branch=master)
## Description ##
**goevr** is a [NEVRA library](http://blog.jasonantman.com/2014/07/how-yum-and-rpm-compare-versions/) written in golang for comparing rpm/dpkg package version.
Documentation is [here](https://godoc.org/github.com/arzonus/goevr).
 
 
## Usage ##
You can use this library for comparing different rpm/dpkg package versions

```go
package main

import (
  "github.com/arzonus/goevr"
  "log"
)

func main() {

    oldVersion := "6.0-16+deb8u2"
    newVersion := "0:6.0-17+deb7u2"
    
    
    if goevr.GT(newVersion, oldVersion) {
      log.Println(newVersion, "greater than", oldVersion )
    } else {
      log.Println(newVersion, "less or equal than", oldVersion)
    }
    
    newEvr := goevr.New(newVersion)
    oldEvr := goevr.New(oldVersion)
    
    if newEvr.GTE(oldEvr) {
      log.Println(newEvr.String(), "greater or equal than", oldEvr.String())
    } else {
      log.Println(newEvr.String(), "less than", oldEvr.String())
    }
    
    e, v, r := goevr.Parse(oldVersion)
    log.Println(e, v, r)    
}

```
### Installation ###
```
go get github.com/arzonus/goevr
```