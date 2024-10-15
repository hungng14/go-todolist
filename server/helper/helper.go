package helper

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func InputUserName() string {
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	return userName
}

type Note struct {
	Name        string
	IsCompleted bool
}

var wg = sync.WaitGroup{}

func InputNotes() []Note {
	var notes = []Note{}
	fmt.Println("How many notes do you want to add?")
	fmt.Print("Enter your number of notes: ")
	var totalNote uint
	fmt.Scan(&totalNote)

	reader := bufio.NewReader(os.Stdin) // Using bufio.NewReader to read full lines

	var countFilledNote uint = 1
	if totalNote > 0 {
		for countFilledNote <= totalNote {

			fmt.Printf("(%v): ", countFilledNote)
			var newNote Note

			noteName, _ := reader.ReadString('\n') // Read entire line including spaces
			noteName = noteName[:len(noteName)-1]  // Remove the newline character
			newNote.Name = noteName
			newNote.IsCompleted = false

			notes = append(notes, newNote)

			fmt.Println()
			countFilledNote += 1

			wg.Add(1)
			go MakeNoteToComplete(&newNote)
		}
	}

	wg.Wait()

	return notes
}

func MakeNoteToComplete(note *Note) {
	time.Sleep(2 * time.Second)

	note.IsCompleted = true

	fmt.Printf("Note name %v moved to completed!\n", note.Name)

	wg.Done()
}
