package v8

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v1.0.4"
	// TODO: MainnetUpgradeHeight defines the Arcis mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 0
	// TODO: TestnetUpgradeHeight defines the Arcis testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 0
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/Ambiplatforms-TORQUE/arcis/releases/download/v1.0.4/arcis_1.0.4_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/Ambiplatforms-TORQUE/arcis/releases/download/v1.0.4/arcis_1.0.4_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/Ambiplatforms-TORQUE/arcis/releases/download/v1.0.4/arcis_1.0.4_Linux_arm64.tar.gz","linux/x86_64":"https://github.com/Ambiplatforms-TORQUE/arcis/releases/download/v1.0.4/arcis_1.0.4_Linux_x86_64.tar.gz","windows/x86_64":"https://github.com/Ambiplatforms-TORQUE/arcis/releases/download/v1.0.4/arcis_1.0.4_Windows_x86_64.zip"}}'`
)
