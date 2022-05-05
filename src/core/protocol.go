package core

type Instruct string

const (
        Run Instruct = "Run"
        Invalid Instruct = "Invalid"
        Kill Instruct = "Kill"
        Skip Instruct = "Skip"
        ForceValid Instruct = "Ready"
)


type Respond string


