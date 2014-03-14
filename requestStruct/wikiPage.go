/**
 * provides functions to process requests to given urls
 * some example:
 * - https://gist.github.com/border/775526
 * - https://stackoverflow.com/questions/16931499/in-go-language-how-do-i-unmarshal-json-to-array-of-object
 * - https://stackoverflow.com/questions/13593519/how-do-i-parse-an-inner-field-in-a-nested-json-object-in-golang
 * - https://stackoverflow.com/questions/17209111/unable-to-parse-a-complex-json-in-golang
 * - http://play.golang.org/p/AEC_TyXE3B
 * - http://play.golang.org/p/TFUgJsWNhq
 * - https://stackoverflow.com/questions/19482612/go-golang-array-type-inside-struct-missing-type-composite-literal
 * - https://gobyexample.com/json
 * - http://mattyjwilliams.blogspot.co.uk/2013/01/using-go-to-unmarshal-json-lists-with.html
 */

package requestStruct

import (
    "github.com/russross/blackfriday"
    // "bytes"
    "log"
    "fmt"
    "strings"
    "regexp"
)

import (
)

type WikiPage struct {
    Query SubType `json:"query"`
}

type SubType struct {
    Pages map[string]ActualPage `json:"pages"`
}

type ActualPage struct {
    PageId int  `json:"pageid"`
    Title string `json:"title"`
    Rev []Revision `json:"revisions"`
}

type Revision struct {
    ContentFormat string `json:"contentformat"`
    ContentModel string `json:"contentmodel"`
    RawContent string `json:"*"`
}

func GetWikiRawText(text string) string {

    // latexRenderer := new(blackfriday.Latex)
    // textByte := []byte(text)

    // var byteBuffer []byte
    // var byteBuffer bytes.Buffer

    // byteBuffer := bytes.Buffer("")

    // out *bytes.Buffer, text []byte)
    // latexRenderer.NormalText(&byteBuffer, textByte)

    // escapeItems := []string {
    //     "[", "]",
    //     "=",
    //     "<", ">",
    //     "{", "}",
    //     "*",
    //     "'",
    // }

    escapeItems := []byte("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

    for i := range escapeItems {
        text = strings.Replace(text, string(escapeItems[i]), "", -1)
    }

    // text = string(stripTagCustom(text, "[", "]", "", false)[:])
    //
    //
    // text = string(stripTagCustom(text, "]", "")[:])
    // text = string(stripTagCustom(text, "<!--", "")[:])


    // dassa := byteBuffer.String()

    // output := blackfriday.MarkdownBasic([]byte(text))
    return text
}

// was used to strip not needed stuff from a text block / snipped
func GetWikiRawTextRegexpr(text string) string {

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

    // re := regexp.MustCompile("<(["+punct+"]+)>")

    re := regexp.MustCompile(regexString)

    return re.ReplaceAllString(text, " ")

}

func GetWikiRawTextOld(text string) string {

    htmlFlags := 0
    htmlFlags |= blackfriday.HTML_OMIT_CONTENTS
    htmlFlags |= blackfriday.HTML_SKIP_IMAGES
    htmlFlags |= blackfriday.HTML_SKIP_HTML
    htmlFlags |= blackfriday.HTML_SKIP_SCRIPT
    htmlFlags |= blackfriday.HTML_SKIP_LINKS
    renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

    // set up the parser
    extensions := 0

    output := blackfriday.Markdown([]byte(text), renderer, extensions)
    das := string(output[:])
    return das
}


// func stripTag(text, tag, newTag string) []byte {
//     closeNewTag := fmt.Sprintf("</%s>", newTag)
//     i := 0
//     for i < len(text) && text[i] != '<' {
//         i++
//     }
//     if i == len(text) {
//         return []byte(text)
//     }
//     found, end := findHtmlTagPos([]byte(text[i:]), tag)
//     closeTag := fmt.Sprintf("</%s>", tag)
//     noOpen := text
//     if found {
//         log.Printf("FOUND SOMETHING")
//         noOpen = text[0:i+1] + newTag + text[end:]
//     }
//         log.Printf("FOUND NOTHING")
//     return []byte(strings.Replace(noOpen, closeTag, closeNewTag, -1))
// }

func stripTagCustom(text string, tagBeginning string, tagClosing string, newTag string, removeContent bool) []byte {
    closeNewTag := fmt.Sprintf("%s", newTag)
    i := 0
    // for i < len(text) && text[i] != '<' {
    //     i++
    // }
    if i == len(text) {
        return []byte(text)
    }
    found, end := findHtmlTagPos([]byte(text[i:]), tagBeginning)
    closeTag := fmt.Sprintf("%s", tagClosing)
    newText := text
    if found {
        log.Printf("FOUND SOMETHING")
        // we will take the beginning of the text till the beginning tag and then the rest
        newText = text[0:i+1] + newTag + text[end:]
    }
        log.Printf("FOUND NOTHING")
    return []byte(strings.Replace(newText, closeTag, closeNewTag, -1))
}


func findHtmlTagPos(tag []byte, tagname string) (bool, int) {
    i := 0
    if i < len(tag) && tag[0] != '<' {
        return false, -1
    }
    i++
    i = skipSpace(tag, i)

    if i < len(tag) && tag[i] == '/' {
        i++
    }

    i = skipSpace(tag, i)
    j := 0
    for ; i < len(tag); i, j = i+1, j+1 {
        if j >= len(tagname) {
            break
        }

        if strings.ToLower(string(tag[i]))[0] != tagname[j] {
            return false, -1
        }
    }

    if i == len(tag) {
        return false, -1
    }

    // Now look for closing '>', but ignore it when it's in any kind of quotes,
    // it might be JavaScript
    inSingleQuote := false
    inDoubleQuote := false
    inGraveQuote := false
    for i < len(tag) {
        switch {
        case tag[i] == '>' && !inSingleQuote && !inDoubleQuote && !inGraveQuote:
            return true, i
        case tag[i] == '\'':
            inSingleQuote = !inSingleQuote
        case tag[i] == '"':
            inDoubleQuote = !inDoubleQuote
        case tag[i] == '`':
            inGraveQuote = !inGraveQuote
        }
        i++
    }

    return false, -1
}

func skipSpace(tag []byte, i int) int {
    for i < len(tag) && isspace(tag[i]) {
        i++
    }
    return i
}

func isspace(c byte) bool {
    return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' || c == '\v'
}
