for $doc in doc("file.xml") for $node in $doc/alpino_ds/node
  return ($node/node/node)[1]