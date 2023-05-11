protoc -I ./proto \
  --proto_path=${ANNOTATIONS} \
  --proto_path=./proto/hello.proto \
  --swagger_out=logtostderr=true:. \
  ./proto/hello.proto