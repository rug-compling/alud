//go:build ignore
// +build ignore

package alud

// recursive
func dependencyLabel(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "dependencyLabel", q, node))
		}
	}()

	depthCheck(q)

	if node.parent.Cat == "top" && node.parent.End == 1000 {
		return "root"
	}

	if node.Rel == "app" {
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return "appos"
		}
		if TEST(q, `$node/../node[@rel="mod" and (@pt or @cat)]`) {
			return "appos" //was "orphan"
		}
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "cmp" {
		return "mark"
	}
	if node.Rel == "crd" {
		if node.Lemma == "zowel" && TEST(q, `$node/../node[@rel="crd" and @lemma="als"]`) {
			return "cc:preconj"
		}
		if node.Lemma == "niet" && TEST(q, `$node/../node[@rel="crd" and @lemma="maar"]`) {
			return "cc:preconj"
		}
		if node.Lemma == "respectievelijk" && TEST(q, `$node/../node[@rel="crd" and @lemma="en"]`) {
			return "cc:preconj"
		}
		if node.Lemma == "ver" && TEST(q, `$node/../node[@rel="crd" and @lemma="en"]`) { // verder A, B, en C
			return "cc:preconj"
		}
		if node.Lemma == "achtereenvolgens" && TEST(q, `$node/../node[@rel="crd" and @lemma="en"]`) { // [PP door [NP achtereenvolgens A en B]]
			return "cc:preconj"
		}
		if node.Lemma == "van" && TEST(q, `$node/../node[@rel="crd" and @lemma="tot"]`) {
			return "cc:preconj"
		}
		if node.Lemma == "noch" && TEST(q, `$node/../node[@rel="crd" and @lemma="noch" and $node/@begin < @begin ]`) {
			return "cc:preconj"
		}
		if node.Lemma == "eerder" && TEST(q, `$node/../node[@rel="crd" and @lemma="dan" and $node/@begin < @begin ]`) {
			return "cc:preconj"
		}
		return "cc"
	}
	if node.Rel == "me" && TEST(q, `$node[@rel="me" and ../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="ADP")]]`) {
		return determineNominalModLabel(node, q)
	}
	if node.Rel == "obcomp" {
		return "advcl"
	}
	if node.Rel == "obj2" { // added case for coordination of aan-PP -- GB 2/3/23
		if node.Cat == "pp" || TEST(q, `$node[@cat="conj" and node[@cat="pp"]/node[@lemma="aan"]]`) {
			return "obl"
		}
		return "iobj"
	}
	if node.Rel == "pobj1" {
		return "expl"
	}
	if node.Rel == "predc" {
		if TEST(q, `$node/../node[@rel=("obj1","se","vc") and (@pt or @cat)] or $node/../node[@rel="hd" and (@pt or @cat) and not(@ud:pos="AUX")]`) {
			if TEST(q, `$node/../@cat="pp"`) { // check for absolutive (met) constructions, https://github.com/UniversalDependencies/docs/issues/408
				if TEST(q, `$node/../../@cat="np"`) {
					return "acl"
				}
				if TEST(q, `$node/../../@cat=("top","du")`) { // met als gevolg een neiging tot...  added du GB 9/3/21
					return dependencyLabel(node.parent, q) //  replaced "root" with this to account for dp-dp cases, where met-PP is second element (hoe doe je dat met sokken aan) GB 18/03/21
				}
				if TEST(q, `$node[../../../@cat=("top","du") and ../@rel="body"]`) { // [cmp alleen dan] [body met een langer 75 mm kanon ], in automatic output GB 22/03/21
					return dependencyLabel(node.parent, q) //
				}
				return "advcl"
			}
			return "xcomp"
		}
		if TEST(q, `$node[../node[@rel="su" and (@pt or @cat)] and ../node[@rel="hd" and not(@pt or @cat)] and ../@rel="cnj"]`) { // [su _ predc] cases, here we assume predc is the hd and predc an orphan GB 8/1/24
			if TEST(q, `$node/../node[@rel="hd" and 
		                               ( @ud:pos="AUX" or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index )]`) {
				return dependencyLabel(node.parent, q)
			}
			return "orphan" // but return orphan if the missing verb is not AUX (ie X werd koning en Y koningin)
		}
		return dependencyLabel(node.parent, q) // covers gapping cases where predc is promoted to head as well (not anymore, see above GB 8/1/24)
		/*
		   hack for now: de keuze is gauw gemaakt
		   was amod, is this more accurate??
		   examples of secondary predicates under xcomp suggests so

		*/
	}
	if node.Rel == "se" {
		return "expl:pv"
	}
	if node.Rel == "su" { // predc can be the head, even if its own hd is empty -- GB 21/2/24
		if TEST(q, `$node[../@rel=("cnj","dp","body","nucl") and 
		                       ../node[@rel="hd" and 
		                               ( @ud:pos="AUX" or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index ) ]
										and ../node[@rel="predc" and (@pt or @cat)] 
						]`) { // gapping with predc
			return subjectLabel(node, q)
		}
		if TEST(q, `$node[../@rel=("cnj","dp","body","nucl") and 
		                  ../node[@rel="hd" and not(@pt or @cat)] and  
		                  not(../node[@rel=("vc","XXXpredc") and (@pt or node[@rel=("hd","cnj","vc")  and (@pt or @cat)] )] )]`) { // gapping, added vc for recursive cases GB 26/2/24
			return dependencyLabel(node.parent, q) // added nucl GB 4/3/21
		}
		if TEST(q, `$node[../@rel="vc" and ../node[@rel="hd" and not(@pt or @cat)]
	                                 and ../parent::node[@rel="cnj"]]`) { // gapping with subj downstairs
			// TODO: ../.. is veranderd in ../parent::node , is dat juist? JA
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
	if aux, err := auxiliary1(node, q); err == nil {
		if aux == "aux:pass" {
			if TEST(q, `$node[../node[@rel="su" and not(@pt or @cat)] and
	                 ../node[@rel="vc" and not(descendant-or-self::node/@pt)] and
	                 ../@rel="cnj"]`) {
				return "conj"
			}
			return "aux:pass"
		}
		if aux == "aux" {
			if TEST(q, `$node[../node[@rel="su" and not(@pt or @cat)] and
	                 ../node[@rel="vc" and not(descendant-or-self::node/@pt)] and
	                 ../@rel="cnj"]`) {
				return "conj"
			}
			return "aux"
		}
		if aux == "cop" {
			if TEST(q, `$node[../node[@rel="su" and not(@pt or @cat)] and
	                 ../node[@rel="predc" and not(descendant-or-self::node/@pt)] and
	                 ../@rel="cnj"]`) {
				return "conj"
			}
			return "cop"
		}
	}
	if node.Rel == "det" {
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return detLabel(node, q)
		}
		if TEST(q, `$node/../node[@rel="mod" and (@pt or @cat)]`) { // gapping
			return detLabel(node, q) // was "orphan"
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
		if TEST(q, `$node[@index = ../../node[@rel="su"]/@index and ../node[@rel="hd" and (@pt or @cat)]]`) {
			return "nsubj:pass" // trees where su (with extraposed material) is spelled out at position of obj1
		} // but not elliptic cases (with empty head) like: Boven de engel is de profeet Zacharia afgebeeld en boven Maria de profeet Micha
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
		if TEST(q, `$node[@pt and ../node/@index = ancestor::node[@cat="whq"]/node[@rel="whd"]/@index]`) { // wat voor constructions, where wat is dislocated
			return "fixed"
		}
		if node == nLeft(FIND(q, `$node/../node[@rel="mwp" and (@pt or @cat)]`)) {
			return dependencyLabel(node.parent, q)
		}
		if TEST(q, `$node/../node[@ud:pos="PROPN"]`) {
			return "flat"
		}
		return "fixed" // v2 mwe-> fixed
	}
	if node.Rel == "cnj" {
		if TEST(q, `$node[@cat="pp"]/node[@rel="obj1" and @index and (@pt or @cat)]`) { // binnen en buiten de kerk --> kerk becomes obl or nmod GB 4/3/24
			return dependencyLabel(node.parent, q)
		}
		if node == nLeft(FIND(q, `$node/../node[@rel="cnj"]`)) { //changed picking first cnj in XML to nLeft test GB 4/3/21
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
		if TEST(q, `$node/../node[@rel="hd" and 
			                      ( @ud:pos=("AUX","ADP")  or $node/ancestor::node[@rel="top"]//node[@ud:pos="AUX"]/@index = @index ) 
			                      ] and
	                not($node/../node[@rel="predc"])`) {
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
	if node.Rel == "mod" && node.parent.Rel == "det" { // [ap/det steeds vnw/meer] invloed
		return "amod"
	}
	if (node.Rel == "mod" || node.Rel == "pc" || node.Rel == "ld") && node.parent.Cat == "np" { // [detp niet veel] meer
		// modification of nomimal heads
		// pc and ld occur in nominalizations
		if TEST(q, `$node/../node[@rel="hd" and (@pt or @cat)]`) {
			return modLabelInsideNp(node, q)
		}
		if node == nLeft(FIND(q, `$node/../node[@rel=("mod") and (@pt or @cat)]`)) { // gapping with multiple mods
			return dependencyLabel(node.parent, q) // gapping, where this mod is the head
		}
		return modLabelInsideNp(node, q) //was "orphan"
	}

	if TEST(q, `$node[@rel=("pc","ld") and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp","cp","whrel")]`) { // give pc priority GB 8/2/21
		// modification of verbal, adjectival heads
		// nb some oti's directly dominate (preceding) modifiers
		// [advp weg ermee ]
		if TEST(q, `$node/../node[@rel=("hd","body") and (@pt or @cat)]`) { // body for mods dangling outside cmp/body: maar niet om ...
			return labelVmod(node, q)
		}
		if TEST(q, `$node/../node[@rel=("su","obj1","predc","vc") and (@pt or @cat)]`) {
			return "orphan"
		}
		return dependencyLabel(node.parent, q) // gapping, where this mod is the head
	}

	if TEST(q, `$node[@rel="mod" and ../@cat=("sv1","smain","ssub","inf","ppres","ppart","oti","ap","advp","cp","whrel")]`) {
		// modification of verbal, adjectival heads
		// nb some oti's directly dominate (preceding) modifiers
		// [advp weg ermee ]
		if TEST(q, `$node/../node[@rel=("hd","body") and (@pt or @cat)]`) { // body for mods dangling outside cmp/body: maar niet om ...
			return labelVmod(node, q)
		}
		if TEST(q, `$node[not(../node[@rel="hd"]) and ../node[@rel="predc" and (@pt or @cat)]]`) { // mod in context of predc but no empty head
			// avoid orphan, because that would not disappear in enhanced
			return labelVmod(node, q)
		}
		if TEST(q, `$node/../node[@rel=("su","obj1","predc","vc","pc") and (@pt or @cat)]`) { // added pc 16/2/21
			return "orphan"
		}
		if TEST(q, `$node[../@rel="body" and ../@cat="ssub"]`) { // wat cool is en wat niet GB 24/1/24
			return "orphan"
		}
		// combined mod/ld/pc for consistency with dephead position & superfluous again, see above. Gb 8/2 BN begin-position checks dont work as these are renumbered internally
		if TEST(q, `$node[@rel=("mod","ld","pc") and ../node[@rel=("mod","pc","ld") and (@pt or @cat)]/@begin < @begin ]`) { // gapping with multiple mods
			return "orphan" //  added (@pt or @cat) to prevent match with an empty mod, eventhough those do not have a @begin attribute
		} // seems like a BUG
		return dependencyLabel(node.parent, q) // gapping, where this mod is the head
	}
	if TEST(q, `$node[@rel="mod" and ../@cat=("pp","part")]`) { // [mod hd/ADP obj1/empty]  --> make mod the external head , added part for automatic parse output GB 31/03/21
		if TEST(q, `$node/../node[@rel="obj1" and (@pt or @cat)]`) {
			return "amod"
		}
		if TEST(q, `$node/../node[@rel="hd" and @ud:pos=("ADV","ADP")]`) { // daarom dus, vlak voor en tijdens de oorlog --> orphan or advmod?
			return "advmod"
		}
		return dependencyLabel(node.parent, q)
	}
	if TEST(q, `$node[@rel="mod" and ../@cat=("detp","advp")]`) {
		return "amod"
	}

	if TEST(q, `$node[@rel="mod" and ../@cat="conj"]`) { // occurs in automatic parse output, should be more finegrained if it happens more often GB 22/03/21
		return "nmod"
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
		if node.udPos == "VERB" || node.udPos == "ADJ" {
			return "advcl"
		}
		if node.udPos == "PRON" || node.udPos == "NOUN" { //floating quantifiers, zich politiek bezoedelen
			return "obl" // , https://github.com/UniversalDependencies/docs/issues/581, https://github.com/UniversalDependencies/UD_French-Sequoia/issues/6,
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
		if node.udPos == "PRON" { // [whd wat] [body nu]
			return "obl"
		}
		return "advmod" // [whq waarom jij]
	}
	if node.Rel == "body" {
		return dependencyLabel(node.parent, q)
	}
	if node.Rel == "--" {
		// debugging, removed last escape for punct/num/sym from this position
		// commented out, as seems to be covered by clause below and gives errors in some cases
		// restricted to PUNCT now...26/03/21
		if TEST(q, `$node[@cat="mwu" and node/@ud:pos=("PUNCT","SYM")]`) { // fix for mwu punctuation in Alpino output
			if TEST(q, `$node/../node[not(@ud:pos="PUNCT" or (@cat="mwu" and node/@ud:pos=("PUNCT","SYM")))]`) {
				return "punct"
			}
			return "root"
		}
		if node.udPos == "PUNCT" {
			if TEST(q, `$node[not(../node[not(@ud:pos="PUNCT")]) and @begin = ../@begin]`) {
				return "root" // just punctuation
			}
			return "punct"
		}
		if node.udPos == "X" { // SYM covered below
			if TEST(q, `$node/../node[@cat]`) {
				return "parataxis" // 1. Jantje is ziek  1-->appos??
			}
			// return "root"
		}
		if node.Lemma == "\\" {
			return "punct" // hack for tagging errors in lassy small 250
		}
		if TEST(q, `$node[@ud:pos="NUM" and ../node[@cat] ]`) {
			return "parataxis" // dangling number 1.
		}
		if TEST(q, `$node[@ud:pos="CCONJ" and ../node[@cat="smain" or @cat="conj"]]`) {
			return "cc"
		}
		// sentence initial or final 'en' etc , merge with statement below??
		if TEST(q, `$node[@ud:pos=("NOUN","PROPN","DET","ADP","ADV","INTJ","PRON","SYM","CCONJ","SCONJ") and ../node[(@pt="spec" and @rel="dp") or @cat=("du","smain","conj","sv1","np","whq","pp","ppart","inf","advp")]]`) {
			return "parataxis" // dangling words, generalize to all POS ? GB 26/03/21
		}

		if TEST(q, `$node[@cat="mwu" or @ud:pos=("ADP","ADV","DET","PRON","CCONJ","SCONJ","NOUN","PROPN","INTJ","NUM","SYM")]`) { // make exception for du nodes as well GB 22/03/21
			if node == nLeft(FIND(q, `$node/../node[(@cat="mwu" or @ud:pos=("ADP","ADV","DET","PRON","CCONJ","SCONJ","NOUN","PROPN","INTJ","NUM","SYM")) 
			                                    and not(../node[@ud:pos=("ADJ","VERB") or @cat=("du","ap","smain","ppart","ti")]) ]`)) {
				return "root"
			}
			return "parataxis"
		}
		if TEST(q, `$node[@ud:pos=("ADJ","VERB")]`) {
			if node == nLeft(FIND(q, `$node/../node[@ud:pos=("ADJ","VERB")]`)) {
				return "root"
			}
			return "parataxis"
		}
		if TEST(q, `$node[not(@ud:pos)]/../@rel="top"`) {
			if TEST(q, `$node/../node[@ud:pos=("VERB","ADJ")]`) {
				return "parataxis"
			}
			return "root"
		}

		//debugging, moved this to last escape option fixing [MISSING PARA] error messages ...

		if TEST(q, `count($node/../node[not(@ud:pos=("PUNCT","SYM","X"))]) < 2`) {
			return "root" // only one non-punct/sym/foreign element in the string
		}
		panic("No label --")
	}

	if node.Rel == "hd" {
		if node.udPos == "ADP" || node.parent.Cat == "pp" { // added pp check for PPs headed by a verb (gezien de structuur....) GB 23/03/21
			if TEST(q, `$node/../node[@rel="predc"]`) {
				return "mark" // absolute met constructie -- analoog aan with X being Y
			}
			if TEST(q, `$node[../node[@rel="obj1" and @index]]`) {
				if TEST(q, `$node[../@rel="cnj" and ../node[@rel="obj1" and not(@pt or @cat)]]`) { // binnen en buiten de kerk, binnen is case of kerk GB 4/3/24
					return "case"
				}
				if TEST(q, `$node[../@rel="cnj" and ../node[@rel="obj1" and (@pt or @cat)] and ../../node[@rel="cnj" and @cat="ap"]]`) { // op weg naar en in het Heilige Land  -- op weg = ap GB 12/3/24
					return "case"
				}
				if TEST(q, `$node[../@rel="cnj"]`) { // binnen en buiten de kerk, buiten is conj of binnen GB 4/3/24
					return "conj"
				}
				return "case" // waar-relatives with P-stranding: waar hij Gabbema voor bedankt
			}
			if TEST(q, `$node/../node[@rel=("obj1","vc","se","me") and (@pt or @cat)]`) { // examples in paqus suggest me case is already covered (advmod), yet leuven/253 gives error without me here..
				return "case" //
			}
			if TEST(q, `$node/../node[@rel="pc"]`) { // superfluous ??
				return dependencyLabel(node.parent, q) // er blijft weinig over van het lijk : over heads a predc and has pc as sister
			}
			if TEST(q, `$node[../node[@rel="mod" and (@pt or @cat)] and ../@cat="pp"]`) { //
				return dependencyLabel(node.parent, q) // [mod om wat te zonnen] in [1] en bij [1 de kleine meertjes]  , changed from parataxis, GB 3/3/21 (consistent with treatment in depheads)
				// actually, this case is redundant now
			}
			return dependencyLabel(node.parent, q) // [predc [mod het beste] af/hd,ADP] here af heads a predc --> go to parent
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
	panic("No label")
}

func determineNominalModLabel(node *nodeType, q *context) string {
	if TEST(q, `$node/../node[@rel="hd" and (@ud:pos="VERB" or @ud:pos="ADJ")]`) {
		return "obl"
	}
	return "nmod"
}

func subjectLabel(subj *nodeType, q *context) string {
	// pass (requires pass aux) and outer (requires copula) never co-occur I assume
	pass := passiveSubject(subj, q)
	outer := outerSubject(subj, q)
	if TEST(q, `$subj[@cat=("whsub","ssub","ti","cp","oti")] or
	            $subj[@cat="conj" and node/@cat=("whsub","ssub","ti","cp","oti")]`) {
		return "csubj" + pass + outer
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
	if TEST(q, `$subj[../@rel=("cnj","dp","body","nucl") and ../node[@rel="hd" and not(@pt or @cat)] and ../node[@rel="predc" and (@pt or @cat)] ]`) {
		return "nsubj" + pass + outer //  predc is the head if it has any content GB 21/2/24
	}
	// do we ever reach the case below, if this is already checked for before calling this function (ie at rel=su level)
	if TEST(q, `$subj[../@rel=("cnj","dp","body","nucl") and ../node[@rel="hd" and not(@pt or @cat)] and 
	                  not(../node[@rel=("vc","predc") and (@pt or node[@rel=("hd","cnj","vc")  and (@pt or @cat)] )   ] )]`) { // gapping, added vc for recursive cases GB 26/2/24
		return dependencyLabel(subj.parent, q) // copied from regular rel=su case GB 24/1/24 but now the predc is an orphan to su, and we should always go to parent??? GB 5/2/24
	}
	return "nsubj" + pass + outer
}

// add :outer to subject in copula construction with clausal predicate (whrel, cp (dat +finite clause), other?)
func outerSubject(subj *nodeType, q *context) string {
	if TEST(q, `$subj/../node[@rel="predc" and (@cat="whrel" or (@cat="conj" and node[@cat="whrel"]))]`) {
		return ":outer"
	}
	if TEST(q, `$subj/../node[@rel="predc" and @cat="cp"]/node[@cat="ssub" or (@cat="conj" and node[@cat="ssub"])]`) {
		return ":outer"
	}
	if TEST(q, `$subj/../node[@rel="predc" and @cat="smain"]`) {
		return ":outer"
	}
	return ""
}

// recursive
func passiveSubject(subj *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "passiveSubject", q, nil, nil, nil, subj))
		}
	}()

	depthCheck(q)

	if aux, err := auxiliary1(n1(FIND(q, `($subj/../node[@rel="hd"])[1]`)), q); err == nil {
		if aux == "aux:pass" { // de carriere had gered kunnen worden
			return ":pass"
		}
		if aux == "aux" && TEST(q, `$subj/@index = $subj/../node[@rel="vc"]/node[@rel="su"]/@index`) {
			return passiveSubject(n1(FIND(q, `$subj/../node[@rel="vc"]/node[@rel="su"]`)), q)
		}
	}
	return ""
}

func detLabel(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "detLabel", q, node))
		}
	}()

	// zijn boek, diens boek, ieders boek, aller landen, Ron's probleem, Fidel Castro's belang, Ditvoorst carriere (ie spelling error?), zuids schoppenaas
	if TEST(q, `$node[@ud:pos = "PRON" and @vwtype="bez"] or
	          $node[@ud:pos = ("PRON", "NOUN") and @naamval="gen"] or
	          $node[@cat="mwu" and node[@spectype="deeleigen"]] or
	          $node[@ud:pos = "PROPN"]`) {
		return "nmod:poss"
	}
	if TEST(q, `$node/@ud:pos = ("DET","PRON","X")`) {
		return "det" // meer // genoeg // the
	}
	if node.Cat == "detp" {
		if TEST(q, `$node/node[@rel="hd" and @ud:pos="NUM"]`) {
			return "nummod"
		}
		if TEST(q, `$node/node[@rel="hd" and @ud:pos=("NOUN","ADJ")]`) {
			return "nmod"
		}
		if TEST(q, `$node/node[@rel="hd" and @ud:pos="PRON" and @vwtype="bez"]`) {
			return "nmod:poss"
		}
		return "det"
	}
	if TEST(q, `$node[@cat=("np","ap") or @ud:pos=("SYM","ADJ","ADV","NOUN") ]`) {
		return "nmod"
	}
	if TEST(q, `$node/@cat = ("mwu","smain")`) {
		return "det"
	}
	// tussen 5 en 6 ..--> almost all PP cases are with tussen NUM and NUM
	if TEST(q, `$node/@cat = "pp"`) {
		return "nummod"
	}
	// nog meer mensen dan anders
	// hetzelfde tijdstip als anders , nogal wat,
	// hij heeft ik weet niet hoeveel boeken
	if node.udPos == "NUM" || node.udPos == "SYM" {
		return "nummod"
	}
	if node.Cat == "conj" {

		return detLabel(n1(FIND(q, `($node/node[@rel="cnj"])[1]`)), q)
	}
	if node.Cat == "cp" { //ik heb boeken gezien [cp/det dan hem] weird...
		return "nmod"
	}
	panic("No label det")
}

