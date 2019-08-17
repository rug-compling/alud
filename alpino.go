package alud

func alpinoRestore(q *context) {
	for i := len(q.swapped) - 1; i >= 0; i-- {
		swap(q.swapped[i][1], q.swapped[i][0])
	}
	for _, n := range q.varindexnodes {
		node := n.(*nodeType)
		if node.udOldState != nil {
			node.Begin = node.udOldState.Begin
			node.End = node.udOldState.End
			node.Word = node.udOldState.Word
			node.Lemma = node.udOldState.Lemma
			node.Postag = node.udOldState.Postag
			node.Pt = node.udOldState.Pt
		}
	}
}

func alpinoFormat(q *context) (alpino string, err error) {
	return
}

func alpinoDo(conllu string, q *context) {

}
