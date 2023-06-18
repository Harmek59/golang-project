package config

type Config struct {
	ScreenWidth   float64
	ScreenHeight  float64
	PlayerWidth   float64
	PlayerHeight  float64
	Gravity       float64
	JumpForce     float64
	CameraYOffset float64
}

var C *Config

func init() {
	C = &Config{
		ScreenWidth:   854,
		ScreenHeight:  480,
		PlayerWidth:   60,
		PlayerHeight:  90,
		Gravity:       1000,
		JumpForce:     500,
		CameraYOffset: 480 / 4,
	}
}
