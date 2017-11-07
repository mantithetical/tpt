package dialect

import (
  "net/http"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "fmt"
  "os"
  "io"
  "bufio"
  "strings"
  "bytes"
  "encoding/json"
  "io/ioutil"
)

const (
  AmericanWordsGist  = "american-words-gist"
  BritishWordsGist   = "british-words-gist"
  ProductIds         = "product-ids"
  BatchSize          = "batch-size"
  TptGraphQlResource = "tpt-graphql-resource"
)

type Corpus struct {
  AmericanWords map[string]bool
  BritishWords  map[string]bool
}

type Product struct {
  Description string
  Id          string
  Name        string
}

var queryStr = `{
    "query": "query productText($productIds: [ID]!) {products(ids: $productIds) {id name description}}",
      "variables": { "productIds": [$input] }
  }`

var corpus = Corpus{make(map[string]bool), make(map[string]bool)}
var detector = NaiveDetector{corpus}

func Detect(cmd *cobra.Command) {
  setupCorpus(cmd)

  if !cmd.Flag(ProductIds).Changed {
    flag(getProducts(viper.GetStringSlice(ProductIds)), detector)
  } else {
    file, err := os.Open(viper.GetString(ProductIds))
    panic(err)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      flag(
        getProducts(
          batch(viper.GetInt(BatchSize), scanner)),
        detector)
    }
  }
}

func flag(products []Product, detector Detector) {
  for _, p := range products {
    fmt.Printf("%-8s %s\n", p.Id, detector.Flag(p))
  }
}

func getProducts(ids []string) []Product {
  q := strings.Replace(queryStr, "$input", strings.Join(ids, ","), 1)
  if viper.GetBool("verbose") {
    fmt.Printf("%s\n\n", q)
  }
  resp, err := http.Post(viper.GetString(TptGraphQlResource), "application/json",
    bytes.NewBuffer([]byte(q)))
  panic(err)

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  panic(err)

  data := map[string]map[string][]Product{"data":
    {"products": make([]Product, 0)}}
  panic(json.Unmarshal(body, &data))
  return data["data"]["products"]
}

func batch(size int, scanner *bufio.Scanner) []string {
  productIds := make([]string, 0)
  productIds = append(productIds, scanner.Text())
  for i := 1; i < size; i++ {
    if scanner.Scan() {
      productIds = append(productIds, scanner.Text())
    }
  }
  panic(scanner.Err())

  if viper.GetBool("verbose") {
    fmt.Printf("\nProduct Ids: %v\n\n", productIds)
  }
  return productIds
}

func setupCorpus(cmd *cobra.Command) {
  loadCorpus(
    fetchCorpus(viper.GetString(AmericanWordsGist), corpus.AmericanWords))
  loadCorpus(
    fetchCorpus(viper.GetString(BritishWordsGist), corpus.BritishWords))
}

// fetch a given raw gist url
func fetchCorpus(gist string, m map[string]bool) (io.ReadCloser,
  map[string]bool) {

  resp, err := http.Get(gist)
  panic(err)
  return resp.Body, m
}

// make corpus from reader
func loadCorpus(reader io.ReadCloser, m map[string]bool) {
  defer reader.Close()
  scanner := bufio.NewScanner(reader)
  for scanner.Scan() {
    m[scanner.Text()] = true
  }
  panic(scanner.Err())
}

func panic(err error) {
  if err != nil {
    fmt.Printf("fetch: %v\n", err)
    os.Exit(1)
  }
}
