//
// GENERATED FILE -- DO NOT EDIT
//

package main

import (
	"sort"
	"strings"
)

type DepType struct {
	head int
	dep  string
}

func enhancedDependencies(q *Context) {

	changed := reconstructEmptyHead(q)

	// add_Edependency_relations
	for _, node := range q.ptnodes {
		// Edependency_relation
		if changed {
			q.depth = 0
			node.udERelation = dependencyLabel(node, q)
			q.depth = 0
			node.udEHeadPosition = externalHeadPosition(node, q)
		} else {
			node.udERelation = node.udRelation
			node.udEHeadPosition = node.udHeadPosition
		}
	}

	for _, node := range q.ptnodes {
		enhancedDependencies1(node, q)
	}
}

func enhancedDependencies1(node *NodeType, q *Context) {
	// iobj2 control : de commissie kan de raad aanbevelen/adviseren/ X te doen
	// rhd : een levend visje dat doorgeslikt moet worden
	q.depth = 0
	var enhanced []DepType
	for {

		if Test(q /* $node[@ud:ERelation=("nsubj","obj","iobj","nsubj:pass")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: node,
					},
					arg2: &Sort{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__ud_3aERelation,
								arg1: &Node{},
							},
							arg2: &Elem{
								DATA: []interface{}{"nsubj", "obj", "iobj", "nsubj:pass"},
								arg1: &Collect{
									ARG:  collect__attributes__ud_3aERelation,
									arg1: &Node{},
								},
							},
						},
					},
				},
			},
		}) { // TODO: klopt dit? exists binnen [ ]
			so := Find(q,
				/* $node/ancestor::node/node[@rel=("su","obj1","obj2") and local:internal_head_position(.) = $node/@end ]/@index */ &XPath{
					arg1: &Sort{
						arg1: &Collect{
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
									arg1: &And{
										arg1: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
											arg2: &Elem{
												DATA: []interface{}{"su", "obj1", "obj2"},
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
											},
										},
										arg2: &Equal{
											ARG: equal__is,
											arg1: &Function{
												ARG: function__local__internal__head__position__1__args,
												arg1: &Arg{
													arg1: &Sort{
														arg1: &Node{},
													},
												},
											},
											arg2: &Collect{
												ARG: collect__attributes__end,
												arg1: &Variable{
													VAR: node,
												},
											},
										},
									},
								},
							},
						},
					},
				})
			if len(so) > 0 {
				soIndex := i1(so)
				enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
				enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                    // self
				enhanced = append(enhanced, distributeConjuncts(node, q)...)                                    // self
				enhanced = append(enhanced, distributeDependents(node, q)...)                                   // self
				enhanced = append(enhanced, xcompControl(node, q, soIndex)...)
				enhanced = append(enhanced, upstairsControl(node, q, soIndex)...)
				enhanced = append(enhanced, passiveVpControl(node, q, soIndex)...)
				break
			}
		}

		rhd := Find(q,
			/* $node/ancestor::node/node[@rel="rhd" and local:internal_head_position(.) = $node/@end ]/@index */ &XPath{
				arg1: &Sort{
					arg1: &Collect{
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
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"rhd"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Function{
											ARG: function__local__internal__head__position__1__args,
											arg1: &Arg{
												arg1: &Sort{
													arg1: &Node{},
												},
											},
										},
										arg2: &Collect{
											ARG: collect__attributes__end,
											arg1: &Variable{
												VAR: node,
											},
										},
									},
								},
							},
						},
					},
				},
			})
		if len(rhd) > 0 {
			rhdIndex := i1(rhd)
			rhdNp := Find(q /* $node/ancestor::node[@cat="np" and node[@rel="mod"]/node[@rel="rhd"]/@index = $rhdIndex] */, &XPath{
				arg1: &Sort{
					arg1: &Collect{
						ARG: collect__ancestors__node,
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
										DATA: []interface{}{"np"},
										arg1: &Collect{
											ARG:  collect__attributes__cat,
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
															DATA: []interface{}{"mod"},
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
														DATA: []interface{}{"rhd"},
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
													},
												},
											},
										},
									},
									arg2: &Variable{
										VAR: rhdIndex,
									},
								},
							},
						},
					},
				},
			})
			// de enige _i die voldoet aan de eisen -- make sure empty heads are covered as well
			if len(rhdNp) > 0 {
				enhanced = []DepType{DepType{head: internalHeadPositionWithGapping(rhdNp, q), dep: "ref"}} // rhdref
				enhanced = append(enhanced, xcompControl(node, q, rhdIndex)...)
				enhanced = append(enhanced, passiveVpControl(node, q, rhdIndex)...)
				break
			}
			// if there is no antecedent, lets keep the basic relation
			enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                    // self
			enhanced = append(enhanced, distributeConjuncts(node, q)...)                                    // self
			enhanced = append(enhanced, distributeDependents(node, q)...)                                   // self
			enhanced = append(enhanced, xcompControl(node, q, rhdIndex)...)
			enhanced = append(enhanced, passiveVpControl(node, q, rhdIndex)...)
			break
		}

		relSister := Find(q /* ($node/../node[@rel="mod" and @cat="rel"]/node[@rel="rhd"]/@index)[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__attributes__index,
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
													ARG:  collect__attributes__cat,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"rel"},
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
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__rel,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"rhd"},
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
					arg2: &Sort{
						arg1: &Function{
							ARG: function__first__0__args,
						},
					},
				},
			},
		})
		if len(relSister) > 0 {
			relSisterIndex := i1(relSister)
			enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                    // self
			enhanced = append(enhanced, distributeConjuncts(node, q)...)                                    // self
			enhanced = append(enhanced, distributeDependents(node, q)...)                                   // self
			enhanced = append(enhanced, xcompControl(node, q, relSisterIndex)...)
			enhanced = append(enhanced, passiveVpControl(node, q, relSisterIndex)...)
			break
		}

		// node.udHeadPosition == UNDERSCORE als resultaat van reconstructEmptyHead()
		if node.udHeadPosition >= 0 || node.udHeadPosition == UNDERSCORE {
			enhanced = []DepType{DepType{head: node.udEHeadPosition, dep: enhanceDependencyLabel(node, q)}} // self
			enhanced = append(enhanced, anaphoricRelpronoun(node, q)...)                                    // self
			enhanced = append(enhanced, distributeConjuncts(node, q)...)                                    // self
			enhanced = append(enhanced, distributeDependents(node, q)...)                                   // self
			break
		}

		enhanced = []DepType{}
		break
	}

	sort.Slice(enhanced, func(i, j int) bool {
		if enhanced[i].head != enhanced[j].head {
			return enhanced[i].head < enhanced[j].head
		}
		return enhanced[i].dep < enhanced[j].dep
	})
	for i := 1; i < len(enhanced); i++ {
		if enhanced[i].head == enhanced[i-1].head && enhanced[i].dep == enhanced[i-1].dep {
			enhanced = append(enhanced[:i], enhanced[1+i:]...)
			i--
		}
	}
	ss := make([]string, 0, len(enhanced))
	for _, e := range enhanced {
		if e.dep != "" {
			ss = append(ss, number(e.head)+":"+e.dep)
		}
	}
	node.udEnhanced = strings.Join(ss, "|")

}

func join(a, b string) string {
	if b == "" {
		return a
	}
	return a + ":" + b
}

func enhanceDependencyLabel(node *NodeType, q *Context) string {
	label := node.udERelation
	if label == "conj" {
		if crd := n1(Find(q, /* ($node/ancestor::node[@cat="conj" and
			   not(.//node[@cat="conj"]//node/@begin = $node/@begin)]/node[@rel="crd"])[1] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Sort{
							arg1: &Collect{
								ARG: collect__child__node,
								arg1: &Collect{
									ARG: collect__ancestors__node,
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
													DATA: []interface{}{"conj"},
													arg1: &Collect{
														ARG:  collect__attributes__cat,
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
																ARG: collect__attributes__begin,
																arg1: &Collect{
																	ARG: collect__descendant__node,
																	arg1: &Collect{
																		ARG: collect__child__node,
																		arg1: &Collect{
																			ARG:  collect__descendant__or__self__type__node,
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
											DATA: []interface{}{"crd"},
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
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
			})); crd != noNode {
			if crd.Lemma != "" {
				return join(label, enhancedLemmaString1(crd, q))
			}
			if crd.Cat == "mwu" {
				return join(label, enhancedLemmaString1(n1(Find(q /* ($crd/node[@rel="mwp"])[1] */, &XPath{
					arg1: &Sort{
						arg1: &Filter{
							arg1: &Sort{
								arg1: &Collect{
									ARG: collect__child__node,
									arg1: &Variable{
										VAR: crd,
									},
									arg2: &Predicate{
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
				})), q))
			}
			return "ERROR_empty_eud_label"
		}
	}

	if label == "nmod" || label == "obl" {
		if casee := n1(Find(q /* ($q.varptnodes[@ud:ERelation="case" and @ud:EHeadPosition=$node/@end])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Filter{
							arg1: &Variable{
								VAR: q.varptnodes,
							},
							arg2: &Sort{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__ud_3aERelation,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"case"},
											arg1: &Collect{
												ARG:  collect__attributes__ud_3aERelation,
												arg1: &Node{},
											},
										},
									},
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__ud_3aEHeadPosition,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG: collect__attributes__end,
											arg1: &Variable{
												VAR: node,
											},
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
		})); casee != noNode {
			return join(label, enhancedLemmaString1(casee, q))
		}
	}

	if label == "advcl" || label == "acl" {
		if mark := n1(Find(q /* ($q.varptnodes[@ud:ERelation=("mark","case") and @ud:EHeadPosition=$node/@end])[1] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Filter{
							arg1: &Variable{
								VAR: q.varptnodes,
							},
							arg2: &Sort{
								arg1: &And{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__ud_3aERelation,
											arg1: &Node{},
										},
										arg2: &Elem{
											DATA: []interface{}{"mark", "case"},
											arg1: &Collect{
												ARG:  collect__attributes__ud_3aERelation,
												arg1: &Node{},
											},
										},
									},
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__ud_3aEHeadPosition,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG: collect__attributes__end,
											arg1: &Variable{
												VAR: node,
											},
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
		})); mark != noNode {
			return join(label, enhancedLemmaString1(mark, q))
		}
	}

	if label != "" {
		return label
	}

	return "ERROR_empty_eud_label"
}

func anaphoricRelpronoun(node *NodeType, q *Context) []DepType {
	// works voor waar, and last() picks waar in 'daar waar' cases
	// dont add anything for hij werd voorzitter, wat hij nog steeds is (otherwise self-reference)
	// for loop ensures correct result if N has 2 acl:relcl dependents
	result := []DepType{}
	for _, a := range Find(q, /* $node/ancestor::node[@cat="np" and local:internal_head_position(.) = $node/@end]/
		   node[@rel="mod"]/node[@rel="rhd"]/descendant-or-self::node[@pt="vnw" and not(@ud:HeadPosition = $node/@end)][last()] */&XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__descendant__or__self__node,
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Collect{
								ARG: collect__ancestors__node,
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
												DATA: []interface{}{"np"},
												arg1: &Collect{
													ARG:  collect__attributes__cat,
													arg1: &Node{},
												},
											},
										},
										arg2: &Equal{
											ARG: equal__is,
											arg1: &Function{
												ARG: function__local__internal__head__position__1__args,
												arg1: &Arg{
													arg1: &Sort{
														arg1: &Node{},
													},
												},
											},
											arg2: &Collect{
												ARG: collect__attributes__end,
												arg1: &Variable{
													VAR: node,
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
						arg2: &Predicate{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__rel,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"rhd"},
									arg1: &Collect{
										ARG:  collect__attributes__rel,
										arg1: &Node{},
									},
								},
							},
						},
					},
					arg2: &Predicate{
						arg1: &Predicate{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__pt,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"vnw"},
										arg1: &Collect{
											ARG:  collect__attributes__pt,
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
													ARG:  collect__attributes__ud_3aHeadPosition,
													arg1: &Node{},
												},
												arg2: &Collect{
													ARG: collect__attributes__end,
													arg1: &Variable{
														VAR: node,
													},
												},
											},
										},
									},
								},
							},
						},
						arg2: &Function{
							ARG: function__last__0__args,
						},
					},
				},
			},
		}) {
		anrel := a.(*NodeType)
		var label string
		if r := anrel.udRelation; r == "nsubj" || r == "nsubj:pass" {
			label = r + ":relsubj"
		} else if r == "obj" || r == "obl" {
			label = r + ":relobj"
		} else {
			label = r
		}
		result = append(result, DepType{head: anrel.udHeadPosition, dep: label})
	}
	return result
}

