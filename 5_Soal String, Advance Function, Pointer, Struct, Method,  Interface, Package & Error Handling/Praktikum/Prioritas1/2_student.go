package main

import (
    "fmt"
    "sort"
)

type Student struct {
    name  string
    nilai []int
}

func (s *Student) addNilai(nilaiInput int) {
    s.nilai = append(s.nilai, nilaiInput)
}

func (s *Student) getRata2() float64 {
    jumlahNilai := 0
    for _, nilai := range s.nilai {
        jumlahNilai += nilai
    }
    return float64(jumlahNilai) / float64(len(s.nilai))
}

func (s *Student) getMinScore() (string, int) {
    nilaiMin := s.nilai[0]
    namaMin := s.name
    for i := 1; i < len(s.nilai); i++ {
        if s.nilai[i] < nilaiMin {
            nilaiMin = s.nilai[i]
            namaMin = s.name
        }
    }
    return namaMin, nilaiMin
}

func (s *Student) getMaxScore() (string, int) {
    nilaiMax := s.nilai[0]
    namaMax := s.name
    for i := 1; i < len(s.nilai); i++ {
        if s.nilai[i] > nilaiMax {
            nilaiMax = s.nilai[i]
            namaMax = s.name
        }
    }
    return namaMax, nilaiMax
}

func main() {
    var students [5]Student

    for i := 0; i < 5; i++ {
        var name string
        var nilai int
        fmt.Printf("Input %d Student's Name: ", i+1)
        fmt.Scanln(&name)
        fmt.Printf("Input %d Student's Score: ", i+1)
        fmt.Scanln(&nilai)

        students[i] = Student{name: name, nilai: []int{nilai}}
    }

    var sumScore int
    for i := 0; i < 5; i++ {
        sumScore += students[i].nilai[0]
    }

	fmt.Println("\nOutput:")

    avgScore := float64(sumScore) / float64(len(students))
    fmt.Printf("Average Score: %d\n", int(avgScore))

    sort.Slice(students[:], func(i, j int) bool {
        return students[i].getRata2() < students[j].getRata2()
    })

    namaMin, nilaiMin := students[0].getMinScore()
    fmt.Printf("Min Score of Students: %s (%d)\n", namaMin, nilaiMin)

    namaMax, nilaiMax := students[4].getMaxScore()
    fmt.Printf("Max Score of Students: %s (%d)\n", namaMax, nilaiMax)
}
