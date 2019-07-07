//
// GENERATED FILE -- DO NOT EDIT
//

package main

func addEdependencyRelations(q *Context) {
	for _, node := range q.ptnodes {
		node.udERelation = dependencyLabel(node, q)
		node.udEHeadPosition = externalHeadPosition(node, q)
	}
}

func enhancedDependencies1(q *Context) {
	// TODO
}
