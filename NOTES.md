Het programma `compile`, dat de Go-bestanden met XPath-expressies compileert
naar zuiver Go, maakt gebruik van het programma `testXPath` dat onderdeel is van
de sources van `libxml2`. Dit is getest met `libxml2` versie 2.9.9.
Een oudere versie werkte niet goed.

----

XPath-expressies in Go kunnen gebruikt worden in de functies `FIND()`
en `TEST()`, waarin de expressie zelf tussen backquotes dient te
staan. Met gewone, dubbele aanhalingstekens werkt het niet.

----

Alleen de indexen `[1]` en `[last()]` zijn geïmplementeerd. Andere
indexen, zoals `[2]` of `[last()-1]` zullen niet werken of een
verkeerd resultaat geven, tenzij de implementatie wordt toegevoegd.
De index `[1]` wordt door de compiler omgezet in `[first()]`.

----

Posities (begin, end, id, head) zijn geïmplementeerd als integers.
Houd hier rekening mee als je wilt weten hoeveel posities twee woorden
van elkaar af liggen.

interne waarde | uitvoer
-------------- | -------------
1000           | 1
2003           | 2.3
4050           | 4.50

----

De compiler maakt gebruik van libxml2, dat alleen XPath versie 1 kan verwerken.

De compiler herkent reeksen zoals deze, en codeert die op een wijze
die libxml2 kan verwerken:

```
    = ("aap", "noot")
```

Het isgelijk-teken en de dubbele aanhalingstekens zijn nodig. Deze reeksen
worden niet herkend:

```
    > ("aap", "noot")
    = ('aap', 'noot')
```

----

Er zit nogal wat complexiteit in de verwerking van predicaten. Je hebt
de ‘gewone’ predicaten, die werken op elke node individueel, en je
hebt de index-predicaten, die werken op een of meer sets van nodes.
Nog complexer is het wanneer de beide gecombineerd worden. Om de
juiste resultaten te krijgen heb ik hier en daar wat houtje-touwtje
gebruikt in de methode `do` van `dCollect`, in de hoop dat de boel niet
breekt. Er is vast een betere manier, maar dan niet zonder simpel
platte sets van nodes doorschuiven van de ene naar de volgende stap
zoals nu gebeurt.

