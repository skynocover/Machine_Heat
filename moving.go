package main

import (
	"math"
)

type Tmoving struct {
	Axiaload float64
	Weight float64
	V float64
	Tballscrew
	Tguide
}

type Tballscrew struct {
	Lead float64	
	Effectiveness float64
	torque float64
	rpm float64
}

type Tguide struct {
	Type float64 //型式 
	R float64 //滾動體半徑
	U float64 //摩擦係數
}

func (moving *Tmoving) guideHeat()float64 {
	if moving.Tguide.Type==0 { //滑動導軌
		return moving.Tguide.U*moving.Weight*moving.V/1000/60
	}else{
		if moving.Tguide.R<=0 {
			ui.Eval(`alert("請填入滾動體半徑");`)
		}
		return moving.Tguide.U*moving.Weight*moving.V/1000/60/moving.Tguide.R
	}
}

func (moving *Tmoving) screwHeat() float64 {
	moving.Tballscrew.rpm = moving.V/moving.Tballscrew.Lead
	moving.Tballscrew.torque = (moving.Axiaload+moving.Tguide.U*moving.Weight)*moving.Tballscrew.Lead/2/math.Pi/moving.Tballscrew.Effectiveness
	return 9.81*2*math.Pi*moving.Tballscrew.rpm/60*moving.Tballscrew.torque/1000
}