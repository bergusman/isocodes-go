package iso15924

type Script struct {
	Alpha4  string
	Numeric string
	Name    string
}

var (
	Adlm = Script{Alpha4: "Adlm", Numeric: "166", Name: "Adlam"}
	Afak = Script{Alpha4: "Afak", Numeric: "439", Name: "Afaka"}
	Aghb = Script{Alpha4: "Aghb", Numeric: "239", Name: "Caucasian Albanian"}
	Ahom = Script{Alpha4: "Ahom", Numeric: "338", Name: "Ahom, Tai Ahom"}
	Arab = Script{Alpha4: "Arab", Numeric: "160", Name: "Arabic"}
	Aran = Script{Alpha4: "Aran", Numeric: "161", Name: "Arabic (Nastaliq variant)"}
	Armi = Script{Alpha4: "Armi", Numeric: "124", Name: "Imperial Aramaic"}
	Armn = Script{Alpha4: "Armn", Numeric: "230", Name: "Armenian"}
	Avst = Script{Alpha4: "Avst", Numeric: "134", Name: "Avestan"}
	Bali = Script{Alpha4: "Bali", Numeric: "360", Name: "Balinese"}
	Bamu = Script{Alpha4: "Bamu", Numeric: "435", Name: "Bamum"}
	Bass = Script{Alpha4: "Bass", Numeric: "259", Name: "Bassa Vah"}
	Batk = Script{Alpha4: "Batk", Numeric: "365", Name: "Batak"}
	Beng = Script{Alpha4: "Beng", Numeric: "325", Name: "Bengali (Bangla)"}
	Bhks = Script{Alpha4: "Bhks", Numeric: "334", Name: "Bhaiksuki"}
	Blis = Script{Alpha4: "Blis", Numeric: "550", Name: "Blissymbols"}
	Bopo = Script{Alpha4: "Bopo", Numeric: "285", Name: "Bopomofo"}
	Brah = Script{Alpha4: "Brah", Numeric: "300", Name: "Brahmi"}
	Brai = Script{Alpha4: "Brai", Numeric: "570", Name: "Braille"}
	Bugi = Script{Alpha4: "Bugi", Numeric: "367", Name: "Buginese"}
	Buhd = Script{Alpha4: "Buhd", Numeric: "372", Name: "Buhid"}
	Cakm = Script{Alpha4: "Cakm", Numeric: "349", Name: "Chakma"}
	Cans = Script{Alpha4: "Cans", Numeric: "440", Name: "Unified Canadian Aboriginal Syllabics"}
	Cari = Script{Alpha4: "Cari", Numeric: "201", Name: "Carian"}
	Cham = Script{Alpha4: "Cham", Numeric: "358", Name: "Cham"}
	Cher = Script{Alpha4: "Cher", Numeric: "445", Name: "Cherokee"}
	Chrs = Script{Alpha4: "Chrs", Numeric: "109", Name: "Chorasmian"}
	Cirt = Script{Alpha4: "Cirt", Numeric: "291", Name: "Cirth"}
	Copt = Script{Alpha4: "Copt", Numeric: "204", Name: "Coptic"}
	Cpmn = Script{Alpha4: "Cpmn", Numeric: "402", Name: "Cypro-Minoan"}
	Cprt = Script{Alpha4: "Cprt", Numeric: "403", Name: "Cypriot syllabary"}
	Cyrl = Script{Alpha4: "Cyrl", Numeric: "220", Name: "Cyrillic"}
	Cyrs = Script{Alpha4: "Cyrs", Numeric: "221", Name: "Cyrillic (Old Church Slavonic variant)"}
	Deva = Script{Alpha4: "Deva", Numeric: "315", Name: "Devanagari (Nagari)"}
	Diak = Script{Alpha4: "Diak", Numeric: "342", Name: "Dives Akuru"}
	Dogr = Script{Alpha4: "Dogr", Numeric: "328", Name: "Dogra"}
	Dsrt = Script{Alpha4: "Dsrt", Numeric: "250", Name: "Deseret (Mormon)"}
	Dupl = Script{Alpha4: "Dupl", Numeric: "755", Name: "Duployan shorthand, Duployan stenography"}
	Egyd = Script{Alpha4: "Egyd", Numeric: "070", Name: "Egyptian demotic"}
	Egyh = Script{Alpha4: "Egyh", Numeric: "060", Name: "Egyptian hieratic"}
	Egyp = Script{Alpha4: "Egyp", Numeric: "050", Name: "Egyptian hieroglyphs"}
	Elba = Script{Alpha4: "Elba", Numeric: "226", Name: "Elbasan"}
	Elym = Script{Alpha4: "Elym", Numeric: "128", Name: "Elymaic"}
	Ethi = Script{Alpha4: "Ethi", Numeric: "430", Name: "Ethiopic (Geʻez)"}
	Geok = Script{Alpha4: "Geok", Numeric: "241", Name: "Khutsuri (Asomtavruli and Nuskhuri)"}
	Geor = Script{Alpha4: "Geor", Numeric: "240", Name: "Georgian (Mkhedruli and Mtavruli)"}
	Glag = Script{Alpha4: "Glag", Numeric: "225", Name: "Glagolitic"}
	Gong = Script{Alpha4: "Gong", Numeric: "312", Name: "Gunjala Gondi"}
	Gonm = Script{Alpha4: "Gonm", Numeric: "313", Name: "Masaram Gondi"}
	Goth = Script{Alpha4: "Goth", Numeric: "206", Name: "Gothic"}
	Gran = Script{Alpha4: "Gran", Numeric: "343", Name: "Grantha"}
	Grek = Script{Alpha4: "Grek", Numeric: "200", Name: "Greek"}
	Gujr = Script{Alpha4: "Gujr", Numeric: "320", Name: "Gujarati"}
	Guru = Script{Alpha4: "Guru", Numeric: "310", Name: "Gurmukhi"}
	Hanb = Script{Alpha4: "Hanb", Numeric: "503", Name: "Han with Bopomofo (alias for Han + Bopomofo)"}
	Hang = Script{Alpha4: "Hang", Numeric: "286", Name: "Hangul (Hangŭl, Hangeul)"}
	Hani = Script{Alpha4: "Hani", Numeric: "500", Name: "Han (Hanzi, Kanji, Hanja)"}
	Hano = Script{Alpha4: "Hano", Numeric: "371", Name: "Hanunoo (Hanunóo)"}
	Hans = Script{Alpha4: "Hans", Numeric: "501", Name: "Han (Simplified variant)"}
	Hant = Script{Alpha4: "Hant", Numeric: "502", Name: "Han (Traditional variant)"}
	Hatr = Script{Alpha4: "Hatr", Numeric: "127", Name: "Hatran"}
	Hebr = Script{Alpha4: "Hebr", Numeric: "125", Name: "Hebrew"}
	Hira = Script{Alpha4: "Hira", Numeric: "410", Name: "Hiragana"}
	Hluw = Script{Alpha4: "Hluw", Numeric: "080", Name: "Anatolian Hieroglyphs (Luwian Hieroglyphs, Hittite Hieroglyphs)"}
	Hmng = Script{Alpha4: "Hmng", Numeric: "450", Name: "Pahawh Hmong"}
	Hmnp = Script{Alpha4: "Hmnp", Numeric: "451", Name: "Nyiakeng Puachue Hmong"}
	Hrkt = Script{Alpha4: "Hrkt", Numeric: "412", Name: "Japanese syllabaries (alias for Hiragana + Katakana)"}
	Hung = Script{Alpha4: "Hung", Numeric: "176", Name: "Old Hungarian (Hungarian Runic)"}
	Inds = Script{Alpha4: "Inds", Numeric: "610", Name: "Indus (Harappan)"}
	Ital = Script{Alpha4: "Ital", Numeric: "210", Name: "Old Italic (Etruscan, Oscan, etc.)"}
	Jamo = Script{Alpha4: "Jamo", Numeric: "284", Name: "Jamo (alias for Jamo subset of Hangul)"}
	Java = Script{Alpha4: "Java", Numeric: "361", Name: "Javanese"}
	Jpan = Script{Alpha4: "Jpan", Numeric: "413", Name: "Japanese (alias for Han + Hiragana + Katakana)"}
	Jurc = Script{Alpha4: "Jurc", Numeric: "510", Name: "Jurchen"}
	Kali = Script{Alpha4: "Kali", Numeric: "357", Name: "Kayah Li"}
	Kana = Script{Alpha4: "Kana", Numeric: "411", Name: "Katakana"}
	Khar = Script{Alpha4: "Khar", Numeric: "305", Name: "Kharoshthi"}
	Khmr = Script{Alpha4: "Khmr", Numeric: "355", Name: "Khmer"}
	Khoj = Script{Alpha4: "Khoj", Numeric: "322", Name: "Khojki"}
	Kitl = Script{Alpha4: "Kitl", Numeric: "505", Name: "Khitan large script"}
	Kits = Script{Alpha4: "Kits", Numeric: "288", Name: "Khitan small script"}
	Knda = Script{Alpha4: "Knda", Numeric: "345", Name: "Kannada"}
	Kore = Script{Alpha4: "Kore", Numeric: "287", Name: "Korean (alias for Hangul + Han)"}
	Kpel = Script{Alpha4: "Kpel", Numeric: "436", Name: "Kpelle"}
	Kthi = Script{Alpha4: "Kthi", Numeric: "317", Name: "Kaithi"}
	Lana = Script{Alpha4: "Lana", Numeric: "351", Name: "Tai Tham (Lanna)"}
	Laoo = Script{Alpha4: "Laoo", Numeric: "356", Name: "Lao"}
	Latf = Script{Alpha4: "Latf", Numeric: "217", Name: "Latin (Fraktur variant)"}
	Latg = Script{Alpha4: "Latg", Numeric: "216", Name: "Latin (Gaelic variant)"}
	Latn = Script{Alpha4: "Latn", Numeric: "215", Name: "Latin"}
	Leke = Script{Alpha4: "Leke", Numeric: "364", Name: "Leke"}
	Lepc = Script{Alpha4: "Lepc", Numeric: "335", Name: "Lepcha (Róng)"}
	Limb = Script{Alpha4: "Limb", Numeric: "336", Name: "Limbu"}
	Lina = Script{Alpha4: "Lina", Numeric: "400", Name: "Linear A"}
	Linb = Script{Alpha4: "Linb", Numeric: "401", Name: "Linear B"}
	Lisu = Script{Alpha4: "Lisu", Numeric: "399", Name: "Lisu (Fraser)"}
	Loma = Script{Alpha4: "Loma", Numeric: "437", Name: "Loma"}
	Lyci = Script{Alpha4: "Lyci", Numeric: "202", Name: "Lycian"}
	Lydi = Script{Alpha4: "Lydi", Numeric: "116", Name: "Lydian"}
	Mahj = Script{Alpha4: "Mahj", Numeric: "314", Name: "Mahajani"}
	Maka = Script{Alpha4: "Maka", Numeric: "366", Name: "Makasar"}
	Mand = Script{Alpha4: "Mand", Numeric: "140", Name: "Mandaic, Mandaean"}
	Mani = Script{Alpha4: "Mani", Numeric: "139", Name: "Manichaean"}
	Marc = Script{Alpha4: "Marc", Numeric: "332", Name: "Marchen"}
	Maya = Script{Alpha4: "Maya", Numeric: "090", Name: "Mayan hieroglyphs"}
	Medf = Script{Alpha4: "Medf", Numeric: "265", Name: "Medefaidrin (Oberi Okaime, Oberi Ɔkaimɛ)"}
	Mend = Script{Alpha4: "Mend", Numeric: "438", Name: "Mende Kikakui"}
	Merc = Script{Alpha4: "Merc", Numeric: "101", Name: "Meroitic Cursive"}
	Mero = Script{Alpha4: "Mero", Numeric: "100", Name: "Meroitic Hieroglyphs"}
	Mlym = Script{Alpha4: "Mlym", Numeric: "347", Name: "Malayalam"}
	Modi = Script{Alpha4: "Modi", Numeric: "324", Name: "Modi, Moḍī"}
	Mong = Script{Alpha4: "Mong", Numeric: "145", Name: "Mongolian"}
	Moon = Script{Alpha4: "Moon", Numeric: "218", Name: "Moon (Moon code, Moon script, Moon type)"}
	Mroo = Script{Alpha4: "Mroo", Numeric: "264", Name: "Mro, Mru"}
	Mtei = Script{Alpha4: "Mtei", Numeric: "337", Name: "Meitei Mayek (Meithei, Meetei)"}
	Mult = Script{Alpha4: "Mult", Numeric: "323", Name: "Multani"}
	Mymr = Script{Alpha4: "Mymr", Numeric: "350", Name: "Myanmar (Burmese)"}
	Nand = Script{Alpha4: "Nand", Numeric: "311", Name: "Nandinagari"}
	Narb = Script{Alpha4: "Narb", Numeric: "106", Name: "Old North Arabian (Ancient North Arabian)"}
	Nbat = Script{Alpha4: "Nbat", Numeric: "159", Name: "Nabataean"}
	Newa = Script{Alpha4: "Newa", Numeric: "333", Name: "Newa, Newar, Newari, Nepāla lipi"}
	Nkdb = Script{Alpha4: "Nkdb", Numeric: "085", Name: "Naxi Dongba (na²¹ɕi³³ to³³ba²¹, Nakhi Tomba)"}
	Nkgb = Script{Alpha4: "Nkgb", Numeric: "420", Name: "Naxi Geba (na²¹ɕi³³ gʌ²¹ba²¹, 'Na-'Khi ²Ggŏ-¹baw, Nakhi Geba)"}
	Nkoo = Script{Alpha4: "Nkoo", Numeric: "165", Name: "N’Ko"}
	Nshu = Script{Alpha4: "Nshu", Numeric: "499", Name: "Nüshu"}
	Ogam = Script{Alpha4: "Ogam", Numeric: "212", Name: "Ogham"}
	Olck = Script{Alpha4: "Olck", Numeric: "261", Name: "Ol Chiki (Ol Cemet’, Ol, Santali)"}
	Orkh = Script{Alpha4: "Orkh", Numeric: "175", Name: "Old Turkic, Orkhon Runic"}
	Orya = Script{Alpha4: "Orya", Numeric: "327", Name: "Oriya (Odia)"}
	Osge = Script{Alpha4: "Osge", Numeric: "219", Name: "Osage"}
	Osma = Script{Alpha4: "Osma", Numeric: "260", Name: "Osmanya"}
	Ougr = Script{Alpha4: "Ougr", Numeric: "143", Name: "Old Uyghur"}
	Palm = Script{Alpha4: "Palm", Numeric: "126", Name: "Palmyrene"}
	Pauc = Script{Alpha4: "Pauc", Numeric: "263", Name: "Pau Cin Hau"}
	Pcun = Script{Alpha4: "Pcun", Numeric: "015", Name: "Proto-Cuneiform"}
	Pelm = Script{Alpha4: "Pelm", Numeric: "016", Name: "Proto-Elamite"}
	Perm = Script{Alpha4: "Perm", Numeric: "227", Name: "Old Permic"}
	Phag = Script{Alpha4: "Phag", Numeric: "331", Name: "Phags-pa"}
	Phli = Script{Alpha4: "Phli", Numeric: "131", Name: "Inscriptional Pahlavi"}
	Phlp = Script{Alpha4: "Phlp", Numeric: "132", Name: "Psalter Pahlavi"}
	Phlv = Script{Alpha4: "Phlv", Numeric: "133", Name: "Book Pahlavi"}
	Phnx = Script{Alpha4: "Phnx", Numeric: "115", Name: "Phoenician"}
	Plrd = Script{Alpha4: "Plrd", Numeric: "282", Name: "Miao (Pollard)"}
	Piqd = Script{Alpha4: "Piqd", Numeric: "293", Name: "Klingon (KLI pIqaD)"}
	Prti = Script{Alpha4: "Prti", Numeric: "130", Name: "Inscriptional Parthian"}
	Psin = Script{Alpha4: "Psin", Numeric: "103", Name: "Proto-Sinaitic"}
	Qaaa = Script{Alpha4: "Qaaa", Numeric: "900", Name: "Reserved for private use (start)"}
	Qabx = Script{Alpha4: "Qabx", Numeric: "949", Name: "Reserved for private use (end)"}
	Ranj = Script{Alpha4: "Ranj", Numeric: "303", Name: "Ranjana"}
	Rjng = Script{Alpha4: "Rjng", Numeric: "363", Name: "Rejang (Redjang, Kaganga)"}
	Rohg = Script{Alpha4: "Rohg", Numeric: "167", Name: "Hanifi Rohingya"}
	Roro = Script{Alpha4: "Roro", Numeric: "620", Name: "Rongorongo"}
	Runr = Script{Alpha4: "Runr", Numeric: "211", Name: "Runic"}
	Samr = Script{Alpha4: "Samr", Numeric: "123", Name: "Samaritan"}
	Sara = Script{Alpha4: "Sara", Numeric: "292", Name: "Sarati"}
	Sarb = Script{Alpha4: "Sarb", Numeric: "105", Name: "Old South Arabian"}
	Saur = Script{Alpha4: "Saur", Numeric: "344", Name: "Saurashtra"}
	Sgnw = Script{Alpha4: "Sgnw", Numeric: "095", Name: "SignWriting"}
	Shaw = Script{Alpha4: "Shaw", Numeric: "281", Name: "Shavian (Shaw)"}
	Shrd = Script{Alpha4: "Shrd", Numeric: "319", Name: "Sharada, Śāradā"}
	Shui = Script{Alpha4: "Shui", Numeric: "530", Name: "Shuishu"}
	Sidd = Script{Alpha4: "Sidd", Numeric: "302", Name: "Siddham, Siddhaṃ, Siddhamātṛkā"}
	Sind = Script{Alpha4: "Sind", Numeric: "318", Name: "Khudawadi, Sindhi"}
	Sinh = Script{Alpha4: "Sinh", Numeric: "348", Name: "Sinhala"}
	Sogd = Script{Alpha4: "Sogd", Numeric: "141", Name: "Sogdian"}
	Sogo = Script{Alpha4: "Sogo", Numeric: "142", Name: "Old Sogdian"}
	Sora = Script{Alpha4: "Sora", Numeric: "398", Name: "Sora Sompeng"}
	Soyo = Script{Alpha4: "Soyo", Numeric: "329", Name: "Soyombo"}
	Sund = Script{Alpha4: "Sund", Numeric: "362", Name: "Sundanese"}
	Sylo = Script{Alpha4: "Sylo", Numeric: "316", Name: "Syloti Nagri"}
	Syrc = Script{Alpha4: "Syrc", Numeric: "135", Name: "Syriac"}
	Syre = Script{Alpha4: "Syre", Numeric: "138", Name: "Syriac (Estrangelo variant)"}
	Syrj = Script{Alpha4: "Syrj", Numeric: "137", Name: "Syriac (Western variant)"}
	Syrn = Script{Alpha4: "Syrn", Numeric: "136", Name: "Syriac (Eastern variant)"}
	Tagb = Script{Alpha4: "Tagb", Numeric: "373", Name: "Tagbanwa"}
	Takr = Script{Alpha4: "Takr", Numeric: "321", Name: "Takri, Ṭākrī, Ṭāṅkrī"}
	Tale = Script{Alpha4: "Tale", Numeric: "353", Name: "Tai Le"}
	Talu = Script{Alpha4: "Talu", Numeric: "354", Name: "New Tai Lue"}
	Taml = Script{Alpha4: "Taml", Numeric: "346", Name: "Tamil"}
	Tang = Script{Alpha4: "Tang", Numeric: "520", Name: "Tangut"}
	Tavt = Script{Alpha4: "Tavt", Numeric: "359", Name: "Tai Viet"}
	Telu = Script{Alpha4: "Telu", Numeric: "340", Name: "Telugu"}
	Teng = Script{Alpha4: "Teng", Numeric: "290", Name: "Tengwar"}
	Tfng = Script{Alpha4: "Tfng", Numeric: "120", Name: "Tifinagh (Berber)"}
	Tglg = Script{Alpha4: "Tglg", Numeric: "370", Name: "Tagalog (Baybayin, Alibata)"}
	Thaa = Script{Alpha4: "Thaa", Numeric: "170", Name: "Thaana"}
	Thai = Script{Alpha4: "Thai", Numeric: "352", Name: "Thai"}
	Tibt = Script{Alpha4: "Tibt", Numeric: "330", Name: "Tibetan"}
	Tirh = Script{Alpha4: "Tirh", Numeric: "326", Name: "Tirhuta"}
	Tnsa = Script{Alpha4: "Tnsa", Numeric: "275", Name: "Tangsa"}
	Toto = Script{Alpha4: "Toto", Numeric: "294", Name: "Toto"}
	Ugar = Script{Alpha4: "Ugar", Numeric: "040", Name: "Ugaritic"}
	Vaii = Script{Alpha4: "Vaii", Numeric: "470", Name: "Vai"}
	Visp = Script{Alpha4: "Visp", Numeric: "280", Name: "Visible Speech"}
	Vith = Script{Alpha4: "Vith", Numeric: "228", Name: "Vithkuqi"}
	Wara = Script{Alpha4: "Wara", Numeric: "262", Name: "Warang Citi (Varang Kshiti)"}
	Wcho = Script{Alpha4: "Wcho", Numeric: "283", Name: "Wancho"}
	Wole = Script{Alpha4: "Wole", Numeric: "480", Name: "Woleai"}
	Xpeo = Script{Alpha4: "Xpeo", Numeric: "030", Name: "Old Persian"}
	Xsux = Script{Alpha4: "Xsux", Numeric: "020", Name: "Cuneiform, Sumero-Akkadian"}
	Yezi = Script{Alpha4: "Yezi", Numeric: "192", Name: "Yezidi"}
	Yiii = Script{Alpha4: "Yiii", Numeric: "460", Name: "Yi"}
	Zanb = Script{Alpha4: "Zanb", Numeric: "339", Name: "Zanabazar Square (Zanabazarin Dörböljin Useg, Xewtee Dörböljin Bicig, Horizontal Square Script)"}
	Zinh = Script{Alpha4: "Zinh", Numeric: "994", Name: "Code for inherited script"}
	Zmth = Script{Alpha4: "Zmth", Numeric: "995", Name: "Mathematical notation"}
	Zsye = Script{Alpha4: "Zsye", Numeric: "993", Name: "Symbols (Emoji variant)"}
	Zsym = Script{Alpha4: "Zsym", Numeric: "996", Name: "Symbols"}
	Zxxx = Script{Alpha4: "Zxxx", Numeric: "997", Name: "Code for unwritten documents"}
	Zyyy = Script{Alpha4: "Zyyy", Numeric: "998", Name: "Code for undetermined script"}
	Zzzz = Script{Alpha4: "Zzzz", Numeric: "999", Name: "Code for uncoded script"}
)

