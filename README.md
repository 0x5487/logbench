# logbench

## Benchmarch

Date: 2023-06-23

```go
goos: linux
goarch: amd64
pkg: logbench
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
BenchmarkDisabledWithoutFields
    /workspaces/logbench/scenario_bench_test.go:30: Logging at a disabled level without any structured context.
BenchmarkDisabledWithoutFields/blackbear/log
BenchmarkDisabledWithoutFields/blackbear/log-16          1000000000          0.4582 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledWithoutFields/apex/log
BenchmarkDisabledWithoutFields/apex/log-16               1000000000          0.6236 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledWithoutFields/Zap
BenchmarkDisabledWithoutFields/Zap-16                    1000000000          1.060 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledWithoutFields/Zap.Check
BenchmarkDisabledWithoutFields/Zap.Check-16              1000000000          0.8960 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledWithoutFields/Zap.Sugar
BenchmarkDisabledWithoutFields/Zap.Sugar-16              83759577         14.52 ns/op       16 B/op        1 allocs/op
BenchmarkDisabledWithoutFields/Zap.SugarFormatting
BenchmarkDisabledWithoutFields/Zap.SugarFormatting-16    10367266        124.1 ns/op      136 B/op        6 allocs/op
BenchmarkDisabledWithoutFields/sirupsen/logrus
BenchmarkDisabledWithoutFields/sirupsen/logrus-16        74684968         14.84 ns/op       16 B/op        1 allocs/op
BenchmarkDisabledWithoutFields/rs/zerolog
BenchmarkDisabledWithoutFields/rs/zerolog-16             1000000000          0.6606 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledAccumulatedContext
    /workspaces/logbench/scenario_bench_test.go:110: Logging at a disabled level with some accumulated context.
BenchmarkDisabledAccumulatedContext/blackbear/log
BenchmarkDisabledAccumulatedContext/blackbear/log-16     1000000000          0.6594 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledAccumulatedContext/apex/log
BenchmarkDisabledAccumulatedContext/apex/log-16          1000000000          0.4470 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledAccumulatedContext/Zap
BenchmarkDisabledAccumulatedContext/Zap-16               885968905          1.315 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledAccumulatedContext/Zap.Check
BenchmarkDisabledAccumulatedContext/Zap.Check-16         1000000000          0.8960 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledAccumulatedContext/Zap.Sugar
BenchmarkDisabledAccumulatedContext/Zap.Sugar-16         79490601         15.35 ns/op       16 B/op        1 allocs/op
BenchmarkDisabledAccumulatedContext/Zap.SugarFormatting
BenchmarkDisabledAccumulatedContext/Zap.SugarFormatting-16          10064905        122.9 ns/op      136 B/op        6 allocs/op
BenchmarkDisabledAccumulatedContext/sirupsen/logrus
BenchmarkDisabledAccumulatedContext/sirupsen/logrus-16              68148162         16.14 ns/op       16 B/op        1 allocs/op
BenchmarkDisabledAccumulatedContext/rs/zerolog
BenchmarkDisabledAccumulatedContext/rs/zerolog-16                   1000000000          0.6394 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledAddingFields
    /workspaces/logbench/scenario_bench_test.go:190: Logging at a disabled level, adding context at each log site.
BenchmarkDisabledAddingFields/blackbear/log
BenchmarkDisabledAddingFields/blackbear/log-16                      45396289         24.09 ns/op       24 B/op        1 allocs/op
BenchmarkDisabledAddingFields/apex/log
BenchmarkDisabledAddingFields/apex/log-16                            1646732        712.1 ns/op      886 B/op       10 allocs/op
BenchmarkDisabledAddingFields/Zap
BenchmarkDisabledAddingFields/Zap-16                                 2871567        452.2 ns/op      736 B/op        5 allocs/op
BenchmarkDisabledAddingFields/Zap.Check
BenchmarkDisabledAddingFields/Zap.Check-16                          1000000000          1.052 ns/op        0 B/op        0 allocs/op
BenchmarkDisabledAddingFields/Zap.Sugar
BenchmarkDisabledAddingFields/Zap.Sugar-16                           8157687        132.5 ns/op      136 B/op        6 allocs/op
BenchmarkDisabledAddingFields/sirupsen/logrus
BenchmarkDisabledAddingFields/sirupsen/logrus-16                      846895       1321 ns/op     1526 B/op       12 allocs/op
BenchmarkDisabledAddingFields/rs/zerolog
BenchmarkDisabledAddingFields/rs/zerolog-16                         11987624         88.59 ns/op       96 B/op        4 allocs/op
BenchmarkShortText
    /workspaces/logbench/scenario_bench_test.go:264: Logging with short text
BenchmarkShortText/blackbear/log
BenchmarkShortText/blackbear/log-16                                  5068552        225.9 ns/op        0 B/op        0 allocs/op
BenchmarkShortText/rs/zerolog
BenchmarkShortText/rs/zerolog-16                                    20440348         53.84 ns/op        0 B/op        0 allocs/op
BenchmarkShortText/Zap
BenchmarkShortText/Zap-16                                            6118238        184.8 ns/op        0 B/op        0 allocs/op
BenchmarkLongText
    /workspaces/logbench/scenario_bench_test.go:299: Logging with long text
BenchmarkLongText/blackbear/log
BenchmarkLongText/blackbear/log-16                                   1425358        856.3 ns/op        0 B/op        0 allocs/op
BenchmarkLongText/rs/zerolog
BenchmarkLongText/rs/zerolog-16                                       788925       1291 ns/op        0 B/op        0 allocs/op
BenchmarkLongText/Zap
BenchmarkLongText/Zap-16                                              179590       6370 ns/op        4 B/op        0 allocs/op
BenchmarkWithoutFields
    /workspaces/logbench/scenario_bench_test.go:334: Logging without any structured context.
BenchmarkWithoutFields/blackbear/log
BenchmarkWithoutFields/blackbear/log-16                              5631736        192.1 ns/op        0 B/op        0 allocs/op
BenchmarkWithoutFields/apex/log
BenchmarkWithoutFields/apex/log-16                                    694462       1584 ns/op      240 B/op        5 allocs/op
BenchmarkWithoutFields/Zap
BenchmarkWithoutFields/Zap-16                                       14240935         93.61 ns/op        0 B/op        0 allocs/op
BenchmarkWithoutFields/Zap.Check
BenchmarkWithoutFields/Zap.Check-16                                 11751316        101.8 ns/op        0 B/op        0 allocs/op
BenchmarkWithoutFields/Zap.CheckSampled
BenchmarkWithoutFields/Zap.CheckSampled-16                          35724284         31.22 ns/op        0 B/op        0 allocs/op
BenchmarkWithoutFields/Zap.Sugar
BenchmarkWithoutFields/Zap.Sugar-16                                  9140786        128.9 ns/op       16 B/op        1 allocs/op
BenchmarkWithoutFields/Zap.SugarFormatting
BenchmarkWithoutFields/Zap.SugarFormatting-16                         307940       3858 ns/op     1936 B/op       58 allocs/op
BenchmarkWithoutFields/go-kit/kit/log
BenchmarkWithoutFields/go-kit/kit/log-16                             2280416        539.3 ns/op      521 B/op        9 allocs/op
BenchmarkWithoutFields/inconshreveable/log15
BenchmarkWithoutFields/inconshreveable/log15-16                       253660       4776 ns/op     1450 B/op       20 allocs/op
BenchmarkWithoutFields/sirupsen/logrus
BenchmarkWithoutFields/sirupsen/logrus-16                             350174       3409 ns/op     1132 B/op       23 allocs/op
BenchmarkWithoutFields/rs/zerolog
BenchmarkWithoutFields/rs/zerolog-16                                33204984         33.50 ns/op        0 B/op        0 allocs/op
BenchmarkWithoutFields/rs/zerolog.Formatting
BenchmarkWithoutFields/rs/zerolog.Formatting-16                       376706       3224 ns/op     1917 B/op       58 allocs/op
BenchmarkWithoutFields/rs/zerolog.Check
BenchmarkWithoutFields/rs/zerolog.Check-16                          31966191         35.06 ns/op        0 B/op        0 allocs/op
BenchmarkAccumulatedContext
    /workspaces/logbench/scenario_bench_test.go:469: Logging with some accumulated context.
BenchmarkAccumulatedContext/blackbear/log
BenchmarkAccumulatedContext/blackbear/log-16                         5906029        189.6 ns/op        0 B/op        0 allocs/op
BenchmarkAccumulatedContext/apex/log
BenchmarkAccumulatedContext/apex/log-16                                55074      21861 ns/op     3304 B/op       53 allocs/op
BenchmarkAccumulatedContext/Zap
BenchmarkAccumulatedContext/Zap-16                                  13164829         94.57 ns/op        0 B/op        0 allocs/op
BenchmarkAccumulatedContext/Zap.Check
BenchmarkAccumulatedContext/Zap.Check-16                            11359444         97.27 ns/op        0 B/op        0 allocs/op
BenchmarkAccumulatedContext/Zap.CheckSampled
BenchmarkAccumulatedContext/Zap.CheckSampled-16                     34143926         30.69 ns/op        0 B/op        0 allocs/op
BenchmarkAccumulatedContext/Zap.Sugar
BenchmarkAccumulatedContext/Zap.Sugar-16                             8790692        126.7 ns/op       16 B/op        1 allocs/op
BenchmarkAccumulatedContext/Zap.SugarFormatting
BenchmarkAccumulatedContext/Zap.SugarFormatting-16                    313186       3963 ns/op     1952 B/op       58 allocs/op
BenchmarkAccumulatedContext/go-kit/kit/log
BenchmarkAccumulatedContext/go-kit/kit/log-16                         189099       5825 ns/op     3700 B/op       56 allocs/op
BenchmarkAccumulatedContext/inconshreveable/log15
BenchmarkAccumulatedContext/inconshreveable/log15-16                   46414      25235 ns/op     3300 B/op       70 allocs/op
BenchmarkAccumulatedContext/sirupsen/logrus
BenchmarkAccumulatedContext/sirupsen/logrus-16                         49592      23228 ns/op     4723 B/op       68 allocs/op
BenchmarkAccumulatedContext/rs/zerolog
BenchmarkAccumulatedContext/rs/zerolog-16                           32183832         37.15 ns/op        0 B/op        0 allocs/op
BenchmarkAccumulatedContext/rs/zerolog.Check
BenchmarkAccumulatedContext/rs/zerolog.Check-16                     28383674         37.95 ns/op        0 B/op        0 allocs/op
BenchmarkAccumulatedContext/rs/zerolog.Formatting
BenchmarkAccumulatedContext/rs/zerolog.Formatting-16                  342412       3298 ns/op     1917 B/op       58 allocs/op
BenchmarkAddingFields
    /workspaces/logbench/scenario_bench_test.go:603: Logging with additional context at each log site.
BenchmarkAddingFields/blackbear/log
BenchmarkAddingFields/blackbear/log-16                                373516       3137 ns/op     2127 B/op       32 allocs/op
BenchmarkAddingFields/apex/log
BenchmarkAddingFields/apex/log-16                                      46994      24019 ns/op     4196 B/op       63 allocs/op
BenchmarkAddingFields/Zap
BenchmarkAddingFields/Zap-16                                          982134       1444 ns/op      745 B/op        5 allocs/op
BenchmarkAddingFields/Zap.Check
BenchmarkAddingFields/Zap.Check-16                                    764864       1517 ns/op      744 B/op        5 allocs/op
BenchmarkAddingFields/Zap.CheckSampled
BenchmarkAddingFields/Zap.CheckSampled-16                            5244549        206.5 ns/op       86 B/op        0 allocs/op
BenchmarkAddingFields/Zap.Sugar
BenchmarkAddingFields/Zap.Sugar-16                                    439082       2303 ns/op     1513 B/op       10 allocs/op
BenchmarkAddingFields/go-kit/kit/log
BenchmarkAddingFields/go-kit/kit/log-16                               200570       5252 ns/op     3334 B/op       57 allocs/op
BenchmarkAddingFields/inconshreveable/log15
BenchmarkAddingFields/inconshreveable/log15-16                         40218      29120 ns/op     6457 B/op       74 allocs/op
BenchmarkAddingFields/sirupsen/logrus
BenchmarkAddingFields/sirupsen/logrus-16                               45156      26164 ns/op     6249 B/op       79 allocs/op
BenchmarkAddingFields/rs/zerolog
BenchmarkAddingFields/rs/zerolog-16                                   302256       3621 ns/op     2532 B/op       32 allocs/op
BenchmarkAddingFields/rs/zerolog.Check
BenchmarkAddingFields/rs/zerolog.Check-16                             290060       3781 ns/op     2531 B/op       32 allocs/op
PASS
 logbench coverage: [no statements]
ok   logbench 83.549s
```
