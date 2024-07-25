package lexer

/*
top-level:
    if we encounter an <at-keyword> -> parse top-level-at-rule
    else if we encounter an <identifier> || <colon> || <square-bracket-open> || <hash> -> parse top-level-selectors until we reach <curly-bracket-open>
    else -> parse error

top-level-selectors:
    always consume until we encounter a <comma> or <curly-bracket-open>

    if first code point is <alpha> -> parse element selector
        if we encounter any non <alpha> -> parse combinator selector
    else if first code point is <period> -> parse class selector
        if we encounter any non <alpha> or <digit> or <hyphen> or <underscore> -> parse combinator selector
    else if first code point is <hash> -> parse id selector
        if we encounter any non <alpha> or <digit> or <hyphen> or <underscore> -> parse combinator selector
    else if first code point is <colon> and second code point is <alpha> -> parse pseudo class selector
        if we encounter any non <alpha> or <hyphen> -> parse combinator selector
    else if first code point is <colon> and second code point is <colon> -> parse pseudo element selector
        if we encounter any non <alpha> or <hyphen> -> parse combinator selector
    else if first code point is <square-bracket-open> -> parse attribute selector until we reach <square-bracket-close>
        if next code point is not <comma> or <curly-bracket-open> -> parse combinator selector
    else -> parse error

top-level-at-rule:
    consume <at>
    if next code point is <alpha> -> parse at-keyword
    else -> parser error

at-keyword:
    consume until no more <alpha> || <hyphen>

    switch on <at-keyword>
    "charset" -> parse at-rule-charset
    "import" -> parse at-rule-import
    "layer" -> parse at-rule-layer
    "namespace" -> parse at-rule-namespace
    else -> parse error
*/
