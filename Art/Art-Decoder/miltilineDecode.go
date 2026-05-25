package main
import ( 
		"fmt" 
		"bufio"
		"os"
		)

func multilineDecode() {
	scanner := bufio.NewScanner(os.Stdin)
		results := []string{}

		for scanner.Scan() {
			line := scanner.Text()

			art, err := readArtCode( line)
		
			if err != nil {
				erroInfo := Red+ Bold + "Error" + Reset
				results = append(results,erroInfo)
				continue
			}
			
			results = append(results,art)
			
		}
		fmt.Println()
		for _, result := range results {
			fmt.Println(result)
		}
		return
}