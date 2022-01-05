module github.com/forbole/bdjuno/v2

go 1.16

require (
	github.com/Sifchain/sifnode v0.0.0-20211110000810-22b3ac797839
	github.com/cosmos/cosmos-sdk v0.42.10
	github.com/forbole/juno/v2 v2.0.0-20211221121955-cf2fcf04394b
	github.com/go-co-op/gocron v1.11.0
	github.com/gogo/protobuf v1.3.3
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.4
	github.com/pelletier/go-toml v1.9.4
	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/tendermint/tendermint => github.com/huichiaotsou/tendermint v1.0.3 // perso fork for customized tendermint

replace github.com/cosmos/cosmos-sdk => github.com/huichiaotsou/cosmos-sdk v1.0.5 // remove final stake / current stake panic

// replace github.com/forbole/juno/v2 => github.com/huichiaotsou/juno/v2 v2.0.0-20211229050548-ceaf3de8e114 // add fix certain blocks cmd

replace github.com/forbole/juno/v2 => github.com/huichiaotsou/juno/v2 v2.0.0-20211231074226-85621cdb1fb4 // add time log
