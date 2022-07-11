go test -v -coverprofile=c.out -bench=.
go tool cover -html=c.out -o coverage.html