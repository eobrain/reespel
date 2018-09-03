# reespel -- Convert English Text to Phonetic Spelling

This command-line utility converts English text to phonetic spelling.

Usage:

```sh
./reespel <input.txt >output.txt
```

It uses the phonetic information in the cmudict file to covert words
into phonemes, and then maps them to a spelling designed to be
intuitive to English readers. The cmudict file must be in the current
directory, or it must be specified with a commant-line flag.

To build and try out this you will need the Go language and the `make`
utility.  After cloning this repo, simply type `make` to build the
`reespel` executable and run it on some examples.
