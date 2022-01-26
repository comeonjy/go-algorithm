## TODO

### 排序算法
```
➜  algorithm git:(master) ✗ go test -v -bench=. -benchmem sort_test.go -test.run BenchmarkSort
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkSort
BenchmarkSort/QuickSort
BenchmarkSort/QuickSort-8                   1730            665652 ns/op               0 B/op          0 allocs/op
BenchmarkSort/sort.Ints
BenchmarkSort/sort.Ints-8                   1074           1105999 ns/op              24 B/op          1 allocs/op
BenchmarkSort/MergeSort
BenchmarkSort/MergeSort-8                    428           3103434 ns/op         3471622 B/op      33089 allocs/op
BenchmarkSort/InsertSort
BenchmarkSort/InsertSort-8                    66          16490736 ns/op               0 B/op          0 allocs/op
BenchmarkSort/SelectSort
BenchmarkSort/SelectSort-8                    12          95424101 ns/op               0 B/op          0 allocs/op
BenchmarkSort/BubbleSort
BenchmarkSort/BubbleSort-8                    10         110265651 ns/op               0 B/op          0 allocs/op
PASS
ok      command-line-arguments  8.711s

```