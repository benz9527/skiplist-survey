// This implementation is based on
// https://github.com/liyue201/gostl/tree/master/ds/skiplist (many thanks),
// added many optimizations, such as:
//
//  - adaptive level
//  - lesser search for prevs when key already exists.
//  - reduce memory allocations
//  - richer interface.
//
// etc.

#pragma once

#include <stddef.h>
#include <time.h>

#include <algorithm>
#include <random>
#include <tuple>
#include <vector>

const int skipListMaxLevel = 40;

int BitLength64(uint64_t n) {
  if (n == 0) {
    return 1;
  }
  return 64 - __builtin_clzl(n);
}

class Rand64 {
public:
  explicit Rand64(uint64_t seed = 1) : i(seed) {}
  uint64_t operator()() {
    return (i = (164603309694725029ull * i) % 14738995463583502973ull);
  }
  uint64_t i = 1;
};

// SkipList is a probabilistic data structure that seem likely to supplant
// balanced trees as the implementation method of choice for many applications.
// Skip list algorithms have the same asymptotic expected time bounds as
// balanced trees and are simpler, faster and use less space.
//
// See https://en.wikipedia.org/wiki/Skip_list for more details.
template <typename K, typename V> class SkipList {
  struct Node;

public:
  // IsEmpty implements the Container interface.
  bool IsEmpty() const { return this->len == 0; }
  size_t Size() const { return this->len; }
  void Clear() {
    this->level = 1;
    this->len = 0;
  }

  bool Has(const K &key) const {
    return const_cast<SkipList *>(this)->FindNode(key) != nullptr;
  }

  const V *Find(const K &key) const {
    auto node = this->FindNode(key);
    if (node != nullptr) {
      return &node->value;
    }
    return nullptr;
  }

  void Insert(const K &key, const V &value) {
    auto &&[node, prevs] = FindInsertPoint(key);

    if (node != nullptr) {
      // Already exist, update the value
      node->value = value;
      return;
    }

    auto level = this->RandomLevel();
    node = NewNode(level, key, value);

    for (auto i = 0; i < std::min(level, this->level); ++i) {
      node->next[i] = prevs[i]->next[i];
      prevs[i]->next[i] = node;
    }

    if (level > this->level) {
      for (auto i = this->level; i < level; ++i) {
        this->head.next[i] = node;
      }
      this->level = level;
    }

    ++this->len;
  }

  bool Remove(const K &key) {
    auto &&[node, prevs] = this->FindRemovePoint(key);
    if (node == nullptr) {
      return false;
    }
    for (auto i = 0; i < node->level; ++i) {
      prevs[i]->next[i] = node->next[i];
    }
    DeleteNode(node);
    while (this->level > 1 && this->head.next[this->level - 1] == nullptr) {
      this->level--;
    }
    this->len--;
    return true;
  }

private:
  Node *NewNode(int level, const K &key, const V &value) {
    auto p = calloc(sizeof(Node) + level * sizeof(Node *), 1);
    return new (p) Node{key, value, level};
  }

  void DeleteNode(Node *node) {
    node->~Node();
    free(node);
  }

  int RandomLevel() {
    auto total = (uint64_t(1) << uint64_t(skipListMaxLevel)) - 1;
    auto k = this->rander() % total;
    auto level = skipListMaxLevel - BitLength64(k) + 1;
    while (level > 3 && (1 << (level - 3)) > this->len) {
      level--;
    }
    return level;
  }

  const Node *FindNode(const K &key) const {
    // This function execute the job of findNode if eq is true, otherwise
    // lowBound. Passing the control variable eq is ugly but it's faster than
    // testing node again outside the function in findNode.
    auto prev = static_cast<const Node *>(&this->head);
    for (auto i = this->level - 1; i >= 0; i--) {
      for (auto cur = prev->next[i]; cur != nullptr; cur = cur->next[i]) {
        if (cur->key == key) {
          return cur;
        }
        if (cur->key > key) {
          // All other node in this level must be greater than the key,
          // search the next level.
          break;
        }
        prev = cur;
      }
    }
    return prev->next[0];
  }

  // findInsertPoint returns (*node, nullptr) to the existed node if the key
  // exists, or (nullptr, []*node) to the previous nodes if the key doesn't
  // exist
  std::tuple<Node *, std::vector<Node *> &> FindInsertPoint(const K &key) {
    auto &prevs = this->prevsCache;
    prevs.resize(this->level);
    auto prev = static_cast<Node *>(&this->head);
    for (auto i = this->level - 1; i >= 0; i--) {
      for (auto next = prev->next[i]; next != nullptr; next = next->next[i]) {
        if (next->key == key) {
          // The key is already existed, prevs are useless because no new node
          // insertion. stop searching.
          return {next, prevs};
        }
        if (next->key > key) {
          // All other node in this level must be greater than the key,
          // search the next level.
          break;
        }
        prev = next;
      }
      prevs[i] = prev;
    }
    return {nullptr, prevs};
  }

  // findRemovePoint finds the node which match the key and it's previous nodes.
  std::tuple<Node *, std::vector<Node *> &> FindRemovePoint(const K &key) {
    auto &prevs = this->FindPrevNodes(key);
    auto node = prevs[0]->next[0];
    if (node == nullptr || node->key != key) {
      return {nullptr, prevs};
    }
    return {node, prevs};
  }

  std::vector<Node *> &FindPrevNodes(const K &key) {
    auto &prevs = this->prevsCache;
    prevs.resize(this->level);
    auto prev = static_cast<Node *>(&this->head);
    for (auto i = this->level - 1; i >= 0; i--) {
      for (auto next = prev->next[i]; next != nullptr; next = next->next[i]) {
        if (next->key >= key) {
          break;
        }
        prev = next;
      }
      prevs[i] = prev;
    }
    return prevs;
  }

private:
  struct Node {
    K key{};
    V value{};
    int level = 0;
    Node *next[0];
  };
  struct HeadNode : Node {
    Node *next[skipListMaxLevel] = {};
  };
  int level = 1;
  size_t len = 0;
  HeadNode head;
  std::vector<Node *> prevsCache;
  // std::mt19937_64 rander{};
  Rand64 rander;
};
