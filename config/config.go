package config

type Config struct {
	ScreenWidth  float64
	ScreenHeight float64
	PlayerWidth  float64
	PlayerHeight float64
	Gravity      float64
	JumpForce    float64
}

var C *Config

func init() {
	C = &Config{
		ScreenWidth:  640,
		ScreenHeight: 480,
		PlayerWidth:  20,
		PlayerHeight: 30,
		Gravity:      2,
		JumpForce:    6,
	}
}
