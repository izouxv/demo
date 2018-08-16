package structs

//区块数据模型
type Block struct {
	Index     int		//是这个块在整个链中的位置
	Timestamp string	//块生成时的时间戳
	BPM       int		//每分钟心跳数，也就是心率
	Hash      string	//块通过 SHA256 算法生成的散列值
	PrevHash  string	//前一个块的 SHA256 散列值
	Difficulty int		//难度值
	Nonce      string	//随机数
}

//结构表示整个链，最简单的表示形式是slice
var BlockChain []*Block


type Message struct {
	BPM int
}