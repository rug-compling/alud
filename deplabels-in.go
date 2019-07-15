// +build ignore

package main

// recursive
func dependencyLabel(node *NodeType, q *Context) string {
	if depthCheck(q, "dependencyLabel") {
		return "ERROR_NO_LABEL"
	}

	if node.parent.Cat == "top" && node.parent.End == 1000 {
		return "root"
	}
	if node.Rel == "app" {
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return "appos"
		}
		if TEST(q, `$node/../node[@rel="mod" and (@pt or @cat)]`) {
			return "orphan"
		}
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "cmp" {
		return "mark"
	}
	if node.Rel == "crd" {
		return "cc"
	}
	if node.Rel == "me" && TEST(q, `$node[@rel="me" and not(../node[@ud:pos="ADP"]) ]`) {
		return determineNominalModLabel(node, q)
	}
	if node.Rel == "obcomp" {
		return "advcl"
	}
	if node.Rel == "obj2" {
		if node.Cat == "pp" {
			return "obl"
		}
		return "iobj"
	}
	if node.Rel == "pobj1" {
		return "expl"
	}
	if node.Rel == "predc" {
		if TEST(q, `$node/../node[@rel=("obj1","se") and (@pt or @cat)] or $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX")]`) {
			if TEST(q, `$node/../@cat="pp"`) { // check for absolutive (met) constructions, https://github.com/UniversalDependencies/docs/issues/408
				if TEST(q, `$node/../../@cat="np"`) {
					return "acl"
				}
				return "advcl"
			}
			return "xcomp"
		}
		return dependencyLabel(node.parent, q) // covers gapping cases where predc is promoted to head as well
		/*
		   hack for now: de keuze is gauw gemaakt
		   was amod, is this more accurate??
		   examples of secondary predicates under xcomp suggests so

		*/
	}
	if node.Rel == "se" {
		return "expl:pv"
	}
	if node.Rel == "su" {
		if TEST(q, `$node[../@rel="cnj" and ../node[@rel="hd" and not(@pt or @cat)]]`) { // gapping
			return dependencyLabel(node.parent, q)
		}
		if TEST(q, `$node[../@rel="vc" and ../node[@rel="hd" and not(@pt or @cat)]
	                                 and ../parent::node[@rel="cnj" and node[@rel="hd" and not(@pt or @cat)]]]`) { // gapping with subj downstairs
			// TODO: ../.. is veranderd in ../parent::node    is dat juist?
			/*
			   In 1909 werd de persoonlijke dienstplicht ingevoerd en in 1913 de algemene persoonlijke dienstplicht .
			   [ hd_i su_j vc [ hd_k [_j pers dienstplicht ]
			*/
			return dependencyLabel(node.parent.parent, q)
		}
		return subjectLabel(node, q)
	}
	if node.Rel == "sup" {
		return "expl"
	}
	if node.Rel == "svp" {
		return "compound:prt" // v2: added prt extension
	}
	aux := auxiliary1(node, q)
	if aux == "aux:pass" {
		if TEST(q, `$node[../node[@rel="su" and not(@pt or @cat)] and
	                 ../node[@rel="vc" and not(@pt or @cat)] and
	                 ../@rel="cnj"]`) {
			return "conj"
		}
		return "aux:pass"
	}
	if aux == "aux" {
		return "aux"
	}
	if aux == "cop" {
		return "cop"
	}
	if node.Rel == "det" {
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return detLabel(node, q)
		}
		if TEST(q, `$node/../node[@rel="mod" and (@pt or @cat)]`) { // gapping
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "obj1" || node.Rel == "me" {
		if TEST(q, `$node/../@cat="pp" or $node/../node[@rel="hd" and @ud:pos="ADP"]`) { // vol vertrouwen , heel de geschiedenis door (cat=ap!)
			if TEST(q, `$node/../node[@rel="predc"]`) { // absolutive met
				return "nsubj"
			}
			return dependencyLabel(node.parent, q)
		}
		if TEST(q, `$node[@index = ../../node[@rel="su"]/@index ]`) {
			return "nsubj:pass" // trees where su (with extraposed material) is spelled out at position of obj1
		}
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return "obj"
		}
		if TEST(q, `$node/../node[@rel="su" and (@pt or @cat)]`) {
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if node.Rel == "mwp" {
		if node.Begin >= 0 && node.Begin == node.parent.Begin {
			return dependencyLabel(node.parent, q)
		}
		if TEST(q, `$node/../node[@ud:pos="PROPN"]`) {
			return "flat"
		}
		return "fixed" // v2 mwe-> fixed
	}
	if node.Rel == "cnj" {
		if node == n1(FIND(q, `$node/../node[@rel="cnj"][1]`)) {
			return dependencyLabel(node.parent, q)
		}
		return "conj"
	}
	if node.Rel == "dp" {
		if node == nLeft(FIND(q, `$node/../node[@rel="dp"]`)) {
			return dependencyLabel(node.parent, q)
		}
		return "parataxis"
	}
	if node.Rel == "tag" || node.Rel == "sat" {
		return "parataxis"
	}
	if node.Rel == "dlink" {
		return "mark"
	}
	if node.Rel == "nucl" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "vc" {
		if TEST(q, `$node/../node[@rel="hd" and @ud:pos=("AUX","ADP")] and not($node/../node[@rel="predc"])`) {
			return dependencyLabel(node.parent, q)
		}
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			if TEST(q, `$node/node[@rel="su" and @index and not(@word or @cat)]
	                           (: or $node[@cat=("ti","oti")] :)
	                           or $node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)]
	                           or $node[@cat="oti"]/node[@cat="ti"]/node[@rel="body"]/node[@rel="su" and @index and not(@word or @cat)]`) {
				return "xcomp"
			}
			if TEST(q, `$node/../@cat="np"`) {
				return "acl" // v2: clausal dependents of nouns always acl
			}
			return "ccomp"
		}
		if TEST(q, `$node/../node[@rel=("su","obj1") and (@pt or @cat)]`) {
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping
	}
	if (node.Rel == "mod" || node.Rel == "pc" || node.Rel == "ld") && node.parent.Cat == "np" { // [detp niet veel] meer
		// modification of nomimal heads
		// pc and ld occur in nominalizations
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return modLabelInsideNp(node, q)
		}
		if node == nLeft(FIND(q, `$node/../node[@rel="mod" and (@pt or @cat)]`)) { // gapping with multiple mods
			return dependencyLabel(node.parent, q) // gapping, where this mod is the head
		}
		return "orphan"
	}
	if TEST(q, `$node[@rel=("mod","pc","ld") and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")]`) {
		// modification of verbal, adjectival heads
		// nb some oti's directly dominate (preceding) modifiers
		// [advp weg ermee ]
		if TEST(q, `$node/../node[@rel=("hd","body") and (@pt or @cat)]`) { // body for mods dangling outside cmp/body: maar niet om ...
			return labelVmod(node, q)
		}
		if TEST(q, `$node/../node[@rel=("su","obj1","predc","vc") and (@pt or @cat)]`) {
			return "orphan"
		}
		if TEST(q, `$node[@rel="mod" and ../node[@rel=("pc","ld")]]`) {
			return "orphan"
		}
		if TEST(q, `$node[@rel="mod" and ../node[@rel="mod"]/@begin < @begin]`) { // gapping with multiple mods
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping, where this mod is the head
	}
	if TEST(q, `$node[@rel="mod" and ../@cat=("pp","detp","advp")]`) {
		return "amod"
	}
	if TEST(q, `$node[@rel="mod" and ../@cat=("cp", "whrel", "whq", "whsub")]`) {
		// [cp net  [cmp als] [body de Belgische regering]], net wat het land nodig heeft
		if TEST(q, `$node/../node[@rel="body"]/node[@rel="hd" and @ud:pos="VERB"]`) {
			return "advmod"
		}
		return "amod"
	}
	if node.Rel == "pc" && node.parent.Cat == "pp" { //  [[hd,mwu aan het hoofd] [pc van X]]
		return "nmod"
	}
	if node.Rel == "hdf" {
		return "case"
	}
	if node.Rel == "predm" {
		if node.udPos == "VERB" {
			return "xcomp"
		}
		if node.udPos != "" {
			return "advmod"
		}
		return "advcl"
	}
	if node.Rel == "rhd" || node.Rel == "whd" {
		if TEST(q, `$node/../node[@rel="body"]//node/@index = $node/@index`) {
			return nonLocalDependencyLabel(
				node,
				n1(FIND(q, `($node/../node[@rel="body"]//node[@index = $node/@index])[1]`)),
				q)
		}
		if node.Cat == "pp" {
			return "nmod" // onder wie michael boogerd
		}
		return "advmod" // [whq waarom jij]
	}
	if node.Rel == "body" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "--" {
		if node.udPos == "PUNCT" {
			if TEST(q, `$node[not(../node[not(@ud:pos="PUNCT")]) and @begin = ../@begin]`) {
				return "root" // just punctuation
			}
			return "punct"
		}
		if node.udPos == "SYM" || node.udPos == "X" {
			if TEST(q, `$node/../node[@cat]`) {
				return "appos" // 1. Jantje is ziek  1-->appos??
			}
			return "root"
		}
		if node.Lemma == "\\" {
			return "punct" // hack for tagging errors in lassy small 250
		}
		/*
			if node.Spectype == "deeleigen" {
				return "punct" // hack for tagging errors in lassy small 250
			}
		*/
		if TEST(q, `$node[@ud:pos="NUM" and ../node[@cat] ]`) {
			return "parataxis" // dangling number 1.
		}
		if TEST(q, `$node[@ud:pos="CCONJ" and ../node[@cat="smain" or @cat="conj"]]`) {
			return "cc"
		}
		// sentence initial or final 'en'
		if TEST(q, `$node[@ud:pos=("NOUN","PROPN","VERB") and ../node[@cat=("du","smain")]]`) {
			return "parataxis" // dangling words
		}
		if TEST(q, `count($node/../node[not(@ud:pos=("PUNCT","SYM","X"))]) < 2`) {
			return "root" // only one non-punct/sym/foreign element in the string
		}
		if node.Cat == "mwu" {
			if node.Begin == node.parent.Begin && node.End == node.parent.End {
				return "root"
			}
			if TEST(q, `$node/node[@ud:pos=("PUNCT","SYM")]`) { // fix for mwu punctuation in Alpino output
				return "punct"
			}
			return "ERROR_NO_LABEL_--"
		}
		if TEST(q, `$node[not(@ud:pos)]/../@rel="top"`) {
			return "root"
		}
		if TEST(q, `$node[@ud:pos="PROPN" and not(../node[@cat]) ]`) {
			return "root" // Arthur .
		}
		if TEST(q, `$node[@ud:pos=("ADP","ADV","ADJ","DET","PRON","CCONJ","NOUN","VERB","INTJ")]`) {
			return "parataxis"
		}
		return "ERROR_NO_LABEL_--"
	}
	if node.Rel == "hd" {
		if node.udPos == "ADP" {
			if TEST(q, `$node/../node[@rel="predc"]`) {
				return "mark" // absolute met constructie -- analoog aan with X being Y
			}
			if TEST(q, `not($node/../node[@rel="pc"])`) {
				return "case" // er blijft weinig over van het lijk : over heads a predc and has pc as sister
			}
			return dependencyLabel(node.parent, q) // not sure about this one
		}
		if TEST(q, `$node[(@ud:pos=("ADJ","X","ADV") or @cat="mwu")
	                            and ../@cat="pp"
	                            and ../node[@rel=("obj1","se","vc")]]`) {
			if TEST(q, `$node[../@rel="cnj" and ../node[@rel="obj1" and @index and not(@pt or @cat)]]`) {
				return "conj"
			}
			return "case" // vol vertrouwen in, Ad U3... , op het gebied van
		}
		if TEST(q, `$node/../node[@rel="hd"]/@begin < $node/@begin`) {
			return "conj"
		}
		return dependencyLabel(node.parent, q)
	}
	return "ERROR_NO_LABEL"
}

func determineNominalModLabel(node *NodeType, q *Context) string {
	if TEST(q, `$node/../node[@rel="hd" and (@ud:pos="VERB" or @ud:pos="ADJ")]`) {
		return "obl"
	}
	return "nmod"
}

func subjectLabel(subj *NodeType, q *Context) string {
	pass := passiveSubject(subj, q)
	if TEST(q, `$subj[@cat=("whsub","ssub","ti","cp","oti")] or
	            $subj[@cat="conj" and node/@cat=("whsub","ssub","ti","cp","oti")]`) {
		return "csubj" + pass
	}
	// weather verbs and other expletive subject verbs
	if TEST(q, `$subj[@lemma="het"] and
	                   ( $subj/../node[@rel="hd" and
	                                   @lemma=("dooien", "gieten", "hagelen", "miezeren",
	                                           "misten", "motregenen", "onweren", "plenzen",
	                                           "regenen", "sneeuwen", "stormen", "stortregenen",
	                                           "ijzelen", "vriezen", "weerlichten", "winteren",
	                                           "zomeren") ] or
	                     $subj/../node[@rel="hd" and
	                                   @lemma="ontbreken" and
	                                   ../node[@rel="pc" and node[@rel="hd" and @lemma="aan"] ] ] or
	                     $subj[@index = ../node//node[@rel="sup"]/@index]    (: het kan voorkomen dat ... :)
	                   )`) {
		return "expl" + pass
	}
	return "nsubj" + pass
}

// recursive
func passiveSubject(subj *NodeType, q *Context) string {
	if depthCheck(q, "passiveSubject") {
		return "ERROR_NO_PASSIVE_SUBJECT"
	}

	aux := auxiliary1(n1(FIND(q, `($subj/../node[@rel="hd"])[1]`)), q)
	if aux == "aux:pass" { // de carriere had gered kunnen worden
		return ":pass"
	}
	if aux == "aux" && TEST(q, `$subj/@index = $subj/../node[@rel="vc"]/node[@rel="su"]/@index`) {
		return passiveSubject(n1(FIND(q, `$subj/../node[@rel="vc"]/node[@rel="su"]`)), q)
	}
	return ""
}

func detLabel(node *NodeType, q *Context) string {
	// zijn boek, diens boek, ieders boek, aller landen, Ron's probleem, Fidel Castro's belang
	if TEST(q, `$node[@ud:pos = "PRON" and @vwtype="bez"] or
	          $node[@ud:pos = ("PRON","PROPN") and @naamval="gen"] or
	          $node[@cat="mwu" and node[@spectype="deeleigen"]]`) {
		return "nmod:poss"
	}
	if TEST(q, `$node/@ud:pos = ("DET","PROPN","NOUN","ADJ","PRON","ADV","X")`) {
		return "det" // meer // genoeg // the
	}
	if TEST(q, `$node/@cat = ("mwu","np","pp","ap","detp","smain")`) {
		return "det"
	}
	// tussen 5 en 6 .., needs more principled solution
	// nog meer mensen dan anders
	// hetzelfde tijdstip als anders , nogal wat,
	// hij heeft ik weet niet hoeveel boeken
	if node.udPos == "NUM" || node.udPos == "SYM" {
		return "nummod"
	}
	if node.Cat == "conj" {
		// TODO: als ik hier 1 vervang door last() dan verdwijnen de verschillen met Saxon, maar het moet echt 1 zijn
		if TEST(q, `$node/node[@rel="cnj"][1]/@ud:pos="NUM"`) {
			return "nummod"
		}
		return "det"
	}
	return "ERROR_NO_LABEL_DET"
}

func modLabelInsideNp(node *NodeType, q *Context) string {
	if TEST(q, `$node[@cat="pp"]/node[@rel="vc"]`) {
		return "acl" // pp containing a clause
	}
	if TEST(q, `$node[@ud:pos="ADJ" or @cat="ap" or node[@cat="conj" and node[@ud:pos="ADJ" or @cat="ap"] ]]`) {
		return "amod"
	}
	if TEST(q, `$node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM","INTJ") ]`) {
		return "nmod"
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if TEST(q, `$node[@cat="detp"]/node[@rel="hd" and @ud:pos="NUM"]`) {
		return "nummod"
	}
	if node.Cat == "detp" {
		return "det" // [detp niet veel] meer error?
	}
	if node.Cat == "rel" || node.Cat == "whrel" {
		return "acl:relcl"
	}
	// v2 added relcl -- whrel= met name waar ...
	if TEST(q, `$node[@cat="cp"]/node[@rel="body" and (@ud:pos = ("NOUN","PROPN") or @cat=("np","conj"))]`) {
		return "nmod"
	}
	// zijn loopbaan [CP als schrijver]
	if TEST(q, `$node[@cat=("cp","sv1","smain","ppres","ppart","inf","ti","oti","du","whq") or @ud:pos="SCONJ"]`) {
		return "acl"
	}
	// oa zinnen tussen haakjes
	if TEST(q, `$node[@ud:pos= ("ADV","ADP","VERB","CCONJ") or @cat="advp"]`) {
		return "amod"
	}
	// VERB= aanstormend etc -> amod, ADV = nagenoeg alle prijzen, slechts 4 euro --> amod
	// CCONJ = opdrachten zoals:   --> amod
	if node.Rel == "det" {
		return "det" // empty determiners in gapping?
	}
	if node.Index > 0 {
		return "ERROR_INDEX_NMOD"
	}
	return "ERROR_NO_LABEL_NMOD"
}

func labelVmod(node *NodeType, q *Context) string {
	if TEST(q, `$node[@cat="pp"]/node[@rel="vc"]`) {
		return "advcl"
	}
	if TEST(q, `$node[ (  node[@rel="hd" and @lemma="door"]
	                                  or (@pt="bw" and ends-with(@lemma,"door"))
	                                  )
	                                  and parent::node[@cat="ppart"]/../node[@rel="hd" and @lemma=("zijn","worden")]
	                                  and ../../node[@rel="su"]/@index = ../node[@rel="obj1"]/@index
	                              ]`) {
		return "obl:agent"
		/*
			but NOT: door rookontwikkeling om het leven gekomen
			-- already filtered by constraint of su/obj1 control
			NOT: bij Bakema is een stoeptegel door de ruit gegooid
			NO/YES: hierdoor werd Prince door het grote publiek ontdekt
		*/
	}
	if TEST(q, `$node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","PRON","PROPN","X","PUNCT","SYM") ]`) {
		return "obl"
	}
	if TEST(q, `$node[@cat=("cp","sv1","smain","ssub","ppres","ppart","ti","oti","inf","du","whq","whrel","rel")]`) {
		return "advcl"
	}
	if TEST(q, `$node[@ud:pos= ("ADJ","ADV","ADP","VERB","SCONJ","INTJ")
	                      or @cat=("advp","ap")
	                      or (@cat="conj" and node/@ud:pos="ADV")]`) {
		return "advmod" // niet of nauwelijks
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if node.Index > 0 {
		return "ERROR_INDEX_VMOD"
	}
	return "ERROR_NO_LABEL_VMOD"
}

// this function is now also used to distribute dependents in coordination in Enhanced UD , so lot more rels and contexts are possible
// and passives, as in " hun taal werd gediscrimineerd en verboden"
func nonLocalDependencyLabel(head, gap *NodeType, q *Context) string {
	if gap.Rel == "su" {
		return subjectLabel(gap, q)
	}
	if gap.Rel == "obj1" {
		if head.Rel == "su" {
			return "nsubj:pass"
		}
		return "obj" // ppart coordination -- see comment above
	}
	if gap.Rel == "obj2" {
		if head.udPos == "ADV" {
			return "advmod"
		}
		return "iobj"
	}
	if gap.Rel == "me" {
		return determineNominalModLabel(gap, q)
	}
	if gap.Rel == "predc" || gap.Rel == "predm" {
		return dependencyLabel(gap, q)
	}
	if gap.Rel == "pc" || gap.Rel == "ld" {
		if TEST(q, `$head/node[@rel="obj1"]`) {
			return "nmod"
		}
		if TEST(q, `$head[@ud:pos=("ADV", "ADP") or @cat=("advp","ap")]`) {
			return "advmod" // waar precies zit je ..
		}
		return "ERROR_NO_LABEL_INDEX_PC"
	}
	if gap.Rel == "sup" || gap.Rel == "pobj1" {
		return "expl" // waar het om gaat is dat hij scoort, het is 1881 en dertien jaar geleden dat ...
	}
	if gap.Rel == "mwp" {
		return dependencyLabel(gap.parent, q) //wat heb je voor boeken gelezen
	}
	if gap.Rel == "vc" {
		return "ccomp" // wat ik me afvraag is of hij komt -- CLEFT
	}
	if TEST(q, `$gap[@rel="mod" and ../@cat=("np","pp")]`) { // voornamelijk in kloosters en door vrouwen
		return modLabelInsideNp(head, q)
	}
	if TEST(q, `$gap[@rel="mod" and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp")]`) {
		return labelVmod(head, q)
	}
	if gap.Rel == "mod" || gap.Rel == "spec" { // spec only used for funny coord
		if TEST(q, `$head[@cat=("pp","np") or @ud:pos=("NOUN","PRON")]`) {
			return "nmod"
		}
		if TEST(q, `$head[@ud:pos=("ADV","ADP") or @cat= ("advp","ap","mwu","conj")]`) {
			return "advmod" // hoe vaak -- AP, daar waar, waar en wanneer, voor als rhd
		}
		return "ERROR_NO_LABEL_INDEX_MOD"
	}
	if gap.Rel == "det" {
		return detLabel(head, q)
	}
	if TEST(q, `$gap[@rel="hd"] and $head[@ud:pos=("ADV","ADP")]`) { // waaronder A, B, en C
		return "case"
	}
	if gap.Rel == "du" || gap.Rel == "dp" {
		return "parataxis"
	}
	return "ERROR_NO_LABEL_INDEX"
}
