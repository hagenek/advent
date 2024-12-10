const demoData = `XMMM
XAMM
SAMM`;

function part1(input: string) {
  const lines = input.split("\n");

  const matrix = lines.map((line) => line.split(""));

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    const lineArray = line.split("");
    for (let j = 0; j < lineArray.length; j++) {
      const char = lineArray[j];
      if (char === "X") {
        console.log("Found X at", i, j);
      }
    }
  }
}

part1(demoData);
