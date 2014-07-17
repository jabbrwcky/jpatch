jsonmerge
=========

Simple utility for merging two JSON files, the first being the source and the second being a 'patch' file.

Usage
-----

```
jsonmerge [-f] [-p] [-o <output file>] <source> <patch>
```

if the output file is not specified 

Known limitations
-----------------

Arrays are not merged right now but are completely replaced.

