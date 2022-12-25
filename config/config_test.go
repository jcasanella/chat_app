package config

import "testing"

func TestGetConfigValid(t *testing.T) {
	portWanted := "8080"

	Init("local")
	config := GetConfig()
	portActual := config.GetString("port")

	if portActual != portWanted {
		t.Errorf("GetConfig(), property port value wrong = %q, want %q", portActual, portWanted)
	}
}
