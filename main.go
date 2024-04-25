package main

import (
	"log"
	tea "github.com/charmbracelet/bubbletea"
)


func main(){
    m := NewModel()
    p := tea.NewProgram(m)

    _,err := p.Run()
    if err != nil{
        log.Fatalln(err)
    }
}

type Model struct{
    title string
}

func NewModel() Model{
    return Model{
        title: "Hello World",
    }
}


func (m Model) Init() tea.Cmd{
    return nil
}


func (m Model) Update(msg tea.Msg) (tea.Model,tea.Cmd){
    return m,nil
}

func (m Model) View() string{
    return m.title
}


