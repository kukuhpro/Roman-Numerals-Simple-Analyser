package client

var (
	KEY_MATERIAL    = "material"
	KEY_UNIT        = "unit"
	KEY_QUANTIFIERS = "quantifiers"
	KEY_QUESTION    = "question"
	KEY_VERB        = "verb"
	KEY_NUMERIC     = "numeric"
	KEY_ROMAN       = "roman"
	KEY_CREDIT      = "credit"
	KEY_MARK        = "mark"

	regexPatternSourceWord = map[string]string{
		KEY_MATERIAL:    "(Silver|Gold|Iron|silver|gold|iron)",
		KEY_UNIT:        "(glob|prok|pish|tegj)",
		KEY_QUANTIFIERS: "(Many|Much|many|much)",
		KEY_QUESTION:    "(How|how)",
		KEY_VERB:        "^(Is|is)$",
		KEY_NUMERIC:     "[0-9]+",
		KEY_ROMAN:       "^(I|V|X|L|C|D|M)$",
		KEY_CREDIT:      "(Credits|credits|Credit|credit)",
		KEY_MARK:        `(\?)`,
	}
)
