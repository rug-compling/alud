//
// GENERATED FILE -- DO NOT EDIT
//

package main

import (
	"sort"
)

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
		if Test(node, q /* $node/../node[@rel="predc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"predc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			// met als titel
			return internalHeadPosition(first(Find(node, q /* $node/../node[@rel="predc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			})), q)
		}
		if Test(node, q /* $node/../node[@rel=("obj1","vc","se","me")] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"obj1", "vc", "se", "me"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			// adding pt/cat enough for gapping cases?
			if Test(node, q /* $node/../node[@rel=("obj1","vc","se","me") and (@pt or @cat)] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"obj1", "vc", "se", "me"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Sort{
									arg1: &Or{
										arg1: &Collect{
											ARG:  collect__attributes__pt,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			}) {
				return internalHeadPositionWithGapping(first(Find(node, q /* $node/../node[@rel=("obj1","vc","se","me")] */, &XPath{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"obj1", "vc", "se", "me"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				})), q)
			}
			if Test(node, q /* $node/../node[@rel=("obj1","vc","se","me") and @index = ancestor::node/node[@rel=("rhd","whd")]/@index] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"obj1", "vc", "se", "me"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__index,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG: collect__attributes__index,
										arg1: &Collect{
											ARG: collect__child__node,
											arg1: &Collect{
												ARG:  collect__ancestors__node,
												arg1: &Node{},
											},
											arg2: &Predicate{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"rhd", "whd"},
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}) {
				return internalHeadPosition(first(Find(node, q, /* $node/ancestor::node/node[@rel=("rhd","whd")
					   and @index = $node/../node[@rel=("obj1","vc","se","me")]/@index] */&XPath{
						arg1: &Sort{
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG: collect__ancestors__node,
									arg1: &Variable{
										ARG: variable__node,
									},
								},
								arg2: &Predicate{
									arg1: &And{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"rhd", "whd"},
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
											},
										},
										arg2: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__index,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG: collect__attributes__index,
												arg1: &Collect{
													ARG: collect__child__node,
													arg1: &Collect{
														ARG: collect__parent__type__node,
														arg1: &Variable{
															ARG: variable__node,
														},
													},
													arg2: &Predicate{
														arg1: &Equal{
															ARG: equal__is,
															arg1: &Collect{
																ARG:  collect__attributes__rel,
																arg1: &Node{},
															},
															arg2: &Elem{
																DATA: []interface{}{"obj1", "vc", "se", "me"},
																arg1: &Collect{
																	ARG:  collect__attributes__rel,
																	arg1: &Node{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					})), q)
			}
			if Test(node, q /* $node/../node[@rel="pobj1"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"pobj1"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}) {
				return internalHeadPosition(first(Find(node, q /* $node/../node[@rel="pobj1"] */, &XPath{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"pobj1"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				})), q)
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
		if Test(node, q /* $node/../node[@rel=("vc","predc") and (@pt or (@cat and node[@pt or @cat]))] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"vc", "predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Sort{
										arg1: &And{
											arg1: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__child__node,
												arg1: &Node{},
												arg2: &Predicate{
													arg1: &Or{
														arg1: &Collect{
															ARG:  collect__attributes__pt,
															arg1: &Node{},
														},
														arg2: &Collect{
															ARG:  collect__attributes__cat,
															arg1: &Node{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			// skip vc with just empty nodes
			return internalHeadPositionWithGapping(first(Find(node, q /* $node/../node[@rel=("vc","predc")] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"vc", "predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			})), q)
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
		if Test(node, q /* $node/../node[@rel="predc" and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPositionWithGapping(first(Find(node, q /* $node/../node[@rel="predc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			})), q)
		}
		if Test(node, q /* $node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__index,
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"predc"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
					arg2: &Collect{
						ARG: collect__attributes__index,
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__ancestors__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"rhd", "whd"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__ancestors__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"rhd", "whd"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__index,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG: collect__attributes__index,
										arg1: &Collect{
											ARG: collect__child__node,
											arg1: &Collect{
												ARG: collect__parent__type__node,
												arg1: &Variable{
													ARG: variable__node,
												},
											},
											arg2: &Predicate{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"predc"},
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return externalHeadPosition(node.parent, q) // gapping, but could it??
	}

	/*
	   else if ($node[@rel=("hd","nucl","body") ] )
	    then if ($node/../node[@rel="hd"]/number(@begin) < $node/number(@begin) )
	          then local:internal_head_position($node/../node[@rel="hd" and number(@begin) < $node/number(@begin)] ) (: dan moet je moet :)
	          else local:external_head_position($node/..)
	*/
	if Test(node, q /* $node[@rel=("hd","nucl","body") ] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"hd", "nucl", "body"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		if Test(node, q /* $node/../node[@rel="hd"]/@begin < $node/@begin */, &XPath{
			arg1: &Sort{
				arg1: &Cmp{
					ARG: cmp__lt,
					arg1: &Collect{
						ARG: collect__attributes__begin,
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"hd"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
					arg2: &Collect{
						ARG: collect__attributes__begin,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/../node[@rel="hd" and @begin < $node/@begin] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"hd"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Cmp{
									ARG: cmp__lt,
									arg1: &Collect{
										ARG:  collect__attributes__begin,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG: collect__attributes__begin,
										arg1: &Variable{
											ARG: variable__node,
										},
									},
								},
							},
						},
					},
				},
			}), q) // dan moet je moet
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
		if Test(node, q /* $node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Sort{
						arg1: &And{
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
								arg2: &Predicate{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"obj1", "se", "vc"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
								},
							},
							arg2: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
								arg2: &Predicate{
									arg1: &And{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"hd"},
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
											},
										},
										arg2: &Sort{
											arg1: &Or{
												arg1: &Collect{
													ARG:  collect__attributes__pt,
													arg1: &Node{},
												},
												arg2: &Collect{
													ARG:  collect__attributes__cat,
													arg1: &Node{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			if Test(node, q /* $node/../node[@rel="hd" and @ud:pos="ADP"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"hd"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"ADP"},
										arg1: &Collect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			}) {
				return externalHeadPosition(node.parent, q) // met als presentator Bruno W , met als gevolg [vc dat ...]
			}
			return internalHeadPosition(Find(node, q /* $node/../node[@rel="hd"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if Test(node, q /* $node/../self::node[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__self__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"np", "ap"},
									arg1: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
							arg2: &Collect{
								ARG:  collect__child__node,
								arg1: &Node{},
								arg2: &Predicate{
									arg1: &And{
										arg1: &And{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"hd"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
											arg2: &Sort{
												arg1: &Or{
													arg1: &Collect{
														ARG:  collect__attributes__pt,
														arg1: &Node{},
													},
													arg2: &Collect{
														ARG:  collect__attributes__cat,
														arg1: &Node{},
													},
												},
											},
										},
										arg2: &Function{
											ARG: function__not__1__args,
											arg1: &Arg{
												arg1: &Sort{
													arg1: &Equal{
														ARG: equal__is,
														arg1: &Collect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &Node{},
														},
														arg2: &Elem{
															DATA: []interface{}{"AUX"},
															arg1: &Collect{
																ARG:  collect__attributes__ud_3apos,
																arg1: &Node{},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			//reduced relatives , make sure head is not empty (ellipsis)
			// also : ap with predc: actief als schrijver
			return internalHeadPosition(Find(node, q /* $node/../node[@rel="hd"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if Test(node, q /* $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"hd"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Sort{
									arg1: &Or{
										arg1: &Collect{
											ARG:  collect__attributes__pt,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
									},
								},
							},
							arg2: &Function{
								ARG: function__not__1__args,
								arg1: &Arg{
									arg1: &Sort{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"AUX", "ADP"},
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) { // [met als titel] -- obj1/vc missing
			return internalHeadPosition(Find(node, q /* $node/../node[@rel="hd"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return externalHeadPosition(node.parent, q) // covers gapping as well?
	}

	/*
	   else if ( $node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	    then if ($node/../node[@rel="predc"] )
	          then local:internal_head_position($node/../node[@rel="predc"])
	          else local:external_head_position($node/..)
	*/
	if Test(node, q /* $node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"obj1", "se", "me"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG: collect__attributes__cat,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
										},
									},
									arg2: &Elem{
										DATA: []interface{}{"pp"},
										arg1: &Collect{
											ARG: collect__attributes__cat,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
									},
								},
								arg2: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG:  collect__parent__type__node,
										arg1: &Node{},
									},
									arg2: &Predicate{
										arg1: &And{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"ADP"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
												},
											},
											arg2: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"hd"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}) {
		if Test(node, q /* $node/../node[@rel="predc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"predc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/../node[@rel="predc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else if ( $node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] )
	    then if ($node/../node[@rel="vc"])
	          then local:internal_head_position($node/../node[@rel="vc"])
	          else local:external_head_position($node/..)
	*/
	if Test(node, q /* $node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"pobj1"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG: collect__attributes__cat,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
										},
									},
									arg2: &Elem{
										DATA: []interface{}{"pp"},
										arg1: &Collect{
											ARG: collect__attributes__cat,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
									},
								},
								arg2: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG:  collect__parent__type__node,
										arg1: &Node{},
									},
									arg2: &Predicate{
										arg1: &And{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"ADP"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
												},
											},
											arg2: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"hd"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}) {
		if Test(node, q /* $node/../node[@rel="vc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"vc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/../node[@rel="vc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"vc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
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
	if Test(node, q /* $node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &And{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"mod"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Function{
								ARG: function__not__1__args,
								arg1: &Arg{
									arg1: &Sort{
										arg1: &Collect{
											ARG: collect__child__node,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
											arg2: &Predicate{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"obj1", "pobj1", "se", "me"},
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG: collect__attributes__cat,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
										},
									},
									arg2: &Elem{
										DATA: []interface{}{"pp"},
										arg1: &Collect{
											ARG: collect__attributes__cat,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
									},
								},
								arg2: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG:  collect__parent__type__node,
										arg1: &Node{},
									},
									arg2: &Predicate{
										arg1: &And{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"hd"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
											arg2: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"ADP"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}) { // mede op grond hiervan
		// daarom dus
		if Test(node, q /* $node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd", "su", "obj1", "vc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPositionWithGapping(node.axParent, q)
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
	if Test(node, q /* $node[@rel=("cnj","dp","mwp")] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"cnj", "dp", "mwp"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		if node == leftmost(Find(node, q /* $node/../node[@rel=("cnj","dp","mwp")] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"cnj", "dp", "mwp"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})) {
			return externalHeadPosition(node.parent, q)
		}
		if node.Rel == "cnj" {
			return headPositionOfConjunction(node, q)
		}
		return internalHeadPositionWithGapping(node.axParent, q)
	}

	/*
	   else if ($node[@rel="cmp" and ../node[@rel="body"]])
	    then local:internal_head_position_with_gapping($node/../node[@rel="body"][1])
	*/
	if Test(node, q /* $node[@rel="cmp" and ../node[@rel="body"]] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"cmp"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG:  collect__parent__type__node,
								arg1: &Node{},
							},
							arg2: &Predicate{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"body"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(first(Find(node, q /* $node/../node[@rel="body"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"body"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})), q)
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
			if Test(node, q /* $node/../node[@cat and not(@cat="mwu")] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Function{
									ARG: function__not__1__args,
									arg1: &Arg{
										arg1: &Sort{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__cat,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"mwu"},
													arg1: &Collect{
														ARG:  collect__attributes__cat,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}) { // fix for multiword punctuation in Alpino output
				return internalHeadPosition(first(Find(node, q /* $node/../node[@cat and not(@cat="mwu")] */, &XPath{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
									arg2: &Function{
										ARG: function__not__1__args,
										arg1: &Arg{
											arg1: &Sort{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__cat,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"mwu"},
														arg1: &Collect{
															ARG:  collect__attributes__cat,
															arg1: &Node{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				})), q)
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
		if Test(node, q, /* $node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
			   and ../node[@rel="--" and
			               not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
			  ] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							ARG: variable__node,
						},
						arg2: &Sort{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"PUNCT", "SYM", "X", "CONJ", "NOUN", "PROPN", "NUM", "ADP", "ADV", "DET", "PRON"},
										arg1: &Collect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &Node{},
										},
									},
								},
								arg2: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG:  collect__parent__type__node,
										arg1: &Node{},
									},
									arg2: &Predicate{
										arg1: &And{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"--"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
											arg2: &Function{
												ARG: function__not__1__args,
												arg1: &Arg{
													arg1: &Sort{
														arg1: &Equal{
															ARG: equal__is,
															arg1: &Collect{
																ARG:  collect__attributes__ud_3apos,
																arg1: &Node{},
															},
															arg2: &Elem{
																DATA: []interface{}{"PUNCT", "SYM", "X", "CONJ", "NOUN", "PROPN", "NUM", "ADP", "ADV", "DET", "PRON"},
																arg1: &Collect{
																	ARG:  collect__attributes__ud_3apos,
																	arg1: &Node{},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}) {
			return internalHeadPositionWithGapping(first(Find(node, q /* $node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"--"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Function{
									ARG: function__not__1__args,
									arg1: &Arg{
										arg1: &Sort{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"PUNCT", "SYM", "X", "CONJ", "NOUN", "ADP", "ADV", "DET", "PROPN", "NUM", "PRON"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			})), q)
		}
		if Test(node, q /* $node/../node[@cat] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Collect{
							ARG:  collect__attributes__cat,
							arg1: &Node{},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(first(Find(node, q /* $node/../node[@cat] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
						},
					},
				},
			})), q)
		}
		if Test(node, q /* $node[@ud:pos="PUNCT" and count(../node) > 1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Sort{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"PUNCT"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
								},
							},
							arg2: &Cmp{
								ARG: cmp__gt,
								arg1: &Function{
									ARG: function__count__1__args,
									arg1: &Arg{
										arg1: &Collect{
											ARG: collect__child__node,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
									},
								},
								arg2: &Elem{
									DATA: []interface{}{1},
									arg1: &Function{
										ARG: function__count__1__args,
										arg1: &Arg{
											arg1: &Collect{
												ARG: collect__child__node,
												arg1: &Collect{
													ARG:  collect__parent__type__node,
													arg1: &Node{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			if Test(node, q /* $node/../node[not(@ud:pos="PUNCT")] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Function{
								ARG: function__not__1__args,
								arg1: &Arg{
									arg1: &Sort{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"PUNCT"},
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}) {
				return internalHeadPosition(first(Find(node, q /* $node/../node[not(@ud:pos="PUNCT")] */, &XPath{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &Function{
									ARG: function__not__1__args,
									arg1: &Arg{
										arg1: &Sort{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"PUNCT"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				})), q)
			}
			if node == leftmost(Find(node, q /* $node/../node[@rel="--" and (@cat or @pt)] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"--"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Sort{
									arg1: &Or{
										arg1: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG:  collect__attributes__pt,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			})) {
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
		if Test(node, q /* $node/../node[@rel="nucl"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"nucl"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPositionWithGapping(Find(node, q /* $node/../node[@rel="nucl"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"nucl"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
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
	if node.Rel == "vc" {
		if Test(node, q, /* $node/../node[@rel="hd" and
			         ( @ud:pos="AUX" or
			           $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
			         )
			     ]
			and not($node/../node[@rel="predc"]) */&XPath{
				arg1: &Sort{
					arg1: &And{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"hd"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"AUX"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
												},
											},
											arg2: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG: collect__attributes__index,
													arg1: &Collect{
														ARG: collect__child__node,
														arg1: &Collect{
															ARG: collect__descendant__or__self__type__node,
															arg1: &Collect{
																ARG: collect__ancestors__node,
																arg1: &Variable{
																	ARG: variable__node,
																},
																arg2: &Predicate{
																	arg1: &Equal{
																		ARG: equal__is,
																		arg1: &Collect{
																			ARG:  collect__attributes__rel,
																			arg1: &Node{},
																		},
																		arg2: &Elem{
																			DATA: []interface{}{"top"},
																			arg1: &Collect{
																				ARG:  collect__attributes__rel,
																				arg1: &Node{},
																			},
																		},
																	},
																},
															},
														},
														arg2: &Predicate{
															arg1: &Equal{
																ARG: equal__is,
																arg1: &Collect{
																	ARG:  collect__attributes__ud_3apos,
																	arg1: &Node{},
																},
																arg2: &Elem{
																	DATA: []interface{}{"AUX"},
																	arg1: &Collect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &Node{},
																	},
																},
															},
														},
													},
												},
												arg2: &Collect{
													ARG:  collect__attributes__index,
													arg1: &Node{},
												},
											},
										},
									},
								},
							},
						},
						arg2: &Function{
							ARG: function__not__1__args,
							arg1: &Arg{
								arg1: &Sort{
									arg1: &Collect{
										ARG: collect__child__node,
										arg1: &Collect{
											ARG: collect__parent__type__node,
											arg1: &Variable{
												ARG: variable__node,
											},
										},
										arg2: &Predicate{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"predc"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}) {
			return externalHeadPosition(node.parent, q)
		}
		if Test(node, q /* $node/../@cat="pp" */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__cat,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
					},
					arg2: &Elem{
						DATA: []interface{}{"pp"},
						arg1: &Collect{
							ARG: collect__attributes__cat,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
						},
					},
				},
			},
		}) { // eraan dat
			return externalHeadPosition(node.parent, q)
		}
		if Test(node, q /* $node/../node[@rel=("hd","su") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd", "su"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else if ($node[@rel="whd" or @rel="rhd"])
	    then if ($node[@index])
	          then local:external_head_position( ($node/../node[@rel="body"]//node[@index = $node/@index ])[1] )
	          else local:internal_head_position($node/../node[@rel="body"])
	*/
	if node.Rel == "whd" || node.Rel == "rhd" {
		if node.Index > 0 {
			return externalHeadPosition(firstnode(Find(node, q /* $node/../node[@rel="body"]//node[@index = $node/@index ] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__descendant__or__self__type__node,
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG: collect__parent__type__node,
									arg1: &Variable{
										ARG: variable__node,
									},
								},
								arg2: &Predicate{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"body"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__index,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG: collect__attributes__index,
									arg1: &Variable{
										ARG: variable__node,
									},
								},
							},
						},
					},
				},
			})), q)
		}
		return internalHeadPosition(Find(node, q /* $node/../node[@rel="body"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"body"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	/*
	   we need to select the original node and not the result of
	      following-cnj-sister, as that has no global context
	      and global context is needed where the hd is an index node...
	      unfortunately, nodes are completely identical in some
	      elliptical cases, select last() as brute force solution
	*/
	/*
	   else if ($node[@rel="crd"])
	    then local:internal_head_position_with_gapping(
	    	           $node/../node[@rel="cnj" and
	          	                 @begin=local:following-cnj-sister($node)/@begin and
	          	                 @end=local:following-cnj-sister($node)/@end
	          	                ][last()] )
	*/
	if node.Rel == "crd" {
		q.vartmp = followingCnjSister(node, q)
		return internalHeadPositionWithGapping(last(Find(node, q, /* $node/../node[@rel="cnj" and
			   @begin=$tmp/@begin and
			   @end=$tmp/@end
			  ] */&XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"cnj"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__begin,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG: collect__attributes__begin,
											arg1: &Variable{
												ARG: variable__tmp,
											},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__end,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG: collect__attributes__end,
										arg1: &Variable{
											ARG: variable__tmp,
										},
									},
								},
							},
						},
					},
				},
			})), q)
	}

	/*
	   else if ($node[@rel="su"])
	    then if ($node/../node[@rel="hd" and (@pt or @cat)]) (: gapping :)
	         then local:internal_head_position_with_gapping($node/..) (: ud head could still be a predc :)
	          (: only for 1 case where verb is missing -- die eigendom ... (no verb)) :)
	         else if ($node[../node[@rel="predc"] and not(../node[@rel="hd"])] )
	              then local:internal_head_position($node/../node[@rel="predc"])
	          	 else local:external_head_position($node/..) (: this probably does no change anything, as we are still attaching to head of left conjunct :)
	*/
	if node.Rel == "su" {
		if Test(node, q /* $node/../node[@rel="hd" and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) { // gapping
			return internalHeadPositionWithGapping(node.axParent, q) // ud head could still be a predc
		}
		// only for 1 case where verb is missing -- die eigendom ... (no verb))
		if Test(node, q /* $node[../node[@rel="predc"] and not(../node[@rel="hd"])] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Sort{
						arg1: &And{
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
								arg2: &Predicate{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"predc"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
								},
							},
							arg2: &Function{
								ARG: function__not__1__args,
								arg1: &Arg{
									arg1: &Sort{
										arg1: &Collect{
											ARG: collect__child__node,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
											arg2: &Predicate{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"hd"},
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/../node[@rel="predc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return externalHeadPosition(node.parent, q) // this probably does no change anything, as we are still attaching to head of left conjunct
	}

	/*
	   else if ($node[@rel="obj1"])
	    then if ($node/../node[@rel=("hd","su") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	          then local:internal_head_position_with_gapping($node/..)
	          else local:external_head_position($node/..)
	*/
	if node.Rel == "obj1" {
		if Test(node, q /* $node/../node[@rel=("hd","su") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd", "su"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) { // gapping, as su but now su could be head as well
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else if ($node[@rel="pc"])
	    then if ($node/../node[@rel=("hd","su","obj1") and (@pt or @cat)]) (: gapping, as su but now su could be head as well :)
	          then local:internal_head_position_with_gapping($node/..)
	          else local:external_head_position($node/..)
	*/
	if node.Rel == "pc" {
		if Test(node, q /* $node/../node[@rel=("hd","su","obj1") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd", "su", "obj1"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) { // gapping, as su but now su could be head as well
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else 1if ($node[@rel="mod"])
	    1then 2if ($node/../node[@rel=("hd","su","obj1","pc","predc","body") and (@pt or @cat)]) (: gapping, as su but now su or obj1  could be head as well :)
	          2then local:internal_head_position_with_gapping($node/..)
	          2else 3if ($node/../node[@rel="mod" and (@cat or @pt)])
	                3then 4if  (deep-equal($node,local:leftmost($node/../node[@rel="mod" and (@pt or @cat)])) ) (: gapping with multiple mods :)
	                      4then local:external_head_position($node/..)
	                      4else local:internal_head_position_with_gapping($node/..)
	                3else 5if ( $node/../../node[@rel="su" and (@pt or @cat)]  )  (: an mod in an otherwise empty tree (after fixing heads in conj) :)
	                     5then local:internal_head_position($node/../../node[@rel="su"])
	                     5else local:external_head_position($node/..) (: an empty mod in an otherwise empty tree
	                                                               -- mod is co-indexed with rhd, rest is elided,
	                                                               LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml :)
	*/
	if node.Rel == "mod" {
		if Test(node, q /* $node/../node[@rel=("hd","su","obj1","pc","predc","body") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd", "su", "obj1", "pc", "predc", "body"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) { // gapping, as su but now su or obj1  could be head as well
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		if Test(node, q /* $node/../node[@rel="mod" and (@cat or @pt)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"mod"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			if node == leftmost(Find(node, q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"mod"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Sort{
									arg1: &Or{
										arg1: &Collect{
											ARG:  collect__attributes__pt,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			})) { // gapping with multiple mods
				return externalHeadPosition(node.parent, q)
			}
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		if Test(node, q /* $node/../../node[@rel="su" and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__node,
							},
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"su"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) { // an mod in an otherwise empty tree (after fixing heads in conj)
			return internalHeadPosition(Find(node, q /* $node/../../node[@rel="su"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									ARG: variable__node,
								},
							},
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"su"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return externalHeadPosition(node.parent, q) /* an empty mod in an otherwise empty tree
		   -- mod is co-indexed with rhd, rest is elided,
		   LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml */
	}

	/*
	   else if ($node[@rel=("app","det")])
	      then if ($node/../node[@rel=("hd","mod") and (@pt or @cat)]) (: gapping with an app (or a det)! :)
	            then local:internal_head_position_with_gapping($node/..)
	            else local:external_head_position($node/..)
	*/
	if node.Rel == "app" || node.Rel == "det" {
		if Test(node, q /* $node/../node[@rel=("hd","mod") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd", "mod"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) { // gapping with an app (or a det)!
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.parent, q)
	}

	/*
	   else if ($node[@rel="top"])
	    then "0"
	*/
	if node.Rel == "top" {
		return 0
	}

	/*
	   else if ( $node[not(@rel="hd")] )
	     then    local:internal_head_position_with_gapping($node/..)
	*/
	if node.Rel != "hd" { // TODO: klopt dit?
		return internalHeadPositionWithGapping(node.axParent, q)
	}

	/*
	   else    'ERROR_NO_EXTERNAL_HEAD'
	*/
	return ERROR_NO_EXTERNAL_HEAD
}

// recursive
func internalHeadPosition(node []interface{}, q *Context) int {
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
	if Test(node, q /* $node[@cat="pp"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__cat,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"pp"},
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		// if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		// then $node/node[@rel="hd"]/@end
		if Test(node, q /* $node/node[@rel=("obj1","pobj1","se")] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"obj1", "pobj1", "se"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(first(Find(node, q /* $node/node[@rel=("obj1","pobj1","se")] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Variable{
							ARG: variable__node,
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"obj1", "pobj1", "se"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			})), q)
		}
		if Test(node, q /* $node/node[@rel="hd"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"hd"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			// if ($node/@cat="mwu")  ( mede [op grond hiervan] )
			//     local:internal_head_position($node/node[@rel="hd"] )
			return internalHeadPosition(Find(node, q /* $node/node[@rel="hd"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Variable{
							ARG: variable__node,
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return internalHeadPosition(first(Find(node, q /* $node/node */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
				},
			},
		})), q)
	}

	/*
	  else if ($node[@cat="mwu"] )
	  then    $node/node[@rel="mwp" and not(../node/number(@begin) < number(@begin))]/@end
	*/
	if Test(node, q /* $node[@cat="mwu"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__cat,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"mwu"},
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		return Find(node, q /* $node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__attributes__end,
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Variable{
							ARG: variable__node,
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"mwp"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Function{
									ARG: function__not__1__args,
									arg1: &Arg{
										arg1: &Sort{
											arg1: &Cmp{
												ARG: cmp__lt,
												arg1: &Collect{
													ARG: collect__attributes__begin,
													arg1: &Collect{
														ARG: collect__child__node,
														arg1: &Collect{
															ARG:  collect__parent__type__node,
															arg1: &Node{},
														},
													},
												},
												arg2: &Collect{
													ARG:  collect__attributes__begin,
													arg1: &Node{},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		})[0].(int)
	}

	/*
	   else if ($node[@cat="conj"])
	   then    local:internal_head_position(local:leftmost($node/node[@rel="cnj"]))
	*/
	if Test(node, q /* $node[@cat="conj"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__cat,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"conj"},
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPosition(left(Find(node, q /* $node/node[@rel="cnj"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"cnj"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})), q)
	}

	/*
	   else if ( $node/node[@rel="predc"] )
	        then if ($node/node[@rel="hd" and @ud:pos="AUX"])
	             then local:internal_head_position($node/node[@rel="predc"])
	             else if ( not($node/node[@rel="hd"]) )      (: cases where copula is missing by accident (ungrammatical, not gapping) :)
	                  then local:internal_head_position($node/node[@rel="predc"])
	                  else local:internal_head_position($node/node[@rel="hd"])
	*/
	if Test(node, q /* $node/node[@rel="predc"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"predc"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		if Test(node, q /* $node/node[@rel="hd" and @ud:pos="AUX"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"hd"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"AUX"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/node[@rel="predc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Variable{
							ARG: variable__node,
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if Test(node, q /* not($node/node[@rel="hd"]) */, &XPath{
			arg1: &Sort{
				arg1: &Function{
					ARG: function__not__1__args,
					arg1: &Arg{
						arg1: &Sort{
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Variable{
									ARG: variable__node,
								},
								arg2: &Predicate{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"hd"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/node[@rel="predc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Variable{
							ARG: variable__node,
						},
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"predc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return internalHeadPosition(Find(node, q /* $node/node[@rel="hd"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"hd"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

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
	if Test(node, q, /* $node[node[@rel="vc"] and
		   node[@rel="hd" and
		        ( @ud:pos="AUX" or
		          $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
		         )
		       ]
		  ] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Sort{
						arg1: &And{
							arg1: &Collect{
								ARG:  collect__child__node,
								arg1: &Node{},
								arg2: &Predicate{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"vc"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
								},
							},
							arg2: &Collect{
								ARG:  collect__child__node,
								arg1: &Node{},
								arg2: &Predicate{
									arg1: &And{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"hd"},
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
											},
										},
										arg2: &Sort{
											arg1: &Or{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"AUX"},
														arg1: &Collect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &Node{},
														},
													},
												},
												arg2: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG: collect__attributes__index,
														arg1: &Collect{
															ARG: collect__child__node,
															arg1: &Collect{
																ARG: collect__descendant__or__self__type__node,
																arg1: &Collect{
																	ARG: collect__ancestors__node,
																	arg1: &Variable{
																		ARG: variable__node,
																	},
																	arg2: &Predicate{
																		arg1: &Equal{
																			ARG: equal__is,
																			arg1: &Collect{
																				ARG:  collect__attributes__rel,
																				arg1: &Node{},
																			},
																			arg2: &Elem{
																				DATA: []interface{}{"top"},
																				arg1: &Collect{
																					ARG:  collect__attributes__rel,
																					arg1: &Node{},
																				},
																			},
																		},
																	},
																},
															},
															arg2: &Predicate{
																arg1: &Equal{
																	ARG: equal__is,
																	arg1: &Collect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &Node{},
																	},
																	arg2: &Elem{
																		DATA: []interface{}{"AUX"},
																		arg1: &Collect{
																			ARG:  collect__attributes__ud_3apos,
																			arg1: &Node{},
																		},
																	},
																},
															},
														},
													},
													arg2: &Collect{
														ARG:  collect__attributes__index,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
		return internalHeadPosition(Find(node, q /* $node/node[@rel="vc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"vc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="hd"])
	   then local:internal_head_position($node/node[@rel="hd"][1])
	*/
	if Test(node, q /* $node/node[@rel="hd"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"hd"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPosition(first(Find(node, q /* $node/node[@rel="hd"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"hd"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})), q)
	}

	/*
	   else if ( $node/node[@rel="body"])
	   then    local:internal_head_position($node/node[@rel="body"][1])
	*/
	if Test(node, q /* $node/node[@rel="body"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"body"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPosition(first(Find(node, q /* $node/node[@rel="body"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"body"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})), q)
	}

	/*
	   else if ( $node/node[@rel="dp"])
	   then    local:internal_head_position(local:leftmost($node/node[@rel="dp"]))
	        (: sometimes co-indexing leads to du's starting at same position ... :)
	*/
	if Test(node, q /* $node/node[@rel="dp"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"dp"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPosition(left(Find(node, q /* $node/node[@rel="dp"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"dp"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})), q)
		// sometimes co-indexing leads to du's starting at same position ...
	}

	/*
	   else if ( $node/node[@rel="nucl"])
	   then    local:internal_head_position($node/node[@rel="nucl"][1])
	*/
	if Test(node, q /* $node/node[@rel="nucl"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"nucl"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPosition(first(Find(node, q /* $node/node[@rel="nucl"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"nucl"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})), q)
	}

	/*
	   else if ( $node/node[@cat="du"]) (: is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case... :)
	   then    local:internal_head_position($node/node[@cat ="du"][1])
	*/
	if Test(node, q /* $node/node[@cat="du"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__cat,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"du"},
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}) { // is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case...
		return internalHeadPosition(first(Find(node, q /* $node/node[@cat ="du"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"du"},
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		})), q)
	}

	/*
	   else if ( $node[@word] )
	   then    $node/@end
	*/
	if Test(node, q /* $node[@word] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &Collect{
						ARG:  collect__attributes__word,
						arg1: &Node{},
					},
				},
			},
		},
	}) {
		return i1(Find(node, q /* $node/@end */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__attributes__end,
					arg1: &Variable{
						ARG: variable__node,
					},
				},
			},
		}))
	}

	/*
	   distinguish empty nodes due to gapping/RNR from nonlocal dependencies
	   use 'empty head' as string to signal precence of gapping/RNR
	*/
	/*
	   else if ($node[@index and not(@word or @cat)] )
	      then if ($node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)] )
	           then local:internal_head_position($node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)] )
	           else "empty head"
	*/
	if Test(node, q /* $node[@index and not(@word or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &And{
						arg1: &Collect{
							ARG:  collect__attributes__index,
							arg1: &Node{},
						},
						arg2: &Function{
							ARG: function__not__1__args,
							arg1: &Arg{
								arg1: &Sort{
									arg1: &Or{
										arg1: &Collect{
											ARG:  collect__attributes__word,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}) {
		if Test(node, q /* $node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__descendant__or__self__type__node,
						arg1: &Collect{
							ARG: collect__ancestors__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"top"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"whd", "rhd"},
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__index,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG: collect__attributes__index,
										arg1: &Variable{
											ARG: variable__node,
										},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Collect{
										ARG:  collect__attributes__word,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(Find(node, q /* $node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__descendant__or__self__type__node,
							arg1: &Collect{
								ARG: collect__ancestors__node,
								arg1: &Variable{
									ARG: variable__node,
								},
								arg2: &Predicate{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"top"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
						arg2: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__index,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG: collect__attributes__index,
										arg1: &Variable{
											ARG: variable__node,
										},
									},
								},
								arg2: &Sort{
									arg1: &Or{
										arg1: &Collect{
											ARG:  collect__attributes__word,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return empty_head
	}

	/*
	   else    'ERROR_NO_INTERNAL_HEAD'
	*/
	return ERROR_NO_INTERNAL_HEAD
}

func internalHeadPositionWithGapping(node []interface{}, q *Context) int {
	if hdPos := internalHeadPosition(node, q); hdPos == empty_head {
		return internalHeadPositionOfGappedConstituent(node, q)
	} else {
		return hdPos
	}
}

func internalHeadPositionOfGappedConstituent(node []interface{}, q *Context) int {
	if depthCheck(q, "internalHeadPositionOfGappedConstituent") {
		return ERROR_NO_INTERNAL_HEAD_IN_GAPPED_CONSTITUENT
	}

	/*
	   if ($node/node[@rel="hd" and (@pt or @cat)])
	    then local:internal_head_position_with_gapping($node/node[@rel="hd"])  (: auxiliaries, prepositions, ... :)
	*/
	if Test(node, q /* $node/node[@rel="hd" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"hd"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* $node/node[@rel="hd"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"hd"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="su" and (@pt or @cat)] )
	         then local:internal_head_position_with_gapping($node/node[@rel="su"]) (: 44 van 87 in lassysmall:)
	*/
	if Test(node, q /* $node/node[@rel="su" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"su"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* $node/node[@rel="su"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"su"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q) // 44 van 87 in lassysmall
	}

	/*
	   else if ($node[@rel="vc" and ../node[@rel="su" and (@pt or @cat)]] )
	     (: subject realized inside the vc = funny serialization :)
	     then local:internal_head_position_with_gapping($node/../node[@rel="su"])
	*/
	if Test(node, q /* $node[@rel="vc" and ../node[@rel="su" and (@pt or @cat)]] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Sort{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"vc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG:  collect__parent__type__node,
								arg1: &Node{},
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"su"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Collect{
												ARG:  collect__attributes__pt,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}) {
		// subject realized inside the vc = funny serialization
		return internalHeadPositionWithGapping(Find(node, q /* $node/../node[@rel="su"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"su"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="obj1" and (@pt or @cat)] )
	         then local:internal_head_position_with_gapping($node/node[@rel="obj1"])
	*/
	if Test(node, q /* $node/node[@rel="obj1" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"obj1"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* $node/node[@rel="obj1"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"obj1"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="predc" and (@pt or @cat)] )
	         then local:internal_head_position_with_gapping($node/node[@rel="predc"])
	*/
	if Test(node, q /* $node/node[@rel="predc" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"predc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* $node/node[@rel="predc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"predc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="vc" and (@pt or @cat)] )
	     then local:internal_head_position_with_gapping($node/node[@rel="vc"][1])
	*/
	if Test(node, q /* $node/node[@rel="vc" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"vc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* $node/node[@rel="vc"][1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"vc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
						arg2: &Elem{
							DATA: []interface{}{1},
						},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="pc" and (@pt or @cat)] )
	     then local:internal_head_position_with_gapping($node/node[@rel="pc"][1])
	*/
	if Test(node, q /* $node/node[@rel="pc" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"pc"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* $node/node[@rel="pc"][1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Predicate{
						arg1: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"pc"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
						arg2: &Elem{
							DATA: []interface{}{1},
						},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="mod" and (@pt or @cat)] )
	         then local:internal_head_position_with_gapping(($node/node[@rel="mod" and (@pt or @cat)])[1])
	*/
	if Test(node, q /* $node/node[@rel="mod" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"mod"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* ($node/node[@rel="mod" and (@pt or @cat)])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"mod"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Collect{
												ARG:  collect__attributes__pt,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &Elem{
						DATA: []interface{}{1},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="app" and (@pt or @cat)] )
	     then local:internal_head_position_with_gapping(($node/node[@rel="app" and (@pt or @cat)])[1])
	*/
	if Test(node, q /* $node/node[@rel="app" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"app"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* ($node/node[@rel="app" and (@pt or @cat)])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"app"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Collect{
												ARG:  collect__attributes__pt,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &Elem{
						DATA: []interface{}{1},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="det" and (@pt or @cat)] )
	     then local:internal_head_position_with_gapping(($node/node[@rel="det" and (@pt or @cat)])[1])
	*/
	if Test(node, q /* $node/node[@rel="det" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"det"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* ($node/node[@rel="det" and (@pt or @cat)])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"det"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Collect{
												ARG:  collect__attributes__pt,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &Elem{
						DATA: []interface{}{1},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="body" and (@pt or @cat)] )
	     then local:internal_head_position_with_gapping(($node/node[@rel="body" and (@pt or @cat)])[1])
	*/
	if Test(node, q /* $node/node[@rel="body" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"body"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* ($node/node[@rel="body" and (@pt or @cat)])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"body"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Collect{
												ARG:  collect__attributes__pt,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &Elem{
						DATA: []interface{}{1},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="cnj" and (@pt or @cat)] )
	     then local:internal_head_position_with_gapping(($node/node[@rel="cnj" and (@pt or @cat)])[1])
	*/
	if Test(node, q /* $node/node[@rel="cnj" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"cnj"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* ($node/node[@rel="cnj" and (@pt or @cat)])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"cnj"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Collect{
												ARG:  collect__attributes__pt,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &Elem{
						DATA: []interface{}{1},
					},
				},
			},
		}), q)
	}

	/*
	   else if ( $node/node[@rel="dp" and (@pt or @cat)] )
	     then local:internal_head_position_with_gapping(($node/node[@rel="dp" and (@pt or @cat)])[1])
	*/
	if Test(node, q /* $node/node[@rel="dp" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					ARG: variable__node,
				},
				arg2: &Predicate{
					arg1: &And{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"dp"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Sort{
							arg1: &Or{
								arg1: &Collect{
									ARG:  collect__attributes__pt,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(Find(node, q /* ($node/node[@rel="dp" and (@pt or @cat)])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"dp"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Sort{
										arg1: &Or{
											arg1: &Collect{
												ARG:  collect__attributes__pt,
												arg1: &Node{},
											},
											arg2: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &Elem{
						DATA: []interface{}{1},
					},
				},
			},
		}), q)
	}

	return ERROR_NO_INTERNAL_HEAD_IN_GAPPED_CONSTITUENT
}

/*
brute force method to be compliant with conj points to the left rule:
if interhdpos($node) < internalhdpos($node/..) then do something ad hoc
because even fixing misplaced heads fails in cases like
Het front der activisten vertoont dan wel een beeld van lusteloosheid , " aan de basis " is en wordt toch veel werk verzet .
*/
func headPositionOfConjunction(node *NodeType, q *Context) int {
	/*
	   declare function local:head_position_of_conjunction($node as element(node)) as xs:string
	   { let $internal_head := local:internal_head_position_with_gapping($node)
	     let $leftmost_conj_daughter := local:leftmost($node/../node[@rel="cnj"])
	     let $leftmost_internal_head := local:internal_head_position_with_gapping($leftmost_conj_daughter)
	     let $endpos_of_leftmost_conj_constituents :=
	                   for $e in $leftmost_conj_daughter/node/@end
	                   where number($e) < number($internal_head)
	                   order by number($e)
	                   return $e
	     return
	     if (number($leftmost_internal_head) < number($internal_head))  (: business as usual :)
	     then $leftmost_internal_head
	     else if ( $endpos_of_leftmost_conj_constituents )
	          then $endpos_of_leftmost_conj_constituents[last()]
	          else ( $leftmost_conj_daughter/node/@end)[1]  (: this should not happen really -- give error msg? :)

	   };
	*/

	internal_head := internalHeadPositionWithGapping([]interface{}{node}, q)
	leftmost_conj_daughter := leftmost(Find(node, q /* $node/../node[@rel="cnj"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Collect{
					ARG: collect__parent__type__node,
					arg1: &Variable{
						ARG: variable__node,
					},
				},
				arg2: &Predicate{
					arg1: &Equal{
						ARG: equal__is,
						arg1: &Collect{
							ARG:  collect__attributes__rel,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"cnj"},
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	}))
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

func followingCnjSister(node *NodeType, q *Context) []interface{} {
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

	sisters := []*NodeType{}
	for _, n := range node.parent.Node {
		if n.Rel == "cnj" /* && n.Begin > node.Begin */ {
			b := Find(n, q /* $node/descendant-or-self::node[@word]/@begin */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__attributes__begin,
						arg1: &Collect{
							ARG: collect__descendant__or__self__node,
							arg1: &Variable{
								ARG: variable__node,
							},
							arg2: &Predicate{
								arg1: &Collect{
									ARG:  collect__attributes__word,
									arg1: &Node{},
								},
							},
						},
					},
				},
			})
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
