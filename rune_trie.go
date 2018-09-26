package main

import (
   "fmt"
   )
type RuneTrie struct {
   Nodes map[rune]*RuneTrie
}

func (trie *RuneTrie) InsertRune(r rune) (*RuneTrie) {
   if child, exists := trie.Nodes[r]; exists {
      // fmt.Println("Returning existing child")
      return child
   }
   // fmt.Println("Creating new child", r)
   newChild := RuneTrie{make(map[rune]*RuneTrie)}
   trie.Nodes[r] = &newChild
   // fmt.Println("Number of nodes in trie:", len(trie.Nodes))
   return &newChild
}

func (trie *RuneTrie) InsertString(s string) {
   n := trie // insertion node
   for _, char := range s {
      // fmt.Printf("n has %d nodes. Inserting %c\n", len(n.Nodes), char)
      n = n.InsertRune(char)
   }
   n.InsertRune('\000')
}


//todo printing is nice, but it would be better to have an enclosure which stores all the strings..
func (trie *RuneTrie) Print(prefix string) {
   //dfs on trie

   for k, v := range trie.Nodes {
      if ( k == '\000' ) {
         fmt.Println(prefix)
      } else {
         v.Print(prefix + string(k))
      }
   }
}

func main() {
   fmt.Println("Hello World!")
   trie := RuneTrie{ make(map[rune]*RuneTrie) }
   //trie.Nodes['h'] = RuneTrie{make(map[rune]RuneTrie)}
   trie.InsertString("hello world")
   trie.InsertString("hello world!")
   trie.InsertString("hen")
   trie.InsertString("henry")
   fmt.Println("Printing trie!")
   trie.Print("")

}