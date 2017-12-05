/***************************************************************************
 *
 * @desc
 * @author <wanyang@wanyang.site>
 * @version 05/12/2017 12:42
 * @project structs
 **************************************************************************/
package trie_tree

type trieNode struct {
	exist    bool
	children [26]*trieNode
	cnt      int
	value    string
}

var root trieNode

func Search(r *trieNode, key string) int {
	for i, v := range key {
		if r.children[v-'a'] == nil {
			return 0
		}
		if i == len(key)-1 {
			if r.children[v-'a'].exist {
				return r.children[v-'a'].cnt
			}
		} else {
			r = r.children[v-'a']
		}
	}
	return 0
}

func AddNode(r *trieNode, key string, start int) {
	l := len(key)
	if start == l {
		return
	}
	t := key[start] - 'a'
	if r.children[t] != nil {
		if start == l-1 {
			if r.children[t].exist == false {
				r.children[t].exist = true
				r.children[t].value = key[:start+1]
			}
			r.children[t].cnt += 1
		} else {
			AddNode(r.children[t], key, start+1)
		}
	} else {
		r.children[t] = &trieNode{}
		if start == l-1 {
			r.children[t].exist = true
			r.children[t].value = key[:start+1]
			r.children[t].cnt = 1
		}
		AddNode(r.children[t], key, start+1)
	}
}
