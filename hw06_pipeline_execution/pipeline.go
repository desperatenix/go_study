package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		bi := make(Bi)

		go func(b Bi, o Out) {
			defer close(b)

			for {
				select {
				case <-done:
					return
				case v, ok := <-o:
					if !ok {
						return
					}
					b <- v
				}
			}
		}(bi, out)

		out = stage(bi)
	}
	return out
}
