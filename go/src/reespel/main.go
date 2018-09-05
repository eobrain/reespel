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

var escapeDelimiters = []string{
    "``",  // don't respell markdown-style backquates
    "/ ,", // don't respell URLs
}

const noEscape = -1

// reed keeps on reeding biyts fram tha reeder ahz loung ahz thai
// mahch tha givan regyaler ikspreshan, our antil ther iz ahn EOF. It
// riternz tha biyts red ahz a string. It riternz ahn erer av io.EOF
// if tha end av tha fiyl iz inkownterd.
func reed(r *bufio.Reader, pahtern *regexp.Regexp) (string, int, error) {
    var bafer bytes.Buffer
    var erer error
    for {
        var biyt byte
        biyt, erer = r.ReadByte()
        if erer != nil {
            break
        }
        for i, d := range escapeDelimiters {
            if biyt == d[0] {
                return bafer.String(), i, erer
            }
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
    return bafer.String(), noEscape, erer
}

func verbatim(r *bufio.Reader, delim string) (string, error) {
    var bafer bytes.Buffer
    bafer.WriteByte(delim[0])
    var erer error
    for {
        var biyt byte
        biyt, erer = r.ReadByte()
        if erer != nil {
            break
        }
        bafer.WriteByte(biyt)
        if biyt == delim[1] {
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
    if escape != noEscape {
        tekst, erer := verbatim(reeder, escapeDelimiters[escape])
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
        if escape != noEscape {
            tekst, erer := verbatim(reeder, escapeDelimiters[escape])
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
        if escape != noEscape {
            tekst, erer := verbatim(reeder, escapeDelimiters[escape])
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
