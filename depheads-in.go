// +build ignore

package alud

import (
	"fmt"
	"sort"
)

// recursive
func externalHeadPosition(nodes []interface{}, q *context, tr []trace) int {
	depthCheck(q, "externalHeadPosition")

	if len(nodes) == 0 {
		tr = append(tr, trace{s: "externalHeadPosition"})
		panic(tracer(fmt.Sprint("External head must have one arg, has ", len(nodes)), tr, q))
	}

	node := nodes[0].(*nodeType)
	tr = append(tr, trace{s: "externalHeadPosition", node: node})

	if node.Rel == "hd" && (node.udPos == "ADP" || node.parent.Cat == "pp") {
		// vol vertrouwen
		if n := FIND(q, `$node/../node[@rel="predc"]`); len(n) > 0 {
			// met als titel
			return internalHeadPosition(n[:1], q, tr)
		}
		if obj1_vc_se_me := FIND(q, `$node/../node[@rel=("obj1","vc","se","me")]`); len(obj1_vc_se_me) > 0 {
			// adding pt/cat enough for gapping cases?
			if TEST(q, `$obj1_vc_se_me[@pt or @cat]`) {
				return internalHeadPositionWithGapping(if1(obj1_vc_se_me), q, tr)
			}
			if TEST(q, `$obj1_vc_se_me[@index = ancestor::node/node[@rel=("rhd","whd")]/@index]`) {
				return internalHeadPosition(FIND(q, `$node/ancestor::node/node[@rel=("rhd","whd")
                                       and @index = $node/../node[@rel=("obj1","vc","se","me")]/@index]`), q, tr)
			}
			if pobj1 := FIND(q, `$node/../node[@rel="pobj1"]`); len(pobj1) > 0 {
				return internalHeadPosition(if1(pobj1), q, tr)
			}
			// in de eerste rond --> typo in LassySmall/Wiki , binnen en [advp later buiten ]
			return externalHeadPosition(node.axParent, q, tr)
		} else {
			return externalHeadPosition(node.axParent, q, tr)
		}
	}

	aux := auxiliary1(node, q)

	if node.Rel == "hd" && (aux == "aux" || aux == "aux:pass") {
		// aux aux:pass cop
		if vc_predc := FIND(q, `$node/../node[@rel=("vc","predc")]`); len(vc_predc) > 0 {
			if TEST(q, `$vc_predc[@pt or (@cat and node[@pt or @cat])]`) {
				// skip vc with just empty nodes
				return internalHeadPositionWithGapping(if1(vc_predc), q, tr)
			}
		}
		// if ($node/../node[@rel="predc"]/@index = $node/../../node[@rel="whd"]/@index)
		//     then local:internal_head_position($node/../../node[@rel="whd"])
		return externalHeadPosition(node.axParent, q, tr) // gapping, but does it ever occur with aux?? with cop: hij was en blijft nog steeds een omstreden figuur
	}

	if node.Rel == "hd" && aux == "cop" {
		predc := FIND(q, `$node/../node[@rel="predc"]`)
		if len(predc) > 0 && TEST(q, `$predc[@pt or @cat]`) {
			return internalHeadPositionWithGapping(if1(predc), q, tr)
		}
		if TEST(q, `$node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index`) {
			return internalHeadPosition(
				FIND(q, `$node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index]`),
				q,
				tr)
		}
		return externalHeadPosition(node.axParent, q, tr) // gapping, but could it??
	}

	if node.Rel == "hd" || node.Rel == "nucl" || node.Rel == "body" {
		if n := FIND(q, `$node/../node[@rel="hd" and @begin < $node/@begin]`); len(n) > 0 {
			return internalHeadPosition(list(n), q, tr) // dan moet je moet
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if node.Rel == "predc" {
		if TEST(q, `$node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]]`) {
			if TEST(q, `$node/../node[@rel="hd" and @ud:pos="ADP"]`) {
				return externalHeadPosition(node.axParent, q, tr) // met als presentator Bruno W , met als gevolg [vc dat ...]
			}
			return internalHeadPosition(FIND(q, `$node/../node[@rel="hd"]`), q, tr)
		}
		if TEST(q, `$node/parent::node[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ]`) {
			//reduced relatives , make sure head is not empty (ellipsis)
			// also : ap with predc: actief als schrijver
			return internalHeadPosition(FIND(q, `$node/../node[@rel="hd"]`), q, tr)
		}
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))]`) { // [met als titel] -- obj1/vc missing
			return internalHeadPosition(FIND(q, `$node/../node[@rel="hd"]`), q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr) // covers gapping as well?
	}

	if TEST(q, `$node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])]`) {
		if predc := FIND(q, `$node/../node[@rel="predc"]`); len(predc) > 0 {
			return internalHeadPosition(predc, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if TEST(q, `$node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])]`) {
		if vc := FIND(q, `$node/../node[@rel="vc"]`); len(vc) > 0 {
			return internalHeadPosition(vc, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if TEST(q, `$node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])]`) { // mede op grond hiervan
		// daarom dus
		if TEST(q, `$node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)]`) {
			return internalHeadPositionWithGapping(node.axParent, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr) // gapping
	}

	if TEST(q, `$node[@rel=("cnj","dp","mwp")]`) {
		if node == nLeft(FIND(q, `$node/../node[@rel=("cnj","dp","mwp")]`)) {
			return externalHeadPosition(node.axParent, q, tr)
		}
		if node.Rel == "cnj" {
			return headPositionOfConjunction(node, q, tr)
		}
		return internalHeadPositionWithGapping(node.axParent, q, tr)
	}

	if TEST(q, `$node[@rel="cmp" and ../node[@rel="body"]]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/../node[@rel="body"][1]`), q, tr)
	}

	if node.Rel == "--" && node.Cat != "" {
		if node.Cat == "mwu" {
			if TEST(q, `$node/../node[@cat and not(@cat="mwu")]`) { // fix for multiword punctuation in Alpino output
				return internalHeadPosition(FIND(q, `$node/../node[@cat and not(@cat="mwu")][1]`), q, tr)
			}
			return externalHeadPosition(node.axParent, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if node.Rel == "--" && node.udPos != "" {
		if TEST(q, `$node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
	                   and ../node[@rel="--" and
	                               not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
	                  ]`) {
			return internalHeadPositionWithGapping(
				FIND(q, `$node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))][1]`),
				q,
				tr)
		}
		if n := FIND(q, `$node/../node[@cat][1]`); len(n) > 0 {
			return internalHeadPosition(n, q, tr)
		}
		if TEST(q, `$node[@ud:pos="PUNCT" and count(../node) > 1]`) {
			if n := FIND(q, `$node/../node[not(@ud:pos="PUNCT")][1]`); len(n) > 0 {
				return internalHeadPosition(n, q, tr)
			}
			if node == nLeft(FIND(q, `$node/../node[@rel="--" and (@cat or @pt)]`)) {
				return externalHeadPosition(node.axParent, q, tr)
			}
			return 1000 // ie end of first punct token
		}
		if node.parent.Begin >= 0 {
			return externalHeadPosition(node.axParent, q, tr)
		}
		panic(tracer("No head found", tr, q))
	}

	if node.Rel == "dlink" || node.Rel == "sat" || node.Rel == "tag" {
		if n := FIND(q, `$node/../node[@rel="nucl"]`); len(n) > 0 {
			return internalHeadPositionWithGapping(n, q, tr)
		}
		panic(tracer("No external head", tr, q))
	}

	if node.Rel == "vc" {
		if TEST(q, `$node/../node[@rel="hd" and
	                            ( @ud:pos="AUX" or
	                              $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                            )
	                        ]
	                   and not($node/../node[@rel="predc"])`) {
			return externalHeadPosition(node.axParent, q, tr)
		}
		if TEST(q, `$node/../@cat="pp"`) { // eraan dat
			return externalHeadPosition(node.axParent, q, tr)
		}
		if TEST(q, `$node/../node[@rel=("hd","su") and (@pt or @cat)]`) {
			return internalHeadPositionWithGapping(node.axParent, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if node.Rel == "whd" || node.Rel == "rhd" {
		if node.Index > 0 {
			return externalHeadPosition(FIND(q, `($node/../node[@rel="body"]//node[@index = $node/@index ])[1]`), q, tr)
		}
		return internalHeadPosition(FIND(q, `$node/../node[@rel="body"]`), q, tr)
	}

	/*
		we need to select the original node and not the result of
		following-cnj-sister, as that has no global context
		and global context is needed where the hd is an index node...
		unfortunately, nodes are completely identical in some
		elliptical cases, select last() as brute force solution
	*/
	if node.Rel == "crd" {
		tmp := followingCnjSister(node, q, tr)
		return internalHeadPositionWithGapping(ifZ(FIND(q, `$node/../node[@rel="cnj" and
	          	                 @begin=$tmp/@begin and
	          	                 @end=$tmp/@end
	          	                ]`)), q, tr)
	}

	if node.Rel == "su" {
		if TEST(q, `$node[../node[@rel="vc"] and ../node[@rel="hd" and 
			                  ( @ud:pos="AUX" or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index ) ]]`) {
			tmp := internalHeadPositionWithGapping(FIND(q, `$node/../node[@rel="vc"]`), q, tr)
			if node.Begin < tmp && tmp <= node.End {
				return externalHeadPosition(node.axParent, q, tr)
			}
			return tmp
		}
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) { // gapping
			return internalHeadPositionWithGapping(node.axParent, q, tr) // ud head could still be a predc
		}
		// only for 1 case where verb is missing -- die eigendom ... (no verb))
		if TEST(q, `$node[../node[@rel="predc"] and not(../node[@rel="hd"])]`) {
			return internalHeadPosition(FIND(q, `$node/../node[@rel="predc"]`), q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr) // this probably does no change anything, as we are still attaching to head of left conjunct
	}

	if node.Rel == "obj1" {
		if TEST(q, `$node/../node[@rel=("hd","su") and (@pt or @cat)]`) { // gapping, as su but now su could be head as well
			return internalHeadPositionWithGapping(node.axParent, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if node.Rel == "pc" {
		if TEST(q, `$node/../node[@rel=("hd","su","obj1") and (@pt or @cat)]`) { // gapping, as su but now su could be head as well
			return internalHeadPositionWithGapping(node.axParent, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if node.Rel == "mod" {
		if TEST(q, `$node/../node[@rel=("hd","su","obj1","pc","predc","body") and (@pt or @cat)]`) { // gapping, as su but now su or obj1  could be head as well
			return internalHeadPositionWithGapping(node.axParent, q, tr)
		}
		if n := FIND(q, `$node/../node[@rel="mod" and (@cat or @pt)]`); len(n) > 0 {
			if node == nLeft(n) { // gapping with multiple mods
				return externalHeadPosition(node.axParent, q, tr)
			}
			return internalHeadPositionWithGapping(node.axParent, q, tr)
		}
		if TEST(q, `$node/../../node[@rel="su" and (@pt or @cat)]`) { // an mod in an otherwise empty tree (after fixing heads in conj)
			return internalHeadPosition(FIND(q, `$node/../../node[@rel="su"]`), q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr) /* an empty mod in an otherwise empty tree
		   -- mod is co-indexed with rhd, rest is elided,
		   LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml */
	}

	if node.Rel == "app" || node.Rel == "det" {
		if TEST(q, `$node/../node[@rel=("hd","mod") and (@pt or @cat)]`) { // gapping with an app (or a det)!
			return internalHeadPositionWithGapping(node.axParent, q, tr)
		}
		return externalHeadPosition(node.axParent, q, tr)
	}

	if node.Rel == "top" {
		return 0
	}

	if node.Rel != "hd" { // TODO: klopt dit?
		return internalHeadPositionWithGapping(node.axParent, q, tr)
	}

	panic(tracer("No external head", tr, q))
}

// recursive
func internalHeadPosition(nodes []interface{}, q *context, tr []trace) int {
	depthCheck(q, "internalHeadPosition")

	if n := len(nodes); n != 1 {
		tr = append(tr, trace{s: "internalHeadPosition"})
		if n == 0 {
			panic(tracer("No internal head position found", tr, q))
		} else if n > 1 {
			panic(tracer("More than one internal head position found", tr, q))
		}
	}
	node := nodes[0]

	if TEST(q, `$node[@cat="pp"]`) {
		// if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		// then $node/node[@rel="hd"]/@end
		if n := FIND(q, `$node/node[@rel=("obj1","pobj1","se")][1]`); len(n) > 0 {
			return internalHeadPosition(n, q, tr)
		}
		if n := FIND(q, `$node/node[@rel="hd"]`); len(n) > 0 {
			// if ($node/@cat="mwu")  ( mede [op grond hiervan] )
			//     local:internal_head_position($node/node[@rel="hd"] )
			return internalHeadPosition(n, q, tr)
		}
		return internalHeadPosition(FIND(q, `$node/node[1]`), q, tr)
	}

	if TEST(q, `$node[@cat="mwu"]`) {
		return FIND(q, `$node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end`)[0].(int)
	}

	if TEST(q, `$node[@cat="conj"]`) {
		return internalHeadPosition(ifLeft(FIND(q, `$node/node[@rel="cnj"]`)), q, tr)
	}

	if predc := FIND(q, `$node/node[@rel="predc"]`); len(predc) > 0 {
		if TEST(q, `$node/node[@rel="hd" and @ud:pos="AUX"]`) {
			return internalHeadPosition(predc, q, tr)
		}
		hd := FIND(q, `$node/node[@rel="hd"]`)
		if len(hd) == 0 { // cases where copula is missing by accident (ungrammatical, not gapping)
			return internalHeadPosition(predc, q, tr)
		}
		return internalHeadPosition(hd, q, tr)
	}

	if TEST(q, `$node[node[@rel="vc"] and
	                  node[@rel="hd" and
	                       ( @ud:pos="AUX" or
	                         $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                        )
	                      ]
	                 ]`) {
		return internalHeadPosition(FIND(q, `$node/node[@rel="vc"]`), q, tr)
	}

	if n := FIND(q, `$node/node[@rel="hd"][1]`); len(n) > 0 {
		return internalHeadPosition(n, q, tr)
	}

	if n := FIND(q, `$node/node[@rel="body"][1]`); len(n) > 0 {
		return internalHeadPosition(n, q, tr)
	}

	if n := FIND(q, `$node/node[@rel="dp"]`); len(n) > 0 {
		return internalHeadPosition(ifLeft(n), q, tr)
		// sometimes co-indexing leads to du's starting at same position ...
	}

	if n := FIND(q, `$node/node[@rel="nucl"]`); len(n) > 0 {
		return internalHeadPosition(if1(n), q, tr)
	}

	if n := FIND(q, `$node/node[@cat="du"]`); len(n) > 0 { // is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case...
		return internalHeadPosition(if1(n), q, tr)
	}

	if TEST(q, `$node[@word]`) {
		return i1(FIND(q, `$node/@end`))
	}

	/*
		distinguish empty nodes due to gapping/RNR from nonlocal dependencies
		use 'empty head' as string to signal precence of gapping/RNR
	*/
	if TEST(q, `$node[@index and not(@word or @cat)]`) {
		if TEST(q, `$node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)]`) {
			return internalHeadPosition(
				FIND(q, `$node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)]`),
				q,
				tr)
		}
		return empty_head
	}

	panic(tracer("No internal head", tr, q))
}

func internalHeadPositionWithGapping(node []interface{}, q *context, tr []trace) int {
	if len(node) == 1 {
		tr = append(tr, trace{s: "internalHeadPositionWithGapping", node: node[0].(*nodeType)})
	} else {
		tr = append(tr, trace{s: fmt.Sprintf("internalHeadPositionWithGapping with %d nodes", len(node))})
	}
	if hdPos := internalHeadPosition(node, q, tr); hdPos == empty_head {
		return internalHeadPositionOfGappedConstituent(node, q, tr)
	} else {
		return hdPos
	}
}

func internalHeadPositionOfGappedConstituent(node []interface{}, q *context, tr []trace) int {
	depthCheck(q, "internalHeadPositionOfGappedConstituent")
	if len(node) == 1 {
	} else {
		tr = append(tr, trace{s: fmt.Sprintf("internalHeadPositionOfGappedConstituent with %d nodes", len(node))})
	}

	if TEST(q, `$node/node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="hd"]`), q, tr) // aux, prepositions
	}

	if TEST(q, `$node[ node[@rel="hd" and @ud:pos="AUX"] and node[@rel=("vc","predc")] ]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel=("vc","predc")]`), q, tr)
	}

	if TEST(q, `$node/node[@rel="su" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="su"]`), q, tr) // 44 van 87 in lassysmall
	}

	if TEST(q, `$node[@rel="vc" and ../node[@rel="su" and (@pt or @cat)]]`) {
		// subject realized inside the vc = funny serialization
		return internalHeadPositionWithGapping(FIND(q, `$node/../node[@rel="su"]`), q, tr)
	}

	if TEST(q, `$node/node[@rel="obj1" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="obj1"]`), q, tr)
	}

	if TEST(q, `$node/node[@rel="predc" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="predc"]`), q, tr)
	}

	if TEST(q, `$node/node[@rel="vc" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="vc"][1]`), q, tr)
	}

	if TEST(q, `$node/node[@rel="pc" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="pc"][1]`), q, tr)
	}

	if n := FIND(q, `$node/node[@rel="mod" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q, tr)
	}

	if n := FIND(q, `$node/node[@rel="app" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q, tr)
	}

	if n := FIND(q, `$node/node[@rel="det" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q, tr)
	}

	if n := FIND(q, `$node/node[@rel="body" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q, tr)
	}

	if n := FIND(q, `$node/node[@rel="cnj" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q, tr)
	}

	if n := FIND(q, `$node/node[@rel="dp" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q, tr)
	}
	if n := FIND(q, `$node/node[@rel="hd" and @ud:pos="ADP"]`); len(n) > 0 { // in en rond Brussel, case not necessary in xquery code (run-time issue?)
		return internalHeadPositionWithGapping(if1(n), q, tr)
	}

	panic(tracer("No internal head in gapped constituent", tr, q))
}

