go-alignment-block-estimate

- a alignment block estimates.
- define the start and the stop of the alignment regions and it computes all the stats from the given block region
- it will only process that block of the alignment and not the complete alignment. 

```
[gauravsablok@ultramarine]~/Desktop/codecreatede/go-alignment-block-estimate% \
go run main.go -h
This estimates the site proportion in your whole genome or gene specific alignment

Usage:
  flags [flags]

Flags:
  -a, --alignmentfile string   a alignment file (default "align")
  -e, --endcoordinate int      end of the alignment block (default 40)
  -h, --help                   help for flags
  -s, --startcoordinate int    start of the alignment block (default 1)
[gauravsablok@ultramarine]~/Desktop/codecreatede/go-alignment-block-estimate% \
go run main.go -a ./samplefile/samplealignment.fasta -s 1 -e 10
The collinearity block for A as a base pattern and T as a mismatch is %d 1
The collinearity block for A as a base pattern G as a mismatch is %d 1
The collinearity block for A as a base pattern and C as a mismatch is %d 0
The collinearity block for T as a base pattern and G as a mismatch is %d 0
The collinearity block for T as a base pattern C as a mismatch is  %d 0
The collinearity block for T as a base pattern and A as a mismatch is %d 0
The collinearity block for G as a base pattern and C as a mismatch is %d 0
The collinearity block for G as a base pattern A as a mismatch is  %d 0
The collinearity block for G as a base pattern and T as a mismatch is %d 0
The collinearity block for C as a base pattern and A as a mismatch is %d 0
The collinearity block for C as a base pattern T as a mismatch is  %d 0
The collinearity block for C as a base pattern and G as a mismatch is %d 0
The number of the unified bases across A are %d 0
The number of the unified bases across T are %d 0
The number of the unified bases across G are %d 0
The number of the unified bases across C are %d 1
```

Gaurav Sablok 

