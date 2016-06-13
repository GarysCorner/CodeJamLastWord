//File:		codejamlastword.go
//Author:	Gary Bezet
//Date:		2016-06-13
//Desc:		This program is designed to solve Google Code Jam "The Last Word"  The problem turned out to be a little too easy, yawn
//Problem:	https://code.google.com/codejam/contest/4304486/dashboard

package main

import (
		"time"
		"fmt"
		"os"
		"flag"
		"bufio"
		"strings"
		"strconv"
	)
	
//global variables
var infileopt, outfileopt string  //input and output filenames
var infile, outfile *os.File  //input and output file pointers
var totalcases int  //number of cases


var testcases []testcase


//structures
type testcase struct {
	casenum int  //case number
	letters []byte  //letters given
	lastword []byte //solution
	proctime time.Duration  //time it took to solve this problem

}


//program entry point
func main() {

	starttime := time.Now()  //start time for stats

	defer infile.Close()
	defer outfile.Close()

	initflags()  //initialize the command line args
	
	openFiles() //open the files
	
	processFile()  //process input file
	
	proctime := time.Now()
	for i := 0; i<totalcases; i++ {  //process all cases
		testcases[i].solve()
		printErrln( "Solved case", testcases[i].casenum, "in", testcases[i].proctime," Ans=", string(testcases[i].lastword) )
	}
	
	printErrln(totalcases, "solved in", time.Now().Sub(proctime), "sending solutions to", outfileopt)
	
	for _, v := range testcases { //print solutions
		fmt.Fprintf(outfile, "Case #%d: %s\n", v.casenum, v.lastword)
	}
	
	printErrln("FINISHED!  Elapsed: ", time.Now().Sub(starttime))
	
}


//get the flags from command line
func initflags() {
	flag.StringVar(&infileopt, "if", "", "Input file (required)")
	flag.StringVar(&outfileopt, "of", "-", "Output file, defaults to stdout" )

	flag.Parse()

	if infileopt == "" {
		printErrln("You must supply an input file\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	

}

//print error to console
func printErrln( line ...interface{} ) {
	fmt.Fprintln( os.Stderr, line... )
}


func openFiles() {
	
	var err error
	
	infile, err = os.Open(infileopt)

	if err != nil {
		printErrln( "Error:  Could not open:  ", infileopt)
		printErrln( "\tError: ", err  )
		os.Exit(2)
	}

	if outfileopt == "-"  {
		outfile = os.Stdout
		outfileopt = "Stdout"
	} else {
		outfile, err = os.Create(outfileopt)

		if err != nil {
			printErrln( "Error:  Could not create:  ", outfileopt)
			printErrln( "\tError: ", err  )
			os.Exit(3)
		} 
	}

	printErrln("InFile:\t", infileopt)
	printErrln("OutFile:\t", outfileopt, "\n")
		
}


func processFile() {  //process the input file into data structure
	proctime := time.Now() //for time to load data

	var err error
	var line string
	
	reader := bufio.NewReader(infile)  //new reader for file
	
	line, err = reader.ReadString('\n')
	if err != nil {
		printErrln("Couldn't read first line from:  ", infileopt)
		printErrln("\tError:  ", err )
		os.Exit(4)
		
	}
	
	totalcases, err = strconv.Atoi( strings.TrimSpace( line ) )
	if err != nil {
		printErrln("Couldn't process first line from:  ", infileopt)
		printErrln("\tError:  ", err )
		os.Exit(4)
		
	}
	
	testcases = make([]testcase, totalcases)  //initialise testcases[]
	for i := 0; i < totalcases; i++ {
		testcases[i].casenum = i + 1  //set casenum
		
		line, err = reader.ReadString('\n')
		if err != nil {
			printErrln("Couldn't read case", i+1, "from", infileopt)
			printErrln("\tError:  ", err)
			os.Exit(5)
		}
		
		testcases[i].letters = []byte(strings.TrimSpace(line))  //load testcase into byte array

		testcases[i].lastword = make([]byte, len(testcases[i].letters))  //allocate solution array
		
		testcases[i].lastword[0] = testcases[i].letters[0]  //time answer with first letter

		
	}
	
	printErrln("File", infileopt,"processed in", time.Now().Sub(proctime))		
}

func (self *testcase) solve() {
	
	starttime := time.Now()  //start the time processing solution started
	
	for i := 1; i < len(self.letters); i++ {  //process all characters
		if self.letters[i] < self.lastword[0] {
			//printErrln("Append", string(self.letters[i]), "to", string(self.lastword))
			self.lastword[i] = self.letters[i]
			
		} else {
			//printErrln("Prepend", string(self.letters[i]), "to", string(self.lastword))
			self.prepend(self.letters[i], i)
			
		}
	}
	
	self.proctime = time.Now().Sub(starttime)  //store the time to process solution

}

func (self *testcase) prepend(letter byte, place int) {
	for i:=place; i>0 ; i-- {  //start at back
		self.lastword[i] = self.lastword[i-1]
	}
	self.lastword[0] = letter
}

