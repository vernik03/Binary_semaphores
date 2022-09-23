package main

import (
	"fmt"
	"math/rand"
	// "os"
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
	s.i--
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

func (a *Agent) Put(c chan Sigarette, sem *Semaphore, sem_agent *Semaphore) {
	sem.Acquire()
	fmt.Println("Putting")
	temp_s := Sigarette{0, 0, 0}
	for i := 0; i < 2; i++ {
		elem := rand.Intn(3)
		switch elem {
		case 0:
			temp_s.paper++
			fmt.Println("Put paper")
		case 1:
			temp_s.tobacco++
			fmt.Println("Put tobacco")
		case 2:
			temp_s.matches++
			fmt.Println("Put matches")
		}
	}
	c <- temp_s
	sem.Release()
	sem_agent.Acquire()
}

func (s *Smoker) Get(c chan Sigarette, sem *Semaphore, sem_agent *Semaphore) {
	
	fmt.Println("Smoker is checking sigarette")
	temp := <-c
	if temp.paper > 0 && s.paper == 0 {
		s.paper++
	} else if temp.tobacco > 0 && s.tobacco == 0 {
		s.tobacco++
	} else if temp.matches > 0 && s.matches == 0 {
		s.matches++
	}else {
		c <- temp
	}
	
	sem.Acquire()
	if s.Check() {
		s.Smoke()
	}
	sem.Release()
	sem_agent.Release()
}

func (s *Smoker) Check() bool {
	return s.paper > 0 && s.tobacco > 0 && s.matches > 0
}

func (s *Smoker) Smoke() {
	fmt.Println("Smoker is smoking")
	s.paper--
	s.tobacco--
	s.matches--
	time.Sleep(1000 * time.Millisecond)
}

func (s *Smoker) LifeOfSmoker(c chan Sigarette, sem *Semaphore, sem_agent *Semaphore) {
	for {
		time.Sleep(time.Millisecond)
		// fmt.Println("LifeOfSmoker")
        s.Get(c, sem, sem_agent);
	}
}

func (s *Agent) LifeOfAgent(c chan Sigarette, sem *Semaphore, sem_agent *Semaphore) {
	for {
		time.Sleep(time.Millisecond)
		// fmt.Println("LifeOfAgent")
        s.Put(c, sem, sem_agent);
	}
}

func main() {
	c := make(chan Sigarette)
	paper := Smoker{1000, 0, 0}
	tobacco := Smoker{0, 1000, 0}
	matches := Smoker{0, 0, 1000}
	agent := Agent{1000, 1000, 1000}
	semaphore := Semaphore{0}
	semaphore_agent := Semaphore{0}

	
	go agent.LifeOfAgent(c, &semaphore, &semaphore_agent)

	go paper.LifeOfSmoker(c, &semaphore, &semaphore_agent)
	go tobacco.LifeOfSmoker(c, &semaphore, 	 &semaphore_agent)
	go matches.LifeOfSmoker(c, &semaphore, 	 &semaphore_agent)

	finished := make(chan bool)
	<- finished

	// semaphore := New(1, 1*time.Second)

	// fmt.Println("Начало")
	// tickets, timeout := 0, 0*time.Second
	// s := New(tickets, timeout)
	

	// if err := s.Acquire(); err != nil {
	// 	if err != ErrNoTickets {
	// 		panic(err)
	// 	}

	// 	fmt.Println("Билетов не осталось, не могу работать")
	// 	os.Exit(1)
	// }
	// fmt.Println("Завершение")

    // go greet(c)

    // c <- "John"
    //fmt.Println("main() stopped")
}
