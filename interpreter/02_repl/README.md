# Monkey Lexer

The lexer'sjob is not to tell us whether code makes sense.
1. Cover tokens.
2. Provoke off-by-one errors.
3. Edge cases at end-of-line.
4. newline handling.
5. multi-difit handling.

* It does not support unicode - see `lexer.go:isLetter()`.
* It does not support floats, hex, octal, etc. - see `lexer.go:isDigit()`.
