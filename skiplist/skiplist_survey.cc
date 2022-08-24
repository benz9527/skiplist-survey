#include "skiplist/skiplist.h"
#include <map>
#include <sys/time.h>

const int start = 100000;
const int end = 3000000;
const int step = 100000;

struct Bytes {
  const char *begin;
  size_t size;
  size_t capacity;
};

Bytes testByteString{"test value", 10, 10};

const int64_t NSINS = 1000000000LL;

int64_t ts_to_ns(timespec ts) { return ts.tv_sec * NSINS + ts.tv_nsec; }

class TimerTracker {
public:
  explicit TimerTracker(int n) : n(n) {
    clock_gettime(CLOCK_MONOTONIC, &start);
  }
  ~TimerTracker() {
    clock_gettime(CLOCK_MONOTONIC, &end);
    int64_t ns_start = ts_to_ns(start);
    int64_t ns_end = ts_to_ns(end);
    int64_t ns_cost = ns_end - ns_start;
    printf("%lld", ns_cost / n);
  }

private:
  int n;
  struct timespec start, end;
};

template <class Tp> inline void DoNotOptimize(Tp const &value) {
  asm volatile("" : : "r,m"(value) : "memory");
}

void chen3fengInserts(int n) {
  SkipList<int, Bytes> list;
  TimerTracker tt(n);

  for (int i = 0; i < n; i++) {
    list.Insert(n - i, testByteString);
  }
}

void chen3fengWorstInserts(int n) {
  SkipList<int, Bytes> list;
  TimerTracker tt(n);
  for (int i = 0; i < n; i++) {
    list.Insert(i, testByteString);
  }
}

void chen3fengAvgSearch(int n) {
  SkipList<int, Bytes> list;
  for (int i = 0; i < n; i++) {
    list.Insert(i, testByteString);
  }

  TimerTracker tt(n);
  for (int i = 0; i < n; i++) {
    DoNotOptimize(list.Find(i));
  }
}

void chen3fengSearchEnd(int n) {
  SkipList<int, Bytes> list;
  for (int i = 0; i < n; i++) {
    list.Insert(i, testByteString);
  }

  TimerTracker tt(n);
  for (int i = 0; i < n; i++) {
    DoNotOptimize(list.Find(n));
  }
}

void chen3fengDelete(int n) {
  SkipList<int, Bytes> list;
  for (int i = 0; i < n; i++) {
    list.Insert(i, testByteString);
  }

  TimerTracker tt(n);

  for (int i = 0; i < n; i++) {
    list.Remove(i);
  }
}

void chen3fengWorstDelete(int n) {
  SkipList<int, Bytes> list;
  for (int i = 0; i < n; i++) {
    list.Insert(i, testByteString);
  }
  TimerTracker tt(n);

  for (int i = 0; i < n; i++) {
    list.Remove(n - i);
  }
}

#define FUNC_ENTRY(func)                                                       \
  { #func, func }

std::map<std::string, void (*)(int)> chen3fengFunctions = {
    FUNC_ENTRY(chen3fengInserts),   FUNC_ENTRY(chen3fengWorstInserts),
    FUNC_ENTRY(chen3fengAvgSearch), FUNC_ENTRY(chen3fengSearchEnd),
    FUNC_ENTRY(chen3fengDelete),    FUNC_ENTRY(chen3fengWorstDelete)};

void iterations(int n) { printf("%d", n); }

void runIterations(const std::string &name, int start, int end, int step,
                   void (*f)(int)) {
  printf("%s", name.c_str());
  for (int i = start; i <= end; i += step) {
    printf(",");
    f(i);
  }
  printf("\n");
}

int main() {
  runIterations("iterations", start, end, step, iterations);
  for (auto &&[fname, f] : chen3fengFunctions) {
    runIterations(fname, start, end, step, f);
  }
}
