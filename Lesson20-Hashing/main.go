package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint64 {
	var retval uint64
	h := fnv.New64a()
	h.Write([]byte(s))
	fmt.Printf("Type = %T \t Value = %v \n", h, h)
	retval = h.Sum64()
	fmt.Printf("Return Value = %v \n", retval)
	fmt.Printf("Size = %v \t BlockSize = %v \n", h.Size(), h.BlockSize())
	h.Reset()
	fmt.Printf("h.Reset() \n")
	retval = h.Sum64()
	fmt.Printf("Return Value = %v \n", retval)

	h.Write([]byte(s))
	fmt.Printf("Type = %T \t Value = %v \n", h, h)
	retval = h.Sum64()
	fmt.Printf("Return Value = %v \n", retval)
	fmt.Printf("Size = %v \t BlockSize = %v \n", h.Size(), h.BlockSize())

	return retval
}

func main() {

	s := `
	I’m in charge of a cage. I know those that won’t.
	I don’t mean can’t. Just won’t. There’s a roster
	for Tuesdays, Fridays. Dogs to die.
	
	The disconsolate, the abandoned, those with recurrent
	symptoms, the incorrigible mutt — oh, a dozen
	choices by way of reasons. Even so,
	
	some won’t. Won’t play along once their number’s
	up. The “rainbow bridge” in the offing
	as the posher clinics put it, a pig’s ear
	
	as a final treat, a venison chew, the profession
	behaving beautifully at a time like this.
	Still, those that won’t. Won’t go nicely, I mean,
	
	with a gaze to melt, a last slobbed lick.
	Those with a soul’s defiance, though embarrassment
	in the lunchroom should you come at that one!
	
	Even after the bag is zipped, you feel it:
	We’re real at the end as you are, buster. We sniff
	the wind. What say if we say it together? Won’t.
	
	Source: Poetry (February 2018)

	`

	fmt.Println(hash(s))
}
