
Dit demonstreert een bug in het gebruik van indexen, zoals in `$node/node/node[1]`

```bash
~ xqilla test1.xq 
<node id="3"/>
<node id="6"/>

~ xqilla test2.xq 
<node id="3"/>

~ ./main 
Test 1
id= 3
id= 4
id= 6
id= 7
Test 2
id= 3
```

De Go-versie geeft in test 1 vier resultaten in plaats van twee.

Met `xpath.go.test` heb ik geprobeerd het op te lossen, maar er
ontstaan alleen maar meer bugs.

Een deel van de oplossing is om een speciaal type te gebruiken voor
indexen, om die te kunnen onderscheiden van gewone integers. Dat kan
door `[1]` te vervangen door `[first()]`. De functies `first()` en
`last()` geven geen `int` maar een `IndexType`.
