package geo_range

type GeoTrieRoot struct {
	root *geoTrieNode
}

func (t *GeoTrieRoot) Insert(key string) bool {
	return t.root.insert(key, 0)
}

func (t *GeoTrieRoot) AllLeaf() []string {
	return t.root.allLeaf()
}

func (t *GeoTrieRoot) Has(key string) bool {
	return t.root.has(key, 0)
}

func NewGeoTrie() *GeoTrieRoot {
	return &GeoTrieRoot{
		root: &geoTrieNode{
			children: make(map[byte]*geoTrieNode),
		},
	}
}

type geoTrieNode struct {
	children map[byte]*geoTrieNode
}

func (t *geoTrieNode) insert(key string, idx int) bool {
	if idx == len(key) {
		return true
	}

	ch := key[idx]
	if _, ok := t.children[ch]; !ok {
		t.children[ch] = &geoTrieNode{children: make(map[byte]*geoTrieNode)}
	}
	return t.children[ch].insert(key, idx+1)
}

func (t *geoTrieNode) allLeaf() []string {
	out := make([]string, 0)
	for ch, child := range t.children {
		paths := child.allLeaf()
		if len(paths) == 0 {
			out = append(out, string(ch))
		} else {
			for _, path := range paths {
				out = append(out, string(ch)+path)
			}
		}
	}
	return out
}

func (t *geoTrieNode) has(key string, idx int) bool {
	if idx == len(key) {
		return true
	}

	ch := key[idx]
	if _, ok := t.children[ch]; !ok {
		return false
	}
	return t.children[ch].has(key, idx+1)
}
