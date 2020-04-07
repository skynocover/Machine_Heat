package main

import (
	//"fmt"
	"math"
	"strconv"
)

var (
	f0 = [][]float64{{0.088, 0.28, 0.8}, {0.13, 0.46, 1}} //f0使用lube以及btype計算
)

type Bearing struct {
	Dm float64
	Pcs float64
	Row float64
	C0 float64
	Angle float64
	Bearingtype float64
}

type Work struct {
	Rpm float64
	Fu float64
	Fr float64
	Lubetype float64
	V float64
}

type topera struct {
	Bearing	
	Work
	result
}

type result struct{
	f0 float64
	p0 float64
	f1 float64
	g1p0 float64
	mv float64
	ml float64
	m float64
	q float64
}

func (opera *topera)calc(){
	opera.p0 = math.Floor(opera.Fu*9.81/math.Tan(opera.Angle*math.Pi/180)*100000+0.5) / 100000 //靜等價賀重
	opera.f0 = f0[int(opera.Bearingtype)][int(opera.Lubetype)]*opera.Row                                       //潤滑定數
	//軸承型式定數
	if opera.Bearingtype == 0 {
		opera.f1 = math.Floor(0.001*opera.Pcs*math.Pow(opera.p0/opera.C0, 0.33)*100000+0.5) / 100000
	} else {
		opera.f1 = 0.0003
	}

	//荷重常數
	opera.g1p0 = math.Floor((0.9*opera.Fu/math.Tan((opera.Angle)*math.Pi/180)-0.1*opera.Fr)*10000+0.5) / 10000
	if opera.g1p0 < opera.Fr {
		opera.g1p0 = opera.Fr
	}

	opera.ml = math.Floor(opera.f1*opera.g1p0*opera.Dm*math.Pow(10, -3)*100000+0.5) / 100000
	opera.mv = math.Floor((opera.Pcs*opera.f0*math.Pow(opera.Dm, 3)*math.Pow(opera.V*opera.Rpm, (0.6666666667))*math.Pow(10, -11))*100000+0.5) / 100000
	opera.m = opera.ml + opera.mv
	opera.q = math.Floor(0.00234*math.Pi*opera.m*opera.Rpm*60*2*1000+0.5) / 1000
}

func (opera *topera) show() {
	ui.Eval(`document.getElementById("mv").innerText = ` + strconv.FormatFloat(opera.mv, 'E', -1, 64))
	ui.Eval(`document.getElementById("ml").innerText = ` + strconv.FormatFloat(opera.ml, 'E', -1, 64))
	ui.Eval(`document.getElementById("q").innerText = ` + strconv.FormatFloat(opera.q, 'E', -1, 64))
	ui.Eval(`document.getElementById("w").innerText = ` + strconv.FormatFloat(opera.q*1.163, 'E', 4, 64))
}