// Glastra en Terlouw verzonnen een list --> nsubj(verzonnen,Glastra) nsubj(verzonnen,Terlouw)
func distributeConjuncts(node *NodeType, q *Context) []DepType {
	if node.udRelation == "conj" {
		coordHead := n1(Find(q, /* $q.varallnodes[@end = $node/@ud:HeadPosition
			   and @ud:Relation=("amod","appos","nmod","nsubj","nsubj:pass","nummod","obj","iobj","obl","obl:agent","advcl")] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							VAR: q.varallnodes,
						},
						arg2: &Sort{
							arg1: &And{
								arg1: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__end,
										arg1: &Node{},
									},
									arg2: &Collect{
										ARG: collect__attributes__ud_3aHeadPosition,
										arg1: &Variable{
											VAR: node,
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__ud_3aRelation,
										arg1: &Node{},
									},
									arg2: &Elem{
										DATA: []interface{}{"amod", "appos", "nmod", "nsubj", "nsubj:pass", "nummod", "obj", "iobj", "obl", "obl:agent", "advcl"},
										arg1: &Collect{
											ARG:  collect__attributes__ud_3aRelation,
											arg1: &Node{},
										},
									},
								},
							},
						},
					},
				},
			}))
		if coordHead != noNode {
			// in A en B vs in A en naast B --> use enh_dep_label($node) in the latter case...
			depLabel := enhanceDependencyLabel(coordHead, q)
			return []DepType{DepType{head: coordHead.udHeadPosition, dep: depLabel}}
		}
	}
	return []DepType{}
}

// de onrust kan een reis vertragen of frustreren  --> obj(vertragen,reis) obj(frustreren,reis)
// todo: passives ze werd ontmanteld en verkocht  su coindexed with two obj1
// done: phrases [np_i [een scoutskameraad] werd .. en _i zocht hem op]
// idem: de hond was gebaseerd op Lassy en verscheen onder de naam Wirel nsubj:pass in conj1, nsubj in conj 2
func distributeDependents(node *NodeType, q *Context) []DepType {
	var phrase *NodeType
	if node.Rel == "hd" {
		if Test(q /* $node/../../@cat="pp" */, &XPath{
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
						DATA: []interface{}{"pp"},
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
		}) { // door het schilderij
			phrase = node.parent.parent
		} else {
			phrase = node.parent
		}
	} else {
		if Test(q /* $node[@rel="mwp" and @begin = ../@begin] */, &XPath{
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
									DATA: []interface{}{"mwp"},
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
			if Test(q /* $node[ ../@rel="obj1" and ../../@cat="pp"] */, &XPath{
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
										DATA: []interface{}{"obj1"},
										arg1: &Collect{
											ARG: collect__attributes__rel,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
									},
								},
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG: collect__attributes__cat,
										arg1: &Collect{
											ARG: collect__parent__type__node,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
									},
									arg2: &Elem{
										DATA: []interface{}{"pp"},
										arg1: &Collect{
											ARG: collect__attributes__cat,
											arg1: &Collect{
												ARG: collect__parent__type__node,
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
			}) {
				phrase = node.parent.parent
			} else {
				if Test(q /* $node[../@rel="hd" and ../../@rel="obj1" and ../../../@cat="pp"] */, &XPath{
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
												DATA: []interface{}{"hd"},
												arg1: &Collect{
													ARG: collect__attributes__rel,
													arg1: &Collect{
														ARG:  collect__parent__type__node,
														arg1: &Node{},
													},
												},
											},
										},
										arg2: &Equal{
											ARG: equal__is,
											arg1: &Collect{
												ARG: collect__attributes__rel,
												arg1: &Collect{
													ARG: collect__parent__type__node,
													arg1: &Collect{
														ARG:  collect__parent__type__node,
														arg1: &Node{},
													},
												},
											},
											arg2: &Elem{
												DATA: []interface{}{"obj1"},
												arg1: &Collect{
													ARG: collect__attributes__rel,
													arg1: &Collect{
														ARG: collect__parent__type__node,
														arg1: &Collect{
															ARG:  collect__parent__type__node,
															arg1: &Node{},
														},
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
												ARG: collect__parent__type__node,
												arg1: &Collect{
													ARG: collect__parent__type__node,
													arg1: &Collect{
														ARG:  collect__parent__type__node,
														arg1: &Node{},
													},
												},
											},
										},
										arg2: &Elem{
											DATA: []interface{}{"pp"},
											arg1: &Collect{
												ARG: collect__attributes__cat,
												arg1: &Collect{
													ARG: collect__parent__type__node,
													arg1: &Collect{
														ARG: collect__parent__type__node,
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
					phrase = node.parent.parent.parent // in en rond het Hoofstedelijk Gewest --> do not distribute Hoofdstedelijk
				} else {
					if Test(q /* $node[ ../@rel="hd" and not( ../../@cat="pp") ] */, &XPath{
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
												DATA: []interface{}{"hd"},
												arg1: &Collect{
													ARG: collect__attributes__rel,
													arg1: &Collect{
														ARG:  collect__parent__type__node,
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
															ARG: collect__attributes__cat,
															arg1: &Collect{
																ARG: collect__parent__type__node,
																arg1: &Collect{
																	ARG:  collect__parent__type__node,
																	arg1: &Node{},
																},
															},
														},
														arg2: &Elem{
															DATA: []interface{}{"pp"},
															arg1: &Collect{
																ARG: collect__attributes__cat,
																arg1: &Collect{
																	ARG: collect__parent__type__node,
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
							},
						},
					}) { // mwu as head, but not complex P
						phrase = node.parent.parent
					} else {
						phrase = node.parent
					}
				}
			}
		} else {
			if Test(q /* $node[@rel="obj1" and ../@cat="pp"] */, &XPath{
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
										DATA: []interface{}{"obj1"},
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
						},
					},
				},
			}) {
				phrase = node.parent
			} else {
				phrase = node
				// do not apply to prepositions and auxiliaries, ever. Too strict?
			}
		}
	}

	if !Test(q /* $phrase[@rel=("obj1","su","mod","pc","det") and @index] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					VAR: phrase,
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
								DATA: []interface{}{"obj1", "su", "mod", "pc", "det"},
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
				},
			},
		},
	}) {
		return []DepType{}
	}

	// TODO: dit xpath kan efficiënter?
	conj_heads := Find(q, /* $node[not(@ud:pos=("ADP","AUX"))]/ancestor::node//node[@rel="cnj"
			     and node[
			    (: @rel=$phrase/@rel
				and -- this constraint is too strict for coord of passives:)
				not(@pt or @cat)]/@index = $phrase/@index
			     and node[@rel=("hd","predc") and not(@ud:pos="AUX") and (@pt or @cat) and	 (: bekende cafes zijn A en B :)
				(: not(@ud:pos=("ADP","AUX")) and not(@cat="mwu") :)
				not(local:internal_head_position(..) = @end and (@ud:pos=("ADP","AUX") or @cat="mwu") )
				]
		      ]
			      (: not coordination of AUX or (complex) Ps :) */&XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__descendant__or__self__type__node,
						arg1: &Collect{
							ARG: collect__ancestors__node,
							arg1: &Filter{
								arg1: &Variable{
									VAR: node,
								},
								arg2: &Sort{
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
														DATA: []interface{}{"ADP", "AUX"},
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
										ARG: collect__attributes__index,
										arg1: &Collect{
											ARG:  collect__child__node,
											arg1: &Node{},
											arg2: &Predicate{
												arg1: &Function{
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
										ARG: collect__attributes__index,
										arg1: &Variable{
											VAR: phrase,
										},
									},
								},
							},
							arg2: &Collect{
								ARG:  collect__child__node,
								arg1: &Node{},
								arg2: &Predicate{
									arg1: &And{
										arg1: &And{
											arg1: &And{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"hd", "predc"},
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
													arg1: &And{
														arg1: &Equal{
															ARG: equal__is,
															arg1: &Function{
																ARG: function__local__internal__head__position__1__args,
																arg1: &Arg{
																	arg1: &Sort{
																		arg1: &Collect{
																			ARG:  collect__parent__type__node,
																			arg1: &Node{},
																		},
																	},
																},
															},
															arg2: &Collect{
																ARG:  collect__attributes__end,
																arg1: &Node{},
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
																		DATA: []interface{}{"ADP", "AUX"},
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
		})
	if len(conj_heads) == 0 {
		return []DepType{}
	}

	udRelation := nonLocalDependencyLabel(phrase, n1(Find(q, /* ($q.varallnodes[@rel="cnj"]/
		   			    node[
		   			    (: @rel=$phrase/@rel and :)
						not(@pt or @cat) and @index=$phrase/@index])[1] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Sort{
						arg1: &Collect{
							ARG: collect__child__node,
							arg1: &Filter{
								arg1: &Variable{
									VAR: q.varallnodes,
								},
								arg2: &Sort{
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
							arg2: &Predicate{
								arg1: &And{
									arg1: &Function{
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
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__index,
											arg1: &Node{},
										},
										arg2: &Collect{
											ARG: collect__attributes__index,
											arg1: &Variable{
												VAR: phrase,
											},
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
		})), q)

	EudRelation := udRelation
	if Test(q /* $udRelation = ("nmod","obl") and $phrase[@cat="pp"]//node[@ud:Relation="case" and @ud:HeadPosition=$node/@end] */, &XPath{
		arg1: &Sort{
			arg1: &And{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Variable{
						VAR: udRelation,
					},
					arg2: &Elem{
						DATA: []interface{}{"nmod", "obl"},
						arg1: &Variable{
							VAR: udRelation,
						},
					},
				},
				arg2: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__descendant__or__self__type__node,
						arg1: &Filter{
							arg1: &Variable{
								VAR: phrase,
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
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__ud_3aRelation,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"case"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3aRelation,
										arg1: &Node{},
									},
								},
							},
							arg2: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__ud_3aHeadPosition,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG: collect__attributes__end,
									arg1: &Variable{
										VAR: node,
									},
								},
							},
						},
					},
				},
			},
		},
	}) {
		EudRelation = udRelation + ":" + enhancedLemmaString(Find(q /* $phrase//node[@ud:Relation="case" and @ud:HeadPosition=$node/@end] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__descendant__or__self__type__node,
						arg1: &Variable{
							VAR: phrase,
						},
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__ud_3aRelation,
									arg1: &Node{},
								},
								arg2: &Elem{
									DATA: []interface{}{"case"},
									arg1: &Collect{
										ARG:  collect__attributes__ud_3aRelation,
										arg1: &Node{},
									},
								},
							},
							arg2: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__ud_3aHeadPosition,
									arg1: &Node{},
								},
								arg2: &Collect{
									ARG: collect__attributes__end,
									arg1: &Variable{
										VAR: node,
									},
								},
							},
						},
					},
				},
			},
		}), q)
	}

	result := []DepType{}
	for _, conj_head := range conj_heads {
		result = append(result, DepType{head: internalHeadPosition([]interface{}{conj_head.(*NodeType)}, q), dep: EudRelation})

	}
	return result
}

// should work in coordinations like te laten reizen en te laten beleven,
// and recursive cases: Andras blijft ontkennen sexuele relaties met Timea te hebben gehad ,
//    .. of hij ook voor hen wilde komen tekenen :)
func xcompControl(node *NodeType, q *Context, so_index int) []DepType {

	result := []DepType{}
	for _, xcomp := range Find(q, /* $node[not(@ud:PronType="Rel")]/ancestor::node//node[(@rel="vc" or (@cat="inf" and @rel="body")) (: covers inf ti oti :)
		   and node[@rel=("hd","predc") and @ud:Relation="xcomp"]  (: vrouwen moeten vertegenwoordigd zijn :)
		   and node[@rel="su" and @index]/@index = $so_index
		  ] */&XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__descendant__or__self__type__node,
						arg1: &Collect{
							ARG: collect__ancestors__node,
							arg1: &Filter{
								arg1: &Variable{
									VAR: node,
								},
								arg2: &Sort{
									arg1: &Function{
										ARG: function__not__1__args,
										arg1: &Arg{
											arg1: &Sort{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__ud_3aPronType,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"Rel"},
														arg1: &Collect{
															ARG:  collect__attributes__ud_3aPronType,
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
					arg2: &Predicate{
						arg1: &And{
							arg1: &And{
								arg1: &Sort{
									arg1: &Or{
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
											arg1: &And{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__cat,
														arg1: &Node{},
													},
													arg2: &Elem{
														DATA: []interface{}{"inf"},
														arg1: &Collect{
															ARG:  collect__attributes__cat,
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
													DATA: []interface{}{"hd", "predc"},
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
												},
											},
											arg2: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__ud_3aRelation,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"xcomp"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3aRelation,
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
										},
									},
								},
								arg2: &Variable{
									VAR: so_index,
								},
							},
						},
					},
				},
			},
		}) {
		result = append(result, DepType{head: internalHeadPosition([]interface{}{xcomp.(*NodeType)}, q), dep: "nsubj:xsubj"})
	}
	return result
}

// alpino NF specific case, controllers with extraposed content are realized downstairs
func upstairsControl(node *NodeType, q *Context, so_index int) []DepType {

	result := []DepType{}
	for _, upstairs := range Find(q, /* $node/ancestor::node[node[@rel="hd" and @ud:pos="VERB"]
		 and node[@rel="vc"]
		 and node[@rel=("su","obj1","obj2") and not(@pt or @cat)]/@index = $so_index
		] */&XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__ancestors__node,
					arg1: &Variable{
						VAR: node,
					},
					arg2: &Predicate{
						arg1: &And{
							arg1: &And{
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
								arg2: &Collect{
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
							},
							arg2: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG: collect__attributes__index,
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
														DATA: []interface{}{"su", "obj1", "obj2"},
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
								arg2: &Variable{
									VAR: so_index,
								},
							},
						},
					},
				},
			},
		}) {
		result = append(result, DepType{head: internalHeadPosition([]interface{}{upstairs.(*NodeType)}, q), dep: "nsubj:xsubj"})
	}
	return result

}

// een koers waarin de Alsemberg moet worden beklommen
func passiveVpControl(node *NodeType, q *Context, so_index int) []DepType {

	result := []DepType{}
	for _, passive_vp := range Find(q, /* $q.varallnodes[@rel="vc" and @cat="ppart"
		   and node[@rel="hd" and @ud:Relation="xcomp"]
		   and node[@rel="obj1" and @index]/@index = $so_index ] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						VAR: q.varallnodes,
					},
					arg2: &Sort{
						arg1: &And{
							arg1: &And{
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
									arg2: &Equal{
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
													ARG:  collect__attributes__ud_3aRelation,
													arg1: &Node{},
												},
												arg2: &Elem{
													DATA: []interface{}{"xcomp"},
													arg1: &Collect{
														ARG:  collect__attributes__ud_3aRelation,
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
										},
									},
								},
								arg2: &Variable{
									VAR: so_index,
								},
							},
						},
					},
				},
			},
		}) {
		result = append(result, DepType{head: internalHeadPosition([]interface{}{passive_vp.(*NodeType)}, q), dep: "nsubj:pass:xsubj"})
	}

	return result
}

func enhancedLemmaString(nodes []interface{}, q *Context) string {
	sort.Slice(nodes, func(i, j int) bool {
		// TODO: NodeType heeft geen head
		return nodes[i].(*NodeType).udEHeadPosition < nodes[j].(*NodeType).udEHeadPosition
	})
	lemmas := make([]string, len(nodes))
	for i, node := range nodes {
		lemmas[i] = enhancedLemmaString1(node.(*NodeType), q)
	}
	return strings.Join(lemmas, "_")
}

func enhancedLemmaString1(node *NodeType, q *Context) string {
	var lemma string
	switch node.Lemma {
	case "a.k.a":
		lemma = "also_known_as"
	case "c.q.":
		lemma = "casu_quo"
	case "dwz.", "d.w.z.":
		lemma = "dat_wil_zeggen"
	case "e.d.":
		lemma = "en_dergelijke"
	case "en/of":
		lemma = "en_of"
	case "enz.":
		lemma = "enzovoort"
	case "etc.":
		lemma = "etcetera"
	case "m.a.w.":
		lemma = "met_andere_woorden"
	case "nl.":
		lemma = "namelijk"
	case "resp.":
		lemma = "respectievelijk"
	case "t/m":
		lemma = "tot_en_met"
	case "t.a.v.":
		lemma = "ten_aanzien_van"
	case "t.g.v.":
		lemma = "ten_gunste_van"
	case "t.n.v.":
		lemma = "ten_name_van"
	case "t.o.v.":
		lemma = "ten_opzichte_van"
	default:
		lemma = node.Lemma
	}
	fixed := Find(q /* $node/../node[@ud:ERelation="fixed"] */, &XPath{
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
							ARG:  collect__attributes__ud_3aERelation,
							arg1: &Node{},
						},
						arg2: &Elem{
							DATA: []interface{}{"fixed"},
							arg1: &Collect{
								ARG:  collect__attributes__ud_3aERelation,
								arg1: &Node{},
							},
						},
					},
				},
			},
		},
	})
	if len(fixed) > 0 {
		sort.Slice(fixed, func(i, j int) bool {
			fi := fixed[i].(*NodeType)
			fj := fixed[j].(*NodeType)
			ei := fi.End
			ej := fj.End
			if fi.udCopiedFrom > 0 {
				ei = fi.udCopiedFrom
			}
			if fj.udCopiedFrom > 0 {
				ej = fj.udCopiedFrom
			}
			return ei < ej
		})
		for _, f := range fixed {
			lemma += "_" + f.(*NodeType).Lemma
		}
	}
	lemma = strings.Replace(lemma, "/", "schuine_streep", -1)
	lemma = strings.Replace(lemma, "-", "_", -1)
	return strings.ToLower(lemma)
}
