/**
 * @author William
 * @version 1.0.0
 * @date 2025-11-27
 * @fileoverview This program calculates the sum of even numbers up to a positive integer.
 */

// variables
let n: number = 0;
let sum: number = 0;

// ask for positive integer
do {
  n = parseInt(prompt("Enter a positive integer: ") || "0");

  if (n <= 0) {
    console.log("Invalid input. Please enter a positive integer:");
  }
} while (n <= 0);

// calculate sum of even numbers
for (let counter = 1; counter <= n; counter++) {
  if (counter % 2 === 0) {
    sum = sum + counter;
  }
}

console.log(`Sum of even numbers from 1 to ${n} is: ${sum}`);

console.log("Done.");
