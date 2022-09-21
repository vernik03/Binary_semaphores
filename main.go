package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not aquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

// Interface содержит поведение семафора, который может быть захвачен (Acquire) и/или освобожден (Release).
type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

struct Smoker_paper {
	int paper = 1000;
	int tobacco = 0;
	int matches = 0;
}

struct Smoker_tobacco {
	int paper = 0;
	int tobacco = 1000;
	int matches = 0;
}

struct Smoker_matches {
	int paper = 0;
	int tobacco = 0;
	int matches = 1000;
}

func main() {
	fmt.Println("Начало")
	tickets, timeout := 0, 0*time.Second
	s := New(tickets, timeout)
	

	if err := s.Acquire(); err != nil {
		if err != ErrNoTickets {
			panic(err)
		}

		fmt.Println("Билетов не осталось, не могу работать")
		os.Exit(1)
	}
	fmt.Println("Завершение")
}
