for $doc in doc("file.xml") for $node in $doc/alpino_ds/node
  return $node/node[@id]/node[@id][1]
