/**
 * A simple codegen to generate the array of array letter
 */
const fs = require('fs/promises');

// If you used this codegen, don't forget to reset
// the letters array just to be a single value.
const letters = 
[
`        
   __ _ 
  / _\` |
 | (_| |
  \\__,_|
        `,
]

;(async () => {
  let output = ''
  for (const letter of letters) {
    let split = letter.split('\n');
    output += `case "":\n`
    output += `return []string{\n`
    for (const s of split) {
      output += '`'+s+'`,\n'
    }
    output += "}\n";
  }
  await fs.writeFile('./../generated.go', output, { encoding: 'utf-8' });
})();