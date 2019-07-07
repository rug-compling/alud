// +build ignore

package main

import (
	"sort"
	"strings"
)

type DepType struct {
	head int
	rel  string
}

func enhancedDependencies(q *Context) {

	reconstructEmptyHead(q)

	for _, node := range q.ptnodes {
		node.udERelation = dependencyLabel(node, q)
		node.udEHeadPosition = externalHeadPosition(node, q)
	}

	for _, node := range q.ptnodes {
		enhancedDependencies1(node, q)
	}
}

func enhancedDependencies1(node *NodeType, q *Context) {
	/*
	   declare function local:enhanced_dependencies1($node as node()) as node() {
	           (: iobj2 control : de commissie kan de raad aanbevelen/adviseren/ X te doen :)
	           (: rhd : een levend visje dat doorgeslikt moet worden :)
	           let $so_index :=
	               $node/ancestor::node/node[@rel=("su","obj1","obj2") and local:internal_head_position(.) = $node/@end ]/@index
	           let $rhd_index :=
	               $node/ancestor::node/node[@rel="rhd" and local:internal_head_position(.) = $node/@end ]/@index
	           let $rhd_np := $node/ancestor::node[@cat="np" and node[@rel="mod"]/node[@rel="rhd"]/@index = $rhd_index]
	           (: de enige _i die voldoet aan de eisen -- make sure empty heads are covered as well :)
	           let $rhdref := <headdep head="{local:internal_head_position_with_gapping($rhd_np)}" dep="ref"/>
	           let $self := ( <headdep head="{$node/@ud:EHeadPosition}" dep="{local:enhance_dependency_label($node)}"/> ,
	                                 local:anaphoric_relpronoun($node),
	                                    local:distribute_conjuncts($node),
	                                        local:distribute_dependents($node) )
	     let $rel_sister_index :=  ($node/../node[@rel="mod" and @cat="rel"]/node[@rel="rhd"]/@index)[1]
	           let $enhanced :=
	                   1if ($node[@ud:ERelation=("nsubj","obj","iobj","nsubj:pass") and exists($so_index)] )
	                   1then local:enhanced_elements_to_string(($self,
	                                                                 local:xcomp-control($node,$so_index),
	                                                                   local:upstairs-control($node,$so_index),
	                                                                     local:passive-vp-control($node,$so_index)) )
	                   1else 2if (exists($rhd_index))
	                        2then 3if (exists($rhd_np))
	                             3then local:enhanced_elements_to_string(($rhdref,
	                                                                      local:xcomp-control($node,$rhd_index),
	                                                                        local:passive-vp-control($node,$rhd_index)) )
	                             3else local:enhanced_elements_to_string(($self,    (: if there is no antecedent, lets keep the basic relation :)
	                                            local:xcomp-control($node,$rhd_index),
	                                              local:passive-vp-control($node,$rhd_index)) )
	                        2else 4if ($rel_sister_index)
	                 4then local:enhanced_elements_to_string(($self,
	                                                         local:xcomp-control($node,$rel_sister_index),
	                                                         local:passive-vp-control($node,$rel_sister_index)) )
	                 4else 5if  ($node/@ud:HeadPosition)
	                                  5then local:enhanced_elements_to_string($self)
	                                  5else ""
	           return
	           <node ud:Enhanced="{$enhanced}">
	              {( $node/@*,
	                       for $child in $node/node return local:enhanced_dependencies1($child)
	          )}
	     </node>
	   };
	*/

	var enhanced []DepType
	for {

		so := FIND(node, q,
			`$node/ancestor::node/node[@rel=("su","obj1","obj2") and local:internal_head_position(.) = $node/@end ]/@index`)
		if len(so) > 0 && TEST(node, q, `$node[@ud:ERelation=("nsubj","obj","iobj","nsubj:pass")]`) {
			soIndex := i1(so)
			enhanced = []DepType{
				DepType{head: node.udEHeadPosition, rel: enhanceDependencyLabel(node, q)},
				anaphoricRelpronoun(node, q),
				distributeConjuncts(node, q),
				distributeDependents(node, q),
				xcompControl(node, q, soIndex),
				upstairsControl(node, q, soIndex),
				passiveVpControl(node, q, soIndex),
			}
			break
		}

		rhd := FIND(node, q,
			`$node/ancestor::node/node[@rel=("su","obj1","obj2") and local:internal_head_position(.) = $node/@end ]/@index`)
		if len(rhd) > 0 {
			rhdIndex := i1(rhd)
			q.varrhdindex = []interface{}{rhdIndex}
			rhdNp := FIND(node, q, `$node/ancestor::node[@cat="np" and node[@rel="mod"]/node[@rel="rhd"]/@index = $rhd_index]`)
			if len(rhdNp) > 0 {
				enhanced = []DepType{
					DepType{head: internalHeadPositionWithGapping(rhdNp, q), rel: "ref"},
					xcompControl(node, q, rhdIndex),
					passiveVpControl(node, q, rhdIndex),
				}
				break
			}
			enhanced = []DepType{
				DepType{head: node.udEHeadPosition, rel: enhanceDependencyLabel(node, q)},
				anaphoricRelpronoun(node, q),
				distributeConjuncts(node, q),
				distributeDependents(node, q),
				xcompControl(node, q, rhdIndex),
				passiveVpControl(node, q, rhdIndex),
			}
			break
		}

		relSister := FIND(node, q, `$node/../node[@rel="mod" and @cat="rel"]/node[@rel="rhd"]/@index`)
		if len(relSister) > 0 {
			relSisterIndex := i1(relSister)
			enhanced = []DepType{
				DepType{head: node.udEHeadPosition, rel: enhanceDependencyLabel(node, q)},
				anaphoricRelpronoun(node, q),
				distributeConjuncts(node, q),
				xcompControl(node, q, relSisterIndex),
				passiveVpControl(node, q, relSisterIndex),
			}
			break
		}

		if node.udHeadPosition > 0 {
			enhanced = []DepType{
				DepType{head: node.udEHeadPosition, rel: enhanceDependencyLabel(node, q)},
				anaphoricRelpronoun(node, q),
				distributeConjuncts(node, q),
			}
			break
		}

		enhanced = []DepType{}
		break
	}

	sort.Slice(enhanced, func(i, j int) bool {
		if enhanced[i].head != enhanced[j].head {
			return enhanced[i].head < enhanced[j].head
		}
		return enhanced[i].rel < enhanced[j].rel
	})
	for i := 1; i < len(enhanced); i++ {
		if enhanced[i].head == enhanced[i-1].head && enhanced[i].rel == enhanced[i-1].rel {
			enhanced = append(enhanced[:i], enhanced[1+i:]...)
			i--
		}
	}
	ss := make([]string, 0, len(enhanced))
	for _, e := range enhanced {
		if e.rel != "" {
			ss = append(ss, number(e.head)+":"+e.rel)
		}
	}
	node.udEnhanced = strings.Join(ss, "|")

}

func enhanceDependencyLabel(node *NodeType, q *Context) string {
	return "TODO"
}

func anaphoricRelpronoun(node *NodeType, q *Context) DepType {
	return DepType{head: TODO, rel: "TODO"}
}

func distributeConjuncts(node *NodeType, q *Context) DepType {
	return DepType{head: TODO, rel: "TODO"}
}

func distributeDependents(node *NodeType, q *Context) DepType {
	return DepType{head: TODO, rel: "TODO"}
}

func xcompControl(node *NodeType, q *Context, idx int) DepType {
	return DepType{head: TODO, rel: "TODO"}
}

func upstairsControl(node *NodeType, q *Context, idx int) DepType {
	return DepType{head: TODO, rel: "TODO"}
}

func passiveVpControl(node *NodeType, q *Context, idx int) DepType {
	return DepType{head: TODO, rel: "TODO"}
}
