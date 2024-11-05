package entity

// Album
// 相册
type Album struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`        // 相册标题
	Artist      string   `json:"artist"`       // 创建人名字
	Summary     string   `json:"summary"`      // 概况
	Description string   `json:"description"`  // 描述
	Number      int      `json:"price"`        // 照片数量
	OrderIndex  int      `json:"order_index"`  // 排序，升序
	FirstPic    *Picture `json:"first_pic"`    // 第一张照片
	LastViewAt  string   `json:"last_view_at"` // 上次查看日期
	_base
}

// Picture
// 相片
type Picture struct {
	ID       string   `json:"id"`
	FileName string   `json:"file_name"` // 文件名
	Suffix   string   `json:"suffix"`    // 后缀名
	Size     int64    `json:"size"`      // 照片大小，字节
	Hidden   bool     `json:"hidden"`    // 是否隐藏
	AbsPath  string   `json:"abs_path"`  // 相对路径
	Tags     []string `json:"tags"`      // 标签
	Location string   `json:"location"`  // 拍摄位置
	RealPath string   `json:"real_path"` // 真实路径
	_base
}
