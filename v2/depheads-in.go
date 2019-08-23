// +build ignore

package alud

import (
	"fmt"
	"sort"
)

// recursive
func externalHeadPosition(nodes []interface{}, q *context) int {
	var node *nodeType

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "externalHeadPosition", q, node))
		}
	}()

	depthCheck(q)

	if len(nodes) == 0 {
		panic(fmt.Sprint("External head must have one arg, has ", len(nodes)))
	}

	node = nodes[0].(*nodeType)

	if node.Rel == "hd" && (node.udPos == "ADP" || node.parent.Cat == "pp") {
		// vol vertrouwen
		if n := FIND(q, `$node/../node[@rel="predc"]`); len(n) > 0 {
			// met als titel
			return internalHeadPosition(n[:1], q)
		}
		if obj1_vc_se_me := FIND(q, `$node/../node[@rel=("obj1","vc","se","me")]`); len(obj1_vc_se_me) > 0 {
			// adding pt/cat enough for gapping cases?
			if TEST(q, `$obj1_vc_se_me[@pt or @cat]`) {
				return internalHeadPositionWithGapping(if1(obj1_vc_se_me), q)
			}
			if TEST(q, `$obj1_vc_se_me[@index = ancestor::node/node[@rel=("rhd","whd")]/@index]`) {
				return internalHeadPosition(FIND(q, `$node/ancestor::node/node[@rel=("rhd","whd")
                                       and @index = $node/../node[@rel=("obj1","vc","se","me")]/@index]`), q)
			}
			if pobj1 := FIND(q, `$node/../node[@rel="pobj1"]`); len(pobj1) > 0 {
				return internalHeadPosition(if1(pobj1), q)
			}
			// in de eerste rond --> typo in LassySmall/Wiki , binnen en [advp later buiten ]
			return externalHeadPosition(node.axParent, q)
		} else {
			return externalHeadPosition(node.axParent, q)
		}
	}

	aux, _ := auxiliary1(node, q) // negeer fout, aux is dan ""

	if node.Rel == "hd" && (aux == "aux" || aux == "aux:pass") {
		// aux aux:pass cop
		if vc_predc := FIND(q, `$node/../node[@rel=("vc","predc")]`); len(vc_predc) > 0 {
			if TEST(q, `$vc_predc[@pt or (@cat and node[@pt or @cat])]`) {
				// skip vc with just empty nodes
				return internalHeadPositionWithGapping(if1(vc_predc), q)
			}
		}
		// if ($node/../node[@rel="predc"]/@index = $node/../../node[@rel="whd"]/@index)
		//     then local:internal_head_position($node/../../node[@rel="whd"])
		return externalHeadPosition(node.axParent, q) // gapping, but does it ever occur with aux?? with cop: hij was en blijft nog steeds een omstreden figuur
	}

	if node.Rel == "hd" && aux == "cop" {
		predc := FIND(q, `$node/../node[@rel="predc"]`)
		if len(predc) > 0 && TEST(q, `$predc[@pt or @cat]`) {
			return internalHeadPositionWithGapping(if1(predc), q)
		}
		if TEST(q, `$node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index`) {
			return internalHeadPosition(
				FIND(q, `$node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index]`),
				q)
		}
		return externalHeadPosition(node.axParent, q) // gapping, but could it??
	}

	if node.Rel == "hd" || node.Rel == "nucl" || node.Rel == "body" {
		if n := FIND(q, `$node/../node[@rel="hd" and @begin < $node/@begin]`); len(n) > 0 {
			return internalHeadPosition(list(n), q) // dan moet je moet
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "predc" {
		if TEST(q, `$node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]]`) {
			if TEST(q, `$node/../node[@rel="hd" and @ud:pos="ADP"]`) {
				return externalHeadPosition(node.axParent, q) // met als presentator Bruno W , met als gevolg [vc dat ...]
			}
			return internalHeadPosition(FIND(q, `$node/../node[@rel="hd"]`), q)
		}
		if TEST(q, `$node/parent::node[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ]`) {
			//reduced relatives , make sure head is not empty (ellipsis)
			// also : ap with predc: actief als schrijver
			return internalHeadPosition(FIND(q, `$node/../node[@rel="hd"]`), q)
		}
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))]`) { // [met als titel] -- obj1/vc missing
			return internalHeadPosition(FIND(q, `$node/../node[@rel="hd"]`), q)
		}
		return externalHeadPosition(node.axParent, q) // covers gapping as well?
	}

	if TEST(q, `$node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])]`) {
		if predc := FIND(q, `$node/../node[@rel="predc"]`); len(predc) > 0 {
			return internalHeadPosition(predc, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if TEST(q, `$node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])]`) {
		if vc := FIND(q, `$node/../node[@rel="vc"]`); len(vc) > 0 {
			return internalHeadPosition(vc, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if TEST(q, `$node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])]`) { // mede op grond hiervan
		// daarom dus
		if TEST(q, `$node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)]`) {
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q) // gapping
	}

	if TEST(q, `$node[@rel=("cnj","dp","mwp")]`) {
		if node == nLeft(FIND(q, `$node/../node[@rel=("cnj","dp","mwp")]`)) {
			return externalHeadPosition(node.axParent, q)
		}
		if node.Rel == "cnj" {
			return headPositionOfConjunction(node, q)
		}
		return internalHeadPositionWithGapping(node.axParent, q)
	}

	if TEST(q, `$node[@rel="cmp" and ../node[@rel="body"]]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/../node[@rel="body"][1]`), q)
	}

	if node.Rel == "--" && node.Cat != "" {
		if node.Cat == "mwu" {
			if TEST(q, `$node/../node[@cat and not(@cat="mwu")]`) { // fix for multiword punctuation in Alpino output
				return internalHeadPosition(FIND(q, `$node/../node[@cat and not(@cat="mwu")][1]`), q)
			}
			return externalHeadPosition(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "--" && node.udPos != "" {
		if TEST(q, `$node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
	                   and ../node[@rel="--" and
	                               not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
	                  ]`) {
			return internalHeadPositionWithGapping(
				FIND(q, `$node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))][1]`),
				q)
		}
		if n := FIND(q, `$node/../node[@cat][1]`); len(n) > 0 {
			return internalHeadPosition(n, q)
		}
		if TEST(q, `$node[@ud:pos="PUNCT" and count(../node) > 1]`) {
			if n := FIND(q, `$node/../node[not(@ud:pos="PUNCT")][1]`); len(n) > 0 {
				return internalHeadPosition(n, q)
			}
			if node == nLeft(FIND(q, `$node/../node[@rel="--" and (@cat or @pt)]`)) {
				return externalHeadPosition(node.axParent, q)
			}
			return 1000 // ie end of first punct token
		}
		if node.parent.Begin >= 0 {
			return externalHeadPosition(node.axParent, q)
		}
		panic("No head found")
	}

	if node.Rel == "dlink" || node.Rel == "sat" || node.Rel == "tag" {
		if n := FIND(q, `$node/../node[@rel="nucl"]`); len(n) > 0 {
			return internalHeadPositionWithGapping(n, q)
		}
		panic("No external head")
	}

	if node.Rel == "vc" {
		if TEST(q, `$node/../node[@rel="hd" and
	                            ( @ud:pos="AUX" or
	                              $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                            )
	                        ]
	                   and not($node/../node[@rel="predc"])`) {
			return externalHeadPosition(node.axParent, q)
		}
		if TEST(q, `$node/../@cat="pp"`) { // eraan dat
			return externalHeadPosition(node.axParent, q)
		}
		if TEST(q, `$node/../node[@rel=("hd","su") and (@pt or @cat)]`) {
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "whd" || node.Rel == "rhd" {
		if node.Index > 0 {
			return externalHeadPosition(FIND(q, `($node/../node[@rel="body"]//node[@index = $node/@index ])[1]`), q)
		}
		return internalHeadPosition(FIND(q, `$node/../node[@rel="body"]`), q)
	}

	/*
		we need to select the original node and not the result of
		following-cnj-sister, as that has no global context
		and global context is needed where the hd is an index node...
		unfortunately, nodes are completely identical in some
		elliptical cases, select last() as brute force solution
	*/
	if node.Rel == "crd" {
		tmp := followingCnjSister(node, q)
		return internalHeadPositionWithGapping(ifZ(FIND(q, `$node/../node[@rel="cnj" and
	          	                 @begin=$tmp/@begin and
	          	                 @end=$tmp/@end
	          	                ]`)), q)
	}

	if node.Rel == "su" {
		if TEST(q, `$node/../node[@rel="vc"] and $node/../node[@rel="hd" and
			                  ( @ud:pos="AUX" or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index ) ]`) {
			tmp := internalHeadPositionWithGapping(node.axParent, q) // testing -- dont go to vc as it has no head sometimes...
			if node.Begin < tmp && tmp <= node.End {                 // maybe the different error handling in go code causes diff with xquery script?
				return externalHeadPosition(node.axParent, q)
			}
			// TODO: vervangen door: return tmp
			return internalHeadPositionWithGapping(node.axParent, q) // dont go to vc directly as it might be empty
		}
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) { // gapping
			return internalHeadPositionWithGapping(node.axParent, q) // ud head could still be a predc
		}
		// only for 1 case where verb is missing -- die eigendom ... (no verb))
		if TEST(q, `$node[../node[@rel="predc" and (@pt or @cat)] and not(../node[@rel="hd" and (@pt or @cat)])]`) {
			return internalHeadPositionWithGapping(FIND(q, `$node/../node[@rel="predc"]`), q)
		}
		return externalHeadPosition(node.axParent, q) // this probably does no change anything, as we are still attaching to head of left conjunct
	}

	if node.Rel == "obj1" {
		if TEST(q, `$node/../node[@rel=("hd","su") and (@pt or @cat)]`) { // gapping, as su but now su could be head as well
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "pc" {
		if TEST(q, `$node/../node[@rel=("hd","su","obj1") and (@pt or @cat)]`) { // gapping, as su but now su could be head as well
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "mod" {
		if TEST(q, `$node/../node[@rel=("hd","su","obj1","pc","predc","body") and (@pt or @cat)]`) { // gapping, as su but now su or obj1  could be head as well
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		if n := FIND(q, `$node/../node[@rel="mod" and (@cat or @pt)]`); len(n) > 0 {
			if node == nLeft(n) { // gapping with multiple mods
				return externalHeadPosition(node.axParent, q)
			}
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		if TEST(q, `$node/../../node[@rel="su" and (@pt or @cat)]`) { // an mod in an otherwise empty tree (after fixing heads in conj)
			return internalHeadPosition(FIND(q, `$node/../../node[@rel="su"]`), q)
		}
		return externalHeadPosition(node.axParent, q) /* an empty mod in an otherwise empty tree
		   -- mod is co-indexed with rhd, rest is elided,
		   LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml */
	}

	if node.Rel == "app" || node.Rel == "det" {
		if TEST(q, `$node/../node[@rel=("hd","mod") and (@pt or @cat)]`) { // gapping with an app (or a det)!
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "top" {
		return 0
	}

	if node.Rel != "hd" { // TODO: klopt dit?
		return internalHeadPositionWithGapping(node.axParent, q)
	}

	panic("No external head")
}

// recursive
func internalHeadPosition(nodes []interface{}, q *context) int {

	defer func() {
		if r := recover(); r != nil {
			var n *nodeType
			if len(nodes) == 1 {
				if n1, ok := nodes[0].(*nodeType); ok {
					n = n1
				}
			}
			panic(trace(r, "internalHeadPosition", q, n))
		}
	}()

	depthCheck(q)

	if n := len(nodes); n != 1 {
		if n == 0 {
			panic("No internal head position found")
		} else if n > 1 {
			panic("More than one internal head position found")
		}
	}
	node := nodes[0]

	if TEST(q, `$node[@cat="pp"]`) {
		// if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		// then $node/node[@rel="hd"]/@end
		if n := FIND(q, `$node/node[@rel=("obj1","pobj1","se")][1]`); len(n) > 0 {
			return internalHeadPosition(n, q)
		}
		if n := FIND(q, `$node/node[@rel="hd"]`); len(n) > 0 {
			// if ($node/@cat="mwu")  ( mede [op grond hiervan] )
			//     local:internal_head_position($node/node[@rel="hd"] )
			return internalHeadPosition(n, q)
		}
		return internalHeadPosition(FIND(q, `$node/node[1]`), q)
	}

	if TEST(q, `$node[@cat="mwu"]`) {
		return FIND(q, `$node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end`)[0].(int)
	}

	if TEST(q, `$node[@cat="conj"]`) {
		return internalHeadPosition(ifLeft(FIND(q, `$node/node[@rel="cnj"]`)), q)
	}

	if predc := FIND(q, `$node/node[@rel="predc"]`); len(predc) > 0 {
		if TEST(q, `$node/node[@rel="hd" and @ud:pos="AUX"]`) {
			return internalHeadPosition(predc, q)
		}
		hd := FIND(q, `$node/node[@rel="hd"]`)
		if len(hd) == 0 { // cases where copula is missing by accident (ungrammatical, not gapping)
			return internalHeadPosition(predc, q)
		}
		return internalHeadPosition(hd, q)
	}

	if TEST(q, `$node[node[@rel="vc"] and
	                  node[@rel="hd" and
	                       ( @ud:pos="AUX" or
	                         $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                        )
	                      ]
	                 ]`) {
		return internalHeadPosition(FIND(q, `$node/node[@rel="vc"]`), q)
	}

	if n := FIND(q, `$node/node[@rel="hd"][1]`); len(n) > 0 {
		return internalHeadPosition(n, q)
	}

	if n := FIND(q, `$node/node[@rel="body"][1]`); len(n) > 0 {
		return internalHeadPosition(n, q)
	}

	if n := FIND(q, `$node/node[@rel="dp"]`); len(n) > 0 {
		return internalHeadPosition(ifLeft(n), q)
		// sometimes co-indexing leads to du's starting at same position ...
	}

	if n := FIND(q, `$node/node[@rel="nucl"]`); len(n) > 0 {
		return internalHeadPosition(if1(n), q)
	}

	if n := FIND(q, `$node/node[@cat="du"]`); len(n) > 0 { // is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case...
		return internalHeadPosition(if1(n), q)
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
				q)
		}
		return empty_head
	}

	panic("No internal head")
}

func internalHeadPositionWithGapping(node []interface{}, q *context) int {

	defer func() {
		if r := recover(); r != nil {
			if len(node) == 1 {
				panic(trace(r, "internalHeadPositionWithGapping", q, node[0].(*nodeType)))
			} else {
				panic(trace(r, fmt.Sprintf("internalHeadPositionWithGapping with %d nodes", len(node)), q))
			}
		}
	}()

	if hdPos := internalHeadPosition(node, q); hdPos == empty_head {
		return internalHeadPositionOfGappedConstituent(node, q)
	} else {
		return hdPos
	}
}

func internalHeadPositionOfGappedConstituent(node []interface{}, q *context) int {

	defer func() {
		if r := recover(); r != nil {
			if len(node) == 1 {
				panic(trace(r, "internalHeadPositionOfGappedConstituent", q, node[0].(*nodeType)))
			} else {
				panic(trace(r, fmt.Sprintf("internalHeadPositionOfGappedConstituent with %d nodes", len(node)), q))
			}
		}
	}()

	depthCheck(q)

	if TEST(q, `$node/node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="hd"]`), q) // aux, prepositions
	}

	if TEST(q, `$node/node[@rel="hd" and @ud:pos="AUX"]`) {
		if TEST(q, `$node/node[@rel=("vc","predc") and (@pt or node[@cat or @pt])]`) {
			return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel=("vc","pred")]`), q) // testing should be vc,pred
		} else {
			return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="hd"]`), q)
		}
	}

	if TEST(q, `$node/node[@rel="su" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="su"]`), q) // 44 van 87 in lassysmall
	}

	if TEST(q, `$node[@rel="vc" and ../node[@rel="su" and (@pt or @cat)]]`) {
		// subject realized inside the vc = funny serialization
		return internalHeadPositionWithGapping(FIND(q, `$node/../node[@rel="su"]`), q)
	}

	if TEST(q, `$node/node[@rel="obj1" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="obj1"]`), q)
	}

	if TEST(q, `$node/node[@rel="predc" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="predc"]`), q)
	}

	if TEST(q, `$node/node[@rel="vc" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="vc"][1]`), q)
	}

	if TEST(q, `$node/node[@rel="pc" and (@pt or @cat)]`) {
		return internalHeadPositionWithGapping(FIND(q, `$node/node[@rel="pc"][1]`), q)
	}

	if n := FIND(q, `$node/node[@rel="mod" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := FIND(q, `$node/node[@rel="app" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := FIND(q, `$node/node[@rel="det" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := FIND(q, `$node/node[@rel="body" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := FIND(q, `$node/node[@rel="cnj" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := FIND(q, `$node/node[@rel="dp" and (@pt or @cat)]`); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}
	if n := FIND(q, `$node/node[@rel="hd" and @ud:pos="ADP"]`); len(n) > 0 { // in en rond Brussel, case not necessary in xquery code (run-time issue?)
		return internalHeadPositionWithGapping(if1(n), q)
	}

	panic("No internal head in gapped constituent")
}

/*
brute force method to be compliant with conj points to the left rule:
if interhdpos($node) < internalhdpos($node/..) then do something ad hoc
because even fixing misplaced heads fails in cases like
Het front der activisten vertoont dan wel een beeld van lusteloosheid , " aan de basis " is en wordt toch veel werk verzet .
*/
func headPositionOfConjunction(node *nodeType, q *context) int {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "headPositionOfConjunction", q, node))
		}
	}()

	internal_head := internalHeadPositionWithGapping([]interface{}{node}, q)
	leftmost_conj_daughter := nLeft(FIND(q, `$node/../node[@rel="cnj"]`))
	leftmost_internal_head := internalHeadPositionWithGapping([]interface{}{leftmost_conj_daughter}, q)

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

func followingCnjSister(node *nodeType, q *context) []interface{} {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "followingCnjSister", q, node))
		}
	}()

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
