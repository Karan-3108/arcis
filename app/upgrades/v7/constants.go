package v7

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v1.0.3a"
	// MainnetUpgradeHeight defines the Arcis mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 0
	// TODO: TestnetUpgradeHeight defines the Arcis testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 0
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/Ambiplatforms-TORQUE/arcis/releases/download/v1.0.3/arcis_1.0.3_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/arcis/arcis/releases/download/v1.0.3/arcis_1.0.3_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/arcis/arcis/releases/download/v1.0.3/arcis_1.0.3_Linux_arm64.tar.gz","linux/x86_64":"https://github.com/arcis/arcis/releases/download/v1.0.3/arcis_1.0.3_Linux_x86_64.tar.gz","windows/x86_64":"https://github.com/arcis/arcis/releases/download/v1.0.3/arcis_1.0.3_Windows_x86_64.zip"}}'`

	// FaucetAddressFrom is the inaccessible secp address of the Testnet Faucet
	FaucetAddressFrom = "arcis1z4ya98ga2xnffn2mhjym7tzlsm49ec23890sze"
	// FaucetAddressTo is the new eth_secp address of the Testnet Faucet
	FaucetAddressTo = "arcis1ujm4z5v9zkdqm70xnptr027gqu90f7lxjr0fch"

	// ContributorAddrFrom is the lost address of an early contributor
	ContributorAddrFrom = "arcis1659xwt0hnu5humgek7scefhnpcm2w6hyvy4fsq"
	// ContributorAddrTo is the new address of an early contributor
	ContributorAddrTo = "arcis1pktlmqrz448cuazl98tqmsj4kjwpqpmaa0cjcf"
)
