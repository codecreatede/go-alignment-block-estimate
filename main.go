package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-23

a phylogenomics alignment block analyzer that extract the defined block with in the position
specified and then perform the computational analysis on the defined block such as the number
of the conserved bases and the unconserved bases and the similarity and the dissimilarity index.

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	alignment string
	start     int
	end       int
)

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This estimates the site proportion in your whole genome or gene specific alignment",
	Run:  flagsFunc,
}

func init() {
	rootCmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
	rootCmd.Flags().IntVarP(&start, "startcoordinate", "s", 1, "start of the alignment block")
	rootCmd.Flags().IntVarP(&end, "endcoordinate", "e", 40, "end of the alignment block")
}

func flagsFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignBlock struct {
		id  string
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	sequenceID := []string{}
	alignmentBlock := []alignBlock{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
		if strings.HasPrefix(string(line), ">") {
			sequenceID = append(sequenceID, string(line))
		}
	}

	for i := 0; i < len(sequenceID); i++ {
		alignmentBlock = append(alignmentBlock, alignBlock{
			id:  string((sequenceID[i])),
			seq: string((sequenceCap[i][start:end])),
		})
	}
	alignmentSelID := []string{}
	alignmentSelSeq := []string{}

	for i := range alignmentBlock {
		alignmentSelID = append(alignmentSelID, alignmentBlock[i].id)
	}
	for j := range alignmentBlock {
		alignmentSelSeq = append(alignmentSelSeq, alignmentBlock[j].seq)
	}

	counterA := 0
	counterT := 0
	counterG := 0
	counterC := 0
	for i := 0; i < len(alignmentSelSeq); i++ {
		for j := 0; j < len(alignmentSelSeq[0]); j++ {
			if string(alignmentSelSeq[i][j]) == "A" && string(alignmentSelSeq[i+1][j]) == "A" {
				counterA++
			}
			if string(alignmentSelSeq[i][j]) == "T" && string(alignmentSelSeq[i+1][j]) == "T" {
				counterT++
			}
			if string(alignmentSelSeq[i][j]) == "G" && string(alignmentSelSeq[i+1][j]) == "G" {
				counterG++
			}
			if string(alignmentSelSeq[i][j]) == "C" && string(alignmentSelSeq[i+1][j]) == "C" {
				counterC++
			}
		}
	}

	counterAT := 0
	counterAG := 0
	counterAC := 0

	for i := 0; i < len(alignmentSelSeq)-1; i++ {
		for j := 0; j < len(alignmentSelSeq[0]); j++ {
			if string(alignmentSelSeq[i][j]) == "A" && string(alignmentSelSeq[i+1][j]) == "T" {
				counterAT++
			}
			if string(alignmentSelSeq[i][j]) == "A" && string(alignmentSelSeq[i+1][j]) == "C" {
				counterAG++
			}
			if string(alignmentSelSeq[i][j]) == "A" && string(alignmentSelSeq[i+1][j]) == "G" {
				counterAC++
			}
		}
	}

	counterTG := 0
	counterTC := 0
	counterTA := 0

	for i := 0; i < len(alignmentSelSeq)-1; i++ {
		for j := 0; j < len(alignmentSelSeq[0]); j++ {
			if string(alignmentSelSeq[i][j]) == "T" && string(alignmentSelSeq[i+1][j]) == "G" {
				counterTA++
			}
			if string(alignmentSelSeq[i][j]) == "T" && string(alignmentSelSeq[i+1][j]) == "C" {
				counterTC++
			}
			if string(alignmentSelSeq[i][j]) == "T" && string(alignmentSelSeq[i+1][j]) == "A" {
				counterTA++
			}
		}
	}

	counterGC := 0
	counterGA := 0
	counterGT := 0

	for i := 0; i < len(alignmentSelSeq)-1; i++ {
		for j := 0; j < len(alignmentSelSeq[0]); j++ {
			if string(alignmentSelSeq[i][j]) == "G" && string(alignmentSelSeq[i+1][j]) == "C" {
				counterGC++
			}
			if string(alignmentSelSeq[i][j]) == "G" && string(alignmentSelSeq[i+1][j]) == "A" {
				counterGA++
			}
			if string(alignmentSelSeq[i][j]) == "G" && string(alignmentSelSeq[i+1][j]) == "T" {
				counterGT++
			}
		}
	}

	counterCA := 0
	counterCT := 0
	counterCG := 0

	for i := 0; i < len(alignmentSelSeq)-1; i++ {
		for j := 0; j < len(alignmentSelSeq[0]); j++ {
			if string(alignmentSelSeq[i][j]) == "C" && string(alignmentSelSeq[i+1][j]) == "A" {
				counterCA++
			}
			if string(alignmentSelSeq[i][j]) == "C" && string(alignmentSelSeq[i+1][j]) == "T" {
				counterCT++
			}
			if string(alignmentSelSeq[i][j]) == "C" && string(alignmentSelSeq[i+1][j]) == "G" {
				counterCG++
			}
		}
	}

	fmt.Println(
		"The collinearity block for A as a base pattern and T as a mismatch is %d",
		counterAT,
	)
	fmt.Println("The collinearity block for A as a base pattern G as a mismatch is %d", counterAG)
	fmt.Println(
		"The collinearity block for A as a base pattern and C as a mismatch is %d",
		counterAC,
	)
	fmt.Println(
		"The collinearity block for T as a base pattern and G as a mismatch is %d",
		counterTG,
	)
	fmt.Println("The collinearity block for T as a base pattern C as a mismatch is  %d", counterTC)
	fmt.Println(
		"The collinearity block for T as a base pattern and A as a mismatch is %d",
		counterTA,
	)
	fmt.Println(
		"The collinearity block for G as a base pattern and C as a mismatch is %d",
		counterGC,
	)
	fmt.Println("The collinearity block for G as a base pattern A as a mismatch is  %d", counterGA)
	fmt.Println(
		"The collinearity block for G as a base pattern and T as a mismatch is %d",
		counterGT,
	)
	fmt.Println(
		"The collinearity block for C as a base pattern and A as a mismatch is %d\n",
		counterCA,
	)
	fmt.Println(
		"The collinearity block for C as a base pattern T as a mismatch is  %d\n",
		counterCT,
	)
	fmt.Println(
		"The collinearity block for C as a base pattern and G as a mismatch is %d\n",
		counterCG,
	)
	fmt.Println("The number of the unified bases across A are %d", counterA)
	fmt.Println("The number of the unified bases across T are %d", counterT)
	fmt.Println("The number of the unified bases across G are %d", counterG)
	fmt.Println("The number of the unified bases across C are %d", counterC)
}
