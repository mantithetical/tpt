# tpt-cli

`tpt cli` is written using the [cobra](https://github.com/spf13/cobra) library.
It currently supports the `detect-dialect` command only. `detect-dialect` naively
identifies if TpT resources contains American words, British words, both or 
none. 

Notes: 
* `NaiveDetector` tokenizes and singularizes input, drops non-letters and 
and performs a case-insensitive string-comparison to the corpora.
* The program does not account for tenses meaning `analyzed` would be ignored 
even though it is an American spelling.
* It also does not try to capture intent in any way meaning spelling mistakes
 and alternate meanings of words are ignored. 
* Data from the GraphQL API is fetched in batches with default being 10. This
 behavior can be modified with the `-s` flag

```bash
>tpt detect-dialect
135502   British English                     humour
536841   Unknown                             
723439   American English                    color
1224072  American English                    center
2723439  Mixed British and American English  license humour
2939135  Unknown        
```

#### Usage

```bash
tpt -h
tpt cli

Usage:
  tpt [command]

Available Commands:
  detect-dialect Identifies dialect of English in TpT resources
  help           Help about any command

Flags:
      --config string   config file (default is $HOME/.tpt.yaml)
  -h, --help            help for tpt
  -v, --verbose         verbose output

Use "tpt [command] --help" for more information about a command.
```

```bash
>tpt detect-dialect -h
TpT has many users from US, Canada, Australia and New Zealand.
Users from these countries may be more likely to buy resources
localized (localised?) to their own version of English.
`detect-dialect` identifies the dialect of English used in resources
based on their title and description.

Usage:
  tpt detect-dialect [flags]

Flags:
  -a, --american-words-gist string   gist containing American words (default "see config file")
  -s, --batch-size int8              number of product ids fetched at a time (default 10)
  -b, --british-words-gist string    gist containing British words (default "see config file")
  -h, --help                         help for detect-dialect
  -p, --product-ids string           file containing product ids (default "see config file")

Global Flags:
      --config string   config file (default is $HOME/.tpt.yaml)
  -v, --verbose         verbose output

```

#### Install

```bash
# Mac OSs only 
>sh resources/install.sh 
```

#### Uninstall

```bash
# Mac OSs only 
>sh resources/uninstall.sh
```

#### Auto-completion

tpt cli supports tab-completion 
