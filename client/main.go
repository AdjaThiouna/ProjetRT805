package main

import (
    "fmt"
    "io/ioutil"
    
    "path/filepath"
)
func main(){
	directory:="./donnees"
	files, err := ioutil.ReadDir(directory)
    if err != nil {
        fmt.Println("Erreur en lisant le répertoire:", err)
        return
    }
	for _, file := range files {
        // Vérifiez si le fichier est un fichier JSON
        if filepath.Ext(file.Name()) == ".json" {
            // Lire le contenu du fichier
            filePath := filepath.Join(directory, file.Name())
            content, err := ioutil.ReadFile(filePath)
            if err != nil {
                fmt.Println("Erreur en lisant le fichier:", err)
                continue
            }
            // Afficher le contenu du fichier JSON
            fmt.Println("Contenu du fichier", file.Name(), ":", string(content))
        }
    }
}