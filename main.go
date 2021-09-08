package main

import "fmt"

func main() {
	// compressed, _ := compress([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus malesuada massa sed dolor dapibus, id tincidunt sem maximus. Integer faucibus volutpat ornare. Ut sed turpis arcu. Vivamus dignissim imperdiet magna mattis volutpat. Etiam non nunc ut orci interdum imperdiet. Maecenas sollicitudin vitae ipsum eu feugiat. Mauris venenatis velit condimentum ligula ultricies convallis. Morbi aliquam in felis ac hendrerit. Maecenas mollis ex egestas cursus vulputate. Maecenas enim ligula, ullamcorper tempor nisi quis, varius varius risus."))
	// decompressed, _ := decompress(compressed)
	// fmt.Println(string(decompressed))

	ptrConfig, _ := createConfiguration()
	ptrConfig.setServiceHashList([]uint64{12, 54, 32, 432})
	ptrConfig.appendServiceHashList(1322)
	ptrConfig.removeServiceHashList(54)
	fmt.Println(*ptrConfig)
}
