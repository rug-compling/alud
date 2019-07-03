package main

//
// TODO: dit kan allemaal efficiënter: meerdere keren zoeken naar zelfde set nodes
//

// recursive
func externalHeadPosition(node *NodeType, q *Context) int {
	if depthCheck(q, "externalHeadPosition") {
		return ERROR_NO_EXTERNAL_HEAD
	}

	/*
	   if ($node[@rel="hd" and (@ud:pos="ADP" or ../@cat="pp") ] )  (: vol vertrouwen :)
	     then  if ($node/../node[@rel="predc"] ) (: met als titel :)
	           then local:internal_head_position(($node/../node[@rel="predc"])[1])
	           else  if ($node/../node[@rel=("obj1","vc","se","me")] )
	                 (: adding pt/cat enough for gapping cases? :)
	                 then  if ( $node/../node[@rel=("obj1","vc","se","me") and (@pt or @cat)] )
	                       then local:internal_head_position_with_gapping(($node/../node[@rel= ("obj1","vc","se","me")])[1])
	                       else  if ( $node/../node[@rel=("obj1","vc","se","me") and @index = ancestor::node/node[@rel=("rhd","whd")]/@index] )
	                             then local:internal_head_position($node/ancestor::node/node[@rel=("rhd","whd")
	                                                           and @index = $node/../node[@rel= ("obj1","vc","se","me")]/@index])[1]

	                             else  if ($node/../node[@rel="pobj1"] )
	                                   then local:internal_head_position(($node/../node[@rel="pobj1"])[1] )
	                         (: in de eerste rond --> typo in LassySmall/Wiki , binnen en [advp later buiten ]:)
	                                   else local:external_head_position($node/..)
	                 else local:external_head_position($node/..)
	*/
	if node.Rel == "hd" && (node.udPos == "ADP" || node.parent.Cat == "pp") {
		// vol vertrouwen
		if TEST(node, q, `$node/../node[@rel="predc"]`) {
			// met als titel
			return internalHeadPosition(firstnode(FIND(node, q, `$node/../node[@rel="predc"]`)), q)
		}
		if TEST(node, q, `$node/../node[@rel=("obj1","vc","se","me")]`) {
			// adding pt/cat enough for gapping cases?
			if TEST(node, q, `$node/../node[@rel=("obj1","vc","se","me") and (@pt or @cat)]`) {
				return internalHeadPositionWithGapping(firstnode(FIND(node, q, `$node/../node[@rel=("obj1","vc","se","me")]`)), q)
			}
			if TEST(node, q, `$node/../node[@rel=("obj1","vc","se","me") and @index = ancestor::node/node[@rel=("rhd","whd")]/@index]`) {
				return internalHeadPosition(firstnode(FIND(node, q, `$node/ancestor::node/node[@rel=("rhd","whd")
	                                                             and @index = $node/../node[@rel=("obj1","vc","se","me")]/@index]`)), q)
			}
			if TEST(node, q, `$node/../node[@rel="pobj1"]`) {
				return internalHeadPosition(firstnode(FIND(node, q, `$node/../node[@rel="pobj1"]`)), q)
			}
			// in de eerste rond --> typo in LassySmall/Wiki , binnen en [advp later buiten ]
			return externalHeadPosition(node.parent, q)
		} else {
			return externalHeadPosition(node.parent, q)
		}
	}

	aux := auxiliary1(node, q)

	/*
	   else if ($node[@rel="hd" and local:auxiliary($node)= ("aux","aux:pass")] ) (: aux aux:pass cop :)
	          then if ($node/../node[@rel=("vc","predc") and (@pt or (@cat and node[@pt or @cat]))])  (: skip vc with just empty nodes :)
	                 then local:internal_head_position_with_gapping(($node/../node[@rel=("vc","predc")])[1])
	          (: else if ($node/../node[@rel="predc"]/@index = $node/../../node[@rel="whd"]/@index)
	               then local:internal_head_position($node/../../node[@rel="whd"]) :)
	                             else local:external_head_position($node/..)  (: gapping, but does it ever occur with aux?? with cop: hij was en blijft nog steeds een omstreden figuur :)
	*/
	if node.Rel == "hd" && (aux == "aux" || aux == "aux:pass") {
		// aux aux:pass cop
		if TEST(node, q, `$node/../node[@rel=("vc","predc") and (@pt or (@cat and node[@pt or @cat]))]`) {
			// skip vc with just empty nodes
			return internalHeadPositionWithGapping(firstnode(FIND(node, q, `$node/../node[@rel=("vc","predc")]`)), q)
		}
		// if ($node/../node[@rel="predc"]/@index = $node/../../node[@rel="whd"]/@index)
		//     then local:internal_head_position($node/../../node[@rel="whd"])
		return externalHeadPosition(node.parent, q) // gapping, but does it ever occur with aux?? with cop: hij was en blijft nog steeds een omstreden figuur
	}

	/*
	   else if ($node[@rel="hd" and local:auxiliary($node) eq 'cop'] )
	          then if ($node/../node[@rel="predc" and (@pt or @cat)])
	                then local:internal_head_position_with_gapping(($node/../node[@rel="predc"])[1])
	                        else if ($node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index)
	                              then local:internal_head_position($node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index] )
	                              else local:external_head_position($node/..)  (: gapping, but could it??:)
	*/
	if node.Rel == "hd" && aux == "cop" {
		if TEST(node, q, `$node/../node[@rel="predc" and (@pt or @cat)]`) {
			return internalHeadPositionWithGapping(firstnode(FIND(node, q, `$node/../node[@rel="predc"]`)), q)
		}
		if TEST(node, q, `$node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index`) {
			return internalHeadPosition(FIND(node, q, `$node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index]`)[0].(*NodeType), q)
		}
		return externalHeadPosition(node.parent, q) // gapping, but could it??
	}

	/*
	   else if ($node[@rel=("hd","nucl","body") ] )
	    then if ($node/../node[@rel="hd"]/number(@begin) < $node/number(@begin) )
	          then local:internal_head_position($node/../node[@rel="hd" and number(@begin) < $node/number(@begin)] ) (: dan moet je moet :)
	          else local:external_head_position($node/..)
	*/
	if TEST(node, q, `$node[@rel=("hd","nucl","body") ]`) {
		if TEST(node, q, `$node/../node[@rel="hd"]/@begin < $node/@begin`) {
			return internalHeadPosition(FIND(node, q, `$node/../node[@rel="hd" and @begin < $node/@begin]`)[0].(*NodeType), q) // dan moet je moet
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else 1if ( $node[@rel="predc"] )
	    1then 2if   ($node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]] )
	          2then 3if ( $node/../node[@rel="hd" and @ud:pos="ADP"] )
	                3then local:external_head_position($node/..) (: met als presentator Bruno W , met als gevolg [vc dat ...]:)
	                3else local:internal_head_position($node/../node[@rel="hd"])
	          2else 4if  ( $node/..[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ]  )
	                          (: reduced relatives , make sure head is not empty (ellipsis) :)
	                          (: also : ap with predc: actief als schrijver :)
	                4then local:internal_head_position($node/../node[@rel="hd"])
	                4else 5if ($node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))] )  (: [met als titel] -- obj1/vc missing :)
	                     5then local:internal_head_position($node/../node[@rel="hd"])
	                     5else local:external_head_position($node/..) (: covers gapping as well? :)
	*/
	if node.Rel == "predc" {
		if TEST(node, q, `$node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]]`) {
			if TEST(node, q, `$node/../node[@rel="hd" and @ud:pos="ADP"]`) {
				return externalHeadPosition(node.parent, q) // met als presentator Bruno W , met als gevolg [vc dat ...]
			}
			return internalHeadPosition(FIND(node, q, `$node/../node[@rel="hd"]`)[0].(*NodeType), q)
		}
		if TEST(node, q, `$node/../self::node[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ]`) {
			//reduced relatives , make sure head is not empty (ellipsis)
			// also : ap with predc: actief als schrijver
			return internalHeadPosition(FIND(node, q, `$node/../node[@rel="hd"]`)[0].(*NodeType), q)
		}
		if TEST(node, q, `$node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))]`) { // [met als titel] -- obj1/vc missing
			return internalHeadPosition(FIND(node, q, `$node/../node[@rel="hd"]`)[0].(*NodeType), q)
		}
		return externalHeadPosition(node.parent, q) // covers gapping as well?
	}

	/*
	   else if ( $node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	    then if ($node/../node[@rel="predc"] )
	          then local:internal_head_position($node/../node[@rel="predc"])
	          else local:external_head_position($node/..)
	*/
	if TEST(node, q, `$node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])]`) {
		if TEST(node, q, `$node/../node[@rel="predc"]`) {
			return internalHeadPosition(FIND(node, q, `$node/../node[@rel="predc"]`)[0].(*NodeType), q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else if ( $node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	    then if ($node/../node[@rel="vc"])
	          then local:internal_head_position($node/../node[@rel="vc"])
	          else local:external_head_position($node/..)
	*/
	if TEST(node, q, `$node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])]`) {
		if TEST(node, q, `$node/../node[@rel="vc"]`) {
			return internalHeadPosition(FIND(node, q, `$node/../node[@rel="vc"]`)[0].(*NodeType), q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else if ($node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])])   (: mede op grond hiervan :)
	    (: daarom dus :)
	         then if ($node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)] )
	               then local:internal_head_position_with_gapping($node/..)
	               else local:external_head_position($node/..) (: gapping :)
	*/
	if TEST(node, q, `$node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])]`) { // mede op grond hiervan
		// daarom dus
		if TEST(node, q, `$node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)]`) {
			return internalHeadPositionWithGapping(node.parent, q)
		}
		return externalHeadPosition(node.parent, q) // gapping
	}

	/*
	   else if ($node[@rel=("cnj","dp","mwp")])
	    then if ( deep-equal($node,local:leftmost($node/../node[@rel=("cnj","dp","mwp")])) )
	          then local:external_head_position($node/..)
	          else if ($node[@rel="cnj"])
	               then local:head_position_of_conjunction($node)
	               else local:internal_head_position_with_gapping($node/..)
	*/
	if TEST(node, q, `$node[@rel=("cnj","dp","mwp")]`) {
		if node == leftmost(FIND(node, q, `$node/../node[@rel=("cnj","dp","mwp")]`)) {
			return externalHeadPosition(node.parent, q)
		}
		if node.Rel == "cnj" {
			return headPositionOfConjunction(node, q)
		}
		return internalHeadPositionWithGapping(node.parent, q)
	}

	/*
	   else if ($node[@rel="cmp" and ../node[@rel="body"]])
	    then local:internal_head_position_with_gapping($node/../node[@rel="body"][1])
	*/
	if TEST(node, q, `$node[@rel="cmp" and ../node[@rel="body"]]`) {
		return internalHeadPositionWithGapping(firstnode(FIND(node, q, `$node/../node[@rel="body"]`)), q)
	}

	/*
	   else if ($node[@rel="--" and @cat] )
	   	then if ($node[@cat="mwu"])
	           then if ($node/../node[@cat and not(@cat="mwu")]  )    (: fix for multiword punctuation in Alpino output :)
	                 then local:internal_head_position($node/../node[@cat and not(@cat="mwu")][1])
	                 else local:external_head_position($node/..)
	     else local:external_head_position($node/..)
	*/
	if node.Rel == "--" && node.Cat != "" {
		if node.Cat == "mwu" {
			if TEST(node, q, `$node/../node[@cat and not(@cat="mwu")]`) { // fix for multiword punctuation in Alpino output
				return internalHeadPosition(firstnode(FIND(node, q, `$node/../node[@cat and not(@cat="mwu")]`)), q)
			}
			return externalHeadPosition(node.parent, q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else 1if ( $node[@rel="--" and @ud:pos] )
	    1then 2if ($node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
	                   and ../node[@rel="--" and
	                               not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
	                  ] )
	          2then local:internal_head_position_with_gapping($node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))][1])
	          2else 3if ( $node/../node[@cat]  )
	                3then local:internal_head_position($node/../node[@cat][1])
	                3else 4if ($node[@ud:pos="PUNCT" and count(../node) > 1])
	                      4then 5if ($node/../node[not(@ud:pos="PUNCT")] )
	                            5then local:internal_head_position($node/../node[not(@ud:pos="PUNCT")][1])
	                            5else 6if ( deep-equal($node,local:leftmost($node/../node[@rel="--" and (@cat or @pt)]) ) )
	                                  6then local:external_head_position($node/..)
	                                  6else "1" (: ie end of first punct token :)
	                      4else 7if ($node/..) 7then local:external_head_position($node/..)
	     7else "ERROR_NO_HEAD_FOUND" (: TODO: juiste else ??? ")
	*/
	if node.Rel == "--" && node.udPos != "" {
		if TEST(node, q, `$node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
	                   and ../node[@rel="--" and
	                               not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
	                  ]`) {
			return internalHeadPositionWithGapping(firstnode(FIND(node, q, `$node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))]`)), q)
		}
		if TEST(node, q, `$node/../node[@cat]`) {
			return internalHeadPosition(firstnode(FIND(node, q, `$node/../node[@cat]`)), q)
		}
		if TEST(node, q, `$node[@ud:pos="PUNCT" and count(../node) > 1]`) {
			if TEST(node, q, `$node/../node[not(@ud:pos="PUNCT")]`) {
				return internalHeadPosition(firstnode(FIND(node, q, `$node/../node[not(@ud:pos="PUNCT")]`)), q)
			}
			if node == leftmost(FIND(node, q, `$node/../node[@rel="--" and (@cat or @pt)]`)) {
				return externalHeadPosition(node.parent, q)
			}
			return 1000 // ie end of first punct token
		}
		if node.parent.Begin >= 0 {
			return externalHeadPosition(node.parent, q)
		}
		return ERROR_NO_HEAD_FOUND
	}

	/*
	   else if ($node[@rel=("dlink","sat","tag")])
	    then if ($node/../node[@rel="nucl"])
	          then local:internal_head_position_with_gapping($node/../node[@rel="nucl"])
	          else "ERROR_NO_EXTERNAL_HEAD"
	*/
	if node.Rel == "dlink" || node.Rel == "sat" || node.Rel == "tag" {
		if TEST(node, q, `$node/../node[@rel="nucl"]`) {
			return internalHeadPositionWithGapping(FIND(node, q, `$node/../node[@rel="nucl"]`)[0].(*NodeType), q)
		}
		return ERROR_NO_EXTERNAL_HEAD
	}

	/*
	   else if ($node[@rel="vc"])
	     then if ($node/../node[@rel="hd" and
	                            ( @ud:pos="AUX" or
	                              $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                            )
	                        ]
	                   and not($node/../node[@rel="predc"]) )
	           then local:external_head_position($node/..)
	           else if ($node/../@cat="pp") (: eraan dat .. :)
	                  then local:external_head_position($node/..)
	                  else if ($node/../node[@rel=("hd","su") and (@pt or @cat)] )
	                       then local:internal_head_position_with_gapping($node/..)
	                       else local:external_head_position($node/..)
	*/

	/*
	   else if ($node[@rel="whd" or @rel="rhd"])
	    then if ($node[@index])
	          then local:external_head_position( ($node/../node[@rel="body"]//node[@index = $node/@index ])[1] )
	          else local:internal_head_position($node/../node[@rel="body"])
	*/

	/*
	   (: we need to select the original node and not the result of
	      following-cnj-sister, as that has no global context
	      and global context is needed where the hd is an index node...
	      unfortunately, nodes are completely identical in some
	      elliptical cases, select last() as brute force solution :)
	     else if ($node[@rel="crd"])
	      then local:internal_head_position_with_gapping(
	      	           $node/../node[@rel="cnj" and
	            	                 @begin=local:following-cnj-sister($node)/@begin and
	            	                 @end=local:following-cnj-sister($node)/@end
	            	                ][last()] )
	*/

	/*
	   else if ($node[@rel="su"])
	    then if ($node/../node[@rel="hd" and (@pt or @cat)]) (: gapping :)
	         then local:internal_head_position_with_gapping($node/..) (: ud head could still be a predc :)
	          (: only for 1 case where verb is missing -- die eigendom ... (no verb)) :)
	         else if ($node[../node[@rel="predc"] and not(../node[@rel="hd"])] )
	              then local:internal_head_position($node/../node[@rel="predc"])
	          	 else local:external_head_position($node/..) (: this probably does no change anything, as we are still attaching to head of left conjunct :)
	*/

	/*
	   else if ($node[@rel="obj1"])
	    then if ($node/../node[@rel=("hd","su") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	          then local:internal_head_position_with_gapping($node/..)
	          else local:external_head_position($node/..)
	*/

	/*
	   else if ($node[@rel="pc"])
	    then if ($node/../node[@rel=("hd","su","obj1") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	          then local:internal_head_position_with_gapping($node/..)
	          else local:external_head_position($node/..)
	*/

	/*
	   else if ($node[@rel="mod"])
	    then if ($node/../node[@rel=("hd","su","obj1","pc","predc","body") and (@pt or @cat)]) (: gapping, as su but now su or obj1  could be head as well :)
	          then local:internal_head_position_with_gapping($node/..)
	          else if ($node/../node[@rel="mod" and (@cat or @pt)])
	                then if  (deep-equal($node,local:leftmost($node/../node[@rel="mod" and (@pt or @cat)])) ) (: gapping with multiple mods :)
	                      then local:external_head_position($node/..)
	                      else local:internal_head_position_with_gapping($node/..)
	                else if ( $node/../../node[@rel="su" and (@pt or @cat)]  )  (: an mod in an otherwise empty tree (after fixing heads in conj) :)
	                     then local:internal_head_position($node/../../node[@rel="su"])
	                     else local:external_head_position($node/..) (: an empty mod in an otherwise empty tree
	                                                               -- mod is co-indexed with rhd, rest is elided,
	                                                               LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml :)
	*/

	/*
	   else if ($node[@rel=("app","det")])
	      then if ($node/../node[@rel=("hd","mod") and (@pt or @cat)]) (: gapping with an app (or a det)! :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)
	*/

	/*
	   else if ($node[@rel="top"])
	    then "0"
	*/

	/*
	   else if ( $node[not(@rel="hd")] )
	     then    local:internal_head_position_with_gapping($node/..)
	*/

	/*
	   else    'ERROR_NO_EXTERNAL_HEAD'
	*/

	return TODO
}

// recursive
func internalHeadPosition(node *NodeType, q *Context) int {
	if depthCheck(q, "internalHeadPosition") {
		return ERROR_NO_INTERNAL_HEAD
	}

	/*
		{ if      ($node[@cat="pp"])
		  then    (: if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		          then $node/node[@rel="hd"]/@end
		          else
		          :)  if ($node/node[@rel=("obj1","pobj1","se")])
		               then local:internal_head_position($node/node[@rel=("obj1","pobj1","se")][1])
		               else if ($node/node[@rel="hd"])
		                    then (: if ($node/@cat="mwu")  ( mede [op grond hiervan] )
		                         then local:internal_head_position($node/node[@rel="hd"] )
		                         else :)
		                         local:internal_head_position($node/node[@rel="hd"])
		                    else local:internal_head_position( $node/node[1] )
	*/
	if node.Cat == "pp" {
		// if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		// then $node/node[@rel="hd"]/@end
		if TEST(node, q, `$node/node[@rel=("obj1","pobj1","se")]`) {
			return internalHeadPosition(firstnode(FIND(node, q, `$node/node[@rel=("obj1","pobj1","se")]`)), q)
		}
		if TEST(node, q, `$node/node[@rel="hd"]`) {
			// if ($node/@cat="mwu")  ( mede [op grond hiervan] )
			//     local:internal_head_position($node/node[@rel="hd"] )
			return internalHeadPosition(FIND(node, q, `$node/node[@rel="hd"]`)[0].(*NodeType), q)
		}
		if len(node.Node) > 0 {
			return internalHeadPosition(node.Node[0], q)
		}
		return internalHeadPosition(noNode, q)
	}

	/*
	  else if ($node[@cat="mwu"] )
	  then    $node/node[@rel="mwp" and not(../node/number(@begin) < number(@begin))]/@end
	*/
	if node.Cat == "mwu" {
		return FIND(node, q, `$node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end`)[0].(int)
	}

	/*
	   else if ($node[@cat="conj"])
	   then    local:internal_head_position(local:leftmost($node/node[@rel="cnj"]))
	*/

	/*
	   else if ( $node/node[@rel="predc"] )
	        then if ($node/node[@rel="hd" and @ud:pos="AUX"])
	             then local:internal_head_position($node/node[@rel="predc"])
	             else if ( not($node/node[@rel="hd"]) )      (: cases where copula is missing by accident (ungrammatical, not gapping) :)
	                  then local:internal_head_position($node/node[@rel="predc"])
	                  else local:internal_head_position($node/node[@rel="hd"])
	*/

	/*
	   else if ($node[node[@rel="vc"] and
	                  node[@rel="hd" and
	                       ( @ud:pos="AUX" or
	                         $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                        )
	                      ]
	                 ]
	           )
	   then    local:internal_head_position($node/node[@rel="vc"])
	*/

	/*
	   else if ( $node/node[@rel="hd"])
	   then local:internal_head_position($node/node[@rel="hd"][1])
	*/

	/*
	   else if ( $node/node[@rel="body"])
	   then    local:internal_head_position($node/node[@rel="body"][1])
	*/

	/*
	   else if ( $node/node[@rel="dp"])
	   then    local:internal_head_position(local:leftmost($node/node[@rel="dp"]))
	        (: sometimes co-indexing leads to du's starting at same position ... :)
	*/

	/*
	   else if ( $node/node[@rel="nucl"])
	   then    local:internal_head_position($node/node[@rel="nucl"][1])
	*/

	/*
	   else if ( $node/node[@cat="du"]) (: is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case... :)
	   then    local:internal_head_position($node/node[@cat ="du"][1])
	*/

	/*
	   else if ( $node[@word] )
	   then    $node/@end
	*/

	/*
	   (: distinguish empty nodes due to gapping/RNR from nonlocal dependencies
	      use 'empty head' as string to signal precence of gapping/RNR :)
	   else if ($node[@index and not(@word or @cat)] )
	      then if ($node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)] )
	           then local:internal_head_position($node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)] )
	           else "empty head"
	*/

	/*
	   else    'ERROR_NO_INTERNAL_HEAD'
	*/

	return TODO
}

// recursive ??
func internalHeadPositionWithGapping(node *NodeType, q *Context) int {
	return TODO
}

func headPositionOfConjunction(node *NodeType, q *Context) int {
	return TODO
}
