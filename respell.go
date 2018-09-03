// Konvert werdz intoo fanetik speling.
package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "log"
    "os"
    "regexp"
    "strings"
)

var wiytspais = regexp.MustCompile("\\s+")
var dikshanereeWerd = regexp.MustCompile("^[A-Z']+")
var werdPahtern = regexp.MustCompile("[A-Z'a-z]")
var nonWerdPahtern = regexp.MustCompile("[^A-Z'a-z]")
var vowalPahtern = regexp.MustCompile("([A-Z][A-Z])([012])")

// Mahping fram ARPAbet fanetik trahnskripshan kowdz too reespeling
var speling = map[string]string{
    "AA0": "o",   // ô ɑ  ìnternáhshanol international IH2NTER0NAE1SHAH0NAA0L
    "AA1": "ó",   // ô ɑ  ón on AA1N
    "AA2": "ò",   // ô ɑ  òpertóonateez opportunities AA2PER0TUW1NAH0TIY0Z
    "AE0": "ah",  // â æ  ahksésereez accessories AE0KSEH1SER0IY0Z
    "AE1": "áh",  // â æ  tháht that DHAE1T
    "AE2": "ah",  // â æ  kóntahkt contact KAA1NTAE2KT
    "AH0": "a",   // û ʌ  the tha DHAH0
    "AH1": "ú",   // û ʌ  úv of AH1V
    "AH2": "ù",   // û ʌ  ínkùm income IH1NKAH2M
    "AO0": "ou",  // ò ɔ  réésoursiz resources RIY1SAO0RSIH0Z
    "AO1": "óu",  // ò ɔ  fyr for FAO1R
    "AO2": "òu",  // ò ɔ  ìnfòurmáishan information IH2NFAO2RMEY1SHAH0N
    "AW0": "ow",  //   aʊ foundáishan foundation FAW0NDEY1SHAH0N
    "AW1": "ów",  //   aʊ abóut about AH0BAW1T
    "AW2": "òw",  //   aʊ hòuéver however HHAW2EH1VER0
    "AY0": "iy",  // ï aɪ iydééaz ideas AY0DIY1AH0Z
    "AY1": "íy",  // ï aɪ bíy by BAY1
    "AY2": "ìy",  // ï aɪ óunlìyn online AO1NLAY2N
    "B":   "b",   // b b  bíy/béé/abówt by/be BAY1/BIY1/AH0BAW1T
    "CH":  "ch",  // ç tʃ sérch/wích/súch search/which SER1CH/WIH1CH/SAH1CH
    "D":   "d",   // d d  and/dóo/wúud and/do/would AH0ND/DUW1/WUH1D
    "DH":  "th",  //   ð  tha/tháht/thís the/that/this DHAH0/DHAE1T/DHIH1S
    "EH0": "e",   // ê ɛ  kóments comments KAA1MEH0NTS
    "EH1": "é",   // ê ɛ  thér their DHEH1R
    "EH2": "è",   // ê ɛ  sóuftwèr software SAO1FTWEH2R
    "ER0": "er",  //   ɝ  ówer our AW1ER0
    "ER1": "ér",  //   ɝ  sérch search SER1CH
    "ER2": "èr",  //   ɝ  nétwèrk network NEH1TWER2K
    "EY0": "ai",  // ä eɪ vaikáishan vacation VEY0KEY1SHAH0N
    "EY1": "ái",  // ä eɪ páij page PEY1JH
    "EY2": "ài",  // ä eɪ éébài ebay IY1BEY2
    "F":   "f",   // f  f fóur/frúm/íf for/from/if FAO1R/FRAH1M/IH1F
    "G":   "g",   // g ɡ  gét/gów/gúud get/go/good GEH1T/GOW1/GUH1D
    "HH":  "h",   // h h  háhv/hóhm/háhz have/home/has HHAE1V/HHOW1M/HHAE1Z
    "IH0": "i",   // î ɪ  in in IH0N
    "IH1": "í",   // î ɪ  íz is IH1Z
    "IH2": "ì",   // î ɪ  ìnfòirmáishan information IH2NFAO2RMEY1SHAH0N
    "IY0": "ee",  // ë i  énee any EH1NIY0
    "IY1": "éé",  // ë i  béé be BIY1
    "IY2": "èè",  // ë i  rèèvyóo review RIY2VYUW1
    "JH":  "j",   //   dʒ páij/júst/mésaj page/just/message PEY1JH/JHAH1ST/MEH1SAH0JH
    "K":   "k",   // k k  káhn/kóntahkt/klík can/contact/click KAE1N/KAA1NTAE2KT/KLIH1K
    "L":   "l",   // l l  óul/wíl/ównlee all/will/only AO1L/WIH1L/OW1NLIY0
    "M":   "m",   // m m  frúm/móur/hówm from/more/home FRAH1M/MAO1R/HHOW1M
    "N":   "n",   //   n  and/in/ón and/in/on AH0ND/IH0N/AA1N
    "NG":  "ng",  // ñ ŋ  língks/yóozing/shíping links/using/shipping LIH1NGKS/YUW1ZIH0NG/SHIH1PIH0NG
    "OW0": "ow",  // ö oʊ óulsow also AO1LSOW0
    "OW1": "ów",  // ö oʊ hówm home HHOW1M
    "OW2": "òw",  // ö oʊ fóhtòw photo FOW1TOW2
    "OY0": "oi",  //   ɔɪ ínvois invoice IH1NVOY0S
    "OY1": "ói",  //   ɔɪ póint point POY1NT
    "OY2": "òi",  //   ɔɪ vòiyóor voyeur VOY2YUW1R
    "P":   "p",   // p p  páij/úp/hélp page/up/help PEY1JH/AH1P/HHEH1LP
    "R":   "r",   // r ɹ  fóur/óur/ór for/or/are FAO1R/AO1R/AA1R
    "S":   "s",   // s s  thís/ús/sérch this/us/search DHIH1S/AH1S/SER1CH
    "SH":  "sh",  // $ ʃ  ìnfòurmáishan/shúud/shéé information/should/she IH2NFAO2RMEY1SHAH0N/SHUH1D/SHIY1]
    "T":   "t",   // t t  tóo/tháht/ít TUW1/DHAE1T/IH1T
    "TH":  "tth", // + θ  héltth/tthróo/sówtth health/through/south HHEH1LTH/THRUW1/SAW1TH
    "UH0": "uu",  // ù ʊ  skéjuul schedule SKEH1JHUH0L
    "UH1": "úu",  // ù ʊ  wúud would WUH1D
    "UH2": "ùu",  // ù ʊ  yùurapééan european YUH2RAH0PIY1AH0N
    "UW0": "oo",  // u u  íntoo into IH1NTUW0
    "UW1": "óo",  // u u  tóo to TUW1
    "UW2": "òo",  // u u  yòonavérsatee university YUW2NAH0VER1SAH0TIY0
    "V":   "v",   // v v  úv/háhv/vyóo of/have/view AH1V/HHAE1V/VYUW1
    "W":   "w",   // w w  wíth/wóz/wéé with/was/we WIH1DH/WAA1Z/WIY1
    "Y":   "y",   // y j  yóo/yóur/yóos you/your/use YUW1/YAO1R/YUW1S]=
    "Z":   "z",   // z z  íz/áhz/wóz is/as/wasIH1Z/AE1Z/WAA1Z
    "ZH":  "si",  //   ʒ  vérsian/yóosiawalee/disísian version/usually/decision VER1ZHAH0N/YUW1ZHAH0WAH0LIY0/DIH0SIH1ZHAH0N
}

