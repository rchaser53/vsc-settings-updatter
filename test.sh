files="$(find . -name "*.go" | grep -v _test.go)"
go run $files