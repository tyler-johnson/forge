package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func ifaceToString(i interface{}) string {
	var str string
	f := toIfaceSlice(i)

	if f != nil {
		for _, v := range f {
			char := v.([]byte)
			str += string(char[:])
		}
	}

	return str
}

var g = &grammar{
	rules: []*rule{
		{
			name: "start",
			pos:  position{line: 26, col: 1, offset: 402},
			expr: &actionExpr{
				pos: position{line: 26, col: 9, offset: 410},
				run: (*parser).callonstart1,
				expr: &seqExpr{
					pos: position{line: 26, col: 9, offset: 410},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 26, col: 9, offset: 410},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 26, col: 11, offset: 412},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 26, col: 13, offset: 414},
								name: "Content",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 26, col: 21, offset: 422},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 30, col: 1, offset: 451},
			expr: &labeledExpr{
				pos:   position{line: 30, col: 11, offset: 461},
				label: "v",
				expr: &choiceExpr{
					pos: position{line: 31, col: 3, offset: 468},
					alternatives: []interface{}{
						&actionExpr{
							pos: position{line: 31, col: 3, offset: 468},
							run: (*parser).callonComment3,
							expr: &seqExpr{
								pos: position{line: 31, col: 3, offset: 468},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 31, col: 3, offset: 468},
										val:        "#",
										ignoreCase: false,
									},
									&labeledExpr{
										pos:   position{line: 31, col: 7, offset: 472},
										label: "v",
										expr: &zeroOrMoreExpr{
											pos: position{line: 31, col: 9, offset: 474},
											expr: &actionExpr{
												pos: position{line: 31, col: 10, offset: 475},
												run: (*parser).callonComment8,
												expr: &seqExpr{
													pos: position{line: 31, col: 10, offset: 475},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 31, col: 10, offset: 475},
															expr: &charClassMatcher{
																pos:        position{line: 31, col: 11, offset: 476},
																val:        "[\\r\\n]",
																chars:      []rune{'\r', '\n'},
																ignoreCase: false,
																inverted:   false,
															},
														},
														&labeledExpr{
															pos:   position{line: 31, col: 18, offset: 483},
															label: "ch",
															expr: &anyMatcher{
																line: 31, col: 21, offset: 486,
															},
														},
													},
												},
											},
										},
									},
									&ruleRefExpr{
										pos:  position{line: 31, col: 44, offset: 509},
										name: "NL",
									},
								},
							},
						},
						&actionExpr{
							pos: position{line: 32, col: 3, offset: 535},
							run: (*parser).callonComment15,
							expr: &seqExpr{
								pos: position{line: 32, col: 3, offset: 535},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 32, col: 3, offset: 535},
										val:        "//",
										ignoreCase: false,
									},
									&labeledExpr{
										pos:   position{line: 32, col: 8, offset: 540},
										label: "v",
										expr: &zeroOrMoreExpr{
											pos: position{line: 32, col: 10, offset: 542},
											expr: &actionExpr{
												pos: position{line: 32, col: 11, offset: 543},
												run: (*parser).callonComment20,
												expr: &seqExpr{
													pos: position{line: 32, col: 11, offset: 543},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 32, col: 11, offset: 543},
															expr: &charClassMatcher{
																pos:        position{line: 32, col: 12, offset: 544},
																val:        "[\\r\\n]",
																chars:      []rune{'\r', '\n'},
																ignoreCase: false,
																inverted:   false,
															},
														},
														&labeledExpr{
															pos:   position{line: 32, col: 19, offset: 551},
															label: "ch",
															expr: &anyMatcher{
																line: 32, col: 22, offset: 554,
															},
														},
													},
												},
											},
										},
									},
									&ruleRefExpr{
										pos:  position{line: 32, col: 45, offset: 577},
										name: "NL",
									},
								},
							},
						},
						&actionExpr{
							pos: position{line: 33, col: 3, offset: 603},
							run: (*parser).callonComment27,
							expr: &seqExpr{
								pos: position{line: 33, col: 3, offset: 603},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 33, col: 3, offset: 603},
										val:        "/*",
										ignoreCase: false,
									},
									&labeledExpr{
										pos:   position{line: 33, col: 8, offset: 608},
										label: "v",
										expr: &zeroOrMoreExpr{
											pos: position{line: 33, col: 10, offset: 610},
											expr: &actionExpr{
												pos: position{line: 33, col: 11, offset: 611},
												run: (*parser).callonComment32,
												expr: &seqExpr{
													pos: position{line: 33, col: 11, offset: 611},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 33, col: 11, offset: 611},
															expr: &litMatcher{
																pos:        position{line: 33, col: 12, offset: 612},
																val:        "*/",
																ignoreCase: false,
															},
														},
														&labeledExpr{
															pos:   position{line: 33, col: 17, offset: 617},
															label: "ch",
															expr: &anyMatcher{
																line: 33, col: 20, offset: 620,
															},
														},
													},
												},
											},
										},
									},
									&litMatcher{
										pos:        position{line: 33, col: 43, offset: 643},
										val:        "*/",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 33, col: 48, offset: 648},
										name: "NL",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Content",
			pos:  position{line: 36, col: 1, offset: 675},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 685},
				run: (*parser).callonContent1,
				expr: &labeledExpr{
					pos:   position{line: 36, col: 11, offset: 685},
					label: "b",
					expr: &zeroOrMoreExpr{
						pos: position{line: 36, col: 13, offset: 687},
						expr: &choiceExpr{
							pos: position{line: 36, col: 14, offset: 688},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 36, col: 14, offset: 688},
									name: "Comment",
								},
								&ruleRefExpr{
									pos:  position{line: 36, col: 24, offset: 698},
									name: "Verb",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Verb",
			pos:  position{line: 50, col: 1, offset: 921},
			expr: &actionExpr{
				pos: position{line: 50, col: 8, offset: 928},
				run: (*parser).callonVerb1,
				expr: &seqExpr{
					pos: position{line: 50, col: 8, offset: 928},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 50, col: 8, offset: 928},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 50, col: 13, offset: 933},
								name: "identity",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 50, col: 22, offset: 942},
							name: "SP",
						},
						&labeledExpr{
							pos:   position{line: 50, col: 25, offset: 945},
							label: "args",
							expr: &zeroOrMoreExpr{
								pos: position{line: 50, col: 30, offset: 950},
								expr: &ruleRefExpr{
									pos:  position{line: 50, col: 31, offset: 951},
									name: "Arg",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 50, col: 37, offset: 957},
							label: "body",
							expr: &zeroOrOneExpr{
								pos: position{line: 50, col: 42, offset: 962},
								expr: &ruleRefExpr{
									pos:  position{line: 50, col: 42, offset: 962},
									name: "Body",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 50, col: 48, offset: 968},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Arg",
			pos:  position{line: 62, col: 1, offset: 1128},
			expr: &actionExpr{
				pos: position{line: 62, col: 7, offset: 1134},
				run: (*parser).callonArg1,
				expr: &seqExpr{
					pos: position{line: 62, col: 7, offset: 1134},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 62, col: 7, offset: 1134},
							label: "k",
							expr: &choiceExpr{
								pos: position{line: 63, col: 3, offset: 1141},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 63, col: 3, offset: 1141},
										name: "null",
									},
									&ruleRefExpr{
										pos:  position{line: 64, col: 3, offset: 1151},
										name: "boolean",
									},
									&ruleRefExpr{
										pos:  position{line: 65, col: 3, offset: 1164},
										name: "number",
									},
									&ruleRefExpr{
										pos:  position{line: 66, col: 3, offset: 1176},
										name: "str",
									},
									&ruleRefExpr{
										pos:  position{line: 67, col: 3, offset: 1185},
										name: "Variable",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 68, col: 3, offset: 1198},
							name: "SP",
						},
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 70, col: 1, offset: 1222},
			expr: &actionExpr{
				pos: position{line: 70, col: 12, offset: 1233},
				run: (*parser).callonVariable1,
				expr: &seqExpr{
					pos: position{line: 70, col: 12, offset: 1233},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 70, col: 12, offset: 1233},
							label: "f",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 14, offset: 1235},
								name: "identity",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 1244},
							label: "r",
							expr: &zeroOrMoreExpr{
								pos: position{line: 70, col: 25, offset: 1246},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 26, offset: 1247},
									name: "CompoundPath",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 41, offset: 1262},
							label: "a",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 43, offset: 1264},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 43, offset: 1264},
									name: "MethodArgs",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "MethodArgs",
			pos:  position{line: 74, col: 1, offset: 1355},
			expr: &actionExpr{
				pos: position{line: 74, col: 14, offset: 1368},
				run: (*parser).callonMethodArgs1,
				expr: &seqExpr{
					pos: position{line: 74, col: 14, offset: 1368},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 74, col: 14, offset: 1368},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 18, offset: 1372},
							name: "SP",
						},
						&labeledExpr{
							pos:   position{line: 74, col: 21, offset: 1375},
							label: "args",
							expr: &zeroOrMoreExpr{
								pos: position{line: 74, col: 26, offset: 1380},
								expr: &ruleRefExpr{
									pos:  position{line: 74, col: 27, offset: 1381},
									name: "Arg",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 74, col: 33, offset: 1387},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "CompoundPath",
			pos:  position{line: 78, col: 1, offset: 1419},
			expr: &choiceExpr{
				pos: position{line: 79, col: 5, offset: 1437},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 79, col: 5, offset: 1437},
						run: (*parser).callonCompoundPath2,
						expr: &seqExpr{
							pos: position{line: 79, col: 5, offset: 1437},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 79, col: 5, offset: 1437},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 79, col: 9, offset: 1441},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 11, offset: 1443},
										name: "identity",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 80, col: 5, offset: 1475},
						run: (*parser).callonCompoundPath7,
						expr: &seqExpr{
							pos: position{line: 80, col: 5, offset: 1475},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 80, col: 5, offset: 1475},
									val:        "[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 80, col: 9, offset: 1479},
									label: "p",
									expr: &choiceExpr{
										pos: position{line: 80, col: 12, offset: 1482},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 80, col: 12, offset: 1482},
												name: "unsigned",
											},
											&ruleRefExpr{
												pos:  position{line: 80, col: 23, offset: 1493},
												name: "str",
											},
											&ruleRefExpr{
												pos:  position{line: 80, col: 29, offset: 1499},
												name: "Variable",
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 80, col: 39, offset: 1509},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Body",
			pos:  position{line: 82, col: 1, offset: 1534},
			expr: &actionExpr{
				pos: position{line: 82, col: 8, offset: 1541},
				run: (*parser).callonBody1,
				expr: &seqExpr{
					pos: position{line: 82, col: 8, offset: 1541},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 82, col: 8, offset: 1541},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 82, col: 12, offset: 1545},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 82, col: 14, offset: 1547},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 82, col: 16, offset: 1549},
								name: "Content",
							},
						},
						&litMatcher{
							pos:        position{line: 82, col: 24, offset: 1557},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "identity",
			pos:  position{line: 86, col: 1, offset: 1586},
			expr: &actionExpr{
				pos: position{line: 86, col: 12, offset: 1597},
				run: (*parser).callonidentity1,
				expr: &seqExpr{
					pos: position{line: 86, col: 12, offset: 1597},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 86, col: 12, offset: 1597},
							label: "f",
							expr: &charClassMatcher{
								pos:        position{line: 86, col: 14, offset: 1599},
								val:        "[a-z_$]i",
								chars:      []rune{'_', '$'},
								ranges:     []rune{'a', 'z'},
								ignoreCase: true,
								inverted:   false,
							},
						},
						&labeledExpr{
							pos:   position{line: 86, col: 23, offset: 1608},
							label: "r",
							expr: &zeroOrMoreExpr{
								pos: position{line: 86, col: 25, offset: 1610},
								expr: &charClassMatcher{
									pos:        position{line: 86, col: 25, offset: 1610},
									val:        "[a-z0-9_$]i",
									chars:      []rune{'_', '$'},
									ranges:     []rune{'a', 'z', '0', '9'},
									ignoreCase: true,
									inverted:   false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "boolean",
			pos:  position{line: 91, col: 1, offset: 1701},
			expr: &choiceExpr{
				pos: position{line: 92, col: 4, offset: 1713},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 92, col: 4, offset: 1713},
						run: (*parser).callonboolean2,
						expr: &litMatcher{
							pos:        position{line: 92, col: 4, offset: 1713},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 93, col: 4, offset: 1746},
						run: (*parser).callonboolean4,
						expr: &litMatcher{
							pos:        position{line: 93, col: 4, offset: 1746},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "number",
			pos:  position{line: 95, col: 1, offset: 1780},
			expr: &actionExpr{
				pos: position{line: 95, col: 10, offset: 1789},
				run: (*parser).callonnumber1,
				expr: &seqExpr{
					pos: position{line: 95, col: 10, offset: 1789},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 95, col: 10, offset: 1789},
							expr: &litMatcher{
								pos:        position{line: 95, col: 10, offset: 1789},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 95, col: 15, offset: 1794},
							expr: &charClassMatcher{
								pos:        position{line: 95, col: 15, offset: 1794},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 95, col: 22, offset: 1801},
							expr: &seqExpr{
								pos: position{line: 95, col: 23, offset: 1802},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 95, col: 23, offset: 1802},
										val:        ".",
										ignoreCase: false,
									},
									&oneOrMoreExpr{
										pos: position{line: 95, col: 27, offset: 1806},
										expr: &charClassMatcher{
											pos:        position{line: 95, col: 27, offset: 1806},
											val:        "[0-9]",
											ranges:     []rune{'0', '9'},
											ignoreCase: false,
											inverted:   false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "unsigned",
			pos:  position{line: 99, col: 1, offset: 1876},
			expr: &actionExpr{
				pos: position{line: 99, col: 12, offset: 1887},
				run: (*parser).callonunsigned1,
				expr: &oneOrMoreExpr{
					pos: position{line: 99, col: 12, offset: 1887},
					expr: &charClassMatcher{
						pos:        position{line: 99, col: 12, offset: 1887},
						val:        "[0-9]",
						ranges:     []rune{'0', '9'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "str",
			pos:  position{line: 103, col: 1, offset: 1958},
			expr: &actionExpr{
				pos: position{line: 103, col: 7, offset: 1964},
				run: (*parser).callonstr1,
				expr: &labeledExpr{
					pos:   position{line: 103, col: 7, offset: 1964},
					label: "v",
					expr: &choiceExpr{
						pos: position{line: 104, col: 2, offset: 1970},
						alternatives: []interface{}{
							&actionExpr{
								pos: position{line: 104, col: 2, offset: 1970},
								run: (*parser).callonstr4,
								expr: &seqExpr{
									pos: position{line: 104, col: 2, offset: 1970},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 104, col: 2, offset: 1970},
											val:        "\"",
											ignoreCase: false,
										},
										&labeledExpr{
											pos:   position{line: 104, col: 7, offset: 1975},
											label: "v",
											expr: &zeroOrMoreExpr{
												pos: position{line: 104, col: 9, offset: 1977},
												expr: &choiceExpr{
													pos: position{line: 104, col: 10, offset: 1978},
													alternatives: []interface{}{
														&ruleRefExpr{
															pos:  position{line: 104, col: 10, offset: 1978},
															name: "escape",
														},
														&charClassMatcher{
															pos:        position{line: 104, col: 19, offset: 1987},
															val:        "[^\"]",
															chars:      []rune{'"'},
															ignoreCase: false,
															inverted:   true,
														},
													},
												},
											},
										},
										&litMatcher{
											pos:        position{line: 104, col: 26, offset: 1994},
											val:        "\"",
											ignoreCase: false,
										},
									},
								},
							},
							&actionExpr{
								pos: position{line: 105, col: 2, offset: 2022},
								run: (*parser).callonstr13,
								expr: &seqExpr{
									pos: position{line: 105, col: 2, offset: 2022},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 105, col: 2, offset: 2022},
											val:        "'",
											ignoreCase: false,
										},
										&labeledExpr{
											pos:   position{line: 105, col: 6, offset: 2026},
											label: "v",
											expr: &zeroOrMoreExpr{
												pos: position{line: 105, col: 8, offset: 2028},
												expr: &choiceExpr{
													pos: position{line: 105, col: 9, offset: 2029},
													alternatives: []interface{}{
														&ruleRefExpr{
															pos:  position{line: 105, col: 9, offset: 2029},
															name: "escape",
														},
														&charClassMatcher{
															pos:        position{line: 105, col: 18, offset: 2038},
															val:        "[^']",
															chars:      []rune{'\''},
															ignoreCase: false,
															inverted:   true,
														},
													},
												},
											},
										},
										&litMatcher{
											pos:        position{line: 105, col: 25, offset: 2045},
											val:        "'",
											ignoreCase: false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "escape",
			pos:  position{line: 110, col: 1, offset: 2111},
			expr: &actionExpr{
				pos: position{line: 110, col: 10, offset: 2120},
				run: (*parser).callonescape1,
				expr: &seqExpr{
					pos: position{line: 110, col: 10, offset: 2120},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 110, col: 10, offset: 2120},
							val:        "\\",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 110, col: 15, offset: 2125},
							label: "char",
							expr: &anyMatcher{
								line: 110, col: 20, offset: 2130,
							},
						},
					},
				},
			},
		},
		{
			name: "null",
			pos:  position{line: 112, col: 1, offset: 2157},
			expr: &choiceExpr{
				pos: position{line: 112, col: 8, offset: 2164},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 112, col: 8, offset: 2164},
						val:        "null",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 112, col: 17, offset: 2173},
						run: (*parser).callonnull3,
						expr: &litMatcher{
							pos:        position{line: 112, col: 17, offset: 2173},
							val:        "nil",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 114, col: 1, offset: 2203},
			expr: &zeroOrMoreExpr{
				pos: position{line: 114, col: 18, offset: 2220},
				expr: &charClassMatcher{
					pos:        position{line: 114, col: 18, offset: 2220},
					val:        "[ \\t\\r\\n]",
					chars:      []rune{' ', '\t', '\r', '\n'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "SP",
			pos:  position{line: 115, col: 1, offset: 2232},
			expr: &zeroOrMoreExpr{
				pos: position{line: 115, col: 6, offset: 2237},
				expr: &charClassMatcher{
					pos:        position{line: 115, col: 6, offset: 2237},
					val:        "[ \\t]",
					chars:      []rune{' ', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "GS",
			pos:  position{line: 116, col: 1, offset: 2245},
			expr: &oneOrMoreExpr{
				pos: position{line: 116, col: 6, offset: 2250},
				expr: &charClassMatcher{
					pos:        position{line: 116, col: 6, offset: 2250},
					val:        "[ \\t]",
					chars:      []rune{' ', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "NL",
			pos:  position{line: 117, col: 1, offset: 2258},
			expr: &seqExpr{
				pos: position{line: 117, col: 6, offset: 2263},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 117, col: 6, offset: 2263},
						name: "SP",
					},
					&zeroOrMoreExpr{
						pos: position{line: 117, col: 9, offset: 2266},
						expr: &charClassMatcher{
							pos:        position{line: 117, col: 9, offset: 2266},
							val:        "[\\r\\n]",
							chars:      []rune{'\r', '\n'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 117, col: 17, offset: 2274},
						name: "_",
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 119, col: 1, offset: 2279},
			expr: &notExpr{
				pos: position{line: 119, col: 7, offset: 2285},
				expr: &anyMatcher{
					line: 119, col: 8, offset: 2286,
				},
			},
		},
	},
}

func (c *current) onstart1(b interface{}) (interface{}, error) {

	return b, nil
}

func (p *parser) callonstart1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstart1(stack["b"])
}

func (c *current) onComment8(ch interface{}) (interface{}, error) {
	return ch, nil
}

func (p *parser) callonComment8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment8(stack["ch"])
}

func (c *current) onComment3(v interface{}) (interface{}, error) {
	return v, nil
}

func (p *parser) callonComment3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment3(stack["v"])
}

func (c *current) onComment20(ch interface{}) (interface{}, error) {
	return ch, nil
}

func (p *parser) callonComment20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment20(stack["ch"])
}

func (c *current) onComment15(v interface{}) (interface{}, error) {
	return v, nil
}

func (p *parser) callonComment15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment15(stack["v"])
}

func (c *current) onComment32(ch interface{}) (interface{}, error) {
	return ch, nil
}

func (p *parser) callonComment32() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment32(stack["ch"])
}

func (c *current) onComment27(v interface{}) (interface{}, error) {
	return v, nil
}

func (p *parser) callonComment27() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment27(stack["v"])
}

func (c *current) onContent1(b interface{}) (interface{}, error) {

	verbs := make([]*verb, 0)

	if b != nil {
		for _, v := range toIfaceSlice(b) {
			if verb, ok := v.(*verb); ok {
				verbs = append(verbs, verb)
			}
		}
	}

	return verbs, nil
}

func (p *parser) callonContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onContent1(stack["b"])
}

func (c *current) onVerb1(name, args, body interface{}) (interface{}, error) {

	n := name.(string)
	a := toIfaceSlice(args)

	var b []*verb
	if body != nil {
		b = body.([]*verb)
	}

	return &verb{n, a, b}, nil
}

func (p *parser) callonVerb1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVerb1(stack["name"], stack["args"], stack["body"])
}

func (c *current) onArg1(k interface{}) (interface{}, error) {
	return k, nil
}

func (p *parser) callonArg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArg1(stack["k"])
}

func (c *current) onVariable1(f, r, a interface{}) (interface{}, error) {

	return &variable{f.(string), toIfaceSlice(r), toIfaceSlice(a)}, nil
}

func (p *parser) callonVariable1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariable1(stack["f"], stack["r"], stack["a"])
}

func (c *current) onMethodArgs1(args interface{}) (interface{}, error) {

	return args, nil
}

func (p *parser) callonMethodArgs1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMethodArgs1(stack["args"])
}

func (c *current) onCompoundPath2(i interface{}) (interface{}, error) {
	return i, nil
}

func (p *parser) callonCompoundPath2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCompoundPath2(stack["i"])
}

func (c *current) onCompoundPath7(p interface{}) (interface{}, error) {
	return p, nil
}

func (p *parser) callonCompoundPath7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCompoundPath7(stack["p"])
}

