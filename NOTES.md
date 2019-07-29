XPath-expressies in Go kunnen gebruikt worden in de functies `FIND()`
en `TEST()`, waarin de expressie zelf tussen backquotes dient te
staan.

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
2004           | 2.4
3050           | 3.50

----

Er zit nogal wat complexiteit in de verwerking van predicaten. Je hebt
de ‘gewone’ predicaten, die werken op elke node individueel, en je
hebt de index-predicaten, die werken op een of meer sets van nodes.
Nog complexer is het wanneer de beide gecombineerd worden. Om de
juiste resultaten te krijgen heb ik hier en daar wat houtje-touwtje
gebruikt in de methode `Do` van `Collect`, in de hoop dat de boel niet
breekt. Er is vast een betere manier, maar dan niet zonder simpel
platte sets van nodes doorschuiven van de ene naar de volgende stap
zoals nu gebeurt.

