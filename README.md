# Van Alpino naar Universal Dependencies

Een geslaagde poging
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

Alternatief: Saxon met optie -explain, zie http://www.saxonica.com/documentation/index.html#!using-xquery/commandline

Voorbeeld:

```
saxon -explain -opt:0 -qs:'for $doc in <xml/> return $doc/aap[@noot]'
```

Resultaat:

```xml
<query>
   <globalVariables/>
   <body>
      <FLWOR baseUri="file:/home/peter/tmp/"
             ns="fn=~ xs=~ saxon=~ err=~ local=http://www.w3.org/2005/xquery-local-functions xsi=~"
             line="1">
         <for var="Q{}doc" slot="0">
            <elem name="xml" nsuri="">
               <empty/>
            </elem>
         </for>
         <return>
            <slash>
               <varRef name="Q{}doc" slot="0"/>
               <filter flags="">
                  <axis name="child" nodeTest="element(Q{}aap)"/>
                  <axis name="attribute" nodeTest="attribute(Q{}noot)"/>
               </filter>
            </slash>
         </return>
      </FLWOR>
   </body>
</query>
```

Andere mogelijke optie:

 * `-opt:-v`


----

## Tests

taak → tijd | XQilla | Saxon | Udapy | Go
------- | ------:| -----:| -----:| -----:
Eindhoven-corpus, fix misplaced heads + add postags + add features |  8:06 |       |       |  0:27
Eindhoven-corpus, alles behalve enhanced dependencies en fixpunct  | 23:00 |       |       |  1:02
Eindhoven-corpus, alles behalve fixpunct                           | 48:04 |  3:26 |       |  1:28
Eindhoven-corpus, alleen fixpunct                                  |       |       |  0:33 |
Eindhoven-corpus, alles                                            |       |       |       |  1:31

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

## Verschillen tussen Saxon en Go


### fix\_misplaced\_heads\_in\_coordination

Hoe vaak moet de functie `fix_misplaced_heads_in_coordination`
gebruikt worden? In het XQuery-script wordt het twee keer gedaan voor
uitvoer naar conll-formaat, en één keer voor uitvoer naar XML-formaat.

In mijn programma gebruik ik de functie zo vaak als nodig tot er geen verschillen
meer optreden. Voor sommige zinnen gaat dat niet goed omdat dingen
blijvend heen en weer worden geschoven. Dan stop ik, en hoop
op het beste, maar dan is vaak toch het eindresultaat anders dan bij
het script.

Zinnen waarbij dit het geval is leveren nogal eens ongeldige
resultaten op, ook bij verwerking door Saxon.

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

### Ingevoegde nodes met identieke IDs

Soms voegt het xquery-script meerdere nieuwe nodes in op dezelfde
plek, en geeft die hetzelfde ID. Dit is niet toegestaan.

In het programma geef ik in dit geval opeenvolgende waardes, 8.1, 8.2
etc. Dit levert uiteraard andere uitkomsten op.

### Fouten

Als er foutmeldingen zitten in de uitvoer van Saxon, dan kunnen daar
in de uitvoer van Go soms andere foutmeldingen staan.
