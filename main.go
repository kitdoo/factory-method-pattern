package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type setList interface {
	setName(name string)
	setPower(power int)
	setTracks(songs string)
	getName() string
	getPower() int
	getTracks() string
}

type setlist struct {
	name  string
	power int
	songs string
}

func (s *setlist) setName(name string) {
	s.name = name
}

func (s *setlist) getName() string {
	return s.name
}

func (s *setlist) setPower(power int) {
	s.power = power
}

func (s *setlist) getPower() int {
	return s.power
}

func (s *setlist) setTracks(songs string) {
	s.songs = songs
}

func (s *setlist) getTracks() string {
	return s.songs
}

func shuffleTracks(power int) string {
	var res []string
	for key, value := range BTSSongs {
		if power-2 <= value && value <= power+1 {
			res = append(res, key)
		}
	}
	rand.Shuffle(len(res),
		func(i, j int) {
			res[i], res[j] = res[j], res[i]
		})
	songs := strings.Join(res[:10], " , ")
	return songs
}

type setListBrazil struct {
	setlist
}

func newSetListBrazil() setList {
	return &setListBrazil{
		setlist: setlist{
			name:  "SanPaulo",
			power: maxPowerForBrazil,
			songs: shuffleTracks(maxPowerForBrazil),
		},
	}
}

type setListRussia struct {
	setlist
}

func newSetListRussia() setList {
	return &setListRussia{
		setlist: setlist{
			name:  "Moscow",
			power: maxPowerForRussia,
			songs: shuffleTracks(maxPowerForBrazil),
		},
	}
}

type setListFrance struct {
	setlist
}

func newSetListFrance() setList {
	return &setListFrance{
		setlist: setlist{
			name:  "Paris",
			power: maxPowerForFrance,
			songs: shuffleTracks(maxPowerForFrance),
		},
	}
}

func getSetlist(country string) (setList, error) {
	if country == "br" {
		return newSetListBrazil(), nil
	}
	if country == "ru" {
		return newSetListRussia(), nil
	}
	if country == "fr" {
		return newSetListFrance(), nil
	}
	return nil, fmt.Errorf("this country ( %s ) is not on the list of countries to visit BTS in the World Tour 2022", country)
}

func main() {
	brazilTracks, _ := getSetlist("br")
	russiaTracks, _ := getSetlist("ru")
	franceTracks, _ := getSetlist("fr")
	usaTracks, err := getSetlist("us")
	if err != nil {
		fmt.Println(err)
	}

	franceTracks.setTracks("Mic Drop, Serendipity, Just Dance, Sweet Night, Idol, Best of me")

	printDetails(brazilTracks)
	printDetails(russiaTracks)
	printDetails(franceTracks)
	printDetails(usaTracks)
}

func printDetails(s setList) {
	if s != nil {
		fmt.Println()
		fmt.Printf("Setlist for city: %s (power: %d )", s.getName(), s.getPower())
		fmt.Println()
		fmt.Printf("Songs: %s", s.getTracks())
		fmt.Println()
	}
}
