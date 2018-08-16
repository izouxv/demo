package structs

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
)

//生成块函数
func GenerateBlock(oldBlock *Block, BPM int) (*Block, error) {
	t := time.Now()
	newBlock :=  &Block{Index:oldBlock.Index + 1,Timestamp:t.String(),BPM:BPM,PrevHash:oldBlock.Hash}
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock, nil
}

//计算数据的 SHA256 散列值
func CalculateHash(block *Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//块校验函数
func IsBlockValid(oldBlock, newBlock *Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

//更新链
func ReplaceChain(newBlocks []*Block) {
	if len(newBlocks) > len(BlockChain) {
		BlockChain = newBlocks
	}
}
