package main

type tmotor struct {
	Type float64 
	Power float64
	Effectiveness float64
}

func (motor *tmotor)q() float64 {
	motor.Effectiveness = motor.effectiveness()
	if motor.Power<=1 {		
		return (1/motor.Effectiveness-1)*motor.Power
	}else{
		return (1-motor.Effectiveness)*motor.Power
	}
}

func (motor *tmotor)effectiveness()float64 {
	if motor.Type==2 {
		return 0.95 
	}else if motor.Type==0{
		
		return motor.aceff()/100
	}else{
		return motor.dceff()/100
	}
}

func (motor *tmotor)aceff()float64{
	if motor.Power<=0.4 {
		return 70.5
	}else if motor.Power <=5.5{
		return (motor.Power-0.4)*((83-70.5)/(5.5-0.4))+83
	}else if motor.Power <=37{
		return (motor.Power-5.5)*((88-83)/(37-5.5))+88
	}else {
		return 88
	}
}

func (motor *tmotor)dceff()float64{
	if motor.Power<=3.7{
		return 80
	}else if motor.Power>= 37{
		return 90
	}else{
		return (motor.Power-3.7)*((90-80)/(37-3.7))+80
	}
}