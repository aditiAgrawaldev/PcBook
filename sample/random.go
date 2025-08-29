package sample

import (
	"github.com/aditiAgrawaldev/pcBook/pb"
	"github.com/google/uuid"
	"math/rand/v2"
)

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.IntN(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"XEON E-2286M",
			"Core i9-pp80HK",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 Pro 2700U",
		"Ryzen 8 Pro 8976P",
	)

}

func randomInt(min int, max int) int {
	return min + rand.IntN(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.IntN(n)]
}

func randomBool() bool {
	return rand.IntN(2) == 1
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GeForce RTX 20",
			"GeForce RTX 30",
		)

	}
	return randomStringFromSet(
		"Radeon RX 9000",
		"Radeon RX 7000",
	)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.IntN(2) == 1 {
		return pb.Screen_IPS
	}

	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenevo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":

		return randomStringFromSet("MacBookPro", "MacBookAir")
	case "Dell":
		return randomStringFromSet("Latitude", "vostro", "XPS")
	default:
		return randomStringFromSet("ThinkPad X1", "ThinkPad P1")
	}
}