/*
brute force method to be compliant with conj points to the left rule:
if interhdpos($node) < internalhdpos($node/..) then do something ad hoc
because even fixing misplaced heads fails in cases like
Het front der activisten vertoont dan wel een beeld van lusteloosheid , " aan de basis " is en wordt toch veel werk verzet .
*/
func headPositionOfConjunction(node *nodeType, q *context, tr []trace) int {

	tr = append(tr, trace{s: "headPositionOfConjunction", node: node})

	internal_head := internalHeadPositionWithGapping([]interface{}{node}, q, tr)
	leftmost_conj_daughter := nLeft(FIND(q, `$node/../node[@rel="cnj"]`))
	leftmost_internal_head := internalHeadPositionWithGapping([]interface{}{leftmost_conj_daughter}, q, tr)

	if leftmost_internal_head < internal_head {
		return leftmost_internal_head
	}

	endpos_of_leftmost_conj_constituents := []int{}
	for _, e := range leftmost_conj_daughter.Node {
		if e.End < internal_head {
			endpos_of_leftmost_conj_constituents = append(endpos_of_leftmost_conj_constituents, e.End)
		}
	}
	if len(endpos_of_leftmost_conj_constituents) == 0 {
		return leftmost_conj_daughter.Node[0].End // this should not happen really -- give error msg?
	}
	sort.Ints(endpos_of_leftmost_conj_constituents)
	return endpos_of_leftmost_conj_constituents[len(endpos_of_leftmost_conj_constituents)-1]
}

