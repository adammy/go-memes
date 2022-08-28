package meme

//go:generate oapi-codegen -old-config-style -templates=../../tools/oapi-codegen/templates -generate types -o types.gen.go -package meme ../../api/meme.yml
//go:generate oapi-codegen -old-config-style -templates=../../tools/oapi-codegen/templates -generate chi-server -o server.gen.go -package meme ../../api/meme.yml
//go:generate oapi-codegen -old-config-style -templates=../../tools/oapi-codegen/templates -generate client -o client.gen.go -package meme ../../api/meme.yml
//go:generate oapi-codegen -old-config-style -generate spec -o spec.gen.go -package meme ../../api/meme.yml
