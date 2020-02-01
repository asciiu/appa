module github.com/asciiu/appa/api

go 1.13.3

replace github.com/asciiu/appa/trade-engine => ../trade-engine

replace github.com/asciiu/appa/lib => ../lib

require (
	github.com/asciiu/appa/lib v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.14
	github.com/micro/go-micro v1.18.0
	golang.org/x/crypto v0.0.0-20200128174031-69ecbb4d6d5d
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa
)
