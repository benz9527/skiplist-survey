cc_library(
    name = 'skiplist',
    hdrs = 'skiplist.h'
)

cc_binary(
    name = 'skiplist_test',
    srcs = 'skiplist_test.cc',
    deps = ':skiplist',
)

cc_benchmark(
    name = 'skiplist_benchmark',
    srcs = 'skiplist_benchmark.cc',
    deps = ':skiplist',
)

cc_benchmark(
    name = 'skiplist_survey',
    srcs = 'skiplist_survey.cc',
    deps = ':skiplist',
)
