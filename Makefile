.PHONY: bff
# generate bff -- make bff
api:
	goctl api go -api ./bff/bff.api -dir ./bff/

.PHONY: user
# generate rpc -- make user
rpc:
	goctl rpc protoc ./user/user.proto --go_out=./user --go-grpc_out=./user --zrpc_out=./user

.PHONY: post
# generate rpc -- make post
rpc:
	goctl rpc protoc ./post/post.proto --go_out=./post --go-grpc_out=./post --zrpc_out=./post

.PHONY: doc
# generate doc -- make doc
doc:
	goctl api plugin -plugin goctl-swagger="swagger -filename bff.json" -api ./bff/bff.api -dir ./doc


.PHONY: model
# generate model
# make model F=xxx.sql
model:$(F)
$(F):
	goctl model mysql ddl -src="./docs/sql/$(F)" -dir="./v1/rpc/dao" --style goZero