func (c *current) onBody1(b interface{}) (interface{}, error) {

	return b, nil
}

func (p *parser) callonBody1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBody1(stack["b"])
}

func (c *current) onidentity1(f, r interface{}) (interface{}, error) {

	str := string(f.([]byte)[:]) + ifaceToString(r)
	return str, nil
}

func (p *parser) callonidentity1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onidentity1(stack["f"], stack["r"])
}

func (c *current) onboolean2() (interface{}, error) {
	return true, nil
}

func (p *parser) callonboolean2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onboolean2()
}

func (c *current) onboolean4() (interface{}, error) {
	return false, nil
}

func (p *parser) callonboolean4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onboolean4()
}

func (c *current) onnumber1() (interface{}, error) {

	return strconv.ParseFloat(string(c.text[:]), 64)
}

func (p *parser) callonnumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onnumber1()
}

func (c *current) onunsigned1() (interface{}, error) {

	return strconv.ParseUint(string(c.text[:]), 10, 64)
}

func (p *parser) callonunsigned1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onunsigned1()
}

func (c *current) onstr4(v interface{}) (interface{}, error) {
	return v, nil
}

func (p *parser) callonstr4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstr4(stack["v"])
}

func (c *current) onstr13(v interface{}) (interface{}, error) {
	return v, nil
}

func (p *parser) callonstr13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstr13(stack["v"])
}

func (c *current) onstr1(v interface{}) (interface{}, error) {

	return ifaceToString(v), nil
}

func (p *parser) callonstr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onstr1(stack["v"])
}

func (c *current) onescape1(char interface{}) (interface{}, error) {
	return char, nil
}

func (p *parser) callonescape1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onescape1(stack["char"])
}

func (c *current) onnull3() (interface{}, error) {
	return nil, nil
}

func (p *parser) callonnull3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onnull3()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
