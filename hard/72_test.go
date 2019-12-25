package hard

import "testing"

func TestMinDistance(t *testing.T) {
	word1, word2 := "dinitrophenylhydrazine", "acetylphenylhydrazine"
	t.Log(MinDistance(word1, word2))
	word1, word2 = "", ""
	t.Log(MinDistance(word1, word2))
}
