package net

const (
	LatestProtocol     = 140
	GameVersion        = "v1.2.6.2"
	GameVersionNetwork = "1.2.6.2"
)

const (
	LoginPacket = 0x01
	PlayStatusPacket = 0x02
	ClientHandshakePacket = 0x03
	ServerHandshakePacket = 0x04
	DisconnectPacket = 0x05
	ResourcePackInfo = 0x06
	ResourcePackStack = 0x06
	ResourcePackClientResponsePacket = 0x08
	TextPacket = 0x09
	SetTimePacket = 0x0a
	StartGamePacket = 0x0b
	AddPlayerPacket = 0x0c
	AddEntityPacket = 0x0d
	RemoveEntityPacket = 0x0e
	AddItemEntityPacket = 0x0f
	AddHangingEntityPacket = 0x10
	TakeItemEntityPacket = 0x11
	MoveEntityPacket = 0x12
	MovePlayerPacket = 0x13
	RiderJumpPacket = 0x14
	UpdateBlockPacket = 0x15
	AddPaintingPacket = 0x16
	ExplodePacket = 0x17
	LevelSoundEventPacket = 0x18
	LevelEventPacket = 0x19
	BlockEventPacket = 0x1a
	EntityEventPacket = 0x1b
	MobEffectPacket = 0x1c
	UpdateAttributesPacket = 0x1d
	InventoryTransactionPacket = 0x1e
	MobEquipmentPacket = 0x1f
	MobArmorEquipmentPacket = 0x20
	InteractPacket = 0x21
	BlockPickRequestPacket = 0x22
	EntityPickRequestPacket = 0x23
	PlayerActionPacket = 0x24
	EntityFallPacket = 0x25
	HurtArmorPacket = 0x26
	SetEntityDataPacket = 0x27
	SetEntityMotionPacket = 0x28
	SetEntityLinkPacket = 0x29
	SetHealthPacket = 0x2a
	SetSpawnPositionPacket = 0x2b
	AnimatePacket = 0x2c
	RespawnPacket = 0x2d
	ContainerOpenPacket = 0x2e
	ContainerClosePacket = 0x2f
	PlayerHotbarPacket = 0x30
	InventoryContentPacket = 0x31
	InventorySlotPacket = 0x32
	ContainerSetDataPacket = 0x33
	CraftingDataPacket = 0x34
	CraftingEventPacket = 0x35
	GuiDataPickItemPacket = 0x36
	AdventureSettingsPacket = 0x37
	BlockEntityDataPacket = 0x38
	PlayerInputPacket = 0x39
	FullChunkDataPacket = 0x3a
	SetCommandsEnabledPacket = 0x3b
	SetDifficultyPacket = 0x3c
	ChangeDimensionPacket = 0x3d
	SetPlayerGameTypePacket = 0x3e
	PlayerListPacket = 0x3f
	SimpleEventPacket = 0x40
	EventPacket = 0x41
	SpawnExperienceOrbPacket = 0x42
	ClientboundMapItemDataPacket = 0x43
	MapInfoRequestPacket = 0x44
	RequestChunkRadiusPacket = 0x45
	ChunkRadiusUpdatedPacket = 0x46
	ItemFrameDropItemPacket = 0x47
	GameRulesChangedPacket = 0x48
	CameraPacket = 0x49
	BossEventPacket = 0x4a
	ShowCreditsPacket = 0x4b
	AvailableCommandsPacket = 0x4c
	CommandRequestPacket = 0x4d
	CommandBlockUpdatePacket = 0x4e
	CommandOutputPacket = 0x4f
	UpdateTradePacket = 0x50
	UpdateEquipPacket = 0x51
	ResourcePackDataInfoPacket = 0x52
	ResourcePackChunkDataPacket = 0x53
	ResourcePackChunkRequestPacket = 0x54
	TransferPacket = 0x55
	PlaySoundPacket = 0x56
	StopSoundPacket = 0x57
	SetTitlePacket = 0x58
	AddBehaviorTreePacket = 0x59
	StructureBlockUpdatePacket = 0x5a
	ShowStoreOfferPacket = 0x5b
	PurchaseReceiptPacket = 0x5c
	PlayerSkinPacket = 0x5d
	SubClientLoginPacket = 0x5e
	WSConnectPacket = 0x5f
	SetLastHurtByPacket = 0x60
	BookEditPacket = 0x61
	NpcRequestPacket = 0x62
	PhotoTransferPacket = 0x63
	ModalFormRequestPacket = 0x64
	ModalFormResponsePacket = 0x65
	ServerSettingsRequestPacket = 0x66
	ServerSettingsResponsePacket = 0x67
	ShowProfilePacket = 0x68
	SetDefaultGameTypePacket = 0x69
)