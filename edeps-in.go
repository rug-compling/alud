// +build ignore

package main

func addEdependencyRelations(q *Context) {
	for _, node := range q.ptnodes {
		node.udERelation = dependencyLabel(node, q)
		node.udEHeadPosition = externalHeadPosition(node, q)
	}
}

func enhancedDependencies1(q *Context) {
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
	                   if ($node[@ud:ERelation=("nsubj","obj","iobj","nsubj:pass") and exists($so_index)] )
	                   then local:enhanced_elements_to_string(($self,
	                                                                 local:xcomp-control($node,$so_index),
	                                                                   local:upstairs-control($node,$so_index),
	                                                                     local:passive-vp-control($node,$so_index)) )
	                   else if (exists($rhd_index))
	                        then if (exists($rhd_np))
	                             then local:enhanced_elements_to_string(($rhdref,
	                                                                      local:xcomp-control($node,$rhd_index),
	                                                                        local:passive-vp-control($node,$rhd_index)) )
	                             else local:enhanced_elements_to_string(($self,    (: if there is no antecedent, lets keep the basic relation :)
	                                            local:xcomp-control($node,$rhd_index),
	                                              local:passive-vp-control($node,$rhd_index)) )
	                        else if ($rel_sister_index)
	                 then local:enhanced_elements_to_string(($self,
	                                                         local:xcomp-control($node,$rel_sister_index),
	                                                         local:passive-vp-control($node,$rel_sister_index)) )
	                 else if  ($node/@ud:HeadPosition)
	                                  then local:enhanced_elements_to_string($self)
	                                  else ""
	           return
	           <node ud:Enhanced="{$enhanced}">
	              {( $node/@*,
	                       for $child in $node/node return local:enhanced_dependencies1($child)
	          )}
	     </node>
	   };
	*/
	// TODO
}
