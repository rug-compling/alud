//
// // THIS IS A GENERATED FILE. DO NOT EDIT.
//

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
		if n := find(q /* $node/../node[@rel="predc"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"predc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}); len(n) > 0 {
			// met als titel
			return internalHeadPositionWithGapping(n[:1], q)
		}
		if obj1_vc_se_me := find(q /* $node/../node[@rel=("obj1","vc","se","me")] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"obj1", "vc", "se", "me"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}); len(obj1_vc_se_me) > 0 {
			if test(q /* $node[../@rel="cnj" and ../node[@index and @rel="obj1"]] */, &xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dSort{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG: collect__attributes__rel,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"cnj"},
										arg1: &dCollect{
											ARG: collect__attributes__rel,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dCollect{
												ARG:  collect__attributes__index,
												arg1: &dNode{},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"obj1"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
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
				if test(q /* $node[../node[@rel="obj1" and (@pt or @cat)] and ../../node[@cat=("pp","ppres") and node[@rel="obj1" and @index and not(@pt or @cat)]]] */, &xPath{
					arg1: &dSort{
						arg1: &dFilter{
							arg1: &dVariable{
								VAR: node,
							},
							arg2: &dSort{
								arg1: &dAnd{
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
										arg2: &dPredicate{
											arg1: &dAnd{
												arg1: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"obj1"},
														arg1: &dCollect{
															ARG:  collect__attributes__rel,
															arg1: &dNode{},
														},
													},
												},
												arg2: &dSort{
													arg1: &dOr{
														arg1: &dCollect{
															ARG:  collect__attributes__pt,
															arg1: &dNode{},
														},
														arg2: &dCollect{
															ARG:  collect__attributes__cat,
															arg1: &dNode{},
														},
													},
												},
											},
										},
									},
									arg2: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG: collect__parent__type__node,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
										arg2: &dPredicate{
											arg1: &dAnd{
												arg1: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"pp", "ppres"},
														arg1: &dCollect{
															ARG:  collect__attributes__cat,
															arg1: &dNode{},
														},
													},
												},
												arg2: &dCollect{
													ARG:  collect__child__node,
													arg1: &dNode{},
													arg2: &dPredicate{
														arg1: &dAnd{
															arg1: &dAnd{
																arg1: &dEqual{
																	ARG: equal__is,
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																	arg2: &dElem{
																		DATA: []interface{}{"obj1"},
																		arg1: &dCollect{
																			ARG:  collect__attributes__rel,
																			arg1: &dNode{},
																		},
																	},
																},
																arg2: &dCollect{
																	ARG:  collect__attributes__index,
																	arg1: &dNode{},
																},
															},
															arg2: &dFunction{
																ARG: function__not__1__args,
																arg1: &dArg{
																	arg1: &dSort{
																		arg1: &dOr{
																			arg1: &dCollect{
																				ARG:  collect__attributes__pt,
																				arg1: &dNode{},
																			},
																			arg2: &dCollect{
																				ARG:  collect__attributes__cat,
																				arg1: &dNode{},
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
						},
					},
				}) {
					// rightmost conj -- attach as conj to first preposition GB 27/2/24
					leftmost_preposition := nLeft(find(q /* $node/../../node[@cat=("pp","ppres") and node[@rel="obj1" and @index and not(@pt or @cat)]]/node[@rel="hd"] */, &xPath{
						arg1: &dSort{
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dCollect{
											ARG: collect__parent__type__node,
											arg1: &dVariable{
												VAR: node,
											},
										},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"pp", "ppres"},
													arg1: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dCollect{
												ARG:  collect__child__node,
												arg1: &dNode{},
												arg2: &dPredicate{
													arg1: &dAnd{
														arg1: &dAnd{
															arg1: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"obj1"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																},
															},
															arg2: &dCollect{
																ARG:  collect__attributes__index,
																arg1: &dNode{},
															},
														},
														arg2: &dFunction{
															ARG: function__not__1__args,
															arg1: &dArg{
																arg1: &dSort{
																	arg1: &dOr{
																		arg1: &dCollect{
																			ARG:  collect__attributes__pt,
																			arg1: &dNode{},
																		},
																		arg2: &dCollect{
																			ARG:  collect__attributes__cat,
																			arg1: &dNode{},
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
								arg2: &dPredicate{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"hd"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
					}))
					return internalHeadPositionWithGapping([]interface{}{leftmost_preposition}, q)
				}
				// otherwise: attach as case to obj1 in 2nd conj
				if test(q /* $node/../../node[@cat=("pp","ppres")]/node[@rel="obj1" and @index and (@pt or @cat)] */, &xPath{
					arg1: &dSort{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dVariable{
											VAR: node,
										},
									},
								},
								arg2: &dPredicate{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"pp", "ppres"},
											arg1: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
										},
									},
								},
							},
							arg2: &dPredicate{
								arg1: &dAnd{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"obj1"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dCollect{
											ARG:  collect__attributes__index,
											arg1: &dNode{},
										},
									},
									arg2: &dSort{
										arg1: &dOr{
											arg1: &dCollect{
												ARG:  collect__attributes__pt,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
					},
				}) {
					return internalHeadPositionWithGapping(find(q,
						/* $node/../../node[@cat=("pp","ppres")]/node[@rel="obj1" and @index and (@pt or @cat)] */ &xPath{
							arg1: &dSort{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG: collect__parent__type__node,
											arg1: &dCollect{
												ARG: collect__parent__type__node,
												arg1: &dVariable{
													VAR: node,
												},
											},
										},
										arg2: &dPredicate{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"pp", "ppres"},
													arg1: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
												},
											},
										},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dAnd{
												arg1: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"obj1"},
														arg1: &dCollect{
															ARG:  collect__attributes__rel,
															arg1: &dNode{},
														},
													},
												},
												arg2: &dCollect{
													ARG:  collect__attributes__index,
													arg1: &dNode{},
												},
											},
											arg2: &dSort{
												arg1: &dOr{
													arg1: &dCollect{
														ARG:  collect__attributes__pt,
														arg1: &dNode{},
													},
													arg2: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
							},
						}), q)
				}
				return externalHeadPosition(node.axParent, q)
			}
			// adding pt/cat enough for gapping cases?
			if test(q /* $obj1_vc_se_me[@pt or @cat] */, &xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: obj1_vc_se_me,
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			}) {
				return internalHeadPositionWithGapping(if1(obj1_vc_se_me), q)
			}
			if test(q /* $obj1_vc_se_me[@index = ancestor::node/node[@rel=("rhd","whd")]/@index] */, &xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: obj1_vc_se_me,
						},
						arg2: &dSort{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__index,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG: collect__attributes__index,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG:  collect__ancestors__node,
											arg1: &dNode{},
										},
										arg2: &dPredicate{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"rhd", "whd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
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
				return internalHeadPosition(find(q, /* $node/ancestor::node/node[@rel=("rhd","whd")
					   and @index = $node/../node[@rel=("obj1","vc","se","me")]/@index] */&xPath{
						arg1: &dSort{
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__ancestors__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"rhd", "whd"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__index,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG: collect__attributes__index,
												arg1: &dCollect{
													ARG: collect__child__node,
													arg1: &dCollect{
														ARG: collect__parent__type__node,
														arg1: &dVariable{
															VAR: node,
														},
													},
													arg2: &dPredicate{
														arg1: &dEqual{
															ARG: equal__is,
															arg1: &dCollect{
																ARG:  collect__attributes__rel,
																arg1: &dNode{},
															},
															arg2: &dElem{
																DATA: []interface{}{"obj1", "vc", "se", "me"},
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
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
			if node.parent.Rel == "cnj" && test(q /* $obj1_vc_se_me[@index = $node/ancestor::node//node[@rel="obj1"]/@index] */, &xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: obj1_vc_se_me,
						},
						arg2: &dSort{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__index,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG: collect__attributes__index,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG: collect__descendant__or__self__type__node,
											arg1: &dCollect{
												ARG: collect__ancestors__node,
												arg1: &dVariable{
													VAR: node,
												},
											},
										},
										arg2: &dPredicate{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"obj1"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
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
			}) { // ten zuiden i en ten noorden van de begraafplaats  GB 26/2/24
				return internalHeadPosition(find(q /* $node/ancestor::node//node[@cat="pp" and node[@rel="obj1" and (@pt or @cat) and @index = $obj1_vc_se_me/@index]]/node[@rel="hd"] */, &xPath{
					arg1: &dSort{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__descendant__or__self__type__node,
									arg1: &dCollect{
										ARG: collect__ancestors__node,
										arg1: &dVariable{
											VAR: node,
										},
									},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"pp"},
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dCollect{
											ARG:  collect__child__node,
											arg1: &dNode{},
											arg2: &dPredicate{
												arg1: &dAnd{
													arg1: &dAnd{
														arg1: &dEqual{
															ARG: equal__is,
															arg1: &dCollect{
																ARG:  collect__attributes__rel,
																arg1: &dNode{},
															},
															arg2: &dElem{
																DATA: []interface{}{"obj1"},
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
																},
															},
														},
														arg2: &dSort{
															arg1: &dOr{
																arg1: &dCollect{
																	ARG:  collect__attributes__pt,
																	arg1: &dNode{},
																},
																arg2: &dCollect{
																	ARG:  collect__attributes__cat,
																	arg1: &dNode{},
																},
															},
														},
													},
													arg2: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG:  collect__attributes__index,
															arg1: &dNode{},
														},
														arg2: &dCollect{
															ARG: collect__attributes__index,
															arg1: &dVariable{
																VAR: obj1_vc_se_me,
															},
														},
													},
												},
											},
										},
									},
								},
							},
							arg2: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"hd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
				}), q)
			}
			// in de eerste rond --> typo in LassySmall/Wiki , binnen en [advp later buiten ]
			return externalHeadPosition(node.axParent, q)
		}
		if pobj1 := find(q /* $node/../node[@rel=("pobj1")] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"pobj1"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}); len(pobj1) > 0 { // removed mod, assume ADP is the head here GB 26/2/21
			return internalHeadPosition(if1(pobj1), q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if aux, err := auxiliary1(node, q); err == nil {
		if node.Rel == "hd" && (aux == "aux" || aux == "aux:pass") {
			// aux aux:pass cop
			if vc_predc := find(q /* $node/../node[@rel=("vc","predc")] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"vc", "predc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}); len(vc_predc) > 0 {
				if test(q /* $vc_predc[@pt or (@cat and node[@pt or @cat])] */, &xPath{
					arg1: &dSort{
						arg1: &dFilter{
							arg1: &dVariable{
								VAR: vc_predc,
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dSort{
										arg1: &dAnd{
											arg1: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG:  collect__child__node,
												arg1: &dNode{},
												arg2: &dPredicate{
													arg1: &dOr{
														arg1: &dCollect{
															ARG:  collect__attributes__pt,
															arg1: &dNode{},
														},
														arg2: &dCollect{
															ARG:  collect__attributes__cat,
															arg1: &dNode{},
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
			predc := find(q /* $node/../node[@rel="predc"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"predc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			})
			if len(predc) > 0 && test(q /* $predc[@pt or @cat] */, &xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: predc,
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			}) {
				return internalHeadPositionWithGapping(if1(predc), q)
			}
			if test(q /* $node/../node[@rel="predc"]/@index = $node/ancestor::node/node[@rel=("rhd","whd")]/@index */, &xPath{
				arg1: &dSort{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG: collect__attributes__index,
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dPredicate{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"predc"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
						arg2: &dCollect{
							ARG: collect__attributes__index,
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__ancestors__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dPredicate{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"rhd", "whd"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
					},
				},
			}) {
				return internalHeadPosition(
					find(q /* $node/ancestor::node/node[@rel=("rhd","whd") and @index = $node/../node[@rel="predc"]/@index] */, &xPath{
						arg1: &dSort{
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__ancestors__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"rhd", "whd"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__index,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG: collect__attributes__index,
												arg1: &dCollect{
													ARG: collect__child__node,
													arg1: &dCollect{
														ARG: collect__parent__type__node,
														arg1: &dVariable{
															VAR: node,
														},
													},
													arg2: &dPredicate{
														arg1: &dEqual{
															ARG: equal__is,
															arg1: &dCollect{
																ARG:  collect__attributes__rel,
																arg1: &dNode{},
															},
															arg2: &dElem{
																DATA: []interface{}{"predc"},
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
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
					}),
					q)
			}
			return externalHeadPosition(node.axParent, q) // gapping, but could it??
		}
	}

	if node.Rel == "hd" || node.Rel == "nucl" || node.Rel == "body" {
		if n := find(q /* $node/../node[@rel="hd" and @begin < $node/@begin] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCmp{
								ARG: cmp__lt,
								arg1: &dCollect{
									ARG:  collect__attributes__begin,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG: collect__attributes__begin,
									arg1: &dVariable{
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
		if test(q /* $node[../node[@rel=("obj1","se","vc")] and ../node[@rel="hd" and (@pt or @cat)]] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
								arg2: &dPredicate{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"obj1", "se", "vc"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
								},
							},
							arg2: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"hd"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dSort{
											arg1: &dOr{
												arg1: &dCollect{
													ARG:  collect__attributes__pt,
													arg1: &dNode{},
												},
												arg2: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
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
			if test(q /* $node/../node[@rel="hd" and @ud:pos="ADP"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"hd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"ADP"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &dNode{},
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
			return internalHeadPosition(find(q /* $node/../node[@rel="hd"][1] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"hd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dFunction{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}), q) //make robust for double hd in CGN GB 19/03/21
		}
		if test(q /* $node/parent::node[@cat=("np","ap") and node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX") ]  ] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__parent__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"np", "ap"},
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCollect{
								ARG:  collect__child__node,
								arg1: &dNode{},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"hd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dSort{
												arg1: &dOr{
													arg1: &dCollect{
														ARG:  collect__attributes__pt,
														arg1: &dNode{},
													},
													arg2: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
												},
											},
										},
										arg2: &dFunction{
											ARG: function__not__1__args,
											arg1: &dArg{
												arg1: &dSort{
													arg1: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &dNode{},
														},
														arg2: &dElem{
															DATA: []interface{}{"AUX"},
															arg1: &dCollect{
																ARG:  collect__attributes__ud_3apos,
																arg1: &dNode{},
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
			return internalHeadPosition(find(q /* $node/../node[@rel="hd"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if test(q /* $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"hd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dSort{
									arg1: &dOr{
										arg1: &dCollect{
											ARG:  collect__attributes__pt,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dFunction{
								ARG: function__not__1__args,
								arg1: &dArg{
									arg1: &dSort{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"AUX", "ADP"},
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
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
			return internalHeadPosition(find(q /* $node/../node[@rel="hd"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if test(q, /* $node[../node[@rel="su" and (@pt or @cat)] and
			   ../node[@rel="hd" and
			           not(@ud:pos="AUX" or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index )]
			   ] */&xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dSort{
							arg1: &dAnd{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"su"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dSort{
												arg1: &dOr{
													arg1: &dCollect{
														ARG:  collect__attributes__pt,
														arg1: &dNode{},
													},
													arg2: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"hd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dFunction{
												ARG: function__not__1__args,
												arg1: &dArg{
													arg1: &dSort{
														arg1: &dOr{
															arg1: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__ud_3apos,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"AUX"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &dNode{},
																	},
																},
															},
															arg2: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG: collect__attributes__index,
																	arg1: &dCollect{
																		ARG: collect__child__node,
																		arg1: &dCollect{
																			ARG: collect__descendant__or__self__type__node,
																			arg1: &dCollect{
																				ARG: collect__ancestors__node,
																				arg1: &dVariable{
																					VAR: node,
																				},
																				arg2: &dPredicate{
																					arg1: &dEqual{
																						ARG: equal__is,
																						arg1: &dCollect{
																							ARG:  collect__attributes__rel,
																							arg1: &dNode{},
																						},
																						arg2: &dElem{
																							DATA: []interface{}{"top"},
																							arg1: &dCollect{
																								ARG:  collect__attributes__rel,
																								arg1: &dNode{},
																							},
																						},
																					},
																				},
																			},
																		},
																		arg2: &dPredicate{
																			arg1: &dEqual{
																				ARG: equal__is,
																				arg1: &dCollect{
																					ARG:  collect__attributes__ud_3apos,
																					arg1: &dNode{},
																				},
																				arg2: &dElem{
																					DATA: []interface{}{"AUX"},
																					arg1: &dCollect{
																						ARG:  collect__attributes__ud_3apos,
																						arg1: &dNode{},
																					},
																				},
																			},
																		},
																	},
																},
																arg2: &dCollect{
																	ARG:  collect__attributes__index,
																	arg1: &dNode{},
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
			return internalHeadPosition(find(q /* $node/../node[@rel="su"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"su"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
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

	if test(q /* $node[@rel=("obj1","se","me") and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"obj1", "se", "me"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG: collect__attributes__cat,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"pp"},
										arg1: &dCollect{
											ARG: collect__attributes__cat,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"ADP"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"hd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
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
		if predc := find(q /* $node/../node[@rel="predc"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"predc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
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

	if test(q /* $node[@rel="pobj1" and (../@cat="pp" or ../node[@ud:pos="ADP" and @rel="hd"])] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"pobj1"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG: collect__attributes__cat,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"pp"},
										arg1: &dCollect{
											ARG: collect__attributes__cat,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"ADP"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"hd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
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
		if vc := find(q /* $node/../node[@rel="vc"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"vc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
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

	if test(q /* $node[@rel="mod" and not(../node[@rel=("obj1","pobj1","se","me")]) and (../@cat="pp" or ../node[@rel="hd" and @ud:pos="ADP"])] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dAnd{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"mod"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dFunction{
								ARG: function__not__1__args,
								arg1: &dArg{
									arg1: &dSort{
										arg1: &dCollect{
											ARG: collect__child__node,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
											arg2: &dPredicate{
												arg1: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"obj1", "pobj1", "se", "me"},
														arg1: &dCollect{
															ARG:  collect__attributes__rel,
															arg1: &dNode{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG: collect__attributes__cat,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"pp"},
										arg1: &dCollect{
											ARG: collect__attributes__cat,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"hd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"ADP"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
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
		if test(q /* $node/../node[@rel=("hd","su","obj1","vc") and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd", "su", "obj1", "vc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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

	if test(q /* $node[@rel=("cnj","dp","mwp")] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__rel,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"cnj", "dp", "mwp"},
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}) {
		if n := find(q /* $node[@pt and ../node/@index]/ancestor::node[@cat="whq"]/node[@rel="whd" and @index = ..//node[@rel="mwp"]/@index] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__ancestors__node,
						arg1: &dFilter{
							arg1: &dVariable{
								VAR: node,
							},
							arg2: &dSort{
								arg1: &dAnd{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG: collect__attributes__index,
										arg1: &dCollect{
											ARG: collect__child__node,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"whq"},
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"whd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__index,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG: collect__attributes__index,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG: collect__descendant__or__self__type__node,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
										arg2: &dPredicate{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"mwp"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
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
		}); len(n) > 0 { // wat voor cases: attach voor to wat
			return internalHeadPosition(n, q)
		}
		if node == nLeft(find(q /* $node/../node[@rel=("cnj","dp","mwp") and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"cnj", "dp", "mwp"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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
		if test(q /* $node[@cat="pp" and node[@index and @rel="obj1" and (@cat or @pt)]] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"pp"},
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCollect{
								ARG:  collect__child__node,
								arg1: &dNode{},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dAnd{
											arg1: &dCollect{
												ARG:  collect__attributes__index,
												arg1: &dNode{},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"obj1"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
										},
										arg2: &dSort{
											arg1: &dOr{
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
												arg2: &dCollect{
													ARG:  collect__attributes__pt,
													arg1: &dNode{},
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
		}) { // coordinated prepositions with indexed obj1 GB 27/2/24
			return externalHeadPosition(node.axParent, q)
		}
		if test(q /* $node[@rel="mwp" and @index and not(@pt)] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"mwp"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__index,
									arg1: &dNode{},
								},
							},
							arg2: &dFunction{
								ARG: function__not__1__args,
								arg1: &dArg{
									arg1: &dSort{
										arg1: &dCollect{
											ARG:  collect__attributes__pt,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
				},
			},
		}) { // wat heb je voor .. where wat is extracted from mwu! TODO: ensure other mwp attach to 'wat' as fixed
			return externalHeadPosition(node.axParent, q)
		}
		if node.Rel == "cnj" {
			return headPositionOfConjunction(node, q)
		}
		return internalHeadPositionWithGapping(node.axParent, q)
	}

	// note that this gives highly counterintuitive results for 'te+AUX,COP', but fixing it gives UD validation errors GB 9-4-21
	// on the other hand, English does the same: ... are going to be bloody days, mark(days,to)
	if test(q /* $node[@rel="cmp" and ../node[@rel="body"]] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"cmp"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG:  collect__parent__type__node,
								arg1: &dNode{},
							},
							arg2: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"body"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
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
		return internalHeadPositionWithGapping(find(q /* $node/../node[@rel="body"][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"body"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if node.Rel == "--" && node.Cat != "" {
		if test(q /* $node/../node[@ud:pos=("VERB","ADJ")] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"VERB", "ADJ"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(find(q /* $node/../node[@ud:pos=("VERB","ADJ")][1] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"VERB", "ADJ"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dFunction{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}), q)
		}
		if node.Cat == "mwu" {
			if test(q /* $node/../node[@cat and not(@cat="mwu")] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
								arg2: &dFunction{
									ARG: function__not__1__args,
									arg1: &dArg{
										arg1: &dSort{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"mwu"},
													arg1: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
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
				return internalHeadPosition(find(q /* $node/../node[@cat and not(@cat="mwu")][1] */, &xPath{
					arg1: &dSort{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
							arg2: &dPredicate{
								arg1: &dPredicate{
									arg1: &dAnd{
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
										arg2: &dFunction{
											ARG: function__not__1__args,
											arg1: &dArg{
												arg1: &dSort{
													arg1: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG:  collect__attributes__cat,
															arg1: &dNode{},
														},
														arg2: &dElem{
															DATA: []interface{}{"mwu"},
															arg1: &dCollect{
																ARG:  collect__attributes__cat,
																arg1: &dNode{},
															},
														},
													},
												},
											},
										},
									},
								},
								arg2: &dFunction{
									ARG: function__first__0__args,
								},
							},
						},
					},
				}), q)
			}
			if node.Begin != node.parent.Begin { // make consistent with deplabel condition GB 9/3/21
				if node == nLeft(find(q /* $node/../node[@cat or @ud:pos = ("CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON","SYM")] */, &xPath{
					arg1: &dSort{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
							arg2: &dPredicate{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
									arg2: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"CONJ", "NOUN", "PROPN", "NUM", "ADP", "ADV", "DET", "PRON", "SYM"},
											arg1: &dCollect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &dNode{},
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
				return internalHeadPosition(find(q /* $node/../node[@cat or @ud:pos = ("CONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON","SYM")][1] */, &xPath{
					arg1: &dSort{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
							arg2: &dPredicate{
								arg1: &dPredicate{
									arg1: &dOr{
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
										arg2: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"CONJ", "NOUN", "PROPN", "NUM", "ADP", "ADV", "DET", "PRON", "SYM"},
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
											},
										},
									},
								},
								arg2: &dFunction{
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
		if n := find(q /* $node[not(@ud:pos=("VERB","ADJ"))]/../node[@rel="--" and @cat=("smain","whq","du")][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dFilter{
							arg1: &dVariable{
								VAR: node,
							},
							arg2: &dSort{
								arg1: &dFunction{
									ARG: function__not__1__args,
									arg1: &dArg{
										arg1: &dSort{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"VERB", "ADJ"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"--"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"smain", "whq", "du"},
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}); len(n) > 0 { // fixing problematic case in dpc
			return internalHeadPositionWithGapping(n, q) // why does internalHeadPositionWithGapping not work here?
		}

		if test(q, /* $node[@ud:pos = ("PUNCT","SYM","X","NOUN","PROPN","SCONJ","CCONJ","NUM","ADP","ADV","DET","PRON")
			   and ../node[@rel="--" and
			               not(@cat="mwu" or @ud:pos=("PUNCT","SYM","X","NOUN","PROPN","SCONJ","CCONJ","NUM","ADP","ADV","DET","PRON")) ]
			  ] */&xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dSort{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"PUNCT", "SYM", "X", "NOUN", "PROPN", "SCONJ", "CCONJ", "NUM", "ADP", "ADV", "DET", "PRON"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"--"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dFunction{
												ARG: function__not__1__args,
												arg1: &dArg{
													arg1: &dSort{
														arg1: &dOr{
															arg1: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__cat,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"mwu"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__cat,
																		arg1: &dNode{},
																	},
																},
															},
															arg2: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__ud_3apos,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"PUNCT", "SYM", "X", "NOUN", "PROPN", "SCONJ", "CCONJ", "NUM", "ADP", "ADV", "DET", "PRON"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &dNode{},
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
				},
			}) {
			return internalHeadPositionWithGapping(
				find(q /* ($node/../node[@rel="--" and not(@cat="mwu" or @ud:pos=("PUNCT","SYM","X","SCONJ","CCONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON"))])[1] */, &xPath{
					arg1: &dSort{
						arg1: &dFilter{
							arg1: &dSort{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dVariable{
											VAR: node,
										},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"--"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dFunction{
												ARG: function__not__1__args,
												arg1: &dArg{
													arg1: &dSort{
														arg1: &dOr{
															arg1: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__cat,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"mwu"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__cat,
																		arg1: &dNode{},
																	},
																},
															},
															arg2: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__ud_3apos,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"PUNCT", "SYM", "X", "SCONJ", "CCONJ", "NOUN", "PROPN", "NUM", "ADP", "ADV", "DET", "PRON"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &dNode{},
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
							arg2: &dSort{
								arg1: &dFunction{
									ARG: function__first__0__args,
								},
							},
						},
					},
				}),
				q)
		}
		// moved this clause to take precedence over next (cat) case for consistency with deplabel GB 22/03/31
		if test(q /* $node[@ud:pos = ("ADJ","VERB") ] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"ADJ", "VERB"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) { // deal with time out parses GB 19/03/21
			if node == nLeft(find(q /* $node/../node[@rel="--" and @ud:pos = ("ADJ","VERB") ] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"--"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"ADJ", "VERB"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &dNode{},
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
			return internalHeadPositionWithGapping(
				find(q /* ($node/../node[@rel="--" and @ud:pos=("ADJ","VERB")])[1] */, &xPath{
					arg1: &dSort{
						arg1: &dFilter{
							arg1: &dSort{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dVariable{
											VAR: node,
										},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"--"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"ADJ", "VERB"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
							},
							arg2: &dSort{
								arg1: &dFunction{
									ARG: function__first__0__args,
								},
							},
						},
					},
				}),
				q)
		}
		if n := find(q /* $node/../node[@cat and not(@cat="mwu")][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dAnd{
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
								arg2: &dFunction{
									ARG: function__not__1__args,
									arg1: &dArg{
										arg1: &dSort{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"mwu"},
													arg1: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}); len(n) > 0 {
			return internalHeadPosition(n, q)
		}

		if test(q /* $node[@ud:pos="PUNCT" and count(../node) > 1] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"PUNCT"},
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCmp{
								ARG: cmp__gt,
								arg1: &dFunction{
									ARG: function__count__1__args,
									arg1: &dArg{
										arg1: &dCollect{
											ARG: collect__child__node,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
								arg2: &dElem{
									DATA: []interface{}{1},
									arg1: &dFunction{
										ARG: function__count__1__args,
										arg1: &dArg{
											arg1: &dCollect{
												ARG: collect__child__node,
												arg1: &dCollect{
													ARG:  collect__parent__type__node,
													arg1: &dNode{},
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
		}) { // moved this case up, seems safer for PUNCT  GB 16/3/21
			if n := find(q /* $node/../node[not(@ud:pos="PUNCT")][1] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dPredicate{
								arg1: &dFunction{
									ARG: function__not__1__args,
									arg1: &dArg{
										arg1: &dSort{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"PUNCT"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
							},
							arg2: &dFunction{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}); len(n) > 0 {
				return internalHeadPosition(n, q)
			}
			if node == nLeft(find(q /* $node/../node[@rel="--" and (@cat or @pt)] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"--"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dSort{
									arg1: &dOr{
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG:  collect__attributes__pt,
											arg1: &dNode{},
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

		if test(q /* $node[@ud:pos = ("SYM","X","NOUN","PROPN","SCONJ","CCONJ","NUM","ADP","ADV","DET","PRON") ] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"SYM", "X", "NOUN", "PROPN", "SCONJ", "CCONJ", "NUM", "ADP", "ADV", "DET", "PRON"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) { // removed PUNCT here as we never want PUNCT as root or PUNCT having deps GB 16/3/21
			if node == nLeft(find(q /* $node/../node[@rel="--" and (@cat="mwu" or @ud:pos = ("SYM","X","NOUN","PROPN","SCONJ","CCONJ","NUM","ADP","ADV","DET","PRON"))] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"--"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dSort{
									arg1: &dOr{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"mwu"},
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"SYM", "X", "NOUN", "PROPN", "SCONJ", "CCONJ", "NUM", "ADP", "ADV", "DET", "PRON"},
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
											},
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
			return internalHeadPositionWithGapping(
				find(q /* ($node/../node[@rel="--" and (@cat="mwu" or @ud:pos=("SYM","X","SCONJ","CCONJ","NOUN","PROPN","NUM","ADP","ADV","DET","PRON")) ])[1] */, &xPath{
					arg1: &dSort{
						arg1: &dFilter{
							arg1: &dSort{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dVariable{
											VAR: node,
										},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"--"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dSort{
												arg1: &dOr{
													arg1: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG:  collect__attributes__cat,
															arg1: &dNode{},
														},
														arg2: &dElem{
															DATA: []interface{}{"mwu"},
															arg1: &dCollect{
																ARG:  collect__attributes__cat,
																arg1: &dNode{},
															},
														},
													},
													arg2: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &dNode{},
														},
														arg2: &dElem{
															DATA: []interface{}{"SYM", "X", "SCONJ", "CCONJ", "NOUN", "PROPN", "NUM", "ADP", "ADV", "DET", "PRON"},
															arg1: &dCollect{
																ARG:  collect__attributes__ud_3apos,
																arg1: &dNode{},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							arg2: &dSort{
								arg1: &dFunction{
									ARG: function__first__0__args,
								},
							},
						},
					},
				}),
				q)
		}

		if node.parent.Begin >= 0 {
			return externalHeadPosition(node.axParent, q)
		}
		panic("No head found")
	}

	if node.Rel == "dlink" || node.Rel == "sat" || node.Rel == "tag" {
		if n := find(q /* $node/../node[@rel="nucl"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"nucl"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}); len(n) > 0 {
			return internalHeadPositionWithGapping(n, q)
		}
		panic("No external head")
	}

	if node.Rel == "vc" && node.parent.Cat == "pp" { // eraan dat, added exception for predc GB 22/03/21
		if test(q /* $node/../node/@rel="predc" */, &xPath{
			arg1: &dSort{
				arg1: &dEqual{
					ARG: equal__is,
					arg1: &dCollect{
						ARG: collect__attributes__rel,
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
						},
					},
					arg2: &dElem{
						DATA: []interface{}{"predc"},
						arg1: &dCollect{
							ARG: collect__attributes__rel,
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPositionWithGapping(find(q /* $node/../node[@rel="predc"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"predc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	// ten ondergaan is a mwp head in extra/972, added cat to first disjunct GB 3/3/21
	if node.Rel == "vc" { // only consider vc as head if it has a head itself or is a word (rare cases where subj-index is missing, also in conjuctions!), otherwise attach orphans to subj GB 16/02/21
		// added vc node as possible embedded node as well, as hd can be hidden there as well -- GB 4/1/24
		if test(q, /* $node[node[@rel="hd" and (@pt or @cat)] or
					              node[@rel=("body","cnj","vc")]/node[@rel="hd" and @pt] or
					              node[@rel="cnj"]/node[@rel="body"]/node[@rel="hd" and @pt] or
					              node[@rel=("body","cnj") and @pt] or
					              @pt
					             ]
					        and $node/../node[@rel="hd" and
			                                   ( @ud:pos="AUX" or
			                                      $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index

			                                   )
			                                 ]
			                and not($node/../node[@rel="predc"]) */&xPath{
				arg1: &dSort{
					arg1: &dAnd{
						arg1: &dAnd{
							arg1: &dFilter{
								arg1: &dVariable{
									VAR: node,
								},
								arg2: &dSort{
									arg1: &dOr{
										arg1: &dOr{
											arg1: &dOr{
												arg1: &dOr{
													arg1: &dCollect{
														ARG:  collect__child__node,
														arg1: &dNode{},
														arg2: &dPredicate{
															arg1: &dAnd{
																arg1: &dEqual{
																	ARG: equal__is,
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																	arg2: &dElem{
																		DATA: []interface{}{"hd"},
																		arg1: &dCollect{
																			ARG:  collect__attributes__rel,
																			arg1: &dNode{},
																		},
																	},
																},
																arg2: &dSort{
																	arg1: &dOr{
																		arg1: &dCollect{
																			ARG:  collect__attributes__pt,
																			arg1: &dNode{},
																		},
																		arg2: &dCollect{
																			ARG:  collect__attributes__cat,
																			arg1: &dNode{},
																		},
																	},
																},
															},
														},
													},
													arg2: &dCollect{
														ARG: collect__child__node,
														arg1: &dCollect{
															ARG:  collect__child__node,
															arg1: &dNode{},
															arg2: &dPredicate{
																arg1: &dEqual{
																	ARG: equal__is,
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																	arg2: &dElem{
																		DATA: []interface{}{"body", "cnj", "vc"},
																		arg1: &dCollect{
																			ARG:  collect__attributes__rel,
																			arg1: &dNode{},
																		},
																	},
																},
															},
														},
														arg2: &dPredicate{
															arg1: &dAnd{
																arg1: &dEqual{
																	ARG: equal__is,
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																	arg2: &dElem{
																		DATA: []interface{}{"hd"},
																		arg1: &dCollect{
																			ARG:  collect__attributes__rel,
																			arg1: &dNode{},
																		},
																	},
																},
																arg2: &dCollect{
																	ARG:  collect__attributes__pt,
																	arg1: &dNode{},
																},
															},
														},
													},
												},
												arg2: &dCollect{
													ARG: collect__child__node,
													arg1: &dCollect{
														ARG: collect__child__node,
														arg1: &dCollect{
															ARG:  collect__child__node,
															arg1: &dNode{},
															arg2: &dPredicate{
																arg1: &dEqual{
																	ARG: equal__is,
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																	arg2: &dElem{
																		DATA: []interface{}{"cnj"},
																		arg1: &dCollect{
																			ARG:  collect__attributes__rel,
																			arg1: &dNode{},
																		},
																	},
																},
															},
														},
														arg2: &dPredicate{
															arg1: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"body"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																},
															},
														},
													},
													arg2: &dPredicate{
														arg1: &dAnd{
															arg1: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"hd"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__rel,
																		arg1: &dNode{},
																	},
																},
															},
															arg2: &dCollect{
																ARG:  collect__attributes__pt,
																arg1: &dNode{},
															},
														},
													},
												},
											},
											arg2: &dCollect{
												ARG:  collect__child__node,
												arg1: &dNode{},
												arg2: &dPredicate{
													arg1: &dAnd{
														arg1: &dEqual{
															ARG: equal__is,
															arg1: &dCollect{
																ARG:  collect__attributes__rel,
																arg1: &dNode{},
															},
															arg2: &dElem{
																DATA: []interface{}{"body", "cnj"},
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
																},
															},
														},
														arg2: &dCollect{
															ARG:  collect__attributes__pt,
															arg1: &dNode{},
														},
													},
												},
											},
										},
										arg2: &dCollect{
											ARG:  collect__attributes__pt,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"hd"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dSort{
											arg1: &dOr{
												arg1: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"AUX"},
														arg1: &dCollect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &dNode{},
														},
													},
												},
												arg2: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG: collect__attributes__index,
														arg1: &dCollect{
															ARG: collect__child__node,
															arg1: &dCollect{
																ARG: collect__descendant__or__self__type__node,
																arg1: &dCollect{
																	ARG: collect__ancestors__node,
																	arg1: &dVariable{
																		VAR: node,
																	},
																	arg2: &dPredicate{
																		arg1: &dEqual{
																			ARG: equal__is,
																			arg1: &dCollect{
																				ARG:  collect__attributes__rel,
																				arg1: &dNode{},
																			},
																			arg2: &dElem{
																				DATA: []interface{}{"top"},
																				arg1: &dCollect{
																					ARG:  collect__attributes__rel,
																					arg1: &dNode{},
																				},
																			},
																		},
																	},
																},
															},
															arg2: &dPredicate{
																arg1: &dEqual{
																	ARG: equal__is,
																	arg1: &dCollect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &dNode{},
																	},
																	arg2: &dElem{
																		DATA: []interface{}{"AUX"},
																		arg1: &dCollect{
																			ARG:  collect__attributes__ud_3apos,
																			arg1: &dNode{},
																		},
																	},
																},
															},
														},
													},
													arg2: &dCollect{
														ARG:  collect__attributes__index,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__not__1__args,
							arg1: &dArg{
								arg1: &dSort{
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG: collect__parent__type__node,
											arg1: &dVariable{
												VAR: node,
											},
										},
										arg2: &dPredicate{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"predc"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
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

		if test(q /* $node/../node[@rel=("hd","su","obj1") and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd", "su", "obj1"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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

	if node.Rel == "whd" || node.Rel == "rhd" { // hack for '[whd hoe [du/1 du [de nevel van de tijd]]'
		if test(q /* $node[../node[@rel="body"]/node[@rel="dp"]/@index = @index ] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__index,
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
										arg2: &dPredicate{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"body"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
										},
									},
									arg2: &dPredicate{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"dp"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
									},
								},
							},
							arg2: &dCollect{
								ARG:  collect__attributes__index,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		}) {
			return externalHeadPosition(node.axParent, q)
		}
		if node.Index > 0 {
			return externalHeadPosition(find(q /* ($node/../node[@rel="body"]//node[@index = $node/@index ])[1] */, &xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dSort{
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__descendant__or__self__type__node,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG: collect__parent__type__node,
											arg1: &dVariable{
												VAR: node,
											},
										},
										arg2: &dPredicate{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"body"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
								arg2: &dPredicate{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__index,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG: collect__attributes__index,
											arg1: &dVariable{
												VAR: node,
											},
										},
									},
								},
							},
						},
						arg2: &dSort{
							arg1: &dFunction{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}), q)
		}
		return internalHeadPosition(find(q /* $node/../node[@rel="body"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"body"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
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
		//NP -> OK, followingCnjSister geeft geen panic
		tmp := followingCnjSister(node, q)
		// coordination of gapped PPs -- attach cc to 2nd prep GB 4/3/24
		if test(q /* $node/../node[@rel="cnj" and @cat="pp"]/node[@rel="obj1" and @index and not(@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"cnj"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"pp"},
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"obj1"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__index,
									arg1: &dNode{},
								},
							},
							arg2: &dFunction{
								ARG: function__not__1__args,
								arg1: &dArg{
									arg1: &dSort{
										arg1: &dOr{
											arg1: &dCollect{
												ARG:  collect__attributes__pt,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
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
			return internalHeadPositionWithGapping(ifZ(find(q, /* $node/../node[@rel="cnj" and
				   @begin=$tmp/@begin and
				   @end=$tmp/@end
				  ]/node[@rel="hd"] */&xPath{
					arg1: &dSort{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"cnj"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__begin,
													arg1: &dNode{},
												},
												arg2: &dCollect{
													ARG: collect__attributes__begin,
													arg1: &dVariable{
														VAR: tmp,
													},
												},
											},
										},
										arg2: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__end,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG: collect__attributes__end,
												arg1: &dVariable{
													VAR: tmp,
												},
											},
										},
									},
								},
							},
							arg2: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"hd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
				})), q)
		}
		return internalHeadPositionWithGapping(ifZ(find(q, /* $node/../node[@rel="cnj" and
			   @begin=$tmp/@begin and
			   @end=$tmp/@end
			  ] */&xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dAnd{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"cnj"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
									arg2: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__begin,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG: collect__attributes__begin,
											arg1: &dVariable{
												VAR: tmp,
											},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__end,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG: collect__attributes__end,
										arg1: &dVariable{
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

	if node.Rel == "su" { // similar condition for empty aux node as for vc case -- GB 21/2/24
		if test(q, /* $node[../node[@rel="predc" and (@pt or @cat)] and ../node[@rel="hd" and
			   ( @ud:pos="AUX" or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index ) ]] */&xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dSort{
							arg1: &dAnd{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"predc"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dSort{
												arg1: &dOr{
													arg1: &dCollect{
														ARG:  collect__attributes__pt,
														arg1: &dNode{},
													},
													arg2: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
									arg2: &dPredicate{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"hd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dSort{
												arg1: &dOr{
													arg1: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &dNode{},
														},
														arg2: &dElem{
															DATA: []interface{}{"AUX"},
															arg1: &dCollect{
																ARG:  collect__attributes__ud_3apos,
																arg1: &dNode{},
															},
														},
													},
													arg2: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG: collect__attributes__index,
															arg1: &dCollect{
																ARG: collect__child__node,
																arg1: &dCollect{
																	ARG: collect__descendant__or__self__type__node,
																	arg1: &dCollect{
																		ARG: collect__ancestors__node,
																		arg1: &dVariable{
																			VAR: node,
																		},
																		arg2: &dPredicate{
																			arg1: &dEqual{
																				ARG: equal__is,
																				arg1: &dCollect{
																					ARG:  collect__attributes__rel,
																					arg1: &dNode{},
																				},
																				arg2: &dElem{
																					DATA: []interface{}{"top"},
																					arg1: &dCollect{
																						ARG:  collect__attributes__rel,
																						arg1: &dNode{},
																					},
																				},
																			},
																		},
																	},
																},
																arg2: &dPredicate{
																	arg1: &dEqual{
																		ARG: equal__is,
																		arg1: &dCollect{
																			ARG:  collect__attributes__ud_3apos,
																			arg1: &dNode{},
																		},
																		arg2: &dElem{
																			DATA: []interface{}{"AUX"},
																			arg1: &dCollect{
																				ARG:  collect__attributes__ud_3apos,
																				arg1: &dNode{},
																			},
																		},
																	},
																},
															},
														},
														arg2: &dCollect{
															ARG:  collect__attributes__index,
															arg1: &dNode{},
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
			return internalHeadPositionWithGapping(find(q /* $node/../node[@rel="predc"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"predc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if test(q, /* $node/../node[@rel="vc"] and $node/../node[@rel="hd" and
			   ( @ud:pos="AUX" or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index ) ] */&xPath{
				arg1: &dSort{
					arg1: &dAnd{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
							arg2: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"vc"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
							},
						},
						arg2: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
							arg2: &dPredicate{
								arg1: &dAnd{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"hd"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
									arg2: &dSort{
										arg1: &dOr{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"AUX"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG: collect__attributes__index,
													arg1: &dCollect{
														ARG: collect__child__node,
														arg1: &dCollect{
															ARG: collect__descendant__or__self__type__node,
															arg1: &dCollect{
																ARG: collect__ancestors__node,
																arg1: &dVariable{
																	VAR: node,
																},
																arg2: &dPredicate{
																	arg1: &dEqual{
																		ARG: equal__is,
																		arg1: &dCollect{
																			ARG:  collect__attributes__rel,
																			arg1: &dNode{},
																		},
																		arg2: &dElem{
																			DATA: []interface{}{"top"},
																			arg1: &dCollect{
																				ARG:  collect__attributes__rel,
																				arg1: &dNode{},
																			},
																		},
																	},
																},
															},
														},
														arg2: &dPredicate{
															arg1: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__ud_3apos,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"AUX"},
																	arg1: &dCollect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &dNode{},
																	},
																},
															},
														},
													},
												},
												arg2: &dCollect{
													ARG:  collect__attributes__index,
													arg1: &dNode{},
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
			//NP -> half opgelost -> zie TODO
			/* tmp, err := internalHeadPositionWithGappingWithError(node.axParent, q) // testing -- dont go to vc as it has no head sometimes...
			if err == nil && node.Begin < tmp && tmp <= node.End {                 // maybe the different error handling in go code causes diff with xquery script?
				return externalHeadPosition(node.axParent, q)
			}
			*/
			if internalHeadPositionWithGapping(node.axParent, q) == internalHeadPositionWithGapping(find(q /* $node */, &xPath{
				arg1: &dSort{
					arg1: &dVariable{
						VAR: node,
					},
				},
			}), q) {
				return externalHeadPosition(node.axParent, q)
			}
			// TODO: dit is gelijk aan tmp... wat als err != nil?
			return internalHeadPositionWithGapping(node.axParent, q) // dont go to vc directly as it might be empty
		}
		if test(q /* $node/../node[@rel="hd" and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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
		// switched off (XXX) for now, as it gives wrong results in [su [predc _ mod]] cases, there, su should be head, and mod an orphan
		// GB 8/1/24
		if test(q /* $node[../node[@rel="XXXXpredc" and (@pt or @cat)] and not(../node[@rel="hd" and (@pt or @cat)])] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"XXXXpredc"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dSort{
											arg1: &dOr{
												arg1: &dCollect{
													ARG:  collect__attributes__pt,
													arg1: &dNode{},
												},
												arg2: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
											},
										},
									},
								},
							},
							arg2: &dFunction{
								ARG: function__not__1__args,
								arg1: &dArg{
									arg1: &dSort{
										arg1: &dCollect{
											ARG: collect__child__node,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
											arg2: &dPredicate{
												arg1: &dAnd{
													arg1: &dEqual{
														ARG: equal__is,
														arg1: &dCollect{
															ARG:  collect__attributes__rel,
															arg1: &dNode{},
														},
														arg2: &dElem{
															DATA: []interface{}{"hd"},
															arg1: &dCollect{
																ARG:  collect__attributes__rel,
																arg1: &dNode{},
															},
														},
													},
													arg2: &dSort{
														arg1: &dOr{
															arg1: &dCollect{
																ARG:  collect__attributes__pt,
																arg1: &dNode{},
															},
															arg2: &dCollect{
																ARG:  collect__attributes__cat,
																arg1: &dNode{},
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
			return internalHeadPositionWithGapping(find(q /* $node/../node[@rel="predc"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"predc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
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
		if test(q /* $node/../node[@rel=("hd","su") and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd", "su"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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

	if node.Rel == "pc" || node.Rel == "ld" {
		if test(q /* $node/../node[@rel=("hd","su","obj1") and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd", "su", "obj1"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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
		if n := find(q /* $node/../node[@rel=("modERR","appERR","pcERR") and (@cat or @pt)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"modERR", "appERR", "pcERR"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}); len(n) > 0 { // whatever comes first 4/2/21: added pc GB
			if node == nLeft(n) { // gapping with multiple mods
				return externalHeadPosition(node.axParent, q)
			}
			return internalHeadPositionWithGapping(node.axParent, q)
		}
		return externalHeadPosition(node.axParent, q)
	}

	if node.Rel == "mod" || node.Rel == "app" {

		if test(q /* $node/../node[( @rel=("su","obj1","predc","body","pc", "cnj") or (@rel="hd" and not(@ud:pos="ADP"))) and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dSort{
								arg1: &dOr{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"su", "obj1", "predc", "body", "pc", "cnj"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
									arg2: &dSort{
										arg1: &dAnd{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"hd"},
													arg1: &dCollect{
														ARG:  collect__attributes__rel,
														arg1: &dNode{},
													},
												},
											},
											arg2: &dFunction{
												ARG: function__not__1__args,
												arg1: &dArg{
													arg1: &dSort{
														arg1: &dEqual{
															ARG: equal__is,
															arg1: &dCollect{
																ARG:  collect__attributes__ud_3apos,
																arg1: &dNode{},
															},
															arg2: &dElem{
																DATA: []interface{}{"ADP"},
																arg1: &dCollect{
																	ARG:  collect__attributes__ud_3apos,
																	arg1: &dNode{},
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
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) { // added pc 16/2/21, added cnj 22/03/21
			// gapping, as su but now su or obj1  could be head as well
			return internalHeadPositionWithGapping(node.axParent, q)
		}

		if test(q /* $node/../node[@rel="hd" and @ud:pos="ADP"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"ADP"},
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) { // vlak voor [index obj1] en na de winter  GB 4/3/21
			return internalHeadPosition(find(q /* $node/../node[@rel="hd"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}

		// if mod in otherwise empty predc --> externalheadposition GB 21/2/24
		if test(q /* $node/../@rel="predc" */, &xPath{
			arg1: &dSort{
				arg1: &dEqual{
					ARG: equal__is,
					arg1: &dCollect{
						ARG: collect__attributes__rel,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: node,
							},
						},
					},
					arg2: &dElem{
						DATA: []interface{}{"predc"},
						arg1: &dCollect{
							ARG: collect__attributes__rel,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
						},
					},
				},
			},
		}) { // a mod in an otherwise empty predc
			return externalHeadPosition(node.axParent, q)
		}
		// ordered case below before multiple mod cases, to deal with su [predc _ mod] cases GB 8/1/24
		if test(q /* $node[not(../@rel="su") and ../../node[@rel="su" and (@pt or @cat)]] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dFunction{
								ARG: function__not__1__args,
								arg1: &dArg{
									arg1: &dSort{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG: collect__attributes__rel,
												arg1: &dCollect{
													ARG:  collect__parent__type__node,
													arg1: &dNode{},
												},
											},
											arg2: &dElem{
												DATA: []interface{}{"su"},
												arg1: &dCollect{
													ARG: collect__attributes__rel,
													arg1: &dCollect{
														ARG:  collect__parent__type__node,
														arg1: &dNode{},
													},
												},
											},
										},
									},
								},
							},
							arg2: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"su"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dSort{
											arg1: &dOr{
												arg1: &dCollect{
													ARG:  collect__attributes__pt,
													arg1: &dNode{},
												},
												arg2: &dCollect{
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
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
		}) { // a mod in an otherwise empty tree (after fixing heads in conj)
			return internalHeadPosition(find(q /* $node/../../node[@rel="su"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"su"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if test(q /* $node/../../@cat="ti" and $node/../../../node[@rel="su" and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dAnd{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG: collect__attributes__cat,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
							},
						},
						arg2: &dElem{
							DATA: []interface{}{"ti"},
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dVariable{
											VAR: node,
										},
									},
								},
							},
						},
					},
					arg2: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
							},
						},
						arg2: &dPredicate{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"su"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dSort{
									arg1: &dOr{
										arg1: &dCollect{
											ARG:  collect__attributes__pt,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
				},
			},
		}) { // an mod in a te-inf VP with a subject GB 20/2/24
			return internalHeadPosition(find(q /* $node/../../../node[@rel="su"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"su"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
		if test(q /* $node[../@cat="ssub" and ../@rel="body" and ../../node[@rel=("whd","rhd") and @pt]] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG: collect__attributes__cat,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"ssub"},
										arg1: &dCollect{
											ARG: collect__attributes__cat,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG: collect__attributes__rel,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"body"},
										arg1: &dCollect{
											ARG: collect__attributes__rel,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
									},
								},
							},
							arg2: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dCollect{
										ARG:  collect__parent__type__node,
										arg1: &dNode{},
									},
								},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"whd", "rhd"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dCollect{
											ARG:  collect__attributes__pt,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(find(q /* $node/../../node[@rel=("whd","rhd")] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"whd", "rhd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q) // wat cool/mogelijk is en wat niet, GB 24/1/24 added rhd GB 26/2/24
		}
		if n := find(q /* $node/../node[@rel=("mod","app") and descendant-or-self::node/@pt] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"mod", "app"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCollect{
								ARG: collect__attributes__pt,
								arg1: &dCollect{
									ARG:  collect__descendant__or__self__node,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}); len(n) > 0 { // whatever comes first 4/2/21: added pc GB
			if node == nLeft(n) { // gapping with multiple mods
				return externalHeadPosition(node.axParent, q)
			}
			return internalHeadPositionWithGapping(node.axParent, q)
		}

		return externalHeadPosition(node.axParent, q) /* an empty mod in an otherwise empty tree
		   -- mod is co-indexed with rhd, rest is elided,
		   LassySmall4/wiki-7064/wiki-7064.p.28.s.3.xml */
	}

	if node.Rel == "app" || node.Rel == "det" || node.Rel == "me" {
		if test(q /* $node/../node[@rel=("hd","mod") and (@pt or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd", "mod"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) { // gapping with an app (or a det)! (or me!)
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

func internalHeadPositionWithError(nodes []interface{}, q *context) (head int, err error) {
	defer func() {
		if r := recover(); r != nil {
			head = error_no_head
			err = fmt.Errorf("NO HEAD")
		}
	}()
	head = internalHeadPosition(nodes, q)
	return // geen argumenten i.v.m. recover
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

	if test(q /* $node[@cat="pp"] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__cat,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"pp"},
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}) {
		// if ($node/node[@rel="hd" and @pt=("bw","n")] )  ( n --> TEMPORARY HACK to fix error where NP is erroneously tagged as PP )
		// then $node/node[@rel="hd"]/@end
		if n := find(q /* $node/node[@rel=("obj1","pobj1","se","vc")][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"obj1", "pobj1", "se", "vc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}); len(n) > 0 {
			return internalHeadPosition(n, q)
		}
		if n := find(q /* $node/node[@rel="hd"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"hd"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
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
		return internalHeadPosition(find(q /* $node/node[1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if test(q /* $node[@cat="mwuERROR"] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__cat,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"mwuERROR"},
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}) {
		// TODO: CHECK THIS
		if f := find(q /* $node/node[@rel="mwp" and @pt] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"mwp"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCollect{
								ARG:  collect__attributes__pt,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		}); len(f) > 0 {
			return internalHeadPosition(find(q /* $node/node[@rel="mwp"][1] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dPredicate{
							arg1: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"mwp"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dFunction{
								ARG: function__first__0__args,
							},
						},
					},
				},
			}), q) //   f[0].(int)
		}
		return i1(find(q /* $node/@end */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__attributes__end,
					arg1: &dVariable{
						VAR: node,
					},
				},
			},
		}))
	}

	if test(q /* $node[@cat="mwu"]/node */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"mwu"},
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPosition(find(q /* $node/node[1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if test(q /* $node[@cat="conj"] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__cat,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"conj"},
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPosition(ifLeft(find(q /* $node/node[@rel="cnj"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"cnj"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		})), q)
	}

	if predc := find(q /* $node/node[@rel="predc"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__rel,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"predc"},
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}); len(predc) > 0 {
		if test(q /* $node/node[@rel="hd" and @ud:pos="AUX"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"AUX"},
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
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
		hd := find(q /* $node/node[@rel="hd"][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}) // make robust in CGN cases GB 19/03/21
		if len(hd) == 0 { // cases where copula is missing by accident (ungrammatical, not gapping)
			return internalHeadPosition(predc, q)
		}
		return internalHeadPosition(hd, q)
	}

	if test(q, /* $node[node[@rel="vc"] and
		   node[@rel="hd" and
		        ( @ud:pos="AUX" or
		          $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index
		         )
		       ]
		  ] */&xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dAnd{
							arg1: &dCollect{
								ARG:  collect__child__node,
								arg1: &dNode{},
								arg2: &dPredicate{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"vc"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
								},
							},
							arg2: &dCollect{
								ARG:  collect__child__node,
								arg1: &dNode{},
								arg2: &dPredicate{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"hd"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dSort{
											arg1: &dOr{
												arg1: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"AUX"},
														arg1: &dCollect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &dNode{},
														},
													},
												},
												arg2: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG: collect__attributes__index,
														arg1: &dCollect{
															ARG: collect__child__node,
															arg1: &dCollect{
																ARG: collect__descendant__or__self__type__node,
																arg1: &dCollect{
																	ARG: collect__ancestors__node,
																	arg1: &dVariable{
																		VAR: node,
																	},
																	arg2: &dPredicate{
																		arg1: &dEqual{
																			ARG: equal__is,
																			arg1: &dCollect{
																				ARG:  collect__attributes__rel,
																				arg1: &dNode{},
																			},
																			arg2: &dElem{
																				DATA: []interface{}{"top"},
																				arg1: &dCollect{
																					ARG:  collect__attributes__rel,
																					arg1: &dNode{},
																				},
																			},
																		},
																	},
																},
															},
															arg2: &dPredicate{
																arg1: &dEqual{
																	ARG: equal__is,
																	arg1: &dCollect{
																		ARG:  collect__attributes__ud_3apos,
																		arg1: &dNode{},
																	},
																	arg2: &dElem{
																		DATA: []interface{}{"AUX"},
																		arg1: &dCollect{
																			ARG:  collect__attributes__ud_3apos,
																			arg1: &dNode{},
																		},
																	},
																},
															},
														},
													},
													arg2: &dCollect{
														ARG:  collect__attributes__index,
														arg1: &dNode{},
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
		return internalHeadPosition(find(q /* $node/node[@rel="vc"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"vc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	if n := find(q /* $node/node[@rel="hd"][1] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"hd"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
					arg2: &dFunction{
						ARG: function__first__0__args,
					},
				},
			},
		},
	}); len(n) > 0 {
		return internalHeadPosition(n, q)
	}

	if n := find(q /* $node/node[@rel="body"][1] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"body"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
					arg2: &dFunction{
						ARG: function__first__0__args,
					},
				},
			},
		},
	}); len(n) > 0 {
		return internalHeadPosition(n, q)
	}

	if n := find(q /* $node/node[@rel="dp"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__rel,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"dp"},
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
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

	if n := find(q /* $node/node[@rel="nucl"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__rel,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"nucl"},
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}); len(n) > 0 {
		return internalHeadPosition(if1(n), q)
	}

	if n := find(q /* $node/node[@cat="du"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__cat,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"du"},
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}); len(n) > 0 { // is this neccesary at all? , only one referring to cat, and causes problems if applied before @rel=dp case...
		return internalHeadPosition(if1(n), q)
	}

	if test(q /* $node[@word] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dCollect{
						ARG:  collect__attributes__word,
						arg1: &dNode{},
					},
				},
			},
		},
	}) {
		return i1(find(q /* $node/@end */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__attributes__end,
					arg1: &dVariable{
						VAR: node,
					},
				},
			},
		}))
	}

	/*
		distinguish empty nodes due to gapping/RNR from nonlocal dependencies
		use 'empty head' as string to signal presence of gapping/RNR
	*/
	if test(q /* $node[@index and not(@word or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dAnd{
						arg1: &dCollect{
							ARG:  collect__attributes__index,
							arg1: &dNode{},
						},
						arg2: &dFunction{
							ARG: function__not__1__args,
							arg1: &dArg{
								arg1: &dSort{
									arg1: &dOr{
										arg1: &dCollect{
											ARG:  collect__attributes__word,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
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
		if test(q /* $node/ancestor::node[@rel="top"]//node[@rel=("whd","rhd") and @index = $node/@index and (@word or @cat)] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__descendant__or__self__type__node,
						arg1: &dCollect{
							ARG: collect__ancestors__node,
							arg1: &dVariable{
								VAR: node,
							},
							arg2: &dPredicate{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"top"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dAnd{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"whd", "rhd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__index,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG: collect__attributes__index,
										arg1: &dVariable{
											VAR: node,
										},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__word,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return internalHeadPosition(
				find(q /* $node/ancestor::node[@rel="top"]//node[@index = $node/@index and (@word or @cat)] */, &xPath{
					arg1: &dSort{
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__descendant__or__self__type__node,
								arg1: &dCollect{
									ARG: collect__ancestors__node,
									arg1: &dVariable{
										VAR: node,
									},
									arg2: &dPredicate{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"top"},
												arg1: &dCollect{
													ARG:  collect__attributes__rel,
													arg1: &dNode{},
												},
											},
										},
									},
								},
							},
							arg2: &dPredicate{
								arg1: &dAnd{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__index,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG: collect__attributes__index,
											arg1: &dVariable{
												VAR: node,
											},
										},
									},
									arg2: &dSort{
										arg1: &dOr{
											arg1: &dCollect{
												ARG:  collect__attributes__word,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
					},
				}),
				q)
		}
		return empty_head
	}

	panic("No internal head")
}

func internalHeadPositionWithGappingWithError(node []interface{}, q *context) (head int, err error) {
	defer func() {
		if r := recover(); r != nil {
			head = error_no_head
			err = fmt.Errorf("NO HEAD")
		}
	}()
	head = internalHeadPositionWithGapping(node, q)
	return // geen argumenten i.v.m. recover
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

	//NP -> OK, werkt zoals het zou moeten
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

	if test(q /* $node[not(@cat="pp")]/node[@rel="hd" and (@pt or @cat) and not(@ud:pos=("AUX","ADP"))] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dFunction{
							ARG: function__not__1__args,
							arg1: &dArg{
								arg1: &dSort{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"pp"},
											arg1: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
					},
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dSort{
								arg1: &dOr{
									arg1: &dCollect{
										ARG:  collect__attributes__pt,
										arg1: &dNode{},
									},
									arg2: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__not__1__args,
							arg1: &dArg{
								arg1: &dSort{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"AUX", "ADP"},
											arg1: &dCollect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &dNode{},
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
		return internalHeadPositionWithGapping(find(q /* $node/node[@rel="hd"][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q) // aux, prepositions   added [1] to make robust for CGN cases GB 19/03/21
	}

	if test(q /* $node/node[@rel="hd" and @ud:pos="AUX"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"hd"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"AUX"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) { // (@pt or @cat or node[@cat or @pt])
		if test(q /* $node/node[@rel=("vc","predc") and descendant-or-self::node/@pt ] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dAnd{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"vc", "predc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCollect{
								ARG: collect__attributes__pt,
								arg1: &dCollect{
									ARG:  collect__descendant__or__self__node,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) { // what does this final node mean ?? added cat instead GB 26/2/21
			return internalHeadPositionWithGapping(find(q /* $node/node[@rel=("vc","predc")] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"vc", "predc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q) // testing should be vc,predc , fixed typo pred -> predc GB 26/2/21
		} else {
			return internalHeadPositionWithGapping(find(q /* $node/node[@rel="hd"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"hd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}), q)
		}
	}

	if test(q /* $node/node[@rel="su" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"su"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(find(q /* $node/node[@rel="su"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"su"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}), q) // 44 van 87 in lassysmall
	}

	if test(q /* $node[@rel="vc" and ../node[@rel="su" and (@pt or @cat)]] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"vc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG:  collect__parent__type__node,
								arg1: &dNode{},
							},
							arg2: &dPredicate{
								arg1: &dAnd{
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"su"},
											arg1: &dCollect{
												ARG:  collect__attributes__rel,
												arg1: &dNode{},
											},
										},
									},
									arg2: &dSort{
										arg1: &dOr{
											arg1: &dCollect{
												ARG:  collect__attributes__pt,
												arg1: &dNode{},
											},
											arg2: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
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
		return internalHeadPositionWithGapping(find(q /* $node/../node[@rel="su"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dCollect{
						ARG: collect__parent__type__node,
						arg1: &dVariable{
							VAR: node,
						},
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"su"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	if test(q /* $node/node[@rel="obj1" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"obj1"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(find(q /* $node/node[@rel="obj1"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"obj1"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	if test(q /* $node/node[@rel="predc" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"predc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(find(q /* $node/node[@rel="predc"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"predc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	if test(q /* $node/node[@rel="vc" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"vc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) {
		return internalHeadPositionWithGapping(find(q /* $node/node[@rel="vc"][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"vc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if test(q /* $node/node[@rel="pc" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"pc"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) { // commenting out this case and added to case below GB 5/2/21 and put back in 8/2
		return internalHeadPositionWithGapping(find(q /* $node/node[@rel="pc"][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"pc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	if test(q /* $node/node[@rel="ld" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"ld"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}) { // fixing EANS case
		return internalHeadPositionWithGapping(find(q /* $node/node[@rel="ld"][1] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dPredicate{
						arg1: &dPredicate{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"ld"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
						arg2: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		}), q)
	}

	// moved the next case up to take precendence over mod rule that follows for consistency with deplabel GB 3/3/21
	if n := find(q /* $node/node[@rel="hd" and @ud:pos="ADP"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"hd"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"ADP"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}); len(n) > 0 { // in en rond Brussel, case not necessary in xquery code (run-time issue?)
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := find(q /* $node[@cat="pp"]/node[@rel="hd" and @cat="mwu"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"pp"},
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"hd"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"mwu"},
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}); len(n) > 0 { // zowel voorafgaand als na afloop van X
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := find(q /* $node/node[@rel=("mod","app","me") and descendant-or-self::node/@pt] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"mod", "app", "me"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dCollect{
							ARG: collect__attributes__pt,
							arg1: &dCollect{
								ARG:  collect__descendant__or__self__node,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}); len(n) > 0 { // pick leftmost
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := find(q /* $node/node[@rel="det" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"det"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
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

	if n := find(q /* $node/node[@rel=("body","nucl") and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"body", "nucl"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}); len(n) > 0 { // added nucl GB 4/3/21
		return internalHeadPositionWithGapping(if1(n), q)
	}

	if n := find(q /* $node/node[@rel="cnj" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"cnj"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
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

	if n := find(q /* $node/node[@rel="dp" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"dp"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
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

	if n := find(q /* $node[@cat="mwu"]/node[@rel="mwp" and (@pt or @cat)] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"mwu"},
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
				arg2: &dPredicate{
					arg1: &dAnd{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"mwp"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dSort{
							arg1: &dOr{
								arg1: &dCollect{
									ARG:  collect__attributes__pt,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		},
	}); len(n) > 0 { // zowel voorafgaand als na afloop van X
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

	//NP wat te doen?
	internal_head := internalHeadPositionWithGapping([]interface{}{node}, q)
	leftmost_conj_daughter := nLeft(find(q /* $node/../node[@rel="cnj"] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dCollect{
					ARG: collect__parent__type__node,
					arg1: &dVariable{
						VAR: node,
					},
				},
				arg2: &dPredicate{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG:  collect__attributes__rel,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"cnj"},
							arg1: &dCollect{
								ARG:  collect__attributes__rel,
								arg1: &dNode{},
							},
						},
					},
				},
			},
		},
	}))
	//NP wat te doen?
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
		// return leftmost_conj_daughter.Node[0].End // this should not happen really -- give error msg?
		// happens in elliptical (right-node-raising) cases like cdb/4863,
		// door zaterdag met 7-9 ... en zondag met 11-1... te winnen , [zaterdag met 11-7 ... te winnen] is seen as single constituent, [zondag met ...] as elliptical conj
		return leftmost_internal_head // accept that conj points in the wrong direction in such cases
	}
	sort.Ints(endpos_of_leftmost_conj_constituents)
	return endpos_of_leftmost_conj_constituents[len(endpos_of_leftmost_conj_constituents)-1]
}

func followingCnjSister(node *nodeType, q *context) []interface{} {

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
			b := find(q /* $n/descendant-or-self::node[@word]/@begin */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__attributes__begin,
						arg1: &dCollect{
							ARG: collect__descendant__or__self__node,
							arg1: &dVariable{
								VAR: n,
							},
							arg2: &dPredicate{
								arg1: &dCollect{
									ARG:  collect__attributes__word,
									arg1: &dNode{},
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
