// +build ignore

package alud

import (
	"fmt"
)

/*
func auxiliary(nodes []*nodeType, q *context) string {
	if len(nodes) != 1 { // TODO: in script staat: = 0
		return "ERROR_AUXILIARY_FUNCTION_TAKES_EXACTLY_ONE_ARG"
	}
	return auxiliary1(nodes[0], q)
}
*/

func auxiliary1(node *nodeType, q *context) (aux string, err error) {

	if node.Pt != "ww" {
		return "", fmt.Errorf("ERROR_NO_VERB")
	}
	if node.Rel != "hd" {
		return "verb", nil
	}

	if TEST(q, `$node[not(../node[@rel=("obj1","se","vc")]) and
			        (: ud documentation suggests 1 cop per lg, van Eynde suggests much more, compromise: the traditional ones :)
			        (: @lemma=("zijn","lijken","blijken","blijven","schijnen","heten","voorkomen","worden","dunken") and :)
			        @lemma="zijn" and
		                 ( contains(@sc,'copula') or
		                   contains(@sc,'pred')   or
		                   contains(@sc,'cleft')  or
		                   ../node[@rel="predc"]
		                 ) ]`) {
		return "cop", nil
	}

	if TEST(q, `$node[@lemma=("zijn","worden") and
	                  ( @sc="passive"  or
	                  	 ( ../node[@rel="vc"] and
	                        ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj1"]/@index or
	                          ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj1"]/@index or
	                          ../node[@rel="vc" and not(@pt or @cat)]/@index =
	                              $q.varindexnodes[@rel="vc" and node[@rel="obj1"]/@index = $node/../node[@rel="su"]/@index]/@index
	                         or not(../node[@rel="su"])
	                         )
	                     )
	                  ) ]`) {
		return "aux:pass", nil
	}

	// krijgen passive with iobj control
	if TEST(q, `$node[@lemma="krijgen" and
	  	              ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj2"]/@index or
	                    ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj2"]/@index
	                  )]`) {
		return "aux:pass", nil
	}

	// alpino has no principled distinction between AUX and VERB, should be TAME verbs semantically, we follow ENGLISH
	// blijken and hoeven removed from list
	if TEST(q, `$node[not(../node[@rel="predc"]) and  (: hij heeft als opdracht stammen uit elkaar te houden  :)
	                 ( starts-with(@sc,'aux') or
	                   ( ../node[@rel="vc"  and
	                              ( @cat=("ppart","inf","ti") or
	                                ( @cat="conj" and node[@rel="cnj" and @cat=("ppart","inf","ti")] ) or
	                                ( @index and not(@pt or @cat))  (: dangling aux in gapped coordination :)
	                              )
	                            ]   and
	                     @lemma=("hebben","kunnen","moeten","mogen","zijn","zullen")
	                   )
	                 )
	               ]`) {
		return "aux", nil
	}

	return "verb", nil
}
