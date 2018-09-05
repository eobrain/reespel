# reespel -- Convert English Text to Phonetic Spelling

This command-line utility converts English text to phonetic spelling.

Usage:

```sh
./reespel <input.txt >output.txt
```

It uses the phonetic information in a dictionary to covert words
into phonemes, and then maps them to a spelling designed to be
intuitive to English readers.

To build and try out this you will need the Go language and the `make`
utility.  After cloning this repo, simply type `make` to build the
`reespel` executable and run it on some examples.

The build actually happens in two stages.

1.   A code generator reads the [The CMU Pronouncing Dictionary][1]
     and wrties a Go-language source file containing an in-momory map
     from English words with the CMY phoneme codes replaced by a new
     phonetic spelling.
 2.  Then the reespel program is compiled using the generated code and
     a main program that filters input text.



[1]: http://www.speech.cs.cmu.edu/cgi-bin/cmudict
