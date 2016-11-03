#!/usr/bin/env bash
rm example_xmlrpc.go
rm xmlrpcgen
go build ../xmlrpcgen
go generate