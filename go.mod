module project.localhost/example

replace project.localhost/purescript-native/ffi-loader => /home/alexa/Workspace/PursGame/purescript-native

replace project.localhost/purescript-native/output => /home/alexa/Workspace/PursGame/output

go 1.13

require (
	github.com/purescript-native/go-runtime v0.0.0-20190907045917-ec626efcf4a1
	github.com/veandco/go-sdl2 v0.3.3
	project.localhost/purescript-native/ffi-loader v0.0.0-00010101000000-000000000000 // indirect
	project.localhost/purescript-native/output v0.0.0-00010101000000-000000000000 // indirect
)
