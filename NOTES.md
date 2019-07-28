XPath-expressies in Go kunnen gebruikt worden in de functies `FIND()`
en `TEST()`, waarin de expressie zelf tussen backquotes dient te
staan.

Alleen de indexen `[1]` en `[last()]` zijn geïmplementeerd. Andere
indexen, zoals `[2]` of `[last()-1]` zullen niet werken of een
verkeerd resultaat geven, tenzij de implementatie wordt toegevoegd.
De index `[1]` wordt door de compiler omgezet in `[first()]`.

Posities (begin, end, id, head) zijn geïmplementeerd als integers:

interne waarde | uitvoer
-------------- | -------------
1000           | 1
2004           | 2.4
3050           | 3.50
