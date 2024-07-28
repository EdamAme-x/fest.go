package lexer

import "strings"

type TokenType int

type Token struct {
	Type     TokenType
	Value    string
	Children []Token
}

func Lexer(code string) []Token {
	splittedCode := strings.Split(code, "\n")

	tokenizedTokens := []Token{} // からの配列を宣言するときには varを用いたほうが良さげだと思うが宗教に任せる
	for i := 0; i < len(splittedCode); i++ {
		tokenizedTokens = append(tokenizedTokens, Tokenize(splittedCode[i]))
	}

	return tokenizedTokens
}

func Tokenize(code string) Token {
	code = strings.TrimSpace(code)
	
	if strings.HasPrefix(code, "#") {
		return Token{
			Type:  LABEL,
			Value: "comment",
		}
	}

	return Token{}
}

const (
	LABEL TokenType = iota
	NAME
	STRUCT
	LINE
	VALUE
)

/* 	  MEMO
MAMIMUMEMO

Union型とかそういうのやるんだったらGenericを用いればいいのでは
@EdamAme-x
*/