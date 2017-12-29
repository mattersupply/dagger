# Dagger

Dagger is a small set of Golang web services/API utilities and middlewares that can be re-used across different projects.

Currently included are:

## CORS middleware
Simple CORS Adapter using the standard CORS package.

## JSON middleware
Currently just writes the response header as JSON.
Future use cases include making sure that the data written to the response is actual JSON.

## JSON response writer addition
Simple function to write a payload of any interface to JSON via json.marshal.
Deals with errors resulting from such operations and writes the response codes to the ResponseWriter.

## Adapter
Interface and function to allow middlewares to be strung together and then called as a series of middlewares on handlers.
This allows us to create higher order functions that handlers can be passed into.
Inspired by https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
