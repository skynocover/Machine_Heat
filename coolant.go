package main

type tworker struct {
	RoughTime float64
	FinishTime float64
	Chipremoval float64
	SpindlePower float64
}

func (worker *tworker)q() float64 {	
	return worker.SpindlePower*0.3*(worker.FinishTime*0.6+worker.RoughTime*1)
}

type tpump struct {
	Power1 float64
	Num1 float64
	Power2 float64
	Num2 float64
	Power3 float64
	Num3 float64
	Power4 float64
	Num4 float64
}

func (pump *tpump)q() float64 {
	return 0.35*(pump.Power1*pump.Num1+pump.Power2*pump.Num2+pump.Power3*pump.Num3+pump.Power4*pump.Num4)
}

type tcoolant struct {
	Surface float64
	Temperature float64
}

func (coolant *tcoolant)q() float64 {
	return 7*coolant.Surface*coolant.Temperature/1000
}