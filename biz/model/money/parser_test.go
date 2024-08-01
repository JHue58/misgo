package money

import (
	"fmt"
	"github.com/jhue/misgo/biz"
	"strings"
	"testing"
)

func TestSplitString(t *testing.T) {
	separators := []string{
		"，",
		"。",
		"，",
		"。",
		"\n",
	}
	content := `
这是一个字符串,还有我,以及你。嗯嗯，ok
啊？
`

	var splitContent []string
	for _, sep := range separators {
		splitContent = strings.Split(content, sep)
		content = ""
		for _, part := range splitContent {
			content += part + " "
		}
		splitContent = strings.Fields(content)
	}
	for _, s := range splitContent {
		fmt.Println(s)
	}

}

func initConf() {
	biz.SetBizConfig(biz.Config{
		MoneyConfig: biz.MoneyConfig{
			ParserConfig: biz.ParserConfig{
				Separators:      []string{",", "，", "。", "\n"},
				ExpenseKeywords: []string{"用了", "花了", "支出", "开销", "花费"},
				IncomeKeywords:  []string{"收入", "得到"},
			},
		},
	})
}

func TestTransactionParser_Parse(t *testing.T) {
	initConf()
	p := NewTransactionParser()
	content := `
	100元，吃饭用了90，健身开销100，吃饭90.66
吃饭99
`
	transactions := p.Parse(content)
	for _, transaction := range transactions {
		fmt.Println(transaction.String())
	}

}

func BenchmarkTransactionParser_Parse(b *testing.B) {
	initConf()
	p := NewTransactionParser()
	content := `
	100元，吃饭用了90，健身开销100，吃饭90
吃饭99
`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(content)
	}
}
