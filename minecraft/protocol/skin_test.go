package protocol

import "testing"

func TestPersonaPieceTintWireNames(t *testing.T) {
	for login, wire := range map[string]string{
		"persona_facial_hair":    "facialhair",
		"persona_face_accessory": "faceaccessory",
		"persona_hand":           "hands",
		"persona_high_pants":     "high_pants",
		"persona_hair":           "hair",
		"persona_left_arm":       "leftarm",
	} {
		if got := personaPieceTintWireType(login); got != wire {
			t.Fatalf("%v: wire %v, want %v", login, got, wire)
		}
		if got := personaPieceTintLoginType(wire); got != login {
			t.Fatalf("%v: login %v, want %v", wire, got, login)
		}
	}
}
