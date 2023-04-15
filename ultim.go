package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Response is a struct to handle the API response
type Response struct {
	Commit struct {
		Commit struct {
			Author struct {
				Name  string `json:"name"`
				Date  string `json:"date"`
				Sha   string `json:"sha"`
			} `json:"author"`
		} `json:"commit"`
	} `json:"commit"`
}

func main() {
	token := os.Args[1]
    author := os.Args[2]
    repo := os.Args[3]
    branch := os.Args[4]

	for {
		//token := "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"
		//author := "aratan"
		//repo := "test"
		apiURL := "https://api.github.com/repos"
		//branch := "dev"

		cmd := exec.Command("curl", "-s", "-H", fmt.Sprintf("Authorization: token %s", token), fmt.Sprintf("%s/%s/%s/branches/%s", apiURL, author, repo, branch))
		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		

		var response Response
		err = json.Unmarshal(output, &response)
		if err != nil {
			log.Fatal(err)
		}
		

		commit := fmt.Sprintf("%s, %s, %s", response.Commit.Commit.Author.Name, response.Commit.Commit.Author.Date, response.Commit.Commit.Author.Sha)
		fmt.Println(commit)

		lastCommitFile := "archivo.txt"
		content, err := ioutil.ReadFile(lastCommitFile)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Fatal(err)
			}
		}

		fileContent := strings.TrimSpace(string(content))
		commit = strings.TrimSpace(commit)

		if commit == fileContent {
			fmt.Println("\033[32mSIN CAMBIOS\033[0m")

			// limpia
				rm := exec.Command("rm", "-rf", repo)
				rmout, err := rm.Output()
				if err != nil {
					fmt.Println("Error:", err)
					
				}
				fmt.Println(string(rmout))
		} else {
			fmt.Println("\033[31m* CAMBIOS *\033[0m")
			err = ioutil.WriteFile(lastCommitFile, []byte(commit), 0644)
			if err != nil {
				log.Fatal(err)
			}


			// comprobar si ecxxister el repo
			

				// Aquí se debe colocar la lógica para clonar el repositorio, construir y desplegar el proyecto.
				
				cloneURL := fmt.Sprintf("https://github.com/%s/%s.git", author, repo)

				// Ejecutar el comando git clone
				git := exec.Command("git", "clone", cloneURL)
				gitout, err := git.Output()
				if err != nil {
					fmt.Println("Error:", err)
					
				}
				fmt.Println(string(gitout))

			// Definir el comando que se va a ejecutar
			cmd := exec.Command("/bin/bash", "-c", "./gitrun.sh")

			// Ejecutar el comando y capturar la salida
			out, err := cmd.Output()

			// Verificar si ocurrió un error al ejecutar el comando
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Imprimir la salida del comando en la consola
			fmt.Println(string(out))


			//	docker build -t nombre_de_la_imagen /ruta/a/mi_carpeta
			// Ejecutar el comando git clone
			docker := exec.Command("sudo","docker", "build", "-t",repo,"./%s",repo)
			dockerout, err := docker.Output()
			if err != nil {
				fmt.Println("Error:", err)
				
			}
			fmt.Println(string(dockerout))

			//}

			// Imprimir la salida del comando en la consola
			//fmt.Println(string(out))
		}
		
		// descomentar despues de pruebas
		//time.Sleep(1 * time.Hour)
		time.Sleep(2 * time.Minute)
			//fmt.Println(output)
	}
}
