# protobuf benchmarks gogo protobuf vs go protobfu v1 v2

- github.com/gogo/protobuf v1.3.1
- github.com/golang/protobuf v1.4.3
- google.golang.org/protobuf v1.23.0

```shell script

# test
go test -v ./

# benchmark
go test -bench=. -run=none

```

## result

goos: windows
goarch: amd64
pkg: github.com/ikaiguang/protobuf_benchmarks
BenchmarkGoPbV1_Marshal-8                 922948              1236 ns/op               296 B/serial          320 B/op          1 allocs/op
BenchmarkGoPbV1_JSONMarshal-8              60913             19536 ns/op               797 B/serial         5513 B/op        141 allocs/op
BenchmarkGoPbV2_Marshal-8                1000000              1216 ns/op               296 B/serial          320 B/op          1 allocs/op
BenchmarkGoPbV2_JSONMarshal-8              60000             19900 ns/op               826 B/serial         4472 B/op        110 allocs/op
BenchmarkGoGoFast_Marshal-8               857203              1403 ns/op               296 B/serial          488 B/op          5 allocs/op
BenchmarkGoGoFast_JSONMarshal-8            18375             65252 ns/op               797 B/serial        25793 B/op        601 allocs/op
BenchmarkGoGo_Marshal-8                   857142              1376 ns/op               296 B/serial          488 B/op          5 allocs/op
BenchmarkGoGo_JSONMarshal-8                18291             65497 ns/op               797 B/serial        25793 B/op        601 allocs/op
PASS
ok      github.com/ikaiguang/protobuf_benchmarks        11.344s

