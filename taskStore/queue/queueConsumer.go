package queue

type Consumer interface {
	ReadFromQueue()
}
