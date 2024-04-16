package main

/* Determine Data Structure */

type xyz struct {
	id string;
	title string;
	description string;
	number int;
	someBoolean boolean;
}

var items = []xyz{
	{id: "Unique Identifier #1", title: "Title 1", description: "Here is some more information about title.", number: 1, someBoolean: true},
	{id: "mov_inception", title: "Inception", description: "This film is about drugs and draeming.", number: 56, someBoolean: false},
	{id: "tv_rickandmorty", title: "Rick & Morty", description: "I turned myself into a pickle, Morty! I'm Pickle Riiiiiiick!", number: 6786789, someBoolean: false},
	{id: "docu_totaltrust", title: "Total Trust", description: "A documentary about surveillance and censorship in China.", number: 562, someBoolean: true},
	{id: "soft_vscode", title: "Visual Studio Code", description: "IDE.", number: 3589, someBoolean: false},
	{id: "web_google", title: "Google", description: "A well-known search engine.", number: 99, someBoolean: true}
}