//
// GENERATED FILE -- DO NOT EDIT
//

package main

/*
func auxiliary(nodes []*NodeType, q *Context) string {
	if len(nodes) != 1 {
		return "ERROR_AUXILIARY_FUNCTION_TAKES_EXACTLY_ONE_ARG"
	}
	return auxiliary1(nodes[0], q)
}
*/

func auxiliary1(node *NodeType, q *Context) string {

	if node.Pt != "ww" {
		return "ERROR_NO_VERB"
	}
	if node.Rel != "hd" {
		return "verb"
	}

	if Test(node, q, /* $node[not(../node[@rel=("obj1","se","vc")]) and
			        (: ud documentation suggests 1 cop per lg, van Eynde suggests much more, compromise: the traditional ones :)
			        @lemma=("zijn","lijken","blijken","blijven","schijnen","heten","voorkomen","worden","dunken") and
		                 ( contains(@sc,'copula') or
		                   contains(@sc,'pred')   or
		                   contains(@sc,'cleft')  or
		                   ../node[@rel="predc"]
		                 ) ] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Sort{
						arg1: &And{
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
													arg1: &Equal{
														ARG: equal__is,
														arg1: &Collect{
															ARG:  collect__attributes__rel,
															arg1: &Node{},
														},
														arg2: &Elem{
															ARG: elem__string__obj1__se__vc,
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
								arg2: &Equal{
									ARG: equal__is,
									arg1: &Collect{
										ARG:  collect__attributes__lemma,
										arg1: &Node{},
									},
									arg2: &Elem{
										ARG: elem__string__zijn__lijken__blijken__blijven__schijnen__het_2e_2e_2e,
										arg1: &Collect{
											ARG:  collect__attributes__lemma,
											arg1: &Node{},
										},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Or{
										arg1: &Or{
											arg1: &Function{
												ARG: function__contains__2__args,
												arg1: &Arg{
													arg1: &Arg{
														arg1: &Sort{
															arg1: &Collect{
																ARG:  collect__attributes__sc,
																arg1: &Node{},
															},
														},
													},
													arg2: &Elem{
														ARG: elem__string__copula,
													},
												},
											},
											arg2: &Function{
												ARG: function__contains__2__args,
												arg1: &Arg{
													arg1: &Arg{
														arg1: &Sort{
															arg1: &Collect{
																ARG:  collect__attributes__sc,
																arg1: &Node{},
															},
														},
													},
													arg2: &Elem{
														ARG: elem__string__pred,
													},
												},
											},
										},
										arg2: &Function{
											ARG: function__contains__2__args,
											arg1: &Arg{
												arg1: &Arg{
													arg1: &Sort{
														arg1: &Collect{
															ARG:  collect__attributes__sc,
															arg1: &Node{},
														},
													},
												},
												arg2: &Elem{
													ARG: elem__string__cleft,
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
													ARG: elem__string__predc,
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
		return "cop"
	}

	if Test(node, q, /* $node[@lemma=("zijn","worden") and
		   ( @sc="passive"  or
		   	 ( ../node[@rel="vc"] and
		         ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj1"]/@index or
		           ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj1"]/@index or
		           ../node[@rel="vc" and not(@pt or @cat)]/@index =
		               $indexnodes[@rel="vc" and node[@rel="obj1"]/@index = $node/../node[@rel="su"]/@index]/@index
		          or not(../node[@rel="su"])
		          )
		      )
		   ) ] */&XPath{
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
									ARG:  collect__attributes__lemma,
									arg1: &Node{},
								},
								arg2: &Elem{
									ARG: elem__string__zijn__worden,
									arg1: &Collect{
										ARG:  collect__attributes__lemma,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
											ARG:  collect__attributes__sc,
											arg1: &Node{},
										},
										arg2: &Elem{
											ARG: elem__string__passive,
											arg1: &Collect{
												ARG:  collect__attributes__sc,
												arg1: &Node{},
											},
										},
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
															ARG: elem__string__vc,
															arg1: &Collect{
																ARG:  collect__attributes__rel,
																arg1: &Node{},
															},
														},
													},
												},
											},
											arg2: &Sort{
												arg1: &Or{
													arg1: &Or{
														arg1: &Or{
															arg1: &Equal{
																ARG: equal__is,
																arg1: &Collect{
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
																					ARG: elem__string__su,
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
																						ARG: elem__string__vc,
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
																					ARG: elem__string__obj1,
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
															arg2: &Equal{
																ARG: equal__is,
																arg1: &Collect{
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
																					ARG: elem__string__su,
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
																			ARG: collect__child__node,
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
																							ARG: elem__string__vc,
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
																						ARG: elem__string__cnj,
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
																					ARG: elem__string__obj1,
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
														arg2: &Equal{
															ARG: equal__is,
															arg1: &Collect{
																ARG: collect__attributes__index,
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
																					ARG: elem__string__vc,
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
																ARG: collect__attributes__index,
																arg1: &Filter{
																	arg1: &Variable{
																		ARG: variable__indexnodes,
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
																					ARG: elem__string__vc,
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
																							arg1: &Equal{
																								ARG: equal__is,
																								arg1: &Collect{
																									ARG:  collect__attributes__rel,
																									arg1: &Node{},
																								},
																								arg2: &Elem{
																									ARG: elem__string__obj1,
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
																									ARG: elem__string__su,
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
																				ARG: elem__string__su,
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
							},
						},
					},
				},
			},
		}) {
		return "aux:pass"
	}

	if Test(node, q, /*
		  (: krijgen passive with iobj control :)
		            $node[@lemma="krijgen" and
		  	              ( ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="obj2"]/@index or
		                    ../node[@rel="su"]/@index = ../node[@rel="vc"]/node[@rel="cnj"]/node[@rel="obj2"]/@index
		                  )] */&XPath{
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
									ARG:  collect__attributes__lemma,
									arg1: &Node{},
								},
								arg2: &Elem{
									ARG: elem__string__krijgen,
									arg1: &Collect{
										ARG:  collect__attributes__lemma,
										arg1: &Node{},
									},
								},
							},
							arg2: &Sort{
								arg1: &Or{
									arg1: &Equal{
										ARG: equal__is,
										arg1: &Collect{
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
															ARG: elem__string__su,
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
																ARG: elem__string__vc,
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
															ARG: elem__string__obj2,
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
									arg2: &Equal{
										ARG: equal__is,
										arg1: &Collect{
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
															ARG: elem__string__su,
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
													ARG: collect__child__node,
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
																	ARG: elem__string__vc,
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
																ARG: elem__string__cnj,
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
															ARG: elem__string__obj2,
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
		return "aux:pass"
	}

	if Test(node, q, /*
		  (: alpino has no principled distinction between AUX and VERB, should be TAME verbs semantically, we follow ENGLISH :)
		          $node[not(../node[@rel="predc"]) and  (: hij heeft als opdracht stammen uit elkaar te houden  :)
		                 ( starts-with(@sc,'aux') or
		                   ( ../node[@rel="vc"  and
		                              ( @cat=("ppart","inf","ti") or
		                                ( @cat="conj" and node[@rel="cnj" and @cat=("ppart","inf","ti")] ) or
		                                ( @index and not(@pt or @cat))  (: dangling aux in gapped coordination :)
		                              )
		                            ]   and
		                     @lemma=("blijken","hebben","hoeven","kunnen","moeten","mogen","zijn","zullen")
		                   )
		                 )
		               ] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
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
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__rel,
														arg1: &Node{},
													},
													arg2: &Elem{
														ARG: elem__string__predc,
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
							arg2: &Sort{
								arg1: &Or{
									arg1: &Function{
										ARG: function__starts__with__2__args,
										arg1: &Arg{
											arg1: &Arg{
												arg1: &Sort{
													arg1: &Collect{
														ARG:  collect__attributes__sc,
														arg1: &Node{},
													},
												},
											},
											arg2: &Elem{
												ARG: elem__string__aux,
											},
										},
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
													arg1: &And{
														arg1: &Equal{
															ARG: equal__is,
															arg1: &Collect{
																ARG:  collect__attributes__rel,
																arg1: &Node{},
															},
															arg2: &Elem{
																ARG: elem__string__vc,
																arg1: &Collect{
																	ARG:  collect__attributes__rel,
																	arg1: &Node{},
																},
															},
														},
														arg2: &Sort{
															arg1: &Or{
																arg1: &Or{
																	arg1: &Equal{
																		ARG: equal__is,
																		arg1: &Collect{
																			ARG:  collect__attributes__cat,
																			arg1: &Node{},
																		},
																		arg2: &Elem{
																			ARG: elem__string__ppart__inf__ti,
																			arg1: &Collect{
																				ARG:  collect__attributes__cat,
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
																					ARG: elem__string__conj,
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
																						arg1: &Equal{
																							ARG: equal__is,
																							arg1: &Collect{
																								ARG:  collect__attributes__rel,
																								arg1: &Node{},
																							},
																							arg2: &Elem{
																								ARG: elem__string__cnj,
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
																								ARG: elem__string__ppart__inf__ti,
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
											arg2: &Equal{
												ARG: equal__is,
												arg1: &Collect{
													ARG:  collect__attributes__lemma,
													arg1: &Node{},
												},
												arg2: &Elem{
													ARG: elem__string__blijken__hebben__hoeven__kunnen__moeten__moge_2e_2e_2e,
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
		}) {
		return "aux"
	}

	/*
	  else if ($node[@pt="ww"] )
	  then "verb"
	*/
	return "verb"

	/*
	  else "ERROR_NO_VERB"
	*/

}
