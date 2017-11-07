package cmd

import (
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "github.com/mantithetical/tpt/dialect"
  "fmt"
)

var AmericanWordsGist string
var BritishWordsGist string
var ProductIdsFile string
var BatchSize int8

func init() {
  TptCmd.AddCommand(detectDialectCmd)

  // Local Flags
  detectDialectCmd.Flags().StringVarP(&AmericanWordsGist,
    dialect.AmericanWordsGist, "a",
    "see config file", "gist containing American words")
  viper.BindPFlag(dialect.AmericanWordsGist,
    detectDialectCmd.Flag(dialect.AmericanWordsGist))

  detectDialectCmd.Flags().StringVarP(&BritishWordsGist,
    dialect.BritishWordsGist, "b",
    "see config file", "gist containing British words")
  viper.BindPFlag(dialect.BritishWordsGist,
    detectDialectCmd.Flag(dialect.BritishWordsGist))

  detectDialectCmd.Flags().StringVarP(&ProductIdsFile, dialect.ProductIds, "p",
    "see config file", "file containing product ids")
  viper.BindPFlag(dialect.ProductIds,
    detectDialectCmd.Flag(dialect.ProductIds))

  detectDialectCmd.Flags().Int8VarP(&BatchSize, dialect.BatchSize, "s",
    10, "number of product ids fetched at a time")
  viper.BindPFlag(dialect.BatchSize,
    detectDialectCmd.Flag(dialect.BatchSize))
}

var detectDialectCmd = &cobra.Command{
  Use:   "detect-dialect",
  Short: "Identifies dialect of English in TpT resources",
  Long: "TpT has many users from US, Canada, Australia and New Zealand.\n" +
    "Users from these countries may be more likely to buy resources\n" +
    "localized (localised?) to their own version of English.\n" +
    "`detect-dialect` identifies the dialect of English used in resources\n" +
    "based on their title and description.",

  Run: func(cmd *cobra.Command, args []string) {
    if Verbose {
      fmt.Printf("%s: %s\n", dialect.AmericanWordsGist,
        viper.GetString(dialect.AmericanWordsGist))
      fmt.Printf("%s: %s\n", dialect.BritishWordsGist,
        viper.GetString(dialect.BritishWordsGist))
      fmt.Printf("%s: %s\n\n", dialect.ProductIds,
        viper.GetString(dialect.ProductIds))
    }
    dialect.Detect(cmd)
  },
}
