//
// GENERATED FILE -- DO NOT EDIT
//

package main

import (
	"sort"
)

// recursive
func externalHeadPosition(nodes []interface{}, q *Context) int {
	if depthCheck(q, "externalHeadPosition") {
		return ERROR_RECURSION_LIMIT
	}

	if len(nodes) == 0 {
		return ERROR_EXTERNAL_HEAD_MUST_HAVE_ONE_ARG
	}

	node := nodes[0].(*NodeType)

	if node.Rel == "hd" && (node.udPos == "ADP" || node.parent.Cat == "pp") {
		// vol vertrouwen
		if n := Find(q /* $node/../node[@rel="predc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		}); len(n) > 0 {
			// met als titel
			return internalHeadPosition(n[:1], q)
		}
		if obj1_vc_se_me := Find(q /* $node/../node[@rel=("obj1","vc","se","me")] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		}); len(obj1_vc_se_me) > 0 {
			// adding pt/cat enough for gapping cases?
			if Test(q /* $obj1_vc_se_me[@pt or @cat] */, &XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: obj1_vc_se_me,
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
			}) {
				return internalHeadPositionWithGapping(if1(obj1_vc_se_me), q)
			}
			if Test(q /* $obj1_vc_se_me[@index = ancestor::node/node[@rel=("rhd","whd")]/@index] */, &XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: obj1_vc_se_me,
						},
						arg2: &Sort{
							arg1: &Equal{
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
			}) {
				return internalHeadPosition(Find(q, /* $node/ancestor::node/node[@rel=("rhd","whd")
					   and @index = $node/../node[@rel=("obj1","vc","se","me")]/@index] */&XPath{
						arg1: &Sort{
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG: collect__ancestors__node,
									arg1: &Variable{
										VAR: node,
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
															VAR: node,
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
					}), q)
			}
			if pobj1 := Find(q /* $node/../node[@rel="pobj1"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
			}); len(pobj1) > 0 {
				return internalHeadPosition(if1(pobj1), q)
			}
			// in de eerste rond --> typo in LassySmall/Wiki , binnen en [advp later buiten ]
			return externalHeadPosition(node.axParent, q)
		} else {
			return externalHeadPosition(node.axParent, q)
		}
	}

	aux := auxiliary1(node, q)

	if node.Rel == "hd" && (aux == "aux" || aux == "aux:pass") {
		// aux aux:pass cop
		if vc_predc := Find(q /* $node/../node[@rel=("vc","predc")] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		}); len(vc_predc) > 0 {
			if Test(q /* $vc_predc[@pt or (@cat and node[@pt or @cat])] */, &XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: vc_predc,
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
			}) {
				// skip vc with just empty nodes
				return internalHeadPositionWithGapping(if1(vc_predc), q)
			}
		}
		// if ($node/../node[@rel="predc"]/@index = $node/../../node[@rel="whd"]/@index)
		//     then local:internal_head_position($node/../../node[@rel="whd"])
		return externalHeadPosition(node.axParent, q) // gapping, but does it ever occur with aux?? with cop: hij was en blijft nog steeds een omstreden figuur
	}

	if node.Rel == "hd" && aux == "cop" {
		predc := Find(q /* $node/../node[@rel="predc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		})
		if len(predc) > 0 && Test(q /* $predc[@pt or @cat] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: predc,
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
		}) {
			return internalHeadPositionWithGapping(if1(predc), q)
		}
		if Test(q /* $node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index */, &XPath{
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
									VAR: node,
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
									VAR: node,
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
			return internalHeadPosition(Find(q /* $node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__ancestors__node,
							arg1: &Variable{
								VAR: node,
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
													VAR: node,
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
		return externalHeadPosition(node.axParent, q) // gapping, but could it??
	}

	if node.Rel == "hd" || node.Rel == "nucl" || node.Rel == "body" {
		if n := Find(q /* $node/../node[@rel="hd" and @begin < $node/@begin] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
										VAR: node,
									},
								},
							},
						},
					},
				},
			},
		}); len(n) > 0 {
			return internalHeadPosition(list(n), q) // dan moet je moet
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "predc" {
		if Test(q /* $node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
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
			if Test(q /* $node/../node[@rel="hd" and @ud:pos="ADP"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
				return externalHeadPosition(node.axParent, q) // met als presentator Bruno W , met als gevolg [vc dat ...]
			}
			return internalHeadPosition(Find(q /* $node/../node[@rel="hd"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
		if Test(q /* $node/parent::node[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__parent__node,
					arg1: &Variable{
						VAR: node,
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
			return internalHeadPosition(Find(q /* $node/../node[@rel="hd"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
		if Test(q /* $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
			return internalHeadPosition(Find(q /* $node/../node[@rel="hd"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
		return externalHeadPosition(node.axParent, q) // covers gapping as well?
	}

	if Test(q /* $node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		if predc := Find(q /* $node/../node[@rel="predc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		}); len(predc) > 0 {
			return internalHeadPosition(predc, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if Test(q /* $node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		if vc := Find(q /* $node/../node[@rel="vc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		}); len(vc) > 0 {
			return internalHeadPosition(vc, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if Test(q /* $node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		if Test(q /* $node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		return externalHeadPosition(node.axParent, q) // gapping
	}

	if Test(q /* $node[@rel=("cnj","dp","mwp")] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		if node == nLeft(Find(q /* $node/../node[@rel=("cnj","dp","mwp")] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
			return externalHeadPosition(node.axParent, q)
		}
		if node.Rel == "cnj" {
			return headPositionOfConjunction(node, q)
		}
		return internalHeadPositionWithGapping(node.axParent, q)
	}

	if Test(q /* $node[@rel="cmp" and ../node[@rel="body"]] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/../node[@rel="body"][1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
						},
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
									DATA: []interface{}{"body"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
						arg2: &Function{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if node.Rel == "--" && node.Cat != "" {
		if node.Cat == "mwu" {
			if Test(q /* $node/../node[@cat and not(@cat="mwu")] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
				return internalHeadPosition(Find(q /* $node/../node[@cat and not(@cat="mwu")][1] */, &XPath{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									VAR: node,
								},
							},
							arg2: &Predicate{
								arg1: &Predicate{
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
								arg2: &Function{
									ARG: function__first__0__args,
								},
							},
						},
					},
				}), q)
			}
			return externalHeadPosition(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "--" && node.udPos != "" {
		if Test(q, /* $node[@ud:pos = ("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")
			   and ../node[@rel="--" and
			               not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ]
			  ] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: node,
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
			return internalHeadPositionWithGapping(Find(q /* $node/../node[@rel="--" and not(@ud:pos=("PUNCT","SYM","X","CONJ","NOUN","ADP","ADV","DET","PROPN","NUM","PRON"))][1] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
							},
						},
						arg2: &Predicate{
							arg1: &Predicate{
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
							arg2: &Function{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}), q)
		}
		if n := Find(q /* $node/../node[@cat][1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
						},
					},
					arg2: &Predicate{
						arg1: &Predicate{
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
						},
						arg2: &Function{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}); len(n) > 0 {
			return internalHeadPosition(n, q)
		}
		if Test(q /* $node[@ud:pos="PUNCT" and count(../node) > 1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
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
			if n := Find(q /* $node/../node[not(@ud:pos="PUNCT")][1] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
							},
						},
						arg2: &Predicate{
							arg1: &Predicate{
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
							arg2: &Function{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}); len(n) > 0 {
				return internalHeadPosition(n, q)
			}
			if node == nLeft(Find(q /* $node/../node[@rel="--" and (@cat or @pt)] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
				return externalHeadPosition(node.axParent, q)
			}
			return 1000 // ie end of first punct token
		}
		if node.parent.Begin >= 0 {
			return externalHeadPosition(node.axParent, q)
		}
		return ERROR_NO_HEAD_FOUND
	}

	if node.Rel == "dlink" || node.Rel == "sat" || node.Rel == "tag" {
		if n := Find(q /* $node/../node[@rel="nucl"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		}); len(n) > 0 {
			return internalHeadPositionWithGapping(n, q)
		}
		return ERROR_NO_EXTERNAL_HEAD
	}

	if node.Rel == "vc" {
		if Test(q, /* $node/../node[@rel="hd" and
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
									VAR: node,
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
																	VAR: node,
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
												VAR: node,
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
			return externalHeadPosition(node.axParent, q)
		}
		if Test(q /* $node/../@cat="pp" */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__cat,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
									VAR: node,
								},
							},
						},
					},
				},
			},
		}) { // eraan dat
			return externalHeadPosition(node.axParent, q)
		}
		if Test(q /* $node/../node[@rel=("hd","su") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "whd" || node.Rel == "rhd" {
		if node.Index > 0 {
			return externalHeadPosition(Find(q /* ($node/../node[@rel="body"]//node[@index = $node/@index ])[1] */, &XPath{
				arg1: &Sort{
					arg1: &Filter{
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
												VAR: node,
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
												VAR: node,
											},
										},
									},
								},
							},
						},
						arg2: &Sort{
							arg1: &Function{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}), q)
		}
		return internalHeadPosition(Find(q /* $node/../node[@rel="body"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
	if node.Rel == "crd" {
		tmp := followingCnjSister(node, q)
		return internalHeadPositionWithGapping(ifZ(Find(q, /* $node/../node[@rel="cnj" and
			   @begin=$tmp/@begin and
			   @end=$tmp/@end
			  ] */&XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
												VAR: tmp,
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
											VAR: tmp,
										},
									},
								},
							},
						},
					},
				},
			})), q)
	}

	if node.Rel == "su" {
		if Test(q /* $node/../node[@rel="hd" and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		if Test(q /* $node[../node[@rel="predc"] and not(../node[@rel="hd"])] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
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
			return internalHeadPosition(Find(q /* $node/../node[@rel="predc"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
		return externalHeadPosition(node.axParent, q) // this probably does no change anything, as we are still attaching to head of left conjunct
	}

	if node.Rel == "obj1" {
		if Test(q /* $node/../node[@rel=("hd","su") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "pc" {
		if Test(q /* $node/../node[@rel=("hd","su","obj1") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "mod" {
		if Test(q /* $node/../node[@rel=("hd","su","obj1","pc","predc","body") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		if n := Find(q /* $node/../node[@rel="mod" and (@cat or @pt)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		}); len(n) > 0 {
			if node == nLeft(n) { // gapping with multiple mods
				return externalHeadPosition(node.axParent, q)
			}
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		if Test(q /* $node/../../node[@rel="su" and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: node,
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
			return internalHeadPosition(Find(q /* $node/../../node[@rel="su"] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									VAR: node,
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
		return externalHeadPosition(node.axParent, q) /* an empty mod in an otherwise empty tree
		   -- mod is co-indexed with rhd, rest is elided,
		   LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml */
	}

	if node.Rel == "app" || node.Rel == "det" {
		if Test(q /* $node/../node[@rel=("hd","mod") and (@pt or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "top" {
		return 0
	}

	if node.Rel != "hd" { // TODO: klopt dit?
		return internalHeadPositionWithGapping(node.axParent, q)
	}

	return ERROR_NO_EXTERNAL_HEAD
}

// recursive
func internalHeadPosition(nodes []interface{}, q *Context) int {
	if depthCheck(q, "internalHeadPosition") {
		return ERROR_RECURSION_LIMIT
	}

	if n := len(nodes); n == 0 {
		return ERROR_NO_INTERNAL_HEAD_POSITION_FOUND
	} else if n > 1 {
		return ERROR_MORE_THAN_ONE_INTERNAL_HEAD_POSITION_FOUND
	}
	node := nodes[0]

	if Test(q /* $node[@cat="pp"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		if n := Find(q /* $node/node[@rel=("obj1","pobj1","se")][1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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
									DATA: []interface{}{"obj1", "pobj1", "se"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
						arg2: &Function{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}); len(n) > 0 {
			return internalHeadPosition(n, q)
		}
		if n := Find(q /* $node/node[@rel="hd"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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
		}); len(n) > 0 {
			// if ($node/@cat="mwu")  ( mede [op grond hiervan] )
			//     local:internal_head_position($node/node[@rel="hd"] )
			return internalHeadPosition(n, q)
		}
		return internalHeadPosition(Find(q /* $node/node[1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
					},
					arg2: &Predicate{
						arg1: &Function{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if Test(q /* $node[@cat="mwu"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		return Find(q /* $node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__attributes__end,
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Variable{
							VAR: node,
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

	if Test(q /* $node[@cat="conj"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPosition(ifLeft(Find(q /* $node/node[@rel="cnj"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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

	if predc := Find(q /* $node/node[@rel="predc"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(predc) > 0 {
		if Test(q /* $node/node[@rel="hd" and @ud:pos="AUX"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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
			return internalHeadPosition(predc, q)
		}
		hd := Find(q /* $node/node[@rel="hd"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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
		})
		if len(hd) == 0 { // cases where copula is missing by accident (ungrammatical, not gapping)
			return internalHeadPosition(predc, q)
		}
		return internalHeadPosition(hd, q)
	}

	if Test(q, /* $node[node[@rel="vc"] and
		   node[@rel="hd" and
		        ( @ud:pos="AUX" or
		          $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
		         )
		       ]
		  ] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
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
																		VAR: node,
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
		return internalHeadPosition(Find(q /* $node/node[@rel="vc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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

	if n := Find(q /* $node/node[@rel="hd"][1] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
								DATA: []interface{}{"hd"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
					arg2: &Function{
						ARG: function__first__0__args,
					},
				},
			},
		},
	}); len(n) > 0 {
		return internalHeadPosition(n, q)
	}

	if n := Find(q /* $node/node[@rel="body"][1] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
								DATA: []interface{}{"body"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
					},
					arg2: &Function{
						ARG: function__first__0__args,
					},
				},
			},
		},
	}); len(n) > 0 {
		return internalHeadPosition(n, q)
	}

	if n := Find(q /* $node/node[@rel="dp"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPosition(ifLeft(n), q)
		// sometimes co-indexing leads to du's starting at same position ...
	}

	if n := Find(q /* $node/node[@rel="nucl"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPosition(if1(n), q)
	}

	if n := Find(q /* $node/node[@cat="du"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 { // is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case...
		return internalHeadPosition(if1(n), q)
	}

	if Test(q /* $node[@word] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		return i1(Find(q /* $node/@end */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__attributes__end,
					arg1: &Variable{
						VAR: node,
					},
				},
			},
		}))
	}

	/*
		distinguish empty nodes due to gapping/RNR from nonlocal dependencies
		use 'empty head' as string to signal precence of gapping/RNR
	*/
	if Test(q /* $node[@index and not(@word or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		if Test(q /* $node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__descendant__or__self__type__node,
						arg1: &Collect{
							ARG: collect__ancestors__node,
							arg1: &Variable{
								VAR: node,
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
											VAR: node,
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
			return internalHeadPosition(Find(q /* $node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__descendant__or__self__type__node,
							arg1: &Collect{
								ARG: collect__ancestors__node,
								arg1: &Variable{
									VAR: node,
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
											VAR: node,
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
		return ERROR_RECURSION_LIMIT
	}

	if Test(q /* $node/node[@rel="hd" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/node[@rel="hd"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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

	if Test(q /* $node/node[@rel="su" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/node[@rel="su"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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

	if Test(q /* $node[@rel="vc" and ../node[@rel="su" and (@pt or @cat)]] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/../node[@rel="su"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__parent__type__node,
						arg1: &Variable{
							VAR: node,
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

	if Test(q /* $node/node[@rel="obj1" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/node[@rel="obj1"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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

	if Test(q /* $node/node[@rel="predc" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/node[@rel="predc"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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

	if Test(q /* $node/node[@rel="vc" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/node[@rel="vc"][1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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
						arg2: &Function{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if Test(q /* $node/node[@rel="pc" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
		return internalHeadPositionWithGapping(Find(q /* $node/node[@rel="pc"][1] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: node,
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
						arg2: &Function{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if n := Find(q /* $node/node[@rel="mod" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := Find(q /* $node/node[@rel="app" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := Find(q /* $node/node[@rel="det" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := Find(q /* $node/node[@rel="body" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := Find(q /* $node/node[@rel="cnj" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := Find(q /* $node/node[@rel="dp" and (@pt or @cat)] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Variable{
					VAR: node,
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
	}); len(n) > 0 {
		return internalHeadPositionWithGapping(if1(n), q)
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

	internal_head := internalHeadPositionWithGapping([]interface{}{node}, q)
	leftmost_conj_daughter := nLeft(Find(q /* $node/../node[@rel="cnj"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Collect{
					ARG: collect__parent__type__node,
					arg1: &Variable{
						VAR: node,
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
			b := Find(q /* $n/descendant-or-self::node[@word]/@begin */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__attributes__begin,
						arg1: &Collect{
							ARG: collect__descendant__or__self__node,
							arg1: &Variable{
								VAR: n,
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