var Scripts = []Script{
	Adlm,
	Afak,
	Aghb,
	Ahom,
	Arab,
	Aran,
	Armi,
	Armn,
	Avst,
	Bali,
	Bamu,
	Bass,
	Batk,
	Beng,
	Bhks,
	Blis,
	Bopo,
	Brah,
	Brai,
	Bugi,
	Buhd,
	Cakm,
	Cans,
	Cari,
	Cham,
	Cher,
	Chrs,
	Cirt,
	Copt,
	Cpmn,
	Cprt,
	Cyrl,
	Cyrs,
	Deva,
	Diak,
	Dogr,
	Dsrt,
	Dupl,
	Egyd,
	Egyh,
	Egyp,
	Elba,
	Elym,
	Ethi,
	Geok,
	Geor,
	Glag,
	Gong,
	Gonm,
	Goth,
	Gran,
	Grek,
	Gujr,
	Guru,
	Hanb,
	Hang,
	Hani,
	Hano,
	Hans,
	Hant,
	Hatr,
	Hebr,
	Hira,
	Hluw,
	Hmng,
	Hmnp,
	Hrkt,
	Hung,
	Inds,
	Ital,
	Jamo,
	Java,
	Jpan,
	Jurc,
	Kali,
	Kana,
	Khar,
	Khmr,
	Khoj,
	Kitl,
	Kits,
	Knda,
	Kore,
	Kpel,
	Kthi,
	Lana,
	Laoo,
	Latf,
	Latg,
	Latn,
	Leke,
	Lepc,
	Limb,
	Lina,
	Linb,
	Lisu,
	Loma,
	Lyci,
	Lydi,
	Mahj,
	Maka,
	Mand,
	Mani,
	Marc,
	Maya,
	Medf,
	Mend,
	Merc,
	Mero,
	Mlym,
	Modi,
	Mong,
	Moon,
	Mroo,
	Mtei,
	Mult,
	Mymr,
	Nand,
	Narb,
	Nbat,
	Newa,
	Nkdb,
	Nkgb,
	Nkoo,
	Nshu,
	Ogam,
	Olck,
	Orkh,
	Orya,
	Osge,
	Osma,
	Ougr,
	Palm,
	Pauc,
	Pcun,
	Pelm,
	Perm,
	Phag,
	Phli,
	Phlp,
	Phlv,
	Phnx,
	Plrd,
	Piqd,
	Prti,
	Psin,
	Qaaa,
	Qabx,
	Ranj,
	Rjng,
	Rohg,
	Roro,
	Runr,
	Samr,
	Sara,
	Sarb,
	Saur,
	Sgnw,
	Shaw,
	Shrd,
	Shui,
	Sidd,
	Sind,
	Sinh,
	Sogd,
	Sogo,
	Sora,
	Soyo,
	Sund,
	Sylo,
	Syrc,
	Syre,
	Syrj,
	Syrn,
	Tagb,
	Takr,
	Tale,
	Talu,
	Taml,
	Tang,
	Tavt,
	Telu,
	Teng,
	Tfng,
	Tglg,
	Thaa,
	Thai,
	Tibt,
	Tirh,
	Tnsa,
	Toto,
	Ugar,
	Vaii,
	Visp,
	Vith,
	Wara,
	Wcho,
	Wole,
	Xpeo,
	Xsux,
	Yezi,
	Yiii,
	Zanb,
	Zinh,
	Zmth,
	Zsye,
	Zsym,
	Zxxx,
	Zyyy,
	Zzzz,
}

