package main

import (
	"log"
	"os"
	"time"
)

type startGoroutineFn func(
	done <-chan any,
	pulseInterval time.Duration,
) (heartbeat <-chan any)

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
	return func(done <-chan any, pulseInterval time.Duration) <-chan any {
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

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	doWork := func(done <-chan any, _ time.Duration) <-chan any {
		log.Println("ward: Hello, I'm irresponsible!")
		go func() {
			<-done
			log.Println("ward: I a halting.")
		}()
		return nil
	}
	doWorkwithSteward := newSteward(4*time.Second, doWork)

	done := make(chan any)
	time.AfterFunc(9*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for range doWorkwithSteward(done, 4*time.Second) {
	}
	log.Println("Done")
}
