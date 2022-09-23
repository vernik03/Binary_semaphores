package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Semaphore struct {
	i int
}

func (s *Semaphore) Acquire() {
	for s.isAcquired() {
		time.Sleep(time.Millisecond)
	}
	s.i++
}

func (s *Semaphore) Release() {
	if s.isAcquired() {
		s.i--
	}
}

func (s *Semaphore) isAcquired() bool {
	return s.i > 0
}


type Sigarette struct {
	paper int
	tobacco int
	matches int
}

type Smoker struct {
	paper int
	tobacco int
	matches int
}

type Agent struct {
	paper int
	tobacco int
	matches int
}

func (a *Agent) Put(c chan Sigarette, sem *Semaphore) {
	sem.Acquire()
	temp_s := Sigarette{0, 0, 0}
	elem := rand.Intn(3)
	switch elem {
	case 0:
		temp_s.tobacco++
		temp_s.matches++
		fmt.Println("🤵 Put tobacco and matches")
	case 1:
		temp_s.paper++
		temp_s.matches++
		fmt.Println("🤵 Put paper and matches")
	case 2:
		temp_s.paper++
		temp_s.tobacco++
		fmt.Println("🤵 Put paper and tobacco")
	}
	time.Sleep(500 * time.Millisecond)
	c <- temp_s
	sem.Release()
}

func (s *Smoker) Get(c chan Sigarette, sem *Semaphore, name string) {
	
	temp := <-c
	fmt.Println("\t🔬 Checking sigarette")
	if temp.paper == 0 && s.paper > 0 {
		sem.Acquire()
		fmt.Println("🚬 Get tobacco and matches")
		s.paper++
		s.Smoke(sem, name)
	} else if temp.tobacco == 0 && s.tobacco > 0 {
		sem.Acquire()
		fmt.Println("🚬 Get paper and matches")
		s.tobacco++
		s.Smoke(sem, name)
	} else if temp.matches == 0 && s.matches > 0 {
		sem.Acquire()
		fmt.Println("🚬 Get paper and tobacco")
		s.matches++
		s.Smoke(sem, name)
	}else {
		c <- temp
	}
}
func (s *Smoker) Smoke(sem *Semaphore, name string) {
	fmt.Println("🚬 Smoking " + name)
	s.paper--
	s.tobacco--
	s.matches--
	time.Sleep(1000 * time.Millisecond)
	sem.Release()
}

func (s *Smoker) LifeOfSmoker(c chan Sigarette, sem *Semaphore, name string) {
	for {
		time.Sleep(time.Millisecond)
        s.Get(c, sem, name);
	}
}

func (s *Agent) LifeOfAgent(c chan Sigarette, sem *Semaphore) {
	for {
		time.Sleep(time.Millisecond)
        s.Put(c, sem);
	}
}

func main() {
	c := make(chan Sigarette)
	paper := Smoker{1000, 0, 0}
	tobacco := Smoker{0, 1000, 0}
	matches := Smoker{0, 0, 1000}
	agent := Agent{1000, 1000, 1000}
	semaphore := Semaphore{0}

	
	go agent.LifeOfAgent(c, &semaphore)

	go paper.LifeOfSmoker(c, &semaphore, "🔵")
	go tobacco.LifeOfSmoker(c, &semaphore, "🟢")
	go matches.LifeOfSmoker(c, &semaphore, "🟡")

	finished := make(chan bool)
	<- finished
}
