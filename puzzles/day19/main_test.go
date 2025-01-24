package main

import (
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	avail, design := ParsePuzzleInput(false, "day19.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part1(avail, design)
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}

func TestParsing(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	expected_avail := []string{
		"r", "wr", "b", "g", "bwu", "rb", "gb", "br",
	}
	slices.Sort(expected_avail)
	sort.Slice(expected_avail, func(i, j int) bool {
		return len(expected_avail[i]) > len(expected_avail[j])
	})
	expected_design := []string{
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}

	if !reflect.DeepEqual(expected_avail, avail) {
		t.Errorf("Expected: %+v. Actual: %+v", expected_avail, avail)
	}
	if !reflect.DeepEqual(expected_design, design) {
		t.Errorf("Expected: %+v. Actual: %+v", expected_design, design)
	}
}

func TestSplitDesign(t *testing.T) {
	test_str := "brwbrwbrwbrwbrwbrw"
	m := regexp.MustCompile("w")

	idx := m.FindAllIndex([]byte(test_str), -1)
	e := []string{
		"br",
		"brwbrwbrwbrwbrw",
	}
	a := splitUpDesign(idx[0], test_str)

	if !reflect.DeepEqual(e, a) {
		t.Errorf("Expected: %+v. Actual: %+v.", e, a)
	}
}

func TestPart1(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 6
	a := Part1(avail, design)
	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCase1(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := map[string]int{
		"br": 1,
		"wr": 1,
		"r":  1,
	}
	a := matchDesign(avail, design[0])

	if !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase2(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := map[string]int{
		"b": 1,
		"g": 2,
		"r": 1,
	}
	a := matchDesign(avail, design[1])

	if !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase3(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := map[string]int{
		"gb": 1,
		"br": 1,
	}
	a := matchDesign(avail, design[2])

	if !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase4(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e1 := map[string]int{
		"r":  1,
		"rb": 1,
		"g":  1,
		"br": 1,
	}
	e2 := map[string]int{
		"r":  2,
		"rb": 1,
		"gb": 1,
	}
	a := matchDesign(avail, design[3])

	if !reflect.DeepEqual(e1, a) && !reflect.DeepEqual(e2, a) {
		t.Errorf("Expected %+v or %+v. Actual %+v", e1, e2, a)
	}
}

func TestCase5(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	a := matchDesign(avail, design[4])

	if a != nil {
		t.Errorf("Expected %+v. Actual %+v", nil, a)
	}
}

func TestCase6(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := map[string]int{
		"bwu": 1,
		"r":   2,
		"g":   1,
	}
	a := matchDesign(avail, design[5])

	if !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase7(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := map[string]int{
		"br": 1,
		"r":  1,
		"g":  1,
	}
	a := matchDesign(avail, design[6])

	if !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase8(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	a := matchDesign(avail, design[7])

	if a != nil {
		t.Errorf("Expected %+v. Actual %+v", nil, a)
	}
}

func TestCase9(t *testing.T) {
	design := "burwurwrwgwbrbggruuuwwwuwuurgubwwguburrgbwgwbw"
	availTowels := strings.Split("grrg, buwuurr, w, bguwu, uuuu, wr, wuw, brbug, buwrg, wwr, uwr, buu, gurw, ubrgg, wggww, urgbr, wur, urur, wguw, buurb, ugbrwwb, gbr, gb, wurww, wgruw, rrwgbur, wrg, wru, urr, wwbgru, bb, rrr, gwurbgb, bgguur, rurru, rwggru, uwbuwbr, rwww, grbbrr, uuw, ubgwu, rwrrr, rwgrbu, rbu, ub, buubgrr, ruu, rggwgr, bgwgwu, wuwgbg, bgugg, bbb, gbwur, buwrubb, uwgr, grrggug, brgu, ubg, ggur, uuwr, gbuu, bwwgrb, wbu, bbrbb, ubbrr, rrwrggb, urbu, ubrb, grb, brww, uuu, uuwg, bbwr, wwgg, buwubg, gwuuru, bguugrwu, rwwugr, rwru, rwwwwg, bgug, urru, bwuug, bugrr, wggr, buw, gub, gwgrgg, brugwg, wug, grbrw, bgbbb, ubwrg, rubr, grrubug, rubuugwr, gbgw, rwgwgg, ubw, bwbwwrb, brbwg, gbgg, rgwrur, guwg, grg, rwwubg, ugu, urbwrg, urg, gugbwbr, gwg, gbg, rruwug, wuwwug, gwgbg, wuuubu, rwuwbb, urrwur, rrgbubw, rub, rgr, uuuww, wrgu, ggg, brwb, rgrwg, bwr, ugwggwb, wggbguw, rwwgb, ubr, gbwg, rubug, bwg, bru, uwu, wbwggbr, ggu, guuub, uur, bww, ruguwgg, wbwrb, brwg, rrw, r, wgrb, ggrb, uwb, uuugbbr, brubb, bggb, urwubrgg, bgbw, rbgrrrgg, ubbg, rww, uuruggrw, rugw, uug, bgbgr, uwrb, wg, gubr, wbb, gwr, wwrw, wu, wwrgu, ugrrgu, bbguu, uwwwg, rrg, uuuwgu, grbbuwg, gguu, gugg, gurg, bgg, bug, ugb, rwr, wgbbg, gwgwuu, wuwbbbw, wgbr, wuggur, gwgr, bbuubuww, uub, br, bggwu, brg, rbg, rwb, brwwwg, ruwrg, wbr, ubwwww, gbww, rbww, rwuwbu, wbuburw, rrgg, uugbub, urrbb, wwgrb, wrub, guu, rgugb, wrubu, rbwgwg, gbgwubr, uwbrbu, wuru, rbbgu, wwg, rbguugr, brb, wbgbrr, guw, rguw, rbrgub, bgu, gurrrg, ggrbg, guwr, bwbb, wbgg, ugggb, rrrbggb, rurg, gug, rwg, urrrb, wbubg, wrbbubw, bwgu, bwurw, urrg, gbw, rrbr, rbrg, brwbww, rbgb, uurb, bwwgwgrg, wrwbuu, bwbgw, rug, ww, grr, rg, rbwbrbrr, rgrg, wwb, rurr, ubwu, rgu, www, gbb, bbw, rbw, gbgbwbw, ur, gguwuggr, gru, uru, rbuw, rrbgr, wuu, wgg, bg, wbuu, rwbggb, uggbg, gbu, bbgruuuw, rb, rggb, ubbuub, ruwubgbr, bub, guuru, ugg, uuggbuw, rbuur, gbrw, gu, wbuugwr, ubb, urbuuwu, wrgub, rwu, uuww, wgwgbg, uww, rwwb, rbbubggg, bgb, rrwuwwg, ubugw, brw, gww, bgr, rbur, rr, buwg, wub, wbwwuu, wuruuwr, wbur, rgbwu, rbgwurg, uwg, uwrrg, bruuw, rgg, gubu, bbbubbb, bgrw, urb, wgw, rrgwu, wubbgr, bbr, bwuwu, rbbu, rrb, wuuwggww, ubwg, uwgub, bgubgru, rwubu, wuwgbgg, ggb, g, bbrwrgwu, wbg, wwru, rbwugr, urwwuu, wwwgug, gbur, bbggw, gwbrru, wrbr, ggr, ubu, gbwugw, u, grur, rgub, bggr, bwrw, wurb, ggggw, gwurrwrw, bgrubw, wgr, rrbrrwb, urwbubgr, grw, rgw, bbrbr, urbrr, brrbruuu, uwbbw, rur, ggw, ruw, uu, ubrw, rurgb, gbrr, gwwbu, ugr, brwubuub, bwb, wuugr, wrrub, uugrbr, bgwr, gwu, gwbw, wbru, gwb, rgrwur, wrwur, bw, bbu, wwu, uwwb, gbbb, grrb, uwuwbwrw, gur, rgbbu, gbwbgg, bbrur, gw, ru, guwwbr, wgubb, ggug, ugw, wwgwugg, wgrw, grrw, uuggb, rwbgbr, wrb, wbw, bbg, rw, uwgbgu, bwgw, gr, rgwb, rbr, rugubb, gwuggbb, ruww, bwu, rwuggr, bubrb, rbbub, uggr, brwugur, wgu, bbwb, gbwbu, rwbubu, ug, wb, rrruu, bwbbrb, rbb, uwwgu, bgw, rwwr, urw, wgb, gg, bruwrbr, uugwgbgu, rru", ", ")
	slices.Sort(availTowels)
	sort.Slice(availTowels, func(i, j int) bool {
		return len(availTowels[i]) > len(availTowels[j])
	})

	a := matchDesign(availTowels, design)
	if a != nil {
		t.Errorf("expected: %+v, actual: %+v", nil, a)
	}
}

func TestCaseAll1(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 2
	a := findAllCombo(avail, design[0])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll2(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a := findAllCombo(avail, design[1])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll3(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 4
	a := findAllCombo(avail, design[2])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll4(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 6
	a := findAllCombo(avail, design[3])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll5(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 0
	a := findAllCombo(avail, design[4])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll6(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a := findAllCombo(avail, design[5])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll7(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 2
	a := findAllCombo(avail, design[6])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll8(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 0
	a := findAllCombo(avail, design[7])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestPart2Sample(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 16
	a := Part2(avail, design)
	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestPart2Puzzle(t *testing.T) {
	design := "ugruuruwwgwguggubuuwrubwuuguwgrbuubrggwbbbggbuubgubgrwbub"
	availTowels := strings.Split("grrg, buwuurr, w, bguwu, uuuu, wr, wuw, brbug, buwrg, wwr, uwr, buu, gurw, ubrgg, wggww, urgbr, wur, urur, wguw, buurb, ugbrwwb, gbr, gb, wurww, wgruw, rrwgbur, wrg, wru, urr, wwbgru, bb, rrr, gwurbgb, bgguur, rurru, rwggru, uwbuwbr, rwww, grbbrr, uuw, ubgwu, rwrrr, rwgrbu, rbu, ub, buubgrr, ruu, rggwgr, bgwgwu, wuwgbg, bgugg, bbb, gbwur, buwrubb, uwgr, grrggug, brgu, ubg, ggur, uuwr, gbuu, bwwgrb, wbu, bbrbb, ubbrr, rrwrggb, urbu, ubrb, grb, brww, uuu, uuwg, bbwr, wwgg, buwubg, gwuuru, bguugrwu, rwwugr, rwru, rwwwwg, bgug, urru, bwuug, bugrr, wggr, buw, gub, gwgrgg, brugwg, wug, grbrw, bgbbb, ubwrg, rubr, grrubug, rubuugwr, gbgw, rwgwgg, ubw, bwbwwrb, brbwg, gbgg, rgwrur, guwg, grg, rwwubg, ugu, urbwrg, urg, gugbwbr, gwg, gbg, rruwug, wuwwug, gwgbg, wuuubu, rwuwbb, urrwur, rrgbubw, rub, rgr, uuuww, wrgu, ggg, brwb, rgrwg, bwr, ugwggwb, wggbguw, rwwgb, ubr, gbwg, rubug, bwg, bru, uwu, wbwggbr, ggu, guuub, uur, bww, ruguwgg, wbwrb, brwg, rrw, r, wgrb, ggrb, uwb, uuugbbr, brubb, bggb, urwubrgg, bgbw, rbgrrrgg, ubbg, rww, uuruggrw, rugw, uug, bgbgr, uwrb, wg, gubr, wbb, gwr, wwrw, wu, wwrgu, ugrrgu, bbguu, uwwwg, rrg, uuuwgu, grbbuwg, gguu, gugg, gurg, bgg, bug, ugb, rwr, wgbbg, gwgwuu, wuwbbbw, wgbr, wuggur, gwgr, bbuubuww, uub, br, bggwu, brg, rbg, rwb, brwwwg, ruwrg, wbr, ubwwww, gbww, rbww, rwuwbu, wbuburw, rrgg, uugbub, urrbb, wwgrb, wrub, guu, rgugb, wrubu, rbwgwg, gbgwubr, uwbrbu, wuru, rbbgu, wwg, rbguugr, brb, wbgbrr, guw, rguw, rbrgub, bgu, gurrrg, ggrbg, guwr, bwbb, wbgg, ugggb, rrrbggb, rurg, gug, rwg, urrrb, wbubg, wrbbubw, bwgu, bwurw, urrg, gbw, rrbr, rbrg, brwbww, rbgb, uurb, bwwgwgrg, wrwbuu, bwbgw, rug, ww, grr, rg, rbwbrbrr, rgrg, wwb, rurr, ubwu, rgu, www, gbb, bbw, rbw, gbgbwbw, ur, gguwuggr, gru, uru, rbuw, rrbgr, wuu, wgg, bg, wbuu, rwbggb, uggbg, gbu, bbgruuuw, rb, rggb, ubbuub, ruwubgbr, bub, guuru, ugg, uuggbuw, rbuur, gbrw, gu, wbuugwr, ubb, urbuuwu, wrgub, rwu, uuww, wgwgbg, uww, rwwb, rbbubggg, bgb, rrwuwwg, ubugw, brw, gww, bgr, rbur, rr, buwg, wub, wbwwuu, wuruuwr, wbur, rgbwu, rbgwurg, uwg, uwrrg, bruuw, rgg, gubu, bbbubbb, bgrw, urb, wgw, rrgwu, wubbgr, bbr, bwuwu, rbbu, rrb, wuuwggww, ubwg, uwgub, bgubgru, rwubu, wuwgbgg, ggb, g, bbrwrgwu, wbg, wwru, rbwugr, urwwuu, wwwgug, gbur, bbggw, gwbrru, wrbr, ggr, ubu, gbwugw, u, grur, rgub, bggr, bwrw, wurb, ggggw, gwurrwrw, bgrubw, wgr, rrbrrwb, urwbubgr, grw, rgw, bbrbr, urbrr, brrbruuu, uwbbw, rur, ggw, ruw, uu, ubrw, rurgb, gbrr, gwwbu, ugr, brwubuub, bwb, wuugr, wrrub, uugrbr, bgwr, gwu, gwbw, wbru, gwb, rgrwur, wrwur, bw, bbu, wwu, uwwb, gbbb, grrb, uwuwbwrw, gur, rgbbu, gbwbgg, bbrur, gw, ru, guwwbr, wgubb, ggug, ugw, wwgwugg, wgrw, grrw, uuggb, rwbgbr, wrb, wbw, bbg, rw, uwgbgu, bwgw, gr, rgwb, rbr, rugubb, gwuggbb, ruww, bwu, rwuggr, bubrb, rbbub, uggr, brwugur, wgu, bbwb, gbwbu, rwbubu, ug, wb, rrruu, bwbbrb, rbb, uwwgu, bgw, rwwr, urw, wgb, gg, bruwrbr, uugwgbgu, rru", ", ")
	slices.Sort(availTowels)
	sort.Slice(availTowels, func(i, j int) bool {
		return len(availTowels[i]) > len(availTowels[j])
	})

	a := findAllCombo(availTowels, design)
	if a == 0 {
		t.Errorf("Expected non-nil. Actual: %+v", a)
	}
}
