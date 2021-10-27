# ASCIITXT

```
Just a simple utilities for creating...

  _____   _____  __  __  _____ 
 |_   _| | ____| \ \/ / |_   _|
   | |   |  _|    \  /    | |  
   | |   | |___   /  \    | |  
   |_|   |_____| /_/\_\   |_|  
                               
  _       ___   _  __  _____   
 | |     |_ _| | |/ / | ____|  
 | |      | |  | ' /  |  _|    
 | |___   | |  | . \  | |___   
 |_____| |___| |_|\_\ |_____|  
                               
  _____   _   _   ___   ____   
 |_   _| | | | | |_ _| / ___|  
   | |   | |_| |  | |  \___ \  
   | |   |  _  |  | |   ___) | 
   |_|   |_| |_| |___| |____/  
                               
The usage is pretty straight forward, on your Go file:

import "github.com/aldy505/asciitxt"

func main() {
  output := asciitxt.New("Hello world", asciitxt.Standard)
  
  // or

  output = asciitxt.WithConfig("Hello world", &asciitxt.Config{
    Style: asciitxt.Standard,
  })
}

What's the asciitxt.Standard, you asked.

Well, I thought it would be nice if we could have more than one style.
But, for now that's a long term plan.

My current goal is to support most unicode letters and signs.

Licensed under MIT License. See the LICENSE file.
```