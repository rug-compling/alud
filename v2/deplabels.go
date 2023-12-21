//
// // THIS IS A GENERATED FILE. DO NOT EDIT.
//

package alud

// recursive
func dependencyLabel(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "dependencyLabel", q, node))
		}
	}()

	depthCheck(q)

	if node.parent.Cat == "top" && node.parent.End == 1000 {
		return "root"
	}

	if node.Rel == "app" {
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
		}) {
			return "appos"
		}
		if test(q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"mod"},
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
			return "appos" //was "orphan"
		}
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "cmp" {
		return "mark"
	}
	if node.Rel == "crd" {
		if node.Lemma == "zowel" && test(q /* $node/../node[@rel="crd" and @lemma="als"] */, &xPath{
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
									DATA: []interface{}{"crd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"als"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return "cc:preconj"
		}
		if node.Lemma == "niet" && test(q /* $node/../node[@rel="crd" and @lemma="maar"] */, &xPath{
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
									DATA: []interface{}{"crd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"maar"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return "cc:preconj"
		}
		if node.Lemma == "respectievelijk" && test(q /* $node/../node[@rel="crd" and @lemma="en"] */, &xPath{
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
									DATA: []interface{}{"crd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"en"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return "cc:preconj"
		}
		if node.Lemma == "ver" && test(q /* $node/../node[@rel="crd" and @lemma="en"] */, &xPath{
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
									DATA: []interface{}{"crd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"en"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) { // verder A, B, en C
			return "cc:preconj"
		}
		if node.Lemma == "achtereenvolgens" && test(q /* $node/../node[@rel="crd" and @lemma="en"] */, &xPath{
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
									DATA: []interface{}{"crd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"en"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) { // [PP door [NP achtereenvolgens A en B]]
			return "cc:preconj"
		}
		if node.Lemma == "van" && test(q /* $node/../node[@rel="crd" and @lemma="tot"] */, &xPath{
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
									DATA: []interface{}{"crd"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"tot"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return "cc:preconj"
		}
		if node.Lemma == "noch" && test(q /* $node/../node[@rel="crd" and @lemma="noch" and $node/@begin < @begin ] */, &xPath{
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
										DATA: []interface{}{"crd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"noch"},
										arg1: &dCollect{
											ARG:  collect__attributes__lemma,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dCmp{
								ARG: cmp__lt,
								arg1: &dCollect{
									ARG: collect__attributes__begin,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__begin,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) {
			return "cc:preconj"
		}
		if node.Lemma == "eerder" && test(q /* $node/../node[@rel="crd" and @lemma="dan" and $node/@begin < @begin ] */, &xPath{
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
										DATA: []interface{}{"crd"},
										arg1: &dCollect{
											ARG:  collect__attributes__rel,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"dan"},
										arg1: &dCollect{
											ARG:  collect__attributes__lemma,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dCmp{
								ARG: cmp__lt,
								arg1: &dCollect{
									ARG: collect__attributes__begin,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dCollect{
									ARG:  collect__attributes__begin,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) {
			return "cc:preconj"
		}
		return "cc"
	}
	if node.Rel == "me" && test(q /* $node[@rel="me" and ../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="ADP")]] */, &xPath{
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
								DATA: []interface{}{"me"},
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
		},
	}) {
		return determineNominalModLabel(node, q)
	}
	if node.Rel == "obcomp" {
		return "advcl"
	}
	if node.Rel == "obj2" { // added case for coordination of aan-PP -- GB 2/3/23
		if node.Cat == "pp" || test(q /* $node[@cat="conj" and node[@cat="pp"]/node[@lemma="aan"]] */, &xPath{
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
									DATA: []interface{}{"conj"},
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__lemma,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"aan"},
											arg1: &dCollect{
												ARG:  collect__attributes__lemma,
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
			return "obl"
		}
		return "iobj"
	}
	if node.Rel == "pobj1" {
		return "expl"
	}
	if node.Rel == "predc" {
		if test(q /* $node/../node[@rel=("obj1","se","vc") and (@pt or @cat)] or $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX")] */, &xPath{
			arg1: &dSort{
				arg1: &dOr{
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
										DATA: []interface{}{"obj1", "se", "vc"},
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
		}) {
			if test(q /* $node/../@cat="pp" */, &xPath{
				arg1: &dSort{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG: collect__attributes__cat,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
						},
						arg2: &dElem{
							DATA: []interface{}{"pp"},
							arg1: &dCollect{
								ARG: collect__attributes__cat,
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
			}) { // check for absolutive (met) constructions, https://github.com/UniversalDependencies/docs/issues/408
				if test(q /* $node/../../@cat="np" */, &xPath{
					arg1: &dSort{
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
								DATA: []interface{}{"np"},
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
					},
				}) {
					return "acl"
				}
				if test(q /* $node/../../@cat=("top","du") */, &xPath{
					arg1: &dSort{
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
								DATA: []interface{}{"top", "du"},
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
					},
				}) { // met als gevolg een neiging tot...  added du GB 9/3/21
					return dependencyLabel(node.parent, q) //  replaced "root" with this to account for dp-dp cases, where met-PP is second element (hoe doe je dat met sokken aan) GB 18/03/21
				}
				if test(q /* $node[../../../@cat=("top","du") and ../@rel="body"] */, &xPath{
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
											ARG: collect__attributes__cat,
											arg1: &dCollect{
												ARG: collect__parent__type__node,
												arg1: &dCollect{
													ARG: collect__parent__type__node,
													arg1: &dCollect{
														ARG:  collect__parent__type__node,
														arg1: &dNode{},
													},
												},
											},
										},
										arg2: &dElem{
											DATA: []interface{}{"top", "du"},
											arg1: &dCollect{
												ARG: collect__attributes__cat,
												arg1: &dCollect{
													ARG: collect__parent__type__node,
													arg1: &dCollect{
														ARG: collect__parent__type__node,
														arg1: &dCollect{
															ARG:  collect__parent__type__node,
															arg1: &dNode{},
														},
													},
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
							},
						},
					},
				}) { // [cmp alleen dan] [body met een langer 75 mm kanon ], in automatic output GB 22/03/21
					return dependencyLabel(node.parent, q) //
				}
				return "advcl"
			}
			return "xcomp"
		}
		return dependencyLabel(node.parent, q) // covers gapping cases where predc is promoted to head as well
		/*
		   hack for now: de keuze is gauw gemaakt
		   was amod, is this more accurate??
		   examples of secondary predicates under xcomp suggests so

		*/
	}
	if node.Rel == "se" {
		return "expl:pv"
	}
	if node.Rel == "su" {
		if test(q /* $node[../@rel=("cnj","dp","body","nucl") and ../node[@rel="hd" and not(@pt or @cat)] and not(../node[@rel=("vc","predc") and (@pt or node[@rel=("hd","cnj")  and (@pt or @cat)] )] )] */, &xPath{
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
										ARG: collect__attributes__rel,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"cnj", "dp", "body", "nucl"},
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
															DATA: []interface{}{"vc", "predc"},
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
																				DATA: []interface{}{"hd", "cnj"},
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
							},
						},
					},
				},
			},
		}) { // gapping
			return dependencyLabel(node.parent, q) // added nucl GB 4/3/21
		}
		if test(q, /* $node[../@rel="vc" and ../node[@rel="hd" and not(@pt or @cat)]
			   and ../parent::node[@rel="cnj"]] */&xPath{
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
											ARG: collect__attributes__rel,
											arg1: &dCollect{
												ARG:  collect__parent__type__node,
												arg1: &dNode{},
											},
										},
										arg2: &dElem{
											DATA: []interface{}{"vc"},
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
								arg2: &dCollect{
									ARG: collect__parent__node,
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
						},
					},
				},
			}) { // gapping with subj downstairs
			// TODO: ../.. is veranderd in ../parent::node , is dat juist? JA
			/*
			   In 1909 werd de persoonlijke dienstplicht ingevoerd en in 1913 de algemene persoonlijke dienstplicht .
			   [ hd_i su_j vc [ hd_k [_j pers dienstplicht ]
			*/
			return dependencyLabel(node.parent.parent, q)
		}
		return subjectLabel(node, q)
	}
	if node.Rel == "sup" {
		return "expl"
	}
	if node.Rel == "svp" {
		return "compound:prt" // v2: added prt extension
	}
	if aux, err := auxiliary1(node, q); err == nil {
		if aux == "aux:pass" {
			if test(q, /* $node[../node[@rel="su" and not(@pt or @cat)] and
				   ../node[@rel="vc" and not(descendant-or-self::node/@pt)] and
				   ../@rel="cnj"] */&xPath{
					arg1: &dSort{
						arg1: &dFilter{
							arg1: &dVariable{
								VAR: node,
							},
							arg2: &dSort{
								arg1: &dAnd{
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
															DATA: []interface{}{"vc"},
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
								},
							},
						},
					},
				}) {
				return "conj"
			}
			return "aux:pass"
		}
		if aux == "aux" {
			if test(q, /* $node[../node[@rel="su" and not(@pt or @cat)] and
				   ../node[@rel="vc" and not(descendant-or-self::node/@pt)] and
				   ../@rel="cnj"] */&xPath{
					arg1: &dSort{
						arg1: &dFilter{
							arg1: &dVariable{
								VAR: node,
							},
							arg2: &dSort{
								arg1: &dAnd{
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
															DATA: []interface{}{"vc"},
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
								},
							},
						},
					},
				}) {
				return "conj"
			}
			return "aux"
		}
		if aux == "cop" {
			return "cop"
		}
	}
	if node.Rel == "det" {
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
		}) {
			return detLabel(node, q)
		}
		if test(q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"mod"},
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
			return detLabel(node, q) // was "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "obj1" || node.Rel == "me" {
		if test(q /* $node/../@cat="pp" or $node/../node[@rel="hd" and @ud:pos="ADP"] */, &xPath{
			arg1: &dSort{
				arg1: &dOr{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG: collect__attributes__cat,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
						},
						arg2: &dElem{
							DATA: []interface{}{"pp"},
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
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
		}) { // vol vertrouwen , heel de geschiedenis door (cat=ap!)
			if test(q /* $node/../node[@rel="predc"] */, &xPath{
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
			}) { // absolutive met
				return "nsubj"
			}
			return dependencyLabel(node.parent, q)
		}
		if test(q /* $node[@index = ../../node[@rel="su"]/@index and ../node[@rel="hd" and (@pt or @cat)]] */, &xPath{
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
									ARG:  collect__attributes__index,
									arg1: &dNode{},
								},
								arg2: &dCollect{
									ARG: collect__attributes__index,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG: collect__parent__type__node,
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
			return "nsubj:pass" // trees where su (with extraposed material) is spelled out at position of obj1
		} // but not elliptic cases (with empty head) like: Boven de engel is de profeet Zacharia afgebeeld en boven Maria de profeet Micha
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
		}) {
			return "obj"
		}
		if test(q /* $node/../node[@rel="su" and (@pt or @cat)] */, &xPath{
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
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "mwp" {
		if node.Begin >= 0 && node.Begin == node.parent.Begin {
			return dependencyLabel(node.parent, q)
		}
		if test(q /* $node[@pt and ../node/@index = ancestor::node[@cat="whq"]/node[@rel="whd"]/@index] */, &xPath{
			arg1: &dSort{
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
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG: collect__attributes__index,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG:  collect__parent__type__node,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dCollect{
									ARG: collect__attributes__index,
									arg1: &dCollect{
										ARG: collect__child__node,
										arg1: &dCollect{
											ARG:  collect__ancestors__node,
											arg1: &dNode{},
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
										},
									},
								},
							},
						},
					},
				},
			},
		}) { // wat voor constructions, where wat is dislocated
			return "fixed"
		}
		if node == nLeft(find(q /* $node/../node[@rel="mwp" and (@pt or @cat)] */, &xPath{
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
		})) {
			return dependencyLabel(node.parent, q)
		}
		if test(q /* $node/../node[@ud:pos="PROPN"] */, &xPath{
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
								DATA: []interface{}{"PROPN"},
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
			return "flat"
		}
		return "fixed" // v2 mwe-> fixed
	}
	if node.Rel == "cnj" {
		if node == nLeft(find(q /* $node/../node[@rel="cnj"] */, &xPath{
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
		})) { //changed picking first cnj in XML to nLeft test GB 4/3/21
			return dependencyLabel(node.parent, q)
		}
		return "conj"
	}
	if node.Rel == "dp" {
		if node == nLeft(find(q /* $node/../node[@rel="dp"] */, &xPath{
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
		})) {
			return dependencyLabel(node.parent, q)
		}
		return "parataxis"
	}
	if node.Rel == "tag" || node.Rel == "sat" {
		return "parataxis"
	}
	if node.Rel == "dlink" {
		return "mark"
	}
	if node.Rel == "nucl" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "vc" {
		if test(q /* $node/../node[@rel="hd" and @ud:pos=("AUX","ADP")] and not($node/../node[@rel="predc"]) */, &xPath{
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
			return dependencyLabel(node.parent, q)
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
		}) {
			if test(q, /* $node/node[@rel="su" and @index and not(@word or @cat)]
				   (: or $node[@cat=("ti","oti")] :)
				   or $node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)]
				   or $node[@cat="oti"]/node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)] */&xPath{
					arg1: &dSort{
						arg1: &dOr{
							arg1: &dOr{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dVariable{
										VAR: node,
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
														DATA: []interface{}{"su"},
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
								arg2: &dCollect{
									ARG: collect__child__node,
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
														DATA: []interface{}{"ti"},
														arg1: &dCollect{
															ARG:  collect__attributes__cat,
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
							arg2: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__child__node,
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
														DATA: []interface{}{"oti"},
														arg1: &dCollect{
															ARG:  collect__attributes__cat,
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
													ARG:  collect__attributes__cat,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"ti"},
													arg1: &dCollect{
														ARG:  collect__attributes__cat,
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
					},
				}) {
				return "xcomp"
			}
			if test(q /* $node/../@cat="np" */, &xPath{
				arg1: &dSort{
					arg1: &dEqual{
						ARG: equal__is,
						arg1: &dCollect{
							ARG: collect__attributes__cat,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: node,
								},
							},
						},
						arg2: &dElem{
							DATA: []interface{}{"np"},
							arg1: &dCollect{
								ARG: collect__attributes__cat,
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
			}) {
				return "acl" // v2: clausal dependents of nouns always acl
			}
			return "ccomp"
		}
		if test(q /* $node/../node[@rel=("su","obj1") and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"su", "obj1"},
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
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "mod" && node.parent.Rel == "det" { // [ap/det steeds vnw/meer] invloed
		return "amod"
	}
	if (node.Rel == "mod" || node.Rel == "pc" || node.Rel == "ld") && node.parent.Cat == "np" { // [detp niet veel] meer
		// modification of nomimal heads
		// pc and ld occur in nominalizations
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
		}) {
			return modLabelInsideNp(node, q)
		}
		if node == nLeft(find(q /* $node/../node[@rel=("mod") and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"mod"},
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
		})) { // gapping with multiple mods
			return dependencyLabel(node.parent, q) // gapping, where this mod is the head
		}
		return modLabelInsideNp(node, q) //was "orphan"
	}

	if test(q /* $node[@rel=("pc","ld") and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp","cp","whrel")] */, &xPath{
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
								DATA: []interface{}{"pc", "ld"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"sv1", "smain", "ssub", "inf", "ppres", "ppart", "oti", "ap", "advp", "cp", "whrel"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) { // give pc priority GB 8/2/21
		// modification of verbal, adjectival heads
		// nb some oti's directly dominate (preceding) modifiers
		// [advp weg ermee ]
		if test(q /* $node/../node[@rel=("hd","body") and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"hd", "body"},
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
		}) { // body for mods dangling outside cmp/body: maar niet om ...
			return labelVmod(node, q)
		}
		if test(q /* $node/../node[@rel=("su","obj1","predc","vc") and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"su", "obj1", "predc", "vc"},
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
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping, where this mod is the head
	}

	if test(q /* $node[@rel=("mod") and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp","cp","whrel")] */, &xPath{
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
								DATA: []interface{}{"mod"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"sv1", "smain", "ssub", "inf", "ppres", "ppart", "oti", "ap", "advp", "cp", "whrel"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) {
		// modification of verbal, adjectival heads
		// nb some oti's directly dominate (preceding) modifiers
		// [advp weg ermee ]
		if test(q /* $node/../node[@rel=("hd","body") and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"hd", "body"},
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
		}) { // body for mods dangling outside cmp/body: maar niet om ...
			return labelVmod(node, q)
		}
		if test(q /* $node[not(../node[@rel="hd"]) and ../node[@rel="predc" and (@pt or @cat)]] */, &xPath{
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
					},
				},
			},
		}) { // mod in context of predc but no empty head
			// avoid orphan, because that would not disappear in enhanced
			return labelVmod(node, q)
		}
		if test(q /* $node/../node[@rel=("su","obj1","predc","vc","pc") and (@pt or @cat)] */, &xPath{
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
									DATA: []interface{}{"su", "obj1", "predc", "vc", "pc"},
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
		}) { // added pc 16/2/21
			return "orphan"
		}
		// combined mod/ld/pc for consistency with dephead position & superfluous again, see above. Gb 8/2 BN begin-position checks dont work as these are renumbered internally
		if test(q /* $node[@rel=("mod","ld","pc") and ../node[@rel=("mod","pc","ld") and (@pt or @cat)]/@begin < @begin ] */, &xPath{
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
									DATA: []interface{}{"mod", "ld", "pc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
							arg2: &dCmp{
								ARG: cmp__lt,
								arg1: &dCollect{
									ARG: collect__attributes__begin,
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
														DATA: []interface{}{"mod", "pc", "ld"},
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
								arg2: &dCollect{
									ARG:  collect__attributes__begin,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) { // gapping with multiple mods
			return "orphan" //  added (@pt or @cat) to prevent match with an empty mod, eventhough those do not have a @begin attribute
		} // seems like a BUG
		return dependencyLabel(node.parent, q) // gapping, where this mod is the head
	}
	if test(q /* $node[@rel="mod" and ../@cat=("pp","part")] */, &xPath{
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
								DATA: []interface{}{"mod"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"pp", "part"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) { // [mod hd/ADP obj1/empty]  --> make mod the external head , added part for automatic parse output GB 31/03/21
		if test(q /* $node/../node[@rel="obj1" and (@pt or @cat)] */, &xPath{
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
			return "amod"
		}
		if test(q /* $node/../node[@rel="hd" and @ud:pos=("ADV","ADP")] */, &xPath{
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
									DATA: []interface{}{"ADV", "ADP"},
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
		}) { // daarom dus, vlak voor en tijdens de oorlog --> orphan or advmod?
			return "advmod"
		}
		return dependencyLabel(node.parent, q)
	}
	if test(q /* $node[@rel="mod" and ../@cat=("detp","advp")] */, &xPath{
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
								DATA: []interface{}{"mod"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"detp", "advp"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) {
		return "amod"
	}

	if test(q /* $node[@rel="mod" and ../@cat="conj"] */, &xPath{
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
								DATA: []interface{}{"mod"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"conj"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) { // occurs in automatic parse output, should be more finegrained if it happens more often GB 22/03/21
		return "nmod"
	}

	if test(q /* $node[@rel="mod" and ../@cat=("cp", "whrel", "whq", "whsub")] */, &xPath{
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
								DATA: []interface{}{"mod"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"cp", "whrel", "whq", "whsub"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) {
		// [cp net  [cmp als] [body de Belgische regering]], net wat het land nodig heeft
		if test(q /* $node/../node[@rel="body"]/node[@rel="hd" and @ud:pos="VERB"] */, &xPath{
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
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"VERB"},
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
			return "advmod"
		}
		return "amod"
	}
	if node.Rel == "pc" && node.parent.Cat == "pp" { //  [[hd,mwu aan het hoofd] [pc van X]]
		return "nmod"
	}
	if node.Rel == "hdf" {
		return "case"
	}
	if node.Rel == "predm" {
		if node.udPos == "VERB" || node.udPos == "ADJ" {
			return "advcl"
		}
		if node.udPos == "PRON" || node.udPos == "NOUN" { //floating quantifiers, zich politiek bezoedelen
			return "obl" // , https://github.com/UniversalDependencies/docs/issues/581, https://github.com/UniversalDependencies/UD_French-Sequoia/issues/6,
		}
		if node.udPos != "" {
			return "advmod"
		}
		return "advcl"
	}
	if node.Rel == "rhd" || node.Rel == "whd" {
		if test(q /* $node/../node[@rel="body"]//node/@index = $node/@index */, &xPath{
			arg1: &dSort{
				arg1: &dEqual{
					ARG: equal__is,
					arg1: &dCollect{
						ARG: collect__attributes__index,
						arg1: &dCollect{
							ARG: collect__descendant__node,
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
					},
					arg2: &dCollect{
						ARG: collect__attributes__index,
						arg1: &dVariable{
							VAR: node,
						},
					},
				},
			},
		}) {
			return nonLocalDependencyLabel(
				node,
				n1(find(q /* ($node/../node[@rel="body"]//node[@index = $node/@index])[1] */, &xPath{
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
				})),
				q)
		}
		if node.Cat == "pp" {
			return "nmod" // onder wie michael boogerd
		}
		if node.udPos == "PRON" { // [whd wat] [body nu]
			return "obl"
		}
		return "advmod" // [whq waarom jij]
	}
	if node.Rel == "body" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "--" {
		// debugging, removed last escape for punct/num/sym from this position
		// commented out, as seems to be covered by clause below and gives errors in some cases
		// restricted to PUNCT now...26/03/21
		if test(q /* $node[@cat="mwu" and node/@ud:pos=("PUNCT","SYM")] */, &xPath{
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
									ARG: collect__attributes__ud_3apos,
									arg1: &dCollect{
										ARG:  collect__child__node,
										arg1: &dNode{},
									},
								},
								arg2: &dElem{
									DATA: []interface{}{"PUNCT", "SYM"},
									arg1: &dCollect{
										ARG: collect__attributes__ud_3apos,
										arg1: &dCollect{
											ARG:  collect__child__node,
											arg1: &dNode{},
										},
									},
								},
							},
						},
					},
				},
			},
		}) { // fix for mwu punctuation in Alpino output
			if test(q /* $node/../node[not(@ud:pos="PUNCT" or (@cat="mwu" and node/@ud:pos=("PUNCT","SYM")))] */, &xPath{
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
							arg1: &dFunction{
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
													DATA: []interface{}{"PUNCT"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
													},
												},
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
															ARG: collect__attributes__ud_3apos,
															arg1: &dCollect{
																ARG:  collect__child__node,
																arg1: &dNode{},
															},
														},
														arg2: &dElem{
															DATA: []interface{}{"PUNCT", "SYM"},
															arg1: &dCollect{
																ARG: collect__attributes__ud_3apos,
																arg1: &dCollect{
																	ARG:  collect__child__node,
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
				return "punct"
			}
			return "root"
		}
		if node.udPos == "PUNCT" {
			if test(q /* $node[not(../node[not(@ud:pos="PUNCT")]) and @begin = ../@begin] */, &xPath{
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
											arg1: &dCollect{
												ARG: collect__child__node,
												arg1: &dCollect{
													ARG:  collect__parent__type__node,
													arg1: &dNode{},
												},
												arg2: &dPredicate{
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
											},
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
			}) {
				return "root" // just punctuation
			}
			return "punct"
		}
		if node.udPos == "X" { // SYM covered below
			if test(q /* $node/../node[@cat] */, &xPath{
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
							arg1: &dCollect{
								ARG:  collect__attributes__cat,
								arg1: &dNode{},
							},
						},
					},
				},
			}) {
				return "parataxis" // 1. Jantje is ziek  1-->appos??
			}
			// return "root"
		}
		if node.Lemma == "\\" {
			return "punct" // hack for tagging errors in lassy small 250
		}
		if test(q /* $node[@ud:pos="NUM" and ../node[@cat] ] */, &xPath{
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
									DATA: []interface{}{"NUM"},
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
			return "parataxis" // dangling number 1.
		}
		if test(q /* $node[@ud:pos="CCONJ" and ../node[@cat="smain" or @cat="conj"]] */, &xPath{
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
									DATA: []interface{}{"CCONJ"},
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
									arg1: &dOr{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"smain"},
												arg1: &dCollect{
													ARG:  collect__attributes__cat,
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
						},
					},
				},
			},
		}) {
			return "cc"
		}
		// sentence initial or final 'en' etc , merge with statement below??
		if test(q /* $node[@ud:pos=("NOUN","PROPN","DET","ADP","ADV","INTJ","PRON","SYM","CCONJ","SCONJ") and ../node[@cat=("du","smain","conj","sv1","np","whq","pp","ppart","inf","advp")]] */, &xPath{
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
									DATA: []interface{}{"NOUN", "PROPN", "DET", "ADP", "ADV", "INTJ", "PRON", "SYM", "CCONJ", "SCONJ"},
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
									arg1: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"du", "smain", "conj", "sv1", "np", "whq", "pp", "ppart", "inf", "advp"},
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
		}) {
			return "parataxis" // dangling words, generalize to all POS ? GB 26/03/21
		}

		if test(q /* $node[@cat="mwu" or @ud:pos=("ADP","ADV","DET","PRON","CCONJ","SCONJ","NOUN","PROPN","INTJ","NUM","SYM")] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
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
									DATA: []interface{}{"ADP", "ADV", "DET", "PRON", "CCONJ", "SCONJ", "NOUN", "PROPN", "INTJ", "NUM", "SYM"},
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
		}) { // make exception for du nodes as well GB 22/03/21
			if node == nLeft(find(q, /* $node/../node[(@cat="mwu" or @ud:pos=("ADP","ADV","DET","PRON","CCONJ","SCONJ","NOUN","PROPN","INTJ","NUM","SYM"))
				   and not(../node[@ud:pos=("ADJ","VERB") or @cat=("du","ap","smain","ppart","ti")]) ] */&xPath{
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
													DATA: []interface{}{"ADP", "ADV", "DET", "PRON", "CCONJ", "SCONJ", "NOUN", "PROPN", "INTJ", "NUM", "SYM"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &dNode{},
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
														arg1: &dOr{
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
															arg2: &dEqual{
																ARG: equal__is,
																arg1: &dCollect{
																	ARG:  collect__attributes__cat,
																	arg1: &dNode{},
																},
																arg2: &dElem{
																	DATA: []interface{}{"du", "ap", "smain", "ppart", "ti"},
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
							},
						},
					},
				})) {
				return "root"
			}
			return "parataxis"
		}
		if test(q /* $node[@ud:pos=("ADJ","VERB")] */, &xPath{
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
		}) {
			if node == nLeft(find(q /* $node/../node[@ud:pos=("ADJ","VERB")] */, &xPath{
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
			})) {
				return "root"
			}
			return "parataxis"
		}
		if test(q /* $node[not(@ud:pos)]/../@rel="top" */, &xPath{
			arg1: &dSort{
				arg1: &dEqual{
					ARG: equal__is,
					arg1: &dCollect{
						ARG: collect__attributes__rel,
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
					arg2: &dElem{
						DATA: []interface{}{"top"},
						arg1: &dCollect{
							ARG: collect__attributes__rel,
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
		}) {
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
				return "parataxis"
			}
			return "root"
		}

		//debugging, moved this to last escape option fixing [MISSING PARA] error messages ...

		if test(q /* count($node/../node[not(@ud:pos=("PUNCT","SYM","X"))]) < 2 */, &xPath{
			arg1: &dSort{
				arg1: &dCmp{
					ARG: cmp__lt,
					arg1: &dFunction{
						ARG: function__count__1__args,
						arg1: &dArg{
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: node,
									},
								},
								arg2: &dPredicate{
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
														DATA: []interface{}{"PUNCT", "SYM", "X"},
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
					arg2: &dElem{
						DATA: []interface{}{2},
						arg1: &dFunction{
							ARG: function__count__1__args,
							arg1: &dArg{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dVariable{
											VAR: node,
										},
									},
									arg2: &dPredicate{
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
															DATA: []interface{}{"PUNCT", "SYM", "X"},
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
			return "root" // only one non-punct/sym/foreign element in the string
		}
		panic("No label --")
	}

	if node.Rel == "hd" {
		if node.udPos == "ADP" || node.parent.Cat == "pp" { // added pp check for PPs headed by a verb (gezien de structuur....) GB 23/03/21
			if test(q /* $node/../node[@rel="predc"] */, &xPath{
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
			}) {
				return "mark" // absolute met constructie -- analoog aan with X being Y
			}
			if test(q /* $node[../node[@rel="obj1" and @index and not(@pt or @cat)] and not(../@rel="cnj")] */, &xPath{
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
								arg2: &dFunction{
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
										},
									},
								},
							},
						},
					},
				},
			}) {
				return "case" // waar-relatives with P-stranding: waar hij Gabbema voor bedankt
			}
			if test(q /* $node/../node[@rel=("obj1","vc","se","me") and (@pt or @cat)] */, &xPath{
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
										DATA: []interface{}{"obj1", "vc", "se", "me"},
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
			}) { // examples in paqus suggest me case is already covered (advmod), yet leuven/253 gives error without me here..
				return "case" //
			}
			if test(q /* $node/../node[@rel="pc"] */, &xPath{
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
									DATA: []interface{}{"pc"},
									arg1: &dCollect{
										ARG:  collect__attributes__rel,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			}) { // superfluous ??
				return dependencyLabel(node.parent, q) // er blijft weinig over van het lijk : over heads a predc and has pc as sister
			}
			if test(q /* $node[../node[@rel="mod" and (@pt or @cat)] and ../@cat="pp"] */, &xPath{
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
													DATA: []interface{}{"mod"},
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
								arg2: &dEqual{
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
							},
						},
					},
				},
			}) { //
				return dependencyLabel(node.parent, q) // [mod om wat te zonnen] in [1] en bij [1 de kleine meertjes]  , changed from parataxis, GB 3/3/21 (consistent with treatment in depheads)
				// actually, this case is redundant now
			}
			return dependencyLabel(node.parent, q) // [predc [mod het beste] af/hd,ADP] here af heads a predc --> go to parent
		}
		if test(q, /* $node[(@ud:pos=("ADJ","X","ADV") or @cat="mwu")
			   and ../@cat="pp"
			   and ../node[@rel=("obj1","se","vc")]] */&xPath{
				arg1: &dSort{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: node,
						},
						arg2: &dSort{
							arg1: &dAnd{
								arg1: &dAnd{
									arg1: &dSort{
										arg1: &dOr{
											arg1: &dEqual{
												ARG: equal__is,
												arg1: &dCollect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"ADJ", "X", "ADV"},
													arg1: &dCollect{
														ARG:  collect__attributes__ud_3apos,
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
									arg2: &dEqual{
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
												DATA: []interface{}{"obj1", "se", "vc"},
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
			if test(q /* $node[../@rel="cnj" and ../node[@rel="obj1" and @index and not(@pt or @cat)]] */, &xPath{
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
			}) {
				return "conj"
			}
			return "case" // vol vertrouwen in, Ad U3... , op het gebied van
		}
		if test(q /* $node/../node[@rel="hd"]/@begin < $node/@begin */, &xPath{
			arg1: &dSort{
				arg1: &dCmp{
					ARG: cmp__lt,
					arg1: &dCollect{
						ARG: collect__attributes__begin,
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
					arg2: &dCollect{
						ARG: collect__attributes__begin,
						arg1: &dVariable{
							VAR: node,
						},
					},
				},
			},
		}) {
			return "conj"
		}
		return dependencyLabel(node.parent, q)
	}
	panic("No label")
}

func determineNominalModLabel(node *nodeType, q *context) string {
	if test(q /* $node/../node[@rel="hd" and (@ud:pos="VERB" or @ud:pos="ADJ")] */, &xPath{
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
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"VERB"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
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
										DATA: []interface{}{"ADJ"},
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
	}) {
		return "obl"
	}
	return "nmod"
}

func subjectLabel(subj *nodeType, q *context) string {
	// pass (requires pass aux) and outer (requires copula) never co-occur I assume
	pass := passiveSubject(subj, q)
	outer := outerSubject(subj, q)
	if test(q, /* $subj[@cat=("whsub","ssub","ti","cp","oti")] or
		   $subj[@cat="conj" and node/@cat=("whsub","ssub","ti","cp","oti")] */&xPath{
			arg1: &dSort{
				arg1: &dOr{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: subj,
						},
						arg2: &dSort{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"whsub", "ssub", "ti", "cp", "oti"},
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
									},
								},
							},
						},
					},
					arg2: &dFilter{
						arg1: &dVariable{
							VAR: subj,
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
										DATA: []interface{}{"conj"},
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
									},
								},
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG: collect__attributes__cat,
										arg1: &dCollect{
											ARG:  collect__child__node,
											arg1: &dNode{},
										},
									},
									arg2: &dElem{
										DATA: []interface{}{"whsub", "ssub", "ti", "cp", "oti"},
										arg1: &dCollect{
											ARG: collect__attributes__cat,
											arg1: &dCollect{
												ARG:  collect__child__node,
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
		return "csubj" + pass + outer
	}
	// weather verbs and other expletive subject verbs
	if test(q, /* $subj[@lemma="het"] and
		   ( $subj/../node[@rel="hd" and
		                   @lemma=("dooien", "gieten", "hagelen", "miezeren",
		                           "misten", "motregenen", "onweren", "plenzen",
		                           "regenen", "sneeuwen", "stormen", "stortregenen",
		                           "ijzelen", "vriezen", "weerlichten", "winteren",
		                           "zomeren") ] or
		     $subj/../node[@rel="hd" and
		                   @lemma="ontbreken" and
		                   ../node[@rel="pc" and node[@rel="hd" and @lemma="aan"] ] ] or
		     $subj[@index = ../node//node[@rel="sup"]/@index]    (: het kan voorkomen dat ... :)
		   ) */&xPath{
			arg1: &dSort{
				arg1: &dAnd{
					arg1: &dFilter{
						arg1: &dVariable{
							VAR: subj,
						},
						arg2: &dSort{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"het"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
					arg2: &dSort{
						arg1: &dOr{
							arg1: &dOr{
								arg1: &dCollect{
									ARG: collect__child__node,
									arg1: &dCollect{
										ARG: collect__parent__type__node,
										arg1: &dVariable{
											VAR: subj,
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
													ARG:  collect__attributes__lemma,
													arg1: &dNode{},
												},
												arg2: &dElem{
													DATA: []interface{}{"dooien", "gieten", "hagelen", "miezeren",
														"misten", "motregenen", "onweren", "plenzen",
														"regenen", "sneeuwen", "stormen", "stortregenen",
														"ijzelen", "vriezen", "weerlichten", "winteren",
														"zomeren"},
													arg1: &dCollect{
														ARG:  collect__attributes__lemma,
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
										arg1: &dVariable{
											VAR: subj,
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
												arg2: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__lemma,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"ontbreken"},
														arg1: &dCollect{
															ARG:  collect__attributes__lemma,
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
																DATA: []interface{}{"pc"},
																arg1: &dCollect{
																	ARG:  collect__attributes__rel,
																	arg1: &dNode{},
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
																	arg2: &dEqual{
																		ARG: equal__is,
																		arg1: &dCollect{
																			ARG:  collect__attributes__lemma,
																			arg1: &dNode{},
																		},
																		arg2: &dElem{
																			DATA: []interface{}{"aan"},
																			arg1: &dCollect{
																				ARG:  collect__attributes__lemma,
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
							arg2: &dFilter{
								arg1: &dVariable{
									VAR: subj,
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
														ARG: collect__child__node,
														arg1: &dCollect{
															ARG:  collect__parent__type__node,
															arg1: &dNode{},
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
															DATA: []interface{}{"sup"},
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
			},
		}) {
		return "expl" + pass
	}
	return "nsubj" + pass + outer
}

// add :outer to subject in copula construction with clausal predicate (whrel, other?)
func outerSubject(subj *nodeType, q *context) string {
	if test(q /* $subj/../node[@rel="predc" and (@cat="whrel" or (@cat="conj" and node[@cat="whrel"]))] */, &xPath{
		arg1: &dSort{
			arg1: &dCollect{
				ARG: collect__child__node,
				arg1: &dCollect{
					ARG: collect__parent__type__node,
					arg1: &dVariable{
						VAR: subj,
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
								DATA: []interface{}{"predc"},
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
										DATA: []interface{}{"whrel"},
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
									},
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
												DATA: []interface{}{"conj"},
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
												arg1: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__cat,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"whrel"},
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
				},
			},
		},
	}) {
		return ":outer"
	}
	return ""
}

// recursive
func passiveSubject(subj *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "passiveSubject", q, nil, nil, nil, subj))
		}
	}()

	depthCheck(q)

	if aux, err := auxiliary1(n1(find(q /* ($subj/../node[@rel="hd"])[1] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: subj,
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
				arg2: &dSort{
					arg1: &dFunction{
						ARG: function__first__0__args,
					},
				},
			},
		},
	})), q); err == nil {
		if aux == "aux:pass" { // de carriere had gered kunnen worden
			return ":pass"
		}
		if aux == "aux" && test(q /* $subj/@index = $subj/../node[@rel="vc"]/node[@rel="su"]/@index */, &xPath{
			arg1: &dSort{
				arg1: &dEqual{
					ARG: equal__is,
					arg1: &dCollect{
						ARG: collect__attributes__index,
						arg1: &dVariable{
							VAR: subj,
						},
					},
					arg2: &dCollect{
						ARG: collect__attributes__index,
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__child__node,
								arg1: &dCollect{
									ARG: collect__parent__type__node,
									arg1: &dVariable{
										VAR: subj,
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
				},
			},
		}) {
			return passiveSubject(n1(find(q /* $subj/../node[@rel="vc"]/node[@rel="su"] */, &xPath{
				arg1: &dSort{
					arg1: &dCollect{
						ARG: collect__child__node,
						arg1: &dCollect{
							ARG: collect__child__node,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: subj,
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
			})), q)
		}
	}
	return ""
}

func detLabel(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "detLabel", q, node))
		}
	}()

	// zijn boek, diens boek, ieders boek, aller landen, Ron's probleem, Fidel Castro's belang, zuids schoppenaas
	if test(q, /* $node[@ud:pos = "PRON" and @vwtype="bez"] or
		   $node[@ud:pos = ("PRON","PROPN","NOUN") and @naamval="gen"] or
		   $node[@cat="mwu" and node[@spectype="deeleigen"]] */&xPath{
			arg1: &dSort{
				arg1: &dOr{
					arg1: &dOr{
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
											DATA: []interface{}{"PRON"},
											arg1: &dCollect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &dNode{},
											},
										},
									},
									arg2: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__vwtype,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"bez"},
											arg1: &dCollect{
												ARG:  collect__attributes__vwtype,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
						arg2: &dFilter{
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
											DATA: []interface{}{"PRON", "PROPN", "NOUN"},
											arg1: &dCollect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &dNode{},
											},
										},
									},
									arg2: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__naamval,
											arg1: &dNode{},
										},
										arg2: &dElem{
											DATA: []interface{}{"gen"},
											arg1: &dCollect{
												ARG:  collect__attributes__naamval,
												arg1: &dNode{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &dFilter{
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
										DATA: []interface{}{"mwu"},
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
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__spectype,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"deeleigen"},
												arg1: &dCollect{
													ARG:  collect__attributes__spectype,
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
		return "nmod:poss"
	}
	if test(q /* $node/@ud:pos = ("DET","PROPN","PRON","X") */, &xPath{
		arg1: &dSort{
			arg1: &dEqual{
				ARG: equal__is,
				arg1: &dCollect{
					ARG: collect__attributes__ud_3apos,
					arg1: &dVariable{
						VAR: node,
					},
				},
				arg2: &dElem{
					DATA: []interface{}{"DET", "PROPN", "PRON", "X"},
					arg1: &dCollect{
						ARG: collect__attributes__ud_3apos,
						arg1: &dVariable{
							VAR: node,
						},
					},
				},
			},
		},
	}) {
		return "det" // meer // genoeg // the
	}
	if node.Cat == "detp" {
		if test(q /* $node/node[@rel="hd" and @ud:pos="NUM"] */, &xPath{
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
									DATA: []interface{}{"NUM"},
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
			return "nummod"
		}
		if test(q /* $node/node[@rel="hd" and @ud:pos=("NOUN","ADJ")] */, &xPath{
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
									DATA: []interface{}{"NOUN", "ADJ"},
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
			return "nmod"
		}
		if test(q /* $node/node[@rel="hd" and @ud:pos="PRON" and @vwtype="bez"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: node,
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
								arg2: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"PRON"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &dNode{},
										},
									},
								},
							},
							arg2: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__vwtype,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"bez"},
									arg1: &dCollect{
										ARG:  collect__attributes__vwtype,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		}) {
			return "nmod:poss"
		}
		return "det"
	}
	if test(q /* $node[@cat=("np","ap") or @ud:pos=("SYM","ADJ","ADV","NOUN") ] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
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
								DATA: []interface{}{"np", "ap"},
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
								DATA: []interface{}{"SYM", "ADJ", "ADV", "NOUN"},
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
		return "nmod"
	}
	if test(q /* $node/@cat = ("mwu","smain") */, &xPath{
		arg1: &dSort{
			arg1: &dEqual{
				ARG: equal__is,
				arg1: &dCollect{
					ARG: collect__attributes__cat,
					arg1: &dVariable{
						VAR: node,
					},
				},
				arg2: &dElem{
					DATA: []interface{}{"mwu", "smain"},
					arg1: &dCollect{
						ARG: collect__attributes__cat,
						arg1: &dVariable{
							VAR: node,
						},
					},
				},
			},
		},
	}) {
		return "det"
	}
	// tussen 5 en 6 ..--> almost all PP cases are with tussen NUM and NUM
	if test(q /* $node/@cat = "pp" */, &xPath{
		arg1: &dSort{
			arg1: &dEqual{
				ARG: equal__is,
				arg1: &dCollect{
					ARG: collect__attributes__cat,
					arg1: &dVariable{
						VAR: node,
					},
				},
				arg2: &dElem{
					DATA: []interface{}{"pp"},
					arg1: &dCollect{
						ARG: collect__attributes__cat,
						arg1: &dVariable{
							VAR: node,
						},
					},
				},
			},
		},
	}) {
		return "nummod"
	}
	// nog meer mensen dan anders
	// hetzelfde tijdstip als anders , nogal wat,
	// hij heeft ik weet niet hoeveel boeken
	if node.udPos == "NUM" || node.udPos == "SYM" {
		return "nummod"
	}
	if node.Cat == "conj" {

		return detLabel(n1(find(q /* ($node/node[@rel="cnj"])[1] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
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
					arg2: &dSort{
						arg1: &dFunction{
							ARG: function__first__0__args,
						},
					},
				},
			},
		})), q)
	}
	if node.Cat == "cp" { //ik heb boeken gezien [cp/det dan hem] weird...
		return "nmod"
	}
	panic("No label det")
}

func modLabelInsideNp(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "modLabelInsideNp", q, node))
		}
	}()

	if test(q /* $node[@cat="pp"]/node[@rel="vc"] */, &xPath{
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
	}) {
		return "acl" // pp containing a clause
	}
	// fixing issue noted by Anouck
	if test(q /* $node[@ud:pos="ADJ" or @cat="ap" or node[@cat="conj" and @begin = node[@ud:pos="ADJ" or @cat="ap"]/@begin ]] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
				},
				arg2: &dSort{
					arg1: &dOr{
						arg1: &dOr{
							arg1: &dEqual{
								ARG: equal__is,
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"ADJ"},
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
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
									DATA: []interface{}{"ap"},
									arg1: &dCollect{
										ARG:  collect__attributes__cat,
										arg1: &dNode{},
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
									arg2: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG:  collect__attributes__begin,
											arg1: &dNode{},
										},
										arg2: &dCollect{
											ARG: collect__attributes__begin,
											arg1: &dCollect{
												ARG:  collect__child__node,
												arg1: &dNode{},
												arg2: &dPredicate{
													arg1: &dOr{
														arg1: &dEqual{
															ARG: equal__is,
															arg1: &dCollect{
																ARG:  collect__attributes__ud_3apos,
																arg1: &dNode{},
															},
															arg2: &dElem{
																DATA: []interface{}{"ADJ"},
																arg1: &dCollect{
																	ARG:  collect__attributes__ud_3apos,
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
																DATA: []interface{}{"ap"},
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
						},
					},
				},
			},
		},
	}) {
		return "amod"
	}
	if test(q /* $node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM","INTJ") ] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
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
								DATA: []interface{}{"pp", "np", "conj", "mwu"},
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
								DATA: []interface{}{"NOUN", "PRON", "PROPN", "X", "PUNCT", "SYM", "INTJ"},
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
		return "nmod"
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if test(q /* $node[@cat="detp"]/node[@rel="hd" and @ud:pos="NUM"] */, &xPath{
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
								DATA: []interface{}{"detp"},
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
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"NUM"},
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
		return "nummod"
	}
	if node.Cat == "detp" {
		return "det" // [detp niet veel] meer error?
	}
	if node.Cat == "rel" || node.Cat == "whrel" || node.Cat == "ssub" {
		return "acl:relcl"
	}
	// v2 added relcl -- whrel= met name waar ...
	if test(q /* $node[@cat="cp"]/node[@rel="body" and (@ud:pos = ("NOUN","PROPN") or @cat=("np","conj"))] */, &xPath{
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
								DATA: []interface{}{"cp"},
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
								DATA: []interface{}{"body"},
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
										DATA: []interface{}{"NOUN", "PROPN"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
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
										DATA: []interface{}{"np", "conj"},
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
	}) {
		return "nmod"
	}
	// zijn loopbaan [CP als schrijver]
	if test(q /* $node[@cat=("cp","sv1","smain","ppres","ppart","inf","ti","oti","du","whq", "whsub") or @ud:pos="SCONJ"] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
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
								DATA: []interface{}{"cp", "sv1", "smain", "ppres", "ppart", "inf", "ti", "oti", "du", "whq", "whsub"},
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
								DATA: []interface{}{"SCONJ"},
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
	}) { // added whsub for robustness in automatic parser output
		return "acl"
	}
	// oa zinnen tussen haakjes
	if test(q /* $node[@ud:pos= ("ADV","ADP","VERB","CCONJ") or @cat="advp"] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
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
								DATA: []interface{}{"ADV", "ADP", "VERB", "CCONJ"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
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
								DATA: []interface{}{"advp"},
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
		return "amod"
	}
	// VERB= aanstormend etc -> amod, ADV = nagenoeg alle prijzen, slechts 4 euro --> amod
	// CCONJ = opdrachten zoals:   --> amod
	if node.Rel == "det" {
		return "det" // empty determiners in gapping?
	}
	if node.Index > 0 {
		panic("Index nmod")
	}
	panic("No label nmod")
}

func labelVmod(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "labelVmod", q, node))
		}
	}()

	if test(q /* $node[@cat="pp"]/node[@rel="vc"] */, &xPath{
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
	}) {
		return "advcl"
	}
	if test(q, /* $node[ (   node[@rel="hd" and @lemma="door"]
			                or node[@rel="cnj"]/node[@rel="hd" and @lemma="door"]
		                    or (@pt="bw" and ends-with(@lemma,"door"))
		                   )
		                 ] */&xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dSort{
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
												arg2: &dEqual{
													ARG: equal__is,
													arg1: &dCollect{
														ARG:  collect__attributes__lemma,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"door"},
														arg1: &dCollect{
															ARG:  collect__attributes__lemma,
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
														ARG:  collect__attributes__lemma,
														arg1: &dNode{},
													},
													arg2: &dElem{
														DATA: []interface{}{"door"},
														arg1: &dCollect{
															ARG:  collect__attributes__lemma,
															arg1: &dNode{},
														},
													},
												},
											},
										},
									},
								},
								arg2: &dSort{
									arg1: &dAnd{
										arg1: &dEqual{
											ARG: equal__is,
											arg1: &dCollect{
												ARG:  collect__attributes__pt,
												arg1: &dNode{},
											},
											arg2: &dElem{
												DATA: []interface{}{"bw"},
												arg1: &dCollect{
													ARG:  collect__attributes__pt,
													arg1: &dNode{},
												},
											},
										},
										arg2: &dFunction{
											ARG: function__ends__with__2__args,
											arg1: &dArg{
												arg1: &dArg{
													arg1: &dSort{
														arg1: &dCollect{
															ARG:  collect__attributes__lemma,
															arg1: &dNode{},
														},
													},
												},
												arg2: &dElem{
													DATA: []interface{}{"door"},
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
		if aux, err := auxiliary1(n1(find(q /* $node/../../node[@rel="hd" and @lemma=("zijn","worden")] */, &xPath{
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
									ARG:  collect__attributes__lemma,
									arg1: &dNode{},
								},
								arg2: &dElem{
									DATA: []interface{}{"zijn", "worden"},
									arg1: &dCollect{
										ARG:  collect__attributes__lemma,
										arg1: &dNode{},
									},
								},
							},
						},
					},
				},
			},
		})), q); err == nil && aux == "aux:pass" {
			return "obl:agent"
		}
	}
	/*
		but NOT: door rookontwikkeling om het leven gekomen
		-- already filtered by constraint of su/obj1 control or @sc=passive (cheating, for impersonal passives) GB 20/04/23
		NOT: bij Bakema is een stoeptegel door de ruit gegooid
		NO/YES: hierdoor werd Prince door het grote publiek ontdekt
		ADDED: coordination cases GB 11/04/23
	*/

	if test(q /* $node[@cat=("pp","conj") and @rel="pc"] */, &xPath{
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
								DATA: []interface{}{"pp", "conj"},
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
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
								DATA: []interface{}{"pc"},
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
	}) { // added GB 20/11/23
		return "obl:arg" // what about enhanced deps (obl:arg:aan ? yes, see CZ treebanks)
	}
	if test(q /* $node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","VERB","PRON","PROPN","X","PUNCT","SYM","ADP") ] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: node,
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
								DATA: []interface{}{"pp", "np", "conj", "mwu"},
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
								DATA: []interface{}{"NOUN", "VERB", "PRON", "PROPN", "X", "PUNCT", "SYM", "ADP"},
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
		return "obl"
	}
	if test(q /* $node[@cat="advp"]/node[@ud:pos=("NOUN","VERB","ADP")] */, &xPath{
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
								DATA: []interface{}{"advp"},
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
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
							ARG:  collect__attributes__ud_3apos,
							arg1: &dNode{},
						},
						arg2: &dElem{
							DATA: []interface{}{"NOUN", "VERB", "ADP"},
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
		return "obl"
	}
	if test(q /* $node[@cat="ap"]/node[@ud:pos="NOUN" or @cat="np"] */, &xPath{
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
								DATA: []interface{}{"ap"},
								arg1: &dCollect{
									ARG:  collect__attributes__cat,
									arg1: &dNode{},
								},
							},
						},
					},
				},
				arg2: &dPredicate{
					arg1: &dOr{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"NOUN"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
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
								DATA: []interface{}{"np"},
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
	}) { // added NP for 'het hele jaar door' h_suite/53 GB 26/02/21
		return "obl"
	}
	if test(q /* $node[@cat=("cp","sv1","smain","ssub","ppres","ppart","ti","oti","inf","du","whq","whrel","rel","whsub")] */, &xPath{
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
							DATA: []interface{}{"cp", "sv1", "smain", "ssub", "ppres", "ppart", "ti", "oti", "inf", "du", "whq", "whrel", "rel", "whsub"},
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
		return "advcl"
	}
	if test(q, /* $node[@ud:pos= ("ADJ","ADV","SCONJ","INTJ")
		   or @cat=("advp","ap")
		   or (@cat="conj" and node/@ud:pos="ADV")] */&xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: node,
					},
					arg2: &dSort{
						arg1: &dOr{
							arg1: &dOr{
								arg1: &dEqual{
									ARG: equal__is,
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &dNode{},
									},
									arg2: &dElem{
										DATA: []interface{}{"ADJ", "ADV", "SCONJ", "INTJ"},
										arg1: &dCollect{
											ARG:  collect__attributes__ud_3apos,
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
										DATA: []interface{}{"advp", "ap"},
										arg1: &dCollect{
											ARG:  collect__attributes__cat,
											arg1: &dNode{},
										},
									},
								},
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
											DATA: []interface{}{"conj"},
											arg1: &dCollect{
												ARG:  collect__attributes__cat,
												arg1: &dNode{},
											},
										},
									},
									arg2: &dEqual{
										ARG: equal__is,
										arg1: &dCollect{
											ARG: collect__attributes__ud_3apos,
											arg1: &dCollect{
												ARG:  collect__child__node,
												arg1: &dNode{},
											},
										},
										arg2: &dElem{
											DATA: []interface{}{"ADV"},
											arg1: &dCollect{
												ARG: collect__attributes__ud_3apos,
												arg1: &dCollect{
													ARG:  collect__child__node,
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
		return "advmod" // niet of nauwelijks
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if node.Index > 0 {
		panic("Index vmod")
	}
	panic("No label vmod")
}

// this function is now also used to distribute dependents in coordination in Enhanced UD , so lot more rels and contexts are possible
// and passives, as in " hun taal werd gediscrimineerd en verboden"
func nonLocalDependencyLabel(head, gap *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "nonLocalDependencyLabel", q, nil, head, gap))
		}
	}()

	if gap.Rel == "su" {
		return subjectLabel(gap, q)
	}
	if gap.Rel == "obj1" {
		if head.Rel == "su" {
			return "nsubj:pass"
		}
		if test(q /* $gap/../@cat="pp" */, &xPath{
			arg1: &dSort{
				arg1: &dEqual{
					ARG: equal__is,
					arg1: &dCollect{
						ARG: collect__attributes__cat,
						arg1: &dCollect{
							ARG: collect__parent__type__node,
							arg1: &dVariable{
								VAR: gap,
							},
						},
					},
					arg2: &dElem{
						DATA: []interface{}{"pp"},
						arg1: &dCollect{
							ARG: collect__attributes__cat,
							arg1: &dCollect{
								ARG: collect__parent__type__node,
								arg1: &dVariable{
									VAR: gap,
								},
							},
						},
					},
				},
			},
		}) { // waar men de patienten [ i naar toe] schuift om hen vervolgens te vergeten . ✤
			return "obl"
		}
		return "obj" // ppart coordination -- see comment above
	}
	if gap.Rel == "obj2" {
		if head.udPos == "ADV" {
			return "advmod"
		}
		if head.udPos == "ADP" { // fixed for questions 'aan wie werd het gegeven' -- GB 2/3/23
			return "obl"
		}
		return "iobj"
	}
	if gap.Rel == "me" {
		return determineNominalModLabel(gap, q)
	}
	if gap.Rel == "predc" || gap.Rel == "predm" {
		return dependencyLabel(gap, q)
	}
	if gap.Rel == "pc" || gap.Rel == "ld" {
		if test(q /* $head/node[@rel="obj1"] */, &xPath{
			arg1: &dSort{
				arg1: &dCollect{
					ARG: collect__child__node,
					arg1: &dVariable{
						VAR: head,
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
		}) { // is dit onzin? de head kan een pp met een np inside zijn, maar dat zegt niets over nmod, dit moet gewoon obl zijn GB 17/11/23
			return "nmod"
		}
		if test(q /* $head[@ud:pos=("ADV", "ADP") or @cat=("advp","ap")] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: head,
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
									DATA: []interface{}{"ADV", "ADP"},
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
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
									DATA: []interface{}{"advp", "ap"},
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
			return "advmod" // waar precies zit je ..
		}
		if test(q /* $head[@ud:pos="PRON"] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: head,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"PRON"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) { // het voorstel, [rhd_i dat] de commissie [ld_i] introk (from automatic parses)
			return "obl"
		}
		panic("No label index PC")
	}
	if gap.Rel == "sup" || gap.Rel == "pobj1" {
		return "expl" // waar het om gaat is dat hij scoort, het is 1881 en dertien jaar geleden dat ...
	}
	if gap.Rel == "mwp" {
		return dependencyLabel(gap.parent, q) //wat heb je voor boeken gelezen
	}
	if gap.Rel == "vc" {
		return "ccomp" // wat ik me afvraag is of hij komt -- CLEFT
	}
	if test(q /* $gap[@rel="mod" and ../@cat=("np","pp")] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: gap,
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
								DATA: []interface{}{"mod"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"np", "pp"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) { // voornamelijk in kloosters en door vrouwen
		return modLabelInsideNp(head, q)
	}
	if test(q /* $gap[@rel="mod" and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")] */, &xPath{
		arg1: &dSort{
			arg1: &dFilter{
				arg1: &dVariable{
					VAR: gap,
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
								DATA: []interface{}{"mod"},
								arg1: &dCollect{
									ARG:  collect__attributes__rel,
									arg1: &dNode{},
								},
							},
						},
						arg2: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG: collect__attributes__cat,
								arg1: &dCollect{
									ARG:  collect__parent__type__node,
									arg1: &dNode{},
								},
							},
							arg2: &dElem{
								DATA: []interface{}{"sv1", "smain", "ssub", "inf", "ppres", "ppart", "oti", "ap", "advp"},
								arg1: &dCollect{
									ARG: collect__attributes__cat,
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
	}) {
		return labelVmod(head, q)
	}
	if gap.Rel == "mod" || gap.Rel == "spec" { // spec only used for funny coord
		if test(q /* $head[@cat=("pp","np") or @ud:pos=("NOUN","PRON")] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: head,
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
									DATA: []interface{}{"pp", "np"},
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
									DATA: []interface{}{"NOUN", "PRON"},
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
			return "nmod"
		}
		if test(q /* $head[@ud:pos=("ADV","ADP") or @cat= ("advp","ap","mwu","conj")] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: head,
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
									DATA: []interface{}{"ADV", "ADP"},
									arg1: &dCollect{
										ARG:  collect__attributes__ud_3apos,
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
									DATA: []interface{}{"advp", "ap", "mwu", "conj"},
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
			return "advmod" // hoe vaak -- AP, daar waar, waar en wanneer, voor als rhd
		}
		if test(q /* $head[@ud:pos="ADJ"] */, &xPath{
			arg1: &dSort{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: head,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"ADJ"},
								arg1: &dCollect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &dNode{},
								},
							},
						},
					},
				},
			},
		}) { // added GB 9/3/21
			return "amod" // zeker_i vier a i vijf miljard
		}
		panic("No label index mod")
	}
	if gap.Rel == "det" {
		return detLabel(head, q)
	}
	if test(q /* $gap[@rel="hd"] and $head[@ud:pos=("ADV","ADP")] */, &xPath{
		arg1: &dSort{
			arg1: &dAnd{
				arg1: &dFilter{
					arg1: &dVariable{
						VAR: gap,
					},
					arg2: &dSort{
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
				arg2: &dFilter{
					arg1: &dVariable{
						VAR: head,
					},
					arg2: &dSort{
						arg1: &dEqual{
							ARG: equal__is,
							arg1: &dCollect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &dNode{},
							},
							arg2: &dElem{
								DATA: []interface{}{"ADV", "ADP"},
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
	}) { // waaronder A, B, en C
		return "case"
	}
	if gap.Rel == "du" || gap.Rel == "dp" {
		return "parataxis"
	}
	panic("No label index")
}