var ScriptsByAlpha4 = map[string]Script{
	"Adlm": Adlm,
	"Afak": Afak,
	"Aghb": Aghb,
	"Ahom": Ahom,
	"Arab": Arab,
	"Aran": Aran,
	"Armi": Armi,
	"Armn": Armn,
	"Avst": Avst,
	"Bali": Bali,
	"Bamu": Bamu,
	"Bass": Bass,
	"Batk": Batk,
	"Beng": Beng,
	"Bhks": Bhks,
	"Blis": Blis,
	"Bopo": Bopo,
	"Brah": Brah,
	"Brai": Brai,
	"Bugi": Bugi,
	"Buhd": Buhd,
	"Cakm": Cakm,
	"Cans": Cans,
	"Cari": Cari,
	"Cham": Cham,
	"Cher": Cher,
	"Chrs": Chrs,
	"Cirt": Cirt,
	"Copt": Copt,
	"Cpmn": Cpmn,
	"Cprt": Cprt,
	"Cyrl": Cyrl,
	"Cyrs": Cyrs,
	"Deva": Deva,
	"Diak": Diak,
	"Dogr": Dogr,
	"Dsrt": Dsrt,
	"Dupl": Dupl,
	"Egyd": Egyd,
	"Egyh": Egyh,
	"Egyp": Egyp,
	"Elba": Elba,
	"Elym": Elym,
	"Ethi": Ethi,
	"Geok": Geok,
	"Geor": Geor,
	"Glag": Glag,
	"Gong": Gong,
	"Gonm": Gonm,
	"Goth": Goth,
	"Gran": Gran,
	"Grek": Grek,
	"Gujr": Gujr,
	"Guru": Guru,
	"Hanb": Hanb,
	"Hang": Hang,
	"Hani": Hani,
	"Hano": Hano,
	"Hans": Hans,
	"Hant": Hant,
	"Hatr": Hatr,
	"Hebr": Hebr,
	"Hira": Hira,
	"Hluw": Hluw,
	"Hmng": Hmng,
	"Hmnp": Hmnp,
	"Hrkt": Hrkt,
	"Hung": Hung,
	"Inds": Inds,
	"Ital": Ital,
	"Jamo": Jamo,
	"Java": Java,
	"Jpan": Jpan,
	"Jurc": Jurc,
	"Kali": Kali,
	"Kana": Kana,
	"Khar": Khar,
	"Khmr": Khmr,
	"Khoj": Khoj,
	"Kitl": Kitl,
	"Kits": Kits,
	"Knda": Knda,
	"Kore": Kore,
	"Kpel": Kpel,
	"Kthi": Kthi,
	"Lana": Lana,
	"Laoo": Laoo,
	"Latf": Latf,
	"Latg": Latg,
	"Latn": Latn,
	"Leke": Leke,
	"Lepc": Lepc,
	"Limb": Limb,
	"Lina": Lina,
	"Linb": Linb,
	"Lisu": Lisu,
	"Loma": Loma,
	"Lyci": Lyci,
	"Lydi": Lydi,
	"Mahj": Mahj,
	"Maka": Maka,
	"Mand": Mand,
	"Mani": Mani,
	"Marc": Marc,
	"Maya": Maya,
	"Medf": Medf,
	"Mend": Mend,
	"Merc": Merc,
	"Mero": Mero,
	"Mlym": Mlym,
	"Modi": Modi,
	"Mong": Mong,
	"Moon": Moon,
	"Mroo": Mroo,
	"Mtei": Mtei,
	"Mult": Mult,
	"Mymr": Mymr,
	"Nand": Nand,
	"Narb": Narb,
	"Nbat": Nbat,
	"Newa": Newa,
	"Nkdb": Nkdb,
	"Nkgb": Nkgb,
	"Nkoo": Nkoo,
	"Nshu": Nshu,
	"Ogam": Ogam,
	"Olck": Olck,
	"Orkh": Orkh,
	"Orya": Orya,
	"Osge": Osge,
	"Osma": Osma,
	"Ougr": Ougr,
	"Palm": Palm,
	"Pauc": Pauc,
	"Pcun": Pcun,
	"Pelm": Pelm,
	"Perm": Perm,
	"Phag": Phag,
	"Phli": Phli,
	"Phlp": Phlp,
	"Phlv": Phlv,
	"Phnx": Phnx,
	"Plrd": Plrd,
	"Piqd": Piqd,
	"Prti": Prti,
	"Psin": Psin,
	"Qaaa": Qaaa,
	"Qabx": Qabx,
	"Ranj": Ranj,
	"Rjng": Rjng,
	"Rohg": Rohg,
	"Roro": Roro,
	"Runr": Runr,
	"Samr": Samr,
	"Sara": Sara,
	"Sarb": Sarb,
	"Saur": Saur,
	"Sgnw": Sgnw,
	"Shaw": Shaw,
	"Shrd": Shrd,
	"Shui": Shui,
	"Sidd": Sidd,
	"Sind": Sind,
	"Sinh": Sinh,
	"Sogd": Sogd,
	"Sogo": Sogo,
	"Sora": Sora,
	"Soyo": Soyo,
	"Sund": Sund,
	"Sylo": Sylo,
	"Syrc": Syrc,
	"Syre": Syre,
	"Syrj": Syrj,
	"Syrn": Syrn,
	"Tagb": Tagb,
	"Takr": Takr,
	"Tale": Tale,
	"Talu": Talu,
	"Taml": Taml,
	"Tang": Tang,
	"Tavt": Tavt,
	"Telu": Telu,
	"Teng": Teng,
	"Tfng": Tfng,
	"Tglg": Tglg,
	"Thaa": Thaa,
	"Thai": Thai,
	"Tibt": Tibt,
	"Tirh": Tirh,
	"Tnsa": Tnsa,
	"Toto": Toto,
	"Ugar": Ugar,
	"Vaii": Vaii,
	"Visp": Visp,
	"Vith": Vith,
	"Wara": Wara,
	"Wcho": Wcho,
	"Wole": Wole,
	"Xpeo": Xpeo,
	"Xsux": Xsux,
	"Yezi": Yezi,
	"Yiii": Yiii,
	"Zanb": Zanb,
	"Zinh": Zinh,
	"Zmth": Zmth,
	"Zsye": Zsye,
	"Zsym": Zsym,
	"Zxxx": Zxxx,
	"Zyyy": Zyyy,
	"Zzzz": Zzzz,
}

