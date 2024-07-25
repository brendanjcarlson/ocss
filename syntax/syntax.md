## Definitions and symbols
```
+ (space-separated)
# (comma-separated)
| (or-condition)
```

## At-rules
### @charset
```
@charset <"charset">;
```



### @color-profile
```
@color-profile [<dashed-indent> | device-cmyk] { <declaration-list> }
```
#### Descriptors
- src
- rendering-intent



### @container
```
@container <container-condition> { <stylesheet> }
```



### @counter-style
```
@counter-style <counter-style-name> { <declaration-list> }

where
    <counter-style-name> ::= <custom-ident>
```
#### Descriptors
- system
- symbols
- additive-symbols
- negative
- prefix
- suffix
- range
- pad
- speak-as
- fallback

### @document DEPRECATED NON-STANDARD



### @font-face
```
@font-face { <declaration-list> }
```
#### Descriptors
- ascent-override
- descent-override
- font-display
- font-family
- font-stretch
- font-style
- font-weight
- font-feature-settings
- font-variation-settings
- line-gap-override
- size-adjust
- src
- unicode-range



### @import
```
@import [ <url> | <string> ] [ layer | layer( <layer-name ) ]? <import-conditions> ;

where
    <url> ::= <url()> | <src()>
    <layer-name> ::= <ident> [ ',' <ident> ]*
    <import-conditions> ::= [ supports( [ <supports-condition> | <declaration> ] ) ]} <media-query-list>?
    <url()> ::= url( <string> <url-modifier>* ) | <url-token>
    <src()> ::= src( <string> <url-modifier>* )
    <supports-condition> ::= not <supports-in-parens> | <supports-in-parens> [ and <supports-in-parens> ]* | <supports-in-parens> [ or <supports-in-parens ]*
    <supports-in-parents> ::= ( <supports-condition ) | <supports-feature> | <general-enclosed>
    <supports-feature> ::= <supports-decl>
    <general-enclosed> ::= [ <function-token> <any-value>? ) ] | [ ( <any-value>? ) ]
    <supports-decl> ::= ( <declaration> )
```



### @keyframes
```
@keyframes <keyframes-name> { <qualified-rule-list> }

where
    <keyframes-name> ::= <custom-ident> | <string>
    <qualified-rule> ::= from | to | ?<timeline-range-name> <percentage> { <declaration-list> }
```



### @layer
```
@layer <layer-name> [ ',' <layer-name> ]*
@layer <layer-name> { <rule-list> }

where 
    <layer-name> ::= <ident>
```



### @media
```
@media <media-query-list> { <rule-list> }
```



### @namespace
```
@namespace <namespace-prefix>? [ <string | <url ] ;

where
    <namespace-prefix> ::= <ident>
    <url> ::= <url()> | <src()>
    <url()> ::= url( <string> <url-modifier>* ) | <url-token>
    <src()> ::= src( <string> <url-modifier>* )
```



### @page
```
@page <page-selector-list>? { <declaration-rule-list> }

where
    <page-selector-list> ::= <page-selector>#
    <page-selector> ::= [ <ident-token>? <pseudo-page>* ]!
    <pseudo-page> ::= ':' [ left | right | first | blank ]
```

#### Margin at-rules
```
@top-left-corner
@top-left
@top-center
@top-right
@top-right-corner
@bottom-left-corner
@bottom-left
@bottom-center
@bottom-right
@bottom-right-corner
@left-top
@left-middle
@left-bottom
@right-top
@right-middle
@right-bottom
```



### @property
```
@property <custom-property-name> { <declaration-list> }
```

#### Descriptors
- syntax
- inherits
- initial-value
