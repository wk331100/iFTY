# iFTY Is short for infinity
<p>
<a href="https://travis-ci.org/wk331100/coder_php_framework"><img src="https://travis-ci.org/wk331100/coder_php_framework.svg" alt="Build Status"></a>
</p>

iFTY Is A HTTP Web Framework Written In Go (Golang). It Based On `FastHTTP` 10x Faster Than `net/http` 

## Introduction
Most Golang Fans Written PHP Many Years, When First Change Their Coding Behavior And Habit, Pain Can Only Be Feeled. Completely Ignoring The Advantages Of Golang Itself.
So, We Need A Framework That Keep the features of PHP language Framework.
Now, `iFTY` Comes!


## Features

- Based On `FastHTTP`, 10X Faster Than `net/http`
- Use `Go Mod` Manage Packages
- Support `Router` And `MiddleWare`
- Support ENV file
- Support`CSM` Structure, Controller(C), Service(S), Model(M)
- Strong early warning mechanism (DB, Cache, Script, Error)
- Write Golang Code Just Like PHP

## Installation

Install From Github
```
go get github.com/wk331100/iFTY
```

Install From Official Website

[Http://www.ifty.tech/download](Http://www.ifty.tech/download)

## Quickstart

#### Start Server

Config File `config/server.go` And Run `go run main.go` 

File `server.go`
```
var ServerConfig = map[string]interface{}{
	"Port" : 8080,
}
```
Run

```
go run main.go
```

#### Directory Structure
```
- app
    - controllers   // Controllers For Business logic
    - middleware    // MiddleWare Before Controller And Response
    - libs          // Customer Functions
    - models        // Data Model
    - services      // Service
- bootstrap
    - application.go  // Framework Bootstrap
- config
    - app.go        // Config Application
    - database.go   // Config Database And Redis 
    - server.go     // Config Server
- routes        // You Can Config Your Api Routers
     api.go     
- system        // Framework Files
- verndor       // Packages Installed By Go Mod
```

#### Add Route
 Add Your Route In `route/api.go`

```
route := new(Route.Route)

//Config Static Route
indexController := new(controllers.IndexController)
route.Get("/test", indexController.List)
route.Post("/test", indexController.Create)
route.Put("/test", indexController.Update)
route.Delete("/test", indexController.Delete)

```

#### Create Controller
Create `IndexController.go` In `app/controllers`
```
package controllers

import (
	"github.com/valyala/fasthttp"
	"github.com/wk331100/iFTY/system/http/response"
)

type IndexController struct {}

func (index *IndexController) List(ctx *fasthttp.RequestCtx){
	response.Json("Hello World", ctx)
}

```

#### Results
response structure
```
{
    "Code": 200,
    "Data": "Hello World",
    "Msg": "Success"
}
``
