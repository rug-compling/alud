package main

// recursive
func externalHeadPosition(node *NodeType, q *Context) int {
	if depthCheck(q, "externalHeadPosition") {
		return ERROR_NO_EXTERNAL_HEAD
	}

	if node.Rel == "hd" && (node.udPos == "ADP" || node.parent.Cat == "pp") {
		// vol vertrouwen
		if TEST(node, q, `$node/../node[@rel="predc"]`) { // TODO: dit kan efficiënter: meerdere keren zoeken naar zelfde set nodes
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

	      else if ($node[@rel=("hd","nucl","body") ] )
	       then if ($node/../node[@rel="hd"]/number(@begin) < $node/number(@begin) )
	             then local:internal_head_position($node/../node[@rel="hd" and number(@begin) < $node/number(@begin)] ) (: dan moet je moet :)
	             else local:external_head_position($node/..)

	      else if ( $node[@rel="predc"] )
	       then if   ($node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]] )
	             then if ( $node/../node[@rel="hd" and @ud:pos="ADP"] )
	                   then local:external_head_position($node/..) (: met als presentator Bruno W , met als gevolg [vc dat ...]:)
	                   else local:internal_head_position($node/../node[@rel="hd"])
	             else if  ( $node/..[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ]  )
	                             (: reduced relatives , make sure head is not empty (ellipsis) :)
	                             (: also : ap with predc: actief als schrijver :)
	                   then local:internal_head_position($node/../node[@rel="hd"])
	                   else if ($node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))] )  (: [met als titel] -- obj1/vc missing :)
	                   	 then local:internal_head_position($node/../node[@rel="hd"])
	                   	 else local:external_head_position($node/..) (: covers gapping as well? :)

	      else if ( $node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	       then if ($node/../node[@rel="predc"] )
	             then local:internal_head_position($node/../node[@rel="predc"])
	             else local:external_head_position($node/..)

	      else if ( $node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	       then if ($node/../node[@rel="vc"])
	             then local:internal_head_position($node/../node[@rel="vc"])
	             else local:external_head_position($node/..)

	      else if ($node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])])   (: mede op grond hiervan :)
	       (: daarom dus :)
	            then if ($node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)] )
	                  then local:internal_head_position_with_gapping($node/..)
	                  else local:external_head_position($node/..) (: gapping :)


	     else if ($node[@rel=("cnj","dp","mwp")])
	      then if ( deep-equal($node,local:leftmost($node/../node[@rel=("cnj","dp","mwp")])) )
	            then local:external_head_position($node/..)
	            else if ($node[@rel="cnj"])
	                 then local:head_position_of_conjunction($node)
	                 else local:internal_head_position_with_gapping($node/..)

	     else if ($node[@rel="cmp" and ../node[@rel="body"]])
	      then local:internal_head_position_with_gapping($node/../node[@rel="body"][1])

	     else if ($node[@rel="--" and @cat] )
	     	then if ($node[@cat="mwu"])
	             then if ($node/../node[@cat and not(@cat="mwu")]  )    (: fix for multiword punctuation in Alpino output :)
	                   then local:internal_head_position($node/../node[@cat and not(@cat="mwu")][1])
	                   else local:external_head_position($node/..)
	       else local:external_head_position($node/..)

	     else if ( $node[@rel="--" and @ud:pos] )
	      then if ($node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
	                     and ../node[@rel="--" and
	                                 not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
	                    ] )
	            then local:internal_head_position_with_gapping($node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))][1])
	            else if ( $node/../node[@cat]  )
	                  then local:internal_head_position($node/../node[@cat][1])
	                  else if ($node[@ud:pos="PUNCT" and count(../node) > 1])
	                        then if ($node/../node[not(@ud:pos="PUNCT")] )
	                              then local:internal_head_position($node/../node[not(@ud:pos="PUNCT")][1])
	                              else if ( deep-equal($node,local:leftmost($node/../node[@rel="--" and (@cat or @pt)]) ) )
	                                    then local:external_head_position($node/..)
	                                    else "1" (: ie end of first punct token :)
	                        else if ($node/..) then local:external_head_position($node/..)
	       else "ERROR_NO_HEAD_FOUND"

	     else if ($node[@rel=("dlink","sat","tag")])
	      then if ($node/../node[@rel="nucl"])
	            then local:internal_head_position_with_gapping($node/../node[@rel="nucl"])
	            else "ERROR_NO_EXTERNAL_HEAD"

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

	     else if ($node[@rel="whd" or @rel="rhd"])
	      then if ($node[@index])
	            then local:external_head_position( ($node/../node[@rel="body"]//node[@index = $node/@index ])[1] )
	            else local:internal_head_position($node/../node[@rel="body"])

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

	     else if ($node[@rel="su"])
	      then if ($node/../node[@rel="hd" and (@pt or @cat)]) (: gapping :)
	           then local:internal_head_position_with_gapping($node/..) (: ud head could still be a predc :)
	            (: only for 1 case where verb is missing -- die eigendom ... (no verb)) :)
	           else if ($node[../node[@rel="predc"] and not(../node[@rel="hd"])] )
	                then local:internal_head_position($node/../node[@rel="predc"])
	            	 else local:external_head_position($node/..) (: this probably does no change anything, as we are still attaching to head of left conjunct :)

	     else if ($node[@rel="obj1"])
	      then if ($node/../node[@rel=("hd","su") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)

	     else if ($node[@rel="pc"])
	      then if ($node/../node[@rel=("hd","su","obj1") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)

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

	   else if ($node[@rel=("app","det")])
	      then if ($node/../node[@rel=("hd","mod") and (@pt or @cat)]) (: gapping with an app (or a det)! :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)

	     else if ($node[@rel="top"])
	      then "0"

	     else if ( $node[not(@rel="hd")] )
	       then    local:internal_head_position_with_gapping($node/..)

	     else    'ERROR_NO_EXTERNAL_HEAD'
	   } ;
	*/
	return TODO
}

// recursive
func internalHeadPosition(node *NodeType, q *Context) int {
	if depthCheck(q, "internalHeadPosition") {
		return ERROR_NO_INTERNAL_HEAD
	}

	if node.Cat == "pp" {
		// if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		// then $node/node[@rel="hd"]/@end
		if TEST(node, q, `$node/node[@rel=("obj1","pobj1","se")]`) {
			return internalHeadPosition(firstnode(FIND(node, q, `$node/node[@rel=("obj1","pobj1","se")]`)), q)
		}
		if TEST(node, q, `$node/node[@rel="hd"]`) {
			// if ($node/@cat="mwu")  ( mede [op grond hiervan] )
			//     local:internal_head_position($node/node[@rel="hd"] )
			return internalHeadPosition(firstnode(FIND(node, q, `$node/node[@rel="hd"]`)), q)
		}
		if len(node.Node) > 0 {
			return internalHeadPosition(node.Node[0], q)
		}
		return internalHeadPosition(noNode, q)
	}

	if node.Cat == "mwu" {
		return firstint(FIND(node, q, `$node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end`))
	}

	/*

	     else if ($node[@cat="conj"])
	     then    local:internal_head_position(local:leftmost($node/node[@rel="cnj"]))

	     else if ( $node/node[@rel="predc"] )
	          then if ($node/node[@rel="hd" and @ud:pos="AUX"])
	               then local:internal_head_position($node/node[@rel="predc"])
	               else if ( not($node/node[@rel="hd"]) )      (: cases where copula is missing by accident (ungrammatical, not gapping) :)
	                    then local:internal_head_position($node/node[@rel="predc"])
	                    else local:internal_head_position($node/node[@rel="hd"])

	     else if ($node[node[@rel="vc"] and
	                    node[@rel="hd" and
	                         ( @ud:pos="AUX" or
	                           $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
	                          )
	                        ]
	                   ]
	             )
	     then    local:internal_head_position($node/node[@rel="vc"])

	     else if ( $node/node[@rel="hd"])
	     then local:internal_head_position($node/node[@rel="hd"][1])

	     else if ( $node/node[@rel="body"])
	     then    local:internal_head_position($node/node[@rel="body"][1])

	     else if ( $node/node[@rel="dp"])
	     then    local:internal_head_position(local:leftmost($node/node[@rel="dp"]))
	          (: sometimes co-indexing leads to du's starting at same position ... :)

	     else if ( $node/node[@rel="nucl"])
	     then    local:internal_head_position($node/node[@rel="nucl"][1])

	     else if ( $node/node[@cat="du"]) (: is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case... :)
	     then    local:internal_head_position($node/node[@cat ="du"][1])

	     else if ( $node[@word] )
	     then    $node/@end

	     (: distinguish empty nodes due to gapping/RNR from nonlocal dependencies
	        use 'empty head' as string to signal precence of gapping/RNR :)
	     else if ($node[@index and not(@word or @cat)] )
	        then if ($node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)] )
	             then local:internal_head_position($node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)] )
	             else "empty head"

	     else    'ERROR_NO_INTERNAL_HEAD'
	   };
	*/
	return TODO
}

// recursive ??
func internalHeadPositionWithGapping(node *NodeType, q *Context) int {
	return TODO
}
