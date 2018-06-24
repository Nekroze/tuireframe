# Purpose

The goal was to provide a way to easily mock up TUI screens for application design. To this end `tuireframe` will take in a description of a screen to mock up and render it in the terminal with full color support.

# Installation

You will need to install [Go][1] to execute the install command:

```
 $ go get github.com/Nekroze/tuireframe/...
```

This should make the `tuireframe` command available to your terminal.

# Usage

The `tuireframe` command takes in two files, a screen file and a meta file. And example of both of these, [example.screen](example.screen) and [example.meta](example.meta), can be found in the source [repository][2].

```
 $ tuireframe example.screen example.meta
```

or if you want to get fancy with bash.

```
 $ tuireframe example.{screen,meta}
```

To exit the rendering press the `q` key.

# Screen Files

A screen file is essentially a plain text file that holds all the characters, glyphs, etc. and their positions and is edited by simply typing them into the places you want them to render at.

This should look like a greyscale version of the final product. This text file is your canvas.

# Meta Files

A meta file is a json document that has a sequence of instructions to decorate the rendering of the screen be modifying cells. Cells have a foreground or `fg` field and a background or `bg` field that contain a color string, the cell can also override the character being displayed with the `char` field.

Instructions describe a point or rectangular area to which a specific cell modification will apply and are executed in sequance facilitating layering. The modifications need not change all of a cell by defining only the values you wish to change.

[1]: https://golang.org
[2]: https://github.com/Nekroze/tuireframe
