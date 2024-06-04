package alud

import (
	"fmt"

	"github.com/jbowtie/gokogiri/xml"
)

/*
func auxiliary(nodes []*nodeType, q *context) string {
	if len(nodes) != 1 { // TODO: in script staat: = 0
		return "ERROR_AUXILIARY_FUNCTION_TAKES_EXACTLY_ONE_ARG"
	}
	return auxiliary1(nodes[0], q)
}
*/

func auxiliary1(node xml.Node, q *context) (aux string, err error) {
	if node.Attr("pt") != "ww" {
		return "", fmt.Errorf("ERROR_NO_VERB")
	}
	if node.Attr("rel") != "hd" {
		return "verb", nil
	}

	if test(
		q,
		`$node[@lemma="zijn" and
                              ../node[@rel="predc"] and
                              not(../node[@rel="obj1" or @rel="se" or @rel="vc"])
                                 ]`,
		"node", node) {
		return "cop", nil
	}

	if test(
		q,
		`$node[(@lemma="zijn" or @lemma="worden") and
                                 ( ../node[@rel="vc"] and
                                        not(../node[@rel="svp"]) and
                                ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj1"]/@index or
                                  ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj1"]/@index or
                                  ../node[@rel="vc" and not(@pt or @cat)]/@index =
                                      $idxnodes[@rel="vc" and node[@rel="obj1"]/@index = $node/../node[@rel="su"]/@index]/@index
                                 or not(../node[@rel="su"])
                                 )
                             )
                           ]`,
		"node", node,
		"idxnodes", q.varidxnodes) { // removed reference to @sc=passive as this is less reliable in automatic parses GB 18/03/21
		return "aux:pass", nil
	}

	// krijgen passive with iobj control
	if test(q, `$node[@lemma="krijgen" and
	  	              ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj2"]/@index or
	                    ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj2"]/@index
	                  )]`, "node", node) {
		return "aux:pass", nil
	}

	// alpino has no principled distinction between AUX and VERB, should be TAME verbs semantically, we follow ENGLISH
	// blijken and hoeven removed from list
	// hij heeft als opdracht stammen uit elkaar te houden  , removed starts-with(sc,aux) as less reliable in automatic parses GB 18/03/21
	// dangling aux in gapped coordination
	// ze hebben thuis nog een varken zitten -- hebben as aci verb takes a obj1 and vc, but is not aux in this construction GB 10/01/23
	// van plan zijn -- exclude svp sister as well (otherwise, svp ends up as compound:prt of the embedded verb ) GB 20/11/23
	if test(q, `$node[(@lemma="hebben" or @lemma="kunnen" or @lemma="moeten" or @lemma="mogen" or @lemma="zijn" or @lemmas="zullen")  and
                      ( ../node[@rel="vc"  and
	                              ( (@cat="ppart" or @cat="inf" or @cat="ti") or
	                                ( @cat="conj" and node[@rel="cnj" and (@cat="ppart" or @cat="inf" or @cat="ti")] ) or
	                                ( @index and not(@pt or @cat))
	                              )
	                            ]
	                   ) and
	                   not(../node[@rel="predc" or @rel="obj1" or @rel="svp")])
	               ]`, "node", node) {
		return "aux", nil
	}

	return "verb", nil
}
