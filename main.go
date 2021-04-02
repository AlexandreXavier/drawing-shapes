package main

import "github.com/fogleman/ln/ln"

func main() {
	// criar a sena
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	// defenir os parametros da camera
	eye := ln.Vector{4, 3, 2}    // posisão camera
	center := ln.Vector{0, 0, 0} // objectivo da camera
	up := ln.Vector{0, 0, 1}     // direcão

	// defenir os parametros do render
	width := 1024.0  // width
	height := 1024.0 // height
	fovy := 50.0     // degrees
	znear := 0.1     // near z plane
	zfar := 10.0     // far z plane
	step := 0.01

	// caminhos 2D
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	// render para uma imagem
	paths.WriteToPNG("out.png", width, height)

	// gravar o caminho em SVG
	paths.WriteToSVG("out.svg", width, height)
}
