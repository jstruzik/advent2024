import { promises as fs } from 'fs';

async function solve(filePath: string): Promise<Array<Array<String>>> {
  try {
    const data = await fs.readFile(filePath, 'utf-8');
    const parsedArray = new Array();
    const regex = /(\d+)\s+(\d+)/gm
    const matches = data.matchAll(regex);
    const list1 = new Array();
    const list2 = new Array();
    for (const match of matches) {
        list1.push(Number(match[1]));
        list2.push(Number(match[2]));
    }
    const locationMap = new Map<number, number>();
    list1.forEach((value, index) => {
        if (!locationMap.has(value)) {
            locationMap.set(value, 0);
        }
    });
    list1.sort();
    list2.sort();
    parsedArray.push(list1, list2);

    let distance = 0;
    for (let i = 0; i < parsedArray[0].length; i++) {
        const localDist = Math.abs(parsedArray[0][i] - parsedArray[1][i]);
        distance = distance + localDist;

        if (locationMap.has(parsedArray[1][i])) {
            const val = Number(locationMap.get(parsedArray[1][i]));
            locationMap.set(parsedArray[1][i], val + 1);
        }
    }
    
    console.log('Distance: ' + distance);

    let simScore = 0;
    locationMap.forEach((value, index) => {
        simScore += value * index;
    });

    console.log('Similarity: ' + simScore);

    return parsedArray;
  } catch (err) {
    console.error('Error reading file:', err);
    throw err;
  }
}

solve('./resources/1.txt');;
