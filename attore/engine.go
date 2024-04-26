package attore

import "log"

type Receiver interface {
	Receive(any)
}

type actor struct {
	r       Receiver
	mailbox chan any
}

func newActor(r Receiver) *actor {
	return &actor{
		r:       r,
		mailbox: make(chan any, 100),
	}
}

func (a *actor) spawn() {
	for {
		msg := <-a.mailbox
		a.r.Receive(msg)
	}
}

type Engine struct {
	actors map[string]*actor
}

func NewEngine() *Engine {
	return &Engine{
		actors: map[string]*actor{},
	}
}

func (e *Engine) Send(pid string, msg any) {
	actor, ok := e.actors[pid]
	if !ok {
		log.Printf("actor with pid [%s] does not exist", pid)
	}

	actor.mailbox <- msg
}

func (e *Engine) Spawn(r Receiver, name string) string {
	log.Println("spawning receiver", name)

	a := newActor(r)
	if _, ok := e.actors[name]; ok {
		panic("actor already in existance")
	}

	e.actors[name] = a
	go a.spawn()

	return name
}
