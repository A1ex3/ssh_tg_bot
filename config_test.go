package main

import "testing"

func TestUnmarshal(t *testing.T){
	config := &Config{}
	cfgunm := config.Unmarshal("config.json")

	if cfgunm != nil{
		t.Error(cfgunm)
	}
}