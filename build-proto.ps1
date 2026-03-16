$ErrorActionPreference = 'Stop'

$gopath = (go env GOPATH)
$env:PATH = $env:PATH + ';' + (Join-Path $gopath 'bin')

$protoRoot = Join-Path $PSScriptRoot 'proto'
$outDir = Join-Path $protoRoot 'generated'

$files = Get-ChildItem -Path $protoRoot -Recurse -Filter *.proto -File |
  Where-Object { $_.FullName -notmatch '\\proto\\generated\\' } |
  ForEach-Object { $_.FullName }

protoc --experimental_allow_proto3_optional `
  --go_out=$outDir --go_opt=paths=source_relative `
  --go-grpc_out=$outDir --go-grpc_opt=paths=source_relative `
  --proto_path=$protoRoot `
  $files

Write-Host "Generated protobuf code into $outDir"