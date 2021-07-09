package lexer

import (
	"ckript/token"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func contains(str string, needle string) bool {
	return strings.Contains(str, needle)
}

func containsRune(str string, char rune) bool {
	return strings.ContainsRune(str, char)
}

func isWhitespace(char rune) bool {
	return unicode.IsSpace(char)
}

func validNumer(str string, base int) bool {
	var _, err = strconv.ParseInt(str, base, 64)
	return err == nil
}

func validFloat(str string) bool {
	var _, err = strconv.ParseFloat(str, 64)
	return err == nil
}

var buildinTypes = [...]string{
	"int", "double", "func", "str", "void", "obj", "arr", "bool",
}

var regexActual = [...]string{
	"\n", "\t", "\a", "\r", "\b", "\v",
}

var regexes = []string{
	`\\n`, `\\t`, `\\a`, `\\r`, `\\b`, `\\v`,
}

func Unescape(str *string) {
	for idx, reg := range regexes {
		*str = strings.ReplaceAll(*str, reg, regexActual[idx])
	}
}

type Lexer struct {
	verbose           bool
	currentLine       int
	deletedSpaces     int
	prevDeletedSpaces int
	ptr               int
	end               int
	code              string
	fileDir           string
	fileName          string
	tokens            []*token.Token
}

func (lexer *Lexer) at(pos int) rune {
	return []rune(lexer.code)[pos]
}

func (lexer *Lexer) cur() rune {
	return lexer.at(lexer.ptr)
}

func (lexer *Lexer) consumeWhitespace() {
	for lexer.ptr != lexer.end && isWhitespace(lexer.cur()) {
		lexer.deletedSpaces++
		if lexer.cur() == '\n' {
			lexer.currentLine++
		}
		lexer.ptr++
	}
}

func (lexer *Lexer) addToken(_type token.TokenType, value string) {
	if lexer.verbose {
		fmt.Printf("Adding token [%v] = '%v'", _type, value)
	}
	lexer.tokens = append(lexer.tokens, token.New(_type, value, 0, ""))
}

const (
	chars  = ".,:;{}[]()~$#"
	chars2 = "=+-*&|/<>!%^"
)

var allowedTokenKeys = [...]string{
	"function", "class", "array", "return", "if", "else", "for", "while",
	"break", "continue", "alloc", "del", "ref", "true", "false", "const",
}

func (lexer *Lexer) Tokenize(code string) []*token.Token {
	lexer.ptr = 0
	lexer.end = len(lexer.code)
	lexer.code = code
	for lexer.ptr != lexer.end {
		lexer.deletedSpaces = 0
		lexer.consumeWhitespace()
		if lexer.ptr == lexer.end {
			break
		}
		var c = lexer.cur()
		if containsRune(chars, c) {
			lexer.addToken(token.TokenType(c), "")
		} else {
			continue
		}
	}
	return lexer.tokens
}

func (lexer *Lexer) ProcessFile(filename string) ([]*token.Token, bool) {
	return make([]*token.Token, 0), false
}
