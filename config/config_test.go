package config

import "testing"

func TestGetConfig(t *testing.T) {
	portWanted := 8080

	Init("local.yaml")
	config := GetConfig()
	portActual := config.GetInt("port")

	if portActual != portWanted {
		t.Errorf("GetConfig(), property port value wrong = %d, want %d", portActual, portWanted)
	}
}
