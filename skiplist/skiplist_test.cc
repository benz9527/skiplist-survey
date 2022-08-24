#include "skiplist/skiplist.h"
#undef NDEBUG
#include <assert.h>

#define ASSERT_EQ(a, b) assert((a) == (b))
#define ASSERT_TRUE(a) assert(a)
#define ASSERT_FALSE(a) assert(!(a))

void TestSkipList_Insert() {
  SkipList<int, int> sl;
  for (size_t i = 0; i < 10000; ++i) {
    ASSERT_EQ(sl.Size(), i);
    ASSERT_FALSE(sl.Has(i));
    sl.Insert(i, i);
    ASSERT_EQ(sl.Size(), i + 1);
    ASSERT_TRUE(sl.Has(i));
  }
}

void TestSkipList_Insert_Dup() {
  SkipList<int, int> sl;
  sl.Insert(1, 1);
  ASSERT_EQ(sl.Size(), 1U);
  sl.Insert(1, 2);
  ASSERT_EQ(sl.Size(), 1U);
}

SkipList<int, int> MakeSkipListN(int n) {
  SkipList<int, int> sl;
  for (int i = 0; i < n; ++i) {
    sl.Insert(i, i);
  }
  return sl;
}

void TestSkipList_Remove() {
  auto sl = MakeSkipListN(100);
  for (int i = 0; i < 100; i++) {
    ASSERT_TRUE(sl.Remove(i));
  }
  ASSERT_TRUE(sl.IsEmpty());
  ASSERT_EQ(sl.Size(), 0U);
}

void TestSkipList_Remove_Nonexist() {
  SkipList<int, int> sl;
  sl.Insert(1, 1);
  sl.Insert(2, 2);
  ASSERT_FALSE(sl.Remove(0));
  ASSERT_FALSE(sl.Remove(3));
  ASSERT_EQ(sl.Size(), 2U);
}

int main() {
  TestSkipList_Insert();
  TestSkipList_Insert_Dup();
  TestSkipList_Remove();
}
