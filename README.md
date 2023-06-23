# logbench

## Benchmarch

> go test -benchmem -run NONE -bench .
goos: windows
goarch: amd64
pkg: logbench
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
BenchmarkDisabledWithoutFields/blackbear/log-16                 1000000000               0.4446 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledWithoutFields/apex/log-16                      1000000000               0.6035 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledWithoutFields/Zap-16                           1000000000               1.052 ns/op           0 B/op          0 allocs/op
BenchmarkDisabledWithoutFields/Zap.Check-16                     1000000000               0.8011 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledWithoutFields/Zap.Sugar-16                     91098477                14.28 ns/op           16 B/op          1 allocs/op
BenchmarkDisabledWithoutFields/Zap.SugarFormatting-16           12394338                94.60 ns/op          136 B/op          6 allocs/op
BenchmarkDisabledWithoutFields/sirupsen/logrus-16               91576671                13.20 ns/op           16 B/op          1 allocs/op
BenchmarkDisabledWithoutFields/rs/zerolog-16                    1000000000               0.5985 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledAccumulatedContext/blackbear/log-16            1000000000               0.4717 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledAccumulatedContext/apex/log-16                 1000000000               0.3330 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledAccumulatedContext/Zap-16                      1000000000               1.108 ns/op           0 B/op          0 allocs/op
BenchmarkDisabledAccumulatedContext/Zap.Check-16                1000000000               0.8785 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledAccumulatedContext/Zap.Sugar-16                83538442                14.11 ns/op           16 B/op          1 allocs/op
BenchmarkDisabledAccumulatedContext/Zap.SugarFormatting-16              12707731                92.10 ns/op          136 B/op          6 allocs/op
BenchmarkDisabledAccumulatedContext/sirupsen/logrus-16                  94267382                11.93 ns/op           16 B/op          1 allocs/op
BenchmarkDisabledAccumulatedContext/rs/zerolog-16                       1000000000               0.6085 ns/op          0 B/op          0 allocs/op
BenchmarkDisabledAddingFields/blackbear/log-16                          62610219                17.51 ns/op           24 B/op          1 allocs/op
BenchmarkDisabledAddingFields/apex/log-16                                2020345               547.6 ns/op           886 B/op         10 allocs/op
BenchmarkDisabledAddingFields/Zap-16                                     2202110               541.1 ns/op           736 B/op          5 allocs/op
BenchmarkDisabledAddingFields/Zap.Check-16                              1000000000               1.033 ns/op           0 B/op          0 allocs/op
BenchmarkDisabledAddingFields/Zap.Sugar-16                              10388600                97.20 ns/op          136 B/op          6 allocs/op
BenchmarkDisabledAddingFields/sirupsen/logrus-16                          985924              1147 ns/op            1525 B/op         12 allocs/op
BenchmarkDisabledAddingFields/rs/zerolog-16                             18026137                76.72 ns/op           96 B/op          4 allocs/op
BenchmarkWithoutFields/blackbear/log-16                                  4998319               210.8 ns/op             0 B/op          0 allocs/op
BenchmarkWithoutFields/apex/log-16                                        586538              1906 ns/op             240 B/op          5 allocs/op
BenchmarkWithoutFields/Zap-16                                           15072517                74.28 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutFields/Zap.Check-16                                     15044889                89.08 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutFields/Zap.CheckSampled-16                              37532488                32.58 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutFields/Zap.Sugar-16                                      6628117               162.9 ns/op            16 B/op          1 allocs/op
BenchmarkWithoutFields/Zap.SugarFormatting-16                             247341              5301 ns/op            1928 B/op         58 allocs/op
BenchmarkWithoutFields/go-kit/kit/log-16                                 2113780               611.7 ns/op           520 B/op          9 allocs/op
BenchmarkWithoutFields/inconshreveable/log15-16                            78669             14007 ns/op            1452 B/op         20 allocs/op
BenchmarkWithoutFields/sirupsen/logrus-16                                 210921              5416 ns/op            1142 B/op         23 allocs/op
BenchmarkWithoutFields/rs/zerolog-16                                    35753017                34.13 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutFields/rs/zerolog.Formatting-16                           246918              4211 ns/op            1916 B/op         58 allocs/op
BenchmarkWithoutFields/rs/zerolog.Check-16                              35194434                34.96 ns/op            0 B/op          0 allocs/op
BenchmarkAccumulatedContext/blackbear/log-16                             4502919               246.1 ns/op             0 B/op          0 allocs/op
BenchmarkAccumulatedContext/apex/log-16                                    47641             25925 ns/op            3312 B/op         53 allocs/op
BenchmarkAccumulatedContext/Zap-16                                      13039480                81.50 ns/op            0 B/op          0 allocs/op
BenchmarkAccumulatedContext/Zap.Check-16                                13418840                95.31 ns/op            0 B/op          0 allocs/op
BenchmarkAccumulatedContext/Zap.CheckSampled-16                         41988782                25.00 ns/op            0 B/op          0 allocs/op
BenchmarkAccumulatedContext/Zap.Sugar-16                                 8489426               130.9 ns/op            16 B/op          1 allocs/op
BenchmarkAccumulatedContext/Zap.SugarFormatting-16                        295059              4221 ns/op            1935 B/op         58 allocs/op
BenchmarkAccumulatedContext/go-kit/kit/log-16                             181990              6479 ns/op            3697 B/op         56 allocs/op
BenchmarkAccumulatedContext/inconshreveable/log15-16                       47438             24916 ns/op            3271 B/op         70 allocs/op
BenchmarkAccumulatedContext/sirupsen/logrus-16                             47217             26157 ns/op            4757 B/op         68 allocs/op
BenchmarkAccumulatedContext/rs/zerolog-16                               38352258                33.64 ns/op            0 B/op          0 allocs/op
BenchmarkAccumulatedContext/rs/zerolog.Check-16                         36227071                33.21 ns/op            0 B/op          0 allocs/op
BenchmarkAccumulatedContext/rs/zerolog.Formatting-16                      308427              3742 ns/op            1917 B/op         58 allocs/op
BenchmarkAddingFields/blackbear/log-16                                    398688              2822 ns/op            2124 B/op         32 allocs/op
BenchmarkAddingFields/apex/log-16                                          44703             26452 ns/op            4206 B/op         63 allocs/op
BenchmarkAddingFields/Zap-16                                             1000000              1379 ns/op             744 B/op          5 allocs/op
BenchmarkAddingFields/Zap.Check-16                                        877089              1425 ns/op             743 B/op          5 allocs/op
BenchmarkAddingFields/Zap.CheckSampled-16                                6310153               169.8 ns/op            84 B/op          0 allocs/op
BenchmarkAddingFields/Zap.Sugar-16                                        666703              1942 ns/op            1502 B/op         10 allocs/op
BenchmarkAddingFields/go-kit/kit/log-16                                   199542              5674 ns/op            3330 B/op         57 allocs/op
BenchmarkAddingFields/inconshreveable/log15-16                             42081             27727 ns/op            6729 B/op         74 allocs/op
BenchmarkAddingFields/sirupsen/logrus-16                                   43677             27821 ns/op            6302 B/op         79 allocs/op
BenchmarkAddingFields/rs/zerolog-16                                       330042              3570 ns/op            2577 B/op         32 allocs/op
BenchmarkAddingFields/rs/zerolog.Check-16                                 350140              3632 ns/op            2577 B/op         32 allocs/op
PASS
ok      logbench        81.689s