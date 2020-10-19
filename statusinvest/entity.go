package statusinvest

import (
	"time"

	"github.com/shopspring/decimal"
)

//StatusInvest data
type StatusInvest struct {
	ID                   string          `json:"id"`
	Name                 string          `json:"Nome"`
	Sector               string          `json:"Sector"`
	SubSector            string          `json:"SubSector" db:"sub_sector"`
	Segment              string          `json:"Segment"`
	Value                decimal.Decimal `json:"Value"`
	VariationLast12Month decimal.Decimal `json:"VariationLast12Month" db:"variation_last_12_month"`
	MinInLast12Month     decimal.Decimal `json:"MinInLast12Month" db:"min_in_last_12_month"`
	MaxInLast12Month     decimal.Decimal `json:"MaxInLast12Month" db:"max_in_last_12_month"`
	PaymentLast12Month   decimal.Decimal `json:"PaymentLast12Month" db:"payment_last_12_month"`
	DividendYield        decimal.Decimal `json:"DividendYield" db:"dividend_yield"`
	VariationInMonth     decimal.Decimal `json:"VariationInMonth" db:"variation_in_month"`
	MinInMonth           decimal.Decimal `json:"MinInMonth" db:"min_in_month"`
	MaxInMonth           decimal.Decimal `json:"MaxInMonth" db:"max_in_month"`
	PVP                  decimal.Decimal `json:"PVP"`
	PL                   decimal.Decimal `json:"PL"`
	PEbitda              decimal.Decimal `json:"PEbitda"`
	PEbit                decimal.Decimal `json:"PEbit"`
	PAtivo               decimal.Decimal `json:"PAtivo"`
	EVEbitda             decimal.Decimal `json:"EVEbitda"`
	EVEbit               decimal.Decimal `json:"EVEbit"`
	PSR                  decimal.Decimal `json:"PSR"`
	PWorkingCapital      decimal.Decimal `json:"PWorkingCapital" db:"p_working_capital"`
	PCurrentAssetsNet    decimal.Decimal `json:"PCurrentAssetsNet" db:"p_current_assets_net"`
	GrossMargin          decimal.Decimal `json:"GrossMargin" db:"gross_margin"`
	EbitdaMargin         decimal.Decimal `json:"EbitdaMargin" db:"ebitda_margin"`
	EbitMargin           decimal.Decimal `json:"EbitMargin" db:"ebit_margin"`
	NetMargin            decimal.Decimal `json:"NetMargin" db:"net_margin"`
	TurnAsset            decimal.Decimal `json:"TurnAsset" db:"turn_asset"`
	ROE                  decimal.Decimal `json:"ROE"`
	ROA                  decimal.Decimal `json:"ROA"`
	ROIC                 decimal.Decimal `json:"ROIC"`
	LPA                  decimal.Decimal `json:"LPA"`
	VPA                  decimal.Decimal `json:"VPA"`
	NetDebtPatrimony     decimal.Decimal `json:"NetDebtPatrimony" db:"net_debt_patrimony"`
	NetDebtEbitda        decimal.Decimal `json:"NetDebtEbitda" db:"net_debt_ebitda"`
	NetDebtEbit          decimal.Decimal `json:"NetDebtEbit" db:"net_debt_ebit"`
	PatrimonyAssets      decimal.Decimal `json:"PatrimonyAssets" db:"patrimony_assets"`
	LiabilitiesAssets    decimal.Decimal `json:"LiabilitiesAssets" db:"liabilities_assets"`
	CurrentLiquidity     decimal.Decimal `json:"CurrentLiquidity" db:"current_liquidity"`
	CAGRRevenue          decimal.Decimal `json:"CAGRRevenue" db:"cagr_revenue"`
	CAGRProfit           decimal.Decimal `json:"CAGRProfit" db:"cagr_profit"`
	CrawledAt            time.Time       `json:"CrawledAt" db:"crawled_at"`
}
