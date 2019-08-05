// +build ignore

package alud

import (
	"fmt"
)

// voorkwam dat LPF opnieuw of SGP voor het eerst in de regering zou komen  -- gapped LD
func fixMisplacedHeadsInCoordination(q *context) {

	if len(q.varindexnodes) == 0 {
		return
	}

	seen := make(map[[2]int]int)
	counter := 0

START:
	for true {
		for _, n1 := range q.varallnodes {
			// FIND op varallnodes niet mogelijk omdat twee keer naar $node wordt verwezen, en dat moet dezelfde node zijn
			for _, n2 := range FIND(q, `
$n1[@rel=("hd","ld") and
      @index and
      (@pt or @cat) and
      ancestor::node[@rel="cnj"] and
      ancestor::node[@cat="conj"]/node[@rel="cnj" and
                                       descendant-or-self::node[@rel=("hd","ld") and
                                                                @index=$n1/@index and
                                                                not(@cat or @pt) and
                                                                ( @begin        = ..//node[@cat or @pt]/@end or
                                                                  @begin - 1000 = ..//node[@cat or @pt]/@end
                                                                )
                                                               ]
                                       ]]`) {
				node2 := n2.(*nodeType)
				for _, n3 := range FIND(q, `
$q.varallnodes[@rel=("hd","ld","vc") and @index and not(@pt or @cat) and
                 ancestor::node[@rel="cnj"]  and
                                    ( @begin        = ..//node[@cat or @pt]/@end or
                                      @begin - 1000 = ..//node[@cat or @pt]/@end
                                     )]`) {
					node3 := n3.(*nodeType)
					if node2.Index == node3.Index {
						pair := [2]int{node2.Id, node3.Id}
						if i, ok := seen[pair]; ok {
							if i == 1 {
								panic(fmt.Sprintf("Loop detected in fixMisplacedHeadsInCoordination: %d %d", node2.Id, node3.Id))
							}
							seen[pair]++
							continue
						}
						seen[pair] = 1
						counter++
						// kopieer inhoud van node2 (niet leeg) naar node3 (leeg)
						id, rel := node3.Id, node3.Rel
						*node3 = *node2
						node3.Id, node3.Rel = id, rel
						// maak node2 leeg
						*node2 = nodeType{
							Begin: node2.Begin,
							End:   node2.End,
							Id:    node2.Id,
							Index: node2.Index,
							Rel:   node2.Rel,
							Node:  []*nodeType{},
						}
						// opnieuw beginnen
						inspect(q)
						continue START
					}
				}
			}
		}
		break
	}
	if counter > 0 {
		q.debugs = append(q.debugs, fmt.Sprintf("fixMisplacedHeadsInCoordination: %d swaps", counter))
	}
}
