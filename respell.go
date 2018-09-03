// Convert words into phometic spelling.
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

var whitespace = regexp.MustCompile("\\s+")
var dictWord = regexp.MustCompile("^[A-Z']+")
var wordPattern = regexp.MustCompile("[A-Z'a-z]")
var nonWordPattern = regexp.MustCompile("[^A-Z'a-z]")

// Mapping from ARPAbet phonetic transcription codes to respelling
var spelling = map[string]string{
    "AA": "o",  // ô ɑ   odd     AA D        od
    "AE": "ah", // â æ   at      AE T        aht
    "AH": "a",  // û ʌ   hut     HH AH T     haht
    "AO": "o",  // ò ɔ   ought   AO T        ot
    "AW": "ow", //   aʊ  cow     K AW        kow
    "AY": "iy", // ï aɪ  hide    HH AY D     hiyd
    "B":  "b",  // b b   be      B IY        bee
    "CH": "ch", // ç tʃ  cheese  CH IY Z     cheez
    "D":  "d",  // d d   dee     D IY        dee
    "DH": "dh", //   ð   thee    DH IY       dhee
    "EH": "eh", // ê ɛ   Ed      EH D        ehd
    "ER": "er", //   ɝ   hurt    HH ER T     hert
    "EY": "ai", // ä eɪ  ate     EY T        ait
    "F":  "f",  // f  f   fee     F IY       fee
    "G":  "g",  // g ɡ   green   G R IY N    green
    "HH": "h",  // h h   he      HH IY       hee
    "IH": "i",  // î ɪ   it      IH T        it
    "IY": "ee", // ë i   eat     IY T        eet
    "JH": "dj", //   dʒ  gee     JH IY       djee
    "K":  "k",  // k k   key     K IY        kee
    "L":  "l",  // l l   lee     L IY        lee
    "M":  "m",  // m m   me      M IY        mee
    "N":  "n",  //   n   knee    N IY        nee
    "NG": "ng", // ñ ŋ   ping    P IH NG     ping
    "OW": "oh", // ö oʊ  oat     OW T        oht
    "OY": "oi", //   ɔɪ  toy     T OY        toi
    "P":  "p",  // p p   pee     P IY        pee
    "R":  "r",  // r ɹ   read    R IY D      reed
    "S":  "ss", // s s   sea     S IY        see
    "SH": "sh", // $ ʃ   she     SH IY       shee
    "T":  "t",  // t t   tea     T IY        tee
    "TH": "th", // + θ   theta   TH EY T AH  thaita
    "UH": "u",  // ù ʊ   hood    HH UH D     hud
    "UW": "oo", // u u   two     T UW        too
    "V":  "v",  // v v   vee     V IY        vee
    "W":  "w",  // w w   we      W IY        wee
    "Y":  "y",  // y j   yield   Y IY L D    yeeld
    "Z":  "z",  // z z   zee     Z IY        zee
    "ZH": "j",  //   ʒ   seizure S IY ZH ER  seejer
}

func readDict() map[string]string {
    dict := make(map[string]string)
    f, err := os.Open("cmudict-0.7b")
    if err != nil {
        log.Fatal(err)

    }
    defer f.Close()
    r := bufio.NewReader(f)

    var n int
    for {
        var line string
        line, err = r.ReadString('\n')
        if err != nil {
            break
        }
        line = strings.TrimSpace(line)
        if len(line) > 0 && line[0] == ';' {
            continue
        }
        toks := whitespace.Split(line, -1)
        if len(toks) < 2 {
            continue
        }
        if !dictWord.MatchString(toks[0]) {
            continue
        }
        dict[toks[0]] = phonetic(toks[1:])
        n++
    }
    if err != io.EOF {
        log.Fatal(err)
    }

    return dict
}

func phonetic(phonemes []string) string {
    var buffer bytes.Buffer
    for _, ph := range phonemes {
        ph = strings.Trim(ph, "012")
        buffer.WriteString(spelling[ph])
    }
    return buffer.String()
}

func respell(dict map[string]string, w string) string {
    ph, ok := dict[strings.ToUpper(strings.Trim(w, ".,"))]
    if !ok {
        return w
    }
    return ph
}

func read(r *bufio.Reader, patt *regexp.Regexp) (string, error) {
    var buffer bytes.Buffer
    var err error
    for {
        var b byte
        b, err = r.ReadByte()
        if err != nil {
            break
        }
        if !patt.Match([]byte{b}) {
            r.UnreadByte()
            break
        }
        buffer.WriteByte(b)
    }
    if err != nil && err != io.EOF {
        log.Fatal(err)
    }
    return buffer.String(), err
}

func main() {
    dict := readDict()

    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    r := bufio.NewReader(f)

    w, err := read(r, nonWordPattern)
    fmt.Printf("%s", w)
    if err == io.EOF {
        return
    }
    for {
        w, err = read(r, wordPattern)
        fmt.Printf("%s", respell(dict, w))
        if err == io.EOF {
            return
        }

        w, err = read(r, nonWordPattern)
        fmt.Printf("%s", w)
        if err == io.EOF {
            return
        }
    }
}