func modLabelInsideNp(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "modLabelInsideNp", q, node))
		}
	}()

	if TEST(q, `$node[@cat="pp"]/node[@rel="vc"]`) {
		return "acl" // pp containing a clause
	}
	// fixing issue noted by Anouck
	if TEST(q, `$node[@ud:pos="ADJ" or @cat="ap" or node[@cat="conj" and @begin = node[@ud:pos="ADJ" or @cat="ap"]/@begin ]]`) {
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
	if node.Cat == "rel" || node.Cat == "whrel" || node.Cat == "ssub" {
		return "acl:relcl"
	}
	// v2 added relcl -- whrel= met name waar ...
	if TEST(q, `$node[@cat="cp"]/node[@rel="body" and (@ud:pos = ("NOUN","PROPN") or @cat=("np","conj"))]`) {
		return "nmod"
	}
	// zijn loopbaan [CP als schrijver]
	if TEST(q, `$node[@cat=("cp","sv1","smain","ppres","ppart","inf","ti","oti","du","whq", "whsub") or @ud:pos="SCONJ"]`) { // added whsub for robustness in automatic parser output
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
		panic("Index nmod")
	}
	panic("No label nmod")
}

func labelVmod(node *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "labelVmod", q, node))
		}
	}()

	if TEST(q, `$node[@cat="pp"]/node[@rel="vc"]`) {
		return "advcl"
	}
	if TEST(q, `$node[ (   node[@rel="hd" and @lemma="door"]
		                or node[@rel="cnj"]/node[@rel="hd" and @lemma="door"]
	                    or (@pt="bw" and ends-with(@lemma,"door"))
	                   )
	                 ]`) {
		if aux, err := auxiliary1(n1(FIND(q, `$node/../../node[@rel="hd" and @lemma=("zijn","worden")]`)), q); err == nil && aux == "aux:pass" {
			return "obl:agent"
		}
	}
	/*
		but NOT: door rookontwikkeling om het leven gekomen
		-- already filtered by constraint of su/obj1 control or @sc=passive (cheating, for impersonal passives) GB 20/04/23
		NOT: bij Bakema is een stoeptegel door de ruit gegooid
		NO/YES: hierdoor werd Prince door het grote publiek ontdekt
		ADDED: coordination cases GB 11/04/23
	*/

	if TEST(q, `$node[@cat=("pp","conj") and @rel="pc"]`) { // added GB 20/11/23
		return "obl:arg" // what about enhanced deps (obl:arg:aan ? yes, see CZ treebanks)
	}
	if TEST(q, `$node[@cat=("pp","np","conj","mwu") or @ud:pos=("NOUN","VERB","PRON","PROPN","X","PUNCT","SYM","ADP") ]`) {
		return "obl"
	}
	if TEST(q, `$node[@cat="advp"]/node[@ud:pos=("NOUN","VERB","ADP")]`) {
		return "obl"
	}
	if TEST(q, `$node[@cat="ap"]/node[@ud:pos="NOUN" or @cat="np"]`) { // added NP for 'het hele jaar door' h_suite/53 GB 26/02/21
		return "obl"
	}
	if TEST(q, `$node[@cat=("cp","sv1","smain","ssub","ppres","ppart","ti","oti","inf","du","whq","whrel","rel","whsub")]`) {
		return "advcl"
	}
	if TEST(q, `$node[@ud:pos= ("ADJ","ADV","SCONJ","INTJ")
	                      or @cat=("advp","ap")
	                      or (@cat="conj" and node/@ud:pos="ADV")]`) {
		return "advmod" // niet of nauwelijks
	}
	if node.udPos == "NUM" {
		return "nummod"
	}
	if node.Index > 0 {
		panic("Index vmod")
	}
	panic("No label vmod")
}

