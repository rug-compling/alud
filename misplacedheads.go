//
// GENERATED FILE -- DO NOT EDIT
//

package main

import (
	"fmt"
	"os"
)

// voorkwam dat LPF opnieuw of SGP voor het eerst in de regering zou komen  -- gapped LD
func fixMisplacedHeadsInCoordination(q *Context) {

	if len(q.varindexnodes) == 0 {
		return
	}

	seen := make(map[[2]int]int)
	counter := 0

START:
	for true {
		for _, n1 := range q.varallnodes {
			// FIND op varallnodes niet mogelijk omdat twee keer naar $node wordt verwezen, en dat moet dezelfde node zijn
			for _, n2 := range Find(q, /*
				$n1[@rel=("hd","ld") and
				      @index and
				      (@pt or @cat) and
				      ancestor::node[@rel="cnj"] and
				      ancestor::node[@cat="conj"]/node[@rel="cnj" and
				                                       descendant-or-self::node[@rel=("hd","ld") and
				                                                                @index=$n1/@index and
				                                                                not(@cat or @pt) and
				                                                                ( @begin        = ..//node[@cat or @pt]/@end or
				                                                                  @begin - 1000 = ..//node[@cat or @pt]/@end
				                                                                )
				                                                               ]
				                                       ]] */&XPath{
					arg1: &Sort{
						arg1: &Filter{
							arg1: &Variable{
								VAR: n1,
							},
							arg2: &Sort{
								arg1: &And{
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
														DATA: []interface{}{"hd", "ld"},
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
										arg2: &Collect{
											ARG:  collect__ancestors__node,
											arg1: &Node{},
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
									arg2: &Collect{
										ARG: collect__child__node,
										arg1: &Collect{
											ARG:  collect__ancestors__node,
											arg1: &Node{},
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
													ARG:  collect__descendant__or__self__node,
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
																			DATA: []interface{}{"hd", "ld"},
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
																				VAR: n1,
																			},
																		},
																	},
																},
																arg2: &Function{
																	ARG: function__not__1__args,
																	arg1: &Arg{
																		arg1: &Sort{
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
															arg2: &Sort{
																arg1: &Or{
																	arg1: &Equal{
																		ARG: equal__is,
																		arg1: &Collect{
																			ARG:  collect__attributes__begin,
																			arg1: &Node{},
																		},
																		arg2: &Collect{
																			ARG: collect__attributes__end,
																			arg1: &Collect{
																				ARG: collect__child__node,
																				arg1: &Collect{
																					ARG: collect__descendant__or__self__type__node,
																					arg1: &Collect{
																						ARG:  collect__parent__type__node,
																						arg1: &Node{},
																					},
																				},
																				arg2: &Predicate{
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
																	arg2: &Equal{
																		ARG: equal__is,
																		arg1: &Plus{
																			ARG: plus__minus,
																			arg1: &Collect{
																				ARG:  collect__attributes__begin,
																				arg1: &Node{},
																			},
																			arg2: &Elem{
																				DATA: []interface{}{1000},
																				arg1: &Collect{
																					ARG:  collect__attributes__begin,
																					arg1: &Node{},
																				},
																			},
																		},
																		arg2: &Collect{
																			ARG: collect__attributes__end,
																			arg1: &Collect{
																				ARG: collect__child__node,
																				arg1: &Collect{
																					ARG: collect__descendant__or__self__type__node,
																					arg1: &Collect{
																						ARG:  collect__parent__type__node,
																						arg1: &Node{},
																					},
																				},
																				arg2: &Predicate{
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
															},
														},
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
				node2 := n2.(*NodeType)
				for _, n3 := range Find(q, /*
					$q.varallnodes[@rel=("hd","ld","vc") and @index and not(@pt or @cat) and
					                 ancestor::node[@rel="cnj"]  and
					                                    ( @begin        = ..//node[@cat or @pt]/@end or
					                                      @begin - 1000 = ..//node[@cat or @pt]/@end
					                                     )] */&XPath{
						arg1: &Sort{
							arg1: &Filter{
								arg1: &Variable{
									VAR: q.varallnodes,
								},
								arg2: &Sort{
									arg1: &And{
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
															DATA: []interface{}{"hd", "ld", "vc"},
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
											arg2: &Collect{
												ARG:  collect__ancestors__node,
												arg1: &Node{},
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
										arg2: &Sort{
											arg1: &Or{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__begin,
														arg1: &Node{},
													},
													arg2: &Collect{
														ARG: collect__attributes__end,
														arg1: &Collect{
															ARG: collect__child__node,
															arg1: &Collect{
																ARG: collect__descendant__or__self__type__node,
																arg1: &Collect{
																	ARG:  collect__parent__type__node,
																	arg1: &Node{},
																},
															},
															arg2: &Predicate{
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
												arg2: &Equal{
													ARG: equal__is,
													arg1: &Plus{
														ARG: plus__minus,
														arg1: &Collect{
															ARG:  collect__attributes__begin,
															arg1: &Node{},
														},
														arg2: &Elem{
															DATA: []interface{}{1000},
															arg1: &Collect{
																ARG:  collect__attributes__begin,
																arg1: &Node{},
															},
														},
													},
													arg2: &Collect{
														ARG: collect__attributes__end,
														arg1: &Collect{
															ARG: collect__child__node,
															arg1: &Collect{
																ARG: collect__descendant__or__self__type__node,
																arg1: &Collect{
																	ARG:  collect__parent__type__node,
																	arg1: &Node{},
																},
															},
															arg2: &Predicate{
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
										},
									},
								},
							},
						},
					}) {
					node3 := n3.(*NodeType)
					if node2.Index == node3.Index {
						pair := [2]int{node2.Id, node3.Id}
						if i, ok := seen[pair]; ok {
							if i == 1 {
								warning := fmt.Sprintf("Loop detected in fixMisplacedHeadsInCoordination: %d %d", node2.Id, node3.Id)
								q.warnings = append(q.warnings, warning)
								fmt.Fprintln(os.Stderr, warning, "in", q.filename)
							}
							seen[pair]++
							continue
						}
						seen[pair] = 1
						counter++
						// kopieer inhoud van node2 (niet leeg) naar node3 (leeg)
						id, rel := node3.Id, node3.Rel
						*node3 = *node2
						node3.Id, node3.Rel = id, rel
						// maak node2 leeg
						*node2 = NodeType{
							Begin: node2.Begin,
							End:   node2.End,
							Id:    node2.Id,
							Index: node2.Index,
							Rel:   node2.Rel,
							Node:  []*NodeType{},
						}
						// opnieuw beginnen
						inspect(q)
						continue START
					}
				}
			}
		}
		break
	}
	if counter > 0 {
		q.debugs = append(q.debugs, fmt.Sprintf("fixMisplacedHeadsInCoordination: %d swaps", counter))
	}
}
