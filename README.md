# gRPC

## Prepare
### Copy google/api/annotations.proto and google/api/http.proto
Refer:  
https://github.com/grpc-ecosystem/grpc-gateway/issues/1935  
```
mkdir -p google/api
wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -O google/api/annotations.proto
wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -O google/api/http.proto
```

## Generate
Refer:  
https://www.kuiki.cn/protobuf-import-package/  
```
protoc -I "./" -I "../../../" \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/fetchnews/kabar/v1/fetchnews.proto
go mod tidy
```

# Clone the project
Just clone and run the command to replace all string in all files, that contains `dwnews`:
```
cd /path/to/your/folder
sed -i 's/foo/bar/g' *
sed -i '' 's/kabar/reuters/g' *.mod *.md ./internal/*/* ./cmd/*/* ./configs/* ./api/*/*/*/*
```
rm git
```
rm -rf .git
```
git init
```
git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/<username>/<repo>.git
git push -u origin main
```
grpc generate
```
protoc -I "./" -I "../../../" \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/fetchnews/kabar/v1/fetchnews.proto
```
go mod tidy
```
go mod tidy
```