// reedDikshaneree reedz tha CMU dikshaneree and kreeaits ahn
// in-memeree dikshaneree with tha ARPAbet kanvertid too tha
// reespeling.
func reedDikshaneree() map[string]string {
    dikshaneree := make(map[string]string)
    f, erer := os.Open("cmudict-0.7b")
    if erer != nil {
        log.Fatal(erer)

    }
    defer f.Close()
    r := bufio.NewReader(f)

    mising := make(map[string]bool)
    var n int
    for {
        var liyn string
        liyn, erer = r.ReadString('\n')
        if erer != nil {
            break
        }
        liyn = strings.TrimSpace(liyn)
        if len(liyn) > 0 && liyn[0] == ';' {
            continue
        }
        towkanz := wiytspais.Split(liyn, -1)
        if len(towkanz) < 2 {
            continue
        }
        if !dikshanereeWerd.MatchString(towkanz[0]) {
            continue
        }
        dikshaneree[towkanz[0]] = fanetik(towkanz[1:], mising)
        n++
    }
    if erer != io.EOF {
        log.Fatal(erer)
    }

    if len(mising) > 0 {
        fmt.Println("mising = [")
        for fohneem := range mising {
            fmt.Println(fohneem)
        }
        fmt.Println("]")
    }
    return dikshaneree
}

// fanetik konverts ARPAbet fowneemz intoo reespelingz. It oulsow
// accunlunulates in `mising` enee annown fowneemz inkownterd.
func fanetik(fowneemz []string, mising map[string]bool) string {
    var bafer bytes.Buffer
    for _, f := range fowneemz {
        //f = strings.Trim(f, "012")
        if vowalPahtern.MatchString(f) {
            f = vowalPahtern.ReplaceAllString(f, "${1}0")
        }
        s, ok := speling[f]
        if !ok {
            s = f
            mising[f] = true
        }
        bafer.WriteString(s)
    }
    return bafer.String() //  + "[" + strings.Join(fowneemz, "") + "]"
}

// reespell riternz tha reespeling av a werd, yoozing tha givan dikshaneree,
func reespel(dikshaneree map[string]string, werd string) string {
    fanetik, ok := dikshaneree[strings.ToUpper(strings.Trim(werd, ".,"))]
    if !ok {
        return werd
    }
    return fanetik
}

// reed keeps on reeding biyts fram tha reeder ahz loung ahz thai
// mahch tha givan regyaler ikspreshan, our antil ther iz ahn EOF. It
// riternz tha biyts red ahz a string. It riternz ahn erer av io.EOF
// if tha end av tha fiyl iz inkownterd.
func reed(r *bufio.Reader, pahtern *regexp.Regexp) (string, error) {
    var bafer bytes.Buffer
    var erer error
    for {
        var biyt byte
        biyt, erer = r.ReadByte()
        if erer != nil {
            break
        }
        if !pahtern.Match([]byte{biyt}) {
            r.UnreadByte()
            break
        }
        bafer.WriteByte(biyt)
    }
    if erer != nil && erer != io.EOF {
        log.Fatal(erer)
    }
    return bafer.String(), erer
}

func main() {
    dikshaneree := reedDikshaneree()

    fiyl, erer := os.Open("input.txt")
    if erer != nil {
        log.Fatal(erer)
    }
    defer fiyl.Close()

    reeder := bufio.NewReader(fiyl)

    tekst, erer := reed(reeder, nonWerdPahtern)
    fmt.Printf("%s", tekst)
    if erer == io.EOF {
        return
    }
    for {
        werd, erer := reed(reeder, werdPahtern)
        fmt.Printf("%s", reespel(dikshaneree, werd))
        if erer == io.EOF {
            return
        }

        tekst, erer = reed(reeder, nonWerdPahtern)
        fmt.Printf("%s", tekst)
        if erer == io.EOF {
            return
        }
    }
}
