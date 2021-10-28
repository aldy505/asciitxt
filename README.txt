     _      ____     ____   ___   ___   _____  __  __  _____ 
    / \    / ___|   / ___| |_ _| |_ _| |_   _| \ \/ / |_   _|
   / _ \   \___ \  | |      | |   | |    | |    \  /    | |  
  / ___ \   ___) | | |___   | |   | |    | |    /  \    | |  
 /_/   \_\ |____/   \____| |___| |___|   |_|   /_/\_\   |_|  
                                                             

Just a simple utility for creating...

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
      output := asciitxt.New("Hello world")
      
      // or

      output = asciitxt.WithConfig("Hello world", asciitxt.Config{
        Style: asciitxt.StyleStandard,
      })
    }

What's the asciitxt.StyleStandard, you asked.

Well, I thought it would be nice if we could have more than one style.
But, for now that's a long term plan.

My current goal is to support most unicode letters and signs.

See Go Package documentation for more details: https://pkg.go.dev/github.com/aldy505/asciitxt

Licensed under MIT License.
See the LICENSE file.
