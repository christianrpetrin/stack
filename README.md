# stack [![Build Status](https://travis-ci.com/ef-ds/stack.svg?branch=master)](https://travis-ci.com/ef-ds/stack) [![codecov](https://codecov.io/gh/ef-ds/stack/branch/master/graph/badge.svg)](https://codecov.io/gh/ef-ds/stack) [![Go Report Card](https://goreportcard.com/badge/github.com/ef-ds/stack)](https://goreportcard.com/report/github.com/ef-ds/stack)  [![GoDoc](https://godoc.org/github.com/ef-ds/stack?status.svg)](https://godoc.org/github.com/ef-ds/stack)

Package stack updates the [ef-ds stack](https://github.com/ef-ds/stack) to operate on [*TestValue](testvaluestack.go) struct instead of the default "interface{}" type.

## Why

This is a test to check the performance impact of a data structure that operates on "interface{}" vs the final, concrete type.

Below commands can be used to check the performance difference of the default interface vs the new TestValue stack.

Run below from the stack main folder.
```
benchstat testdata/BenchmarkMicroserviceInterfaceStack.txt testdata/BenchmarkMicroserviceTestValueStack.txt
benchstat testdata/BenchmarkFillInterfaceStack.txt testdata/BenchmarkFillTestValueStack.txt
benchstat testdata/BenchmarkRefillInterfaceStack.txt testdata/BenchmarkRefillTestValueStack.txt
benchstat testdata/BenchmarkRefillFullInterfaceStack.txt testdata/BenchmarkRefillFullTestValueStack.txt
benchstat testdata/BenchmarkSlowIncreaseInterfaceStack.txt testdata/BenchmarkSlowIncreaseTestValueStack.txt
benchstat testdata/BenchmarkSlowDecreaseInterfaceStack.txt testdata/BenchmarkSlowDecreaseTestValueStack.txt
benchstat testdata/BenchmarkStableInterfaceStack.txt testdata/BenchmarkStableTestValueStack.txt
```

## Results

```
benchstat testdata/BenchmarkMicroserviceInterfaceStack.txt testdata/BenchmarkMicroserviceTestValueStack.txt
name        old time/op    new time/op    delta
/0-4          4.35ns ± 0%    4.36ns ± 1%     ~     (p=0.728 n=10+10)
/1-4           282ns ± 1%     274ns ± 1%   -2.88%  (p=0.000 n=9+10)
/10-4         2.03µs ± 0%    1.94µs ± 0%   -4.30%  (p=0.000 n=9+10)
/100-4        18.5µs ± 0%    17.7µs ± 1%   -4.75%  (p=0.000 n=9+10)
/1000-4        175µs ± 1%     168µs ± 1%   -3.91%  (p=0.000 n=9+10)
/10000-4      1.79ms ± 1%    1.71ms ± 1%   -4.60%  (p=0.000 n=9+9)
/100000-4     19.9ms ± 1%    18.7ms ± 1%   -6.11%  (p=0.000 n=9+10)
/1000000-4     207ms ± 3%     206ms ± 3%     ~     (p=0.529 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           0.00B          0.00B          ~     (all equal)
/1-4            224B ± 0%      192B ± 0%  -14.29%  (p=0.000 n=10+10)
/10-4         1.49kB ± 0%    1.33kB ± 0%  -10.75%  (p=0.000 n=10+10)
/100-4        16.7kB ± 0%    14.0kB ± 0%  -16.30%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     123kB ± 0%   -8.15%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.21MB ± 0%   -6.56%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    12.0MB ± 0%   -6.29%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     120MB ± 0%   -6.25%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            0.00           0.00          ~     (all equal)
/1-4            9.00 ± 0%      9.00 ± 0%     ~     (all equal)
/10-4           73.0 ± 0%      73.0 ± 0%     ~     (all equal)
/100-4           705 ± 0%       705 ± 0%     ~     (all equal)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.0k ± 0%     70.0k ± 0%     ~     (all equal)
/100000-4       700k ± 0%      700k ± 0%     ~     (all equal)
/1000000-4     7.00M ± 0%     7.00M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkFillInterfaceStack.txt testdata/BenchmarkFillTestValueStack.txt
name        old time/op    new time/op    delta
/0-4          1.17ns ± 3%    1.45ns ± 1%  +24.19%  (p=0.000 n=10+10)
/1-4           127ns ± 1%     119ns ± 1%   -6.23%  (p=0.000 n=9+10)
/10-4          463ns ± 1%     421ns ± 1%   -9.04%  (p=0.000 n=10+9)
/100-4        3.86µs ± 1%    3.31µs ± 0%  -14.28%  (p=0.000 n=10+10)
/1000-4       29.0µs ± 1%    26.8µs ± 1%   -7.73%  (p=0.000 n=10+10)
/10000-4       279µs ± 1%     263µs ± 1%   -5.75%  (p=0.000 n=9+9)
/100000-4     3.05ms ± 2%    3.46ms ± 8%  +13.41%  (p=0.000 n=10+10)
/1000000-4    74.3ms ±42%    85.7ms ±53%     ~     (p=0.353 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           0.00B          0.00B          ~     (all equal)
/1-4            128B ± 0%       96B ± 0%  -25.00%  (p=0.000 n=10+10)
/10-4           528B ± 0%      368B ± 0%  -30.30%  (p=0.000 n=10+10)
/100-4        7.09kB ± 0%    4.37kB ± 0%  -38.37%  (p=0.000 n=10+10)
/1000-4       37.9kB ± 0%    27.0kB ± 0%  -28.81%  (p=0.000 n=10+10)
/10000-4       330kB ± 0%     245kB ± 0%  -25.67%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    2.41MB ± 0%  -25.05%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    24.1MB ± 0%  -24.97%  (p=0.002 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            0.00           0.00          ~     (all equal)
/1-4            3.00 ± 0%      3.00 ± 0%     ~     (all equal)
/10-4           13.0 ± 0%      13.0 ± 0%     ~     (all equal)
/100-4           105 ± 0%       105 ± 0%     ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkRefillInterfaceStack.txt testdata/BenchmarkRefillTestValueStack.txt
name        old time/op    new time/op    delta
/0-4          16.3ns ± 1%    13.0ns ± 0%  -19.97%  (p=0.000 n=10+9)
/1-4           275ns ± 1%     262ns ± 1%   -4.62%  (p=0.000 n=10+9)
/10-4         2.60µs ± 1%    2.55µs ± 2%   -2.04%  (p=0.000 n=10+10)
/100-4        25.2µs ± 2%    24.8µs ± 0%   -1.84%  (p=0.000 n=10+9)
/1000-4        249µs ± 1%     246µs ± 2%     ~     (p=0.052 n=10+10)
/10000-4      2.59ms ± 4%    3.08ms ± 7%  +19.27%  (p=0.000 n=10+10)
/100000-4     40.3ms ± 5%    27.7ms ± 8%  -31.28%  (p=0.000 n=9+10)
/1000000-4     485ms ±29%     387ms ±10%  -20.16%  (p=0.003 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           0.00B          0.00B          ~     (all equal)
/1-4            160B ± 0%      160B ± 0%     ~     (all equal)
/10-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/100-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/1000-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.000 n=10+10)
/10000-4      1.60MB ± 0%    1.60MB ± 0%   -0.01%  (p=0.000 n=10+10)
/100000-4     16.0MB ± 0%    16.0MB ± 0%   -0.15%  (p=0.000 n=10+8)
/1000000-4     166MB ± 2%     163MB ± 0%   -1.84%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            0.00           0.00          ~     (all equal)
/1-4            10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/10-4            100 ± 0%       100 ± 0%     ~     (all equal)
/100-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/1000-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/10000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/100000-4      1.00M ± 0%     1.00M ± 0%     ~     (p=0.087 n=10+10)
/1000000-4     10.0M ± 0%     10.0M ± 0%     ~     (p=0.786 n=10+9)
```

```
benchstat testdata/BenchmarkRefillFullInterfaceStack.txt testdata/BenchmarkRefillFullTestValueStack.txt
name        old time/op    new time/op    delta
/0-4          13.4ns ± 1%    13.1ns ± 1%   -2.54%  (p=0.000 n=10+10)
/1-4           293ns ± 3%     306ns ± 7%   +4.33%  (p=0.001 n=10+10)
/10-4         2.84µs ± 5%    3.08µs ± 5%   +8.43%  (p=0.000 n=10+10)
/100-4        27.0µs ± 1%    29.3µs ± 6%   +8.43%  (p=0.000 n=9+10)
/1000-4        272µs ± 3%     300µs ± 9%  +10.42%  (p=0.000 n=10+10)
/10000-4      2.70ms ± 1%    2.76ms ± 4%   +2.17%  (p=0.011 n=10+10)
/100000-4     37.8ms ± 4%    42.1ms ± 4%  +11.56%  (p=0.000 n=10+10)
/1000000-4     632ms ± 1%     385ms ±13%  -39.05%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           0.00B          0.00B          ~     (all equal)
/1-4            160B ± 0%      160B ± 0%     ~     (all equal)
/10-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/100-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/1000-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/10000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/100000-4     16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)
/1000000-4     160MB ± 0%     160MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/0-4            0.00           0.00          ~     (all equal)
/1-4            10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/10-4            100 ± 0%       100 ± 0%     ~     (all equal)
/100-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/1000-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/10000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/100000-4      1.00M ± 0%     1.00M ± 0%     ~     (all equal)
/1000000-4     10.0M ± 0%     10.0M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkSlowIncreaseInterfaceStack.txt testdata/BenchmarkSlowIncreaseTestValueStack.txt
name        old time/op    new time/op    delta
/0-4          1.16ns ± 0%    1.16ns ± 2%     ~     (p=0.776 n=8+9)
/1-4           151ns ± 1%     144ns ± 0%   -4.76%  (p=0.000 n=9+10)
/10-4          725ns ± 2%     682ns ± 1%   -5.96%  (p=0.000 n=9+10)
/100-4        6.48µs ± 4%    5.70µs ± 1%  -12.11%  (p=0.000 n=10+10)
/1000-4       56.3µs ± 4%    50.4µs ± 1%  -10.56%  (p=0.000 n=10+10)
/10000-4       554µs ± 8%     508µs ± 1%   -8.41%  (p=0.000 n=10+10)
/100000-4     7.93ms ±10%    4.94ms ± 1%  -37.66%  (p=0.000 n=10+10)
/1000000-4    76.7ms ±12%    73.1ms ±17%     ~     (p=0.387 n=9+9)

name        old alloc/op   new alloc/op   delta
/0-4           0.00B          0.00B          ~     (all equal)
/1-4            144B ± 0%      112B ± 0%  -22.22%  (p=0.000 n=10+10)
/10-4           688B ± 0%      528B ± 0%  -23.26%  (p=0.000 n=10+10)
/100-4        8.69kB ± 0%    5.97kB ± 0%  -31.31%  (p=0.000 n=10+10)
/1000-4       53.9kB ± 0%    43.0kB ± 0%  -20.26%  (p=0.000 n=10+10)
/10000-4       490kB ± 0%     405kB ± 0%  -17.28%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.01MB ± 0%  -16.73%  (p=0.000 n=10+9)
/1000000-4    48.1MB ± 0%    40.1MB ± 0%  -16.66%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            0.00           0.00          ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           23.0 ± 0%      23.0 ± 0%     ~     (all equal)
/100-4           205 ± 0%       205 ± 0%     ~     (all equal)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.0k ± 0%     20.0k ± 0%     ~     (all equal)
/100000-4       200k ± 0%      200k ± 0%     ~     (all equal)
/1000000-4     2.00M ± 0%     2.00M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkSlowDecreaseInterfaceStack.txt testdata/BenchmarkSlowDecreaseTestValueStack.txt
name        old time/op    new time/op    delta
/0-4          0.89ns ± 6%    0.58ns ± 0%  -35.05%  (p=0.000 n=10+8)
/1-4          30.5ns ± 7%    28.2ns ± 1%   -7.48%  (p=0.000 n=9+10)
/10-4          299ns ± 2%     287ns ± 1%   -3.95%  (p=0.000 n=10+10)
/100-4        2.90µs ± 1%    2.82µs ± 2%   -2.70%  (p=0.000 n=9+9)
/1000-4       28.3µs ± 1%    27.9µs ± 1%   -1.54%  (p=0.000 n=9+10)
/10000-4       283µs ± 1%     279µs ± 0%   -1.47%  (p=0.000 n=10+8)
/100000-4     2.82ms ± 1%    2.78ms ± 1%   -1.25%  (p=0.000 n=10+10)
/1000000-4    28.1ms ± 1%    27.8ms ± 0%   -1.19%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/0-4           0.00B          0.00B          ~     (all equal)
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/0-4            0.00           0.00          ~     (all equal)
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkStableInterfaceStack.txt testdata/BenchmarkStableTestValueStack.txt
name        old time/op    new time/op    delta
/0-4          0.58ns ± 0%    0.87ns ± 1%  +50.69%  (p=0.000 n=9+10)
/1-4          28.0ns ± 3%    29.2ns ± 8%   +4.43%  (p=0.014 n=10+10)
/10-4          280ns ± 3%     294ns ± 6%   +4.96%  (p=0.000 n=9+9)
/100-4        2.75µs ± 2%    2.88µs ± 4%   +4.62%  (p=0.000 n=10+9)
/1000-4       27.4µs ± 4%    30.1µs ± 4%   +9.84%  (p=0.000 n=10+10)
/10000-4       275µs ± 3%     286µs ± 6%   +3.77%  (p=0.009 n=10+10)
/100000-4     2.73ms ± 4%    2.80ms ± 5%     ~     (p=0.105 n=10+10)
/1000000-4    27.4ms ± 3%    28.4ms ± 7%     ~     (p=0.052 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           0.00B          0.00B          ~     (all equal)
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/0-4            0.00           0.00          ~     (all equal)
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```
