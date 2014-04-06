/**
 * provides functions to create wording context relations
 */

package util

import (
    "log"
    "strings"
    "regexp"
    "bytes"
    "html/template"
    "syscall"
)

/**
 * will actually do the http get request
 */
func GetSnippetsFromText(text string) []string {
    re := regexp.MustCompile("\n|\r")
    snippets := re.Split(text, -1)

    return snippets
}

// could be removed?
func RemovePunctationFromText(text string) string {
    punct := `[!//#$%&()*+,-./:;<=>?@[]^_{|}~]`
    // punct := `[ ! // # $ % & ( ) * + ,  - . / : ; < = > ? @ [ ] ^ _ { | } ~ ]`

    regexString := "\\s*"+            // discard possible leading whitespace
          "("+               // start capture group #1
            "\\.{3}"+            // ellipsis (must appear before punct)
          "|"+               // alternator
            "\\s+\\-\\s+"+       // hyphenated words (must appear before punct)
          "|"+               // alternator
            "\\s+\"(?:\\s+)?"+   // compound words (must appear before punct)
          "|"+               // alternator
            "\\s+"+              // other words
          "|"+               // alternator
            "["+punct+"]"+        // punct
          ")"

    re := regexp.MustCompile(regexString)
    return re.ReplaceAllString(text, " ")
}

func RemovePunctationFromTextCustom(text string) string {
    escapeItems := []byte("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

    for i := range escapeItems {
        text = strings.Replace(text, string(escapeItems[i]), "", -1)
    }
    return text
}

func CleanUpSnippets(snippets []string) []string {
    for i := range snippets {
        snippets[i] = RemoveHtmlTags(snippets[i])
        snippets[i] = TrimUrls(snippets[i])
        snippets[i] = TrimDoubleBracketsLeaveContent(snippets[i])
        snippets[i] = RemovePunctationFromTextCustom(snippets[i])
        snippets[i] = TrimWhitespaces(snippets[i])
        snippets[i] = strings.ToLower(snippets[i])
    }

    return snippets
}

func GetWordsFromSnippet(text string) []string {
    re := regexp.MustCompile(` `)
    words := re.Split(text, -1)

    return words
}

func TrimUrls(text string) string {
    re := regexp.MustCompile("((([A-Za-z]{3,9}:(?:\\/\\/)?)(?:[\\-;:&=\\+\\$,\\w]+@)?[A-Za-z0-9\\.\\-]+|(?:www\\.|[\\-;:&=\\+\\$,\\w]+@)[A-Za-z0-9\\.\\-]+)((?:\\/[\\+~%\\/\\.\\w\\-_]*)?\\??(?:[\\-\\+=&;%@\\.\\w_]*)#?(?:[\\.\\!\\/\\\\w]*))?)")
    return re.ReplaceAllString(text, "")
}

func TrimWhitespaces(text string) string {
    re := regexp.MustCompile("^[ \t]+|[ \t]+")
    return re.ReplaceAllString(text, " ")
}

func RemoveBracketsWithContent(text string) string {
    re := regexp.MustCompile("\\[.*\\]")
    return re.ReplaceAllString(text, " ")
}

func TrimDoubleBracketsLeaveContent(text string) string {
    re := regexp.MustCompile("[\\[(.)\\]]")
    return re.ReplaceAllString(text, " ")
}

func RemoveHtmlTags(s string) (output string) {
    buffer := bytes.NewBufferString("")
    insideHtmlTag := false

    // we will loop trough every char and only readd these that are not a html tag or a content of one
    for _, val := range s {
        switch val {
            case '<':
                insideHtmlTag = true
            case '>':
                insideHtmlTag = false
            default:
                if !insideHtmlTag {
                    buffer.WriteRune(val)
                }
        }
    }

    output = buffer.String()
    output = template.HTMLEscapeString(output)

    // wee need to escape the following html specific syntax
    output = strings.Replace(output, "&nbsp;", " ", -1)
    output = strings.Replace(output, "&quot;", " ", -1)
    output = strings.Replace(output, "&apos;", " ", -1)
    output = strings.Replace(output, "&#34;", " ", -1)
    output = strings.Replace(output, "&#39;", " ", -1)
    output = strings.Replace(output, "&amp; ", " ", -1)
    output = strings.Replace(output, "&amp;amp; ", " ", -1)

    return output
}

func MaximumUlimit() {
    // via https://stackoverflow.com/questions/17817204/how-to-set-ulimit-n-from-a-golang-program
    var rLimit syscall.Rlimit
    err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        log.Println("Error Getting Rlimit ", err)
    }
    log.Println(rLimit)
    rLimit.Max = 999999
    rLimit.Cur = 999999
    err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        log.Println("Error Setting Rlimit ", err)
    }
    err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
    if err != nil {
        log.Println("Error Getting Rlimit ", err)
    }
    log.Println("Rlimit Final", rLimit)
}
