import fs from 'fs';
import readline from 'readline';

async function solve(filePath: string): Promise<void> {
  try {
    let result: number = 0;
    const fileStream = fs.createReadStream(filePath);
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity,
    });

    lineLoop:
    for await (const line of rl) {
        const numbers = line.split(' ').map(Number);
        let isPositive = true;
        let dampenerUsed = false;
        numberLoop:
        for (let i = 0; i < numbers.length - 1; i++) {
            const diff = numbers[i + 1] - numbers[i];
            const absDiff = Math.abs(diff);
            if (i === 0) {
                isPositive = diff > 0;
            }
            // Skip any disqualifiying lines
            if (absDiff < 1 || absDiff > 3 || (isPositive === true && diff < 0) || isPositive === false && diff > 0) {
              if (dampenerUsed) {
                continue lineLoop;
              }  
              dampenerUsed = true;
            }
        }
        result++;
    }

    console.log('Number of safe results: ' + result);
  } catch (err) {
    console.error('Error reading file:', err);
    throw err;
  }
}

solve('./resources/2.txt');;
