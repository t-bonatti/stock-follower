package statusinvest

import (
	"github.com/jmoiron/sqlx"
)

//Repository repository interface
type Repository interface {
	Create(StatusInvest) error
}

type repo struct {
	db *sqlx.DB
}

//NewStatusInvestRepository Return a new repository
func NewStatusInvestRepository(db *sqlx.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(entity StatusInvest) (err error) {
	_, err = r.db.Exec("INSERT INTO statusinvest(name,sector,sub_sector,segment,value,variation_last_12_month,min_in_last_12_month,max_in_last_12_month,payment_last_12_month,dividend_yield,variation_in_month,min_in_month,max_in_month,pvp,pl,pebitda,pebit,pativo,evebitda,evebit,psr,p_working_capital,p_current_assets_net,gross_margin,ebitda_margin,ebit_margin,net_margin,turn_asset,roe,roa,roic,lpa,vpa,net_debt_patrimony,net_debt_ebitda,net_debt_ebit,patrimony_assets,liabilities_assets,current_liquidity,cagr_revenue,cagr_profit,crawled_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42)",
		entity.Name,
		entity.Sector,
		entity.SubSector,
		entity.Segment,
		entity.Value,
		entity.VariationLast12Month,
		entity.MinInLast12Month,
		entity.MaxInLast12Month,
		entity.PaymentLast12Month,
		entity.DividendYield,
		entity.VariationInMonth,
		entity.MinInMonth,
		entity.MaxInMonth,
		entity.PVP,
		entity.PL,
		entity.PEbitda,
		entity.PEbit,
		entity.PAtivo,
		entity.EVEbitda,
		entity.EVEbit,
		entity.PSR,
		entity.PWorkingCapital,
		entity.PCurrentAssetsNet,
		entity.GrossMargin,
		entity.EbitdaMargin,
		entity.EbitMargin,
		entity.NetMargin,
		entity.TurnAsset,
		entity.ROE,
		entity.ROA,
		entity.ROIC,
		entity.LPA,
		entity.VPA,
		entity.NetDebtPatrimony,
		entity.NetDebtEbitda,
		entity.NetDebtEbit,
		entity.PatrimonyAssets,
		entity.LiabilitiesAssets,
		entity.CurrentLiquidity,
		entity.CAGRRevenue,
		entity.CAGRProfit,
		entity.CrawledAt,
	)
	return
}
