//
// GENERATED FILE -- DO NOT EDIT
//

package main

import (
	//	"github.com/kr/pretty"

	"fmt"
	"os"
	"sort"
)

// voorkwam dat LPF opnieuw of SGP voor het eerst in de regering zou komen  -- gapped LD
func fixMisplacedHeadsInCoordination(q *Context) {

	if len(q.varindexnodes) == 0 {
		return
	}

	counter := 0

START:
	for true {
		for _, n1 := range q.varallnodes {
			// FIND op varallnodes niet mogelijk omdat twee keer naar $node wordt verwezen, en dat moet dezelfde node zijn
			for _, n2 := range Find(n1, q, /*
				$node[@rel=("hd","ld") and
				      @index and
				      (@pt or @cat) and
				      ancestor::node[@rel="cnj"] and
				      ancestor::node[@cat="conj"]/node[@rel="cnj" and
				                                       descendant-or-self::node[@rel=("hd","ld") and
				                                                                @index=$node/@index and
				                                                                not(@cat or @pt) and
				                                                                ( @begin        = ..//node[@cat or @pt]/@end or
				                                                                  @begin - 1000 = ..//node[@cat or @pt]/@end
				                                                                )
				                                                               ]
				                                       ]] */&XPath{
					arg1: &Sort{
						arg1: &Filter{
							arg1: &Variable{
								ARG: variable__node,
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
														ARG: elem__string__hd__ld,
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
														ARG: elem__string__cnj,
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
														ARG: elem__string__conj,
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
														ARG: elem__string__cnj,
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
																			ARG: elem__string__hd__ld,
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
																				ARG: elem__number__1000,
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
				for _, n3 := range Find(q.varallnodes, q, /*
					$node[@rel=("hd","ld","vc") and @index and not(@pt or @cat) and
					                 ancestor::node[@rel="cnj"]  and
					                                    ( @begin        = ..//node[@cat or @pt]/@end or
					                                      @begin - 1000 = ..//node[@cat or @pt]/@end
					                                     )] */&XPath{
						arg1: &Sort{
							arg1: &Filter{
								arg1: &Variable{
									ARG: variable__node,
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
															ARG: elem__string__hd__ld__vc,
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
															ARG: elem__string__cnj,
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
															ARG: elem__number__1000,
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
						counter++
						if counter == 100 {
							q.warnings = append(q.warnings, "Swap limit for fixMisplacedHeadsInCoordination")
							fmt.Fprintln(os.Stderr, "WARNING: Swap limit for fixMisplacedHeadsInCoordination in "+q.filename)
							break START
						}
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

func addPosTags(q *Context) {
	for _, node := range q.ptnodes {
		node.udPos = universalPosTags(node, q)
	}
}

func universalPosTags(node *NodeType, q *Context) string {
	pt := node.Pt
	rel := node.Rel

	if pt == "let" {
		if rel == "--" {
			for _, n := range node.parent.Node {
				if n.Pt != "let" || n.Begin < node.Begin {
					return "PUNCT"
				}
			}
		}
		return "SYM"
	}
	if pt == "adj" {
		if rel == "det" {
			return "DET"
		}
		if rel == "hd" && node.parent.Cat == "pp" {
			// vol vertrouwen
			return "ADP"
		}
		if rel == "crd" {
			// respectievelijk
			return "CCONJ"
		}
		return "ADJ" // exceptions forced by 2.4 validation
	}
	if pt == "bw" {
		return "ADV"
	}
	if pt == "lid" {
		return "DET"
	}
	if pt == "n" {
		if node.Ntype == "eigen" {
			return "PROPN"
		}
		return "NOUN"
	}
	if pt == "spec" {
		if node.Spectype == "deeleigen" {
			return "PROPN"
		}
		if node.Spectype == "symb" {
			return "SYM"
		}
		return "X" // afk vreemd afgebr enof meta
	}
	if pt == "tsw" {
		return "INTJ"
	}
	if pt == "tw" {
		if node.Numtype == "rang" {
			return "ADJ"
		}
		return "NUM"
	}
	if pt == "vz" {
		return "ADP" // v2: do not use PART for SVPs and complementizers
	}
	if pt == "vnw" {
		if rel == "det" && node.Vwtype != "bez" {
			return "DET"
		}
		if node.Pdtype == "adv-pron" {
			if rel == "pobj1" {
				return "PRON"
			}
			return "ADV"
		}
		if (rel == "mod" || (rel == "hd" && node.parent.Rel == "mod")) && node.Pdtype == "grad" {
			// veel minder
			return "ADV"
		}
		return "PRON"
	}
	if pt == "vg" {
		if node.Conjtype == "neven" {
			return "CCONJ" // V2: CONJ ==> CCONJ
		}
		return "SCONJ"
	}
	if pt == "ww" {
		if auxiliary1(node, q) == "verb" {
			return "VERB"
		}
		return "AUX" // v2: cop and aux:pass --> AUX  (already in place in v1?)
	}
	return "ERROR_NO_POS_FOUND"
}

func addFeatures(q *Context) {
	for _, node := range q.ptnodes {
		switch node.udPos {
		case "NOUN", "PROPN":
			nominalFeatues(node, q)
		case "ADJ":
			adjectiveFeatures(node, q)
		case "PRON":
			pronounFeatures(node, q)
		case "VERB", "AUX":
			verbalFeatures(node, q)
		case "DET":
			determinerFeatures(node, q)
		case "X":
			specialFeatures(node, q)
		}
	}
}

func nominalFeatues(node *NodeType, q *Context) {
	switch node.Genus {
	case "zijd":
		node.udGender = "Com"
	case "onz":
		node.udGender = "Neut"
	case "genus":
		node.udGender = "Com,Neut"
	case "":
		node.udGender = ""
	default:
		node.udGender = "ERROR_IRREGULAR_GENDER"

	}
	switch node.Getal {
	case "ev":
		node.udNumber = "Sing"
	case "mv":
		node.udNumber = "Plur"
	case "":
		node.udNumber = ""
	default:
		node.udNumber = "ERROR_IRREGULAR_NUMBER"
	}
	switch node.Graad {
	case "dim":
		node.udDegree = ""
	case "basis":
		node.udDegree = ""
	case "":
		node.udDegree = ""
	default:
		node.udDegree = "ERROR_IRREGULAR_DEGREE"
	}
}

func adjectiveFeatures(node *NodeType, q *Context) {
	switch node.Graad {
	case "basis":
		node.udDegree = "Pos"
	case "comp":
		node.udDegree = "Cmp"
	case "sup":
		node.udDegree = "Sup"
	case "dim": // netjes
		node.udDegree = "Pos"
	case "":
		node.udDegree = ""
	default:
		node.udDegree = "ERROR_IRREGULAR_DEGREE"
	}
}

func pronounFeatures(node *NodeType, q *Context) {
	switch node.Vwtype {
	case "refl":
		node.udPronType = "Prs"
		node.udReflex = "Yes"
	case "bez":
		node.udPronType = "Prs"
		node.udPoss = "Yes"
	case "pers", "pr":
		node.udPronType = "Prs"
	case "recip":
		node.udPronType = "Rcp"
	case "vb":
		node.udPronType = "Int"
	case "aanw":
		node.udPronType = "Dem"
	case "onbep":
		node.udPronType = "Ind"
	case "betr":
		node.udPronType = "Rel"
	case "excl": // occurs only once
		node.udPronType = ""
	case "":
		node.udPronType = ""
	default:
		node.udPronType = "ERROR_IRREGULAR_PRONTYPE"
	}

	switch node.Persoon {
	case "1":
		node.udPerson = "1"
	case "2", "2b", "2v":
		node.udPerson = "2"
	case "3", "3o", "3v", "3p", "3m":
		node.udPerson = "3"
	case "persoon":
		node.udPerson = ""
	case "":
		node.udPerson = ""
	default:
		node.udPerson = "ERROR_IRREGULAR_PERSON"
	}

	switch node.Naamval {
	case "nomin":
		node.udCase = "Nom"
	case "obl":
		node.udCase = "Acc"
	case "gen":
		node.udCase = "Gen"
	case "dat": // van dien aard
		node.udCase = "Dat"
	case "stan":
		node.udCase = ""
	case "":
		node.udCase = ""
	default:
		node.udCase = "ERROR_IRREGULAR_CASE"

	}
}

func verbalFeatures(node *NodeType, q *Context) {

	switch node.Wvorm {
	case "pv":
		node.udVerbForm = "Fin"
	case "inf":
		node.udVerbForm = "Inf"
	case "vd", "od":
		node.udVerbForm = "Part"
	case "":
		node.udVerbForm = ""
	default:
		node.udVerbForm = "ERROR_IRREGULAR_VERBFORM"
	}

	switch node.Pvtijd {
	case "verl":
		node.udTense = "Past"
	case "tgw", "conj":
		node.udTense = "Pres"
	case "":
		node.udTense = ""
	default:
		node.udTense = "ERROR_IRREGULAR_TENSE"
	}

	switch node.Pvagr {
	case "ev", "met-t":
		node.udNumber = "Sing"
	case "mv":
		node.udNumber = "Plur"
	case "":
		node.udNumber = ""
	default:
		node.udNumber = "ERROR_IRREGULAR_NUMBER"
	}
}

func determinerFeatures(node *NodeType, q *Context) {
	switch node.Lwtype {
	case "bep":
		node.udDefinite = "Def"
	case "onbep":
		node.udDefinite = "Ind"
	case "":
		node.udDefinite = ""
	default:
		node.udDefinite = "ERROR_IRREGULAR_DEFINITE"
	}
}

func specialFeatures(node *NodeType, q *Context) {
	switch node.Spectype {
	case "vreemd":
		node.udForeign = "Yes"
	case "afk":
		node.udAbbr = "Yes"
	}
}

func addDependencyRelations(q *Context) {
	for _, node := range q.ptnodes {
		q.depth = 0
		node.udRelation = dependencyLabel(node, q)
		q.depth = 0
		node.udHeadPosition = externalHeadPosition(node, q)
	}
}

// recursive
func dependencyLabel(node *NodeType, q *Context) string {
	if depthCheck(q, "dependencyLabel") {
		return "ERROR_NO_LABEL"
	}

	if node.parent.Cat == "top" && node.parent.End == 1000 {
		return "root"
	}
	if node.Rel == "app" {
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
									ARG: elem__string__hd,
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
		if Test(node, q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &XPath{
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
									ARG: elem__string__mod,
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
	if node.Rel == "me" && Test(node, q /* $node[@rel="me" and not(../node[@ud:pos="ADP"]) ] */, &XPath{
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
								ARG: elem__string__me,
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
													ARG: elem__string__ADP,
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
		if Test(node, q /* $node/../node[@rel=("obj1","se") and (@pt or @cat)] or $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX")] */, &XPath{
			arg1: &Sort{
				arg1: &Or{
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
										ARG: elem__string__obj1__se,
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
											ARG: elem__string__hd,
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
													ARG: elem__string__AUX,
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
							ARG: elem__string__pp,
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
			}) { // check for absolutive (met) constructions, https://github.com/UniversalDependencies/docs/issues/408
				if Test(node, q /* $node/../../@cat="np" */, &XPath{
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
											ARG: variable__node,
										},
									},
								},
							},
							arg2: &Elem{
								ARG: elem__string__np,
								arg1: &Collect{
									ARG: collect__attributes__cat,
									arg1: &Collect{
										ARG: collect__parent__type__node,
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
		if Test(node, q /* $node[../@rel="cnj" and ../node[@rel="hd" and not(@pt or @cat)]] */, &XPath{
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
									ARG: collect__attributes__rel,
									arg1: &Collect{
										ARG:  collect__parent__type__node,
										arg1: &Node{},
									},
								},
								arg2: &Elem{
									ARG: elem__string__cnj,
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
												ARG: elem__string__hd,
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
		if Test(node, q, /* $node[../@rel="vc" and ../node[@rel="hd" and not(@pt or @cat)]
			   and ../../self::node[@rel="cnj" and node[@rel="hd" and not(@pt or @cat)]]] */&XPath{
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
											ARG: collect__attributes__rel,
											arg1: &Collect{
												ARG:  collect__parent__type__node,
												arg1: &Node{},
											},
										},
										arg2: &Elem{
											ARG: elem__string__vc,
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
														ARG: elem__string__hd,
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
									ARG: collect__self__node,
									arg1: &Collect{
										ARG: collect__parent__type__node,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
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
													ARG: elem__string__cnj,
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
																ARG: elem__string__hd,
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
			// TODO: ../.. is veranderd in ../../self::node    is dat juist?
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
		if Test(node, q, /* $node[../node[@rel="su" and not(@pt or @cat)] and
			   ../node[@rel="vc" and not(@pt or @cat)] and
			   ../@rel="cnj"] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							ARG: variable__node,
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
														ARG: elem__string__su,
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
										ARG: elem__string__cnj,
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
									ARG: elem__string__hd,
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
		if Test(node, q /* $node/../node[@rel="mod" and (@pt or @cat)] */, &XPath{
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
									ARG: elem__string__mod,
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
		if Test(node, q /* $node/../@cat="pp" or $node/../node[@rel="hd" and @ud:pos="ADP"] */, &XPath{
			arg1: &Sort{
				arg1: &Or{
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
							ARG: elem__string__pp,
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
					arg2: &Collect{
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
										ARG: elem__string__hd,
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
										ARG: elem__string__ADP,
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
			}) { // absolutive met
				return "nsubj"
			}
			return dependencyLabel(node.parent, q)
		}
		if Test(node, q /* $node[@index = ../../node[@rel="su"]/@index ] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
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
		}) {
			return "nsubj:pass" // trees where su (with extraposed material) is spelled out at position of obj1
		}
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
									ARG: elem__string__hd,
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
		if Test(node, q /* $node/../node[@rel="su" and (@pt or @cat)] */, &XPath{
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
									ARG: elem__string__su,
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
		if Test(node, q /* $node/../node[@ud:pos="PROPN"] */, &XPath{
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
								ARG:  collect__attributes__ud_3apos,
								arg1: &Node{},
							},
							arg2: &Elem{
								ARG: elem__string__PROPN,
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
		// TODO als ik hier firstnode gebruik i.p.v. leftmost (wat zou moeten) gaat het vaak mis
		if node == leftmost(Find(node, q /* $node/../node[@rel="cnj"] */, &XPath{
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
								ARG: elem__string__cnj,
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
		return "conj"
	}
	if node.Rel == "dp" {
		if node == leftmost(Find(node, q /* $node/../node[@rel="dp"] */, &XPath{
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
								ARG: elem__string__dp,
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
		if Test(node, q /* $node/../node[@rel="hd" and @ud:pos=("AUX","ADP")] and not($node/../node[@rel="predc"]) */, &XPath{
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
										ARG: elem__string__hd,
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
										ARG: elem__string__AUX__ADP,
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
		}) {
			return dependencyLabel(node.parent, q)
		}
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
									ARG: elem__string__hd,
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
			if Test(node, q, /* $node/node[@rel="su" and @index and not(@word or @cat)]
				   (: or $node[@cat=("ti","oti")] :)
				   or $node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)]
				   or $node[@cat="oti"]/node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)] */&XPath{
					arg1: &Sort{
						arg1: &Or{
							arg1: &Or{
								arg1: &Collect{
									ARG: collect__child__node,
									arg1: &Variable{
										ARG: variable__node,
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
														ARG: elem__string__su,
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
														ARG: elem__string__ti,
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
													ARG: elem__string__body,
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
														ARG: elem__string__su,
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
														ARG: elem__string__oti,
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
													ARG: elem__string__ti,
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
												ARG: elem__string__body,
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
													ARG: elem__string__su,
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
			if Test(node, q /* $node/../@cat="np" */, &XPath{
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
							ARG: elem__string__np,
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
			}) {
				return "acl" // v2: clausal dependents of nouns always acl
			}
			return "ccomp"
		}
		if Test(node, q /* $node/../node[@rel=("su","obj1") and (@pt or @cat)] */, &XPath{
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
									ARG: elem__string__su__obj1,
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
									ARG: elem__string__hd,
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
									ARG: elem__string__mod,
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
	if Test(node, q /* $node[@rel=("mod","pc","ld") and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")] */, &XPath{
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
								ARG: elem__string__mod__pc__ld,
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
								ARG: elem__string__sv1__smain__ssub__inf__ppres__ppart__oti__ap__ad_2e_2e_2e,
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
		if Test(node, q /* $node/../node[@rel=("hd","body") and (@pt or @cat)] */, &XPath{
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
									ARG: elem__string__hd__body,
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
		if Test(node, q /* $node/../node[@rel=("su","obj1","predc","vc") and (@pt or @cat)] */, &XPath{
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
									ARG: elem__string__su__obj1__predc__vc,
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
		if Test(node, q /* $node[@rel="mod" and ../node[@rel=("pc","ld")]] */, &XPath{
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
									ARG: elem__string__mod,
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
											ARG: elem__string__pc__ld,
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
		if Test(node, q /* $node[@rel="mod" and ../node[@rel="mod"]/@begin < @begin] */, &XPath{
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
									ARG: elem__string__mod,
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
													ARG: elem__string__mod,
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
	if Test(node, q /* $node[@rel="mod" and ../@cat=("pp","detp","advp")] */, &XPath{
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
								ARG: elem__string__mod,
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
								ARG: elem__string__pp__detp__advp,
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
	if Test(node, q /* $node[@rel="mod" and ../@cat=("cp", "whrel", "whq", "whsub")] */, &XPath{
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
								ARG: elem__string__mod,
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
								ARG: elem__string__cp__whrel__whq__whsub,
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
		if Test(node, q /* $node/../node[@rel="body"]/node[@rel="hd" and @ud:pos="VERB"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
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
									ARG: elem__string__body,
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
									ARG: elem__string__hd,
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
									ARG: elem__string__VERB,
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
		if Test(node, q /* $node/../node[@rel="body"]//node/@index = $node/@index */, &XPath{
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
											ARG: elem__string__body,
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
							ARG: variable__node,
						},
					},
				},
			},
		}) {
			// index is een int groter dan 0
			return nonLocalDependencyLabel(
				node,
				firstnode(Find(node, q /* $node/../node[@rel="body"]//node[@index = $node/@index] */, &XPath{
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
												ARG: elem__string__body,
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
			if Test(node, q /* $node[not(../node[not(@ud:pos="PUNCT")]) and @begin = ../@begin] */, &XPath{
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
																		ARG: elem__string__PUNCT,
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
		if Test(node, q /* $node[@ud:pos="NUM" and ../node[@cat] ] */, &XPath{
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
									ARG: elem__string__NUM,
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
		if Test(node, q /* $node[@ud:pos="CCONJ" and ../node[@cat="smain" or @cat="conj"]] */, &XPath{
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
									ARG: elem__string__CCONJ,
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
												ARG: elem__string__smain,
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
												ARG: elem__string__conj,
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
		if Test(node, q /* $node[@ud:pos=("NOUN","PROPN","VERB") and ../node[@cat=("du","smain")]] */, &XPath{
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
									ARG: elem__string__NOUN__PROPN__VERB,
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
											ARG: elem__string__du__smain,
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
		if len(Find(node, q /* $node/../node[not(@ud:pos=("PUNCT","SYM","X"))] */, &XPath{
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
											ARG: elem__string__PUNCT__SYM__X,
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
		})) < 2 {
			return "root" // only one non-punct/sym/foreign element in the string
		}
		if node.Cat == "mwu" {
			if node.Begin == node.parent.Begin && node.End == node.parent.End {
				return "root"
			}
			if Test(node, q /* $node/node[@ud:pos=("PUNCT","SYM")] */, &XPath{
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
									ARG:  collect__attributes__ud_3apos,
									arg1: &Node{},
								},
								arg2: &Elem{
									ARG: elem__string__PUNCT__SYM,
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
		if Test(node, q /* $node[not(@ud:pos)]/../@rel="top" */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__rel,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Filter{
								arg1: &Variable{
									ARG: variable__node,
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
						ARG: elem__string__top,
						arg1: &Collect{
							ARG: collect__attributes__rel,
							arg1: &Collect{
								ARG: collect__parent__type__node,
								arg1: &Filter{
									arg1: &Variable{
										ARG: variable__node,
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
		if Test(node, q /* $node[@ud:pos="PROPN" and not(../node[@cat]) ] */, &XPath{
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
									ARG: elem__string__PROPN,
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
		if Test(node, q /* $node[@ud:pos=("ADP","ADV","ADJ","DET","PRON","CCONJ","NOUN","VERB","INTJ")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
					},
					arg2: &Sort{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &Node{},
							},
							arg2: &Elem{
								ARG: elem__string__ADP__ADV__ADJ__DET__PRON__CCONJ__NOUN__VERB__INT_2e_2e_2e,
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
			}) {
				return "mark" // absolute met constructie -- analoog aan with X being Y
			}
			if Test(node, q /* not($node/../node[@rel="pc"]) */, &XPath{
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
												ARG: elem__string__pc,
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
		if Test(node, q, /* $node[(@ud:pos=("ADJ","X","ADV") or @cat="mwu")
			   and ../@cat="pp"
			   and ../node[@rel=("obj1","se","vc")]] */&XPath{
				arg1: &Sort{
					arg1: &Filter{
						arg1: &Variable{
							ARG: variable__node,
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
													ARG: elem__string__ADJ__X__ADV,
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
													ARG: elem__string__mwu,
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
											ARG: elem__string__pp,
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
				},
			}) {
			if Test(node, q /* $node[../@rel="cnj" and ../node[@rel="obj1" and @index and not(@pt or @cat)]] */, &XPath{
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
										ARG: collect__attributes__rel,
										arg1: &Collect{
											ARG:  collect__parent__type__node,
											arg1: &Node{},
										},
									},
									arg2: &Elem{
										ARG: elem__string__cnj,
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
														ARG: elem__string__obj1,
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
										ARG: elem__string__hd,
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
			return "conj"
		}
		return dependencyLabel(node.parent, q)
	}
	return "ERROR_NO_LABEL"
}

func determineNominalModLabel(node *NodeType, q *Context) string {
	if Test(node, q /* $node/../node[@rel="hd" and (@ud:pos="VERB" or @ud:pos="ADJ")] */, &XPath{
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
								ARG: elem__string__hd,
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
										ARG: elem__string__VERB,
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
										ARG: elem__string__ADJ,
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
	if SUTest(subj, q, /* $subj[@cat=("whsub","ssub","ti","cp","oti")] or
		   $subj[@cat="conj" and node/@cat=("whsub","ssub","ti","cp","oti")] */&XPath{
			arg1: &Sort{
				arg1: &Or{
					arg1: &Filter{
						arg1: &Variable{
							ARG: variable__subj,
						},
						arg2: &Sort{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__cat,
									arg1: &Node{},
								},
								arg2: &Elem{
									ARG: elem__string__whsub__ssub__ti__cp__oti,
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
							ARG: variable__subj,
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
										ARG: elem__string__whsub__ssub__ti__cp__oti,
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
	if SUTest(subj, q, /* $subj[@lemma="het"] and
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
							ARG: variable__subj,
						},
						arg2: &Sort{
							arg1: &Equal{
								ARG: equal__is,
								arg1: &Collect{
									ARG:  collect__attributes__lemma,
									arg1: &Node{},
								},
								arg2: &Elem{
									ARG: elem__string__het,
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
											ARG: variable__subj,
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
													ARG: elem__string__hd,
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
													ARG: elem__string__dooien__gieten__hagelen__miezeren__misten__mo_2e_2e_2e,
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
											ARG: variable__subj,
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
														ARG: elem__string__hd,
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
														ARG: elem__string__ontbreken,
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
																ARG: elem__string__pc,
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
																			ARG: elem__string__hd,
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
																			ARG: elem__string__aan,
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
									ARG: variable__subj,
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
															ARG: elem__string__sup,
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

	aux := auxiliary1(firstnode(SUFind(subj, q /* $subj/../node[@rel="hd"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
				arg1: &Collect{
					ARG: collect__parent__type__node,
					arg1: &Variable{
						ARG: variable__subj,
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
							ARG: elem__string__hd,
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
	if aux == "aux:pass" { // de carriere had gered kunnen worden
		return ":pass"
	}
	if aux == "aux" && SUTest(subj, q /* $subj/@index = $subj/../node[@rel="vc"]/node[@rel="su"]/@index */, &XPath{
		arg1: &Sort{
			arg1: &Equal{
				ARG: equal__is,
				arg1: &Collect{
					ARG: collect__attributes__index,
					arg1: &Variable{
						ARG: variable__subj,
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
									ARG: variable__subj,
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
	}) {
		return passiveSubject(firstnode(SUFind(subj, q /* $subj/../node[@rel="vc"]/node[@rel="su"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Collect{
						ARG: collect__child__node,
						arg1: &Collect{
							ARG: collect__parent__type__node,
							arg1: &Variable{
								ARG: variable__subj,
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
		})), q)
	}
	return ""
}

func detLabel(node *NodeType, q *Context) string {
	// zijn boek, diens boek, ieders boek, aller landen, Ron's probleem, Fidel Castro's belang
	if Test(node, q, /* $node[@ud:pos = "PRON" and @vwtype="bez"] or
		   $node[@ud:pos = ("PRON","PROPN") and @naamval="gen"] or
		   $node[@cat="mwu" and node[@spectype="deeleigen"]] */&XPath{
			arg1: &Sort{
				arg1: &Or{
					arg1: &Or{
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
											ARG: elem__string__PRON,
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
											ARG: elem__string__bez,
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
											ARG: elem__string__PRON__PROPN,
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
											ARG: elem__string__gen,
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
							ARG: variable__node,
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
										ARG: elem__string__mwu,
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
												ARG: elem__string__deeleigen,
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
	if Test(node, q /* $node/@ud:pos = ("DET","PROPN","NOUN","ADJ","PRON","ADV","X") */, &XPath{
		arg1: &Sort{
			arg1: &Equal{
				ARG: equal__is,
				arg1: &Collect{
					ARG: collect__attributes__ud_3apos,
					arg1: &Variable{
						ARG: variable__node,
					},
				},
				arg2: &Elem{
					ARG: elem__string__DET__PROPN__NOUN__ADJ__PRON__ADV__X,
					arg1: &Collect{
						ARG: collect__attributes__ud_3apos,
						arg1: &Variable{
							ARG: variable__node,
						},
					},
				},
			},
		},
	}) {
		return "det" // meer  genoeg the
	}
	if Test(node, q /* $node/@cat = ("mwu","np","pp","ap","detp","smain") */, &XPath{
		arg1: &Sort{
			arg1: &Equal{
				ARG: equal__is,
				arg1: &Collect{
					ARG: collect__attributes__cat,
					arg1: &Variable{
						ARG: variable__node,
					},
				},
				arg2: &Elem{
					ARG: elem__string__mwu__np__pp__ap__detp__smain,
					arg1: &Collect{
						ARG: collect__attributes__cat,
						arg1: &Variable{
							ARG: variable__node,
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
		if Test(node, q /* $node/node[@rel="cnj"][1]/@ud:pos="NUM" */, &XPath{
			arg1: &Sort{
				arg1: &Equal{
					ARG: equal__is,
					arg1: &Collect{
						ARG: collect__attributes__ud_3apos,
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
											ARG: elem__string__cnj,
											arg1: &Collect{
												ARG:  collect__attributes__rel,
												arg1: &Node{},
											},
										},
									},
								},
								arg2: &Elem{
									ARG: elem__number__1,
								},
							},
						},
					},
					arg2: &Elem{
						ARG: elem__string__NUM,
						arg1: &Collect{
							ARG: collect__attributes__ud_3apos,
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
												ARG: elem__string__cnj,
												arg1: &Collect{
													ARG:  collect__attributes__rel,
													arg1: &Node{},
												},
											},
										},
									},
									arg2: &Elem{
										ARG: elem__number__1,
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
	if Test(node, q /* $node[@cat="pp"]/node[@rel="vc"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
								ARG: elem__string__pp,
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
							ARG: elem__string__vc,
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
	if Test(node, q /* $node[@ud:pos="ADJ" or @cat="ap" or node[@cat="conj" and node[@ud:pos="ADJ" or @cat="ap"] ]] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
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
									ARG: elem__string__ADJ,
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
									ARG: elem__string__ap,
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
											arg1: &Or{
												arg1: &Equal{
													ARG: equal__is,
													arg1: &Collect{
														ARG:  collect__attributes__ud_3apos,
														arg1: &Node{},
													},
													arg2: &Elem{
														ARG: elem__string__ADJ,
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
														ARG: elem__string__ap,
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
	if Test(node, q /* $node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM","INTJ") ] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
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
								ARG: elem__string__pp__np__conj__mwu,
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
								ARG: elem__string__NOUN__PRON__PROPN__X__PUNCT__SYM__INTJ,
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
	if Test(node, q /* $node[@cat="detp"]/node[@rel="hd" and @ud:pos="NUM"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
								ARG: elem__string__detp,
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
								ARG: elem__string__hd,
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
								ARG: elem__string__NUM,
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
	if Test(node, q /* $node[@cat="cp"]/node[@rel="body" and (@ud:pos = ("NOUN","PROPN") or @cat=("np","conj"))] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
								ARG: elem__string__cp,
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
								ARG: elem__string__body,
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
										ARG: elem__string__NOUN__PROPN,
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
										ARG: elem__string__np__conj,
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
	if Test(node, q /* $node[@cat=("cp","sv1","smain","ppres","ppart","inf","ti","oti","du","whq") or @ud:pos="SCONJ"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
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
								ARG: elem__string__cp__sv1__smain__ppres__ppart__inf__ti__oti__du__w_2e_2e_2e,
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
								ARG: elem__string__SCONJ,
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
	if Test(node, q /* $node[@ud:pos= ("ADV","ADP","VERB","CCONJ") or @cat="advp"] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
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
								ARG: elem__string__ADV__ADP__VERB__CCONJ,
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
								ARG: elem__string__advp,
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
	if Test(node, q /* $node[@cat="pp"]/node[@rel="vc"] */, &XPath{
		arg1: &Sort{
			arg1: &Collect{
				ARG: collect__child__node,
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
								ARG: elem__string__pp,
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
							ARG: elem__string__vc,
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
	if Test(node, q, /* $node[ (  node[@rel="hd" and @lemma="door"]
		    or (@pt="bw" and ends-with(@lemma,"door"))
		    )
		    and ../self::*[@cat="ppart"]/../node[@rel="hd" and @lemma=("zijn","worden")]
		    and ../../node[@rel="su"]/@index = ../node[@rel="obj1"]/@index
		] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
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
															ARG: elem__string__hd,
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
															ARG: elem__string__door,
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
														ARG: elem__string__bw,
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
															ARG: elem__string__door,
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
											ARG: collect__self__all__node,
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
														ARG: elem__string__ppart,
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
													ARG: elem__string__hd,
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
													ARG: elem__string__zijn__worden,
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
	if Test(node, q /* $node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM") ] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__node,
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
								ARG: elem__string__pp__np__conj__mwu,
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
								ARG: elem__string__NOUN__PRON__PROPN__X__PUNCT__SYM,
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
	if Test(node, q /* $node[@cat=("cp","sv1","smain","ssub","ppres","ppart","ti","oti","inf","du","whq","whrel","rel")] */, &XPath{
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
							ARG: elem__string__cp__sv1__smain__ssub__ppres__ppart__ti__oti__inf_2e_2e_2e,
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
	if Test(node, q, /* $node[@ud:pos= ("ADJ","ADV","ADP","VERB","SCONJ","INTJ")
		   or @cat=("advp","ap")
		   or (@cat="conj" and node/@ud:pos="ADV")] */&XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__node,
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
										ARG: elem__string__ADJ__ADV__ADP__VERB__SCONJ__INTJ,
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
										ARG: elem__string__advp__ap,
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
											ARG: elem__string__conj,
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
											ARG: elem__string__ADV,
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
		if HGTest(head, gap, q /* $head/node[@rel="obj1"] */, &XPath{
			arg1: &Sort{
				arg1: &Collect{
					ARG: collect__child__node,
					arg1: &Variable{
						ARG: variable__head,
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
		}) {
			return "nmod"
		}
		if HGTest(head, gap, q /* $head[@ud:pos=("ADV", "ADP") or @cat=("advp","ap")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__head,
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
									ARG: elem__string__ADV__ADP,
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
									ARG: elem__string__advp__ap,
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
	if HGTest(head, gap, q /* $gap[@rel="mod" and ../@cat=("np","pp")] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__gap,
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
								ARG: elem__string__mod,
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
								ARG: elem__string__np__pp,
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
	if HGTest(head, gap, q /* $gap[@rel="mod" and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")] */, &XPath{
		arg1: &Sort{
			arg1: &Filter{
				arg1: &Variable{
					ARG: variable__gap,
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
								ARG: elem__string__mod,
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
								ARG: elem__string__sv1__smain__ssub__inf__ppres__ppart__oti__ap__ad_2e_2e_2e,
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
		if HGTest(head, gap, q /* $head[@cat=("pp","np") or @ud:pos=("NOUN","PRON")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__head,
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
									ARG: elem__string__pp__np,
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
									ARG: elem__string__NOUN__PRON,
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
		if HGTest(head, gap, q /* $head[@ud:pos=("ADV","ADP") or @cat= ("advp","ap","mwu","conj")] */, &XPath{
			arg1: &Sort{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__head,
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
									ARG: elem__string__ADV__ADP,
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
									ARG: elem__string__advp__ap__mwu__conj,
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
	if HGTest(head, gap, q /* $gap[@rel="hd"] and $head[@ud:pos=("ADV","ADP")] */, &XPath{
		arg1: &Sort{
			arg1: &And{
				arg1: &Filter{
					arg1: &Variable{
						ARG: variable__gap,
					},
					arg2: &Sort{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__rel,
								arg1: &Node{},
							},
							arg2: &Elem{
								ARG: elem__string__hd,
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
						ARG: variable__head,
					},
					arg2: &Sort{
						arg1: &Equal{
							ARG: equal__is,
							arg1: &Collect{
								ARG:  collect__attributes__ud_3apos,
								arg1: &Node{},
							},
							arg2: &Elem{
								ARG: elem__string__ADV__ADP,
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

// recursive
func externalHeadPosition(node *NodeType, q *Context) int {
	if depthCheck(q, "externalHeadPosition") {
		return ERROR_NO_EXTERNAL_HEAD
	}

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
		}) { // TODO: dit kan efficiënter: meerdere keren zoeken naar zelfde set nodes
			// met als titel
			return internalHeadPosition(firstnode(Find(node, q /* $node/../node[@rel="predc"] */, &XPath{
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
								ARG: elem__string__obj1__vc__se__me,
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
										ARG: elem__string__obj1__vc__se__me,
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
				return internalHeadPositionWithGapping(firstnode(Find(node, q /* $node/../node[@rel=("obj1","vc","se","me")] */, &XPath{
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
										ARG: elem__string__obj1__vc__se__me,
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
										ARG: elem__string__obj1__vc__se__me,
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
														ARG: elem__string__rhd__whd,
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
				return internalHeadPosition(firstnode(Find(node, q, /* $node/ancestor::node/node[@rel=("rhd","whd")
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
												ARG: elem__string__rhd__whd,
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
																ARG: elem__string__obj1__vc__se__me,
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
									ARG: elem__string__pobj1,
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
				return internalHeadPosition(firstnode(Find(node, q /* $node/../node[@rel="pobj1"] */, &XPath{
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
										ARG: elem__string__pobj1,
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
									ARG: elem__string__vc__predc,
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
			return internalHeadPositionWithGapping(firstnode(Find(node, q /* $node/../node[@rel=("vc","predc")] */, &XPath{
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
									ARG: elem__string__vc__predc,
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
								ARG: elem__string__obj1__pobj1__se,
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
			return internalHeadPosition(firstnode(Find(node, q /* $node/node[@rel=("obj1","pobj1","se")] */, &XPath{
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
									ARG: elem__string__obj1__pobj1__se,
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
								ARG: elem__string__hd,
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
			return internalHeadPosition(firstnode(Find(node, q /* $node/node[@rel="hd"] */, &XPath{
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
									ARG: elem__string__hd,
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
		if len(node.Node) > 0 {
			return internalHeadPosition(node.Node[0], q)
		}
		return internalHeadPosition(noNode, q)
	}

	if node.Cat == "mwu" {
		return firstint(Find(node, q /* $node/node[@rel="mwp" and not(../node/@begin < @begin)]/@end */, &XPath{
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
										ARG: elem__string__mwp,
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
		}))
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

// recursive ??
func enhancedDependencies(q *Context) {
	// TODO

	// na elke invoeging van nieuwe node: inspect(q) ??
}

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

////////////////////////////////////////////////////////////////

func leftmost(nodes []interface{}) *NodeType {
	switch len(nodes) {
	case 0:
		return noNode
	case 1:
		return nodes[0].(*NodeType)
	}
	sort.Slice(nodes, func(i, j int) bool {
		// solve cases where begin is identical (hij is en blijft omstreden)??
		ii := nodes[i].(*NodeType)
		jj := nodes[j].(*NodeType)
		if ii.Begin != jj.Begin {
			return ii.Begin < jj.Begin // ints
		}
		if ii.End != jj.End {
			return ii.End < jj.End // ints
		}
		return ii.Id > jj.Id // ints
	})
	return nodes[0].(*NodeType)
}

func firstnode(nodes []interface{}) *NodeType {
	return leftmost(nodes)
	/*
		if len(nodes) > 0 {
			return nodes[0].(*NodeType)
		}
		return noNode
	*/
}

func firstint(ii []interface{}) int {
	if len(ii) > 0 {
		return ii[0].(int)
	}
	return ERROR_NO_VALUE
}

func depthCheck(q *Context, s string) bool {
	q.depth++
	if q.depth < 1000 {
		return false
	}
	q.warnings = append(q.warnings, "Recursion depth limit for "+s)
	fmt.Fprintln(os.Stderr, "WARNING: Recursion depth limit for %s in %s", s, q.filename)
	return true
}
