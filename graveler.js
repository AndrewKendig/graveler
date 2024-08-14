// Completion Time for 1000000: 0 hours, 1 minutes, 14 seconds

var startTime = Date.now();
var rollcount = 1000000;
var highest = 0;
for (var i = 0; i < rollcount; i++) {
	var counts = [0,0,0,0]
	for (var k = 0; k < 231; k++) {
		counts[Math.floor(Math.random() * 4)] += 1;
	}
	for (var j = 0; j < counts.length; j++) {
		if (counts[j] > highest) {
			highest = counts[j];
		}
	}
}
var endTime = Date.now();

var durationMs = endTime - startTime;
var durationSeconds = Math.floor(durationMs / 1000);
var hours = Math.floor(durationSeconds / 3600);
var minutes = Math.floor((durationSeconds % 3600) / 60);
var seconds = durationSeconds % 60;

console.log(`Duration: ${hours} hours, ${minutes} minutes, ${seconds} seconds`);
console.log(`Highest count of roll: ${highest}`);
console.log(`Roll Count: ${rollcount}`)