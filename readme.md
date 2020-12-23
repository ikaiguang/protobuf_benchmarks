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

## Results

2020-12-23 Results with go version go1.14.11 windows/amd64 on a 3.60 GHz Intel Core i7-7700 16GB

benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGoPbV1_Marshal-8                |     922948 |   1236 ns/op |   296 | 320 |   1.14 |   27319 |    3.86
BenchmarkGoPbV1_JSONMarshal-8            |      60913 |  19536 ns/op |   797 | 5513 |   1.19 |    4854 |    3.54
BenchmarkGoPbV2_Marshal-8                |    1000000 |   1216 ns/op |   296 | 320 |   1.22 |   29600 |    3.80
BenchmarkGoPbV2_JSONMarshal-8            |      60000 |  19900 ns/op |   826 | 4472 |   1.19 |    4956 |    4.45
BenchmarkGoGoFast_Marshal-8              |     857203 |   1403 ns/op |   296 | 488 |   1.20 |   25373 |    2.88
BenchmarkGoGoFast_JSONMarshal-8          |      18375 |  65252 ns/op |   797 | 25793 |   1.20 |    1464 |    2.53
BenchmarkGoGo_Marshal-8                  |     857142 |   1376 ns/op |   296 | 488 |   1.18 |   25371 |    2.82
BenchmarkGoGo_JSONMarshal-8              |      18291 |  65497 ns/op |   797 | 25793 |   1.20 |    1457 |    2.54


Totals:


benchmark                                    | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
---------------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
BenchmarkGoPbV2_-8                       |    1000000 |   1216 ns/op |   296 | 320 |   1.22 |   29600 |    3.80
BenchmarkGoPbV1_-8                       |     922948 |   1236 ns/op |   296 | 320 |   1.14 |   27319 |    3.86
BenchmarkGoGo_-8                         |     857142 |   1376 ns/op |   296 | 488 |   1.18 |   25371 |    2.82
BenchmarkGoGoFast_-8                     |     857203 |   1403 ns/op |   296 | 488 |   1.20 |   25373 |    2.88
BenchmarkGoPbV1_JSON-8                   |      60913 |  19536 ns/op |   797 | 5513 |   1.19 |    4854 |    3.54
BenchmarkGoPbV2_JSON-8                   |      60000 |  19900 ns/op |   826 | 4472 |   1.19 |    4956 |    4.45
BenchmarkGoGoFast_JSON-8                 |      18375 |  65252 ns/op |   797 | 25793 |   1.20 |    1464 |    2.53
BenchmarkGoGo_JSON-8                     |      18291 |  65497 ns/op |   797 | 25793 |   1.20 |    1457 |    2.54
