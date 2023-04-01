package main

import "fmt"

type Car struct {
    Type   string
    FuelIn float64
}

func estimasiJarak(mobil Car) float64 {
    return mobil.FuelIn * 1.5
}

func main() {
    mobil1 := Car{"Avanza", 10.5}
    estimasiJarak1 := estimasiJarak(mobil1)
    fmt.Printf("%s dapat melakukan perjalanan sekitar %.2f km dengan %.2f liter dari pertalite\n", mobil1.Type, estimasiJarak1, mobil1.FuelIn)

	mobil2 := Car{"xenia", 23.5}
    estimasiJarak2 := estimasiJarak(mobil2)
    fmt.Printf("%s dapat melakukan perjalanan sekitar %.2f km dengan %.2f liter dari pertalite\n", mobil2.Type, estimasiJarak2, mobil2.FuelIn)
}
