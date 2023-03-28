# Van Alpino naar Universal Dependencies

A [Go](https://go.dev/) package for deriving [Universal Dependencies](https://universaldependencies.org/) from Dutch sentences parsed with [Alpino](https://www.let.rug.nl/vannoord/alp/Alpino/).

[![GoDoc](https://godoc.org/github.com/rug-compling/alud/v2?status.svg)](https://godoc.org/github.com/rug-compling/alud/v2)
![Technology Readiness Stage 4/4 - The technology is complete, stable and deployed in production scenarios for end-users](https://w3id.org/research-technology-readiness-levels/Stage4Complete.svg)

Example input:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<alpino_ds version="1.6">
  <node begin="0" cat="top" end="8" id="0" rel="top">
    <node begin="0" cat="smain" end="7" id="1" rel="--">
      <node begin="0" cat="pp" end="3" id="2" rel="mod">
        <node begin="0" end="1" frame="preposition(in,[])"
              his="normal" his_1="decap" his_1_1="normal" id="3"
              lcat="pp" lemma="in" pos="prep" postag="VZ(init)"
              pt="vz" rel="hd" root="in" sense="in" vztype="init"
              word="In"/>
        <node begin="1" cat="np" end="3" id="4" rel="obj1">
          <node begin="1" end="2"
                frame="determiner(het,nwh,nmod,pro,nparg,wkpro)"
                his="normal" his_1="normal" id="5" infl="het"
                lcat="detp" lemma="het" lwtype="bep" naamval="stan"
                npagr="evon" pos="det" postag="LID(bep,stan,evon)"
                pt="lid" rel="det" root="het" sense="het" wh="nwh"
                word="het"/>
          <node begin="2" end="3" frame="proper_name(sg,'LOC')"
                genus="onz" getal="ev" graad="basis" his="normal"
                his_1="names_dictionary" id="6" lcat="np"
                lemma="Waddengebied" naamval="stan" neclass="LOC"
                ntype="eigen" num="sg" pos="name"
                postag="N(eigen,ev,basis,onz,stan)" pt="n" rel="hd"
                rnum="sg" root="Waddengebied" sense="Waddengebied"
                word="Waddengebied"/>
        </node>
      </node>
      <node begin="3" end="4" frame="verb(unacc,sg_heeft,copula)"
            his="normal" his_1="normal" id="7" infl="sg_heeft"
            lcat="smain" lemma="zijn" pos="verb"
            postag="WW(pv,tgw,ev)" pt="ww" pvagr="ev" pvtijd="tgw"
            rel="hd" root="ben" sc="copula" sense="ben"
            stype="declarative" tense="present" word="is" wvorm="pv"/>
      <node begin="4" cat="np" end="6" id="8" rel="su">
        <node begin="4" end="5" frame="determiner(de)" his="normal"
              his_1="normal" id="9" infl="de" lcat="detp" lemma="de"
              lwtype="bep" naamval="stan" npagr="rest" pos="det"
              postag="LID(bep,stan,rest)" pt="lid" rel="det" root="de"
              sense="de" word="de"/>
        <node begin="5" end="6" frame="noun(de,count,sg)" gen="de"
              genus="zijd" getal="ev" graad="basis" his="normal"
              his_1="normal" id="10" lcat="np" lemma="wind"
              naamval="stan" ntype="soort" num="sg" pos="noun"
              postag="N(soort,ev,basis,zijd,stan)" pt="n" rel="hd"
              rnum="sg" root="wind" sense="wind" word="wind"/>
      </node>
      <node aform="base" begin="6" buiging="zonder" end="7"
            frame="adjective(no_e(nonadv))" graad="basis" his="normal"
            his_1="normal" id="11" infl="no_e" lcat="ap"
            lemma="veranderlijk" pos="adj" positie="vrij"
            postag="ADJ(vrij,basis,zonder)" pt="adj" rel="predc"
            root="veranderlijk" sense="veranderlijk" vform="adj"
            word="veranderlijk"/>
    </node>
    <node begin="7" end="8" frame="punct(punt)" his="normal"
          his_1="normal" id="12" lcat="punct" lemma="." pos="punct"
          postag="LET()" pt="let" rel="--" root="." sense="."
          special="punt" word="."/>
  </node>
  <sentence sentid="knmi.1">In het Waddengebied is de wind veranderlijk .</sentence>
</alpino_ds>
```

Output (reformatted):

```
# source = knmi.1.xml
# sent_id = knmi.1
# text = In het Waddengebied is de wind veranderlijk.
# auto = ALUD 1
1   In              in              ADP     VZ|init                     _                                       3   case    3:case      _
2   het             het             DET     LID|bep|stan|evon           Definite=Def                            3   det     3:det       _
3   Waddengebied    Waddengebied    PROPN   N|eigen|ev|basis|onz|stan   Gender=Neut|Number=Sing                 7   obl     7:obl:in    _
4   is              zijn            AUX     WW|pv|tgw|ev                Number=Sing|Tense=Pres|VerbForm=Fin     7   cop     7:cop       _
5   de              de              DET     LID|bep|stan|rest           Definite=Def                            6   det     6:det       _
6   wind            wind            NOUN    N|soort|ev|basis|zijd|stan  Gender=Com|Number=Sing                  7   nsubj   7:nsubj     _
7   veranderlijk    veranderlijk    ADJ     ADJ|vrij|basis|zonder       Degree=Pos                              0   root    0:root      SpaceAfter=No
8   .               .               PUNCT   LET                         _                                       7   punct   7:punct     _

```


For visualisation of the result, see http://www.let.rug.nl/kleiweg/conllu/

## Other functionality

 * Derive Universal Dependencies and insert into the alpino_ds format
 * Insert given Universal Dependencies into the alpino_ds format

## Example programs

The package comes with two example programs, `alud` and `alud-dact`.

To install the demo program `alud-dact`, you need to have
[Oracle Berkeley DB XML](https://www.oracle.com/database/berkeley-db/xml.html)
installed. With `alud-dact` you can process [dact](https://rug-compling.github.io/dact/) files.

Install both:

```
make bin
```

... or install `alud` only:

```
make alud
```

... or install `alud-dact` only:

```
make alud-dact
```

You will find the programs in one of these directories:

 1. in `$GOBIN` if variable is set
 2. else in `$GOPATH/bin` if variable is set
 3. else in `~/go/bin`
