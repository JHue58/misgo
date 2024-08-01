package money

import (
	"github.com/jhue/misgo/biz"
	"regexp"
	"strconv"
	"strings"
)

//var (
//	separators      = []string{",", "，", "。", "\n"}
//	expenseKeywords = []string{"用了", "花了", "支出","开销"}
//	incomeKeywords  = []string{"收入", "得到"}
//)

type TransactionParser struct {
	amountRe *regexp.Regexp
}

var DefaultTransactionParser = NewTransactionParser()

func NewTransactionParser() *TransactionParser {
	amountRe := regexp.MustCompile(`(\d+(?:\.\d+)?)\s*(?:元|块钱)?`)

	return &TransactionParser{
		amountRe: amountRe,
	}
}

func (p *TransactionParser) Parse(content string) (result []*Transaction) {
	config := biz.GetBizConfig().ParserConfig

	var splitContent []string
	for _, sep := range config.Separators {
		splitContent = strings.Split(content, sep)
		content = ""
		for _, part := range splitContent {
			content += part + " "
		}
		splitContent = strings.Fields(content)
	}

	// 遍历分割后的字串

	for _, part := range splitContent {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		result = append(result, p.parseOne(part))
	}

	// 返回结果
	return

}

func (p *TransactionParser) parseOne(part string) (transaction *Transaction) {
	config := biz.GetBizConfig().ParserConfig

	transaction = &Transaction{}
	// ① 解析金额
	amountMatch := p.amountRe.FindStringSubmatch(part)
	if len(amountMatch) > 0 {
		amount, _ := strconv.ParseFloat(amountMatch[1], 64)
		transaction.Amount = float32(amount)
	}

	// ② 查找支出或收入
	transactionType := "支出"
	for _, keyword := range config.IncomeKeywords {
		if strings.Contains(part, keyword) {
			transactionType = "收入"
			part = strings.Replace(part, keyword, "", -1)
		}
	}
	for _, keyword := range config.ExpenseKeywords {
		if strings.Contains(part, keyword) {
			transactionType = "支出"
			part = strings.Replace(part, keyword, "", -1)
		}
	}
	transaction.Type = transactionType

	// ③ 移除已解析的内容
	part = p.amountRe.ReplaceAllString(part, "")
	part = strings.Replace(part, transactionType, "", -1)
	part = strings.TrimSpace(part)

	// ④ 处理剩余的字符串
	category := part
	if len(strings.Split(part, " ")) > 1 {
		category = strings.Split(part, " ")[0]
	}

	transaction.Category = category

	return
}
