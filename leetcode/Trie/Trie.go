package main

/**
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。


示例：
输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]
解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True

提示：
1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
*/
type Trie struct {
	// 标识该节点是否是字符串的结束节点
	isEnd bool
	// 当前节点的孩子节点
	next [26]*Trie
}

// 设置当前节点为一个字符串的结束节点
func (this *Trie) setIsEnd(isEnd bool) {
	this.isEnd = isEnd
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	b := []byte(word)
	node := this
	for _, v := range b {
		u := v - 'a'
		if node.next[u] == nil {
			node.next[u] = &Trie{}
		}
		node = node.next[u]
	}
	node.setIsEnd(true)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	b := []byte(word)
	node := this
	for _, v := range b {
		u := v - 'a'
		if node.next[u] == nil {
			return false
		}
		node = node.next[u]
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	b := []byte(prefix)
	node := this
	for _, v := range b {
		u := v - 'a'
		if node.next[u] == nil {
			return false
		}
		node = node.next[u]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/**
class Trie {
    // 构造Trie树节点
    class TrieNode {

        // 标识该节点是否是字符串的结束节点
        boolean isEnd = false;
        // 当前节点的孩子节点
        TrieNode[] next = new TrieNode[26];

        // 设置当前节点为一个字符串的结束节点
        public void setIsEnd(boolean isEnd) {
            this.isEnd = isEnd;
        }

    }

    // root为根节点
    TrieNode root;

    public Trie() {
        root = new TrieNode();
    }

    public void insert(String word) {

        char[] chs = word.toCharArray();
        // 表示从根节点开始向下构建
        TrieNode node = root;
        for (int i = 0; i < chs.length; i++) {

            int u = chs[i] - 'a';
            // 判断node的子节点集合中是否已经存在了chs[i], 不存在则创建
            if (node.next[u] == null)
                node.next[u] = new TrieNode();
            // 继续向下构建
            node = node.next[u];

        }

        // 当前TrieNode节点是一个字符串的结尾
        node.setIsEnd(true);

    }

    public boolean search(String word) {

        char[] chs = word.toCharArray();
        // 表示从根节点开始向下构建
        TrieNode node = root;
        for (int i = 0; i < chs.length; i++) {

            int u = chs[i] - 'a';
            // 判断node的子节点集合中是否已经存在了chs[i], 不存在则创建
            if (node.next[u] == null)
                return false;
            // 继续向下构建
            node = node.next[u];

        }

        // 当前TrieNode节点是否一个字符串的结尾
        return node.isEnd;

    }

    public boolean startsWith(String prefix) {
        char[] chs = prefix.toCharArray();
        // 表示从根节点开始向下构建
        TrieNode node = root;
        for (int i = 0; i < chs.length; i++) {

            int u = chs[i] - 'a';
            // 判断node的子节点集合中是否已经存在了chs[i], 不存在则创建
            if (node.next[u] == null)
                return false;
            // 继续向下构建
            node = node.next[u];

        }

        // 前缀查找成功
        return true;

    }
}
*/
func main() {

}
