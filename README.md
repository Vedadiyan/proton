![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![Go report](https://goreportcard.com/badge/github.com/vedadiyan/proton)](https://goreportcard.com/report/github.com/vedadiyan/proton)

## A Protobuffer Querying Language

Proton is a "Protobuffer Querying Language" designed to provide automatic mapping between Protobuffer messages and other data structures. The original motivation behind creating Proton was to minimize development effort needed for integrating third party RESTful APIs in Go microservices. The idea is to leave modelling and mapping data to data engineers so that developers could focus more on getting the microservice ready fast. 

## Proton is compiled by Go AST package 
Proton follows the exact same syntax as Go and it is compiled by the Go AST package. It has been designed so because the Go AST is a trustworthy package and no complexities relating to code compiling would be introduced to Proton's code base. 
## Proton is Extensible 
Proton consists of a frontend that can be controlled by the user. This means that more functions can be injected to Proton when necessary. Only Go language is supported for injecting new functions. 

# Documentation 
Proton allows mapping between different kinds of field types and it uses two approaches for mapping fields: 

 1. Select: the `select` attribute allows for focusing the read scope on a particular object 
 2. Query: the `query` attribute provides a SQL like syntax for making complex mapping 

To be continued...

 
