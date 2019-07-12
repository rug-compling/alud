# Van Alpino naar Universal Dependencies

Een poging
[universal_dependencies.xq](https://github.com/gossebouma/lassy2ud) om
te zetten naar Go.

----

Met het programma `testXPath` (zit bij de bronbestanden van libxml2) kun je xpath-expressies compileren en
het resultaat printen. Voorbeeld:

```
testXPath --tree '$doc/aap[@noot]'
```

Resultaat:

```
  SORT
    COLLECT  'child' 'name' 'node' aap
      VARIABLE doc
      PREDICATE
        COLLECT  'attributes' 'name' 'node' noot
          NODE
```

Zie functie `xmlXPathDebugDumpStepOp` in `libxml2-`_*_`/xpath.c`

Die uitvoer laat zich makkelijk omzetten naar code in Go:

```go
func process(doc *Obj) (interface{}, error) {
    xpath := &XPath{
        arg1: &Sort{
            arg1: &Collect{
                ARG: collect_child_name_node_aap,
                arg1: &Variable{
                    ARG: variable_doc,
                },
                arg2: &Predicate{
                    arg1: &Collect{
                        ARG:  collect_attributes_name_node_noot,
                        arg1: &Node{},
                    },
                },
            },
        },
    }
    return xpath.Do(doc)
}
```

Vergelijk
[invoer](https://github.com/pebbe/unidep/blob/master/auxiliary-in.go)
met
[uitvoer](https://github.com/pebbe/unidep/blob/master/auxiliary.go)



----

## Tests

taak \ tijd | XQilla | Go
------- | ------:| -----:
Eindhoven-corpus, fix misplaced heads + add postags + add features | 08:06 | 00:29
Eindhoven-corpus, alles behalve enhanced dependencies en fixpunct | 23:00 | 01:24

----

## Verschillen tussen Saxon en XQilla

De resultaten van Saxon en XQilla verschillen van elkaar, in de
toevoeging die enhanced UDs krijgen. Bijvoorbeeld, bestand
`1057.xml` uit het corpus Alpino Treebank, daarin krijgt het woord
*parlementsverkiezingen* van Saxon de DEPS `9:advcl:evenals` en van XQilla
`9:advcl:bij`. 

Nog een paar bestanden die verschillen opleveren, uit hetzelde corpus:

 * 1017.xml
 * 1022.xml
 * 1024.xml
 * 1055.xml

In totaal levert dit corpus 525 verschillen.

De Go-versie geeft op dit punt dezelfde resultaten als Saxon.

----

## fix\_misplaced\_heads\_in\_coordination

Hoe vaak moet de functie `fix_misplaced_heads_in_coordination`
gebruikt worden? In het XQuery-script wordt het twee keer gedaan voor
uitvoer naar conll-formaat, en één keer voor uitvoer naar XML-formaat.

In mijn programma gebruik ik de functie zo vaak als nodig tot er geen verschillen
meer optreden. Voor sommige zinnen gaat dat niet goed omdat dingen
blijvend heen en weer worden geschoven. Dan breek ik het af, en hoop
op het beste.

Zinnen waarbij dit het geval is leveren nogal eens ongeldige
resultaten, ook bij verwerking door XQilla.

Dit komt niet voor in het Eindhoven-corpus, maar wel vaak in het corpus
LassySmall, zo'n 13 keer:

 1. WR-P-E-I-0000006366.p.8.s.4 — Invalid HEAD value ERROR\_NO\_INTERNAL\_HEAD\_IN\_GAPPED\_CONSTITUENT
 1. WR-P-E-I-0000020972.p.4.s.143 — An element has two attributes with the same expanded name
 1. WR-P-E-I-0000020972.p.4.s.150
 1. WR-P-E-I-0000020972.p.4.s.164 — An element has two attributes with the same expanded name
 1. WR-P-E-I-0000020972.p.4.s.177 — An element has two attributes with the same expanded name
 1. WR-P-E-I-0000020972.p.4.s.192 — An element has two attributes with the same expanded name
 1. WR-P-E-I-0000020972.p.4.s.200 — An element has two attributes with the same expanded name
 1. WR-P-E-I-0000020972.p.4.s.215 — An element has two attributes with the same expanded name
 1. WR-P-E-I-0000050211.p.1.s.188 — Unordered ID numbers 20.1, 20.1
 1. WR-P-E-I-0000050381.p.1.s.682.2
 1. WR-P-P-I-0000000098.p.4.s.2
 1. WR-P-P-I-0000000106.p.17.s.4
 1. dpc-ibm-001316-nl-sen.p.37.s.1

----

## Verschillen tussen ~~XQilla~~ Saxon en Go

De Go-versie geeft momenteel drie verschillen voor **DEPREL** ten opzichte van XQilla
voor het Eindhoven-corpus.

Daarnaast zijn er ook wat verschillen voor **DEPS**. De Go-versie krijgt hier soms
meer waardes.

Wat DEPREL betreft: voor
[drie verschillen](https://paqu.let.rug.nl:8068/xpath?db=eindhoven&xpath=%2F%2Fsentence%5B%40sentid%3D%28"cdb-6322"%2C"gbl-5437"%2C"obl-594"%29%5D),
hierin komen deze combinaties voor:

 * zestig- a zeventigduizend katholieken
 * vijf- tot zeshonderd bladzijden
 * een stuk of 6 , 7

De woorden *zestig-*, *vijf-* en *6* krijgen van XQilla de DEPREL
`det`. Van de Go-versie krijgen ze de waarde `nummod`.

In het xquery-script komt deze test voor:

```
    then if ($node/node[@rel="cnj"][1]/@ud:pos="NUM" )
        then "nummod"
        else "det"
```

Als ik in de Go-versie `[1]` vervang door `[last()]` dan verdwijnen de
verschillen. Het heeft dus(?) te maken met hoe de implementatie intern
de tussenresultaten ordent.

Wat DEPS betreft: De extra waardes lijken te komen door deze expressie
in de functie anaphoricRelpronoun():

```
    $node/ancestor::node[@cat="np" and local:internal_head_position(.) = $node/@end]/
        node[@rel="mod"]/node[@rel="rhd"]/descendant-or-self::node[@pt="vnw" and not(@ud:HeadPosition = $node/@end)][last()]
```

Blijkbaar moet voor XPath versie 1 de hele expresie voor de laatste
index tussen haakjes:

```
    ($node/ancestor::node[@cat="np" and local:internal_head_position(.) = $node/@end]/
        node[@rel="mod"]/node[@rel="rhd"]/descendant-or-self::node[@pt="vnw" and not(@ud:HeadPosition = $node/@end)])[last()]
```

In dit geval zijn de extra expressies weg, maar soms verdwijnt de
verkeerde. Hier hebben we dus(?) weer te maken met verschillen in
volgorde.
