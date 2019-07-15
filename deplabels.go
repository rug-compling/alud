//
// GENERATED FILE -- DO NOT EDIT
//

package main

// recursive
func dependencyLabel(node *NodeType, q *Context) string {
	if depthCheck(q, "dependencyLabel") {
		return "ERROR_NO_LABEL"
	}

	if node.parent.Cat == "top" && node.parent.End == 1000 {
		return "root"
	}
	if node.Rel == "app" {
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
		}) {
			return "appos"
		}
		if Test(q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &XPath{
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
			return "orphan"
		}
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "cmp" {
		return "mark"
	}
	if node.Rel == "crd" {
		return "cc"
	}
	if node.Rel == "me" && Test(q /* $node[@rel="me" and not(../node[@ud:pos="ADP"]) ] */, &XPath{
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
								DATA: []interface{}{"me"},
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
	}) {
		return determineNominalModLabel(node, q)
	}
	if node.Rel == "obcomp" {
		return "advcl"
	}
	if node.Rel == "obj2" {
		if node.Cat == "pp" {
			return "obl"
		}
		return "iobj"
	}
	if node.Rel == "pobj1" {
		return "expl"
	}
	if node.Rel == "predc" {
		if Test(q /* $node/../node[@rel=("obj1","se") and (@pt or @cat)] or $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX")] */, &XPath{
			arg1: &Sort{
				arg1: &Or{
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
										DATA: []interface{}{"obj1", "se"},
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
					arg2: &Collect{
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
		}) {
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
			}) { // check for absolutive (met) constructions, https://github.com/UniversalDependencies/docs/issues/408
				if Test(q /* $node/../../@cat="np" */, &XPath{
					arg1: &Sort{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG: collect__attributes__cat,
								arg1: &Collect{
									ARG: collect__parent__type__node,
									arg1: &Collect{
										ARG: collect__parent__type__node,
										arg1: &Variable{
											VAR: node,
										},
									},
								},
							},
							arg2: &Elem{
								DATA: []interface{}{"np"},
								arg1: &Collect{
									ARG: collect__attributes__cat,
									arg1: &Collect{
										ARG: collect__parent__type__node,
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
					},
				}) {
					return "acl"
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
		if Test(q /* $node[../@rel="cnj" and ../node[@rel="hd" and not(@pt or @cat)]] */, &XPath{
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
									ARG: collect__attributes__rel,
									arg1: &Collect{
										ARG:  collect__parent__type__node,
										arg1: &Node{},
									},
								},
								arg2: &Elem{
									DATA: []interface{}{"cnj"},
									arg1: &Collect{
										ARG: collect__attributes__rel,
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
										arg2: &Function{
											ARG: function__not__1__args,
											arg1: &Arg{
												arg1: &Sort{
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
		}) { // gapping
			return dependencyLabel(node.parent, q)
		}
		if Test(q, /* $node[../@rel="vc" and ../node[@rel="hd" and not(@pt or @cat)]
			   and ../parent::node[@rel="cnj" and node[@rel="hd" and not(@pt or @cat)]]] */&XPath{
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
											ARG: collect__attributes__rel,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
										arg2: &Elem{
											DATA: []interface{}{"vc"},
											arg1: &Collect{
												ARG: collect__attributes__rel,
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
												arg2: &Function{
													ARG: function__not__1__args,
													arg1: &Arg{
														arg1: &Sort{
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
								arg2: &Collect{
									ARG: collect__parent__node,
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
													DATA: []interface{}{"cnj"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
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
														arg2: &Function{
															ARG: function__not__1__args,
															arg1: &Arg{
																arg1: &Sort{
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
						},
					},
				},
			}) { // gapping with subj downstairs
			// TODO: ../.. is veranderd in ../parent::node    is dat juist?
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
	aux := auxiliary1(node, q)
	if aux == "aux:pass" {
		if Test(q, /* $node[../node[@rel="su" and not(@pt or @cat)] and
			   ../node[@rel="vc" and not(@pt or @cat)] and
			   ../@rel="cnj"] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: node,
						},
						arg2: &Sort{
							arg1: &And{
								arg1: &And{
									arg1: &Collect{
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
												arg2: &Function{
													ARG: function__not__1__args,
													arg1: &Arg{
														arg1: &Sort{
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
														DATA: []interface{}{"vc"},
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
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG: collect__attributes__rel,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
										},
									},
									arg2: &Elem{
										DATA: []interface{}{"cnj"},
										arg1: &Collect{
											ARG: collect__attributes__rel,
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
			}) {
			return "conj"
		}
		return "aux:pass"
	}
	if aux == "aux" {
		return "aux"
	}
	if aux == "cop" {
		return "cop"
	}
	if node.Rel == "det" {
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
		}) {
			return detLabel(node, q)
		}
		if Test(q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &XPath{
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
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "obj1" || node.Rel == "me" {
		if Test(q /* $node/../@cat="pp" or $node/../node[@rel="hd" and @ud:pos="ADP"] */, &XPath{
			arg1: &Sort{
				arg1: &Or{
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
					arg2: &Collect{
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
			},
		}) { // vol vertrouwen , heel de geschiedenis door (cat=ap!)
			if Test(q /* $node/../node[@rel="predc"] */, &XPath{
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
			}) { // absolutive met
				return "nsubj"
			}
			return dependencyLabel(node.parent, q)
		}
		if Test(q /* $node[@index = ../../node[@rel="su"]/@index ] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
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
										ARG: collect__parent__type__node,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
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
						},
					},
				},
			},
		}) {
			return "nsubj:pass" // trees where su (with extraposed material) is spelled out at position of obj1
		}
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
		}) {
			return "obj"
		}
		if Test(q /* $node/../node[@rel="su" and (@pt or @cat)] */, &XPath{
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
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "mwp" {
		if node.Begin >= 0 && node.Begin == node.parent.Begin {
			return dependencyLabel(node.parent, q)
		}
		if Test(q /* $node/../node[@ud:pos="PROPN"] */, &XPath{
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
								ARG:  collect__attributes__ud_3apos,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"PROPN"},
								arg1: &Collect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
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
		if node == n1(Find(q /* $node/../node[@rel="cnj"][1] */, &XPath{
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
									DATA: []interface{}{"cnj"},
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
		})) {
			return dependencyLabel(node.parent, q)
		}
		return "conj"
	}
	if node.Rel == "dp" {
		if node == nLeft(Find(q /* $node/../node[@rel="dp"] */, &XPath{
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
		if Test(q /* $node/../node[@rel="hd" and @ud:pos=("AUX","ADP")] and not($node/../node[@rel="predc"]) */, &XPath{
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
								arg2: &Equal{
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
			return dependencyLabel(node.parent, q)
		}
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
		}) {
			if Test(q, /* $node/node[@rel="su" and @index and not(@word or @cat)]
				   (: or $node[@cat=("ti","oti")] :)
				   or $node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)]
				   or $node[@cat="oti"]/node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)] */&XPath{
					arg1: &Sort{
						arg1: &Or{
							arg1: &Or{
								arg1: &Collect{
									ARG: collect__child__node,
									arg1: &Variable{
										VAR: node,
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
														DATA: []interface{}{"su"},
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
													},
												},
												arg2: &Collect{
													ARG:  collect__attributes__index,
													arg1: &Node{},
												},
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
								arg2: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG: collect__child__node,
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
														DATA: []interface{}{"ti"},
														arg1: &Collect{
															ARG:  collect__attributes__cat,
															arg1: &Node{},
														},
													},
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
													DATA: []interface{}{"body"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
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
														DATA: []interface{}{"su"},
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
													},
												},
												arg2: &Collect{
													ARG:  collect__attributes__index,
													arg1: &Node{},
												},
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
							arg2: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG: collect__child__node,
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
														DATA: []interface{}{"oti"},
														arg1: &Collect{
															ARG:  collect__attributes__cat,
															arg1: &Node{},
														},
													},
												},
											},
										},
										arg2: &Predicate{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__cat,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"ti"},
													arg1: &Collect{
														ARG:  collect__attributes__cat,
														arg1: &Node{},
													},
												},
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
												DATA: []interface{}{"body"},
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
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
													DATA: []interface{}{"su"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
											arg2: &Collect{
												ARG:  collect__attributes__index,
												arg1: &Node{},
											},
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
					},
				}) {
				return "xcomp"
			}
			if Test(q /* $node/../@cat="np" */, &XPath{
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
							DATA: []interface{}{"np"},
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
			}) {
				return "acl" // v2: clausal dependents of nouns always acl
			}
			return "ccomp"
		}
		if Test(q /* $node/../node[@rel=("su","obj1") and (@pt or @cat)] */, &XPath{
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
									DATA: []interface{}{"su", "obj1"},
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
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if (node.Rel == "mod" || node.Rel == "pc" || node.Rel == "ld") && node.parent.Cat == "np" { // [detp niet veel] meer
		// modification of nomimal heads
		// pc and ld occur in nominalizations
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
		}) {
			return modLabelInsideNp(node, q)
		}
		if node == nLeft(Find(q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &XPath{
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
			return dependencyLabel(node.parent, q) // gapping, where this mod is the head
		}
		return "orphan"
	}
	if Test(q /* $node[@rel=("mod","pc","ld") and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")] */, &XPath{
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
								DATA: []interface{}{"mod", "pc", "ld"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG: collect__attributes__cat,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
							},
							arg2: &Elem{
								DATA: []interface{}{"sv1", "smain", "ssub", "inf", "ppres", "ppart", "oti", "ap", "advp"},
								arg1: &Collect{
									ARG: collect__attributes__cat,
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
	}) {
		// modification of verbal, adjectival heads
		// nb some oti's directly dominate (preceding) modifiers
		// [advp weg ermee ]
		if Test(q /* $node/../node[@rel=("hd","body") and (@pt or @cat)] */, &XPath{
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
									DATA: []interface{}{"hd", "body"},
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
		}) { // body for mods dangling outside cmp/body: maar niet om ...
			return labelVmod(node, q)
		}
		if Test(q /* $node/../node[@rel=("su","obj1","predc","vc") and (@pt or @cat)] */, &XPath{
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
									DATA: []interface{}{"su", "obj1", "predc", "vc"},
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
			return "orphan"
		}
		if Test(q /* $node[@rel="mod" and ../node[@rel=("pc","ld")]] */, &XPath{
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
									DATA: []interface{}{"mod"},
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
											DATA: []interface{}{"pc", "ld"},
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
			return "orphan"
		}
		if Test(q /* $node[@rel="mod" and ../node[@rel="mod"]/@begin < @begin] */, &XPath{
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
									DATA: []interface{}{"mod"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
							arg2: &Cmp{
								ARG: cmp__lt,
								arg1: &Collect{
									ARG: collect__attributes__begin,
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
													DATA: []interface{}{"mod"},
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
									ARG:  collect__attributes__begin,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) { // gapping with multiple mods
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping, where this mod is the head
	}
	if Test(q /* $node[@rel="mod" and ../@cat=("pp","detp","advp")] */, &XPath{
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
								DATA: []interface{}{"mod"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG: collect__attributes__cat,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
							},
							arg2: &Elem{
								DATA: []interface{}{"pp", "detp", "advp"},
								arg1: &Collect{
									ARG: collect__attributes__cat,
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
	}) {
		return "amod"
	}
	if Test(q /* $node[@rel="mod" and ../@cat=("cp", "whrel", "whq", "whsub")] */, &XPath{
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
								DATA: []interface{}{"mod"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG: collect__attributes__cat,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
							},
							arg2: &Elem{
								DATA: []interface{}{"cp", "whrel", "whq", "whsub"},
								arg1: &Collect{
									ARG: collect__attributes__cat,
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
	}) {
		// [cp net  [cmp als] [body de Belgische regering]], net wat het land nodig heeft
		if Test(q /* $node/../node[@rel="body"]/node[@rel="hd" and @ud:pos="VERB"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
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
									DATA: []interface{}{"VERB"},
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
		if node.udPos == "VERB" {
			return "xcomp"
		}
		if node.udPos != "" {
			return "advmod"
		}
		return "advcl"
	}
	if node.Rel == "rhd" || node.Rel == "whd" {
		if Test(q /* $node/../node[@rel="body"]//node/@index = $node/@index */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__index,
						arg1: &Collect{
							ARG: collect__descendant__node,
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
					},
					arg2: &Collect{
						ARG: collect__attributes__index,
						arg1: &Variable{
							VAR: node,
						},
					},
				},
			},
		}) {
			return nonLocalDependencyLabel(
				node,
				n1(Find(q /* ($node/../node[@rel="body"]//node[@index = $node/@index])[1] */, &XPath{
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
							arg2: &Elem{
								DATA: []interface{}{1},
							},
						},
					},
				})),
				q)
		}
		if node.Cat == "pp" {
			return "nmod" // onder wie michael boogerd
		}
		return "advmod" // [whq waarom jij]
	}
	if node.Rel == "body" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "--" {
		if node.udPos == "PUNCT" {
			if Test(q /* $node[not(../node[not(@ud:pos="PUNCT")]) and @begin = ../@begin] */, &XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: node,
						},
						arg2: &Sort{
							arg1: &And{
								arg1: &Function{
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
			}) {
				return "root" // just punctuation
			}
			return "punct"
		}
		if node.udPos == "SYM" || node.udPos == "X" {
			if Test(q /* $node/../node[@cat] */, &XPath{
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
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
						},
					},
				},
			}) {
				return "appos" // 1. Jantje is ziek  1-->appos??
			}
			return "root"
		}
		if node.Lemma == "\\" {
			return "punct" // hack for tagging errors in lassy small 250
		}
		/*
			if node.Spectype == "deeleigen" {
				return "punct" // hack for tagging errors in lassy small 250
			}
		*/
		if Test(q /* $node[@ud:pos="NUM" and ../node[@cat] ] */, &XPath{
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
									DATA: []interface{}{"NUM"},
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
		}) {
			return "parataxis" // dangling number 1.
		}
		if Test(q /* $node[@ud:pos="CCONJ" and ../node[@cat="smain" or @cat="conj"]] */, &XPath{
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
									DATA: []interface{}{"CCONJ"},
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
									arg1: &Or{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__cat,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"smain"},
												arg1: &Collect{
													ARG:  collect__attributes__cat,
													arg1: &Node{},
												},
											},
										},
										arg2: &Equal{
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
						},
					},
				},
			},
		}) {
			return "cc"
		}
		// sentence initial or final 'en'
		if Test(q /* $node[@ud:pos=("NOUN","PROPN","VERB") and ../node[@cat=("du","smain")]] */, &XPath{
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
									DATA: []interface{}{"NOUN", "PROPN", "VERB"},
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
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"du", "smain"},
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
		}) {
			return "parataxis" // dangling words
		}
		if Test(q /* count($node/../node[not(@ud:pos=("PUNCT","SYM","X"))]) < 2 */, &XPath{
			arg1: &Sort{
				arg1: &Cmp{
					ARG: cmp__lt,
					arg1: &Function{
						ARG: function__count__1__args,
						arg1: &Arg{
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG: collect__parent__type__node,
									arg1: &Variable{
										VAR: node,
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
														DATA: []interface{}{"PUNCT", "SYM", "X"},
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
					arg2: &Elem{
						DATA: []interface{}{2},
						arg1: &Function{
							ARG: function__count__1__args,
							arg1: &Arg{
								arg1: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG: collect__parent__type__node,
										arg1: &Variable{
											VAR: node,
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
															DATA: []interface{}{"PUNCT", "SYM", "X"},
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
			return "root" // only one non-punct/sym/foreign element in the string
		}
		if node.Cat == "mwu" {
			if node.Begin == node.parent.Begin && node.End == node.parent.End {
				return "root"
			}
			if Test(q /* $node/node[@ud:pos=("PUNCT","SYM")] */, &XPath{
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
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"PUNCT", "SYM"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
								},
							},
						},
					},
				},
			}) { // fix for mwu punctuation in Alpino output
				return "punct"
			}
			return "ERROR_NO_LABEL_--"
		}
		if Test(q /* $node[not(@ud:pos)]/../@rel="top" */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__rel,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Filter{
								arg1: &Variable{
									VAR: node,
								},
								arg2: &Sort{
									arg1: &Function{
										ARG: function__not__1__args,
										arg1: &Arg{
											arg1: &Sort{
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
					arg2: &Elem{
						DATA: []interface{}{"top"},
						arg1: &Collect{
							ARG: collect__attributes__rel,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Filter{
									arg1: &Variable{
										VAR: node,
									},
									arg2: &Sort{
										arg1: &Function{
											ARG: function__not__1__args,
											arg1: &Arg{
												arg1: &Sort{
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
		}) {
			return "root"
		}
		if Test(q /* $node[@ud:pos="PROPN" and not(../node[@cat]) ] */, &XPath{
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
									DATA: []interface{}{"PROPN"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
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
		}) {
			return "root" // Arthur .
		}
		if Test(q /* $node[@ud:pos=("ADP","ADV","ADJ","DET","PRON","CCONJ","NOUN","VERB","INTJ")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
					},
					arg2: &Sort{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"ADP", "ADV", "ADJ", "DET", "PRON", "CCONJ", "NOUN", "VERB", "INTJ"},
								arg1: &Collect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) {
			return "parataxis"
		}
		return "ERROR_NO_LABEL_--"
	}
	if node.Rel == "hd" {
		if node.udPos == "ADP" {
			if Test(q /* $node/../node[@rel="predc"] */, &XPath{
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
			}) {
				return "mark" // absolute met constructie -- analoog aan with X being Y
			}
			if Test(q /* not($node/../node[@rel="pc"]) */, &XPath{
				arg1: &Sort{
					arg1: &Function{
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
												DATA: []interface{}{"pc"},
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
				return "case" // er blijft weinig over van het lijk : over heads a predc and has pc as sister
			}
			return dependencyLabel(node.parent, q) // not sure about this one
		}
		if Test(q, /* $node[(@ud:pos=("ADJ","X","ADV") or @cat="mwu")
			   and ../@cat="pp"
			   and ../node[@rel=("obj1","se","vc")]] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: node,
						},
						arg2: &Sort{
							arg1: &And{
								arg1: &And{
									arg1: &Sort{
										arg1: &Or{
											arg1: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3apos,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"ADJ", "X", "ADV"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
												},
											},
											arg2: &Equal{
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
									arg2: &Equal{
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
												DATA: []interface{}{"obj1", "se", "vc"},
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
			if Test(q /* $node[../@rel="cnj" and ../node[@rel="obj1" and @index and not(@pt or @cat)]] */, &XPath{
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
										ARG: collect__attributes__rel,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
										},
									},
									arg2: &Elem{
										DATA: []interface{}{"cnj"},
										arg1: &Collect{
											ARG: collect__attributes__rel,
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
												arg2: &Collect{
													ARG:  collect__attributes__index,
													arg1: &Node{},
												},
											},
											arg2: &Function{
												ARG: function__not__1__args,
												arg1: &Arg{
													arg1: &Sort{
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
				return "conj"
			}
			return "case" // vol vertrouwen in, Ad U3... , op het gebied van
		}
		if Test(q /* $node/../node[@rel="hd"]/@begin < $node/@begin */, &XPath{
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
					arg2: &Collect{
						ARG: collect__attributes__begin,
						arg1: &Variable{
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
	return "ERROR_NO_LABEL"
}

func determineNominalModLabel(node *NodeType, q *Context) string {
	if Test(q /* $node/../node[@rel="hd" and (@ud:pos="VERB" or @ud:pos="ADJ")] */, &XPath{
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
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"VERB"},
										arg1: &Collect{
											ARG:  collect__attributes__ud_3apos,
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
										DATA: []interface{}{"ADJ"},
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
		return "obl"
	}
	return "nmod"
}

func subjectLabel(subj *NodeType, q *Context) string {
	pass := passiveSubject(subj, q)
	if Test(q, /* $subj[@cat=("whsub","ssub","ti","cp","oti")] or
		   $subj[@cat="conj" and node/@cat=("whsub","ssub","ti","cp","oti")] */&XPath{
			arg1: &Sort{
				arg1: &Or{
					arg1: &Filter{
						arg1: &Variable{
							VAR: subj,
						},
						arg2: &Sort{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"whsub", "ssub", "ti", "cp", "oti"},
									arg1: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
								},
							},
						},
					},
					arg2: &Filter{
						arg1: &Variable{
							VAR: subj,
						},
						arg2: &Sort{
							arg1: &And{
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
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG: collect__attributes__cat,
										arg1: &Collect{
											ARG:  collect__child__node,
											arg1: &Node{},
										},
									},
									arg2: &Elem{
										DATA: []interface{}{"whsub", "ssub", "ti", "cp", "oti"},
										arg1: &Collect{
											ARG: collect__attributes__cat,
											arg1: &Collect{
												ARG:  collect__child__node,
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
		return "csubj" + pass
	}
	// weather verbs and other expletive subject verbs
	if Test(q, /* $subj[@lemma="het"] and
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
		   ) */&XPath{
			arg1: &Sort{
				arg1: &And{
					arg1: &Filter{
						arg1: &Variable{
							VAR: subj,
						},
						arg2: &Sort{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__lemma,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"het"},
									arg1: &Collect{
										ARG:  collect__attributes__lemma,
										arg1: &Node{},
									},
								},
							},
						},
					},
					arg2: &Sort{
						arg1: &Or{
							arg1: &Or{
								arg1: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG: collect__parent__type__node,
										arg1: &Variable{
											VAR: subj,
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
													ARG:  collect__attributes__lemma,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"dooien", "gieten", "hagelen", "miezeren",
														"misten", "motregenen", "onweren", "plenzen",
														"regenen", "sneeuwen", "stormen", "stortregenen",
														"ijzelen", "vriezen", "weerlichten", "winteren",
														"zomeren"},
													arg1: &Collect{
														ARG:  collect__attributes__lemma,
														arg1: &Node{},
													},
												},
											},
										},
									},
								},
								arg2: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG: collect__parent__type__node,
										arg1: &Variable{
											VAR: subj,
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
												arg2: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__lemma,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"ontbreken"},
														arg1: &Collect{
															ARG:  collect__attributes__lemma,
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
																DATA: []interface{}{"pc"},
																arg1: &Collect{
																	ARG:  collect__attributes__rel,
																	arg1: &Node{},
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
																	arg2: &Equal{
																		ARG: equal__is,
																		arg1: &Collect{
																			ARG:  collect__attributes__lemma,
																			arg1: &Node{},
																		},
																		arg2: &Elem{
																			DATA: []interface{}{"aan"},
																			arg1: &Collect{
																				ARG:  collect__attributes__lemma,
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
							arg2: &Filter{
								arg1: &Variable{
									VAR: subj,
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
													ARG: collect__descendant__or__self__type__node,
													arg1: &Collect{
														ARG: collect__child__node,
														arg1: &Collect{
															ARG:  collect__parent__type__node,
															arg1: &Node{},
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
															DATA: []interface{}{"sup"},
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
			},
		}) {
		return "expl" + pass
	}
	return "nsubj" + pass
}

// recursive
func passiveSubject(subj *NodeType, q *Context) string {
	if depthCheck(q, "passiveSubject") {
		return "ERROR_NO_PASSIVE_SUBJECT"
	}

	aux := auxiliary1(n1(Find(q /* ($subj/../node[@rel="hd"])[1] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: subj,
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
				arg2: &Elem{
					DATA: []interface{}{1},
				},
			},
		},
	})), q)
	if aux == "aux:pass" { // de carriere had gered kunnen worden
		return ":pass"
	}
	if aux == "aux" && Test(q /* $subj/@index = $subj/../node[@rel="vc"]/node[@rel="su"]/@index */, &XPath{
		arg1: &Sort{
			arg1: &Equal{
				ARG: equal__is,
				arg1: &Collect{
					ARG: collect__attributes__index,
					arg1: &Variable{
						VAR: subj,
					},
				},
				arg2: &Collect{
					ARG: collect__attributes__index,
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Variable{
									VAR: subj,
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
			},
		},
	}) {
		return passiveSubject(n1(Find(q /* $subj/../node[@rel="vc"]/node[@rel="su"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								VAR: subj,
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
		})), q)
	}
	return ""
}

func detLabel(node *NodeType, q *Context) string {
	// zijn boek, diens boek, ieders boek, aller landen, Ron's probleem, Fidel Castro's belang
	if Test(q, /* $node[@ud:pos = "PRON" and @vwtype="bez"] or
		   $node[@ud:pos = ("PRON","PROPN") and @naamval="gen"] or
		   $node[@cat="mwu" and node[@spectype="deeleigen"]] */&XPath{
			arg1: &Sort{
				arg1: &Or{
					arg1: &Or{
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
											DATA: []interface{}{"PRON"},
											arg1: &Collect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &Node{},
											},
										},
									},
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__vwtype,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"bez"},
											arg1: &Collect{
												ARG:  collect__attributes__vwtype,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
						arg2: &Filter{
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
											DATA: []interface{}{"PRON", "PROPN"},
											arg1: &Collect{
												ARG:  collect__attributes__ud_3apos,
												arg1: &Node{},
											},
										},
									},
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__naamval,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"gen"},
											arg1: &Collect{
												ARG:  collect__attributes__naamval,
												arg1: &Node{},
											},
										},
									},
								},
							},
						},
					},
					arg2: &Filter{
						arg1: &Variable{
							VAR: node,
						},
						arg2: &Sort{
							arg1: &And{
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
								arg2: &Collect{
									ARG:  collect__child__node,
									arg1: &Node{},
									arg2: &Predicate{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__spectype,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"deeleigen"},
												arg1: &Collect{
													ARG:  collect__attributes__spectype,
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
		return "nmod:poss"
	}
	if Test(q /* $node/@ud:pos = ("DET","PROPN","NOUN","ADJ","PRON","ADV","X") */, &XPath{
		arg1: &Sort{
			arg1: &Equal{
				ARG: equal__is,
				arg1: &Collect{
					ARG: collect__attributes__ud_3apos,
					arg1: &Variable{
						VAR: node,
					},
				},
				arg2: &Elem{
					DATA: []interface{}{"DET", "PROPN", "NOUN", "ADJ", "PRON", "ADV", "X"},
					arg1: &Collect{
						ARG: collect__attributes__ud_3apos,
						arg1: &Variable{
							VAR: node,
						},
					},
				},
			},
		},
	}) {
		return "det" // meer // genoeg // the
	}
	if Test(q /* $node/@cat = ("mwu","np","pp","ap","detp","smain") */, &XPath{
		arg1: &Sort{
			arg1: &Equal{
				ARG: equal__is,
				arg1: &Collect{
					ARG: collect__attributes__cat,
					arg1: &Variable{
						VAR: node,
					},
				},
				arg2: &Elem{
					DATA: []interface{}{"mwu", "np", "pp", "ap", "detp", "smain"},
					arg1: &Collect{
						ARG: collect__attributes__cat,
						arg1: &Variable{
							VAR: node,
						},
					},
				},
			},
		},
	}) {
		return "det"
	}
	// tussen 5 en 6 .., needs more principled solution
	// nog meer mensen dan anders
	// hetzelfde tijdstip als anders , nogal wat,
	// hij heeft ik weet niet hoeveel boeken
	if node.udPos == "NUM" || node.udPos == "SYM" {
		return "nummod"
	}
	if node.Cat == "conj" {
		// TODO: als ik hier 1 vervang door last() dan verdwijnen de verschillen met Saxon, maar het moet echt 1 zijn
		if Test(q /* $node/node[@rel="cnj"][1]/@ud:pos="NUM" */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__ud_3apos,
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
											DATA: []interface{}{"cnj"},
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
					arg2: &Elem{
						DATA: []interface{}{"NUM"},
						arg1: &Collect{
							ARG: collect__attributes__ud_3apos,
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
												DATA: []interface{}{"cnj"},
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
					},
				},
			},
		}) {
			return "nummod"
		}
		return "det"
	}
	return "ERROR_NO_LABEL_DET"
}

func modLabelInsideNp(node *NodeType, q *Context) string {
	if Test(q /* $node[@cat="pp"]/node[@rel="vc"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
		return "acl" // pp containing a clause
	}
	if Test(q /* $node[@ud:pos="ADJ" or @cat="ap" or node[@cat="conj" and node[@ud:pos="ADJ" or @cat="ap"] ]] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
				},
				arg2: &Sort{
					arg1: &Or{
						arg1: &Or{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"ADJ"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
								},
							},
							arg2: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"ap"},
									arg1: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
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
									arg2: &Collect{
										ARG:  collect__child__node,
										arg1: &Node{},
										arg2: &Predicate{
											arg1: &Or{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"ADJ"},
														arg1: &Collect{
															ARG:  collect__attributes__ud_3apos,
															arg1: &Node{},
														},
													},
												},
												arg2: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__cat,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"ap"},
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
				},
			},
		},
	}) {
		return "amod"
	}
	if Test(q /* $node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM","INTJ") ] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
				},
				arg2: &Sort{
					arg1: &Or{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"pp", "np", "conj", "mwu"},
								arg1: &Collect{
									ARG:  collect__attributes__cat,
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
								DATA: []interface{}{"NOUN", "PRON", "PROPN", "X", "PUNCT", "SYM", "INTJ"},
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
		return "nmod"
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if Test(q /* $node[@cat="detp"]/node[@rel="hd" and @ud:pos="NUM"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
								DATA: []interface{}{"detp"},
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
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
								DATA: []interface{}{"NUM"},
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
		return "nummod"
	}
	if node.Cat == "detp" {
		return "det" // [detp niet veel] meer error?
	}
	if node.Cat == "rel" || node.Cat == "whrel" {
		return "acl:relcl"
	}
	// v2 added relcl -- whrel= met name waar ...
	if Test(q /* $node[@cat="cp"]/node[@rel="body" and (@ud:pos = ("NOUN","PROPN") or @cat=("np","conj"))] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
								DATA: []interface{}{"cp"},
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
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
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"NOUN", "PROPN"},
										arg1: &Collect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &Node{},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"np", "conj"},
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
	}) {
		return "nmod"
	}
	// zijn loopbaan [CP als schrijver]
	if Test(q /* $node[@cat=("cp","sv1","smain","ppres","ppart","inf","ti","oti","du","whq") or @ud:pos="SCONJ"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
				},
				arg2: &Sort{
					arg1: &Or{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"cp", "sv1", "smain", "ppres", "ppart", "inf", "ti", "oti", "du", "whq"},
								arg1: &Collect{
									ARG:  collect__attributes__cat,
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
								DATA: []interface{}{"SCONJ"},
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
		return "acl"
	}
	// oa zinnen tussen haakjes
	if Test(q /* $node[@ud:pos= ("ADV","ADP","VERB","CCONJ") or @cat="advp"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
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
								DATA: []interface{}{"ADV", "ADP", "VERB", "CCONJ"},
								arg1: &Collect{
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
								},
							},
						},
						arg2: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"advp"},
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
	}) {
		return "amod"
	}
	// VERB= aanstormend etc -> amod, ADV = nagenoeg alle prijzen, slechts 4 euro --> amod
	// CCONJ = opdrachten zoals:   --> amod
	if node.Rel == "det" {
		return "det" // empty determiners in gapping?
	}
	if node.Index > 0 {
		return "ERROR_INDEX_NMOD"
	}
	return "ERROR_NO_LABEL_NMOD"
}

func labelVmod(node *NodeType, q *Context) string {
	if Test(q /* $node[@cat="pp"]/node[@rel="vc"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
		return "advcl"
	}
	if Test(q, /* $node[ (  node[@rel="hd" and @lemma="door"]
		    or (@pt="bw" and ends-with(@lemma,"door"))
		    )
		    and parent::node[@cat="ppart"]/../node[@rel="hd" and @lemma=("zijn","worden")]
		    and ../../node[@rel="su"]/@index = ../node[@rel="obj1"]/@index
		] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
					},
					arg2: &Sort{
						arg1: &And{
							arg1: &And{
								arg1: &Sort{
									arg1: &Or{
										arg1: &Collect{
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
													arg2: &Equal{
														ARG: equal__is,
														arg1: &Collect{
															ARG:  collect__attributes__lemma,
															arg1: &Node{},
														},
														arg2: &Elem{
															DATA: []interface{}{"door"},
															arg1: &Collect{
																ARG:  collect__attributes__lemma,
																arg1: &Node{},
															},
														},
													},
												},
											},
										},
										arg2: &Sort{
											arg1: &And{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__pt,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"bw"},
														arg1: &Collect{
															ARG:  collect__attributes__pt,
															arg1: &Node{},
														},
													},
												},
												arg2: &Function{
													ARG: function__ends__with__2__args,
													arg1: &Arg{
														arg1: &Arg{
															arg1: &Sort{
																arg1: &Collect{
																	ARG:  collect__attributes__lemma,
																	arg1: &Node{},
																},
															},
														},
														arg2: &Elem{
															DATA: []interface{}{"door"},
														},
													},
												},
											},
										},
									},
								},
								arg2: &Collect{
									ARG: collect__child__node,
									arg1: &Collect{
										ARG: collect__parent__type__node,
										arg1: &Collect{
											ARG:  collect__parent__node,
											arg1: &Node{},
											arg2: &Predicate{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__cat,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"ppart"},
														arg1: &Collect{
															ARG:  collect__attributes__cat,
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
													ARG:  collect__attributes__lemma,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"zijn", "worden"},
													arg1: &Collect{
														ARG:  collect__attributes__lemma,
														arg1: &Node{},
													},
												},
											},
										},
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
											ARG: collect__parent__type__node,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
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
								arg2: &Collect{
									ARG: collect__attributes__index,
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
							},
						},
					},
				},
			},
		}) {
		return "obl:agent"
		/*
			but NOT: door rookontwikkeling om het leven gekomen
			-- already filtered by constraint of su/obj1 control
			NOT: bij Bakema is een stoeptegel door de ruit gegooid
			NO/YES: hierdoor werd Prince door het grote publiek ontdekt
		*/
	}
	if Test(q /* $node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM") ] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: node,
				},
				arg2: &Sort{
					arg1: &Or{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__cat,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"pp", "np", "conj", "mwu"},
								arg1: &Collect{
									ARG:  collect__attributes__cat,
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
								DATA: []interface{}{"NOUN", "PRON", "PROPN", "X", "PUNCT", "SYM"},
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
		return "obl"
	}
	if Test(q /* $node[@cat=("cp","sv1","smain","ssub","ppres","ppart","ti","oti","inf","du","whq","whrel","rel")] */, &XPath{
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
							DATA: []interface{}{"cp", "sv1", "smain", "ssub", "ppres", "ppart", "ti", "oti", "inf", "du", "whq", "whrel", "rel"},
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
		return "advcl"
	}
	if Test(q, /* $node[@ud:pos= ("ADJ","ADV","ADP","VERB","SCONJ","INTJ")
		   or @cat=("advp","ap")
		   or (@cat="conj" and node/@ud:pos="ADV")] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
					},
					arg2: &Sort{
						arg1: &Or{
							arg1: &Or{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"ADJ", "ADV", "ADP", "VERB", "SCONJ", "INTJ"},
										arg1: &Collect{
											ARG:  collect__attributes__ud_3apos,
											arg1: &Node{},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__cat,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"advp", "ap"},
										arg1: &Collect{
											ARG:  collect__attributes__cat,
											arg1: &Node{},
										},
									},
								},
							},
							arg2: &Sort{
								arg1: &And{
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
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG: collect__attributes__ud_3apos,
											arg1: &Collect{
												ARG:  collect__child__node,
												arg1: &Node{},
											},
										},
										arg2: &Elem{
											DATA: []interface{}{"ADV"},
											arg1: &Collect{
												ARG: collect__attributes__ud_3apos,
												arg1: &Collect{
													ARG:  collect__child__node,
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
		return "advmod" // niet of nauwelijks
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if node.Index > 0 {
		return "ERROR_INDEX_VMOD"
	}
	return "ERROR_NO_LABEL_VMOD"
}

// this function is now also used to distribute dependents in coordination in Enhanced UD , so lot more rels and contexts are possible
// and passives, as in " hun taal werd gediscrimineerd en verboden"
func nonLocalDependencyLabel(head, gap *NodeType, q *Context) string {
	if gap.Rel == "su" {
		return subjectLabel(gap, q)
	}
	if gap.Rel == "obj1" {
		if head.Rel == "su" {
			return "nsubj:pass"
		}
		return "obj" // ppart coordination -- see comment above
	}
	if gap.Rel == "obj2" {
		if head.udPos == "ADV" {
			return "advmod"
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
		if Test(q /* $head/node[@rel="obj1"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						VAR: head,
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
		}) {
			return "nmod"
		}
		if Test(q /* $head[@ud:pos=("ADV", "ADP") or @cat=("advp","ap")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: head,
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
									DATA: []interface{}{"ADV", "ADP"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
								},
							},
							arg2: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"advp", "ap"},
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
		}) {
			return "advmod" // waar precies zit je ..
		}
		return "ERROR_NO_LABEL_INDEX_PC"
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
	if Test(q /* $gap[@rel="mod" and ../@cat=("np","pp")] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: gap,
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
								DATA: []interface{}{"mod"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG: collect__attributes__cat,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
							},
							arg2: &Elem{
								DATA: []interface{}{"np", "pp"},
								arg1: &Collect{
									ARG: collect__attributes__cat,
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
	}) { // voornamelijk in kloosters en door vrouwen
		return modLabelInsideNp(head, q)
	}
	if Test(q /* $gap[@rel="mod" and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: gap,
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
								DATA: []interface{}{"mod"},
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
							},
						},
						arg2: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG: collect__attributes__cat,
								arg1: &Collect{
									ARG:  collect__parent__type__node,
									arg1: &Node{},
								},
							},
							arg2: &Elem{
								DATA: []interface{}{"sv1", "smain", "ssub", "inf", "ppres", "ppart", "oti", "ap", "advp"},
								arg1: &Collect{
									ARG: collect__attributes__cat,
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
	}) {
		return labelVmod(head, q)
	}
	if gap.Rel == "mod" || gap.Rel == "spec" { // spec only used for funny coord
		if Test(q /* $head[@cat=("pp","np") or @ud:pos=("NOUN","PRON")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: head,
					},
					arg2: &Sort{
						arg1: &Or{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"pp", "np"},
									arg1: &Collect{
										ARG:  collect__attributes__cat,
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
									DATA: []interface{}{"NOUN", "PRON"},
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
			return "nmod"
		}
		if Test(q /* $head[@ud:pos=("ADV","ADP") or @cat= ("advp","ap","mwu","conj")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: head,
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
									DATA: []interface{}{"ADV", "ADP"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3apos,
										arg1: &Node{},
									},
								},
							},
							arg2: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"advp", "ap", "mwu", "conj"},
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
		}) {
			return "advmod" // hoe vaak -- AP, daar waar, waar en wanneer, voor als rhd
		}
		return "ERROR_NO_LABEL_INDEX_MOD"
	}
	if gap.Rel == "det" {
		return detLabel(head, q)
	}
	if Test(q /* $gap[@rel="hd"] and $head[@ud:pos=("ADV","ADP")] */, &XPath{
		arg1: &Sort{
			arg1: &And{
				arg1: &Filter{
					arg1: &Variable{
						VAR: gap,
					},
					arg2: &Sort{
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
				arg2: &Filter{
					arg1: &Variable{
						VAR: head,
					},
					arg2: &Sort{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"ADV", "ADP"},
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
	}) { // waaronder A, B, en C
		return "case"
	}
	if gap.Rel == "du" || gap.Rel == "dp" {
		return "parataxis"
	}
	return "ERROR_NO_LABEL_INDEX"
}
