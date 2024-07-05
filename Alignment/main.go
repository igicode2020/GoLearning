package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sequence alignment!")
	SARSALlignement()
}

func Hemoglobin() {
	zebrafish := ReadFASTAFile("Data/Hemoglobin/Danio_rerio_hemoglobin.fasta")
	cow := ReadFASTAFile("Data/Hemoglobin/Bos_taurus_hemoglobin.fasta")

	gorilla := ReadFASTAFile("Data/Hemoglobin/Gorilla_gorilla_hemoglovin.fasta")

	human := ReadFASTAFile("Data/Hemoglovin/Homo_sapiens_hemoglovin.fasta")
	
	match := 1.0
	mismatch := 1.0
	gap := 3.0
}

func SARSALlignement string{
	sars := ReadFASTAFile("Data/Cornaviruses/SARS-CoV_genome.fasta")

	sars2 := ReadFASTAFile("Data/Cornaviruses/SARS-CoV-2_genome.fasta")

	fmt.println("Geonomes read")
	match := 1.0
	mismatch := 1.0
	gap := 3.0
	GlobalAlignment(sars1, sars2, match, mismatch, gap)
	fmt.Println("Global alignment run. Printing file now")

	WriteAlignmentToFASTA(sarsAlignment, "output/sars_genome_alignmeent.txt")
}
