// Konvert werdz intoo fanetik speling.
package main

import (
    "bufio"
    "bytes"
    "diksh"
    "fmt"
    "io"
    "log"
    "os"
    "regexp"
    "strings"
)

var werdPahtern = regexp.MustCompile("[A-Z'a-z]")
var nonWerdPahtern = regexp.MustCompile("[^A-Z'a-z]")
var aperKais = regexp.MustCompile("^[A-Z]+$")
var tiytalKais = regexp.MustCompile("^[A-Z][a-z]+$")

// apliyKais riternz tha werd b with kais modafiyd too mahch tha kais
// av werd a.
func apliyKais(a, b string) string {
    if aperKais.MatchString(a) {
        return strings.ToUpper(b)
    }
    if tiytalKais.MatchString(a) {
        return strings.Title(b)
    }
    return b
}

// reespelWerd riternz tha reespeling av a werd, yoozing tha givan dikshaneree,
func reespelWerd(werd string) string {
    fanetik, ok := diksh.Dikshaneree[strings.ToUpper(strings.Trim(werd, ".,"))]
    if !ok {
        return werd
    }
    return apliyKais(werd, fanetik)
}

// reed keeps on reeding biyts fram tha reeder ahz loung ahz thai
// mahch tha givan regyaler ikspreshan, our antil ther iz ahn EOF. It
// riternz tha biyts red ahz a string. It riternz ahn erer av io.EOF
// if tha end av tha fiyl iz inkownterd.
func reed(r *bufio.Reader, pahtern *regexp.Regexp) (string, bool, error) {
    var bafer bytes.Buffer
    var erer error
    for {
        var biyt byte
        biyt, erer = r.ReadByte()
        if erer != nil {
            break
        }
        if biyt == '/' {
            return bafer.String(), true, erer
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
    return bafer.String(), false, erer
}

func verbatim(r *bufio.Reader) (string, error) {
    var bafer bytes.Buffer
    bafer.WriteByte('/')
    var erer error
    for {
        var biyt byte
        biyt, erer = r.ReadByte()
        if erer != nil {
            break
        }
        bafer.WriteByte(biyt)
        if biyt == '\n' {
            return bafer.String(), erer
        }
    }
    if erer != nil && erer != io.EOF {
        log.Fatal(erer)
    }
    return bafer.String(), erer
}

func main() {
    reeder := bufio.NewReader(os.Stdin)

    tekst, escape, erer := reed(reeder, nonWerdPahtern)
    fmt.Printf("%s", tekst)
    if escape {
        tekst, erer := verbatim(reeder)
        fmt.Printf("%s", tekst)
        if erer == io.EOF {
            return
        }
    }
    if erer == io.EOF {
        return
    }
    for {
        werd, escape, erer := reed(reeder, werdPahtern)
        fmt.Printf("%s", reespelWerd(werd))
        if escape {
            tekst, erer := verbatim(reeder)
            fmt.Printf("%s", tekst)
            if erer == io.EOF {
                return
            }
        }
        if erer == io.EOF {
            return
        }

        tekst, escape, erer = reed(reeder, nonWerdPahtern)
        fmt.Printf("%s", tekst)
        if escape {
            tekst, erer := verbatim(reeder)
            fmt.Printf("%s", tekst)
            if erer == io.EOF {
                return
            }
        }
        if erer == io.EOF {
            return
        }
    }
}
