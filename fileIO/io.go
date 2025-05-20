package fileIO

import (
	"TerminalAI/logger"
	"fmt"
	"os"
	"sync"
)

var f *os.File
var confFile *os.File
var once sync.Once
var chatNumber string
var tempChatNumber = make([]byte, 1)
var err error
var currentChat string = ""

func InitFileIO() {
	once.Do(func() {

		confFile, err = os.OpenFile("./fileIO/config.txt", os.O_RDWR, 0644)
		if err != nil {
			logger.Logger().Error("Error reading config file:", err)
			return
		}

		_, err := confFile.Read(tempChatNumber)
		if err != nil {
			logger.Logger().Error("Error reading chat number from config file:", err)
			return
		}
		chatNumber = string(tempChatNumber)

		f, err = os.Create("./chats/chat_history_" + chatNumber + ".txt")
		if err != nil {
			logger.Logger().Error("Error creating chat history file:", err)
			return
		}

		logger.Logger().Info("Chat history file created: ./chats/chat_history_"+chatNumber+".txt")
	})
}

func CloseFile() {

	tempChatNumber[0]++
	confFile, err = os.Create("./fileIO/config.txt")
	if err != nil {
		logger.Logger().Error("Error recreating config file to overwrite:", err)
		return
	}

	_, err = confFile.Write(tempChatNumber)
	if err != nil {
		logger.Logger().Error("Error writing chat number to config file:", err)
		return
	}

	err = confFile.Sync()
	if err != nil {
		logger.Logger().Error("Error syncing config file:", err)
		return
	}

	if confFile != nil {
		err := confFile.Close()
		if err != nil {
			logger.Logger().Error("Error closing config file:", err.Error())
		} else {
			logger.Logger().Info("Config file closed successfully")
		}
		confFile = nil
	}

	if f != nil {
		err := f.Close()
		if err != nil {
			logger.Logger().Error("Error closing file:", err.Error())
		} else {
			logger.Logger().Info("File closed successfully")
		}
		f = nil
	}
}

func WriteToFile(data string) {
	InitFileIO()
	if f == nil {
		logger.Logger().Error("File not initialized")
		return
	}

	currentChat += data + "\n"
	
	if _, err := f.WriteString(data + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)	
		return
	}
}

func GetChat() string{
	return currentChat
}