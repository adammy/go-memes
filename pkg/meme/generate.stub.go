//go:build stub
// +build stub

package meme

//go:generate oapi-codegen -old-config-style -templates=../../tools/oapi-codegen/stub-templates -generate chi-server -o server.go -package meme ../../api/meme.yml
