for $doc in doc("file.xml") for $node in $doc/alpino_ds/node
  let $nodes := $node/node
  return $nodes[@id][last()]/node[@id>0][1]
