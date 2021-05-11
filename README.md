# Installation

`$ go install`

## Usage 

`$ hyper-go path-to-config.yml`

### Configuration

```yaml
settings:
  default_modifier: Mod3

keys:
  - key: c
    type: run-or-raise # Tries to find window base on WM_CLASS Instance on raises first match, otherwise command is executed
    class: google-chrome
    command: google-chrome
  - key: d
    type: type # Types value of text into active window
    text: $
  - key: h
    type: sequence # Sends key sequence in value to active window
    value: Ctrl-Shift-Tab
  - key: r
    type: command # Executes command
    command: rofi -selected-row 0 -combi-modi window,run,drun,ssh -modi combi -show combi    
```