var ScriptsByNumeric = map[string]Script{
	"166": Adlm,
	"439": Afak,
	"239": Aghb,
	"338": Ahom,
	"160": Arab,
	"161": Aran,
	"124": Armi,
	"230": Armn,
	"134": Avst,
	"360": Bali,
	"435": Bamu,
	"259": Bass,
	"365": Batk,
	"325": Beng,
	"334": Bhks,
	"550": Blis,
	"285": Bopo,
	"300": Brah,
	"570": Brai,
	"367": Bugi,
	"372": Buhd,
	"349": Cakm,
	"440": Cans,
	"201": Cari,
	"358": Cham,
	"445": Cher,
	"109": Chrs,
	"291": Cirt,
	"204": Copt,
	"402": Cpmn,
	"403": Cprt,
	"220": Cyrl,
	"221": Cyrs,
	"315": Deva,
	"342": Diak,
	"328": Dogr,
	"250": Dsrt,
	"755": Dupl,
	"070": Egyd,
	"060": Egyh,
	"050": Egyp,
	"226": Elba,
	"128": Elym,
	"430": Ethi,
	"241": Geok,
	"240": Geor,
	"225": Glag,
	"312": Gong,
	"313": Gonm,
	"206": Goth,
	"343": Gran,
	"200": Grek,
	"320": Gujr,
	"310": Guru,
	"503": Hanb,
	"286": Hang,
	"500": Hani,
	"371": Hano,
	"501": Hans,
	"502": Hant,
	"127": Hatr,
	"125": Hebr,
	"410": Hira,
	"080": Hluw,
	"450": Hmng,
	"451": Hmnp,
	"412": Hrkt,
	"176": Hung,
	"610": Inds,
	"210": Ital,
	"284": Jamo,
	"361": Java,
	"413": Jpan,
	"510": Jurc,
	"357": Kali,
	"411": Kana,
	"305": Khar,
	"355": Khmr,
	"322": Khoj,
	"505": Kitl,
	"288": Kits,
	"345": Knda,
	"287": Kore,
	"436": Kpel,
	"317": Kthi,
	"351": Lana,
	"356": Laoo,
	"217": Latf,
	"216": Latg,
	"215": Latn,
	"364": Leke,
	"335": Lepc,
	"336": Limb,
	"400": Lina,
	"401": Linb,
	"399": Lisu,
	"437": Loma,
	"202": Lyci,
	"116": Lydi,
	"314": Mahj,
	"366": Maka,
	"140": Mand,
	"139": Mani,
	"332": Marc,
	"090": Maya,
	"265": Medf,
	"438": Mend,
	"101": Merc,
	"100": Mero,
	"347": Mlym,
	"324": Modi,
	"145": Mong,
	"218": Moon,
	"264": Mroo,
	"337": Mtei,
	"323": Mult,
	"350": Mymr,
	"311": Nand,
	"106": Narb,
	"159": Nbat,
	"333": Newa,
	"085": Nkdb,
	"420": Nkgb,
	"165": Nkoo,
	"499": Nshu,
	"212": Ogam,
	"261": Olck,
	"175": Orkh,
	"327": Orya,
	"219": Osge,
	"260": Osma,
	"143": Ougr,
	"126": Palm,
	"263": Pauc,
	"015": Pcun,
	"016": Pelm,
	"227": Perm,
	"331": Phag,
	"131": Phli,
	"132": Phlp,
	"133": Phlv,
	"115": Phnx,
	"282": Plrd,
	"293": Piqd,
	"130": Prti,
	"103": Psin,
	"900": Qaaa,
	"949": Qabx,
	"303": Ranj,
	"363": Rjng,
	"167": Rohg,
	"620": Roro,
	"211": Runr,
	"123": Samr,
	"292": Sara,
	"105": Sarb,
	"344": Saur,
	"095": Sgnw,
	"281": Shaw,
	"319": Shrd,
	"530": Shui,
	"302": Sidd,
	"318": Sind,
	"348": Sinh,
	"141": Sogd,
	"142": Sogo,
	"398": Sora,
	"329": Soyo,
	"362": Sund,
	"316": Sylo,
	"135": Syrc,
	"138": Syre,
	"137": Syrj,
	"136": Syrn,
	"373": Tagb,
	"321": Takr,
	"353": Tale,
	"354": Talu,
	"346": Taml,
	"520": Tang,
	"359": Tavt,
	"340": Telu,
	"290": Teng,
	"120": Tfng,
	"370": Tglg,
	"170": Thaa,
	"352": Thai,
	"330": Tibt,
	"326": Tirh,
	"275": Tnsa,
	"294": Toto,
	"040": Ugar,
	"470": Vaii,
	"280": Visp,
	"228": Vith,
	"262": Wara,
	"283": Wcho,
	"480": Wole,
	"030": Xpeo,
	"020": Xsux,
	"192": Yezi,
	"460": Yiii,
	"339": Zanb,
	"994": Zinh,
	"995": Zmth,
	"993": Zsye,
	"996": Zsym,
	"997": Zxxx,
	"998": Zyyy,
	"999": Zzzz,
}
