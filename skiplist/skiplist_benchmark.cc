#include "skiplist/skiplist.h"

#include "thirdparty/benchmark/benchmark.h"

const int benchInitSize = 1000000;
const int benchBatchSize = 10;

SkipList<int, int> MakeSkipListN(int n) {
  SkipList<int, int> sl;
  for (int i = 0; i < n; ++i) {
    sl.Insert(i, i);
  }
  return sl;
}

std::map<int, int> MakeMapN(int n) {
  std::map<int, int> m;
  for (auto i = 0; i < n; ++i) {
    m[i] = i;
  }
  return m;
}

void BenchmarkSkipList_Insert(benchmark::State &state) {
  auto start = benchInitSize;
  state.PauseTiming();
  auto sl = MakeSkipListN(start);
  state.ResumeTiming();
  for (auto _ : state) {
    for (auto i = 0; i < benchBatchSize; i++) {
      sl.Insert(start + i, i);
    }
    start += benchBatchSize;
  }
}
BENCHMARK(BenchmarkSkipList_Insert);

void BenchmarkSkipList_Insert_Dup(benchmark::State &state) {
  auto start = benchInitSize;
  state.PauseTiming();
  auto sl = MakeSkipListN(start);
  state.ResumeTiming();
  for (auto _ : state) {
    for (auto i = 0; i < benchBatchSize; i++) {
      sl.Insert(i, i);
    }
    start += benchBatchSize;
  }
}
BENCHMARK(BenchmarkSkipList_Insert_Dup);

void BenchmarkMap_Insert_Dup(benchmark::State &state) {
  state.PauseTiming();
  auto m = MakeMapN(benchInitSize);
  state.ResumeTiming();
  for (auto _ : state) {
    for (auto i = 0; i < benchBatchSize; i++) {
      m[i] = i;
    }
  }
}
BENCHMARK(BenchmarkMap_Insert_Dup);

void BenchmarkMap_Find(benchmark::State &state) {
  state.PauseTiming();
  auto m = MakeMapN(benchInitSize);
  state.ResumeTiming();
  for (auto _ : state) {
    for (auto i = 0; i < benchBatchSize; i++) {
      int x = m[i];
      (void)x;
    }
  }
}
BENCHMARK(BenchmarkMap_Find);
