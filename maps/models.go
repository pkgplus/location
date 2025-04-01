package maps

type Geocode struct {
	Province string `json:"province"` // 地址所在的省份名，例如：北京市
	City     string `json:"city"`     // 地址所在的城市名，例如：北京市
	Address  string `json:"address"`  // 参考地址信息: 省份＋城市＋区县＋城镇＋乡村＋街道＋门牌号码
	Location string `json:"location"` // 坐标点，格式为"经度,纬度"
	Level    string `json:"level"`    // 匹配级别
}
