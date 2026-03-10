package sample

import (
	"crypto/rand"

	"github.com/62Kinsley/GoApplication/pb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func NewCPU() *pb.CPU {
	return &pb.CPU{
		Brand: randomCPUBrand(),
		Name: randomCPUName(brandomCPUBrand()),
		NumberCores: uint32(randomInt(2, 8)),
		NumberThreads: uint32(randomInt(4, 16)),
		MinGhz: randomFloat64(2.0, 3.5),
		MaxGhz: randomFloat64(3.5, 5.0),
	
	}
}