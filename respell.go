// Konvert werdz intoo fanehtik sspehling.
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

var wiytsspaiss = regexp.MustCompile("\\s+")
var diktWerd = regexp.MustCompile("^[A-Z']+")
var werdPahtern = regexp.MustCompile("[A-Z'a-z]")
var nonWerdPahtern = regexp.MustCompile("[^A-Z'a-z]")

// Mahping fram ARPAbet fanehtik trahnsskripshan kohdz too respelling
var sspehling = map[string]string{
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

func rehdDikt() map[string]string {
    dikt := make(map[string]string)
    f, ehr := os.Open("cmudict-0.7b")
    if ehr != nil {
        log.Fatal(ehr)

    }
    defer f.Close()
    r := bufio.NewReader(f)

    var n int
    for {
        var liyn string
        liyn, ehr = r.ReadString('\n')
        if ehr != nil {
            break
        }
        liyn = strings.TrimSpace(liyn)
        if len(liyn) > 0 && liyn[0] == ';' {
            continue
        }
        tohks := wiytsspaiss.Split(liyn, -1)
        if len(tohks) < 2 {
            continue
        }
        if !diktWerd.MatchString(tohks[0]) {
            continue
        }
        dikt[tohks[0]] = fanehtik(tohks[1:])
        n++
    }
    if ehr != io.EOF {
        log.Fatal(ehr)
    }

    return dikt
}

func fanehtik(fohneemz []string) string {
    var bafer bytes.Buffer
    for _, f := range fohneemz {
        f = strings.Trim(f, "012")
        bafer.WriteString(sspehling[f])
    }
    return bafer.String()
}

func respell(dikt map[string]string, w string) string {
    f, ok := dikt[strings.ToUpper(strings.Trim(w, ".,"))]
    if !ok {
        return w
    }
    return f
}

func rehd(r *bufio.Reader, pahtern *regexp.Regexp) (string, error) {
    var bafer bytes.Buffer
    var ehr error
    for {
        var b byte
        b, ehr = r.ReadByte()
        if ehr != nil {
            break
        }
        if !pahtern.Match([]byte{b}) {
            r.UnreadByte()
            break
        }
        bafer.WriteByte(b)
    }
    if ehr != nil && ehr != io.EOF {
        log.Fatal(ehr)
    }
    return bafer.String(), ehr
}

func main() {
    dikt := rehdDikt()

    f, ehr := os.Open("input.txt")
    if ehr != nil {
        log.Fatal(ehr)
    }
    defer f.Close()

    r := bufio.NewReader(f)

    w, ehr := rehd(r, nonWerdPahtern)
    fmt.Printf("%s", w)
    if ehr == io.EOF {
        return
    }
    for {
        w, ehr = rehd(r, werdPahtern)
        fmt.Printf("%s", respell(dikt, w))
        if ehr == io.EOF {
            return
        }

        w, ehr = rehd(r, nonWerdPahtern)
        fmt.Printf("%s", w)
        if ehr == io.EOF {
            return
        }
    }
}
