
Dit demonstreerde een bug in het gebruik van indexen, zoals in `$node/node/node[@id][1]`

*Bug is inmiddels verholpen, maar code is misschien nuttig op een later moment.*

```bash
~ xqilla test1.xq
<node id="3"/>
<node id="6"/>

~ xqilla test2.xq
<node id="3"/>
<node id="6"/>

~ xqilla test3.xq
<node id="3"/>

~ ./main 
Test 1
id= 3
id= 6
Test 2
id= 3
id= 4
id= 6
id= 7
Test 3
id= 3
```

De Go-versie ~~geeft~~ gaf in test 2 vier resultaten in plaats van twee.