// this function is now also used to distribute dependents in coordination in Enhanced UD , so lot more rels and contexts are possible
// and passives, as in " hun taal werd gediscrimineerd en verboden"
func nonLocalDependencyLabel(head, gap *nodeType, q *context) string {

	defer func() {
		if r := recover(); r != nil {
			panic(trace(r, "nonLocalDependencyLabel", q, nil, head, gap))
		}
	}()

	if gap.Rel == "su" {
		return subjectLabel(gap, q)
	}
	if gap.Rel == "obj1" {
		if head.Rel == "su" {
			return "nsubj:pass"
		}
		if TEST(q, `$gap/../@cat="pp"`) { // waar men de patienten [ i naar toe] schuift om hen vervolgens te vergeten . ✤
			return "obl"
		}
		return "obj" // ppart coordination -- see comment above
	}
	if gap.Rel == "obj2" {
		if head.udPos == "ADV" {
			return "advmod"
		}
		if head.udPos == "ADP" { // fixed for questions 'aan wie werd het gegeven' -- GB 2/3/23
			return "obl"
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
		if TEST(q, `$head/node[@rel="obj1"]`) { // is dit onzin? de head kan een pp met een np inside zijn, maar dat zegt niets over nmod, dit moet gewoon obl zijn GB 17/11/23
			return "nmod"
		}
		if TEST(q, `$head[@ud:pos=("ADV", "ADP") or @cat=("advp","ap")]`) {
			return "advmod" // waar precies zit je ..
		}
		if TEST(q, `$head[@ud:pos="PRON"]`) { // het voorstel, [rhd_i dat] de commissie [ld_i] introk (from automatic parses)
			return "obl"
		}
		panic("No label index PC")
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
		if TEST(q, `$head[@ud:pos="ADJ"]`) { // added GB 9/3/21
			return "amod" // zeker_i vier a i vijf miljard
		}
		panic("No label index mod")
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
	panic("No label index")
}
