package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type startGoroutineFn func(
	done <-chan any,
	pulseInterval time.Duration,
) (heartbeat <-chan any)

func doWorkFn(done <-chan any, intList ...int) (startGoroutineFn, <-chan any) {
	intChanCh := make(chan (<-chan any))
	intCh := bridge(done, intChanCh)
	doWork := func(done <-chan any, pulseInterval time.Duration) <-chan any {
		intCh := make(chan any)
		heartBeat := make(chan any)
		go func() {
			defer close(intCh)
			select {
			case intChanCh <- intCh:
			case <-done:
				return
			}

			pulse := time.Tick(pulseInterval)

			for {
			valuLoop:
				for _, intVal := range intList {
					if intVal < 0 {
						log.Printf("negative value: %v\n", intVal)
						return
					}

					for {
						select {
						case <-pulse:
							select {
							case heartBeat <- struct{}{}:
							default:
							}
						case intCh <- intVal:
							continue valuLoop
						case <-done:
							return
						}
					}
				}
			}
		}()
		return heartBeat
	}
	return doWork, intCh
}

func bridge(done <-chan any, chanCh <-chan <-chan any) <-chan any {
	valCh := make(chan any)
	go func() {
		defer close(valCh)
		for {
			var ch <-chan any
			select {
			case mabyCh, ok := <-chanCh:
				if !ok {
					return
				}
				ch = mabyCh
			case <-done:
				return
			}

			for val := range orDone(done, ch) {
				select {
				case valCh <- val:
				case <-done:
				}
			}
		}
	}()

	return valCh
}

func orDone(done, valCh <-chan any) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case v, ok := <-valCh:
				if !ok {
					return
				}
				select {
				case ch <- v:
				case <-done:
				}
			}
		}
	}()
	return ch
}

func or(channels ...<-chan any) <-chan any {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan any)
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func newSteward(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn {
	return func(
		done <-chan any,
		pulseInterval time.Duration,
	) <-chan any {
		heartbeat := make(chan any)
		go func() {
			defer close(heartbeat)

			var wardDone chan any
			var wardHeartbeat <-chan any
			startWard := func() {
				wardDone = make(chan any)
				wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2)
			}
			startWard()
			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)

				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()
		return heartbeat
	}
}

func take(done <-chan any, valCh <-chan any, num int) <-chan any {
	ch := make(chan any)
	go func() {
		defer close(ch)
		for range num {
			select {
			case <-done:
				return
			case ch <- <-valCh:
			}
		}
	}()
	return ch
}

func main() {
	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)

	done := make(chan any)
	defer close(done)

	doWork, intCh := doWorkFn(done, 1, 2, -1, 3, 4, 5)
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork)
	doWorkWithSteward(done, 1*time.Hour)

	for intVal := range take(done, intCh, 6) {
		fmt.Printf("Received: %v\n", intVal)
	}

}
