## Go port of the popular Python Jellyfish library 
[![Build Status](https://travis-ci.com/UjjwalAyyangar/go-jellyfish.svg?token=TQXyyNBwokx5jjkwmmpH&branch=master)](https://travis-ci.com/UjjwalAyyangar/go-jellyfish) [![Go Report Card](https://goreportcard.com/badge/github.com/UjjwalAyyangar/go-jellyfish)](https://goreportcard.com/report/github.com/UjjwalAyyangar/go-jellyfish)

Written by James Turk and Michael Stephens in Python

It is a library for doing approximate and phonetic matching of strings.

Python source is available at http://github.com/jamesturk/jellyfish.

**Will be adding tests for this version soon.**

## Included Algorithms

### String comparison:

* Levenshtein Distance
* Damerau-Levenshtein Distance
* Jaro Distance
* Jaro-Winkler Distance
* Match Rating Approach Comparison
* Hamming Distance

### Phonetic encoding:

* American Soundex
* Metaphone
* NYSIIS (New York State Identification and Intelligence System) (Requires Testing)
* Match Rating Codex (Requires Testing)

### Example Usage

```
    jellyfish.Levenshtein_distance("jellyfish","smellyfish")
    jellyfish.Match_rating_comparision("Jellyfish","Smellyfish")
```

refer to [example.go](https://github.com/UjjwalAyyangar/go-jellyfish/blob/master/examples/example.go)
