package main

import "fmt"

type Curso struct {
	titulo string
	videos [] Video
}

type Video struct {
	titulo string
	curso Curso
}

func main() {
	video1 := Video{titulo: "01-Intro"}
	video2 := Video{titulo: "02-Master"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: [] Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(curso)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
