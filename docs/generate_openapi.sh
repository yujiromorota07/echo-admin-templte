#!/bin/sh
oapi-codegen -include-tags login -generate "types" -package login main.yaml > "../api/generated/login/types_login.gen.go"
oapi-codegen -include-tags login -generate "server" -package login main.yaml > "../api/generated/login/servers_login.gen.go"

oapi-codegen -include-tags account -generate "types" -package account main.yaml > "../api/generated/account/types_account.gen.go"
oapi-codegen -include-tags account -generate "server" -package account main.yaml > "../api/generated/account/servers_account.gen.go"