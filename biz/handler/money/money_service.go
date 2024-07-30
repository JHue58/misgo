// Code generated by hertz generator.

package money

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jhue/misgo/biz"
	"github.com/jhue/misgo/biz/model/base"
	money "github.com/jhue/misgo/biz/model/money"
	"github.com/jhue/misgo/biz/request"
	"github.com/jhue/misgo/db"
	moneyM "github.com/jhue/misgo/db/model/money"
	"github.com/jhue/misgo/db/model/user"
	"github.com/jhue/misgo/internal/mislog"
	"strings"
	"time"
)

// TransactionPut .
// @router api/money [PUT]
func TransactionPut(ctx context.Context, c *app.RequestContext) {
	bizCtx := request.NewBizContext(c)
	var err error
	var req money.TransactionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		bizCtx.ParmaError(err)
		return
	}
	u, ok := user.ExtractUser(ctx)
	if !ok {
		bizCtx.ParmaError(base.UIDError)
		return
	}
	if req.OneTransaction == nil {
		bizCtx.ParmaError(money.ErrTransactionError)
		return
	}
	req.OneTransaction.Filter()
	err = req.OneTransaction.Validation()
	if err != nil {
		bizCtx.ParmaError(err)
		return
	}

	d := db.Get()
	entry := moneyM.Transaction{
		UserID:   u.ID,
		Type:     req.OneTransaction.Type,
		Category: req.OneTransaction.Category,
		Amount:   float64(req.OneTransaction.Amount),
		Note:     req.OneTransaction.Note,
		Time:     time.Unix(req.OneTransaction.Time, 0),
	}
	d.Create(&entry)

	bizCtx.Success()
	mislog.DefaultLogger.Infof("Money Success [Name] %s \n", u.Name)
}

// TransactionGet .
// @router api/money [POST]
func TransactionGet(ctx context.Context, c *app.RequestContext) {
	bizCtx := request.NewBizContext(c)
	var err error
	var req money.TransactionGetReq
	err = c.BindAndValidate(&req)
	if err != nil {
		bizCtx.ParmaError(err)
		return
	}
	u, ok := user.ExtractUser(ctx)
	if !ok {
		bizCtx.ParmaError(base.UIDError)
		return
	}
	config := biz.GetBizConfig().MoneyConfig

	d := db.Get()
	if req.Count > config.MaxGetCount {
		req.Count = config.MaxGetCount
	}
	startDate, endDate, dateFormat := req.GetDateRange()
	if startDate.After(endDate) {
		bizCtx.ParmaError(money.ErrStartTimeBeforeEndTime)
		return
	}
	b := make([]*moneyM.Transaction, 0, req.Count)
	var query string
	if req.Condition != "" {
		query = fmt.Sprintf("user_id = ? AND %s AND (time >= ? AND time <= ?)", req.Condition)
	} else {
		query = fmt.Sprintf("user_id = ? AND (time >= ? AND time <= ?)")
	}

	var count int64
	var maxTransaction moneyM.Transaction
	res := d.Model(&maxTransaction).Where(query, u.ID, startDate, endDate).Order("amount desc").Count(&count).Limit(1).Find(&maxTransaction)
	if res.Error != nil {
		bizCtx.DBError(res.Error)
		return
	}

	res = d.Where(query, u.ID, startDate, endDate).Order(req.Order).Offset(int(req.Start)).Limit(int(req.Count)).Find(&b)
	if res.Error != nil {
		bizCtx.DBError(res.Error)
		return
	}
	if len(b) == 0 && count > 0 {
		bizCtx.SuccessWithMsg("finished")
		return
	}
	mdTitle := fmt.Sprintf("### %s\n#### 总共有`%d`笔交易,最大笔的交易为\n`%s` \n", dateFormat, count, maxTransaction.String())

	resp := money.TransactionResp{
		Count:        count,
		ReportMD:     ToReportMarkDown(mdTitle, b),
		Transactions: nil,
	}
	bizCtx.Response(&resp)
	mislog.DefaultLogger.Infof("MoneyGet Success [Name] %s [Returns] %d\n", u.Name, len(b))
}

func ToReportMarkDown(title string, b []*moneyM.Transaction) string {
	builder := strings.Builder{}
	builder.WriteString(title)
	builder.WriteString("|日期|类型|用途/来源|金额|备注| \n |-------|-------|-------|-------|-------|\n")
	for _, transaction := range b {
		builder.WriteString(transaction.Markdown())
	}
	return builder.String()

}