package greedy

import (
	"github.com/truongnhukhang/hellogo/object"
	"github.com/truongnhukhang/hellogo/queue"
	"strconv"
)

type WordNode struct {
	value  string
	weight int
	hash   string
	left   *WordNode
	right  *WordNode
}

func (wn *WordNode) CompareWith(node interface{}) int {
	tempNode := node.(*WordNode)
	return wn.weight - tempNode.weight
}

func (wn *WordNode) String() string {
	return string(wn.value) + " : " + strconv.Itoa(wn.weight) + " : " + wn.hash
}

type HuffmanCoding struct {
	Context     string
	HuffmanCode map[string]string
}

func (hc *HuffmanCoding) CompressData() string {
	wordMap := map[uint8]*WordNode{}
	data := hc.Context
	createCharCountMap(data, &wordMap)
	wordNodeList := []object.Object{}
	// put all WordNode to the PriorityQueue
	wordNodeQueue := queue.PriorityQueue{wordNodeList}
	for _, v := range wordMap {
		teamNode := v
		wordNodeQueue.Put(teamNode)
	}
	// start buff Huffman Tree
	for wordNodeQueue.Size() > 1 {
		left := wordNodeQueue.Poll()
		right := wordNodeQueue.Poll()
		parentNode := WordNode{}
		parentNode.value = ""
		parentNode.weight = left.(*WordNode).weight + right.(*WordNode).weight
		parentNode.left = left.(*WordNode)
		parentNode.right = right.(*WordNode)
		wordNodeQueue.Put(&parentNode)
	}
	root := (wordNodeQueue.Poll()).(*WordNode)
	hc.HuffmanCode = map[string]string{}
	updateWordHash(root, "", &hc.HuffmanCode)
	result := ""
	for i := 0; i < len(data); i++ {
		result = result + wordMap[data[i]].hash
	}
	return result
}

func createCharCountMap(data string, wordMap *map[uint8]*WordNode) {
	for i := 0; i < len(data); i++ {
		if (*wordMap)[data[i]] == nil {
			tempWord := WordNode{
				value: string(data[i]), weight: 1, left: nil, right: nil,
			}
			(*wordMap)[data[i]] = &tempWord
		} else {
			tempWord := (*wordMap)[data[i]]
			tempWord.weight = tempWord.weight + 1
		}
	}
}

func (hc *HuffmanCoding) UncompressData(hashData string) string {
	result := ""
	tempWordCode := ""
	for i := 0; i < len(hashData); i++ {
		tempWordCode = tempWordCode + string(hashData[i])
		if hc.HuffmanCode[tempWordCode] != "" {
			result = result + hc.HuffmanCode[tempWordCode]
			tempWordCode = ""
		}
	}
	return result
}

func updateWordHash(node *WordNode, hashValue string, huffmanMap *map[string]string) {
	left := node.left
	right := node.right
	if left != nil {
		updateWordHash(left, hashValue+"0", huffmanMap)
	}
	if right != nil {
		updateWordHash(right, hashValue+"1", huffmanMap)
	}
	if left == nil && right == nil {
		node.hash = hashValue
		(*huffmanMap)[hashValue] = node.value
	}
}
