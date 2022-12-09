package meta

//
//// Meta 元数据，用于context metadata
//type Meta struct {
//	UID string `json:"uid"`
//}
//
//// FromIncomingCtx 获取元数据（Grpc 内部通讯使用）
//func FromIncomingCtx(ctx context.Context) *Meta {
//	data := &Meta{ChannelID: constants.DefaultChannelID}
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return data
//	}
//	if realIPs := md.Get(constants.RealIP); len(realIPs) > 0 {
//		data.RealIP = conversion.String(realIPs[0])
//	}
//
//	if channelIDs := md.Get(constants.ChannelID); len(channelIDs) > 0 {
//		data.ChannelID = channelIDs[0]
//	}
//
//	if deviceTypes := md.Get(constants.DeviceType); len(deviceTypes) > 0 {
//		data.DeviceType = conversion.Int32(deviceTypes[0])
//	}
//	return data
//}