func followingCnjSister(node *nodeType, q *context, tr []trace) []interface{} {
	tr = append(tr, trace{s: "followingCnjSister", node: node})
	/*
	   declare function local:following-cnj-sister($node as element(node)) as element(node)
	   { let $annotated-sisters :=
	         for $sister in $node/../node[@rel="cnj"]
	         return
	            <begin-node begin="{local:begin-position-of-first-word($sister)}">
	              {$sister}
	            </begin-node>

	     let $sorted-sisters :=
	         for $sister in $annotated-sisters
	         (: where $sister[number(@begin) > $node/number(@begin)] :)
	         order by $sister/number(@begin)
	         return $sister
	     return
	         if  ($sorted-sisters[number(@begin) > $node/number(@begin)] )
	         then ($sorted-sisters[number(@begin) > $node/number(@begin)]/node)[1]
	         else $sorted-sisters[1]/node

	   };
	*/

	// TODO: klopt dit ???

	sisters := []*nodeType{}
	for _, n := range node.parent.Node {
		if n.Rel == "cnj" /* && n.Begin > node.Begin */ {
			b := FIND(q, `$n/descendant-or-self::node[@word]/@begin`)
			if len(b) == 0 {
				n.udFirstWordBegin = 1000000
			} else {
				sort.Slice(b, func(i, j int) bool {
					return b[i].(int) < b[j].(int)
				})
				n.udFirstWordBegin = b[0].(int)
			}
			sisters = append(sisters, n)
		}
	}
	sort.Slice(sisters, func(i, j int) bool {
		return sisters[i].udFirstWordBegin < sisters[j].udFirstWordBegin
	})
	for _, n := range sisters {
		if n.udFirstWordBegin > node.Begin {
			return []interface{}{n}
		}
	}
	if len(sisters) > 0 {
		return []interface{}{sisters[0]}
	}
	return []interface{}{}
}
