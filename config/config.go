package config

const DefaultFuturesHost = "api.hbdm.com"
const DefaultSpotHost = "api-aws.huobi.pro"

const DefaultId = "api"
const DefaultCid = "cid"

var PrivatePaths = []string{
	"/linear-swap-notification",
	"/ws/v2",
	"/ws/v1",
}

const UsageAgent = "litter man"

type Config struct {
	Host      string
	AccessKey string
	SecretKey string
	Debug     bool
}
