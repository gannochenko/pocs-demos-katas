const bubbleSort = <T>(arr: T[]): T[] => {
	for (let i = 0; i < arr.length; i++) {
		for (let j = i; j < arr.length; j++) {
			if (arr[i] > arr[j]) {
				const tmp = arr[i];
				arr[i] = arr[j];
				arr[j] = tmp;
			}
		}
	}

	return arr;
};

function main() {
	const arr = [4,1,7,5,6];

	const sortedArr = bubbleSort(arr);
	console.log(sortedArr);
}

main();
